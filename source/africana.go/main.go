package main

import (
    "os"
    "fmt"
    "bufio"
    "os/exec"
    "strings"
    "bcolors"
    "internals/RunNmap"
)

func main() {
    screenClear(); printBanner()
    scanner := bufio.NewScanner(os.Stdin)
    for {
        printMenu()
        fmt.Print("Enter your choice: ")
        scanner.Scan()
        userInput := scanner.Text()

        switch strings.ToLower(userInput) {
        case "1":
            screenClear(); RunNmap()
        case "2":
            screenClear(); runDnsRecon()
        case "3":
            runBoth()
        case "0":
            screenClear(); exitBanner()
            return
        default:
            fmt.Println("Invalid choice. Please select 'nmap', 'dnsrecon', 'both', or 'exit'.")
        }
    }
}

func printBanner() {
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

func exitBanner() {
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

func printMenu() {
    fmt.Println(bcolors.RED + bcolors.BOLD +"        ~>[ Choose what to do from the menu below ]<~ \n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Port Scanning                 6. Start Nikto Scaning ]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Dns Reconning                 7. Fero File Searching ]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Web Technologies              8. Automation Scanning ]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Nikto Vuln Scanning           9. View all scann logs ]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Nuclei Vuln Scanning          0. Exit Africana Tool  ]\n" + bcolors.ENDC)
}

func screenClear() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("Error running creenClear:", err)
    }
}

func runBoth() {
    fmt.Println("Running both nmap and dnsrecon...")
    runNmap()
    runDnsRecon()
}
