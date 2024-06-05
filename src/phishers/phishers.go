package phishers

import (
    "subprocess"
)


func GoPhish() {
    subprocess.Popen(`gophish`)
}

func GoodGinx() {
    subprocess.Popen(`evilginx2`)
}

func ZPhisher() {
    subprocess.Popen(`cd /root/.africana/africana-base/AdvPhishing; bash AdvPhishing.sh`)
}

func SetoolKit() {
    subprocess.Popen(`cd /root/.africana/africana-base/set/; python3 setoolkit`)
}

func AnonPhisher() {
    subprocess.Popen(`cd /root/.africana/africana-base/anonphisher; bash anonphisher.sh`)
}

func CyberPhish() {
    subprocess.Popen(`cd /root/.africana/africana-base/cyberphish; python3 cyberphish.py`)
}

