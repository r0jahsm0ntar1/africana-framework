package crackers

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "bcolors"
    "subprocess"
)

var(
    userInput       string
    userTarget      string
    userScript      string
    userOutput      string
    userPcap        string
    userHashes      string
    userWordPass    string
    userWordlists   string
    reader = bufio.NewReader(os.Stdin)
)

//Online crackers
func HydraSsh() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHcracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single User or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userName, _ := reader.ReadString('\n')
        userName = strings.TrimSpace(userName)
        if userName == "" {
            userName = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SSHcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.africana/output/HydraSsh_outfile.txt -u ssh://%s`, userName, userWordPass, userTarget); fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userNames, _ := reader.ReadString('\n')
        userNames = strings.TrimSpace(userNames)
        if userNames == "" {
            userNames = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SSHcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SSHcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.africana/output/HydraSsh_outfile.txt -u ssh://%s`, userNames, userWordPass, userTarget); fmt.Println()
        return
    }
}

func HydraFtp() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPcracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single User or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userName, _ := reader.ReadString('\n')
        userName = strings.TrimSpace(userName)
        if userName == "" {
            userName = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPcracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "FTPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.africana/output/HydraFtp_outfile.txt -u ftp://%s`, userName, userWordPass, userTarget); fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userNames, _ := reader.ReadString('\n')
        userNames = strings.TrimSpace(userNames)
        if userNames == "" {
            userNames = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "FTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "FTPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.africana/output/HydraFtp_outfile.txt -u ftp://%s`, userNames, userWordPass, userTarget); fmt.Println()
        return
    }
}

func HydraSmb() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single User or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userName, _ := reader.ReadString('\n')
        userName = strings.TrimSpace(userName)
        if userName == "" {
            userName = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SMBcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.africana/output/HydraSmb_outfile.txt -u smb://%s`, userName, userWordPass, userTarget); fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userNames, _ := reader.ReadString('\n')
        userNames = strings.TrimSpace(userNames)
        if userNames == "" {
            userNames = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMBCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SMBCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.africana/output/HydraSmb_outfile.txt -u smb://%s`, userNames, userWordPass, userTarget); fmt.Println()
        return
    }
}

func HydraRdp() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single User or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userName, _ := reader.ReadString('\n')
        userName = strings.TrimSpace(userName)
        if userName == "" {
            userName = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "RDPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.africana/output/HydraRdp_outfile.txt -u rdp://%s`, userName, userWordPass, userTarget); fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userNames, _ := reader.ReadString('\n')
        userNames = strings.TrimSpace(userNames)
        if userNames == "" {
            userNames = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "RDPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "RDPCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.africana/output/HydraRdp_outfile.txt -u rdp://%s`, userNames, userWordPass, userTarget); fmt.Println()
        return
    }
}

func HydraLdap() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single User or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userName, _ := reader.ReadString('\n')
        userName = strings.TrimSpace(userName)
        if userName == "" {
            userName = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "LDAPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.africana/output/HydraLdap_outfile.txt -u ldap://%s`, userName, userWordPass, userTarget); fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userNames, _ := reader.ReadString('\n')
        userNames = strings.TrimSpace(userNames)
        if userNames == "" {
            userNames = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "LDAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "LDAPCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.africana/output/HydraLdap_outfile.txt -u ldap://%s`, userNames, userWordPass, userTarget); fmt.Println()
        return
    }
}

func HydraSmtp() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single User or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userName, _ := reader.ReadString('\n')
        userName = strings.TrimSpace(userName)
        if userName == "" {
            userName = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SMTPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.africana/output/HydraSmtp_outfile.txt -u smtp://%s`, userName, userWordPass, userTarget); fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userNames, _ := reader.ReadString('\n')
        userNames = strings.TrimSpace(userNames)
        if userNames == "" {
            userNames = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SMTPCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.africana/output/HydraSmtp_outfile.txt -u smtp://%s`, userNames, userWordPass, userTarget); fmt.Println()
        return
    }
}

func HydraSnmtp() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single User or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userName, _ := reader.ReadString('\n')
        userName = strings.TrimSpace(userName)
        if userName == "" {
            userName = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SNMTPCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.africana/output/HydraSnmtp_outfile.txt -u snmtp://%s`, userName, userWordPass, userTarget); fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userNames, _ := reader.ReadString('\n')
        userNames = strings.TrimSpace(userNames)
        if userNames == "" {
            userNames = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "SNMTPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "SNMTPCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.africana/output/HydraSnmtp_outfile.txt -u snmtp://%s`, userNames, userWordPass, userTarget); fmt.Println()
        return
    }
}

func HydraTelnet() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single User or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userName, _ := reader.ReadString('\n')
        userName = strings.TrimSpace(userName)
        if userName == "" {
            userName = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "FTPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.africana/output/HydraTelnet_outfile.txt -u telnet://%s`, userName, userWordPass, userTarget); fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userNames, _ := reader.ReadString('\n')
        userNames = strings.TrimSpace(userNames)
        if userNames == "" {
            userNames = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "TELNETCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "TELNETCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.africana/output/HydraTelnet_outfile.txt -u telnet://%s`, userNames, userWordPass, userTarget); fmt.Println()
        return
    }
}

func HydraHttps() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "use " + bcolors.YELLOW + "1. " + bcolors.RED + "Single User or " + bcolors.YELLOW + "2. " + bcolors.RED + "Wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userInput)
    switch userInput {
    case "1":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "admin" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userName, _ := reader.ReadString('\n')
        userName = strings.TrimSpace(userName)
        if userName == "" {
            userName = "admin"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "FTPcracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -l %s -P %s -f -o /root/.africana/output/HydraHttps_outfile.txt -u https://%s`, userName, userWordPass, userTarget); fmt.Println()
        return
    case "2":
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "User Names to Crack " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userNames, _ := reader.ReadString('\n')
        userNames = strings.TrimSpace(userNames)
        if userNames == "" {
            userNames = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HTTPSCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userWordPass, _ := reader.ReadString('\n')
        userWordPass = strings.TrimSpace(userWordPass)
        if userWordPass == "" {
            userWordPass = "/usr/share/wordlists/rockyou.txt"
        }
        fmt.Printf(bcolors.ENDC + "\n[" + bcolors.DARKCYAN + "HTTPSCracker has began" + bcolors.ENDC + "] [" + bcolors.GREEN + "Target" + bcolors.ENDC + ": " + bcolors.YELLOW + "%s" + bcolors.ENDC + "]\n\n" + bcolors.ENDC, userTarget)
        subprocess.Popen(`hydra -t 4 -L %s -P %s -f -o /root/.africana/output/HydraHttps_outfile.txt -u http://%s`, userNames, userWordPass, userTarget); fmt.Println()
        return
    }
}

func CyberBrute() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "ALL/SCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "TARGET " + bcolors.PURPLE + "To attack" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    subprocess.Popen(`cd /root/.africana/africana-base/cyberbrute; bash cyberbrute %s`, userTarget); fmt.Println()
}

func HashBuster() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "HASHCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your hash " + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userHashes)
    subprocess.Popen(`cd /root/.africana/africana-base/hash-buster; python3 cracker.py -t 10 %s`, userHashes); fmt.Println()
}

//Offline Crackers
func AirCrackng() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "PCAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your .pcap " + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userPcap)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "PCAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    subprocess.Popen(`aircrack-ng %s -w %s`, userPcap, userWordlists); fmt.Println()
}

func JohnCrackng() {
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "PCAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Full path to your .pcap " + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userPcap)
    fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "PCAPCracker " + bcolors.ENDC + bcolors.ITALIC + "set " + bcolors.RED + "Path to Pass wordlist " + bcolors.PURPLE + "default " + bcolors.ENDC + "= " + bcolors.YELLOW + bcolors.ITALIC + "Rockyou.txt" + bcolors.ENDC + bcolors.BLUE + ")" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "\n╰─🐼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    subprocess.Popen(`john %s --wordlist=%s`, userPcap, userWordlists); fmt.Println()
}
