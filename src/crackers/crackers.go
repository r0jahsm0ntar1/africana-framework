package crackers

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "bcolors"
    "subprocess"
)

var userInput, userTarget, userScript, userOutput, userPcap, userHashes, userWordlists string

//Online crackers
func HydraSsh() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.RED + "Bruteforcing SSH password" + bcolors.BLUE + ")" + bcolors.BLUE + " -> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenThree(`hydra -L /usr/share/wordlists/rockyou.txt -P %s -f -o /root/.africana/output/Hydrassh_outfile.txt -u ssh://%s`, userWordlists, userTarget)
}

func HydraFtp() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.RED + "Bruteforcing FTP password" + bcolors.BLUE + ")" + bcolors.BLUE + " -> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenThree(`hydra -L /usr/share/wordlists/rockyou.txt -P %s -f -o /root/.africana/output/Hydraftp_outfile.txt -u ftp://%s`, userWordlists, userTarget)
}

func HydraSmb() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.RED + "Bruteforcing SMB password" + bcolors.BLUE + ")" + bcolors.BLUE + " -> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenThree(`hydra -L /usr/share/wordlists/rockyou.txt -P %s -f -o /root/.africana/output/Hydrasmb_outfile.txt -u smb://%s`, userWordlists, userTarget)
}

func HydraRdp() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.RED + "Bruteforcing RDP password" + bcolors.BLUE + ")" + bcolors.BLUE + " -> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenThree(`hydra -L /usr/share/wordlists/rockyou.txt -P %s -f -o /root/.africana/output/Hydrardp_outfile.txt -u rdp://%s`, userWordlists, userTarget)
}

func HydraLdap() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.RED + "Bruteforcing LDAP password" + bcolors.BLUE + ")" + bcolors.BLUE + " -> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenThree(`hydra -L /usr/share/wordlists/rockyou.txt -P %s -f -o /root/.africana/output/Hydraldap_outfile.txt -u ldap://%s`, userWordlists, userTarget)
}

func HydraSmtp() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.RED + "Bruteforcing SMTP password" + bcolors.BLUE + ")" + bcolors.BLUE + " -> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenThree(`hydra -L /usr/share/wordlists/rockyou.txt -P %s -f -o /root/.africana/output/Hydrasmtp_outfile.txt -u smtp://%s`, userWordlists, userTarget)
}

func HydraSnmtp() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.RED + "Bruteforcing SNMTP password" + bcolors.BLUE + ")" + bcolors.BLUE + " -> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenThree(`hydra -L /usr/share/wordlists/rockyou.txt -P %s -f -o /root/.africana/output/Hydrasnmtp_outfile.txt -u snmtp://%s`, userWordlists, userTarget)
}

func HydraTelnet() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.RED + "Bruteforcing TELNET password" + bcolors.BLUE + ")" + bcolors.BLUE + " -> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenThree(`hydra -L /usr/share/wordlists/rockyou.txt -P %s -f -o /root/.africana/output/Hydratelnet_outfile.txt -u telnet://%s`, userWordlists, userTarget)
}

func HydraHttps() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    fmt.Printf(bcolors.BLUE + "\n(" + bcolors.RED + "Bruteforcing HTTP/S password" + bcolors.BLUE + ")" + bcolors.BLUE + " -> " + bcolors.BLUE + "(" + bcolors.YELLOW + "%s", userTarget + bcolors.BLUE + ")\n" + bcolors.ENDC)
    subprocess.PopenThree(`hydra -L /usr/share/wordlists/rockyou.txt -P %s -f -o /root/.africana/output/Hydrahttps_outfile.txt -u https://%s`, userWordlists, userTarget)
}

func CyberBrute() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Target IP" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    subprocess.PopenTwo(`cd /root/.africana/africana-base/cyberbrute; bash cyberbrute %s`, userTarget)
}

func HashBuster() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.DARKCYAN + "Full path to your .pcap" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userHashes)
    subprocess.PopenTwo(`cd /root/.africana/africana-base/hash-buster; python3 cracker.py -t 10 %s`, userHashes)
}

func AirCrackng() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.DARKCYAN + "Full path to your .pcap" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userPcap)
    fmt.Printf(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.RED + "Script: " + bcolors.YELLOW + "%s", userPcap + bcolors.GREEN + ")\n" + bcolors.ENDC)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    subprocess.PopenThree(`aircrack-ng %s -w %s`, userPcap, userWordlists)
}

//Offline Crackers
func JohnCrackng() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.DARKCYAN + "Full path to your .pcap" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userPcap)
    fmt.Printf(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.RED + "Script: " + bcolors.YELLOW + "%s", userPcap + bcolors.GREEN + ")\n" + bcolors.ENDC)
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Path to wordlist" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.RED + "Default:" + bcolors.YELLOW + "Rockyou.txt" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userWordlists, _ := reader.ReadString('\n')
    userWordlists = strings.TrimSpace(userWordlists)
    if userWordlists == "" {
        userWordlists = "/usr/share/wordlists/rockyou.txt"
    }
    subprocess.PopenThree(`john %s --wordlist=%s`, userPcap, userWordlists)
}
