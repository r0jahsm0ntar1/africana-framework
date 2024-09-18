package subprocess

import (
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

    logDir := filepath.Join("/root/.africana", "logs")
    if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
        fmt.Fprintln(os.Stderr, "Error creating log directory:", err)
        return
    }

    var err error
    logFile, err = os.OpenFile(filepath.Join(logDir, "command_history.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error opening log file:", err)
    }
}

func Popen(command string, args ...interface{}) {
    process := fmt.Sprintf(command, args...)

    logCommand(process)

    cmd := exec.Command(shell, flag, process)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt)

    if err := cmd.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    go handleInterrupt(cmd, sigs)

    if err := cmd.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }

    signal.Stop(sigs)
    close(sigs)
}

func logCommand(command string) {
    mu.Lock()
    defer mu.Unlock()

    if logFile != nil {
        if _, err := logFile.WriteString(command + "\n"); err != nil {
            fmt.Fprintln(os.Stderr, "Error writing to log file:", err)
        }
    }
}

func handleInterrupt(cmd *exec.Cmd, sigs chan os.Signal) {
    <-sigs
    if cmd.Process != nil {
        if err := cmd.Process.Signal(os.Interrupt); err != nil {
            if !isProcessFinished(err) {
                fmt.Fprintln(os.Stderr, "Error sending interrupt signal:", err)
            }
        }
    }
}

func isProcessFinished(err error) bool {
    return err == os.ErrProcessDone
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
