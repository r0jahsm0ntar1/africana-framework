//John 3:16

package crackers

import(
    "os"
    "fmt"
    "utils"
    "menus"
    "strconv"
    "banners"
    "strings"
    "bcolors"
    "scriptures"
    "subprocess"
)

var(
    Function string
)

var defaultValues = map[string]string{

    "rhost":    "",
    "rhosts":   "",
    "proxies":  "",
    "function": "",
    "output":   utils.OutPutDir,
    "password": utils.PassWord,
    "wordlist": utils.WordsList,
}

type stringMatcher struct {
    names  []string
    action func()
}

func CrackersPentest() {
    for {
        fmt.Printf("%s%s%safr%s%s crackers(%s%ssrc/pentest_%s.fn%s)%s > %s", bcolors.Endc, bcolors.Underline, bcolors.Bold, subprocess.Version, bcolors.Endc, bcolors.Bold, bcolors.BrightRed, Function, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
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

        {[]string{"info"}, menus.HelpInfoCrackers},
        {[]string{"m", "menu"}, menus.MenuSix},
        {[]string{"option", "options", "show option", "show options"}, func() {menus.CrackersOptions(utils.CrMode, utils.RHost, utils.WordsList, utils.UserName, utils.PassWord)}},
        {[]string{"func", "funcs", "functions", "show func", "list funcs", "show funcs", "show function", "list function", "list functions", "show functions", "module", "modules", "list module", "show module", "list modules", "show modules", "show all", "list all"}, menus.ListCrackersFunctions},

        {[]string{"1", "run 1", "use 1", "exec 1", "start 1", "launch 1", "exploit 1", "execute 1", "run ssh", "use ssh", "exec ssh", "start ssh", "launch ssh", "exploit ssh", "execute ssh"}, func() {CrackersPenFunctions("ssh")}},
        {[]string{"? 1", "info 1", "help 1", "ssh", "info ssh", "help ssh"}, menus.HelpInfoFeatures},

        {[]string{"2", "run 2", "use 2", "exec 2", "start 2", "launch 2", "exploit 2", "execute 2", "run ftp", "use ftp", "exec ftp", "start ftp", "launch ftp", "exploit ftp", "execute ftp"}, func() {CrackersPenFunctions("ftp")}},
        {[]string{"? 2", "info 2", "help 2", "ftp", "info ftp", "help ftp"}, menus.HelpInfoFeatures},

        {[]string{"3", "run 3", "use 3", "exec 3", "start 3", "launch 3", "exploit 3", "execute 3", "run smb", "use smb", "exec smb", "start smb", "launch smb", "exploit smb", "execute smb"}, func() {CrackersPenFunctions("smb")}},
        {[]string{"? 3", "info 3", "help 3", "smb", "info smb", "help smb"}, menus.HelpInfoFeatures},

        {[]string{"4", "run 4", "use 4", "exec 4", "start 4", "launch 4", "exploit 4", "execute 4", "run rdp", "use rdp", "exec rdp", "start rdp", "launch rdp", "exploit rdp", "execute rdp"}, func() {CrackersPenFunctions("rdp")}},
        {[]string{"? 4", "info 4", "help 4", "rdp", "info rdp", "help rdp"}, menus.HelpInfoFeatures},

        {[]string{"5", "run 5", "use 5", "exec 5", "start 5", "launch 5", "exploit 5", "execute 5", "run ldap", "use ldap", "exec ldap", "start ldap", "launch ldap", "exploit ldap", "execute ldap"}, func() {CrackersPenFunctions("ldap")}},
        {[]string{"? 5", "info 5", "help 5", "ldap", "info ldap", "help ldap"}, menus.HelpInfoFeatures},

        {[]string{"6", "run 6", "use 6", "exec 6", "start 6", "launch 6", "exploit 6", "execute 6", "run smtp", "use smtp", "exec smtp", "start smtp", "launch smtp", "exploit smtp", "execute smtp"}, func() {CrackersPenFunctions("smtp")}},
        {[]string{"? 6", "info 6", "help 6", "smtp", "info smtp", "help smtp"}, menus.HelpInfoFeatures},

        {[]string{"7", "run 7", "use 7", "exec 7", "start 7", "launch 7", "exploit 7", "execute 7", "run telnet", "use telnet", "exec telnet", "start telnet", "launch telnet", "exploit telnet", "execute telnet"}, func() {CrackersPenFunctions("telnet")}},
        {[]string{"? 7", "info 7", "help 7", "telnet", "info telnet", "help telnet"}, menus.HelpInfoFeatures},

        {[]string{"8", "run 8", "use 8", "exec 8", "start 8", "launch 8", "exploit 8", "execute 8", "run https", "use https", "exec https", "start https", "launch https", "exploit https", "execute https"}, func() {CrackersPenFunctions("https")}},
        {[]string{"? 8", "info 8", "help 8", "https", "info https", "help https"}, menus.HelpInfoFeatures},

        {[]string{"9", "run 9", "use 9", "exec 9", "start 9", "launch 9", "exploit 9", "execute 9", "run pcap", "use pcap", "exec pcap", "start pcap", "launch pcap", "exploit pcap", "execute pcap"}, func() {CrackersPenFunctions("pcap")}},
        {[]string{"? 9", "info 9", "help 9", "pcap", "info pcap", "help pcap"}, menus.HelpInfoFeatures},

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
      "iface": &utils.IFace,
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
      "interface": &utils.IFace,
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

      "output": &utils.CrackersLogs,
      "outputlog": &utils.CrackersLogs,
      "outputlogs": &utils.CrackersLogs,

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
        maxWidth += 1

        cols := 5
        if len(suggestions) < cols {
            cols = len(suggestions)
        }

        for i := 0; i < len(suggestions); i += cols {
            for j := 0; j < cols && i+j < len(suggestions); j++ {
                fmt.Printf(" -> %s%-*s%s", bcolors.Green, maxWidth, suggestions[i+j], bcolors.Endc)
            }
            fmt.Println()
        }
        fmt.Println()
        return
    }

    fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Available keys:%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)
    maxWidth := 0
    for _, k := range validKeys {
        if len(k) > maxWidth {
            maxWidth = len(k)
        }
    }
    maxWidth += 1

    cols := 5
    for i := 0; i < len(validKeys); i += cols {
        for j := 0; j < cols && i+j < len(validKeys); j++ {
            fmt.Printf(" -> %s%-*s%s", bcolors.Green, maxWidth, validKeys[i+j], bcolors.Endc)
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

      "output": &utils.CrackersLogs,
      "outputlog": &utils.CrackersLogs,
      "outputlogs": &utils.CrackersLogs,
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
        maxWidth += 1

        cols := 5
        if len(suggestions) < cols {
            cols = len(suggestions)
        }
        
        for i := 0; i < len(suggestions); i += cols {
            for j := 0; j < cols && i+j < len(suggestions); j++ {
                fmt.Printf(" -> %s%-*s%s", bcolors.Green, maxWidth, suggestions[i+j], bcolors.Endc)
            }
            fmt.Println()
        }
        fmt.Println()
        return
    }

    fmt.Printf("%s[!] %sKey '%s%s%s' is invalid. Available keys:%s\n\n", bcolors.Yellow, bcolors.Endc, bcolors.Bold, key, bcolors.Endc, bcolors.Endc)

    maxWidth := 0
    for _, k := range validKeys {
        if len(k) > maxWidth {
            maxWidth = len(k)
        }
    }
    maxWidth += 1

    cols := 5
    for i := 0; i < len(validKeys); i += cols {
        for j := 0; j < cols && i+j < len(validKeys); j++ {
            fmt.Printf(" -> %s%-*s%s", bcolors.Green, maxWidth, validKeys[i+j], bcolors.Endc)
        }
        fmt.Println()
    }
}

func executeFunction() {
    if Function == ""{
        fmt.Printf("\n%s[!] %sNo MODULE was set. Use %s'show modules' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    CrackersPenFunctions(Function, utils.RHost)
}

func CrackersPenFunctions(Function string, args ...interface{}) {

    if utils.Proxies != "" {
        menus.PrintSelected(menus.PrintOptions{
            MODE: utils.CrMode,
            IFACE: utils.IFace,
            //LPORT: utils.LPort,
            //HPORT: utils.HPort,
            RHOST: utils.RHost,
            //LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            OUTPUTLOGS: utils.CrackersLogs,
            PROXIES: utils.Proxies,
            FUNCTION: Function,
            //RECONDIR: utils.ReconDir,
            //LISTENER: utils.Listener,
            PROTOCOL: utils.Protocol,
            //TOOLSDIR: utils.ToolsDir,
            //INNERICON: utils.InnerIcon,
            //OUTERICON: utils.OuterIcon,
            //BUILDNAME: utils.BuildName,
            //OBFUSCATOR: utils.Obfuscator,
        }, true, false)

        if err := utils.SetProxy(utils.Proxies); err != nil {
            //
        }
    } else {
        menus.PrintSelected(menus.PrintOptions{
            MODE: utils.CrMode,
            IFACE: utils.IFace,
            //LPORT: utils.LPort,
            //HPORT: utils.HPort,
            RHOST: utils.RHost,
            //LHOST: utils.LHost,
            //DISTRO: utils.Distro,
            //SCRIPT: utils.Script,
            OUTPUTLOGS: utils.CrackersLogs,
            FUNCTION: Function,
            //RECONDIR: utils.ReconDir,
            //LISTENER: utils.Listener,
            PROTOCOL: utils.Protocol,
            //TOOLSDIR: utils.ToolsDir,
            //INNERICON: utils.InnerIcon,
            //OUTERICON: utils.OuterIcon,
            //BUILDNAME: utils.BuildName,
            //OBFUSCATOR: utils.Obfuscator,
        }, true, false)
    }

    commands := map[string]func(){
      "ssh":    func() {HydraSsh(utils.RHost)},
      "ftp":    func() {HydraFtp(utils.RHost)},
      "smb":    func() {HydraSmb(utils.RHost)},
      "rdp":    func() {HydraRdp(utils.RHost)},
      "ldap":   func() {HydraLdap(utils.RHost)},
      "telnet": func() {HydraTelnet(utils.RHost)},
      "http":   func() {HydraHttp(utils.RHost)},
      "https":  func() {HydraHttp(utils.RHost)},
      "pcap":   func() {AirCrackng(utils.Pcap)},
      "ntlm":   func() {HydraNtlm(utils.RHost)},
      "hash":   func() {HashBuster(utils.Hashes)},

        "1": func() {HydraSsh(utils.RHost)},
        "2": func() {HydraFtp(utils.RHost)},
        "3": func() {HydraSmb(utils.RHost)},
        "4": func() {HydraRdp(utils.RHost)},
        "5": func() {HydraLdap(utils.RHost)},
        "6": func() {HydraTelnet(utils.RHost)},
        "7": func() {HydraHttp(utils.RHost)},
        "8": func() {HydraHttp(utils.RHost)},
        "9": func() {AirCrackng(utils.Pcap)},
       "10": func() {HydraNtlm(utils.RHost)},
       "11": func() {HashBuster(utils.Hashes)},
    }

    textCommands := []string{"ssh", "ftp", "smb", "rdp", "ldap", "smtp", "telnet", "http", "https", "pcap", "ntlm", "hash"}

    if action, exists := commands[Function]; exists {
        action()
        return
    }

    if num, err := strconv.Atoi(Function); err == nil {
        fmt.Printf("%s[!] %sNumber %s%d%s is invalid. Valid numbers are 1-10.\n", bcolors.Yellow, bcolors.Endc, bcolors.Red, num, bcolors.Endc)
        menus.ListCrackersFunctions()
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
    menus.ListCrackersFunctions()
}

func AirCrackng(Pcap string) {
    subprocess.Run("aircrack-ng %s -w %s", utils.Pcap, utils.WordsList)
    fmt.Println()
}

func JohnCrackng(Pcap string) {
    subprocess.Run("john %s --wordlist=%s", utils.Pcap, utils.WordsList)
    fmt.Println()
}

func CyberBrute(RHost string) {
    subprocess.Run("cd %s/cyberbrute; bash cyberbrute.sh %s", utils.CrackersTools, utils.RHost)
    fmt.Println()
}

func HashBuster(Hashes string) {
    subprocess.Run("cd %s/hash-buster; %s cracker.py -t 10 %s", utils.CrackersTools, utils.VenvPython, utils.Hashes)
    fmt.Println()
}

func HydraSsh(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'show options' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sSSHcracker running againist: %s\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("hydra -t 4 -l %s -P %s -f -o %s/HydraSsh_outfile.txt -u ssh://%s", utils.WordsList, utils.PassWord, utils.CrackersLogs, utils.RHost)
    return
}

func HydraFtp(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'show options' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sFTPcracker running againist: %s\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("hydra -t 4 -l %s -P %s -f -o %s/HydraFtp_outfile.txt -u ftp://%s", utils.WordsList, utils.PassWord, utils.CrackersLogs, utils.RHost)
    return
}

func HydraSmb(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'show options' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sSMBcracker running againist: %s\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("hydra -t 4 -L %s -P %s -f -o %s/HydraSmb_outfile.txt -u smb://%s", utils.WordsList, utils.PassWord, utils.CrackersLogs, utils.RHost)
    return
}

func HydraRdp(RHost string) {
    fmt.Printf("\n%s[*] %sRDPCracker running againist: %s\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("hydra -t 4 -L %s -P %s -f -o %s/HydraRdp_outfile.txt -u rdp://%s", utils.WordsList, utils.PassWord, utils.CrackersLogs, utils.RHost)
    return
}

func HydraLdap(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'show options' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sLDAPcracker running againist: %s\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("hydra -t 4 -L %s -P %s -f -o %s/HydraLdap_outfile.txt -u ldap://%s", utils.WordsList, utils.PassWord, utils.CrackersLogs, utils.RHost)
    return
}

func HydraSmtp(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'show options' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sSMTPCracker running againist: %s\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("hydra -t 4 -L %s -P %s -f -o %s/HydraSmtp_outfile.txt -u smtp://%s", utils.WordsList, utils.PassWord, utils.CrackersLogs, utils.RHost)
    return
}

func HydraSnmtp(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'show options' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sSNMTPCracker running againist: %s\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("hydra -t 4 -L %s -P %s -f -o %s/HydraSnmtp_outfile.txt -u snmtp://%s", utils.WordsList, utils.PassWord, utils.CrackersLogs, utils.RHost)
    return
}

func HydraTelnet(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'show options' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sTELNETCracker running againist: %s\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("hydra -t 4 -L %s -P %s -f -o %s/HydraTelnet_outfile.txt -u telnet://%s", utils.WordsList, utils.PassWord, utils.CrackersLogs, utils.RHost)
    return
}

func HydraHttp(RHost string) {
    if utils.RHost == "" {
        fmt.Printf("\n%s[!] %sMissing required parameters RHOST. Use %s'show options' %sfor details.\n", bcolors.BrightRed, bcolors.Endc, bcolors.BrightGreen, bcolors.Endc)
        return
    }
    fmt.Printf("\n%s[*] %sFTPcracker running againist: %s\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("hydra -t 4 -l %s -P %s -f -o %s/HydraHttps_outfile.txt -u https://%s", utils.WordsList, utils.PassWord, utils.CrackersLogs, utils.RHost)
    return
}


func HydraNtlm(RHost string) {
    fmt.Printf("\n%s[*] %sNTLMcracker running againist: %s\n", bcolors.Green, bcolors.Endc, utils.RHost)
    subprocess.Run("hydra -t 4 -l %s -P %s -f -o %s/HydraNtlm_outfile.txt -u https://%s", utils.WordsList, utils.PassWord, utils.CrackersLogs, utils.RHost)
    return
}
