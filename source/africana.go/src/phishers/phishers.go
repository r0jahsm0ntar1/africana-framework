package phishers

import (
    "fmt"
    "bcolors"
    "subprocess"
)


func GoPhish() {
    subprocess.Popen(`gophish`)
}

func GoodGinx() {
    subprocess.Popen(`evilginx2`)
}

func ZPhisher() {
    subprocess.Popen(`cd ~/.africana/externals/AdvPhishing; bash AdvPhishing.sh`)
}

func SetoolKit() {
    subprocess.Popen(`cd ~/.africana/externals/set/; python3 setoolkit`)
}

func AnonPhisher() {
    subprocess.Popen(`cd ~/.africana/externals/anonphisher; bash anonphisher.sh`)
}

func CyberPhish() {
    subprocess.Popen(`cd ~/.africana/externals/cyberphish; python3 cyberphish.py`)
}

func UpsenTools() {
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")" + bcolors.GREEN + "(" + bcolors.DARKCYAN + "Not yet implimented but comming soon" + bcolors.GREEN + ")" + bcolors.ENDC)
}