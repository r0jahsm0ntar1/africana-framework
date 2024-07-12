package subprocess

import (
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "runtime"
    "syscall"
)

var (
    shell string
    flag  string
    sh    *exec.Cmd // Store the exec.Cmd object globally
)

func init() {
    // Initialize shell and flag based on runtime OS
    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/c"
    default:
        shell = "bash"
        flag = "-c"
    }
}

// Popen starts a subprocess with the given command and waits for termination.
func Popen(command string, args ...interface{}) {
    // Format command with provided arguments
    process := fmt.Sprintf(command, args...)

    // Create exec.Command object
    sh = exec.Command(shell, flag, process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    // Channel to capture interrupt signals
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

    // Start the command
    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    // Goroutine to listen for interrupt signals and forward them to the subprocess
    go func() {
        for sig := range sigs {
            if sh.Process != nil {
                sh.Process.Signal(sig)
            }
        }
    }()

    // Wait for the command to finish
    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }

    signal.Stop(sigs) // Stop capturing interrupt signals
    close(sigs)
}
