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
        fmt.Println("Error opening file:", err)
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
        fmt.Println("Error reading file:", err)
    }
}

func Contributors() {
    filename := "/root/.africana/africana-base/cheetsheets/hacktricks.txt"
    delay := 0 * time.Millisecond
    readFileLetterByLetter(filename, delay)

    fmt.Printf(bcolors.BOLD + `
 Africana name        Developer's name        Original_name
 -------------        ----------------        -------------` + bcolors.ENDC)

    fmt.Printf(bcolors.ENDC + bcolors.BOLD + `
1. africana...........Rojahs Montari....africana-framework]
` + bcolors.ENDC)

    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + `
Tor system setup` + bcolors.ENDC)
    fmt.Printf(bcolors.ENDC + `
1. anonsurf...........Salim Zaved Karim............neutron]
` + bcolors.ENDC)

    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + `
Internal-Network-Attack` + bcolors.ENDC)

   internals := bcolors.ENDC + `
3. bettercap.....................................bettercap]
4. nmap...............................................nmap]
5. metasploit-framework.........................msfconsole]
6. smbmap...........................................smbmap]
7. rpcclient.....................................rpcclient]
8. beef-xss.......................................beef-xss]
9. responder.....................................responder]
10 toxssin.........................................toxssin]
` + bcolors.ENDC
    for _, i := range internals {
        fmt.Print(string(i))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + `
Generate Undetectable Malware` + bcolors.ENDC)
   exploits := bcolors.ENDC + `
1. Blackjack..............t3l3machus...............Villain]
2. ShellzGen................4ndr34z.................Shellz]
3. PowerJoker...................................PowerJoker]
4. MeterPeter...................................MeterPeter]
5. Havoc C2...............@C5pider................Havoc C2]
6. Teardroid.....................................Teardroid]
7. AndroidRAT...................................AndroidRAT]
8. Chameleon.....................................Chameleon]
9. Gh0x0st.................Gh0x0st....Invoke-PSObfuscation]
` + bcolors.ENDC
    for _, m := range exploits {
        fmt.Print(string(m))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + `
WiFi Attack Vectors` + bcolors.ENDC)
   wireless := bcolors.ENDC + `
1. Wifite...........................................Wifite]
2. Wifipumpkin3................................Wifipumpkin]
3. Airgeddon.....................................Airgeddon]
` + bcolors.ENDC
    for _, w := range wireless {
        fmt.Print(string(w))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + `
Crack Hash, Pcap & Brute Passwords` + bcolors.ENDC)
   crackers := bcolors.ENDC + `
1. Aircrack_ng.................................Aircrack_ng]
2. John...............................................John]
3. Hash-Buster.........Somdev Sangwan..........Hash-Buster]
` + bcolors.ENDC
    for _, c := range crackers {
        fmt.Print(string(c))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + `
Social-Engineering Attacks` + bcolors.ENDC)
   phishers := bcolors.ENDC + `
1. Gophish.........................................Gophish]
2. Good Ginx.....................................Evil Ginx]
3. AdvPhishing.................................AdvPhishing]
4. Setoolkit...........David Kennedy.............Setoolkit]
5. Anonphisher.................................Anonphisher]
6. Cyberphish...................................Cyberphish]
` + bcolors.ENDC
    for _, p := range phishers {
        fmt.Print(string(p))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + `
Website Attack Vectors` + bcolors.ENDC)
        webs := bcolors.ENDC + `
1. Musker.................Bing A.I.................Proxyes]
2. wafw00f.........................................wafw00f]
3. Dnsrecon.......................................Dnsrecon]
4. Seekolver..............Krypteria..............Seekolver]
5. whatweb.........................................Whatweb]
6. Harvester..................................TheHarvester]
7. Param_spider................................Paramspider]
8. Ssl_scan........................................Sslscan]
9. Gobuster.......................................Gobuster]
10. Nuclei..........................................Nuclei]
11. Nikto............................................Nikto]
12. Bbot..............................................Bbot]
13. Uniscan........................................Uniscan]
14. Sqlmap..........................................Sqlmap]
15. Commix..........................................Commix]
16. Katana..........................................Katana]
17. Xsser............................................Xsser]
18. Nettacker....................................Nettacker]
19. Jok3r..................koutto....................Jok3r]
20. Osmedeus...............j3ssie.................Osmedeus]
21. Ufonet.................epsylon..................Ufonet]
` + bcolors.ENDC
    for _, w := range webs {
        fmt.Print(string(w))
        time.Sleep(3 * time.Millisecond)
    }

    fmt.Printf(bcolors.UNDERL + bcolors.BOLD + `
Code of conduct` + bcolors.ENDC)
    caution := bcolors.ENDC + `
[ Africana-framework is written purely for Good & not evil]
[ The author of africana-framework is Rojahs Montari from.]
[ cyberafrics a cybersecurity organisation in Africa Kenya]
[What is there 4 U 2 gain the whole world & loose your....]
[soul? Be smart your Creator has good plans for you.......]
[Life Tip.: Defeat the devil by fasting & praying.........]
[Email....: rojahsmontari@gmail.com.......................]
[YouTube..: https://youtube.com/@RojahsMontari............]
` + bcolors.ENDC
    for _, c := range caution {
        fmt.Print(string(c))
        time.Sleep(3 * time.Millisecond)
    }
}

func Developer() {
    fmt.Printf(bcolors.BOLD + bcolors.UNDERL + `
About the author` + bcolors.ENDC)
    developer := bcolors.ENDC + `
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
` + bcolors.ENDC
    for _, d := range developer {
        fmt.Print(string(d))
        time.Sleep(3 * time.Millisecond)
    }
    fmt.Printf(`%s
     The devil has no power over you Christian soldier.%s
`, bcolors.BLINK, bcolors.ENDC)

    fmt.Printf(`
    %s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sReturn to main menu.%s

`, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.BOLD + bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD + bcolors.UNDERL, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.BOLD + bcolors.UNDERL, bcolors.UNDERL, bcolors.ENDC)
}
