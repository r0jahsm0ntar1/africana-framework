package subprocess

import (
    "fmt"
    "os"
    "os/exec"
)

func Popen(one string) {
    process := fmt.Sprintf(one)
    cmd := exec.Command("bash", "-c", process)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := cmd.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenTwo(one, two string) {
    process := fmt.Sprintf(one, two)
    cmd := exec.Command("bash", "-c", process)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }
    if err := cmd.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenThree(one, two, three string) {
    process := fmt.Sprintf(one, two, three)
    cmd := exec.Command("bash", "-c", process)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := cmd.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenFour(one, two, three, four string) {
    process := fmt.Sprintf(one, two, three, four)
    cmd := exec.Command("bash", "-c", process)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := cmd.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenFive(one, two, three, four, five string) {
    process := fmt.Sprintf(one, two, three, four, five)
    cmd := exec.Command("bash", "-c", process)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := cmd.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}

func PopenSix(one, two, three, four, five, six string) {
    process := fmt.Sprintf(one, two, three, four, five, six)
    cmd := exec.Command("bash", "-c", process)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        fmt.Fprintln(os.Stderr, "Error starting process:", err)
        return
    }

    if err := cmd.Wait(); err != nil {
        fmt.Fprintln(os.Stderr, "process finished with error:", err)
    }
}
