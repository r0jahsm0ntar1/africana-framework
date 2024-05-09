package utils

import (
    "os"
    "net"
    "fmt"
    "subprocess"
)

func GetDefaultIP() (string, error) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }
    for _, iface := range interfaces {
        if iface.Flags&net.FlagUp == 0 {
            continue
        }
        if iface.Flags&net.FlagLoopback != 0 {
            continue
        }
        addrs, err := iface.Addrs()
        if err != nil {
            return "", err
        }
        for _, addr := range addrs {
            var ip net.IP
            switch v := addr.(type) {
            case *net.IPNet:
                ip = v.IP
            case *net.IPAddr:
                ip = v.IP
            }
            if ip == nil || ip.IsLoopback() {
                continue
            }
            ip = ip.To4()
            if ip == nil {
                continue
            }
            return ip.String(), nil
        }
    }
    return "", fmt.Errorf("no active network interface found")
}

func WordLists() {
    filePath := "/usr/share/wordlists/rockyou.txt"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`gunzip /usr/share/wordlists/rockyou.txt.gz`)
    }
}
