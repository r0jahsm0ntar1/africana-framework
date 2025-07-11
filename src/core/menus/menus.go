//John 3:16

package menus

import (
    "fmt"
    "utils"
    "bcolors"
    "strings"
)

var (
    Interfaces, _ = utils.IFaces()
    LHost, _      = utils.GetDefaultIP()
    Gateway, _    = utils.GetDefaultGatewayIP()
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordListDir = utils.DirLocations()
)

var DefaultTableConfig = TableConfig{
    NameWidth:    15,
    SettingWidth: 18,
    ReqWidth:     9,
}

type TableConfig struct {
    NameWidth     int
    SettingWidth  int
    ReqWidth      int
}

type PrintOptions struct {
    ICON       string
    LPORT      string
    IFACE      string
    HPORT      string
    LHOST      string
    BUILD      string
    OUTPUT     string
    TOOLS_DIR  string
    PROTOCOL   string
    SCRIPT     string
    COREICON   string
    LISTENER   string
    BUILDNAME  string
    OBFUSCATOR string
}


type ModuleHelpInfo struct {
    Name          string
    Function      string
    Platform      string
    Arch          string
    Privileged    string
    License       string
    Rank          string
    Disclosed     string
    ProvidedBy    string
    CreatedBy     string
    TestedDistros string
    CheckSupport  string
    Description   string
    Example       string
    Options       string
}

func modulesHelp(info ModuleHelpInfo) {
    fmt.Printf(`
       %sName%s: %s
   %sFunction%s: %s
   %sPlatform%s: %s
       %sArch%s: %s
 %sPrivileged%s: %s
    %sLicense%s: %s
       %sRank%s: %s
  %sDisclosed%s: %s

%sProvided by%s: %s
 %sCreated by%s: %s

%sTested Distros%s:
   Id  Name
   --  ----
-> 0   %s

%sCheck supported%s:
  %s
`, bcolors.Bold, bcolors.Endc, info.Name, bcolors.Bold, bcolors.Endc, info.Function, bcolors.Bold, bcolors.Endc, info.Platform, bcolors.Bold, bcolors.Endc, info.Arch, bcolors.Bold, bcolors.Endc, info.Privileged, bcolors.Bold, bcolors.Endc, info.License, bcolors.Bold, bcolors.Endc, info.Rank, bcolors.Bold, bcolors.Endc, info.Disclosed, bcolors.Bold, bcolors.Endc, info.ProvidedBy, bcolors.Bold, bcolors.Endc, info.CreatedBy, bcolors.Bold, bcolors.Endc, info.TestedDistros, bcolors.Bold, bcolors.Endc, info.CheckSupport)
}

func generateOptions(modulePath string, options []string) {
    fmt.Printf(`
%sModule options %s(%s):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
`, bcolors.Bold, bcolors.Endc, modulePath, bcolors.Bold, bcolors.Endc)

    for _, opt := range options {
        fmt.Printf(opt)
    }
}

func modulesUsage(info ModuleHelpInfo) {
    if info.Options != "" {
        fmt.Printf("\n%sOptions%s:\n%s\n", bcolors.Bold, bcolors.Endc, info.Options)
    }

    if info.Description != "" {
        fmt.Printf(`

%sDescription%s:
  %s
`, bcolors.Bold, bcolors.Endc, info.Description)
    }

    if info.Example != "" {
        fmt.Printf(`
%sex. %s%susage%s:
--  -----
%s
`, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, info.Example)
    }
}

func FormatRow(nameWidth, settingWidth, reqWidth int, name, value, required, desc string) string {
    return fmt.Sprintf("  %-*s  %-*s  %-*s  %s", nameWidth, name, settingWidth, value, reqWidth, required, desc)
}

func calculateMaxWidths(rows [][]string) (int, int, int) {
    maxName := len("Name")
    maxSetting := len("Current Setting")
    maxReq := len("Required")

    for _, row := range rows {
        if len(row) >= 1 && len(row[0]) > maxName {
            maxName = len(row[0])
        }
        if len(row) >= 2 && len(row[1]) > maxSetting {
            maxSetting = len(row[1])
        }
        if len(row) >= 3 && len(row[2]) > maxReq {
            maxReq = len(row[2])
        }
    }

    if maxName < len("Name") {
        maxName = len("Name")
    }
    if maxSetting < len("Current Setting") {
        maxSetting = len("Current Setting")
    }
    if maxReq < len("Required") {
        maxReq = len("Required")
    }

    return maxName, maxSetting, maxReq
}

func FormatModuleHeader(modulePath string) string {
    return fmt.Sprintf("\nModule Options (%s%s%s):\n", bcolors.Italic, modulePath, bcolors.Endc)
}

func FormatColumnHeaders(nameWidth, settingWidth, reqWidth int) string {
    return FormatRow(nameWidth, settingWidth, reqWidth, "Name", "Current Setting", "Required", "Description") + "\n" + FormatRow(nameWidth, settingWidth, reqWidth, "----", "---------------", "--------", "-----------")
}

func FormatFooter(info ModuleHelpInfo) string {
    if info.Example != "" {
        return fmt.Sprintf(`
%sex. %s%susage%s:
--  -----
%s
View the full module info with the %s'info'%s or %s'show modules'%s, to get list of modules.
`, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, info.Example, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
    }
    return ""
}

func FormatModuleOptions(modulePath string, rows [][]string, moduleInfo ModuleHelpInfo) string {
    var builder strings.Builder

    maxName, maxSetting, maxReq := calculateMaxWidths(rows)

    builder.WriteString(FormatModuleHeader(modulePath))
    builder.WriteString("\n")
    builder.WriteString(FormatColumnHeaders(maxName, maxSetting, maxReq))
    builder.WriteString("\n")

    for _, row := range rows {
        if len(row) >= 4 {
            builder.WriteString(FormatRow(maxName, maxSetting, maxReq, row[0], row[1], row[2], row[3]))
            builder.WriteString("\n")
        }
    }

    builder.WriteString(FormatFooter(moduleInfo))

    return builder.String()
}

func FormatModuleOptionsWithConfig(modulePath string, config TableConfig, rows [][]string, moduleInfo ModuleHelpInfo) string {
    maxName, maxSetting, maxReq := calculateMaxWidths(rows)

    if config.NameWidth > maxName {
        maxName = config.NameWidth
    }
    if config.SettingWidth > maxSetting {
        maxSetting = config.SettingWidth
    }
    if config.ReqWidth > maxReq {
        maxReq = config.ReqWidth
    }

    var builder strings.Builder
    builder.WriteString(FormatModuleHeader(modulePath))
    builder.WriteString("\n")
    builder.WriteString(FormatColumnHeaders(maxName, maxSetting, maxReq))
    builder.WriteString("\n")

    for _, row := range rows {
        if len(row) >= 4 {
            builder.WriteString(FormatRow(maxName, maxSetting, maxReq, row[0], row[1], row[2], row[3]))
            builder.WriteString("\n")
        }
    }

    builder.WriteString(FormatFooter(moduleInfo))
    return builder.String()
}

func generateMenu(menuItems []string, helpText string, backOption bool) {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

`, bcolors.Italic, bcolors.Underl, bcolors.Bold, bcolors.Endc)

    maxLength := 0
    for _, item := range menuItems {
        if len(item) > maxLength {
            maxLength = len(item)
        }
    }

    for i, item := range menuItems {
        num := fmt.Sprintf("%2d", i + 1)
        fmt.Printf("%s %s. %s%s\n", bcolors.BrightBlue, num, bcolors.Endc, item)
    }

    if backOption {
        fmt.Printf(`
%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.BrightBlue, bcolors.Endc, bcolors.Underl, bcolors.Bold, bcolors.Italic, bcolors.Endc,  bcolors.BrightBlue, bcolors.Endc, bcolors.Italic, bcolors.Bold, bcolors.Underl, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Italic, bcolors.Underl, bcolors.Endc)
    } else {
        fmt.Printf(`
%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s10. %s%s%s%sGet Bonus Packages.%s

`, bcolors.BrightBlue, bcolors.Endc, bcolors.Underl, bcolors.Bold, bcolors.Italic, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Italic, bcolors.Bold, bcolors.Underl, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Italic, bcolors.Bold, bcolors.Underl, bcolors.Endc)
    }
}

func PrintSelected(opts PrintOptions, startWithNewLine bool, endWithNewLine bool) {
    printedAny := false

    printIfSet := func(name, value string) {
        if value != "" {
            fmt.Printf("%s => %s\n", name, value)
            printedAny = true
        }
    }

    if opts.ICON != "" || opts.IFACE != "" || opts.COREICON != "" || opts.LHOST != "" || opts.LPORT != "" || opts.HPORT != "" || opts.SCRIPT != "" || opts.PROTOCOL != "" || opts.LISTENER != "" || opts.TOOLS_DIR != "" || opts.OUTPUT != "" || opts.BUILDNAME != ""  || opts.OBFUSCATOR != "" {
        fmt.Println()
    }

    printIfSet("ICON", opts.ICON)
    printIfSet("IFACE", opts.IFACE)
    printIfSet("LPORT", opts.LPORT)
    printIfSet("HPORT", opts.HPORT)
    printIfSet("LHOST", opts.LHOST)
    printIfSet("SCRIPT", opts.SCRIPT)
    printIfSet("OUTPUT", opts.OUTPUT)
    printIfSet("LISTENER", opts.LISTENER)
    printIfSet("PROTOCOL", opts.PROTOCOL)
    printIfSet("COREICON", opts.COREICON)
    printIfSet("TOOLS_DIR", opts.TOOLS_DIR)
    printIfSet("BUILD_NAME", opts.BUILDNAME)
    printIfSet("OBFUSCATOR", opts.OBFUSCATOR)

    if printedAny && endWithNewLine {
        fmt.Println()
    }
}

func MenuZero() {
    items := []string{
        "Install, Update and View Logs",
        "System Security Configuration",
        "Local Network Penetration Testing",
        "Backdoors and Trojans Generating Engines",
        "Wireless Networks Attacks Vector",
        "PassWords, Hashes and .Pcap Crackers",
        "Social-Engineering and Phishings Vectors",
        "Website Penetration Testing Vectors",
        "Credits, and about the author",
    }
    generateMenu(items, "", false)
}

func MenuOne() {
    items := []string{
        "Kali",
        "Ubuntu",
        "Arch",
        "Macos",
        "Android",
        "Windows",
        "Update",
        "Repair",
        "UnInstall",
    }
    generateMenu(items, "", true)
}

func MenuTwo() {
    items := []string{
        "Install all necesary anonymity tools and configure them",
        "Start anonimity through tor",
        "Check if all anonimity softwares are up and running",
        "Check for external tor IP and status",
        "Display all outgoing traffick through proxies",
        "Restart all tor services and reconnect afresh",
        "Change your current exit node",
        "Bring back all default Iptables and other configs files",
        "Disconect from tor network and Kill all anonimty services",
    }
    generateMenu(items, "", true)
}

func MenuThree() {
    items := []string{
        "Discover targets",
        "Discover open ports",
        "Start Vuln scan",
        "Active directory recon",
        "Exploit active directory",
        "Sniff packets",
        "Launch killer Responder",
        "Launch M.I.B",
        "Start XSS Injection ",
    }
    generateMenu(items, "", true)
}

func MenuFour() {
    items := []string{
        "androrat",
        "teardroid",
        "blackjack",
        "hoaxshell",
        "shellz",
        "ghost",
        "chameleon",
        "lithaldll",
        "regsniper",
    }
    generateMenu(items, "", true)
}

func MenuFive() {
    items := []string{
        "wifite",
        "fluxion",
        "bettercap",
        "airgeddon",
        "wifipumpkin",
        "wifipumpkin3",
        "Coming soon",
        "Coming soon",
        "Coming soon",
    }
    generateMenu(items, "", true)
}

func MenuSix() {
    items := []string{
          "ssh",
          "ftp",
          "smb",
          "rdp",
          "ldap",
          "smtp",
          "telnet",
          "http/https",
          "all services",
          "hash",
          "pcap",
          "ntlm",
    }
    generateMenu(items, "", true)
}

func MenuSeven() {
    items := []string{
        "gophish",
        "goodginx",
        "zphisher",
        "blackeye",
        "advphisher",
        "darkphish",
        "shellphish",
        "setoolkit",
        "thc -> or TheHackerChoice",
    }
    generateMenu(items, "", true)
}

func MenuEight() {
    items := []string{
        "Network Mapping",
        "Subdomain enumeration",
        "Dns and asn lookup",
        "Web technology detection",
        "Wayback and asset discovery",
        "Fuzzing and Scanning",
        "Secret and leak detection",
        "Vulnerability scanning (XSS, SQLi, CSRF, SSRF, IDOR)",
        "Automation (Bug bounty hunting techniques)",
    }
    generateMenu(items, "", true)
}

func ListMainFunctions() {
    fmt.Printf(`
  # %sModule      Description%s
  - ------      -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"  setups", "Manage framework installation, updates, dependencies and complete removal."},
        {"torsocks", "Enforce Tor network routing for all traffic with leak protection."},
        {"networks", "Conduct MITM, DNS spoofing and network vulnerability assessments."},
        {"exploits", "Create stealth payloads and establish encrypted C2 connections. Free inplants."},
        {"wireless", "Perform WPA2 cracking, rogue AP attacks and spectrum analysis. Can & Bluetooth support."},
        {"crackers", "Bruteforce services and crack hashes using advanced techniques."},
        {"phishers", "Deploy convincing phishing sites with automated data capture."},
        {"websites", "Scan for vulnerabilities and exploit web application flaws. With bugbounty free recon."},
        {" credits", "View acknowledgments for contributors and integrated tools."},
        {"  verses", "Access scriptural references for spiritual growth. With  ethical hacking guidance."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "This module constains all sub-modules for africana-framework.",
        Example:          "    set MODULE setups\n    run\n",
    }
    modulesUsage(info)
}

func ListSetupsDistros() {
    fmt.Printf(`
  # %sDistro      Description%s
  - -----       -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"    kali", "Necessary tools will be installed in in Kali-Linux (Debian distros). Use this it is stable."},
        {"  ubuntu", "This module will install africana-framework in Ubuntu-Linux."},
        {"    arch", "Tools will be installed in any Arch-Linux Distros using Blackarch repo."},
        {"   macos", "Under development. Install africana on MackingTosh systems."},
        {" android", "Install africana-framework in Termux using chroot environment."},
        {" windows", "Under development. But can run if tools well installed using commando vm."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "This are the Distros by which can fully support africana-framework.",
        Example:          "    set distro kali\n    set module install\n    run\n",
    }
    modulesUsage(info)
}

func ListSetupsFunction() {
    fmt.Printf(`
  # %sModule      Description%s
  - ------      -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {" install", "Installs africana in selected distro."},
        {"  update", "Get new release of africana-framework from github and install it."},
        {"  repair", "Repair africana-framework if broken or with issues."},
        {"    auto", "Auto detect system and do the necessary."},
        {"  delete", "Completely uninstall africana-framework from system. alias to 'uninstall'."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the functions that can be interacted in setup module.",
        Example:          "    set MODULE install\n    run\n",
    }
    modulesUsage(info)
}

func ListTorsocksFunctions() {
    fmt.Printf(`
  # %sModule        Description%s
  - ------        -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"  setups", "Installs dnsmasq, squid, privoxy and tor (also set configs)."},
        {"  vanish", "Start anonymizing the system through tor network."},
        {"  status", "Check if using tor (Show if all anononimty services are up and running)."},
        {"   torip", "Check current tor IP address."},
        {"  chains", "View traffic logs from squid, privoxy, to tor."},
        {"  reload", "Completely restart torsocks and connect to a different exit-node."},
        {"exitnode", "Connect to a different exit-node."},
        {" restore", "Backup and resets Iptables to default."},
        {"    stop", "Get back to the surface-web."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the functions that torsocks module performs.",
        Example:          "    set MODULE discover\n    run\n",
    }
    modulesUsage(info)
}

func ListInternalFunctions() {
    fmt.Printf(`
  # %sModule      Description%s
  - ------      -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"discover", "Discover all devices connected to the network. Lets you locate targets."},
        {"portscan", "Get open ports on the target you have set."},
        {"vulnscan", "Perform vulnerbility scan on open ports of the target you have set."},
        {"enumscan", "Recon for S.M.B deep information on the target set."},
        {"smbexplo", "Launch known vulnerbility exploits on the target's S.M.B services."},
        {"psniffer", "Sniff all Packets from connected devices to the router(Perform M.I.T.M)."},
        {"responda", "Start Killer Responder that configs all required fields to get you a reverse shell on windows. Supports IPv6."},
        {"beefkill", "Launch Beef-xss and Bettercap/ Ettercp For effective (M.I.B attacks)."},
        {"toxssinx", "Get Shell through XSS Injection to packets in the wire. Works like beef-xss."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the functions that networks module performs.",
        Example:          "    set MODULE discover\n    run\n",
    }
    modulesUsage(info)
}

func ListExploitsFunctions() {
    fmt.Printf(`
  # %sModule       Description%s
  - ------       -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {" androrat", "It is another Android Listener. It is cool but only works on android 4 to 9. Suppoorts android < 14 but less functions."},
        {"teardroid", "Andoird Listener. Needs alot of online configuration but the best for now."},
        {"blackjack", "It is a tool derived from villain. It supports both tcp, http and https reverse shells. Supports evasions and bypasses almost all avs."},
        {"hoaxshell", "A Killer FUD that bypasses most avs."},
        {"   shellz", "Supports all distro reverse shells generation that supports both tcp, https and https connection. Launches variety of listeners."},
        {"    ghost", "It is a giant powershell evasion tool that beats almost all AVS. Try it out you will love it."},
        {"chameleon", "It is a python framework evasion tool that beats almost all AVS. Try it out you will love it."},
        {"regsniper", "Just like lithaldll but the persisten mechanisim is through regestry keys."},
        {"lithaldll", "Generates undetected backdoor with embeded hoax reverse shell. Has .dll persistent mechanisims targetting OneDrive."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the functions that exploits module performs.",
        Example:          "    set MODULE hoaxshell\n    run\n",
    }
    modulesUsage(info)
}

func ListListenersFunctions() {
    fmt.Printf(`
  # %sModule        Description%s
  - ------        -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"  blackjack", "It is a tool derived from villain. It supports both tcp, http and https reverse shells. Supports evasions and bypasses almost all avs."},
        {"  hoaxshell", "A Killer FUD that bypasses most avs."},
        {"       ncat", "Starts ncat to listen for either tcp or http(s) protocol for powershell reverse shell. Best launched with lithaldll or regsniper modules."},
        {" metasploit", "Launches metasploit to listen for either tcp or http(s) protocol for powershell reverse shell. Best launched with lithaldll or regsniper modules."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the listeners you can call to handle reverse connections.",
        Example:          "    set MODULE lithaldll\n    set listener blackjack\n    run\n",
    }
    modulesUsage(info)
}

func ListObfscatorsFunctions() {
    fmt.Printf(`
  # %sModule       Description%s
  - ------       -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"    ghost", "This tool introduces is a methodology where you can target and obfuscate the individual components of a script with randomized variations."},
        {"    psobf", "It is a giant powershell evasion tool that beats almost all AVS. Try it out you will love it."},
        {"   vulkan", "Vulkan is able to obfuscate powershell scripts in order to make them undetectable against antivirus."},
        {"chameleon", "It is a python framework evasion tool that beats almost all AVS. Try it out you will love it."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the functions that obfscator module performs.",
        Example:          "    set script Full_path_to_your .ps1\n    run\n",
    }
    modulesUsage(info)
}

func ListIcons() {
    fmt.Printf(`
  # %sIcon         Description%s
  - ----         -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"   access","Microsoft access document icon for your output malware."},
        {"  autorun","Autorun standard icon to disguise your malicious malware."},
        {"    excel","Microsoft excel document icon for your output malware."},
        {"      pdf","Standard .pdf icon for your malicious backdoor."},
        {"  project","Another icon found in microsoft office tools."},
        {"publisher","Microsoft publisher document icon for your output malware."},
        {"   redrat","A simple icon for a rat. This will fortunetly look suspicious."},
        {"    visio","Microsoft access document icon for your output malware."},
        {"      vlc","An icon that will output your malware as a vlc vidoe player."},
        {"     word","Microsoft word document icon for your output malware."},
        {" defender","Microsoft defender icon for your output malware."},
        {"kaspersky","Kapersky antivirus icon for your output malware."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the icons that can be used to embede backdoors.",
        Example:          "    set icon access\n",
    }
    modulesUsage(info)
}

func ListWirelessFunctions() {
    fmt.Printf(`
  # %sModule       Description%s
  - ------       -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"   wifite","Wifite is a tool to audit WEP or WPA encrypted wireless networks."},
        {"  fluxion","Fluxion is a tool to audit WEP or WPA encrypted wireless networks. Only manual option is supported by now."},
        {"bettercap","Bettercap is a tool to audit Internal network and wirekless network like, WEP or WPA encrypted wireless networks."},
        {"airgeddon","Airgeddon Fluxion is a tool to audit WEP or WPA encrypted wireless networks. Only manual option is supported by now."},
        {"wifipumpk","wifipumpkin, Is a Powerful framework for rogue access point attack. This option run automated mode directly."},
        {"autopumpk","This option runs wifipumpkin3 in manual mode where africana sets everything for you."},
        {"  To come",""},
        {"  To come",""},
        {"  To come",""},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the functions that wireless module performs.",
        Example:          "    set MODULE wifite or -> use wifite\n    run\n",
    }
    modulesUsage(info)
}

func ListWebsitesFunctions() {
    fmt.Printf(`
  # %sModule      Description%s
  - ------      -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"  netmap","Network Mapping and portscaning."},
        {"enumscan","Subdomain enumeration."},
        {"dnsrecon","Dns and asn lookup."},
        {"techscan","Web technology detection."},
        {"asetscan","Wayback and asset discovery."},
        {"fuzzscan","Digg for root files from the server."},
        {"leakscan","Secret and leak detection."},
        {"vulnscan","Vulnerability scanning (XSS, SQLi, CSRF, SSRF, IDOR)."},
        {"  bounty","Bug bounty hunting techniques."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the functions that websites module performs.",
        Example:          "    set MODULE enumscan or -> use enumscan\n    set rhosts https://example.com\n    run\n",
    }
    modulesUsage(info)
}

func ListCrackersFunctions() {
    fmt.Printf(`
  # %sModule     Description%s
  - ------     -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"    ssh","Automated Bruteforcer for SSH.(mode online)."},
        {"    ftp","Bruteforcer for FTP.(mode online)."},
        {"    smb","Hydra Bruteforcer for SMB.(mode online)."},
        {"    rdp","Bruteforcer for RDP.(mode online)."},
        {"   ldap","Hydra Bruteforcer for LDAP.(mode online)."},
        {"   smtp","Bruteforcer for SMTP.(mode online)."},
        {" telnet","Bruteforcer for TELNET.(mode online)."},
        {"http(s)","Hydra Bruteforcer for HTTP/ HTTPS. Forms needed.(mode online)."},
        {"   pcap","Crack captured .pcap files. Full location is needed for .pcap.(mod offline)."},
        {"   ntlm","Crack ntlm file.Full location is needed for .ntlm.(mod offline)."},
        {"   hash","Auto identify hash and start bruteforcing.(mod offline)."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-7s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the functions that crackers module performs.",
        Example:          "    set mode online\n    set module ssh\n    set wordlist -> (Give full path)\n    run\n",
    }
    modulesUsage(info)
}

func ListPhishersFunctions() {
    fmt.Printf(`
  # %sModule        Description%s
  - ------        -----------`, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"   gophish", "It is a phishing framework with a Web UI https://127.0.0.1:3333. Africana will launch it for you. Default  is: admin, Default password is: kali-gophish."},
        {"  goodginx", "Goodginx is an advanced phishing framework with insane configurations.Default name evilginx2. Bypasses alot of security features like OTP."},
        {"  zphisher", "A nice framework with alot of templets. Also bypasses OTP with ngrock support."},
        {"  blackeye", "Writen in bash and full of phishing templets. Just check it out."},
        {"advphisher", "Wide range of phishing templets."},
        {" darkphish", "Bypasses OTP with Wide range of phishing templets."},
        {"shellphish", "Start Killer Responder that configs all required fields to get you a reverse shell on windows. Supports IPv6."},
        {" setoolkit", "This tool is equiped with alot of social engeneering. Supports cloning of actual websites."},
        {"       thc", "This tool creates a templete of your interest imidietly but needs you to start your server and generate a link for phishing."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Description:      "These are the functions that phishers module performs.",
        Example:          "    set MODULE gophish\n    run\n",
    }
    modulesUsage(info)
}

//Help Info functions
func HelpInfoMain() {
    info := ModuleHelpInfo{
        Name:          "main",
        Function:      "src/core/africana",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
    }
    modulesHelp(info)
    ListMainFunctions()
}

func HelpInfoSetups() {
    info := ModuleHelpInfo{
        Name:          "setups",
        Function:      "src/core/setups",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This modules enables you to Install, uninstall, update and mentain africana-framework.",
    }
    modulesHelp(info)
    ListSetupsFunction()
}

func HelpInfoTorsocks() {
    info := ModuleHelpInfo{
        Name:          "torsocks",
        Function:      "src/securities/torsocks.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Torsocks is a tool that strictly configures Iptables, Tor, Dsnsmasq, Privoxy and Squid to work together in order to completely anonimize your system through Tor network.",
    }
    modulesHelp(info)
    ListTorsocksFunctions()
}

func HelpInfoKali() {
    info := ModuleHelpInfo{
        Name:          "kali",
        Function:      "src/core/setups_kali.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a module to install africana-framework in kali-linux a stable debian based distro that has a wide comunity support to avoid package breaks and missing dependencies use kali for africana.",
    }
    modulesHelp(info)
}

func HelpInfoArch() {
    info := ModuleHelpInfo{
        Name:          "arch",
        Function:      "src/core/setups_arch.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a module to install africana-framework in arch based distros. Arch is well established and all tools could be installed with blackman an intergration of black-arch in any arch-linux distro. No errors reported. africana can run well in arch-linux distros.",
    }
    modulesHelp(info)
}

func HelpInfoUbuntu() {
    info := ModuleHelpInfo{
        Name:          "ubuntu",
        Function:      "src/core/setups_ubuntu.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a module to install africana-framework ubuntu which is a good distor but has alot of problems while installing kali-linux packages. To avoid issues like dependencies problems, Pleas use docker image or install kali-linux in Ubuntu docker then install africana.",
    }
    modulesHelp(info)
}

func HelpInfoMacos() {
    info := ModuleHelpInfo{
        Name:          "macos",
        Function:      "src/core/setups_macos.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a module to install africana-framework Macos which is a good distor but has alot of problems while installing kali-linux packages. To avoid issues like dependencies problems, Pleas use docker image or install kali-linux in Macos docker then install africana.",
    }
    modulesHelp(info)
}


func HelpInfoWindows() {
    info := ModuleHelpInfo{
        Name:          "windows",
        Function:      "src/core/setups_windows.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a module to install africana-framework Windows which is a good distor but has alot of problems while installing kali-linux packages. To avoid issues like dependencies problems, Pleas use docker image or install kali-linux in Macos docker then install africana.",
    }
    modulesHelp(info)
}

func HelpInfoAndroid() {
    info := ModuleHelpInfo{
        Name:          "android",
        Function:      "src/core/setups_android.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function to install africana-framework in android devices. Kali-linux will be installed in termux then kali-linux in chroot environment that will set all dependencies for africana-framework.",
    }
    modulesHelp(info)
}


func HelpInfoUpdate() {
    info := ModuleHelpInfo{
        Name:          "update",
        Function:      "src/core/setups_update.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function to update and mentain africana-framework.",
    }
    modulesHelp(info)
}

func HelpInfoUninstall() {
    info := ModuleHelpInfo{
        Name:          "uninstall",
        Function:      "src/core/setups_uninstall.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function to Uninstall africana completelty from your system with all it's dependencies. Incase of a bug, email me at rojahsmontari@gmail.com",
    }
    modulesHelp(info)
}

func HelpInfoAuto() {
    info := ModuleHelpInfo{
        Name:          "auto",
        Function:      "src/core/setups_auto.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function to auto select distro and install africana with all it's dependencies.",
    }
    modulesHelp(info)
}


func HelpInfoRepair() {
    info := ModuleHelpInfo{
        Name:          "repair",
        Function:      "src/core/setups_repair.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a function repairs africana incase it is broken.",
    }
    modulesHelp(info)
}


func HelpInfoClearLogs() {
    info := ModuleHelpInfo{
        Name:          "clearlogs",
        Function:      "src/core/setups_clearlogs.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will clear all your logs that has been recorded from the last time you cleaned the log folder.",
    }
    modulesHelp(info)
}

func HelpInfoTorsocksSetups() {
    info := ModuleHelpInfo{
        Name:          "setups",
        Function:      "src/securities/setups.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will install dnsmasq, squid, privoxy and tor. It will (also set configs) so that all your local traffick will go through privoxy -> squid -> then tor network. It is done with great care and integrity for super securities.",
    }
    modulesHelp(info)
}

func HelpInfoTorsocksVanish() {
    info := ModuleHelpInfo{
        Name:          "vanish",
        Function:      "src/securities/vanish.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will start services like changemacc to change maccadress in a random way then start dnsmasq, squid, privoxy and tor. It will (also set configs) so that all your local traffick will go through privoxy > squid > then tor network.",
    }
    modulesHelp(info)
}

func HelpInfoTorsocksStatus() {
    info := ModuleHelpInfo{
        Name:          "status",
        Function:      "src/securities/status.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will query the system to see if macchanger, dnsmasq, squid, privoxy and tor are working correctly and if all traffic that goes through privoxy > squid > then tor network. It is done with great care and integrity for super securities.",
    }
    modulesHelp(info)
}


func HelpInfoTorsocksTorIp() {
    info := ModuleHelpInfo{
        Name:          "torip",
        Function:      "src/securities/torip.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will check for your external IP. It querries tor website for your gateway IP. If your system's proxy is correctly configured, then you will get a congratulation mesage from tor website.",
    }
    modulesHelp(info)
}


func HelpInfoTorsocksChains() {
    info := ModuleHelpInfo{
        Name:          "chains",
        Function:      "src/securities/chains.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will querry  /var/log/privoxy/log to follow all logs living your system through squid, privoxy to tor.",
    }
    modulesHelp(info)
}


func HelpInfoTorsocksReload() {
    info := ModuleHelpInfo{
        Name:          "reload",
        Function:      "src/securities/reload.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will stop all tor services and restart again acuring new exitnodes and torchains.",
    }
    modulesHelp(info)
}

func HelpInfoTorsocksExitnode() {
    info := ModuleHelpInfo{
        Name:          "exitnode",
        Function:      "src/securities/exitnode.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will shufle the exit nodes to new ones. If you see your nrtwork is slow, This module can help to find a fast one.",
    }
    modulesHelp(info)
}


func HelpInfoTorsocksRestore() {
    info := ModuleHelpInfo{
        Name:          "restore",
        Function:      "src/securities/restore.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will restore your Iptables to default. If the Function was killed instantly and IPTABLES were not set as intended, This module will help you fix the lack off internet connection.",
    }
    modulesHelp(info)
}


func HelpInfoTorsocksStop() {
    info := ModuleHelpInfo{
        Name:          "stop",
        Function:      "src/securities/stop.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will restore your Iptables to default. If the Function was killed instantly and IPTABLES were not set as intended, This module will help you fix the lack off internet connection.",
    }
    modulesHelp(info)
}

func HelpInfoNetworks() {
    info := ModuleHelpInfo{
        Name:          "Networks",
        Function:      "src/networks.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This is the Network module that contains all internal networks attacks functions.",
    }
    modulesHelp(info)
    ListInternalFunctions()
}

//Exploits
func HelpInfoChameleon() {
    info := ModuleHelpInfo{
        Name:          "exploits",
        Function:      "src/exploits/chameleon.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will help you in obfuscating any powershell script in order to bypass any AV durring execution.",
    }
    modulesHelp(info)
}

func HelpInfoExploits() {
    info := ModuleHelpInfo{
        Name:          "exploits",
        Function:      "src/networks/exploits.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This is the Exploits module that contains all Listener, backdoors and obfsicators functions.",
        Example:          "    set MODULE exploits\n    run\n",
    }
    modulesHelp(info)
    ListExploitsFunctions()
}

//Crackers
func HelpInfoCrackers() {
    info := ModuleHelpInfo{
        Name:          "crackers",
        Function:      "src/crackers.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Crackers is a module enriched with creative attacking faces to help redtemers in successfully cracking or brutforce passwords from services online or local encripted files.",
    }
    modulesHelp(info)
    ListCrackersFunctions()
}


//phishers
func HelpInfoPhishers() {
    info := ModuleHelpInfo{
        Name:          "phishers",
        Function:      "src/phishers/phishers.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Phishers is a module enriched with creative attacking faces to help redtemers in successfully Perform phishing attacks with ease.",
    }
    modulesHelp(info)
    ListPhishersFunctions()
}

func HelpInfoGoPhish() {
    info := ModuleHelpInfo{
        Name:          "gophish",
        Function:      "src/phishers/gophish.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function that enables the redteamers to perform phising attacks on various bases.",
        Example:          "    set MODULE gophish\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoGoodGinx() { 
    info := ModuleHelpInfo{
        Name:          "goodginx",
        Function:      "src/phishers/goodginx.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function that enables the redteamers to perform phising attacks on various bases.",
        Example:          "    set MODULE goodginx\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoZPhisher() {
    info := ModuleHelpInfo{
        Name:          "zphisher",
        Function:      "src/phishers/zphisher.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function that enables the redteamers to perform phising attacks on various bases.",
        Example:          "    set MODULE zphisher\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoBlackEye() {
    info := ModuleHelpInfo{
        Name:          "blackeye",
        Function:      "src/phishers/blackeye.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function that enables the redteamers to perform phising attacks on various bases.",
        Example:          "    set MODULE blackeye\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoAdvPhisher() {
    info := ModuleHelpInfo{
        Name:          "advphisher",
        Function:      "src/phishers/advphisher.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function that enables the redteamers to perform phising attacks on various bases.",
        Example:          "    set MODULE advphisher\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoDarkPhish() { 
    info := ModuleHelpInfo{
        Name:          "darkphish",
        Function:      "src/phishers/darkphish.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function that enables the redteamers to perform phising attacks on various bases.",
        Example:          "    set MODULE darkphish\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoShellPhish() {
    info := ModuleHelpInfo{
        Name:          "shellphish",
        Function:      "src/phishers/shellphish.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function that enables the redteamers to perform phising attacks on various bases.",
        Example:          "    set MODULE shellphish\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoSetoolKit() {
    info := ModuleHelpInfo{
        Name:          "setoolkit",
        Function:      "src/phishers/setoolkit.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function that enables the redteamers to perform phising attacks on various bases.",
        Example:          "    set MODULE setoolkit\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoThc() {
    info := ModuleHelpInfo{
        Name:          "thc",
        Function:      "src/phishers/thc.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a Function that enables the redteamers to perform phising attacks on various bases.",
        Example:          "    set MODULE thc\n    run\n",
    }
    modulesHelp(info)
}

//websites
func HelpInfoWebsites() {
    info := ModuleHelpInfo{
        Name:          "websites",
        Function:      "src/websites/bug_bounty.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
    }
    modulesHelp(info)
    ListWebsitesFunctions()
}

func HelpInfoEnumScan() {
    info := ModuleHelpInfo{
        Name:          "enumscan",
        Function:      "src/websites/enumscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE enumscan\n    set rhosts https://example.com\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoDnsRecon() {
    info := ModuleHelpInfo{
        Name:          "dnsrecon",
        Function:      "src/websites/dnsrecon.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE dnsrecon\n    set rhosts https://example.com\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoPortScan() {
    info := ModuleHelpInfo{
        Name:          "portscan",
        Function:      "src/websites/portscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE dnsrecon\n    set rhosts https://example.com\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoTechScan() {
    info := ModuleHelpInfo{
        Name:          "techscan",
        Function:      "src/websites/techscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE techscan\n    set rhosts https://example.com\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoFuzzScan() {
    info := ModuleHelpInfo{
        Name:          "fuzzscan",
        Function:      "src/websites/fuzzscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE fuzzscan\n    set rhosts https://example.com\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoLeakScan() {
    info := ModuleHelpInfo{
        Name:          "leakscan",
        Function:      "src/websites/leakscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE leakscan\n    set rhosts https://example.com\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoVulnScan() {
    info := ModuleHelpInfo{
        Name:          "vulnscan",
        Function:      "src/websites/vulnscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE vulnscan\n    set rhosts https://example.com\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoAsetScan() {
    info := ModuleHelpInfo{
        Name:          "asetscan",
        Function:      "src/websites/asetscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE asetscan\n    set rhosts https://example.com\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoAutoScan() {
    info := ModuleHelpInfo{
        Name:          "autoscan",
        Function:      "src/websites/autoscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE autoscan\n    set rhosts https://example.com\n    run\n",
    }
    modulesHelp(info)
}

//credits
func HelpInfoCredits() {
    info := ModuleHelpInfo{
        Name:          "credits",
        Function:      "src/websites/credits.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE credits\n    run\n",
    }
    modulesHelp(info)
}


//verses
func HelpInfoVerses() {
    info := ModuleHelpInfo{
        Name:          "verses",
        Function:      "src/websites/verses.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE verses\n    run\n",
    }
    modulesHelp(info)
}


//networks
func HelpInfoDiscover() {
    info := ModuleHelpInfo{
        Name:          "discover",
        Function:      "src/networks/discover.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will scan for all connected devices in the network given using bettercap then arrange the targets in a table for you to select one to attack further.",
        Example:          "    set MODULE discover\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoInEnumScan() {
    info := ModuleHelpInfo{
        Name:          "enumscan",
        Function:      "src/websites/enumscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease. It consists off recons, vulners, ddos among others.",
        Example:          "    set MODULE enumscan\n    set rhosts https://example.com\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoInPortScan() {
    info := ModuleHelpInfo{
        Name:          "portscan",
        Function:      "src/networks/portscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will scan all open ports of the target to reveal open ports.",
        Example:          "    set MODULE portscan\n    run\n",
    }
    modulesHelp(info)
}


func HelpInfoInVulnScan() {
    info := ModuleHelpInfo{
        Name:          "vulnscan",
        Function:      "src/networks/vulnscan.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will Perform vulnerbility scan on open ports of the target you have set.",
        Example:          "    set MODULE vulnscan\n    run\n",
    }
    modulesHelp(info)
}


func HelpInfoSmbExplo() {
    info := ModuleHelpInfo{
        Name:          "smbexplo",
        Function:      "src/networks/smbexplo.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will Launch known vulnerbility exploits on the target's S.M.B services.",
        Example:          "    set MODULE smbexplo\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoPSniffer() {
    info := ModuleHelpInfo{
        Name:          "psniffer",
        Function:      "src/networks/psniffer.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will Sniff all Packets from connected devices to the router(Perform M.I.T.M).",
        Example:          "    set MODULE psniffer\n    run\n",
    }
    modulesHelp(info)
}

func HelpInfoResponder(Mode, LPort, RHost, LHost string) {
    info := ModuleHelpInfo{
        Name:          "responder",
        Function:      "src/networks/responder.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will Launch reponder asking for your LHOST, Configuring Wpadscript and weponizing it self. Attack supports alot of windows recent version.",
    }
    modulesHelp(info)
    ResponderOptions(Mode, LPort, RHost, LHost)
}

func HelpInfoBeefKill(Mode, LPort, Spoofer, RHost, LHost string) {
    info := ModuleHelpInfo{
        Name:          "beefkill",
        Function:      "src/networks/beefkill.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will Launch a Combination of both beef-xss and bettercap in a unique way to inject hook.js in either one or all targets. All settings are done for you.",
        Example:          "    set MODULE beefkill\n    run\n",
    }
    modulesHelp(info)
    BeefKillOptions(Mode, LPort, Spoofer, RHost, LHost)
}


func HelpInfoToxssInx(Mode, LPort, Spoofer, RHost, LHost string) {
    info := ModuleHelpInfo{
        Name:          "toxssinx",
        Function:      "src/networks/toxssinx.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "This module will try to Get you a revers Shell through XSS Injection. Still Working on this Option.",
    }
    modulesHelp(info)
    ToxssInxOptions(Mode, LPort, Spoofer, RHost, LHost)
}

func HelpInfoWireless() {
    info := ModuleHelpInfo{
        Name:          "wireless",
        Function:      "src/networks/wireless.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
    }
    modulesHelp(info)
    ListWirelessFunctions()
}

func HelpInfoBlackJack(LHost string) {
    info := ModuleHelpInfo{
        Name:          "blackjack",
        Function:      "src/exploits/blackajck.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        ProvidedBy:    "r0jahsm0ntar1",
        CreatedBy:     "t3l3machus",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a tool derived from villain framework. It supports both tcp, http and https reverse shells. It has inbuild evasions and bypasses almost all avs. It is the best for now.",
    }
    modulesHelp(info)
    BlackJackOptions(LHost)
}

func HelpInfoShellz(LHost, LPort, Protocol string) {
    info := ModuleHelpInfo{
        Name:          "shellz",
        Function:      "src/exploits/shellz.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        ProvidedBy:    "r0jahsm0ntar1",
        CreatedBy:     "t3l3machus",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.",
        Example:          "    set MODULE shellz\n    run\n",
    }
    modulesHelp(info)
    ShellzOptions(LHost, LPort, Protocol)
}

func HelpInfoHoaxShell(LHost, LPort, Protocol string) {
    info := ModuleHelpInfo{
        Name:          "hoaxshell",
        Function:      "src/exploits/hoaxshell.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        ProvidedBy:    "r0jahsm0ntar1",
        CreatedBy:     "t3l3machus",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a tool like villein or blackjack that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.",
    }
    modulesHelp(info)
    HoaxshellOptions(LHost, LPort, Protocol)
}


func HelpInfoLithalDll(Icon, LHost, LPort, HPort, Protocol, OutPut, Listener string) {
    info := ModuleHelpInfo{
        Name:          "lithaldll",
        Function:      "src/exploits/lithaldll.fn",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        ProvidedBy:    "r0jahsm0ntar1",
        CreatedBy:     "Israel Isreal",
        TestedDistros: "All Distros",
        CheckSupport:  "Yes",
        Description:   "It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.",
    }
    modulesHelp(info)
    LithalDllOptions(Icon, LHost, LPort, HPort, Protocol, OutPut, Listener)
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

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  ICON           pdf              yes       The icon to use to disguise your backdoor with.
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, LHost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Androids

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, LHost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  SCRIPT         none             yes       Full location of the powershell script to be obfsicated.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, LHost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  SCRIPT         none             yes       Powershell script to obfsicate inorder to bypass AVs.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, LHost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All IOS

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, LHost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, LHost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoRegSniper() {
    fmt.Printf(`
       %sName%s: regsniper
   %sFunction%s: src/exploits/lithaldll.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  ICON           pdf              yes       The icon to use to disguise your backdoor with.
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, LHost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Androids

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, LHost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Windows

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, LHost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}


func HelpInfoWifite() {
    fmt.Printf(`
       %sName%s: wifite
   %sFunction%s: src/exploits/wifite.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Wi-fi

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  MODE           auto             yes       Attacking mode to use. (auto or manual)
  IFACE          wlan0            yes       Mainly needed for monitoring and deuthing ect.
  WORDLISTS      rockyou.txt      yes       WordListDir location for cracking captured handshakes.

%sDescription%s:
-----------
  Wifite is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit. This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

%sex. %s%susage%s:
--  -----

  set function wifite
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoFluxion() {
    fmt.Printf(`
       %sName%s: fluxion
   %sFunction%s: src/exploits/fluxion.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Wi-fi

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  MODE           auto             yes       Attacking mode to use. (auto or manual)
  IFACE          wlan0            yes       Mainly needed for monitoring and deuthing ect.
  WORDLISTS      rockyou.txt      yes       WordListDir location for cracking captured handshakes.

%sDescription%s:
-----------
  Fluxion is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit. This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

%sex. %s%susage%s:
--  -----

  set function fluxion
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoBetterCap() {
    fmt.Printf(`
       %sName%s: bettercap
   %sFunction%s: src/exploits/bettercap.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Wi-fi

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  MODE           auto             yes       Attacking mode to use. (auto or manual)
  IFACE          wlan0            yes       Mainly needed for monitoring and deuthing ect.
  WORDLISTS      rockyou.txt      yes       WordListDir location for cracking captured handshakes.

%sDescription%s:
-----------
  Bettercap is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit. This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

%sex. %s%susage%s:
--  -----

  set function bettercap
  set mode auto
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoAirGeddon() {
    fmt.Printf(`
       %sName%s: airgeddon
   %sFunction%s: src/exploits/airgeddon.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Wi-fi

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  MODE           auto             yes       Attacking mode to use. (auto or manual)
  IFACE          wlan0            yes       Mainly needed for monitoring and deuthing ect.
  WORDLISTS      rockyou.txt      yes       WordListDir location for cracking captured handshakes.

%sDescription%s:
-----------
  Airgeddon is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit. This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

%sex. %s%susage%s:
--  -----

  set function airgeddon
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoWifiPumpkin() {
    fmt.Printf(`
       %sName%s: wifipumpkin
   %sFunction%s: src/exploits/wifipumpkin.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: No
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2025

%sProvided by%s:
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Wi-fi

%sBasic options%s:
----- --------
  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  MODE           auto             yes       Attacking mode to use. (auto or manual)
  SSID           ->               yes       The fake name of your wifi for the clients to see. Default = 'End times ministries'
  IFACE          wlan0            yes       Mainly needed for monitoring and deuthing ect.
  WORDLISTS      rockyou.txt      yes       WordListDir location for cracking captured handshakes.

%sDescription%s:
-----------
  Fluxion is a tool to audit WEP or WPA encrypted wireless networks. It uses aircrack-ng, pyrit, reaver, tshark tools to perform the audit. This tool is customizable to be automated with only a few arguments and can be trusted to run without supervision.

%sex. %s%susage%s:
--  -----

  set function wifipumpkin
  set mode auto
  set ssid ----
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}


func MainOptions() {
    rows := [][]string{
        {"MODULE", "none", "yes", "A module to interact with."},
    }

    fmt.Println(FormatModuleOptions(
        "main/africana_run.fn",
        rows,
        ModuleHelpInfo{Example: "  set MODULE setups\n  run\n"},
    ))
}

func SetupsOptions(Distro, Function string) {
    rows := [][]string{
        {"DISTRO", Distro, "yes", "Distro to install africana on. Supported distros: (arch, ubuntu, macos, android, windows)."},
        {"FUNCTION", Function, "yes", "The function to execute e.g. (install, update, repair, uninstall)."},
    }

    fmt.Println(FormatModuleOptions(
        "setups/setups_launcher.fn",
        rows,
        ModuleHelpInfo{Example: "  set DISTRO " + Distro + "\n  set FUNCTION install\n  run\n"},
    ))
}

func NetworksOptions(Mode, IFace, RHost, Passwd, LHost, Gateway, Spoofer, Proxies, FakeDns, Function string) {
    rows := [][]string{
        {"MODE", Mode, "yes", "Kind of attack to perform (single or all). 'single' will attack one rhost; 'all' will attack entire subnet."},
        {"IFACE", IFace, "yes", "Interface to use for penetration testing."},
        {"RHOST", RHost, "yes", "Alias to (RHOSTS, TARGET, TARGETS). The target to perform functions on."},
        {"LHOST", LHost, "yes", "Needed when using responder, beefkill, or toxssinx to handle reverse connections."},
        {"PASSWD", Passwd, "yes", "Password for beef-xss login page. Default user: beef."},
        {"GATEWAY", Gateway, "yes", "Default router IP address. Needed when running functions like beefkill."},
        {"SPOOFER", Spoofer, "yes", "Tool for spoofing target host to our domain. (e.g., bettercap)."},
        {"PROXIES", Proxies, "no", "Use proxies for traffic (discover, portscan, vulnscan, etc.)."},
        {"FAKEDNS", FakeDns, "yes", "DNS to resolve during spoofing (e.g., for beefkill)."},
        {"FUNCTION", Function, "yes", "Function to perform: portscan, vulnscan, enumscan, smbexpl, psniffer, responder, beefkill."},
    }

    fmt.Println(FormatModuleOptions(
        "src/networks/networks_pentest.fn",
        rows,
        ModuleHelpInfo{Example: "  set IFACE " + IFace + "\n  run\n"},
    ))
}

func ExploitsOptions(Icon, CoreIcon, LHost, LPort, HPort, Script, BuildName, Function, Proxy, BuildDir, Listener, Protocol, Obfuscator string) {
    rows := [][]string{
        {"ICON", Icon, "yes", "Icon to be used while generating backdoors."},
        {"LPORT", LPort, "yes", "Listener port to handle beacons."},
        {"HPORT", HPort, "yes", "Port for file smaglers in blackjack function."},
        {"LHOST", LHost, "yes", "Mainly needed when generating backdoors and launching listeners."},
        {"BUILD", BuildName, "yes", "Output name of the backdoor to be generated."},
        {"SCRIPT", Script, "yes", "Path to PowerShell script to be obfuscated."},
        {"OUTPUT", BuildDir, "yes", "Output location for generated backdoor."},
        {"PROXIES", Proxy, "no", "Route traffic through proxies if desired."},
        {"COREICON", CoreIcon, "yes", "THe hidden backdoor's icon. The one the user will see in the system if curious."},
        {"PROTOCOL", Protocol, "yes", "Communication protocol (tcp, http, https)."},
        {"FUNCTION", Function, "yes", "Function to perform (ghost, shellz, etc.)."},
        {"LISTENER", Listener, "yes", "Listener for callback connections."},
        {"OBFUSCATOR", Obfuscator, "yes", "The tool to obfusicate your backdoor. Use 'show obfuscator' to list them."},
    }

    fmt.Println(FormatModuleOptions(
        "src/exploits/backdoor_pentest.fn",
        rows,
        ModuleHelpInfo{Example: "  set LPORT " + LPort + "\n  set MODULE " + Function + "\n  set LHOST " + LHost + "\n  run\n"},
    ))
}

func BlackJackOptions(LHost string) {
    rows := [][]string{
        {"LHOST", LHost, "yes", "Listener host address."},
        {"LPORT", "9999", "yes", "Listener port to handle beacons."},
        {"HPORT", "3333", "yes", "Port for file smaglers in blacjack."},
        {"SCRIPT", "none", "yes", "Powershell script location."},
        {"PROXIES", "none", "no", "Traffic through proxies."},
        {"PROTOCOL", "tcp", "yes", "Communication protocol."},
    }

    fmt.Println(FormatModuleOptions(
        "src/exploits/blackjack_listener.fn",
        rows,
        ModuleHelpInfo{Example: "  set LHOST " + LHost + "\n  run\n"},
    ))
}

func ShellzOptions(LHost, LPort, Protocol string) {
    rows := [][]string{
        {"LHOST", LHost, "yes", "Mainly needed when generating backdoors."},
        {"LPORT", LPort, "yes", "Listener port to handle beacons."},
        {"PROTOCOL", Protocol, "yes", "Protocol for host communication (tcp, http, https)."},
    }

    fmt.Println(FormatModuleOptions(
        "src/exploits/shellz_listener.fn",
        rows,
        ModuleHelpInfo{Example: "  set LHOST " + LHost + "\n  run\n"},
    ))
}

func AndroRatOptions(LHost, LPort, HPort, BuildName, OutPutDir, Proxies, Protocol string) {
    rows := [][]string{
        {"LHOST", LHost, "yes", "Mainly needed when generating backdoors."},
        {"LPORT", LPort, "yes", "Listener port to handle beacons."},
        {"HPORT", HPort, "yes", "The port to handle file smaglers in blacjack function."},
        {"BUILD", BuildName, "yes", "OutPut name of the backdoor to be generated."},
        {"OUTPUT", OutPutDir, "yes", "OutPut location."},
        {"PROXIES", Proxies, "no", "Run traffic through proxies."},
        {"PROTOCOL", Protocol, "yes", "Protocol for host communication (tcp, http, https)."},
    }

    fmt.Println(FormatModuleOptions(
        "src/exploits/androrat_listener.fn",
        rows,
        ModuleHelpInfo{Example: "  set LHOST " + LHost + "\n  run\n"},
    ))
}

func HoaxshellOptions(LHost, LPort, Protocol string) {
    rows := [][]string{
        {"LHOST", LHost, "yes", "Mainly needed when generating backdoors."},
        {"LPORT", LPort, "yes", "Listener port to handle beacons."},
        {"PROTOCOL", Protocol, "yes", "Protocol for host communication (tcp, http, https)."},
    }

    fmt.Println(FormatModuleOptions(
        "src/exploits/hoaxshell_listener.fn",
        rows,
        ModuleHelpInfo{Example: "  set LHOST " + LHost + "\n  run\n"},
    ))
}

func LithalDllOptions(Icon, LHost, LPort, HPort, Protocol, OutPut, Listener string) {
    rows := [][]string{
        {"ICON", Icon, "yes", "Mainly needed to disguise generated backdoors."},
        {"LPORT", LPort, "yes", "Listener port to handle beacons."},
        {"HPORT", HPort, "yes", "Https Listener port to handle file smaggling."},
        {"LHOST", LHost, "yes", "Your local machine IP. Mainly needed when generating backdoors and handling reverse connections."},
        {"OUTPUT", OutPut, "yes", "Mainly needed when generating backdoors."},
        {"PROTOCOL", Protocol, "yes", "Protocol for host communication (tcp, http, https)."},
    }

    fmt.Println(FormatModuleOptions(
        "src/exploits/lithaldll_listener.fn",
        rows,
        ModuleHelpInfo{Example: "  set LHOST " + LHost + "\n  run\n"},
    ))
}

func WirelessOptions(IFace, LHost, OutPutDir string) {
    rows := [][]string{
        {"IFACE", IFace, "yes", "Mainly needed when generating backdoors."},
        {"LHOST", LHost, "yes", "Listener port to handle beacons."},
        {"OUTPUT", OutPutDir, "yes", "OutPut location."},
    }

    fmt.Println(FormatModuleOptions(
        "src/wirelss/wireless_pentest.fn",
        rows,
        ModuleHelpInfo{Example: "  set IFACE " + IFace + "\n  run\n"},
    ))
}

func CrackersOptions(Mode, RHost, WordListDir, UserName, PassWord string) { 
    rows := [][]string{
        {"MODE", Mode, "yes", "Attack mode (online/offline)."},
        {"RHOST", RHost, "yes", "Target host."},
        {"USERNAME", UserName, "yes", "Single username to test."},
        {"PASSWORD", PassWord, "yes", "Single password to test."},
        {"WORDLIST", WordListDir, "yes", "Path to wordlist."},
    }

    fmt.Println(FormatModuleOptions(
        "src/crackers/passwords_pentest.fn",
        rows,
        ModuleHelpInfo{Example: "  set MODE " + Mode + "\n  run\n"},
    ))
}

func TorsocksOptions() {
    rows := [][]string{
        {"FUNCTION", "none", "yes", "The function to execute. ex. -> (setups, vanish, exitnode, status, ipaddress, restore, reload, chains, stop)"},
    }

    fmt.Println(FormatModuleOptions(
        "src/securities/torrsocks_setup.fn",
        rows,
        ModuleHelpInfo{Example: "  set VANISH vanish\n  run\n"},
    ))
}

func ResponderOptions(Mode, LPort, RHost, LHost string) {
    rows := [][]string{
        {"MODE", Mode, "yes", "Mode to launch. All to attack all subnet, Single to attack only one target."},
        {"LPORT", LPort, "yes", "Listener port to handle beacons."},
        {"RHOST", RHost, "no", "Target to attack. Host's ipadress."},
        {"LHOST", LHost, "no", "Mainly needed when generating backdoors and launching Listeners. Default to your eth0 ip."},
    }

    fmt.Println(FormatModuleOptions(
        "src/networks/responder.fn",
        rows,
        ModuleHelpInfo{Example: "  set MODULE responder\n  set LHOST " + LHost + "\n  run\n"},
    ))
}

func BeefKillOptions(Mode, LPort, Spoofer, RHost, LHost string) {
    rows := [][]string{
        {"MODE", Mode, "yes", "Mode to launch. All to attack all subnet, Single to attack only one target."},
        {"LPORT", LPort, "yes", "Listener port to handle beacons."},
        {"RHOST", RHost, "no", "Target to attack. Host's ipadress."},
        {"LHOST", LHost, "no", "Mainly needed when generating backdoors and launching Listeners. Default to your eth0 ip."},
        {"SPOOFER", Spoofer, "yes", "Listener port to handle beacons."},
    }

    fmt.Println(FormatModuleOptions(
        "src/networks/beefkill.fn",
        rows,
        ModuleHelpInfo{Example: "  set MODULE beefkill\n  set LHOST " + LHost + "\n  run\n"},
    ))
}

func ToxssInxOptions(Mode, LPort, Spoofer, RHost, LHost string) {
    rows := [][]string{
        {"MODE", Mode, "yes", "Mode to launch. All to attack all subnet, Single to attack only one target."},
        {"LPORT", LPort, "yes", "Listener port to handle beacons."},
        {"RHOST", RHost, "no", "Target to attack. Host's ipadress."},
        {"LHOST", LHost, "no", "Mainly needed when generating backdoors and launching Listeners. Default to your eth0 ip."},
        {"SPOOFER", Spoofer, "yes", "Listener port to handle beacons."},
    }

    fmt.Println(FormatModuleOptions(
        "src/networks/toxssinx.fn",
        rows,
        ModuleHelpInfo{Example: "  set MODULE toxssinx\n  set LHOST " + LHost + "\n  run\n"},
    ))
}

func WebsitesOptions(RHost, ResultDir, Proxies, UserName, PassWord, WordListDir string) {
    rows := [][]string{
        {"RHOST", RHost, "yes", "The host to be attacked eg. website."},
        {"OUTPUT", ResultDir, "no", "OutPut location."},
        {"PROXIES", Proxies, "no", "Just incase you want to run your traffic through proxies.."},
        {"USERNAME", UserName, "no", "Single user name to attack on a give service."},
        {"PASSWORD", PassWord, "no", "Single password to use while attacking a given name or password."},
        {"WORDLIST", WordListDir, "yes", "A list of user names or passwords to be used. -> (Give full path)."},
    }

    fmt.Println(FormatModuleOptions(
        "src/websites/bugbounty_pentest.fn",
        rows,
        ModuleHelpInfo{Example: "  set WORDLIST " + WordListDir + "\n  run\n"},
    ))
}

func PhishersOptions(Mode, RHost, Proxies string) {
    rows := [][]string{
        {"MODE", Mode, "no", "Mode to use. Single attack single target or all to attack entire subnet."},
        {"RHOST", RHost, "no", "Mainly needed when generating backdoors."},
        {"PROXIES", Proxies, "no", "Just incase you want to run your traffic through proxies.."},
    }

    fmt.Println(FormatModuleOptions(
        "src/phishers/phishers_pentest.fn",
        rows,
        ModuleHelpInfo{Example: "  set MODULE setoolkit\n  run\n"},
    ))
}


func HelpInfoRun() {
    fmt.Printf(`
%sUsage%s: run [Function]

%sRun a Function%s:
--- - --------
  show modules to list modules or <show all>

%sCommand      Description%s
-------      -----------
run          Enables you to Interact with a Function by name or search term/index. Alias to "use", "exec", "start", "launch", "exploit", "execute".

%sex. %s%susage%s:
--  -----

  run setups or execute setups

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfOptions() {
    fmt.Printf(`
%sUsage%s: options -> to show list of options for a given function or module.
           Same as show options, when 'option' command is invoked does the same.

%sex. %s%susage%s:
--  -----

  show options

Just like show options command, options shows commands in which a sub Function can run in a given sub menu.

`, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoShow() {
    fmt.Printf(`
%sUsage%s: show [all], [modules], [options], [functions]

%sex. %s%susage%s:
--  -----

  show modules

%s[*]%s Valid parameters for the %s"show" %scommand are: all, modules, options
%s[*]%s Additional module-specific parameters are: missing, advanced, evasion, targets, actions

`, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc)
}

func HelpInfoUse() {
    fmt.Printf(`
%sUsage%s: use <name|term>

Just like start and run, use enables you to Interact with a Function by name or search term/index. If a Function
name is not found, it will be treated as a search term. An index from the previous search results can be selected if desired.

%sex. %s%susage%s:
--  -----

  use setups, torsocks, networks, exploits, wireless, phishers, websites, credits, verses

  use setups
  use <name>

`, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoStart() {
    fmt.Printf(`
%sUsage%s: start <name|term>

Just like use and run, start enables you to Interact with a Function by name or search term/index.
If a Function name is not found, it will be treated as a search term. An index from the previous search
results can be selected if desired.

%sex. %s%susage%s:
--  -----

  start setups, torsocks, networks, exploits, wireless, phishers, websites, credits, verses

  start setups
  start <name>

`, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoList() {
    fmt.Printf(`
%sUsage%s: list <name|term>

Show all modules or sub-Functions in a specific face you are. If you are in modules setups,
list modules command, will list all sub-modules available in that Function.

%sex. %s%susage%s:
--  -----

  list modules

%s[*]%s Valid parameters for the "list" command are: modules, functions or all.
%s[*]%s Additional Function-specific parameters are: missing, advanced.

`, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc)
}

func HelpInfoSet() {
    fmt.Printf(`
%sUsage%s: set [options] [name] [value]

Set the given option to value. If value is omitted, print the current value. If both are omitted,
print options that are currently set. ex.(set LPORT 9001)

`, bcolors.Bold, bcolors.Endc)
}

func HelpInfoUnSet() {
    fmt.Printf(`
%sUsage%s: unset [value]

Unset the given option. If value is omitted, print the current value. If both are omitted,
print options that are currently set. Unset is an alias to Delete.(unset option).

`, bcolors.Bold, bcolors.Endc)
}

func HelpInfo() {
    fmt.Printf(`
%sUsage%s: info <Function name>

%sex. %s%susage%s:
--  -----

  info setups -> or info [ 1. setups   2. torsocks 3. networks 4. exploits 5. wireless ]
                         [ 5. crackers 6. phishers 7. websites 8. credits  9. verse.   ] -> Integer input support. keys are. (1 2 3 4 5 6 7 8 9 0 and 10)

%sDescription%s:
-----------

  Queries the supplied Function or modules for information. If no Function is given, show info for the currently active Function.

`, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoTips() {
    fmt.Printf(`
%sUsage%s: getips <Function name>

%sex. %s%susage%s:
--  -----

  info setups, torsocks, networks, exploits, wireless, phishers, webs, verses.

%sDescription%s:
-----------
  Gives you special information about a Function and how to successfully use it in.
real life scenario.

`, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoFeatures() {
    fmt.Printf(`
%sUsage%s: features

%sex. %s%susage%s:
--  -----
  features -> or show features

%sDescription%s:
-----------

Prints feature plans for africana. Ex. modules to be added, what to be corrected.

`, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoExecute() {
    fmt.Printf(`
%sOptions%s:
-------
    %srun%s : Launch the selected Function. Alias to (start, execute)

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoMenuMain() {
    fmt.Printf(`
Usage: ./afrconsole [options]

Common options:
    -a, --auto          Start africana in automation mode 'guide by menu'.

Setup options:
    -i, --install       Launch installation menu to install needed dependencies.
    -u, --update        Update africana framework and and all it's tools.

Framework options:
    -l, -L, --logs      Show logs in terminal.
    -v, -V, --version   Show version.

Console options:
    -t, --torsocks      Launch torsocks menu to torify your system.
    -n, --networks      Start internal network attacks.
    -e, --exploits      Generate undetectebal R.A.Ts and (Launch listeners for all systems. Evasions also included).
    -w, --wireless      From wifi, bluetooth, cantools and other wireles attack vectors.
    -c, --crackers      Crack(NTLMS, HASHES, PCAPS) and bruteforce(SSH, FTP, SMB, RPC etc.).
    -p, --phishers      Perform almost all advanced Phishing attacks.
    -x, --websites      Launch Web Penetration engines with free bugbounty automation function.
    -k, --credits       Show who developes and mentains africana-framework and (third party tools developers).
    -s, --verses        Scirptures. Launch chosen Bible verses as used in the framework.
    -g, --guide         Watch tutarials on %sYouTube %s: %s%shttps://youtube.com/@r0jahsm0ntar1/?sub_confirmation=1%s.
    -q, --quite         Start africana without banner and missing tools checking.
    -h, --help          Show this help message and exit.

`, bcolors.Green, bcolors.Endc, bcolors.Italic, bcolors.Underl, bcolors.Endc)
}

func HelpInfoMenuZero() {
    fmt.Printf(`
    %sCommand             Description%s
    -------             -----------
    ?                   Help menu. Alias to (h, help).
    banner              Display an awesome africana banner.
    clear               Clear the working screen or use with flag ('history' to clear history).
    exit                Exit the console.
    features            Display the list of not yet released features that can be opted in to.
    guide               Watch tutarials on %sYouTube %s: %s%shttps://youtube.com/@r0jahsm0ntar1/?sub_confirmation=1%s.
    history             Show command history.
    menu                Print the menu of the current phase. Alias to letter(m).
    quit                Exit the console.
    set                 Sets a context-specific variable to a value.
    sleep               Do nothing for the specified number of seconds.
    tips                Show a list of useful productivity tips.
    version             Show the framework and console library version numbers.

%sFunction Commands%s
-------- --------

    %sCommand             Description%s
    -------             -----------
    advanced            Displays advanced options for one or more modules.
    back                Move back from the current context.
    info                Displays information about one or more modules.
    list                List the Function stack.
    options             Displays global options or for one or more modules.
    run                 Run a selected Function.
    search              Searches Function names and descriptions.
    show                Displays modules of a given type, or all modules.
    use                 Interact with a module by name or search term/index.

%sDeveloper Commands%s
--------- --------

    %sCommand             Description%s
    -------             -----------
    bash                 Start Intaractive shell in africana. Alias to ('sh, zsh, cmd, powershell').
    junks                Display what is in the output directory.
    logs                 Display framework.log paged to the end if possible.
    time                 Time how long it takes to run a particular command.

For more info on a specific command, use %s<command> -h %sor %shelp <command>%s.

`, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Italic, bcolors.Underl, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
}

func UpsentTools() {
    fmt.Println("\n%s[!] %sChoice selected not implemented yet!, but coming soon!", bcolors.Yellow, bcolors.Endc)
}

