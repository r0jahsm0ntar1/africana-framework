package subprocess

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
)

var shell, flag string

func Popen(one string) {

    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/C"
    default:
        shell = "bash"
        flag = "-c"
    }

    process := fmt.Sprintf(one)
    sh := exec.Command(shell, flag, process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }
}

func PopenTwo(one, two string) {

    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/C"
    default:
        shell = "bash"
        flag = "-c"
    }

    process := fmt.Sprintf(one, two)
    sh := exec.Command(shell, flag, process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }
}

func PopenThree(one, two, three string) {

    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/C"
    default:
        shell = "bash"
        flag = "-c"
    }

    process := fmt.Sprintf(one, two, three)
    sh := exec.Command(shell, flag, process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }
}

func PopenFour(one, two, three, four string) {

    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/C"
    default:
        shell = "bash"
        flag = "-c"
    }

    process := fmt.Sprintf(one, two, three, four)
    sh := exec.Command(shell, flag, process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }
}

func PopenFive(one, two, three, four, five string) {

    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/C"
    default:
        shell = "bash"
        flag = "-c"
    }

    process := fmt.Sprintf(one, two, three, four, five)
    sh := exec.Command(shell, flag, process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }
}

func PopenSix(one, two, three, four, five, six string) {

    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/C"
    default:
        shell = "bash"
        flag = "-c"
    }

    process := fmt.Sprintf(one, two, three, four, five, six)
    sh := exec.Command(shell, flag, process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }
}

func PopenSeven(one, two, three, four, five, six, seven string) {

    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/C"
    default:
        shell = "bash"
        flag = "-c"
    }

    process := fmt.Sprintf(one, two, three, four, five, six, seven)
    sh := exec.Command(shell, flag, process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "Process finished with error:", err)
    }
}
