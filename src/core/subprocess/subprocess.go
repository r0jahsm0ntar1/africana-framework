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
    "bcolors"
)

var (
    shell      string
    flag       string
    logFile    *os.File
    mu         sync.Mutex
    logDir     = "/root/.africana/logs"
    initialDir string
    currentDir string
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

    if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
        fmt.Fprintln(os.Stderr, bcolors.RED+"[!] "+bcolors.ENDC+"Error creating log directory:", err)
        return
    }
    openLogFile()
}

// Open or create a log file for storing command history
func openLogFile() {
    mu.Lock()
    defer mu.Unlock()

    if logFile != nil {
        logFile.Close()
    }

    var err error
    logFile, err = os.OpenFile(filepath.Join(logDir, "command_history.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Fprintln(os.Stderr, bcolors.RED+"[!] "+bcolors.ENDC+"Error opening log file: ", err)
    }
}

// Execute a command using the system shell
func Popen(command string, args ...interface{}) {
    process := fmt.Sprintf(command, args...)

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

// Execute a command fully, handling process execution and signals
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
        fmt.Fprintln(os.Stderr, bcolors.RED+"[!] "+bcolors.ENDC+"Error starting command:", err)
        return
    }

    if err := cmd.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, bcolors.YELLOW+"[!] "+bcolors.ENDC+"Process incomplete:", err)
    }

    signal.Stop(sigs)
    close(sigs)
}

// Change the current working directory
func changeDirectory(newDir string) {
    if newDir == "" || newDir == "~" {
        homeDir, err := os.UserHomeDir()
        if err != nil {
            fmt.Println(bcolors.RED+"[!] "+bcolors.ENDC+"Failed to get home directory:", err)
            return
        }
        newDir = homeDir
    } else if strings.HasPrefix(newDir, "~/") {
        homeDir, err := os.UserHomeDir()
        if err != nil {
            fmt.Println(bcolors.RED+"[!] "+bcolors.ENDC+"Failed to get home directory:", err)
            return
        }
        newDir = filepath.Join(homeDir, newDir[2:])
    }

    if !filepath.IsAbs(newDir) {
        newDir = filepath.Join(currentDir, newDir)
    }

    if stat, err := os.Stat(newDir); err == nil && stat.IsDir() {
        currentDir = newDir
        fmt.Println(bcolors.GREEN+"[*] "+bcolors.ENDC+"Changed directory to:", currentDir)
    } else {
        fmt.Println(bcolors.RED+"[!] "+bcolors.ENDC+"Invalid directory:", newDir)
    }
}

// Log executed commands to a file
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

// Close the log file to prevent memory leaks
func CloseLogFile() {
    mu.Lock()
    defer mu.Unlock()

    if logFile != nil {
        _ = logFile.Close()
    }
}

// Print command history from the log file
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
    fmt.Println(bcolors.BLUE+"[*] "+bcolors.ENDC+"exec: history\n")
    lineNumber := 1
    for scanner.Scan() {
        fmt.Printf("%d  %s\n", lineNumber, scanner.Text())
        lineNumber++
    }
}

// Clear the command history log file
func ClearHistory() {
    logFilePath := filepath.Join(logDir, "command_history.log")
    err := os.Remove(logFilePath)
    if err != nil {
        fmt.Println(bcolors.RED+"[!] "+bcolors.ENDC+"Error clearing history:", err)
    } else {
        fmt.Println(bcolors.BLUE+"[*] "+bcolors.ENDC+"exec: clear history\n\n"+bcolors.GREEN+"[*] "+bcolors.ENDC+"History cleared successfully.")
        openLogFile()
    }
}
