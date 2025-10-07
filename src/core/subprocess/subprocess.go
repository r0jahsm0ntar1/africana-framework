//John 3:16

package subprocess

import (
    "os"
    "fmt"
    "sync"
    "bufio"
    "os/exec"
    "os/user"
    "bcolors"
    "runtime"
    "strings"
    "os/signal"
    "path/filepath"
)

var (
    err error
    Version = "3"
    mu sync.Mutex
    logFile *os.File
    Dversion = Version + ".0.5-dev"
    logDir = filepath.Join(BaseDir, "logs")
    BaseDir, _, _, _, _, _, _, _ = DirLocations()
    LogFile = filepath.Join(logDir, "log_history.log")
    HistoryFile = filepath.Join(logDir, "command_history.log")
    flag, shell, process, baseDir, initialDir, currentDir string
)

type RunOptions struct {
    ReturnError bool
}

func init() {
    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/c"
    default:
        shell = "bash"
        flag = "-c"
    }

    dir, err := os.Getwd()
    if err == nil {
        initialDir = dir
        currentDir = dir
    }

}

func CheckRoot()bool {
    return os.Geteuid() == 0
}

func GetHomeDir() string {
    usr, err := user.Current()
    if err != nil {
        fmt.Printf("%s%s[!] %sError: getting current user ", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
        return ""
    }
    return usr.HomeDir
}

func creatLogDir() {
    if !CheckRoot(){
        homeDir := GetHomeDir()
        if homeDir == "" {
            return
        }
        logDir = filepath.Join(homeDir, "/.afr" + Version, "logs")
    }

    if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
        msg, _ := fmt.Printf("%s%s[!] %sError creating log directory ... %s%s_", bcolors.Bold, bcolors.Red, bcolors.Endc, bcolors.Endc, err)
        fmt.Fprintln(os.Stderr, msg)
        return
    }
}

func Run(command string, args ...interface{}) error {
    return runInternal(command, false, args...)
}

func RunWithError(command string, args ...interface{}) error {
    return runInternal(command, true, args...)
}

func runInternal(command string, returnError bool, args ...interface{}) error {
    creatLogDir()
    openLogFile()

    if len(args) > 0 {
        process = fmt.Sprintf(command, args...)
    } else {
        process = command
    }

    if process != "" {
        logCommand(process)
    }

    var executionErr error

    if strings.Contains(process, "&&") || strings.Contains(process, ";") {
        executionErr = executeFullCommandWithError(process)
    } else {
        parts := strings.Split(process, "&&")
        for i, part := range parts {
            part = strings.TrimSpace(part)
            if strings.HasPrefix(part, "cd ") {
                cdCommand := strings.TrimSpace(strings.TrimPrefix(part, "cd "))
                changeDirectory(cdCommand)
                if i == len(parts)-1 {
                    executionErr = nil
                    break
                }
            } else {
                executionErr = executeFullCommandWithError(part)
                if executionErr != nil {
                    break
                }
            }
        }
    }

    if executionErr != nil {
        if returnError {
            return executionErr
        } else {
            fmt.Printf("%s%s[!] %sFailed: %v\n", bcolors.Bold, bcolors.Red, bcolors.Endc, executionErr)
        }
    }
    
    return nil
}

func executeFullCommandWithError(command string) error {
    cmd := exec.Command(shell, flag, command)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Dir = currentDir

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt)

    go func() {
        <-sigs
        if cmd.Process != nil {
            _ = cmd.Process.Signal(os.Interrupt)
        }
    }()

    if err := cmd.Start(); err != nil {
        return fmt.Errorf("error starting command: %w", err)
    }

    if err := cmd.Wait(); err != nil {
        return fmt.Errorf("process incomplete: %w", err)
    }

    signal.Stop(sigs)
    close(sigs)
    return nil
}

func openLogFile() {
    mu.Lock()
    defer mu.Unlock()

    if logFile != nil {
        logFile.Close()
    }

    logFile, err = os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        msg, _ := fmt.Printf("%s%s[!] %sError: opening log file ... %s_", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
        fmt.Fprintln(os.Stderr, msg)
    }
}

func logCommand(command string) {
    mu.Lock()
    defer mu.Unlock()

    if logFile == nil {
        openLogFile()
    }

    if command != "" {
        _, _ = logFile.WriteString(command + "\n")
    }
}

func CloseLogFile() {
    mu.Lock()
    defer mu.Unlock()

    if logFile != nil {
        _ = logFile.Close()
    }
}

func changeDirectory(newDir string) {
    if newDir == "" || newDir == "~" {
        newDir = "/root"
    } else if strings.HasPrefix(newDir, "~/") {
        newDir = "/root" + newDir[1:]
    }

    if !filepath.IsAbs(newDir) {
        newDir = filepath.Join(currentDir, newDir)
    }

    if stat, err := os.Stat(newDir); err == nil && stat.IsDir() {
        currentDir = newDir
        fmt.Printf("%s%s[*] %sChanged directory to: %s%s%s\n", bcolors.Bold, bcolors.Green, bcolors.Endc, bcolors.Blue, currentDir, bcolors.Endc)
    } else {
        fmt.Printf("%s%s[!] %sInvalid directory: %s%s%s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, bcolors.Blue, newDir, bcolors.Endc)
    }
}

func ShowLogs() {
    if _, err := os.Stat(LogFile); os.IsNotExist(err) {
        return
    }

    file, err := os.Open(LogFile)
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    fmt.Printf("%s%s[*] %sexec: history\n\n", bcolors.Bold, bcolors.BrightBlue, bcolors.Endc)
    lineNumber := 1
    for scanner.Scan() {
        fmt.Printf("%s%d. %s%s\n", bcolors.BrightBlue, lineNumber, bcolors.Endc, scanner.Text())
        lineNumber++
    }
}

func ShowHistory() {
    if _, err := os.Stat(HistoryFile); os.IsNotExist(err) {
        return
    }

    file, err := os.Open(HistoryFile)
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    fmt.Printf("%s%s[*] %sexec: history\n\n", bcolors.Bold, bcolors.BrightBlue, bcolors.Endc)
    lineNumber := 1
    for scanner.Scan() {
        fmt.Printf("%s%d. %s%s\n", bcolors.BrightBlue, lineNumber, bcolors.Endc, scanner.Text())
        lineNumber++
    }
}

func ClearLogs() {
    err := os.Remove(LogFile)
    if err != nil {
        fmt.Printf("%s%s[!] %sError: clearing history:", bcolors.Bold, bcolors.BrightRed, bcolors.Endc, err)
    } else {
        fmt.Printf("%s%s[+] %sexec: clear history\n%s%s[*] %sHistory cleared ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc, bcolors.Bold, bcolors.BrightGreen, bcolors.Endc)
        openLogFile()
    }
}

func ClearHistory() {
    err := os.Remove(HistoryFile)
    if err != nil {
        fmt.Printf("%s%s[!] %sError: clearing history:", bcolors.Bold, bcolors.BrightRed, bcolors.Endc, err)
    } else {
        fmt.Printf("%s%s[+] %sexec: clear history\n%s%s[*] %sHistory cleared ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc, bcolors.Bold, bcolors.BrightGreen, bcolors.Endc)
        openLogFile()
    }
}

func DirLocations() (string, string, string, string, string, string, string, string) {
    if runtime.GOOS == "windows" {
        baseDir = filepath.Join(os.Getenv("ProgramData"), ".afr" + Version)
    } else {
        baseDir = filepath.Join("/root", ".afr" + Version)
    }

    if !CheckRoot(){
        if homeDir := GetHomeDir(); homeDir != "" {
            baseDir = filepath.Join(homeDir, ".afr" + Version)
        } else {
            return "", "", "", "", "", "", "", ""
        }
    }

    certDir := filepath.Join(baseDir, "certs")
    toolsDir := filepath.Join(baseDir, "africana-base")

    return baseDir, toolsDir, certDir, filepath.Join(certDir, "afr_cert.pem"), filepath.Join(certDir, "afr_key.pem"), filepath.Join(baseDir, "output"), filepath.Join(toolsDir, "wordlists"), "/usr/share/wordlists/rockyou.txt"
}

func ValidateToolsDir() error {
    _, toolsDir, _, _, _, _, _, _ := DirLocations()

    if _, err := os.Stat(toolsDir); os.IsNotExist(err) {
        return fmt.Errorf("'%s%s%s' doesn't exist!", bcolors.Bold, toolsDir, bcolors.Endc)
    }
    return nil
}
