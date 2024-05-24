package guides

import (
    "fmt"
    "time"
    "bcolors"
)

func Credits() {
    fmt.Printf(`%s
                     _,._
                 __.'   _)
                <_,)'.-"a\
                  /' (    \
      _.-----..,-'   ('"--^
     //              |
    (|   ';      ,   |
      \   ;.----/  ,/
       ) // /   | |\ \
       \ \\'\   | |/ /      %sJesus Christ%s
        \ \\ \  | |\/  %sThe Lamb that was slain%s
         '" '"  '""         %sfor our sins.%s
    `, bcolors.BLUE, bcolors.GREEN + bcolors.GREEN, 
bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN,
 bcolors.ENDC)
    fmt.Printf(bcolors.YELLOW + `
         __                 _____ _____     _     _
      __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
     |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
     |_____|___|___|___|___|_____|__|__|_| |_|___|_|
                         ¯\_(ツ)_/¯
    ` + bcolors.ENDC)

    fmt.Printf(bcolors.GREEN + `
                🛰️   Africana-Framework
    ` + bcolors.ENDC)

    fmt.Printf(bcolors.BLUE + `
 Africana name        Developer's name       Original_name
    ` + bcolors.ENDC)

    credits := bcolors.DARKCYAN + `
1. africana-framework..Rojahs montari..africana-framework ]

               2.    Tor system setup
1. tor_vanish.......Salim Zaved Karim..............Neutron]

               3. Internal-Network-Attack
3. bettercap.....................................bettercap]
4. nmap...............................................nmap]
5. metasploit-framework.........................msfconsole]
6. smbmap...........................................smbmap]
7. rpcclient.....................................rpcclient]
8. beef-xss.......................................beef-xss]
9. responder.....................................responder]
10 toxssin.........................................toxssin]

               4. Generate Undetectable Malware
1. Blackjack..............t3l3machus...............Villain]
2. ShellzGen................4ndr34z.................Shellz]
3. PowerJoker...................................PowerJoker]
4. MeterPeter...................................MeterPeter]
5. Havoc C2...............@C5pider................Havoc C2]
6. Teardroid.....................................Teardroid]
7. AndroidRAT...................................AndroidRAT]
8. Chameleon.....................................Chameleon]
9. Gh0x0st.................Gh0x0st....Invoke-PSObfuscation]

               5.     WiFi Attack Vectors
1. Wifite...........................................Wifite]
2. Wifipumpkin3................................Wifipumpkin]
3. Airgeddon.....................................Airgeddon]

               6. Crack Hash, Pcap & Brute Passwords
1. Aircrack_ng.................................Aircrack_ng]
2. John...............................................John]
3. Hash-Buster.........Somdev Sangwan..........Hash-Buster]

               7.   Social-Engineering Attacks
1. Gophish.........................................Gophish]
2. Good Ginx.....................................Evil Ginx]
3. AdvPhishing.................................AdvPhishing]
4. Setoolkit...........David Kennedy.............Setoolkit]
5. Anonphisher.................................Anonphisher]
6. Cyberphish...................................Cyberphish]

               8.     Website Attack Vectors
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
    for _, c := range credits {
        fmt.Print(string(c))
        time.Sleep(01 * time.Millisecond)
    }

}

func Developer() {
    fmt.Printf(bcolors.YELLOW + `
                   🛰️ About the author
    ` + bcolors.ENDC)
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
[.........................................................]
[ Africana-framework is written purely for Good & not evil]
[ The author of africana-framework is Rojahs Montari from.]
[ cyberafrics a cybersecurity organisation in Africa Kenya]
[ Special thanks to the following people whose third party]
[ tools have been used to contribute to the creation of...]
[.........................................................]
[ What is there 4 U 2 gain the whole world & loose your...]
[ soul? Be smart your Creator has good plans for you......]
[.....Life Tip.: Defeat the devil by fasting & praying....]
[---------------------------------------------------------]
` + bcolors.ENDC
    for _, d := range developer {
        fmt.Print(string(d))
        time.Sleep(03 * time.Millisecond)
    }
    fmt.Print(bcolors.RED + `
[ Email....: rojahsmontari@gmail.com......................]
[ YouTube..: https://youtube.com/@RojahsMontari...........]
    ` + bcolors.ENDC)
    fmt.Printf(bcolors.GREEN + `
             Type 0. To.Exit & Go To Main Menu

` + bcolors.ENDC)

}
