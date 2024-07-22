package phishers

import (
    "fmt"
    "subprocess"
)

func GoPhish() {
    subprocess.Popen(`gophish`); fmt.Println()
}

func GoodGinx() {
    subprocess.Popen(`evilginx2`); fmt.Println()
}

func ZPhisher() {
    subprocess.Popen(`cd /root/.africana/africana-base/zphisher; bash zphisher.sh`); fmt.Println()
}

func Darkphish() {
    subprocess.Popen(`cd /root/.africana/africana-base/darkphish; python3 dark-phish.py`); fmt.Println()
}

func SetoolKit() {
    subprocess.Popen(`cd /root/.africana/africana-base/set/; python3 setoolkit`); fmt.Println()
}

func AnonPhisher() {
    subprocess.Popen(`cd /root/.africana/africana-base/anonphisher; bash anonphisher.sh`); fmt.Println()
}

func CyberPhish() {
    subprocess.Popen(`cd /root/.africana/africana-base/cyberphish; python3 cyberphish.py`); fmt.Println()
}
