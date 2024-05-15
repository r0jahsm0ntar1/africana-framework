package subprocess

import (
    "fmt"
    "os"
    "os/exec"
)

func Popen(one string) {
    process := fmt.Sprintf(one)
    sh := exec.Command("bash", "-c", process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenCmd(one string) {
    process := fmt.Sprintf(one)
    cm := exec.Command("cmd", "/c", process)
    cm.Stdin = os.Stdin
    cm.Stdout = os.Stdout
    cm.Stderr = os.Stderr

    if err := cm.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := cm.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenTwo(one, two string) {
    process := fmt.Sprintf(one, two)
    sh := exec.Command("bash", "-c", process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr
    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }
    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenThree(one, two, three string) {
    process := fmt.Sprintf(one, two, three)
    sh := exec.Command("bash", "-c", process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenFour(one, two, three, four string) {
    process := fmt.Sprintf(one, two, three, four)
    sh := exec.Command("bash", "-c", process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenFive(one, two, three, four, five string) {
    process := fmt.Sprintf(one, two, three, four, five)
    sh := exec.Command("bash", "-c", process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenSix(one, two, three, four, five, six string) {
    process := fmt.Sprintf(one, two, three, four, five, six)
    sh := exec.Command("bash", "-c", process)
    sh.Stdin = os.Stdin
    sh.Stdout = os.Stdout
    sh.Stderr = os.Stderr

    if err := sh.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := sh.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}
