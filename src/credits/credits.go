//John 3:16

package credits

import(
    "os"
    "fmt"
    "time"
    "bufio"
    "bcolors"
)

func readFileLetterByLetter(filename string, delay time.Duration) {
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        return
    }

    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("%s[!] %sError opening file: ", bcolors.BrightRed, bcolors.Endc, err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        for _, letter := range line {
            fmt.Print(string(letter))
            time.Sleep(delay)
        }
        fmt.Println()
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("%s[!] %sError reading file: ", bcolors.BrightRed, bcolors.Endc, err)
    }
}

func Contributors() {
    filename := "/root/.afr3/africana-base/hacktricks/roadmap.txt"
    delay := 0 * time.Millisecond
    readFileLetterByLetter(filename, delay)

    fmt.Printf(`%s%s%s
Torsocks & anonimity%s`, bcolors.Bold, bcolors.Underl, bcolors.BrightBlue, bcolors.Endc)
    fmt.Printf(bcolors.Endc + `
1. torsocks
` + bcolors.Endc)

    fmt.Printf(`
%s%s%sInternal network attack vectors%s`, bcolors.Bold, bcolors.Underl, bcolors.BrightBlue, bcolors.Endc)
   networks := bcolors.Endc + `
3. bettercap
4. nmap
5. metasploit
6. smbmap
7. rpcclient
8. beef-xss
9. responder
10 toxssin
` + bcolors.Endc
    for _, i := range networks {
        fmt.Print(string(i))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(`
%s%s%sBackdoor & c2 attack vectors%s`, bcolors.Bold, bcolors.Underl, bcolors.BrightBlue, bcolors.Endc)
   exploits := bcolors.Endc + `
1. Blackjack by t3l3machus
2. ShellzGen by 4ndr34z
3. PowerJoker
4. MeterPeter
5. Havoc C2 by @C5pider
6. Teardroid
7. AndroidRAT
8. Chameleon
9. Gh0x0st
` + bcolors.Endc
    for _, m := range exploits {
        fmt.Print(string(m))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(`
%s%s%sWireless networks attack vectors%s`, bcolors.Bold, bcolors.Underl, bcolors.BrightBlue, bcolors.Endc)
   wireless := bcolors.Endc + `
1. Wifite
2. Wifipumpkin3
3. Airgeddon
` + bcolors.Endc
    for _, w := range wireless {
        fmt.Print(string(w))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(`
%s%s%sCrackers for PassWords, Hash & .Pcap%s`, bcolors.Bold, bcolors.Underl, bcolors.BrightBlue, bcolors.Endc)
   crackers := bcolors.Endc + `
1. Aircrack_ng
2. John
3. Hash-Buster by Somdev Sangwan
` + bcolors.Endc
    for _, c := range crackers {
        fmt.Print(string(c))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(`
%s%s%sSocial-engineering attacks%s`, bcolors.Bold, bcolors.Underl, bcolors.BrightBlue, bcolors.Endc)
   phishers := bcolors.Endc + `
1. Gophish
2. Good Ginx
3. AdvPhishing
4. Setoolkit by David Kennedy
5. Anonphisher
6. Cyberphish
` + bcolors.Endc
    for _, p := range phishers {
        fmt.Print(string(p))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(`
%s%s%sWebsite attack vectors%s`, bcolors.Bold, bcolors.Underl, bcolors.BrightBlue, bcolors.Endc)
        webs := bcolors.Endc + `
1. Musker   A.I Proxies
2. wafw00f
3. Dnsrecon
4. Seekolver by Krypteria
5. whatweb
6. Harvester
7. Paramspider
8. Sslscan
9. Gobuster
10. Nuclei
11. Nikto
12. Bbot
13. Uniscan
14. Sqlmap
15. Commix
16. Katana
17. Xsser
18. Nettacker
19. Jok3r by koutto
20. Osmedeus by j3ssie
21. Ufonet by epsylon
` + bcolors.Endc
    for _, w := range webs {
        fmt.Print(string(w))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(`
%s%s%sCode of conduct%s`, bcolors.Bold, bcolors.Underl, bcolors.BrightBlue, bcolors.Endc)
    caution := bcolors.Endc + `
[ Africana-framework is written purely for Good & not evil]
[ The author of africana-framework is Rojahs Montari from.]
[ cyberafrics a cybersecurity organisation in Africa Kenya]
[ What is there 4 U 2 gain the whole world & loose your...]
[ soul? Be smart your Creator has good plans for you......]
[ The devil has no power over you Christian soldier.......]
[ Email....: rojahsmontari@gmail.com......................]
[ YouTube..: https://youtube.com/@RojahsMontari...........]
` + bcolors.Endc
    for _, c := range caution {
        fmt.Print(string(c))
        time.Sleep(3 * time.Millisecond)
    }
}

func Developer() {
    fmt.Printf(`
%s%s%sAbout the author%s`, bcolors.Bold, bcolors.Underl, bcolors.BrightBlue, bcolors.Endc)
    developer := `
[ I am Rojahs Montari a Devoted Christian & Pentester.....]
[ One might describe me as an erudite.....................]
[ & perspicacious individual, a connoisseur of cybernetic ]
[ security and digital fortification. My acumen in........]
[ devising the 🦠 africana-framework bespeaks a sagacious.]
[ approach to ethological hacking and vulnerability.......]
[ reconnaissance. As a virtuoso of cryptographic endeavors]
[ My pedagogical content disseminates esoteric knowledge,.]
[ fostering a coterie of aspirant ethical hackers. My.....]
[ prolific contributions to the technological milieu......]
[ underscore a quintessential commitment to advancing.....]
[ cybersecurity paradigms.................................]

[ Great thanks to my Family & my one & only loving wife ..]
[ Naomi Waceke for her love and emotional support.........]

`
    for _, d := range developer {
        fmt.Print(string(d))
        time.Sleep(3 * time.Millisecond)
    }
    fmt.Printf(`Life Tip.: %s%sDefeat the devil by fasting and praying ...%s
`, bcolors.BrightGreen, bcolors.Blink, bcolors.Endc)
}

func Creditors() {
    Contributors(); Developer()
}
