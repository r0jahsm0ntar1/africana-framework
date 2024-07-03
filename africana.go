package main

import (
    "os"
    "fmt"
    "utils"
    "menus"
    "bufio"
    "setups"
    "credits"
    "strings"
    "banners"
    "net/url"
    "runtime"
    "bcolors"
    "butchers"
    "wireless"
    "crackers"
    "phishers"
    "io/ioutil"
    "internals"
    "agreements"
    "scriptures"
    "securities"
    "webattackers"
)

var (
    userInput   string
    userTarget  string
    userTargetx string
    proxyURL    string
)

func sudo() {
    switch runtime.GOOS {
    case "windows":
        userAgreements()
        return
    default:
        if os.Geteuid() != 0 {
            utils.ClearScreen(); banners.Banner(); scriptures.Verse()
            fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🦊africana must be run as root. Try " + bcolors.RED + "sudo africana\n" + bcolors.ENDC)
            return
        }else{
        userAgreements()
        return
        }
    }
}

func userAgreements() {
    filePath := "/root/.africana/agreements/"
    content := "📜 User accepted the terms."
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        utils.ClearScreen(); agreements.Covenant()
        for {
            fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "╰─🦊" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            fmt.Scan(&userInput)
            switch userInput {
                case "y", "Y", "yes", "Yes", "YES":
                    if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
                        fmt.Println("Error creating file:", err)
                        return
                    }
                    err = ioutil.WriteFile(filePath + "covenant.txt", []byte(content), os.ModePerm)
                    if err != nil {
                        fmt.Println("Error writing to file:", err)
                        return
                    }
                    utils.InitiLize(); genesis()
                    return
                case "n", "N", "no", "No", "NO":
                    os.Exit(0)
                default:
                    fmt.Println(bcolors.GREEN + "           (" + bcolors.DARKCYAN + "Choices are [" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + "]" + bcolors.GREEN + ")" + bcolors.ENDC)
                }
            }
    } else {
        utils.InitiLize(); genesis()
        return
    }
}

//1. Install or update africana-framework..(Start here )🩺
func systemSetups() {
    utils.ClearScreen(); banners.Banner(); menus.MenuOne()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🩺" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            africanaFramework()
            return
        case "1":
            utils.ClearScreen(); banners.Banner(); setups.KaliSetups()
            return
        case "2":
            utils.ClearScreen(); banners.Banner(); setups.UbuntuSetups()
            return
        case "3":
            break
        case "4":
            utils.ClearScreen(); banners.Banner(); setups.WindowsSetup()
            return
        case "5":
            utils.ClearScreen(); banners.Banner(); setups.RemoveSetups()
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuOne()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuOne()
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 4" + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//2. System Security Configuration.........(Setup tor &)🎭
func anonymitySetups() {
    utils.ClearScreen(); banners.Banner(); menus.MenuTwo()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🎭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            africanaFramework()
            return
        case "1":
            securities.AnonymitySetups  (           )
        case "2":
            securities.AnonymityStart   (           )
        case "3":
            securities.AnonymityStop    (           )
        case "4":
            securities.AnonymityStatus  (           )
        case "5":
            securities.AnonymityChains  (           )
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuTwo()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuTwo()
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 5" + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//3. Local Network Attack Vectors..........(Mitm, sniff)🐹
func internaltargetInput() {
    utils.ClearScreen(); banners.Banner(); internals.InternalScanner()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.ORANGE + "Select your" + bcolors.PURPLE + "(҂`_´) " + bcolors.RED + "target!!🎯" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐹" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
}

func internalAttackers() {
    banners.Banner(); menus.MenuThree()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐹" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            africanaFramework()
            return
        case "1":
            internaltargetInput         (           )
            utils.ClearScreen()
            banners.Banner(); menus.MenuThree   (  )
        case "2":
            internals.NmapPortscan       (userTarget)
        case "3":
            internals.NmapVulnscan       (userTarget)
        case "4":
            internals.SmbVulnscan        (userTarget)
            internals.SmbMapscan         (userTarget)
            internals.RpcEnumscan        (userTarget)
        case "5":
            internals.SmbVulnscan        (userTarget)
            internals.SmbExploit         (userTarget)
        case "6":
            internals.PacketSniffer      (userTarget)
            internalAttackers()
            return
        case "7":
            internals.PacketsResponder  (           )
        case "8":
            internals.BeefBettercap      (userTarget)
        case "9":
            internals.RpcEnumscan        (userTarget)
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuThree()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuThree()
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//4. Generate Undetectable Backdoors.......(C2 & shells)🐭
func malwareGenerators() {
    utils.ClearScreen(); banners.Banner(); menus.MenuFour()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            africanaFramework()
            return
        case "1":
            butchers.BlackJack  (           )
        case "2":
            butchers.Shellz     (           )
        case "3":
            butchers.PowerJoker (           )
        case "4":
            butchers.MeterPeter (           )
        case "5":
             butchers.Havoc     (           )
        case "6":
            butchers.TearDroid  (           )
        case "7":
            butchers.AndroRat   (           )
        case "8":
            butchers.Chameleon  (           )
        case "9":
            butchers.ChameLeon  (           )
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuFour()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuFour()
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//5. WiFi Attack Vectors...................(Wifite, air)🍍
func wirelesAttackers() {
    utils.ClearScreen(); banners.Banner(); menus.MenuFive()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🍍" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            africanaFramework()
            return
        case "1":
            wireless.WifiteAuto         (           )
        case "2":
            wireless.BettercapAuto      (           )
        case "3":
            wireless.WifiPumpkin3Auto   (           )
        case "4":
            wireless.AirGeddon          (           )
        case "5":
            wireless.AirGeddon          (           )
        case "6":
            wireless.AirGeddon          (           )
        case "7":
            wireless.WifiPumpkin3       (           )
        case "8":
            wireless.AirGeddon          (           )
        case "9":
            wireless.AirGeddon          (           )
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuFive()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuFive()
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//6. Crack Hash, Pcap & Brute Passwords....(Hashcat, jo)🔐
func passwordsCrackers() {
    utils.ClearScreen(); banners.Banner(); menus.MenuSix()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🔐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            africanaFramework()
            return
        //Online crackers
        case "1":
            utils.ClearScreen(); banners.Banner(); menus.MenuSixOne()
            for {
                fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "╰─🔐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0", "e", "E", "exit", "Exit", "EXIT":
                    passwordsCrackers()
                    return
                case "1":
                    crackers.HydraSsh   (           )
                case "2":
                    crackers.HydraFtp   (           )
                case "3":
                    crackers.HydraSmb   (           )
                case "4":
                    crackers.HydraRdp   (           )
                case "5":
                    crackers.HydraLdap  (           )
                case "6":
                    crackers.HydraSmtp  (           )
                case "7":
                    crackers.HydraTelnet(           )
                case "8":
                    crackers.HydraHttps (           )
                case "9":
                    crackers.CyberBrute (           )
                case "99", "m", "M", "menu", "Menu", "MENU":
                    menus.MenuSixOne()
                case "00", "h", "H", "help", "Help", "HELP":
                    menus.HelpMenuSixOne()
                default:
                    fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //Offline crackers
        case "2":
            utils.ClearScreen(); banners.Banner(); menus.MenuSixTwo()
            for {
                fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "╰─🔐" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0", "e", "E", "exit", "Exit", "EXIT":
                    passwordsCrackers   (           )
                    return
                case "1":
                    crackers.AirCrackng (           )
                case "2":
                    crackers.JohnCrackng(           )
                case "8":
                    crackers.HashBuster (           )
                case "99", "m", "M", "menu", "Menu", "MENU":
                    menus.MenuSixOne()
                case "00", "h", "H", "help", "Help", "HELP":
                    menus.HelpMenuSixTwo()
                default:
                    fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.HelpMenuSixTwo()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuSix()
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 2" + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//7. Social-Engineering Attacks............(Gophish, gi)🪝
func credsPhishers() {
    utils.ClearScreen(); banners.Banner(); menus.MenuSeven()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🪝" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            africanaFramework()
            return
        case "1":
            phishers.GoPhish    (           )
        case "2":
            phishers.GoodGinx   (           )
        case "3":
            phishers.ZPhisher   (           )
        case "4":
            phishers.SetoolKit  (           )
        case "5":
            phishers.AnonPhisher(           )
        case "6":
            phishers.CyberPhish (           )
        case "7":
            utils.UpsentTools   (           )
        case "8":
            utils.UpsentTools   (           )
        case "9":
            utils.UpsentTools   (           )
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuSeven()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuSeven()
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//8. Website Attack Vectors................(Proxy, Conf)🎭

func askForProxy() (*url.URL) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "Enter proxy URL🛰️ Eg:" + bcolors.YELLOW + "(http://localhost:80)\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🎭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    proxyStr, _ := reader.ReadString('\n')
    proxyStr = strings.TrimSpace(proxyStr)

    proxyURL, err := url.Parse(proxyStr)
    if err != nil {
        return nil
    }

    if proxyURL.Scheme != "http" && proxyURL.Scheme != "https" && proxyURL.Scheme != "socks5" && proxyURL.Scheme != "socks4" {
        askForProxy()
    }

    return proxyURL
}

func setProxyEnv(proxyURL *url.URL) error {
    if err := os.Setenv("HTTP_PROXY", proxyURL.String()); err != nil {
        return err
    }
    if err := os.Setenv("HTTPS_PROXY", proxyURL.String()); err != nil {
        return err
    }
    // SOCKS5 proxy is not directly supported by environment variables in all tools,
    // but you can use tools like `proxychains` to enable SOCKS5 support.
    return nil
}

//8. Website Attack Vectors................(userTarget, set)🎭
func websiteUserTarget() {
    utils.ClearScreen(); banners.Banner()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "Enter Target:" + bcolors.DARKCYAN + "Either 📡HTTP(S)//: HOSTNAME or IP🌍" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🎯" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userTarget)
    userTargetx = strings.TrimSpace (userTarget)
    userTargetx = strings.TrimPrefix(userTargetx, "http://")
    userTargetx = strings.TrimPrefix(userTargetx, "https://")
    userTargetx = strings.TrimPrefix(userTargetx, "www.")
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "Use proxies:" + bcolors.DARKCYAN + "Eg: SOCK4, SOCKS5, HTTP or HTTPS " + bcolors.YELLOW + "(y/n)\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🎭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "y", "yes", "Y", "Yes", "YES":
            proxyURL := askForProxy()

            if err := setProxyEnv(proxyURL); err != nil {
                fmt.Println("Error setting proxy environment variables:", err)
                return
            }
            return
        case "n", "no", "N", "No", "NO":
            return
        default:
            fmt.Println(bcolors.GREEN + "           (" + bcolors.DARKCYAN + "Choices are [" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + "]" + bcolors.GREEN + ")" + bcolors.ENDC)
        }
    }
}

//8. Website Attack Vectors................(Attack, Phase)🌍
func websitesAttackers() {
    utils.ClearScreen(); banners.Banner(); menus.MenuEight()

    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🪲" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            africanaFramework()
            return
        //1. Start Passive Web recon & Subdomain Enumration.....🌍
        case "1":
            webattackers.WafW00f(userTargetx); webattackers.DnsRecon(userTargetx); webattackers.WhatWeb(userTargetx); webattackers.Nuclei(userTargetx)
        //2. Gather e-mails & subdomain namesfrom public sources🪰
        case "2":
            webattackers.Harvester (userTarget)
        //3. Start Bruteforcing Host's Root Files...............🦑
        case "3":
            webattackers. Gobuster (userTarget)
        //4. Start SQL, XSS & SSRF Detection & Eploitation......💉
        case "4":
            utils.ClearScreen(); banners.Banner(); menus.MenuEightFour()
            for {
                fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "╰─💉" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0", "e", "E", "exit", "Exit", "EXIT":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.SqlmapAuto (userTarget)
                case "2":
                    webattackers.CommixAuto (userTarget)
                    webattackers.XsserAuto  (userTarget)
                    webattackers.KatanaAuto (userTarget)
                case "3":
                    webattackers.SqlmapMan  (           )
                case "4":
                    webattackers.CommixMan  (           )
                    webattackers.XsserMan   (           )
                case "99", "m", "M", "menu", "Menu", "MENU":
                    menus.MenuEightFour()
                case "00", "h", "H", "help", "Help", "HELP":
                    menus.HelpMenuEightFour()
                default:
                    fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 4" + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //5. Launch OWASP Nettacker project MainMenu............🦣
        case "5": 
            utils.ClearScreen(); banners.Banner(); menus.MenuEightFive()
            for {
                fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "╰─🦣" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0", "e", "E", "exit", "Exit", "EXIT":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.NetTacker1 (userTarget)
                case "2":
                    webattackers.NetTacker2 (userTarget)
                case "3":
                    webattackers.NetTacker3 (userTarget)
                case "4":
                    webattackers.NetTacker4 (userTarget)
                case "5":
                    webattackers.NetTacker5 (userTarget)
                case "6":
                    webattackers.NetTacker6 (userTarget)
                case "7":
                    webattackers.NetTacker7 (userTarget)
                case "8":
                    webattackers.NetTacker8 (userTarget)
                case "9":
                    webattackers.NetTacker9(           )
                case "99", "m", "M", "menu", "Menu", "MENU":
                    menus.MenuEightFive()
                case "00", "h", "H", "help", "Help", "HELP":
                    menus.HelpMenuEightFive()
                default:
                    fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //6. Jok3r v3.5 Insane Mult Reconing Engine Main Menu...🃏
        case "6":
            utils.ClearScreen(); banners.Banner(); menus.MenuEightSix()
            for {
                fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "╰─🃏" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0", "e", "E", "exit", "Exit", "EXIT":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.Jok3r1(           )
                case "2":
                    webattackers.Jok3r2(           )
                case "3":
                    webattackers.Jok3r3(           )
                case "4":
                    webattackers.Jok3r4(           )
                case "5":
                    webattackers.Jok3r5 (userTarget)
                case "6":
                    webattackers.Jok3r6 (userTarget)
                case "7":
                    webattackers.Jok3r7 (userTarget)
                case "8":
                    webattackers.Jok3r8(           )
                case "9":
                    webattackers.Jok3r9(           )
                case "99", "m", "M", "menu", "Menu", "MENU":
                    menus.MenuEightSix()
                case "00", "h", "H", "help", "Help", "HELP":
                    menus.HelpMenuEightSix()
                default:
                    fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //7. Osmedeus Next Generation Workflow Engine Main Menu.🐨
        case "7":
            utils.ClearScreen(); banners.Banner(); menus.MenuEightSeven()
            for {
                fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "╰─🐨" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0", "e", "E", "exit", "Exit", "EXIT":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.Osmedeus1(           )
                case "2":
                    webattackers.Osmedeus2 (userTarget)
                case "3":
                    webattackers.Osmedeus3 (userTarget)
                case "4":
                    webattackers.Osmedeus4 (userTarget)
                case "5":
                    webattackers.Osmedeus5 (userTarget)
                case "6":
                    webattackers.Osmedeus6 (userTarget)
                case "7":
                    webattackers.Osmedeus7 (userTarget)
                case "8":
                    webattackers.Osmedeus8(           )
                case "9":
                    webattackers.Osmedeus9(           )
                case "99", "m", "M", "menu", "Menu", "MENU":
                    menus.MenuEightSeven()
                case "00", "h", "H", "help", "Help", "HELP":
                    menus.HelpMenuEightSeven()
                default:
                    fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //8. Ufonet Next Generation DDOS Tool Main Menu.........🦠
        case "8":
            utils.ClearScreen(); banners.Banner(); menus.MenuEightEight()
            for {
                fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
                fmt.Printf(bcolors.BLUE + "╰─🦠" + bcolors.GREEN + "❯ " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0", "e", "E", "exit", "Exit", "EXIT":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.Ufonet1(           )
                case "2":
                    webattackers.Ufonet2(           )
                case "3":
                    webattackers.Ufonet3 (userTarget)
                case "4":
                    webattackers.Ufonet4 (userTarget)
                case "5":
                    webattackers.Ufonet5 (userTarget)
                case "6":
                    webattackers.Ufonet6 (userTarget)
                case "7":
                    webattackers.Ufonet7(           )
                case "8":
                    webattackers.Ufonet8(           )
                case "9":
                    webattackers.Ufonet9 (userTarget)
                case "99", "m", "M", "menu", "Menu", "MENU":
                    menus.MenuEightEight()
                case "00", "h", "H", "help", "Help", "HELP":
                    menus.HelpMenuEightEight()
                default:
                    fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //9. Launch Heavy Automation Attacks On The Host........🤖
        case "9":
            webattackers.WafW00f(userTargetx); webattackers.WhatWeb(userTargetx); webattackers.DnsRecon(userTargetx); webattackers.Nuclei(userTargetx); webattackers.SeekOlver (userTarget); webattackers.Gobuster (userTarget); webattackers.Osmedeus3 (userTarget); webattackers.ParamSpider (userTarget); webattackers.SqlmapAuto (userTarget); webattackers.CommixAuto (userTarget); webattackers.KatanaAuto (userTarget); webattackers.XsserAuto (userTarget); webattackers.Nikto (userTarget); webattackers.Uniscan (userTarget)
        case "99", "m", "M", "menu", "Menu", "MENU":
            menus.MenuEight()
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuEight()
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//9. Help, Credits, Tricks and About.......(🕊  ︻╦╤─JC❤sU)
func creditsGivers() {
    utils.ClearScreen(); banners.Banner(); scriptures.Verse(); credits.Contributors(); credits.Developer()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + " Enter 0, e, E, exit, Exit or EXIT " + bcolors.ENDC + "to go back" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─📚" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            africanaFramework()
            return
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please select " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to Go to menu" + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//0. Exit africana-framework...............(Try option 00)
func scriptureNarators() {
    utils.ClearScreen(); scriptures.TheScriptures(); scriptures.CommandMents()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─✍🏼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            africanaFramework()
            return
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please select " + bcolors.YELLOW + "🦝00. or" + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 & Go back " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//Africana-Framework ......................(The rolling 9)
func africanaFramework() {
    utils.ClearScreen(); banners.Banner(); menus.MenuZero()
    for {
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🦊" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0", "e", "E", "exit", "Exit", "EXIT":
            utils.ClearScreen   (           )
            banners.Banner      (           )
            scriptures.Verse    (           )
            return
        case "1":
            systemSetups        (           )
            return
        case "2":
            anonymitySetups     (           )
            return
        case "3":
            internaltargetInput (           )
            internalAttackers   (           )
            return
        case "4":
            malwareGenerators   (           )
            return
        case "5":
            wirelesAttackers    (           )
            return
        case "6":
            passwordsCrackers   (           )
            return
        case "7":
            credsPhishers       (           )
            return
        case "8":
            websiteUserTarget   (           )
            websitesAttackers   (           )
            return
        case "9":
            creditsGivers       (           )
            return
        case "99":
            scriptureNarators   (           )
        case "00", "h", "H", "help", "Help", "HELP":
            menus.HelpMenuZero()
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please try option " + bcolors.GREEN + "00 " + bcolors.YELLOW  + "or " + bcolors.BLUE + "(" + bcolors.DARKCYAN + "0 to 9" + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//...........................................(Main runner..)
func genesis() {
    if len(os.Args) < 2 {
        banners.Banner      (           )
        menus.HelpMenuMain  (           )
        return
    }
    command := os.Args[1]
    switch command {
    case "-v",           "--version":
        banners.Banner      (           )
        return
    case "-u",            "--update":
        systemSetups        (           )
        return
    case "-0", "-a",        "--auto":
        africanaFramework   (           )
        return
    case "-1", "-i",     "--install":
        systemSetups        (           )
        return
    case "-2", "-t",    "--anonsurf":
        anonymitySetups     (           )
        return
    case "-3", "-n",    "--networks":
        internaltargetInput (           )
        internalAttackers   (           )
        return
    case "-4", "-m",    "--malwares":
        malwareGenerators   (           )
        return
    case "-5", "-w",    "--wireless":
        wirelesAttackers    (           )
        return
    case "-6", "-p",    "--crackers":
        passwordsCrackers   (           )
        return
    case "-7", "-f",    "--phishers":
        credsPhishers       (           )
        return
    case "-8", "-s",    "--websites":
        websiteUserTarget   (           )
        websitesAttackers   (           )
        return
    case "-9", "-c",     "--credits":
        creditsGivers       (           )
        return
    case "-99", "-b", "--scriptures":
        scriptureNarators   (           )
        return
    case "-00", "-h",       "--help":
        banners.Banner      (           )
        menus.HelpMenuMain  (           )
    default:
        banners.Banner      (           )
        menus.HelpMenuMain  (           )
    }
}

func main() {
     sudo()
}
