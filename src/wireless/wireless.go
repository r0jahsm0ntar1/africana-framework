package wireless

import (
    "os"
    "fmt"
    "utils"
    "bufio"
    "strings"
    "bcolors"
    "subprocess"
)

var userInput, userIface, userSsid string

func AirGeddon() {
    subprocess.Popen(`airgeddon`)
}

func WifiPumpkin3() {
    filePath := "/root/.config/wifipumpkin3/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`wifipumpkin3`); utils.ClearScreen()
    }
    subprocess.Popen(`wifipumpkin3`)
}

func WifiteAuto() {
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Println()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Wireless card to use?" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "wlan0" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🍍" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userIface, _ := reader.ReadString('\n')
    userIface = strings.TrimSpace(userIface)
    if userIface == "" {
        userIface = "wlan0"
    }
    subprocess.Popen(`wifite -i %s --ignore-locks --keep-ivs -p 1339 -mac --random-mac -v -inf --bully --pmkid --dic /usr/share/wordlists/rockyou.txt --require-fakeauth --nodeauth --pmkid-timeout 120`, userIface)
}

func BettercapAuto() {
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Println()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Wireless card to use?" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "wlan0" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🍍" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userIface, _ := reader.ReadString('\n')
    userIface = strings.TrimSpace(userIface)
    if userIface == "" {
        userIface = "wlan0"
    }
    subprocess.Popen(`airmon-ng check kill; service NetworkManager restart; ip link set %s down; iw dev %s set type monitor; ip link set %s up; iw %s set txpower fixed 3737373737373; service NetworkManager start; sleep 3; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; wifi.recon on; wifi.show; set wifi.show.sort clients desc; set ticker.commands clear; wifi.show; wifi.assoc all; wifi.assoc all wifi.handshakes.file /usr/local/opt/handshakes; wifi.deauth all"`, userIface, userIface, userIface, userIface, userIface)
}

func WifiPumpkin3Auto() {
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Println()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Wireless card to use?" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "wlan0" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🍍" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userIface, _ := reader.ReadString('\n')
    userIface = strings.TrimSpace(userIface)
    if userIface == "" {
        userIface = "wlan0"
    }
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Name to set your wifi?" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "Jesus is the answer" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🍍" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userSsid, _ := reader.ReadString('\n')
    userSsid = strings.TrimSpace(userSsid)
    if userSsid == "" {
       userSsid = "Jesus is the answer"
    }
    filePath := "/root/.config/wifipumpkin3/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`wifipumpkin3; clear`)
    }
    subprocess.Popen(`wifipumpkin3 --xpulp "set interface %s; set ssid '%s'; set proxy noproxy; start"`, userIface, userSsid)
}
