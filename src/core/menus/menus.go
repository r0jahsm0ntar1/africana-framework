package menus

import(
    "fmt"
    "utils"
    "bcolors"
)
var (
    Lhost, _ = utils.GetDefaultIP()
    Gateway, _ = utils.GetDefaultGatewayIP()
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, WordList = utils.DirLocations()
)

func MenuZero() { 
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sInstall, Update & View Logs
%s 2. %sSystem Security Configuration
%s 3. %sLocal Network Penetration Testing
%s 4. %sBackdoors & Trojans Generating Engines
%s 5. %sWireless Networks Attacks Vector
%s 6. %sPasswords, Hashes & .Pcap Crackers
%s 7. %sSocial-Engineering & Phishings Vectors
%s 8. %sWebsite Penetration Testing Vectors
%s 9. %sCredits, & about the author

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s99. %s%s%s%sGet Bonus Packages.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuOne() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sauto
%s 2. %skali
%s 3. %sarch
%s 4. %smacos
%s 5. %subuntu
%s 6. %sandroid
%s 7. %swindows
%s 8. %supdates
%s 9. %suninstall

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuTwo() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %ssetups
%s 2. %svanish
%s 3. %sstatus
%s 4. %storip
%s 5. %schains
%s 6. %sreload
%s 7. %sexitnode
%s 8. %srestore
%s 9. %sstop

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuThree() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sDiscover targets
%s 2. %sDiscover open ports
%s 3. %sStart Vuln scan
%s 4. %sActive directory recon
%s 5. %sExploit active directory
%s 6. %sSniff packets
%s 7. %sLaunch killer Responder
%s 8. %sLaunch M.I.B
%s 9. %sStart XSS Injection 

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuThreeOne() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sLaunch eternalblue exploit on the target

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuThreeThree() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sLaunch hosts dns using ettercap
%s 2. %sLaunch hosts dns using bettercap

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuThreeFour() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sInject beef-xss.js to early selected target
%s 2. %sInject beef-xss.js to all connected devices

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuFouR() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %slinux
%s 2. %smackos
%s 3. %swindows
%s 4. %siphones
%s 5. %sandroids
%s 6. %swebsites
%s 7. %sevasions
%s 8. %sbackdoors
%s 9. %slisteners

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuFour() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sghost
%s 2. %sshellz
%s 3. %slistener
%s 4. %sandrorat
%s 5. %steardroid
%s 6. %sblackjack
%s 7. %shoaxshell
%s 8. %snoisemaker
%s 9. %scodebreaker
%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuFive() {
    fmt.Printf(`

%s%s%sSelect a number from the table below.%s

%s 1. %swifite
%s 2. %sfluxion
%s 3. %sbettercap
%s 4. %sairgeddon
%s 5. %swifipumpkin
%s 6. %swifipumpkin3
%s 7. %sto add
%s 8. %sto add
%s 9. %sto add

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuSix() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sonline services.(bruteforcing remote services like ssh, ftp etc.)
%s 2. %soffline services.(cracking of hashes/ntlms and .pcap)
%s 3. %sto add
%s 4. %sto add
%s 5. %sto add
%s 6. %sto add
%s 7. %sto add
%s 8. %sto add
%s 9. %sto add

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func HelpInfoSix() {
    fmt.Println(bcolors.YELLOW + "                 Online " + bcolors.ENDC + ": 1. Attack services that are online. Ex SMB, SSH, FTP, HTTP, HTTPS etc." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                Offline " + bcolors.ENDC + ": 2. Attack .pcaps, hashes, captured ntlms or documents with passwords." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Exit & Go to Main Menu " + bcolors.ENDC + ": 0. Exit and go back to previous menu." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "   Current working Menu " + bcolors.ENDC + ": 99. Print current working menu and its options." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                   Help " + bcolors.ENDC + ": 00. Found in all faces of africana-framework. It will guide you through. Use it if lost or confused." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "    Launch shell or cmd " + bcolors.ENDC + "    .bash, shell, pwsh, cmd or sh commands launches system intaractive shell. One can do all you do with terminal here.." + bcolors.ENDC)
}

func MenuSixOne() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sssh
%s 2. %sftp
%s 3. %ssmb
%s 4. %srdp
%s 5. %sldap
%s 6. %ssmtp
%s 7. %stelnet
%s 8. %shttp/https
%s 9. %sall services

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func HelpSixOne() {
    fmt.Println(bcolors.YELLOW + "                    SSH " + bcolors.ENDC + ": 1. Automated Bruteforcer for SSH using rockyou.txt wordlists." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                    FTP " + bcolors.ENDC + ": 2. Bruteforcer for FTP using rockyou.txt wordlists." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                    SMB " + bcolors.ENDC + ": 3. Hydra Bruteforcer for SMB using rockyou.txt wordlists." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                    RDP " + bcolors.ENDC + ": 4. Bruteforcer for RDP using rockyou.txt wordlists." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                   LDAP " + bcolors.ENDC + ": 5. Hydra Bruteforcer for LDAP using rockyou.txt wordlists." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                   SMTP " + bcolors.ENDC + ": 6. Bruteforcer for SMTP using rockyou.txt wordlists." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                 TELNET " + bcolors.ENDC + ": 7. Bruteforcer for TELNET using rockyou.txt wordlists." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "            HTTP/ HTTPS " + bcolors.ENDC + ": 8. Hydra Bruteforcer for HTTP/ HTTPS using rockyou.txt wordlists. Forms needed." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "           All SERVICES " + bcolors.ENDC + ": 9. The Automatic Bruteforce Tool for all opened services. Works nice and automatic." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Exit & Go to Main Menu " + bcolors.ENDC + ": 0. Exit and go back to previous menu." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "   Current working Menu " + bcolors.ENDC + ": 99. Print current working menu and its options." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                   Help " + bcolors.ENDC + ": 00. Found in all faces of africana-framework. It will guide you through. Use it if lost or confused." + bcolors.ENDC)
}

func MenuSixTwo() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sjohn
%s 2. %saircrack_ng
%s 3. %sto add
%s 4. %sto add
%s 5. %sto add
%s 6. %sto add
%s 7. %sto add
%s 8. %shash_buster
%s 9. %sto add

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func HelpSixTwo() {
    fmt.Println(bcolors.YELLOW + "            Aircrack_ng " + bcolors.ENDC + ": 1. Automated. It will prompt you for location of the file e.g. ntlm file to crack using rockyou.txt wordlists. Give full path." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                   John " + bcolors.ENDC + ": 2. Automated. It will prompt you for location of the file e.g. .pcap file to crack using rockyou.txt wordlists. Give full path." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                 to add " + bcolors.ENDC + ": 3. Still working on this option." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                 to add " + bcolors.ENDC + ": 4. Still working on this option." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                 to add " + bcolors.ENDC + ": 5. Still working on this option." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                 to add " + bcolors.ENDC + ": 6. Still working on this option." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                 to add " + bcolors.ENDC + ": 7. Still working on this option." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "            Hash-Buster " + bcolors.ENDC + ": 8. It is a tool to crack hashes. It need internet connection to search for the hashes before caracking them." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                 to add " + bcolors.ENDC + ": 9. Still working on this option." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Exit & Go to Main Menu " + bcolors.ENDC + ": 0. Exit and go back to previous menu." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "   Current working Menu " + bcolors.ENDC + ": 99. Print current working menu and its options." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                   Help " + bcolors.ENDC + ": 00. Found in all faces of africana-framework. It will guide you through. Use it if lost or confused." + bcolors.ENDC)
}

func MenuSeven() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sgophish
%s 2. %sgoodginx
%s 3. %szphisher
%s 4. %sblackeye
%s 5. %sadvnphish
%s 6. %sdarkphish
%s 7. %sshellphish
%s 8. %ssetoolkit
%s 9. %sthehackerchoice

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuSevenOne() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sninja phish           attack early selected target
%s 2. %sgoodginx              attack all Connected devices
%s 3. %sto add
%s 4. %sto add
%s 5. %sto add
%s 6. %sto add
%s 7. %sto add
%s 8. %sto add
%s 9. %sto add
%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuEight() {
    fmt.Printf(`

%s%s%sSelect a number from the table below.%s

%s 1. %sSubdomain enumeration
%s 2. %sDns and asn lookup
%s 3. %sNetwork Mapping
%s 4. %sWeb technology detection
%s 5. %sWayback and asset discovery
%s 6. %sFuzzing and Scanning
%s 7. %sSecret and leak detection
%s 8. %sVulnerability scanning (XSS, SQLi, CSRF, SSRF, IDOR)
%s 9. %sAutomation (Bug bounty hunting techniques)

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func HelpEight() {
    fmt.Println(bcolors.YELLOW + "                Web recon " + bcolors.ENDC + ": 1. Automated. Start Passive Web recon & Subdomain Enumration. Chosen tools with less noise are used to gather information on agiven domain before Intrusion begins." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "           Gather e-mails " + bcolors.ENDC + ": 2. Theharvester tool will try to harvest emails with  names that could be used for phishing, creat wordlists or also as login names for admin panels." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "       Start Bruteforcing " + bcolors.ENDC + ": 3. In this phase, Gobuster tool is used to show us hidden files in the website and also try to find login pages." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "  Start 💉SQL, XSS & SSRF " + bcolors.ENDC + ": 4. This is an injection phase. Chosen tools are configured to find injection vulns. The choice contain embeded submenu that supports both automation and manual injection tackticks." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "          OWASP Nettacker " + bcolors.ENDC + ": 5. Is a fullset framework to automate all website vulnerbilities. An embeded submenu within this option will guide you on how to attack website In an advanced mode.(Intrusive phase)." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                    Jok3r " + bcolors.ENDC + ": 6. Like Nettacker, Joker is a fullset framework. It is intergrated with a collection of several third party tools. The embeded menu in this option will guide you through." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "               Osmedeus🔮 " + bcolors.ENDC + ": 7. It is also a rich pentest framework to help you bug bount websites. It is fast written in go. The embeded menu in this choice will guigde you." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "            Ufonet 🦠DDOS " + bcolors.ENDC + ": 8. Danger!!! This tool kills websites completely. Use it at your own risk. The embeded menu on this choice will guide you." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Heavy Automation Attacks " + bcolors.ENDC + ": 9. Chosen tools are selected in this option to run simulteniously giving you all desirable Outcomes. Other tools are added that are not in the above options." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "   Exit & Go to Main Menu " + bcolors.ENDC + ": 0. Exit and go back to previous menu." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "     Current working Menu " + bcolors.ENDC + ": 99. Print current working menu and its options." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                     Help " + bcolors.ENDC + ": 00. Found in all faces of africana-framework. It will guide you through. Use it if lost or confused." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "      Launch shell or cmd " + bcolors.ENDC + "    .bash, shell, pwsh, cmd or sh commands launches system intaractive shell. One can do all you do with terminal here.." + bcolors.ENDC)
}

func MenuEightFour() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %ssql_auto
%s 2. %sxss_auto
%s 3. %ssql_manual
%s 4. %sxss_manual
%s 5. %snettacker
%s 6. %sto add
%s 7. %sto add
%s 8. %sto add
%s 9. %sto add

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func HelpEightFour() {
    fmt.Println(bcolors.YELLOW + "       Sql injection(A) " + bcolors.ENDC + ": 1.(Automated) Sqlmap tool is configured to run with evasion templets to avoid detection and give you desirable outcomes.(Intrusive)." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "       Xss injection(A) " + bcolors.ENDC + ": 2.(Automated) Acompination of tools are used to try and find xss injection vulnerbility on the website.(Intrusive)." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "       Sql injection(M) " + bcolors.ENDC + ": 3.(Manual) In this option you will be required to enter sql website injection required data manually." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "       Xss injection(M) " + bcolors.ENDC + ": 4.(Manual) In this option you will be required to enter xss website injection required data manually." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Exit & Go to Main Menu " + bcolors.ENDC + ": 0. Exit and go back to previous menu." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "   Current working Menu " + bcolors.ENDC + ": 99. Print current working menu and its options." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                   Help " + bcolors.ENDC + ": 00. Found in all faces of africana-framework. It will guide you through. Use it if lost or confused." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "    Launch shell or cmd " + bcolors.ENDC + "    .bash, shell, pwsh, cmd or sh commands launches system intaractive shell. One can do all you do with terminal here.." + bcolors.ENDC)
}

func MenuEightFive() {
    fmt.Println(bcolors.BLUE + "        ~>( 🪲" + bcolors.ENDC + bcolors.ITALIC + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " )<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "\n[ 1. Start Nettacker port & web content discovery.......📡] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Recon, find and scan subdomains....................🐾] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Launch admin_scan to find admin panel..............🦨] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Start insane information gathering on host.........🧄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Run only recon & vuln scan security checks.........🦍] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Run CVE scans on the target host...................🍹] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Search for(critical vulns & easy to exploit)......🍄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. Automate all Functions & security checks on target...🤖] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. View scaned Nettacker report list..................🥝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Go back............................................🧬] " + bcolors.ENDC)
    fmt.Println(bcolors.Colors() + ` ¯\_(ツ)_/¯                     ` + bcolors.YELLOW + "🎱99. " + bcolors.DARKCYAN + "Guide! " + bcolors.YELLOW + "🎲00. " + bcolors.DARKCYAN + "Help!" + bcolors.BLUE + "📜)" + bcolors.ENDC)
}

func HelpEightFive() {
    fmt.Println(bcolors.YELLOW + "  Nettacker port & web  " + bcolors.ENDC + ": 1. Scans for all open ports and services. Recons for hidden root files in the target host." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "        Scan subdomains " + bcolors.ENDC + ": 2. Digs for all sudomains related to your target for alternative pentesting." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "      Launch admin_scan " + bcolors.ENDC + ": 3. Scans the given target in an aim of finding the login panels where you can try SQL injections & password bruteforcing." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "  Insane info gathering " + bcolors.ENDC + ": 4. The tool will recon for all necessary information on the website essential for your pentesting." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "    Vuln scansec checks " + bcolors.ENDC + ": 5. Check for known vulnerbilities on the website that are easy to exploit." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "              CVE scans " + bcolors.ENDC + ": 6. Scan the website for a collection of discovered CVES ready and easy to attack." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "        Critical vulns  " + bcolors.ENDC + ": 7. The scanner will scan for critical & easy to exploit bugs like SQLinjection, XSS and Others." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Automate all Functions " + bcolors.ENDC + ": 8. For lazy hackers this option got your needs. It will compine all the above tacticks to one and show you findings." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "     View scanned report" + bcolors.ENDC + ": 9. Show all logs of scanned domains with findings." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Exit & Go to Main Menu " + bcolors.ENDC + ": 0. Exit and go back to previous menu." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "   Current working Menu " + bcolors.ENDC + ": 99. Print current working menu and its options." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                   Help " + bcolors.ENDC + ": 00. Found in all faces of africana-framework. It will guide you through. Use it if lost or confused." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "    Launch shell or cmd " + bcolors.ENDC + "    .bash, shell, pwsh, cmd or sh commands launches system intaractive shell. One can do all you do with terminal here.." + bcolors.ENDC)
}

func MenuEightSix() {
    fmt.Println(bcolors.BLUE + "        ~>( 🪲" + bcolors.ENDC + bcolors.ITALIC + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " )<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "\n[ 1. Install all jok3r tools....." + bcolors.DARKCYAN + "Pleas start here if not" + bcolors.BLUE + "🃏] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Update all the tools in the toolbox................🍄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Show all the tools in the toolbox..................🍒] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Show supported products for all services...........🍵] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Run only recon & vulnscansecurity checks...........🦍] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Run all & intense security checks against an URL...🤖] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Search for(critical vulns & easy to exploit)......🌈] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. View the full results from the security checks.....🧄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. Clean database & delete results....................🥑] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Go back............................................🧬] " + bcolors.ENDC)
    fmt.Println(bcolors.Colors() + ` ¯\_(ツ)_/¯                     ` + bcolors.YELLOW + "🎱99. " + bcolors.DARKCYAN + "Guide! " + bcolors.YELLOW + "🎲00. " + bcolors.DARKCYAN + "Help!" + bcolors.BLUE + "📜)" + bcolors.ENDC)
}

func HelpEightSix() {
    fmt.Println(bcolors.YELLOW + "    Install jok3r tools " + bcolors.ENDC + ": 1. Jok3r comes without third party tools preinstalled. This option will make sure that all necessary tools needed by Jok3r are installed for better results.(Run once)." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "          Update toolbox " + bcolors.ENDC + ": 2. If all third party are installed, you could update the tools using this option.(Not necessary in new install.)" + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "      Show all the tools " + bcolors.ENDC + ": 3. Health check. See if all tools are installed correctly and are working properly before engaging in hacking." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Show supported products " + bcolors.ENDC + ": 4. Jok3r will show you all supported products and services that it is made to attack." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Vulnscansecurity checks " + bcolors.ENDC + ": 5. Wide range of vulnerbility checks and bugs on the given target will be scanned by Jok3r as it store outcome in its database." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Intense security checks " + bcolors.ENDC + ": 6. A heavy vulnerbility scanning on the target host will be done storing all logs in Jok3r database." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "          Critical vulns " + bcolors.ENDC + ": 7. Jok3r will Search for(critical vulns & easy to exploit) avoiding heavy intrusions on your target." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "       View full results " + bcolors.ENDC + ": 8. All logs for the scanner will be displayed showing even the targets and their vulns." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "      Clean the database " + bcolors.ENDC + ": 9. Delete all results and logs for previous scanns in the databases." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "  Exit & Go to Main Menu " + bcolors.ENDC + ": 0. Exit and go back to previous menu." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "    Current working Menu " + bcolors.ENDC + ": 99. Print current working menu and its options." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                    Help " + bcolors.ENDC + ": 00. Found in all faces of africana-framework. It will guide you through. Use it if lost or confused." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "     Launch shell or cmd " + bcolors.ENDC + "    .bash, shell, pwsh, cmd or sh commands launches system intaractive shell. One can do all you do with terminal here.." + bcolors.ENDC)
}

func MenuEightSeven() {
    fmt.Println(bcolors.BLUE + "        ~>( 🪲" + bcolors.ENDC + bcolors.ITALIC + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " )<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "\n[ 1. Update Osmedeus & Run diagnostics to check config..🍄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Start a simple scan with other flow................🦠] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Launch vuln and directory scan on domains..........🦍] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Scan list of targets(Full path of target needed)..🥐] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Cloud - Run scan in Distributed Cloud mode.........🔮] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Performing Full web vuln & secret scan on host.....🤖] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Update the vulnerability database before attacking.📜] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. Start web UI server................................🌍] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. View scanned osmedeus report list..................🌈] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Go back............................................🧬] " + bcolors.ENDC)
    fmt.Println(bcolors.Colors() + ` ¯\_(ツ)_/¯                     ` + bcolors.YELLOW + "🎱99. " + bcolors.DARKCYAN + "Guide! " + bcolors.YELLOW + "🎲00. " + bcolors.DARKCYAN + "Help!" + bcolors.BLUE + "📜)" + bcolors.ENDC)
}

func HelpEightSeven() {
    fmt.Println(bcolors.YELLOW + "         Update Osmedeus " + bcolors.ENDC + ": 1. Osmedeus comes without third party tools preinstalled. This option will make sure that all necessary tools needed by Osmedeus are installed for better results.(Run once)." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "             Simple scan " + bcolors.ENDC + ": 2. The tool will Start a simple scan with other flow like discovering domains and ports.(Reconing)." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "      Vuln and directory " + bcolors.ENDC + ": 3. Osmedeus will check for vulnerbilities and discover hidden root directories disclosing related subdomains on your target." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "    Scan list of targets " + bcolors.ENDC + ": 4. If you have alist of several targets you would need to scan at once. Mybe a harvested subdomains, use this option." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "               Cloud-Run " + bcolors.ENDC + ": 5. In this option, Osmedeus will allow you to scan your target using cloud database." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "   Performing Full recon " + bcolors.ENDC + ": 6. Osmedeus will run in automated mode performing all required security checks and data mining in this mode storing logs for review." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Update vuln db & attack " + bcolors.ENDC + ": 7. Osmedeus will Update the vulnerability database before attacking your selected target." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "     Start web UI server " + bcolors.ENDC + ": 8. Launch a web UI which you can interact with Osmedeus in a graphickal paspective." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "     View scanned report " + bcolors.ENDC + ": 9. Reports and all logs for the scanner will be displayed showing even the targets and their vulns." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "  Exit & Go to Main Menu " + bcolors.ENDC + ": 0. Exit and go back to previous menu." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "    Current working Menu " + bcolors.ENDC + ": 99. Print current working menu and its options." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                    Help " + bcolors.ENDC + ": 00. Found in all faces of africana-framework. It will guide you through. Use it if lost or confused." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "     Launch shell or cmd " + bcolors.ENDC + "    .bash, shell, pwsh, cmd or sh commands launches system intaractive shell. One can do all you do with terminal here.." + bcolors.ENDC)
}

func MenuEightEight() {
    fmt.Println(bcolors.BLUE + "        ~>( 🪲" + bcolors.ENDC + bcolors.ITALIC + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " )<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "\n[ 1. Download list of " + bcolors.RED + "Bots " + bcolors.BLUE + "from" + bcolors.YELLOW + " Community " + bcolors.BLUE + "server........🍄]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Test if bots are alive.............................📡] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Launch palantir....................." + bcolors.DARKCYAN + "Palantir 3.14.." + bcolors.BLUE + "🦍] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Launch socking_waves................" + bcolors.DARKCYAN + "Knockout!......" + bcolors.BLUE + "🦠] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Launch xcom-1......................." + bcolors.DARKCYAN + "Only DDoS......" + bcolors.BLUE + "🥐] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Launch xcom-2......................." + bcolors.DARKCYAN + "Only DoS......." + bcolors.BLUE + "🍵] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Launch ufonet-gui..................." + bcolors.DARKCYAN + "Gui on browser." + bcolors.BLUE + "🧄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. Start Grider........................" + bcolors.DARKCYAN + "Grider........." + bcolors.BLUE + "🐝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. Launch Armageddon!.................." + bcolors.DARKCYAN + "Launch All!!..." + bcolors.BLUE + "🥝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Go back............................................🦨] " + bcolors.ENDC)
    fmt.Println(bcolors.Colors() + ` ¯\_(ツ)_/¯                     ` + bcolors.YELLOW + "🎱99. " + bcolors.DARKCYAN + "Guide! " + bcolors.YELLOW + "🎲00. " + bcolors.DARKCYAN + "Help!" + bcolors.BLUE + "📜)" + bcolors.ENDC)
}

func HelpEightEight() {
    fmt.Println(bcolors.YELLOW + "  Download list of Bots " + bcolors.ENDC + ": 1. Ufonet will download a list of zombies machines from Community servers that will be used durring DDOS attacks.." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Test if bots are alive " + bcolors.ENDC + ": 2. After downloding the zombies, it is good to test them if they are alive and working for a better results." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "        Launch palantir " + bcolors.ENDC + ": 3. It is a kind of ddos attack that can be launched in non-root mode." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "   Launch socking_waves " + bcolors.ENDC + ": 4. Another DDOS attack mode that yields instant Knockout!." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "          Launch xcom-1 " + bcolors.ENDC + ": 5. Ufonet launches only DDoS mode without a combination of other modes." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "          Launch xcom-2 " + bcolors.ENDC + ": 6. Only DDoS mode without a combination of other attack modes." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "             Ufonet-gui " + bcolors.ENDC + ": 7. Launch a web UI which you can interact with ufonet in a graphickal paspective." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "           Start Grider " + bcolors.ENDC + ": 8. Share your zombies. WARNING: this *ADVANCED* function is *NOT* secure, proceed if you really want to." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "     Launch Armageddon! " + bcolors.ENDC + ": 9. Very dangerous & Destructive. Use it at your own risk.(Attack With All!)." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + " Exit & Go to Main Menu " + bcolors.ENDC + ": 0. Exit and go back to previous menu." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "   Current working Menu " + bcolors.ENDC + ": 99. Print current working menu and its options." + bcolors.ENDC)
    fmt.Println(bcolors.YELLOW + "                   Help " + bcolors.ENDC + ": 00. Found in all faces of africana-framework. It will guide you through. Use it if lost or confused." + bcolors.ENDC)
}


func UpsentTools() {
    fmt.Println(bcolors.YELLOW + "\n[!] " + bcolors.ENDC + "Choice selected not implimented yet!, but comming soon!" + bcolors.ENDC)
}

func HelpInfoMenuMain() {
    fmt.Printf(`
Usage: ./africana [options]

Common options:
    -a, --auto          Start africana in automation mode 'start from main menu'

Setup options:
    -i, --install       Launch Installation menu to install neede dependencies
    -u, --update        Update africana and africana-base tools

Framework options:
    -l, -L, --logs      Show logs in terminal
    -v, -V, --version   Show version

Console options:
    -t, --torsocks      Launch torsocks menu to torify your system.
    -n, --networks      Start internal network attacks.
    -e, --exploits      Generate undetectebal R.A.Ts and (Launch c2s for all systems. Evasions also included)
    -w, --wireless      From wifi, bluetooth, cantools and other wireles attack vectors.
    -c, --crackers      Crack(NTLMS, HASHES, PCAPS) & bruteforce(SSH, FTP, SMB, RPC etc.)
    -p, --phishers      Perform almost all advanced Phishing attacks.
    -x, --websites      Launch Web Penetration engines with free bugbounty automation function.
    -k, --credits       Show who developes and mentains africana-framework and (third party tools developers)
    -b, --verses        Launch Bible verses in an uniform way manner as used in the framework.
    -g, --guide         Watch tutarials on %sYouTube %s: %s%shttps://youtube.com/@RojahsMontari%s.
    -h, --help          Show this message.
`, bcolors.DARKGREEN, bcolors.ENDC, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func HelpInfoMenuZero() {
    fmt.Printf(`
    %sCommand             Description%s
    -------             -----------
    ?                   Help menu. Alias to (h, help)
    banner              Display an awesome africana banner
    clear               Clear the working screen or (flag 'history' clear history)
    exit                Exit the console
    features            Display the list of not yet released features that can be opted in to
    guide               Watch tutarials on %sYouTube %s: %s%shttps://youtube.com/@RojahsMontari%s.
    history             Show command history
    menu                Print the menu of the current phase. Alias to letter(m)
    quit                Exit the console
    set                 Sets a context-specific variable to a value
    sleep               Do nothing for the specified number of seconds
    tips                Show a list of useful productivity tips
    version             Show the framework and console library version numbers

%sFunction Commands%s
---------------

    %sCommand             Description%s
    -------             -----------
    advanced            Displays advanced options for one or more modules
    back                Move back from the current context
    info                Displays information about one or more modules
    list                List the Function stack
    options             Displays global options or for one or more modules
    run                 Run a selected Function
    search              Searches Function names and descriptions
    show                Displays modules of a given type, or all modules
    use                 Interact with a module by name or search term/index

%sDeveloper Commands%s
------------------

    %sCommand             Description%s
    -------             -----------
    bash                 Start Intaractive shell in africana(other flags 'sh, zsh, cmd, powershell')
    logs                 Display framework.log paged to the end if possible
    time                 Time how long it takes to run a particular command

For more info on a specific command, use %s<command> -h %sor %shelp <command>%s.

`, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func HelpInfoRun() {
    fmt.Printf(`
Usage: run [Function]

Run a Function
    show modules to list modules or <show all>

%sCommand      Description%s
-------      -----------
run enables you to Interact with a Function by name or search term/index.
      alias to "use", "exec", "start", "launch", "exploit", "execute".

Examples:
    run setups or execute setups

`, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfOptions() {
    fmt.Printf(`
Usage: options

Examples:
    options alias to [show options]

Just like show options command, options shows commands in which a sub Function can run in a given submenu.

`)
}

func HelpInfoShow() {
    fmt.Printf(`
Usage: show [all], [modules], [options]

Examples:
  show modules

%s[*]%s Valid parameters for the %s"show" %scommand are: all, modules, options
%s[*]%s Additional module-specific parameters are: missing, advanced, evasion, targets, actions

`, bcolors.BLUE, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC)
}

func HelpInfoUse() {
    fmt.Printf(`
Usage: use <name|term>

Just like start & run, use enables you to Interact with a Function by name or search term/index.
If a Function name is not found, it will be treated as a search term.
An index from the previous search results can be selected if desired.

Examples:
  use setups, torsocks, networks, exploits, wireless, phishers, websites, credits, verses

  use setups
  use <name>

`)
}

func HelpInfoStart() {
    fmt.Printf(`
Usage: start <name|term>

Just like use & run, start enables you to Interact with a Function by name or search term/index.
If a Function name is not found, it will be treated as a search term.
An index from the previous search results can be selected if desired.

Examples:
  start setups, torsocks, networks, exploits, wireless, phishers, websites, credits, verses

  start setups
  start <name>

`)
}

func HelpInfoList() {
    fmt.Printf(`
Usage: list <name|term>

Show all modules or sub-Functions in a specific face you are. If you are in modules
setups, list modules command, will list all sub-modules available in that Function.

Examples:
  list modules

%s[*]%s Valid parameters for the "list" command are: modules, all
%s[*]%s Additional Function-specific parameters are: missing, advanced

`, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC)
}

func HelpInfoSet() {
    fmt.Printf(`
Usage: set [options] [name] [value]

Set the given option to value. If value is omitted, print the current value.
If both are omitted, print options that are currently set.

If run from a Function context, this will set the value in the Function's
datastore: Use -g to operate on the global datastore.

If setting a PAYLOAD, this command can take an index from 'show payloads'.

OPTIONS:

    -c, --clear   Clear the values, explicitly setting to nil(default)
    -g, --global  Operate on global datastore variables
    -h, --help    Help banner.

`)
}

func HelpInfo() {
    fmt.Printf(`
Usage: info <Function name>

Examples:
  info setups, torsocks, internals, exploits, wireless, phishers, webs, verses

Queries the supplied Function or modules for information. If no Function is given,
show info for the currently active Function.

`)
}

func HelpInfoTips() {
    fmt.Printf(`
Usage: getips <Function name>

Examples:
  info setups, torsocks, internals, exploits, wireless, phishers, webs, verses

Gives you special information about a Function and how to successfully use it in
real life scenario.

`)
}

func HelpInfoFeatures() {
    fmt.Printf(`
Usage: features

Prints feature plans for africana. Ex. modules to be added, what to be corrected

`)
}

func HelpSetupsInfo() {
    fmt.Printf(`
Usage: info <Function name>

Examples:
  info setups, torsocks, internals, exploits, wireless, phishers, webs, verses

Queries the supplied Function or modules for information. If no Function is given, show info for the currently active Function.

`)
}

func HelpUsageSetups() {
    fmt.Printf(`
Usage: run [setups]

Description:
  This Function enables you to Install, uninstall, update & mentain africana

`)
}

func HelpInfoExecute() {
    fmt.Printf(`
%sOptions%s:
-------
    %srun%s : Launch the selected Function. Alias to (start, execute)

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func ListMainFunctions() {
    fmt.Printf(`
%sMain Functions%s:
--------------

  # %sName       Description%s
  - ----       -----------
  %s1. %s   setups%s: Install, Update, Repair or Uninstall africana-framework.
  %s2. %s torsocks%s: Configure the system for strictly tight anonymity.(fix dns leak, changemac in auto boot)
  %s3. %s networks%s: Start internal network attacks. From sniffing, injecting packets to hooking hosts.
  %s4. %s exploits%s: Generate undetectebal R.A.Ts and (Launch next gen c2s to attack all systems. Evasions also included)
  %s5. %s wireless%s: From wifies, bluetooths, cantools and other wireless networks attack vectors.
  %s6. %s crackers%s: Crack (NTLMS, HASHES, PCAPS) & bruteforce (SSH, FTP, SMB, RPC etc.)
  %s7. %s phishers%s: Perform almost all kinds of Phishing attacks. With OPT Bypass.
  %s8. %s websites%s: Launch Web Penetration testing engines with full free bugbounty automation kit.
  %s9. %s  credits%s: Show who developes and mentains africana-framework with (third party tools developers)
  %s99.%s   verses%s: Launch Bible verses in an uniform way manner as used in the framework.

%sex. %s%susage%s:
--  ------

    set module setups or -> use setups

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoMain() {
    fmt.Printf(`
       %sName%s: main
   %sFunction%s: src/core/africana
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Distros

%sDescription%s:
-----------
  This is the main module containing all africana-framework functions.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
 ListMainFunctions()
}

func MainOptions() {
    fmt.Printf(`
%sModule options %s(main/africana_run.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  module         none             yes       Module to interact with. supported modules. -> (setups, torsocks, networks, exploits, wireless, crackers, phishers, websites, credits, verse).
                                            Integer support.   -> (1 2 3 4 5 6 7 8 9 0 and 99)
%sex. %s%susage%s:
--  ------
  set module setups or -> use 1, use setups, execute 1, execute setups .etc

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func HelpInfoCrackers() {
    fmt.Printf(`
       Name: Crackers
     module: /src/crackers
   Platform: All linux/ Macos
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All linux/ macos

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  Target         none             Yes       If online attacks, a target is required to be attacked
  Wordlists      rockyou.txt      Yes       Wordlists to be used for dictionary attacks
  Pcap file      none             Yes       If offline, a file with encription needed to be bruteforced

Description:
  Crackers is a module enriched with creative attacking faces to help redtemers in successfully cracking or brutforce passwords from services online or local encripted files.

`)
}

func HelpInfoPhishers() {
    fmt.Printf(`
       Name: Phishers
     module: /src/phishers
   Platform: All linux/ Macos
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All linux/ macos

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  Interface      eth0             No        If Ninja Phish is used then an interface is required
  Target         none             No        If Ninja Phish sub-module is to be used

Description:
  Phishers is a module enriched with creative attacking faces to help redtemers in successfully Perform phishing attacks with ease.

`)
}

func HelpInfoWebsites() {
    fmt.Printf(`
       Name: Websites
     module: /src/websites
   Platform: All linux/ Macos
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All websites

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  Target         none             Yes       Set a target to be attacked
  Proxies        none             No        Set proxies to intercept traffick using Burp or force all traffick through tor

Description:
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos, sql/xss/command/_injectors among others

`)
}

func HelpInfoCredits() {
    fmt.Printf(`
       Name: Credits
     module: /src/credits
   Platform: All Distros
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
          none

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------

Description:
  Credits is a module to print on third party tools with their authors. It enables africana to acknowledge each developer for his/her hard work

`)
}

func HelpInfoVerses() {
    fmt.Printf(`
       Name: Verses
     module: /src/verses
   Platform: All Distros
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
          none

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------

Description:
  Verses is a module to narate the Bible scriptures one by one. It enables africana developer to acknowledge our LORD GOD JESUS CHRIST for Creating everything including you, me & everything.

`)
}

func HelpInfoSetups() {
    fmt.Printf(`
       %sName%s: setups
   %sFunction%s: src/core/setups
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sTested Distros%s:
------ -------
      Id  Name
      --  ----
   -> 0   All Distros

%sDescription%s:
-----------
  This modules enables you to Install, uninstall, update & mentain africana-framework.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
     ListSetupsDistros()
}

func ListSetupsDistros() {
    fmt.Printf(`
%sSupported Distros%s:
--------- -------

  # %sName       Description%s
  - ----       -----------
  %s1. %s     kali%s: Necessary tools will be installed in in Kali-Linux(Debian distros). (Use this it is stable)
  %s2. %s     arch%s: Tools will be installed in any Arch-Linux Distros using Blackarch repo.
  %s3. %s   ubuntu%s: This function will install africana-framework in Ubuntu-Linux.
  %s4. %s    macos%s: Under development. Install africana on MackingTosh systems.
  %s5. %s  android%s: Install africana-framework in Termux using chroot environment.
  %s6. %s  windows%s: Under development. But can run if tools well installed using commando vm.

%sex. %s%susage%s:
--  ------

  set distro kali
  set function install -> other functions include (update, repair and uninstall)
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func ListSetupsFunction() {
    fmt.Printf(`
%sSetups functions%s:
----------------

  # %sName       Description%s
  - ----       -----------
  %s1. %s     auto%s: Auto detect system and do the necessary.
  %s2. %s   repair%s: Repair africana-framework if broken or with issues.
  %s3. %s   update%s: Get new release of africana-framework from github and install it.
  %s4. %s  install%s: Installs africana in selected distro.

%sex. %s%susage%s:
--  ------
  set distro kali
  set function install
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func SetupsOptions() {
    fmt.Printf(`
%sModule options %s(setups/setups_launcher.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  distro         none             yes       Distro to install africana on.    -> supported distros. (arch, ubuntu, macos, android, windows).
  function       none             yes       The function to execute.          -> ex. (Install, update, repair or uninstall).
  run            none             yes       To execute the function. Alias to -> (start, execute, exec, launch).

%sex. %s%susage%s:
--  ------
  set distro kali
  set function install
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}


func SetupOptions() {
    fmt.Printf(`
Use: run [function]
   example: run kali, start kali or kali
   
  commands: run, show, start, list
            run takes flags like [run kali]
      flag: kali, ubuntu, arch, windows, android, viewlogs, browserlogs, clearlogs or uninstall

Description:
  Select an action you want to perform using the above modules Install, uninstall, update & mentain africana

`)
}

func HelpInfoKali() {
    fmt.Printf(`
       Name: kali
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All Distros

Description:
  It is a module to install africana-framework in kali-linux a stable debian based distro that
  has a wide comunity support to avoid package breaks and missing dependencies use kali for africana.

`)
}

func HelpInfoArch() {
    fmt.Printf(`
       Name: arch
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All Distros

Description:
  It is a module to install africana-framework in arch based distros. Arch is well established
  and all tools could be installed with blackman an intergration of black-arch in any arch-linux distro.
  No errors reported. africana can run well in arch-linux distros.

`)
}

func HelpInfoUbuntu() {
    fmt.Printf(`
       Name: ubuntu
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All Distros

Description:
  It is a module to install africana-framework ubuntu which is a good distor but has alot
  of problems while installing kali-linux packages. To avoid issues like dependencies problems,
  Pleas use docker image or install kali-linux in Ubuntu docker then install africana.

`)
}

func HelpInfoMacos() {
    fmt.Printf(`
       Name: macos
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All Distros

Description:
  It is a module to install africana-framework Macos which is a good distor but has alot
  of problems while installing kali-linux packages. To avoid issues like dependencies problems,
  Pleas use docker image or install kali-linux in Macos docker then install africana.

`)
}

func HelpInfoWindows() {
    fmt.Printf(`
       Name: windows
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All windows

Description:
  Under development. But can run if tools well installed. Just give me time. To intergarte this Function
  with commando vm for windows.

`)
}

func HelpInfoAndroid() {
    fmt.Printf(`
       Name: android
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   Android

Description:
  It is a Function to install africana-framework in android devices. Kali-linux will be installed in
  termux then kali-linux in chroot environment that will set all dependencies for africana-framework.

`)
}

func HelpInfoUpdate() {
    fmt.Printf(`
       Name: update
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   all

Description:
  It is a Function to update africana-framework

`)
}

func HelpInfoTermLogs() {
    fmt.Printf(`
       Name: termlogs
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All Distros

Description:
  This Function will show you all logs in your terminal that has been recorded from the last time you cleaned the log folder.

`)
}

func HelpInfoBrowserLogs() {
    fmt.Printf(`
       Name: browserlogs
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All Distros

Description:
  This Function will show you all logs in the browser that has been recorded from the last time you cleaned the log folder.

`)
}

func HelpInfoClearLogs() {
    fmt.Printf(`
       Name: clearlogs
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All Distros

Description:
  This Function will clear all your logs that has been recorded from the last time you cleaned the log folder.

`)
}

func HelpInfoUninstaller() {
    fmt.Printf(`
       Name: uninstall
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All Distros

Description:
  Uninstall africana completelty from your system with all it's dependencies. Incase of a bug, email me at rojahsmontari@gmail.com

`)
}

func HelpInfoAuto() {
    fmt.Printf(`
       Name: auto
   Function: src/core/setups
   Platform: All
       Arch: x64, x86, amd_64, android
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Normal
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All Distros

Description:
  Uninstall africana completelty from your system with all it's dependencies. Incase of a bug, email me at rojahsmontari@gmail.com

`)
}

func ListTorsocksFunctions() {
    fmt.Printf(`
%sTorsocks Functions%s:
-------- ---------

  # %sName       Description%s
  - ----       -----------
  %s1. %s   setups%s: Installs dnsmasq, squid, privoxy and tor (also set configs).
  %s2. %s   vanish%s: Start anonymizing the system through tor network.
  %s3. %s   status%s: Check if using tor (Show if all anononimty services are up and running).
  %s4. %s    torip%s: Check current tor IP address.
  %s5. %s   chains%s: View traffic logs from squid, privoxy, to tor.
  %s6. %s   reload%s: Completely restart torsocks & connect to a different exit-node.
  %s7. %s exitnode%s: Connect to a different exit-node.
  %s8. %s  restore%s: Backup and resets Iptables to default.
  %s9. %s     stop%s: Get back to the surface-web.

%sex. %s%susage%s:
--  ------

    set function vanish
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTorsocks() {
    fmt.Printf(`
       %sName%s: torsocks
   %sFunction%s: src/securities/torsocks
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: Yes
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024-03-14

%sProvided by%s: <rojahsmontari@gmail.com>
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Distros

%sBasic options%s:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  function     none             yes       The function to execute. ex. -> (setups, vanish, status, torip, chains, reload, exitnode, restore and stop)

%sex. %s%susage%s:
--  ------
  set function start
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  Torsocks is a tool that strictly configures Iptables, Tor, Dsnsmasq, Privoxy & Squid to work together in order to completely anonimize your system through Tor network.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
    ListTorsocksFunctions()
}


func TorsocksOptions() {
    fmt.Printf(`
%sModule options %s(src/securities/torrsocks_setup.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  function       none             yes       The function to execute. ex.      -> (setups, start, exitnode, status, ipaddress, restore, reload, chains, stop)
  run            none             yes       To execute the function. Alias to -> (start, execute, exec, launch).

%sSupported Distros%s:
-----------------

   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  ------
  set function start
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func HelpInfoTorsocksSetups() {
    fmt.Printf(`
       Name: setups
   Function: src/torsocks
   Platform: Windows
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  setups         none             Yes       Installs tor, dnsmasq, squid, privoxy & also set configs

Description:
  This Function will install dnsmasq, squid, privoxy and tor. It will (also set configs) so that all your local traffick will go 
  through privoxy > squid > then tor network. It is done with great care and integrity for super securities.

`)
}

func HelpInfoTorsocksVanish() {
    fmt.Printf(`
       Name: start
   Function: src/torsocks
   Platform: Windows
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  start          none             Yes       Start anonymizing

Description:
  This Function will start services like changemacc to change maccadress in a random way then start dnsmasq, squid, privoxy and tor. 
  It will (also set configs) so that all your local traffick will go through privoxy > squid > then tor network.

`)
}

func HelpInfoTorsocksExitnode() {
    fmt.Printf(`
       Name: exitnode
   Function: src/torsocks
   Platform: Windows
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  exitnode       none             Yes       Shuffle exit-node

Description:
  This Function will shufle the exit nodes to new ones. If you see your nrtwork is slow, this Function can help to find a fast one.

`)
}

func HelpInfoTorsocksStatus() {
    fmt.Printf(`
       Name: status
   Function: src/torsocks
   Platform: Windows
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  status         none             Yes       Check if using tor (Show if all anononimty services are up and running)

Description:
  This Function will query the system to see if macchanger, dnsmasq, squid, privoxy and tor are working correctly and if
  all traffic goes through privoxy > squid > then tor network. It is done with great care and integrity for super securities.

`)
}

func HelpInfoTorsocksIpaddress() {
    fmt.Printf(`
       Name: ipaddress
   Function: src/torsocks
   Platform: Windows
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  ipaddress      none             Yes       Check current external IP address

Description:
  This Function will check for your external IP. It querries tor website for your gateway IP. If your system's proxy is correctly
  configured, then you will get a congratulation mesage from tor website.

`)
}

func HelpInfoTorsocksRestore() {
    fmt.Printf(`
       Name: restore
   Function: src/torsocks
   Platform: Windows
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  restore        none             Yes       Back up & resets Iptables to default

Description:
  This Function will restore your Iptables to default. If the Function was killed instantly and IPTABLES were not set as intended, this Function
  will help you fix the lack off internet connection.

`)
}

func HelpInfoTorsocksReload() {
    fmt.Printf(`
       Name: reload
   Function: src/torsocks
   Platform: Windows
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  reload         none             Yes       Restart torsocks & connect to a different exit-node

Description:
  For a better security, this Function enables you to shuffle into different exit nodes in a tor proxy network.

`)
}

func HelpInfoTorsocksChains() {
    fmt.Printf(`
       Name: chains
   Function: src/torsocks
   Platform: Windows
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  chains         none            Yes       View traffic logs from squid, privoxy, to tor

Description:
  This Function will querry  /var/log/privoxy/log to follow all logs living your system through squid, privoxy to tor.

`)
}

func HelpInfoTorsocksStop() {
    fmt.Printf(`
       Name: stop
   Function: src/torsocks
   Platform: Windows
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  stop           none               Yes       Exit and Get back to the surface-web

Description:
  This Function will stop macchanger, dnsmasq, squid, privoxy, tor then restore Iptables, and other edited configs back to default.
  This will live your system as it was without proxies and anyone can intercept your connections even your ISP.

`)
}

func ListInternalFunctions() {
    fmt.Printf(`
%setworks Functions%s:
-------- ---------

  # %sName       Description%s
  - ----       -----------
  %s1. %sdiscover%s:  Discover all devices connected to the network. Lets you locate targets.
  %s2. %sportscan%s:  Get open ports on the target you have set.
  %s3. %svulnscan%s:  Perform vulnerbility scan on open ports of the target you have set.
  %s4. %senumscan%s:  Digg for S.M.B deep information on the target set.
  %s5. %ssmbexplo%s:  Launch known vulnerbility exploits on the target's S.M.B services.
  %s6. %spsniffer%s:  Sniff all Packets from connected devices to the router(Perform M.I.T.M).
  %s7. %sresponder%s: Start Killer Responder that configs all required fields to get you a reverse shell on windows. Supports IPv6.
  %s8. %sbeefninja%s: Launch Beef-xss and Bettercap/ Ettercp For effective (M.I.B attacks).
  %s9. %sxsshooker%s: Get Shell through XSS Injection to packets in the wire.(To come).

%sex. %s%susage%s:
--  ------

    set function discover
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoNetworks() {
    fmt.Printf(`
       %sName%s: Networks
   %sFunction%s: src/internals
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Distros

%sDescription%s:
-----------
  This is the Network module that contains all internal networks attacks functions.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
    ListInternalFunctions()
}

func NetworksOptions() {
    fmt.Printf(`
%sModule options %s(src/internals/networks_pentest.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  MODE           single           yes       Kind of attack to perform (single or all) single will attack single rhost, all for entire subnetmask.
  IFACE          eth0             yes       Interface to use for penetration testing.
  RHOST          none             yes       Alias to (RHOSTS, TARGET, TARGETS) The target to perform functions on.
  PASSWD         Jesus            yes       The password to set for beef-xss login page with the default user:beef
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed if you need to use (responder, beefninja and xsshooker to handle reverse calls.
  GATEWAY        ->               yes       %sDefault%s: %s%s%s. The default roouter ipaddres. Will be needed when running functions like (beefninja).
  SPOOFER        ettercap         yes       The tool to be used for spoofing target host to our domain. the other (bettercap).
  PROXIES        none             no        Just incase you want to run your traffic through proxies. -> (discover, portscan, vulnscan, enumscan, smbexpl, psniffer, responder, beefninja, xsshooker).
  FAKEDNS        *                yes       Dns to be resolved when performing dns spoof. Mainly needed when beefninja function is at hand.
  FUNCTION       none             yes       The function you want network module to perform. ex. (portscan, vulnscan, enumscan, smbexpl, psniffer, responder, beefninja).

%sSupported Distros%s:
-----------------

   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  ------
  set function start
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC,       bcolors.BOLD, bcolors.ENDC, bcolors.YELLOW, Lhost, bcolors.ENDC,           bcolors.BOLD, bcolors.ENDC, bcolors.YELLOW, Gateway, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func HelpInfoDiscover() {
    fmt.Printf(`
       Name: discover
   Function: src/internals
   Platform: All
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  1 discover     none             No        Discover Targets

Description:
  This Function will scann for all connected devices in the network given using bettercap then arrange the targets
  in a table for you to select one to attack further

`)
}

func HelpInfoPortScan() {
    fmt.Printf(`
       Name: portscan
   Function: src/internals
   Platform: All
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  2 portscan     none             No        Port discovery on the target

Description:
  This Function will scan all open ports of the target to reveal open ports. The tool used is nmap
  command = nmap -p- Target

`)
}

func HelpInfoVulnScan() {
    fmt.Printf(`
       Name: vulnscan
   Function: src/internals
   Platform: All
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  2 vulnscan     none             No        Vuln' Scan the Target

Description:
  This Function will Scan for known vulnerbility on the target that may be an easy win

`)
}

func HelpInfoEnumScan() {
    fmt.Printf(`
       Name: enumscan
   Function: src/internals
   Platform: All
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  4 enumscan     none             No        S.M.B enumration on the target

Description:
  This Function will Automate active directory ports, scan them and check if there are vulnerbilities and shares that could be exploited further

`)
}

func HelpInfoExplScan() {
    fmt.Printf(`
       Name: smbexplo
   Function: src/internals
   Platform: All
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  5 smbexplo     none             No        Exploit S.M.B on target

Description:
  This Function will Scout the target active directory services and try to auto exploit them

`)
}

func HelpInfoPsniffer() {
    fmt.Printf(`
       Name: psniffer
   Function: src/internals
   Platform: All
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  6 psniffer     none             No        Sniff Internal Packets

Description:
  This Function will sniff all traffic from the connected devices to the router showing you on terminal. From tcp, udp, http and https

`)
}

func HelpInfoResponder() {
    fmt.Printf(`
       Name: responder
   Function: src/internals
   Platform: All
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  7 responder    none             No        Lunch Responder with IPv6 Support

Description:
  This Function will Launch reponder asking for your LHOST, Configuring Wpadscript and weponizing it self. Attack supports alot of windows recent version

`)
}

func OptionsResponder() {
    fmt.Printf(`
%sModule Options %s(src/networks/responder.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  RHOST          none             yes       Target to attack. Host's ipadress.
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  PROXIES        none             no        Just incase you want to run your traffic through proxies.

%sSupported Distros%s:
-----------------

   Id  Name
   --  ----
   0   All Browsers

%sex. %s%susage%s:
--  ------
  set LHOST 127.0.0.1
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func HelpInfoBeefNinja() {
    fmt.Printf(`
       Name: beefninja
   Function: src/internals
   Platform: All
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  8 beefninja    none             No        Start Beef & Bettercap For M.I.B

Description:
  This Function will Launch a Combination of both beef and bettercap in a unique way to inject hook.js in either one or all targets. All settings are done for you

`)
}

func OptionsBeefNinja() {
    fmt.Printf(`
%sModule Options %s(src/networks/beefninja.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  RHOST          none             yes       Target to attack. Host's ipadress.
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  PROXIES        none             no        Just incase you want to run your traffic through proxies.
  SPOOFER        ettercap         yes       Tool to be used to spoof dns and repond to them. ex. (bettercap)

%sSupported Distros%s:
-----------------

   Id  Name
   --  ----
   0   All Browsers

%sex. %s%susage%s:
--  ------
  set LHOST 127.0.0.1
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func HelpInfoXssHoocker() {
    fmt.Printf(`
       Name: xsshooker
   Function: src/internals
   Platform: All
       Arch: x64, x86
 Privileged: Yes
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>       0

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  9 xsshooker    none             No        Get Shell through XSS Injection

Description:
  This Function will try to Get you a revers Shell through XSS Injection. Still Working on this Option.

`)
}

func OptionsXssHoocker() {
    fmt.Printf(`
%sModule Options %s(src/internals/xsshooker.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  RHOST          none             yes       Target to attack. Host's ipadress.
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  PROXIES        none             no        Just incase you want to run your traffic through proxies.

%sSupported Distros%s:
-----------------

   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  ------
  set LHOST 127.0.0.1
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func ListExploitsFunctions() {
    fmt.Printf(`
%setworks Functions%s:
-------- ---------

  # %sName       Description%s
  - ----       -----------
  %s1. %s      ghost%s: It is a giant powershell evasion tool that beats almost all AVS. Try it out you will love it.
  %s2. %s     shellz%s: Supports all distro reverse shells generation that supports both tcp, https & https connection. Launches variety of listeners.
  %s3. %s   listener%s: Launch any of your disearable c2 with costumized LHOST, LPORT etc.
  %s4. %s   androrat%s: It is another Android C2. It is cool but only works on android 4 to 9. Suppoorts android < 14 but less functions.
  %s5. %s  teardroid%s: Andoird C2. Needs alot of online configuration but the best for now.
  %s6. %s  blackjack%s: It is a tool derived from villain. It supports both tcp, http & https reverse shells. Supports evasions & bypasses almost all avs.
  %s7. %s  hoaxshell%s: A Killer FUD that bypasses most avs.
  %s8. %s noisemaker%s: Generates undetected backdoor with embeded hoaxreverse shell. Has .dll persistent mechanisims.
  %s9. %scodebreaker%s: Just like noisemake but the persisten mechanisim is through regestry keys.

%sex. %s%susage%s:
--  ------

    set function discover
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoExploits() {
    fmt.Printf(`
       %sName%s: Exploits
   %sFunction%s: src/exploits
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Distros

%sDescription%s:
-----------
  This is the Exploits module that contains all c2, backdoors and obfsicatorsfunctions.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
    ListExploitsFunctions()
}

func ExploitsOptions() {
    fmt.Printf(`
%sModule Options %s(src/exploits/backdoor_pentest.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  C2             blackjack        yes       The default c2 to handle your call back connections. (ncat, hoaxshell, metasploit .etc)
  ICON           vlc              yes       Icon to be used while generating backdoors using (noisemakers and codebreakers)
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  BUILD          africana_malware yes       Output name of the backdoor to be generated. Neede when using(androrat, noisemaker and codebreaker)
  SCRIPT         none             yes       Your powershell script location to be opfsicated. Mostly neede when using (ghos).
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROXIES        none             no        Just incase you want to run your traffic through proxies.
  PROTOCOL       tcp              yes       The kind of protocol to be use while communicating to your host machine. (tcp, http or https).
  FUNCTION       none             yes       The function you want network module to perform. ex. (ghost, shellz, listene, androrat, teardroid, blackjack, hoaxshell, noisemaker, codebreaker).

%sSupported Distros%s:
-----------------

   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  ------
  set function start
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func HelpInfoBlackJack() {
    fmt.Printf(`
       %sName%s: blackjack
   %sFunction%s: src/exploits/blackajck.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s: t3l3machus
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool derived from villain framework. It supports both tcp, http & https reverse shells. It has inbuild evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}


func HelpInfoShellz() {
    fmt.Printf(`
       %sName%s: shellz
   %sFunction%s: src/exploits/shellz.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s: sandres
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoHoaxShell() {
    fmt.Printf(`
       %sName%s: hoaxshell
   %sFunction%s: src/exploits/hoaxshell.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s: t3l3machus
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool like villein or blackjack that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoNoiseMaker() {
    fmt.Printf(`
       %sName%s: noisemaker
   %sFunction%s: src/exploits/noisemaker.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s: t3l3machus
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  ICON           vlc              yes       The icon to use to disguise your backdoor with.
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoHavoc() {
    fmt.Printf(`
       %sName%s: havoc
   %sFunction%s: src/exploits/havoc.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s: t3l3machus
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  ICON           vlc              yes       The icon to use to disguise your backdoor with.
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTearNdroid() {
    fmt.Printf(`
       %sName%s: tearndroid
   %sFunction%s: src/exploits/teandroid.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Androids

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func BlackJackOptions() {
    fmt.Printf(`
%sModule Options %s(src/exploits/blackjack_c2.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  SCRIPT         none             yes       Your powershell script location to be opfsicated. Mostly neede when using (ghos).
  PROXIES        none             no        Just incase you want to run your traffic through proxies.
  PROTOCOL       tcp              yes       The kind of protocol to be use while communicating to your host machine. (tcp, http or https).

%sSupported Distros%s:
-----------------

   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  ------
  set LHOST 127.0.0.1
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func HelpInfoChameLeon() {
     fmt.Printf(`
       %sName%s: chameleon
   %sFunction%s: src/exploits/chameleon.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s: t3l3machus
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  SCRIPT         none             yes       Full location of the powershell script to be obfsicated.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoGhost() {
    fmt.Printf(`
       %sName%s: ghost
   %sFunction%s: src/exploits/ghost.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  SCRIPT         none             yes       Powershell script to obfsicate inorder to bypass AVs.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoSeaShell() {
    fmt.Printf(`
       %sName%s: seashell
   %sFunction%s: src/exploits/seashell.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All IOS

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoListener() {
    fmt.Printf(`
       %sName%s: listener
   %sFunction%s: src/exploits/listener.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoCodeBreaker() {
    fmt.Printf(`
       %sName%s: codebreaker
   %sFunction%s: src/exploits/noisemaker.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  ICON           vlc              yes       The icon to use to disguise your backdoor with.
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTearDroid() {
    fmt.Printf(`
       %sName%s: tearndroid
   %sFunction%s: src/exploits/tearndroid.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Androids

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoAndroRat() {
    fmt.Printf(`
       %sName%s: androrat
   %sFunction%s: src/exploits/androrat.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}


func AndroRatOptions() {
    fmt.Printf(`
%sModule Options %s(src/exploits/androrat_c2.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  BUILD          africana_malware yes       Output name of the backdoor to be generated. Neede when using(androrat, noisemaker and codebreaker)
  SCRIPT         none             yes       Your powershell script location to be opfsicated. Mostly neede when using (ghos).
  OUTPUT         ->               yes       %sDefault%s: %s%s%s%s. Location you want your generated backdoor to be placed.
  PROXIES        none             no        Just incase you want to run your traffic through proxies.
  PROTOCOL       tcp              yes       The kind of protocol to be use while communicating to your host machine. (tcp, http or https).

%sSupported Distros%s:
-----------------

   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  ------
  set LHOST 127.0.0.1
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func ListWirelessFunctions() {
    fmt.Printf(`
%setworks Functions%s:
-------- ---------

  # %sName       Description%s
  - ----       -----------
  %s1. %swifite      %s: Wifite is a tool to audit WEP or WPA encrypted wireless networks.
  %s2. %sfluxion     %s: Fluxion is a tool to audit WEP or WPA encrypted wireless networks. Only manual option is supported by now.
  %s3. %sbettercap   %s: Bettercap is a tool to audit Internal network & wirekless network like, WEP or WPA encrypted wireless networks.
  %s4. %sairgeddon   %s: Airgeddon Fluxion is a tool to audit WEP or WPA encrypted wireless networks. Only manual option is supported by now.
  %s5. %swifipumpkin %s: wifipumpkin - Is a Powerful framework for rogue access point attack. This option run automated mode directly.
  %s6. %sKillerpukin %s: This option runs wifipumpkin3 in manual mode where africana sets everything for you.
  %s7. %sTo add      %s:
  %s8. %sTo add      %s:
  %s9. %sTo add      %s:

%sex. %s%susage%s:
--  ------

    set function wifite
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoWireless() {
    fmt.Printf(`
       %sName%s: wireless
   %sFunction%s: src/wireless
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSuported Distros%s:
      Id  Name
      --  ----
   -> 0   All Distros

%sDescription%s:
-----------
  This is the wireless module containing all wireless pentesting functions.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
    ListWirelessFunctions()
}

func WirelessOptions() {
    fmt.Printf(`
%sModule Options %s(src/wirelss/wireless_pentest.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  IFACE          ->               yes       %sDefault%s: %s%s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  SCRIPT         none             yes       Your powershell script location to be opfsicated. Mostly neede when using (ghos).
  PROXIES        none             no        Just incase you want to run your traffic through proxies.
  PROTOCOL       tcp              yes       The kind of protocol to be use while communicating to your host machine. (tcp, http or https).

%sSupported Distros%s:
-----------------

   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  ------
  set IFACE wlan0
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ITALIC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func HelpInfoWifite() {
    fmt.Printf(`
       Name: wifite
   Function: src/wireless
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by kimocoder
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   WiFi

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  wifite         -a               yes       Automated attack to all wireless networks.
  wifite         -m               yes       Manual attack to all wireless networks.

Description:
    Wifite is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit.
    This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

`)
}

func HelpInfoFluxion() {
    fmt.Printf(`
       Name: fluxion
   Function: src/wireless
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   WiFi

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  fluxion        none             no        Manual attack to all wireless networks.

Description:
    Fluxion is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit.
    This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

`)
}

func HelpInfoBetterCap() {
    fmt.Printf(`
       Name: bettercap
   Function: src/wireless
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   WiFi

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  bettercap      -a               yes       Automated attack to all wireless networks.
  bettercap      -m               yes       Manual attack to all wireless networks.

Description:
    Bettercap is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit.
    This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

`)
}

func HelpInfoAirGeddon() {
    fmt.Printf(`
       Name: airgeddon
   Function: src/wireless
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by v1s1t0r
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   WiFi

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  airgeddon      none             no        Manual attack to all wireless networks.

Description:
    Airgeddon is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit.
    This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

`)
}

func HelpInfoWifiPumpkin() {
    fmt.Printf(`
       Name: wifipumpkin
   Function: src/wireless
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by P0cL4bs Team
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   WiFi

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  wifipumpkin    -a               yes       Automated attack to all wireless networks.
  wifipumpkin    -m               yes       Manual attack to all wireless networks.

Description:
    Wifipumpkin is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit.
    This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

`)
}

func HelpInfoWifiPumpkin3() {
    fmt.Printf(`
       Name: wifipumpkin
   Function: src/wireless
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by P0cL4bs Team
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   WiFi

Check supported:

Basic options:
  Name           Current Setting  Required  Description 
  ----           ---------------  --------  -----------
  wifipumpkin    -a               yes       Automated attack to all wireless networks.
  wifipumpkin    -m               yes       Manual attack to all wireless networks.

Description:
    Wifipumpkin is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit.
    This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

`)
}

func ListCrackersFunctions() {
    fmt.Printf(`
Crackers Functions:
----------------

  #   Name         Description
  -   ----         -----------

  1 online        : It contains password bruteforcers that attack online services like ssh, ftp, smb, rdp, ldap, http, https etc.
                     SubFunctions: include hydra, rainbow.

  2 offline       : Supports all cracking offline encrypted files like .zip, .rar, .pcap, .ntlm, .MD4, .MD5, SHA1-SHA3, Kerberos etc.
                     SubFunctions: include hashcat, aircrack-ng.

`)
}

func ListPhishersFunctions() {
    fmt.Printf(`
Phishers Functions:
----------------

  #   Name         Description
  -   ----         -----------

  1 gophish       : It is a phishing framework with a Web UI https://127.0.0.1:3333. Africana will launch it for you. Default  is: admin, Default password is: kali-gophish.
  2 goodginx      : It is an advanced phishing framework with insane configurations.Default name evilginx2. Bypasses alot of security features like OTP.
  3 zphisher      : A nice framework with alot of templets. Also bypasses OTP with ngrock support.
  4 blackeye      :
  5 advnphish     :
  6 darkphish     :
  7 shellphish    : Supports otp bypass. Wide range of phishing templets.
  8 setoolkit     : This tool is equiped with alot of social eneneering. Supports cloning of actual websites.
  9 thehackerchoice:This tool creates a templete of your interest imidietly but needs you to start your server and generate a link for phishing.

`)
}

func HelpInfoGophish() {
    fmt.Printf(`
       Name: gophish
   Function: /usr/share/gophish
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  none           _______________  ________  ___________

Description:
  It is a Function that enables the redteamers to perform phising attacks on various bases.

`)
}

func HelpInfoGoodGinx() { 
    fmt.Printf(`
       Name: goodginx
   Function: /usr/share/eginx
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  none           _______________  ________  ___________

Description:
  It is a Function that enables the redteamers to perform phising attacks on various bases.

`)
}

func HelpInfoZphisher() {
    fmt.Printf(`
       Name: zphisher
   Function: /root/.afr3/africana-base/zphisher
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  none           _______________  ________  ___________

Description:
  It is a Function that enables the redteamers to perform phising attacks on various bases.

`)
}

func HelpInfoBlackEye() {
    fmt.Printf(`
       Name: blackeye
   Function: /root/.afr3/africana-base/phishers/blackeye
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  none           _______________  ________  ___________

Description:
  It is a Function that enables the redteamers to perform phising attacks on various bases.

`)
}

func HelpInfoAdvnPhish() {
    fmt.Printf(`
       Name: advnphish
   Function: /root/.afr3/africana-base/phishers/advnphish
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  none           _______________  ________  ___________

Description:
  It is a Function that enables the redteamers to perform phising attacks on various bases.

`)
}

func HelpInfoDarkPhish() { 
    fmt.Printf(`
       Name: darkphish
   Function: /root/.afr3/africana-base/phishers/darkphish
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  none           _______________  ________  ___________

Description:
  It is a Function that enables the redteamers to perform phising attacks on various bases.

`)
}

func HelpInfoShellPhish() {
    fmt.Printf(`
       Name: shellphish
   Function: /root/.afr3/africana-base/phishers/shellphish
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  none           _______________  ________  ___________

Description:
  It is a Function that enables the redteamers to perform phising attacks on various bases.

`)
}

func HelpInfoSetToolkit() {
    fmt.Printf(`
       Name: setoolkit
   Function: /root/.afr3/africana-base/phishers/set
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  none           _______________  ________  ___________

Description:
  It is a Function that enables the redteamers to perform phising attacks on various bases.

`)
}

func HelpInfoTheHackerChoice() {
    fmt.Printf(`
       Name: thehackerchoice
   Function: /root/.afr3/africana-base/phishers/thehackerchoice
   Platform: All
       Arch: x64, x86
 Privileged: No
    License: Africana Framework License(BSD)
       Rank: Insane
  Disclosed: 2024

Provided by:
  Created by _____________
   Edited by r0jahsm0ntar1

Available targets:
      Id  Name
      --  ----
  =>  0   All

Check supported:

Basic options:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  none           _______________  ________  ___________

Description:
  It is a Function that enables the redteamers to perform phising attacks on various bases.

`)
}
