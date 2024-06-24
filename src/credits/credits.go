package credits

import (
    "fmt"
    "time"
    "bcolors"
)

func Contributors() {
    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
                  🦊Africana-Framework` + bcolors.ENDC)
    credits := bcolors.DARKCYAN + `
[ Special thanks to the following people whose third party]
[ tools have been modified and used to contribute to the..]
[ creation of africana-framework:.........................]
` + bcolors.ENDC
    for _, c := range credits {
        fmt.Print(string(c))
        time.Sleep(28 * time.Millisecond)
    }
    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
 Africana name        Developer's name       Original_name` + bcolors.ENDC)

    fmt.Printf(bcolors.DARKGREEN + `
1. africana-framework..Rojahs montari..africana-framework ]
` + bcolors.ENDC)

    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
                 2. Tor system setup` + bcolors.ENDC)

    fmt.Printf(bcolors.DARKGREEN + `
1. tor_vanish.......Salim Zaved Karim..............Neutron]
` + bcolors.ENDC)

    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
               3. Internal-Network-Attack` + bcolors.ENDC)
   fmt.Printf(bcolors.DARKGREEN + `
3. bettercap.....................................bettercap]
4. nmap...............................................nmap]
5. metasploit-framework.........................msfconsole]
6. smbmap...........................................smbmap]
7. rpcclient.....................................rpcclient]
8. beef-xss.......................................beef-xss]
9. responder.....................................responder]
10 toxssin.........................................toxssin]
` + bcolors.ENDC)

    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
             4. Generate Undetectable Malware` + bcolors.ENDC)
   fmt.Printf(bcolors.DARKGREEN + `
1. Blackjack..............t3l3machus...............Villain]
2. ShellzGen................4ndr34z.................Shellz]
3. PowerJoker...................................PowerJoker]
4. MeterPeter...................................MeterPeter]
5. Havoc C2...............@C5pider................Havoc C2]
6. Teardroid.....................................Teardroid]
7. AndroidRAT...................................AndroidRAT]
8. Chameleon.....................................Chameleon]
9. Gh0x0st.................Gh0x0st....Invoke-PSObfuscation]
` + bcolors.ENDC)

    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
                 5. WiFi Attack Vectors` + bcolors.ENDC)
   fmt.Printf(bcolors.DARKGREEN + `
1. Wifite...........................................Wifite]
2. Wifipumpkin3................................Wifipumpkin]
3. Airgeddon.....................................Airgeddon]
` + bcolors.ENDC)

    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
           6. Crack Hash, Pcap & Brute Passwords` + bcolors.ENDC)
   fmt.Printf(bcolors.DARKGREEN + `
1. Aircrack_ng.................................Aircrack_ng]
2. John...............................................John]
3. Hash-Buster.........Somdev Sangwan..........Hash-Buster]
` + bcolors.ENDC)

    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
               7. Social-Engineering Attacks` + bcolors.ENDC)
   fmt.Printf(bcolors.DARKGREEN + `
1. Gophish.........................................Gophish]
2. Good Ginx.....................................Evil Ginx]
3. AdvPhishing.................................AdvPhishing]
4. Setoolkit...........David Kennedy.............Setoolkit]
5. Anonphisher.................................Anonphisher]
6. Cyberphish...................................Cyberphish]
` + bcolors.ENDC)

    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
                8. Website Attack Vectors` + bcolors.ENDC)

    fmt.Printf(bcolors.DARKGREEN + `
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
` + bcolors.ENDC)

}

func Developer() {
    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
                 🛰️ About the author` + bcolors.ENDC)
    developer := bcolors.BLUE + `
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
    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + `
                 🛰️ code of conduct` + bcolors.ENDC)
    caution := bcolors.PURPLE + `
[ Africana-framework is written purely for Good & not evil]
[ The author of africana-framework is Rojahs Montari from.]
[ cyberafrics a cybersecurity organisation in Africa Kenya]
` + bcolors.ENDC
    for _, c := range caution {
        fmt.Print(string(c))
        time.Sleep(3 * time.Millisecond)
    }
    fmt.Printf(bcolors.ENDC + bcolors.ITALIC + bcolors.BLINK + `
        🛰️ The devil has no power over you soldier` + bcolors.ENDC)
    fmt.Printf(bcolors.ORANGE + bcolors.ITALIC + `
[ What is there 4 U 2 gain the whole world & loose your...]
[ soul? Be smart your Creator has good plans for you......]
[ Life Tip.: Defeat the devil by fasting & praying........]
[ Email....: rojahsmontari@gmail.com......................]
[ YouTube..: https://youtube.com/@RojahsMontari...........]` + bcolors.ENDC)
    fmt.Printf(bcolors.DARKGREEN + `
         __                 _____ _____     _     _
      __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
     |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
     |_____|___|___|___|___|_____|__|__|_| |_|___|_|` + bcolors.ENDC)
    fmt.Printf(bcolors.ENDC + bcolors.BLINK + `
                         ¯\_(ツ)_/¯

` + bcolors.ENDC)

}
