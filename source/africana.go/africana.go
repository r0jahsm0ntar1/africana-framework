package main

import (
    "os"
    "fmt"
    "utils"
    //"os/signal"
    "io/ioutil"
    //"syscall"
    "subprocess"
    "bcolors"
    "agreements"
    "banners"
    "menus"
    "internals"
    "butchers"
    "wireless"
    "crackers"
    "phishers"
    "webattackers"
)

var userInput, userTarget string

func screenClear() {
    subprocess.Popen("clear")
}

func userAgreements() {
    filePath := "/root/.africana/agreements/"
    content := "User accepted"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        screenClear(); agreements.Covenant()
        for {
            fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana: " + bcolors.DARKCYAN + "Do you agree to the terms of service [y/n]:" + bcolors.GREEN + ")# " + bcolors.ENDC)
            fmt.Scan(&userInput)
            switch userInput {
                case "y", "Y", "yes", "YES":
                    if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
                        fmt.Println("Error creating file:", err)
                        return
                    }
                    err = ioutil.WriteFile(filePath + "covenant.txt", []byte(content), os.ModePerm)
                    if err != nil {
                        fmt.Println("Error writing to file:", err)
                        return
                    }
                    utils.WordLists(); africanaFramework()
                    return
                case "n", "N", "no", "NO":
                    os.Exit(0)
                default:
                    fmt.Print(bcolors.GREEN + "           (" + bcolors.DARKCYAN + "Choices are [" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + "]" + bcolors.GREEN + ")" + bcolors.ENDC)
                }
            }
    } else {
        africanaFramework()
        return
    }
}

//1. Install or update africana-framework..(Start here )🐞
func africanaSetups() {
    fmt.Println("agreements.....")
}

//2. System Security Configuration.........(Setup tor &)🐈
func anonymitySetups() {
    fmt.Println("agreements.....")
}

//3. Local Network Attack Vectors..........(Mitm, sniff)🐹
func internalUserTarget() {
    screenClear(); banners.BannerEight(); internals.InternalScanner()
    fmt.Printf(bcolors.BLUE + "\n" + bcolors.YELLOW + `¯\_(ツ)_/¯ ` + bcolors.GREEN + "Select target IP🎯 to be " + bcolors.RED + "Attacked!! \n" + bcolors.ENDC)
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
}

func internalAttackers() {
    screenClear(); banners.Banner(); menus.MenuThree()
    for {
        fmt.Printf(bcolors.GREEN + "\n" + `¯\_(ツ)_/¯ Attacking!! ` + bcolors.RED + "︻╦╤─ ─ 🎯 " + bcolors.YELLOW + "%s\n" , userTarget + bcolors.ENDC)
        fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            africanaFramework()
            return
        case "1":
            internalUserTarget(); internalAttackers()
        case "2":
            internals.NmapPortscan    (userTarget)
        case "3":
            internals.NmapVulnscan    (userTarget)
        case "4":
            internals.SmbVulnscan     (userTarget)
            internals.SmbMapscan      (userTarget)
            internals.RpcEnumscan     (userTarget)
        case "5":
            internals.SmbVulnscan     (userTarget)
            internals.SmbExploit      (userTarget)
        case "6":
            internals.PacketSniffer   (userTarget)
        case "7":
            internals.PacketsResponder()
        case "8":
            internals.BeefBettercap   (userTarget)
        case "9":
            internals.RpcEnumscan     (userTarget)
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 9 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//4. Generate Undetectable Backdoors.......(C2 & shells)🐭
func malwareGenerators() {
    screenClear()
    for {
        banners.Banner(); menus.MenuFour()
        fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            africanaFramework()
            return
        case "1":
            butchers. BlackJack()
        case "2":
            butchers.    Shellz()
        case "3":
            butchers.PowerJoker()
        case "4":
            butchers.MeterPeter()
        case "5":
             butchers.    Havoc()
        case "6":
            butchers. TearDroid()
        case "7":
            butchers.  AndroRat()
        case "8":
            butchers. Chameleon()
        case "9":
            butchers. ChameLeon()
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 9 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//5. WiFi Attack Vectors...................(Wifite, air)🦝
func wirelesAttackers() {
    screenClear()
    for {
        banners.Banner(); menus.MenuFive()
        fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            africanaFramework()
            return
        case "1":
            wireless.WifiteAuto()
        case "2":
            wireless.BettercapAuto()
        case "3":
            wireless.WifiPumpkin3Auto()
        case "4":
            wireless.AirGeddon()
        case "5":
            wireless.AirGeddon()
        case "6":
            wireless.AirGeddon()
        case "7":
            wireless.WifiPumpkin3()
        case "8":
            wireless.AirGeddon()
        case "9":
            wireless.AirGeddon()
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 9 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//6. Crack Hash, Pcap & Brute Passwords....(Hashcat, jo)🐙
func passwordsCrackers() {
    screenClear()
    for {
        banners.Banner(); menus.MenuSix()
        fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            africanaFramework()
            return
        //Online crackers
        case "1":
            screenClear()
            for {
                banners.Banner(); menus.MenuSixOne()
                fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0":
                    passwordsCrackers()
                    return
                case "1":
                    crackers.HydraSsh()
                case "2":
                    crackers.HydraFtp()
                case "3":
                    crackers.HydraSmb()
                case "4":
                    crackers.HydraRdp()
                case "5":
                    crackers.HydraLdp()
                case "6":
                    crackers.HydraSmp()
                case "7":
                    crackers.HydraTnt()
                case "8":
                    crackers.HydraHtp()
                case "9":
                    crackers.CyberBrute()
                default:
                    fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 9 " + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //Offline crackers
        case "2":
            screenClear()
            for {
                banners.Banner(); menus.MenuSixTwo()
                fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0":
                    passwordsCrackers()
                    return
                case "1":
                    crackers.AirCrackng()
                case "2":
                    crackers.JohnCrackng()
                case "8":
                    crackers.HashBuster()
                default:
                    fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 9 " + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//7. Social-Engineering Attacks............(Gophish, gi)🧪
func credsPhishers() {
    screenClear()
    for {
        banners.Banner(); menus.MenuSeven()
        fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            africanaFramework()
            return
        case "1":
            phishers.    GoPhish()
        case "2":
            phishers.   GoodGinx()
        case "3":
            phishers.   ZPhisher()
        case "4":
            phishers.  SetoolKit()
        case "5":
            phishers.AnonPhisher()
        case "6":
            phishers. CyberPhish()
        case "7":
            phishers. UpsenTools()
        case "8":
            phishers. UpsenTools()
        case "9":
            phishers. UpsenTools()
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 9 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//8. Website Attack Vectors................(Osmedeus, j)🌍
func websiteUserTarget() {
    screenClear(); banners.Banner()
    fmt.Print(bcolors.BLUE + "\n Enter Your Target: ~> " + bcolors.RED + "Either HTTP(S)//: HOSTNAME or IP 🎯\n" + bcolors.ENDC)
    fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
    fmt.Scan(&userTarget)
}

func websitesAttackers() {
    screenClear(); banners.Banner(); menus.MenuEight()
    for {
        fmt.Printf(bcolors.GREEN + `¯\_(ツ)_/¯ Attacking!! ` + bcolors.RED + "︻╦╤─ ─ 🎯 " + bcolors.YELLOW + "%s\n" , userTarget + bcolors.ENDC)
        fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            africanaFramework()
            return
        //1. Start Passive Web recon & Subdomain Enumration.....🌍
        case "1":
            webattackers.WafW00f(userTarget); webattackers.DnsRecon(userTarget); webattackers.WhatWeb(userTarget); webattackers.Nuclei(userTarget)
        //2. Gather e-mails & subdomain namesfrom public sources🪰
        case "2":
            webattackers.Harvester(userTarget)
        //3. Start Bruteforcing Host's Root Files...............🚀
        case "3":
            webattackers. Gobuster(userTarget)
        //4. Start SQL, XSS & SSRF Detection & Eploitation......🐌
        case "4":
            screenClear(); banners.Banner();   menus.MenuEightFour()
            for {
                fmt.Printf(bcolors.GREEN + `¯\_(ツ)_/¯ Attacking!! ` + bcolors.RED + "︻╦╤─ ─ 🎯 " + bcolors.YELLOW + "%s\n" , userTarget + bcolors.ENDC)
                fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.SqlmapAuto(userTarget)
                case "2":
                    webattackers.CommixAuto(userTarget); webattackers.XsserAuto(userTarget); webattackers.KatanaAuto(userTarget)
                case "3":
                    webattackers.SqlmapMan()
                case "4":
                    webattackers.CommixMan(); webattackers.XsserMan()
                default:
                    fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 4 " + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //5. Launch OWASP Nettacker project MainMenu............🦣
        case "5": 
            screenClear(); banners.Banner();  menus. MenuEightFive()
            for {
                fmt.Printf(bcolors.GREEN + `¯\_(ツ)_/¯ Attacking!! ` + bcolors.RED + "︻╦╤─ ─ 🎯 " + bcolors.YELLOW + "%s\n" , userTarget + bcolors.ENDC)
                fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.SqlmapAuto(userTarget)
                case "2":
                    webattackers.CommixAuto(userTarget)
                case "3":
                    webattackers.SqlmapMan()
                case "4":
                    webattackers.CommixMan()
                case "5":
                    webattackers.CommixAuto(userTarget)
                case "6":
                    webattackers.SqlmapMan()
                case "7":
                    webattackers.CommixMan()
                case "8":
                    webattackers.SqlmapMan()
                case "9":
                    webattackers.CommixMan()
                default:
                    fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 4 " + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //6. Jok3r v3.5 Insane Mult Reconing Engine Main Menu...👊
        case "6":
            screenClear(); banners.Banner();    menus.MenuEightSix()
            for {
                fmt.Printf(bcolors.GREEN + `¯\_(ツ)_/¯ Attacking!! ` + bcolors.RED + "︻╦╤─ ─ 🎯 " + bcolors.YELLOW + "%s\n" , userTarget + bcolors.ENDC)
                fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.SqlmapAuto(userTarget)
                case "2":
                    webattackers.CommixAuto(userTarget)
                case "3":
                    webattackers.SqlmapMan()
                case "4":
                    webattackers.CommixMan()
                case "5":
                    webattackers.CommixAuto(userTarget)
                case "6":
                    webattackers.SqlmapMan()
                case "7":
                    webattackers.CommixMan()
                case "8":
                    webattackers.SqlmapMan()
                case "9":
                    webattackers.CommixMan()
                default:
                    fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 4 " + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //7. Osmedeus Next Generation Workflow Engine Main Menu.🍈
        case "7":
            screenClear(); banners.Banner();  menus.MenuEightSeven()
            for {
                fmt.Printf(bcolors.GREEN + `¯\_(ツ)_/¯ Attacking!! ` + bcolors.RED + "︻╦╤─ ─ 🎯 " + bcolors.YELLOW + "%s\n" , userTarget + bcolors.ENDC)
                fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.SqlmapAuto(userTarget)
                case "2":
                    webattackers.CommixAuto(userTarget)
                case "3":
                    webattackers.SqlmapMan()
                case "4":
                    webattackers.CommixMan()
                case "5":
                    webattackers.CommixAuto(userTarget)
                case "6":
                    webattackers.SqlmapMan()
                case "7":
                    webattackers.CommixMan()
                case "8":
                    webattackers.SqlmapMan()
                case "9":
                    webattackers.CommixMan()
                default:
                    fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 4 " + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //8. Ufonet Next Generation DDOS Tool Main Menu.........🦠
        case "8":
            screenClear(); banners.Banner();  menus.MenuEightEight()
            for {
                fmt.Printf(bcolors.GREEN + `¯\_(ツ)_/¯ Attacking!! ` + bcolors.RED + "︻╦╤─ ─ 🎯 " + bcolors.YELLOW + "%s\n" , userTarget + bcolors.ENDC)
                fmt.Print(bcolors.GREEN + "\n(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
                fmt.Scan(&userInput)
                switch userInput {
                case "0":
                    websitesAttackers()
                    return
                case "1":
                    webattackers.SqlmapAuto(userTarget)
                case "2":
                    webattackers.CommixAuto(userTarget)
                case "3":
                    webattackers.SqlmapMan()
                case "4":
                    webattackers.CommixMan()
                case "5":
                    webattackers.CommixAuto(userTarget)
                case "6":
                    webattackers.SqlmapMan()
                case "7":
                    webattackers.CommixMan()
                case "8":
                    webattackers.SqlmapMan()
                case "9":
                    webattackers.CommixMan()
                default:
                    fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 4 " + bcolors.BLUE + ")" + bcolors.ENDC)
                }
            }
        //9. Launch Heavy Automation Attacks On The Host........🍄
        case "9":
            screenClear(); banners.Banner()// Heavy Automation Web Attack
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 9 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

//9. Help, Credits, Tricks and About.......(🕊  ︻╦╤─JC❤sU)
func creditsGivers() {
    screenClear(); menus.MenuNine()
}

//0. Exit africana-framework...............(Try option 99)
func scriptureNarators() {
    screenClear(); menus.MenuNineNine()
}

//Africana-Framework ......................(The rolling 9)
func africanaFramework() {
    screenClear(); banners.Banner(); menus.MenuZero()
    for {
        fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "0":
            screenClear(); banners.Banner()
            return
        case "1":
            africanaSetups()
            return
        case "2":
            anonymitySetups()
            return
        case "3":
            internalUserTarget(); internalAttackers()
            return
        case "4":
            malwareGenerators()
            return
        case "5":
            wirelesAttackers()
            return
        case "6":
            passwordsCrackers()
            return
        case "7":
            credsPhishers()
            return
        case "8":
            websiteUserTarget(); websitesAttackers()
            return
        case "9":
            creditsGivers()
            return
        case "99":
            scriptureNarators()
            return
        default:
            fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 9 " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

// ¯\_(ツ)_/¯..............................(Main runner..)
func main(){
    userAgreements()
    return
}
