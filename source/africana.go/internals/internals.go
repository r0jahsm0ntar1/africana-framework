package internals

import (
    "os"
    "fmt"
    "os/exec"
    "bcolors"
)

func RunNmap() {
    fmt.Println(bcolors.GREEN + "Running nmap..." + bcolors.ENDC)
    cmd := exec.Command("nmap", "-p-", "-v", "localhost")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
            fmt.Println("Error running nmap:", err)
    }
}

func RunDnsRecon() {
    fmt.Println(bcolors.GREEN + "Running dnsrecon..." + bcolors.ENDC)
    cmd := exec.Command("dnsrecon", "-d", "example.com")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
            fmt.Println("Error running dnsrecon:", err)
    }
}

func RunBoth() {
    fmt.Println("Running both nmap and dnsrecon...")
    RunNmap()
    RunDnsRecon()
}
