package phishers

import (
    "fmt"
    "time"
    "subprocess"
)

func GoPhish() {
    Logs := fmt.Sprintf("/root/.africana/logs/GoPhish.Log-%s.txt", time.Now().Format("20060102.150405"))
    subprocess.PopenTwo(`script -q -c 'gophish' -O %s`, Logs); fmt.Println()
}

func GoodGinx() {
    Logs := fmt.Sprintf("/root/.africana/logs/GoodGinx.Log-%s.txt", time.Now().Format("20060102.150405"))
    subprocess.PopenTwo(`script -q -c 'evilginx2' -O %s`, Logs); fmt.Println()
}

func ZPhisher() {
    Logs := fmt.Sprintf("/root/.africana/logs/ZPhisher.Log-%s.txt", time.Now().Format("20060102.150405"))
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/zphisher; bash zphisher.sh' -O %s`, Logs); fmt.Println()
}

func SetoolKit() {
    Logs := fmt.Sprintf("/root/.africana/logs/SetoolKit.Log-%s.txt", time.Now().Format("20060102.150405"))
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/set/; python3 setoolkit' -O %s`, Logs); fmt.Println()
}

func AnonPhisher() {
    Logs := fmt.Sprintf("/root/.africana/logs/AnonPhisher.Log-%s.txt", time.Now().Format("20060102.150405"))
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/anonphisher; bash anonphisher.sh' -O %s`, Logs); fmt.Println()
}

func CyberPhish() {
    Logs := fmt.Sprintf("/root/.africana/logs/CyberPhish.Log-%s.txt", time.Now().Format("20060102.150405"))
    subprocess.PopenTwo(`script -q -c 'cd /root/.africana/africana-base/cyberphish; python3 cyberphish.py' -O %s`, Logs); fmt.Println()
}
