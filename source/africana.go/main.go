package main

import(
    "os"
    "fmt"
    "os/exec"
    "syscall"
    "bcolors"
)

func banner() {
    fmt.Println(bcolors.BLUE + `                     _,._                         ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `                 __.'   _)                        ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `                <_,)'.-'a\                        ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `                  /' (    \                       ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `      _.-----..,-'   (''--^                       ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `     //              |                            ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `    (|   ';      ,   |                            ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `      \   ;.----/  ,/                             ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `       ) // /   | |\ \                            ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `       \ \\'\   | |/ /     Jesus Christ           ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `        \ \\ \  | |\/  The Lamb that was slain    ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `         '" '"  '"'        for our sins.          ` + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + `        __                 _____ _____     _     _   ` + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + `     __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_ ` + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + `    |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|` + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + `    |_____|___|___|___|___|_____|__|__|_| |_|___|_|  ` + bcolors.ENDC)
    fmt.Println("                                                                                  \n")
}

func exit() {
    fmt.Print(`
                   _,.---.---.---.--.._ 
               _.-' '--.'---.'---'-. _,'--.._
              /'--._ .'.     '.     ','-.'-._\
             ||   \  '.'---.__'__..-'. ,''-._/
        _  ,'\ '-._\   \    '.    '_.-'-._,''-.
     ,'   '-_ \/ '-.'--.\    _\_.-'\__.-'-.'-._'.
    (_.o> ,--. '._/'--.-',--'  \_.-'       \'-._ \
     '---'    '._ '---._/__,----'           '-. '-\
               /_, ,  _..-'                    '-._\
               \_, \/ ._(
                \_, \/ ._\
                 '._,\/ ._\
                   '._// ./'-._
                     '-._-_-_.-'
    `)
    fmt.Println(bcolors.GREEN + "\n                  wake up, Christian             " + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + "           Lord God Jesus Christ L❤️.VE'S you      " + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + "                follow the white Pigeon.           " + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + "                   knock, knock, knock,            " + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + "                       Man Of God.               \n" + bcolors.ENDC)
}

func menu() {
    fmt.Println(bcolors.RED + bcolors.BOLD +"        ~>[ Choose what to do from the menu below ]<~ \n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Port Scanning                 6. Start Nikto Scaning ]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Dns Reconning                 7. Fero File Searching ]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Web Technologies              8. Automation Scanning ]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Nikto Vuln Scanning           9. View all scann logs ]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Nuclei Vuln Scanning          0. Exit Africana Tool  ]\n" + bcolors.ENDC)
}

func clear() {
    Cmd := exec.Command("clear")
    Out, err := Cmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println(string(Out))
    _, err = exec.Command("clear").Output()
    if err != nil {
        switch e := err.(type) {
        case *exec.Error:
            fmt.Println("failed executing:", err)
        case *exec.ExitError:
            fmt.Println("command exit rc =", e.ExitCode())
        default:
            panic(err)
        }
    }
}

func nmap() {
    binary, lookErr := exec.LookPath("nmap")
    if lookErr != nil {
        panic(lookErr)
    }
    var host string
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.RED + " type > host" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&host)
    args := []string{"nmap", "-p-", "-vv", host}
    env  := os.Environ()
    execErr:= syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}

func dnsr() {
    binary, lookErr := exec.LookPath("dnsrecon")
    if lookErr != nil {
        panic(lookErr)
    }
    var host string
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.RED + " type > host" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&host)
    args := []string{"dnsrecon", "-a", "-d", host}
    env  := os.Environ()
    execErr:= syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}

func main() {
    clear()
    banner()
    menu()
    var choice string
    fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&choice)
    switch choice {
    case "1":
        fmt.Print("\n"); nmap()
    case "2":
        fmt.Print("\n"); dnsr()
    case "0":
        clear(); exit()
        break
    default:
        fmt.Println(bcolors.ENDC + " ~{ " + bcolors.RED + "Poor choice of selection!. Please select int -> " + bcolors.DARKCYAN + " from 0. to 9. " + bcolors.ENDC +  "}~" + bcolors.ENDC)
    }
}
