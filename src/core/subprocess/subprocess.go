package subprocess

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "path/filepath"
    "runtime"
    "strings"
    "sync"
    "os/user"
    "bcolors"
)

var (
    err error
    mu sync.Mutex
    logFile *os.File
    logDir = "/root/.afr3/logs/"
    flag, shell, process, initialDir, currentDir string
)

// Initialize package variables and setup logging directory
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

// IsRoot checks if the current user is root
func IsRoot() bool {
    return os.Geteuid() == 0
}

// GetHomeDir returns the current user's home directory
func GetHomeDir() string {
    usr, err := user.Current()
    if err != nil {
        fmt.Printf("Error getting current user:", err)
        return ""
    }
    return usr.HomeDir
}

// createLogDir creates the log directory if it doesn't exist
func creatLogDir() {
    if !IsRoot() {
        homeDir := GetHomeDir()
        if homeDir == "" {
            return
        }
        logDir = homeDir + "/.afr3/logs/"
    }

    if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
        msg, _ := fmt.Printf("%s[!] %sError creating log directory ... %s%s_", bcolors.BrightRed, bcolors.Endc, bcolors.Endc, err)
        fmt.Fprintln(os.Stderr, msg)
        return
    }
}

// Popen executes a shell command and logs it
func Popen(command string, args ...interface{}) {
    creatLogDir(); openLogFile()
    if len(args) > 0 {
        process = fmt.Sprintf(command, args...)
    } else {
        process = command
    }

    if process != "" {
        logCommand(process)
    }

    if strings.Contains(process, "&&") || strings.Contains(process, ";") {
        executeFullCommand(process)
        return
    }

    parts := strings.Split(process, "&&")

    for i, part := range parts {
        part = strings.TrimSpace(part)
        if strings.HasPrefix(part, "cd ") {
            cdCommand := strings.TrimSpace(strings.TrimPrefix(part, "cd "))
            changeDirectory(cdCommand)
            if i == len(parts)-1 {
                return
            }
        } else {
            executeFullCommand(part)
        }
    }
}

// executeFullCommand executes a single shell command
func executeFullCommand(command string) {
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
        msg, _ := fmt.Printf("%s[!] %sError starting command ... %s%s_", bcolors.BrightRed, bcolors.BrightBlue, bcolors.Endc, err)
        fmt.Fprintln(os.Stderr, msg)
        return
    }

    if err := cmd.Wait(); err != nil {
        msg, _ := fmt.Printf("%s[!] %sProcess is incomplete ... %s%s_", bcolors.BrightRed, bcolors.Endc, bcolors.Endc, err)
        fmt.Fprintln(os.Stderr, msg)
    }

    signal.Stop(sigs)
    close(sigs)
}

// openLogFile opens the log file for writing
func openLogFile() {
    mu.Lock()
    defer mu.Unlock()

    if logFile != nil {
        logFile.Close()
    }

    logFile, err = os.OpenFile(filepath.Join(logDir, "command_history.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        msg, _ := fmt.Printf("%s[!] %sError opening log file ... %s%s_", bcolors.BrightRed, bcolors.Endc, bcolors.Endc, err)
        fmt.Fprintln(os.Stderr, msg)
    }
}

// logCommand logs a command to the log file
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

// CloseLogFile closes the log file
func CloseLogFile() {
    mu.Lock()
    defer mu.Unlock()

    if logFile != nil {
        _ = logFile.Close()
    }
}

// changeDirectory changes the current working directory
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
        fmt.Printf("%s[*] %sChanged directory to: %s%s%s\n", bcolors.Green, bcolors.Endc, bcolors.BrightBlue, currentDir, bcolors.Endc)
    } else {
        fmt.Printf("%s[!] %sInvalid directory: %s%s%s\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightBlue, newDir, bcolors.Endc)
    }
}

// LogHistory prints the command history
func LogHistory() {
    logFilePath := filepath.Join(logDir, "command_history.log")
    if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
        return
    }

    file, err := os.Open(logFilePath)
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    fmt.Printf("%s[*] %sexec: history\n\n", bcolors.BrightBlue, bcolors.Endc)
    lineNumber := 1
    for scanner.Scan() {
        fmt.Printf("%d  %s\n", lineNumber, scanner.Text())
        lineNumber++
    }
}

// ClearHistory clears the command history
func ClearHistory() {
    logFilePath := filepath.Join(logDir, "command_history.log")
    err := os.Remove(logFilePath)
    if err != nil {
        fmt.Printf("%s[!] %sError clearing history:", bcolors.BrightRed, bcolors.Endc, err)
    } else {
        fmt.Printf("%s[*] %sexec: clear history\n\n%s[*] %sHistory cleared ...\n", bcolors.BrightBlue, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        openLogFile()
    }
}
