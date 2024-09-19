package subprocess

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "path/filepath"
    "runtime"
    "sync"
)

var (
    shell   string
    flag    string
    logFile *os.File
    mu      sync.Mutex
    logDir  = "/root/.africana/logs"
)

func init() {
    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/c"
    default:
        shell = "bash"
        flag = "-c"
    }

    if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
        fmt.Fprintln(os.Stderr, "Error creating log directory:", err)
        return
    }
    openLogFile()
}

func openLogFile() {
    mu.Lock()
    defer mu.Unlock()

    if logFile != nil {
        logFile.Close()
    }

    var err error
    logFile, err = os.OpenFile(filepath.Join(logDir, "command_history.log"),
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error opening log file:", err)
    }
}

func Popen(command string, args ...interface{}) {
    process := fmt.Sprintf(command, args...)

    if process != "" {
        logCommand(process)
    }

    cmd := exec.Command(shell, flag, process)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt)

    go func() {
        <-sigs
        if cmd.Process != nil {
            if err := cmd.Process.Signal(os.Interrupt); err != nil && !isProcessFinished(err) {
                fmt.Fprintln(os.Stderr, "Error sending interrupt:", err)
            }
        }
    }()

    if err := cmd.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := cmd.Wait(); err != nil {
        //fmt.Printf(bcolors.RED + "[*] " + bcolors.ENDC + "Unknown command: %s. Run the %shelp %scommand for more details.\n", command, bcolors.DARKGREEN, bcolors.ENDC)
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }
    signal.Stop(sigs)
    close(sigs)
}

func isProcessFinished(err error) bool {
    return err == os.ErrProcessDone || err.Error() == "os: process already finished"
}

func logCommand(command string) {
    mu.Lock()
    defer mu.Unlock()

    if logFile == nil {
        openLogFile()
    }

    if command != "" {
        if _, err := logFile.WriteString(command + "\n"); err != nil {
            fmt.Fprintln(os.Stderr, "Error writing to log file:", err)
        }
    }
}

func CloseLogFile() {
    mu.Lock()
    defer mu.Unlock()

    if logFile != nil {
        if err := logFile.Close(); err != nil {
            fmt.Fprintln(os.Stderr, "Error closing log file:", err)
        }
    }
}

func History() {
    logFilePath := filepath.Join(logDir, "command_history.log")
    if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
        fmt.Println("History file does not exist.")
        return
    }

    file, err := os.Open(logFilePath)
    if err != nil {
        fmt.Println("Error opening history file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    fmt.Println("[*] exec: history")
    lineNumber := 1
    for scanner.Scan() {
        fmt.Printf("%d  %s\n", lineNumber, scanner.Text())
        lineNumber++
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading history:", err)
    }
}

func ClearHistory() {
    logFilePath := filepath.Join(logDir, "command_history.log")

    err := os.Remove(logFilePath)
    if err != nil {
        fmt.Println("Error clearing history:", err)
    } else {
        fmt.Println("Command history cleared.")
        openLogFile()
    }
}
