package subprocess

import (
    "os"
    "fmt"
    "time"
    "bufio"
    "bcolors"
    "runtime"
    "strings"
    "syscall"
    "os/exec"
    "os/signal"
    "path/filepath"
)

var (
    shell string
    flag  string
    sh    *exec.Cmd
    reader = bufio.NewReader(os.Stdin)
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
}

func Popen(command string, args ...interface{}) {
    process := fmt.Sprintf(command, args...)

    logDir := "/root/.africana/logs"
    if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
        fmt.Fprintln(os.Stderr, "Error creating log directory:", err)
        return
    }

    logFileName := fmt.Sprintf("session_%s.log", time.Now().Format("2006-01-02_15-04-05"))
    logFileTime := fmt.Sprintf("logtime_%s.tme", time.Now().Format("2006-01-02_15-04-05"))
    logTimePath := filepath.Join(logDir, logFileTime)
    logFilePath := filepath.Join(logDir, logFileName)

    var scriptCommand string
    if runtime.GOOS == "windows" {
        scriptCommand = process
    } else {
        escapeCommand := fmt.Sprintf("%q", process)
        scriptCommand = fmt.Sprintf("script -q -c %s --log-timing %s --log-out %s", escapeCommand, logTimePath, logFilePath)
    }

    sh = exec.Command(shell, flag, scriptCommand)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    go func() {
        for sig := range sigs {
            if sh.Process != nil {
                sh.Process.Signal(sig)
            }
        }
    }()

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }

    signal.Stop(sigs)
    close(sigs)
}

func InteractiveShell() {
    currentDir, err := os.Getwd()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error getting current directory:", err)
        return
    }

    fmt.Println(bcolors.ITALIC + "\n   🐧System " + bcolors.DARKGREEN + "'interactive' " + bcolors.RED + "'shell' " + bcolors.ENDC + bcolors.ITALIC + "Type exit to quit!" + bcolors.ENDC)

    for {
        fmt.Print(bcolors.BLUE + "\n╭─(" + bcolors.ORANGE + bcolors.ITALIC + currentDir + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Print(bcolors.BLUE + "\n╰─🐚" + bcolors.GREEN + "❯ " + bcolors.ENDC)

        command, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error reading command:", err)
            continue
        }

        command = strings.TrimSpace(command)

        if command == "exit" {
            break
        }

        if strings.HasPrefix(command, "cd ") {
            dir := strings.TrimSpace(strings.TrimPrefix(command, "cd "))
            if dir == "" {
                dir = os.Getenv("HOME")
            }
            if err := os.Chdir(dir); err != nil {
                fmt.Fprintln(os.Stderr, "Error changing directory:", err)
            } else {
                currentDir, _ = os.Getwd()
            }
            continue
        }

        sh = exec.Command(shell, flag, command)
        sh.Stdin = os.Stdin
        sh.Stdout = os.Stdout
        sh.Stderr = os.Stderr

        if err := sh.Start(); err != nil {
            fmt.Fprintln(os.Stderr, "Error starting process:", err)
            continue
        }

        if err := sh.Wait(); err != nil {
            fmt.Fprintln(os.Stderr, "Process finished with error:", err)
        }
    }
}

