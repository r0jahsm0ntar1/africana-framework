package menus

import (
    "fmt"
    "utils"
    "bcolors"
    "strings"
)

var (
    Interfaces, _ = utils.Ifaces()
    Lhost, _      = utils.GetDefaultIP()
    Gateway, _    = utils.GetDefaultGatewayIP()
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordList = utils.DirLocations()
)

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
    Description   string
    Example       string
    Options       string
}

type TableConfig struct {
    NameWidth    int
    SettingWidth int
    ReqWidth     int
}

var DefaultTableConfig = TableConfig{
    NameWidth:    15,
    SettingWidth: 18,
    ReqWidth:     9,
}

func FormatRow(config TableConfig, name, value, required, desc string) string {
    return fmt.Sprintf(
        "  %-*s  %-*s  %-*s  %s",
        config.NameWidth, name,
        config.SettingWidth, value,
        config.ReqWidth, required,
        desc,
    )
}

func FormatModuleHeader(modulePath string) string {
    return fmt.Sprintf(
        "\n%sModule Options %s(%s%s%s):\n",
        bcolors.Bold, bcolors.Endc,
        bcolors.Green, modulePath, bcolors.Endc,
    )
}

func FormatColumnHeaders(config TableConfig) string {
    return FormatRow(config, "Name", "Current Setting", "Required", "Description") + "\n" +
        FormatRow(config, "----", "---------------", "--------", "-----------")
}

func FormatFooter() string {
    return fmt.Sprintf(`
%sSupported Distros%s:
--------- --------
   Id  Name
   --  ----
   0   All Distros

%sex. %susage%s:
--  -----
  show options
  set LHOST 127.0.0.1
  run

View the full module info with the %s'info'%s or %s'info -d'%s command.
`,
        bcolors.Bold, bcolors.Endc,
        bcolors.BrightBlue, bcolors.Bold, bcolors.Endc,
        bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc,
    )
}

func FormatModuleOptions(modulePath string, config TableConfig, rows [][]string) string {
    var builder strings.Builder

    builder.WriteString(FormatModuleHeader(modulePath))
    builder.WriteString("\n")
    builder.WriteString(FormatColumnHeaders(config))
    builder.WriteString("\n")

    for _, row := range rows {
        if len(row) == 4 {
            builder.WriteString(FormatRow(config, row[0], row[1], row[2], row[3]))
            builder.WriteString("\n")
        }
    }

    builder.WriteString(FormatFooter())

    return builder.String()
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
------ -------
   Id  Name
   --  ----
-> 0   %s

%sDescription%s:
-----------
  %s
`, bcolors.Bold, bcolors.Endc, info.Name, bcolors.Bold, bcolors.Endc, info.Function, bcolors.Bold, bcolors.Endc, info.Platform, bcolors.Bold, bcolors.Endc, info.Arch, bcolors.Bold, bcolors.Endc, info.Privileged, bcolors.Bold, bcolors.Endc, info.License, bcolors.Bold, bcolors.Endc, info.Rank, bcolors.Bold, bcolors.Endc, info.Disclosed, bcolors.Bold, bcolors.Endc, info.ProvidedBy, bcolors.Bold, bcolors.Endc, info.CreatedBy, bcolors.Bold, bcolors.Endc, info.TestedDistros, bcolors.Bold, bcolors.Endc, info.Description)

    if info.Options != "" {
        fmt.Printf("\n%sOptions%s:\n%s\n", bcolors.Bold, bcolors.Endc, info.Options)
    }

    if info.Example != "" {
        fmt.Printf(`
%sex. %s%susage%s:
--  -----
%s
`, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, info.Example)
    }
}

func modulesUsage(info ModuleHelpInfo) {
    if info.Example != "" {
        fmt.Printf(`

%sex. %s%susage%s:
--  -----
%s
`, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, info.Example)
    }
}

func generateMenu(menuItems []string, helpText string, backOption bool) {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

`, bcolors.Italic, bcolors.Underl, bcolors.Bold, bcolors.Endc)

    for i, item := range menuItems {
        fmt.Printf("%s %d. %s%s\n", bcolors.BrightBlue, i + 1, bcolors.Endc, item)
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

func generateOptions(modulePath string, options []string) {
    fmt.Printf(`
%sModule options %s(%s):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
`,bcolors.Bold, bcolors.Endc, modulePath, bcolors.Bold, bcolors.Endc)

    for _, opt := range options {
        fmt.Printf(opt)
    }
}

func MenuZero() {
    items := []string{
        "Install, Update and View Logs",
        "System Security Configuration",
        "Local Network Penetration Testing",
        "Backdoors and Trojans Generating Engines",
        "Wireless Networks Attacks Vector",
        "Passwords, Hashes and .Pcap Crackers",
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
        "dllbackdoor",
        "regbackdoor",
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
%sMain Modules%s:
---- -------

  # %sName        Description%s
  - ----        -----------`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"  setups", "Manage framework installation, updates, dependencies and complete removal."},
        {"torsocks", "Enforce Tor network routing for all traffic with leak protection."},
        {"networks", "Conduct MITM, DNS spoofing and network vulnerability assessments."},
        {"exploits", "Create stealth payloads and establish encrypted C2 connections."},
        {"wireless", "Perform WPA2 cracking, rogue AP attacks and spectrum analysis."},
        {"crackers", "Bruteforce services and crack hashes using advanced techniques."},
        {"phishers", "Deploy convincing phishing sites with automated data capture."},
        {"websites", "Scan for vulnerabilities and exploit web application flaws."},
        {" credits", "View acknowledgments for contributors and integrated tools."},
        {"  verses", "Access scriptural references for ethical hacking guidance."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
        }
    info := ModuleHelpInfo{
        Example: "    set module verses\n    run\n",
    }
    modulesUsage(info)
}

func ListSetupsDistros() {
    fmt.Printf(`
%sSupported Distros%s:
--------- ---------

  # %sName        Description%s
  - ----        -----------`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)

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
        Example:          "    set distro kali\n    set module install\n    run\n",
    }
    modulesUsage(info)
}

func ListSetupsFunction() {
    fmt.Printf(`
%sSetups Functions%s:
------ ---------

  # %sName        Description%s
  - ----        -----------`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)

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
        Example:          "    set module install\n    run\n",
    }
    modulesUsage(info)
}

func ListTorsocksFunctions() {
    fmt.Printf(`
%sTorsocks Functions%s:
-------- ---------

  # %sName        Description%s
  - ----        -----------`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"  setups", "Installs dnsmasq, squid, privoxy and tor (also set configs)"},
        {"  vanish", "Start anonymizing the system through tor network"},
        {"  status", "Check if using tor (Show if all anononimty services are up and running)"},
        {"   torip", "Check current tor IP address"},
        {"  chains", "View traffic logs from squid, privoxy, to tor"},
        {"  reload", "Completely restart torsocks and connect to a different exit-node"},
        {"exitnode", "Connect to a different exit-node"},
        {" restore", "Backup and resets Iptables to default"},
        {"    stop", "Get back to the surface-web"},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
        }
    info := ModuleHelpInfo{
        Example:          "    set module discover\n    run\n",
    }
    modulesUsage(info)
}

func ListInternalFunctions() {
    fmt.Printf(`
%sNetworks Functions%s:
-------- ---------

  # %sName        Description%s
  - ----        -----------`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"discover", "Discover all devices connected to the network. Lets you locate targets."},
        {"portscan", "Get open ports on the target you have set."},
        {"vulnscan", "Perform vulnerbility scan on open ports of the target you have set."},
        {"enumscan", "Recon for S.M.B deep information on the target set."},
        {"smbexpl0", "Launch known vulnerbility exploits on the target's S.M.B services."},
        {"psniffer", "Sniff all Packets from connected devices to the router(Perform M.I.T.M)."},
        {"respond4", "Start Killer Responder that configs all required fields to get you a reverse shell on windows. Supports IPv6."},
        {"beefkill", "Launch Beef-xss and Bettercap/ Ettercp For effective (M.I.B attacks)."},
        {"toxssin1", "Get Shell through XSS Injection to packets in the wire.(To come)."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
        }
    info := ModuleHelpInfo{
        Example:          "    set module discover\n    run\n",
    }
    modulesUsage(info)
}

func ListExploitsFunctions() {
    fmt.Printf(`
%sExploits Functions%s:
-------- ---------

  # %sName           Description%s
  - ----           -----------`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"   androrat", "It is another Android Listener. It is cool but only works on android 4 to 9. Suppoorts android < 14 but less functions"},
        {"  teardroid", "Andoird Listener. Needs alot of online configuration but the best for now"},
        {"  blackjack", "It is a tool derived from villain. It supports both tcp, http and https reverse shells. Supports evasions and bypasses almost all avs"},
        {"  hoaxshell", "A Killer FUD that bypasses most avs"},
        {"     shellz", "Supports all distro reverse shells generation that supports both tcp, https and https connection. Launches variety of listeners"},
        {"      ghost", "It is a giant powershell evasion tool that beats almost all AVS. Try it out you will love it"},
        {"  chameleon", "It is a python framework evasion tool that beats almost all AVS. Try it out you will love it"},
        {"dllbackdoor", "Generates undetected backdoor with embeded hoaxreverse shell. Has .dll persistent mechanisims"},
        {"regbackdoor", "Just like dllbackdoor but the persisten mechanisim is through regestry keys"},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
        }
    info := ModuleHelpInfo{
        Example:          "    set module hoaxshell\n    run\n",
    }
    modulesUsage(info)
}

func ListListenersFunctions() {
    fmt.Printf(`
%sListeners Functions%s:
--------- ---------

  # %sName           Description%s
  - ----           -----------`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"  blackjack", "It is a tool derived from villain. It supports both tcp, http and https reverse shells. Supports evasions and bypasses almost all avs"},
        {"  hoaxshell", "A Killer FUD that bypasses most avs"},
        {"       ncat", "Starts ncat to listen for either tcp or http(s) protocol for powershell reverse shell. Best launched with dllbackdoor or regbackdoor modules"},
        {" metasploit", "Launches metasploit to listen for either tcp or http(s) protocol for powershell reverse shell. Best launched with dllbackdoor or regbackdoor modules"},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
        }
    info := ModuleHelpInfo{
        Example:          "    set module dllbackdoor\n    set listener blackjack\n    run\n",
    }
    modulesUsage(info)
}

func ListObfscatorsFunctions() {
    fmt.Printf(`
%sObfscators Functions%s:
---------- ---------

  # %sName     Description%s
  - ----     -----------`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"ghost", "It is a giant powershell evasion tool that beats almost all AVS. Try it out you will love it"},
        {"chameleon", "It is a python framework evasion tool that beats almost all AVS. Try it out you will love it"},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
        }
    info := ModuleHelpInfo{
        Example:          "    set script Full_path_to_your .ps1\n    run\n",
    }
    modulesUsage(info)
}


func ListIcons() {
    fmt.Printf(`
%sList of Icons%s:
---- -- -----

  # %sName          Description%s
  - ----          -----------`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"    access","Microsoft access document icon for your output malware"},
        {"   autorun","Autorun standard icon to disguise your malicious malware"},
        {"     excel","Microsoft excel document icon for your output malware"},
        {"       pdf","Standard .pdf icon for your malicious backdoor"},
        {"   project","Another icon found in microsoft office tools"},
        {" publisher","Microsoft publisher document icon for your output malware"},
        {"    redrat","A simple icon for a rat. This will fortunetly look suspicious"},
        {"     visio","Microsoft access document icon for your output malware"},
        {"       vlc","An icon that will output your malware as a vlc vidoe player"},
        {"     word","Microsoft word document icon for your output malware"},
        {" defender","Microsoft defender icon for your output malware"},
        {"kaspersky","Kapersky antivirus icon for your output malware"},
    }

    for i, item := range items {
        fmt.Printf(`
  %s%d. %s%-3s%s > %s`, bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
    }
    info := ModuleHelpInfo{
        Example:          "    set icon access\n",
    }
    modulesUsage(info)
}

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
        Description:   "This is the main module containing all africana-framework functions.",
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
        Description:   "This is the Network module that contains all internal networks attacks functions.",
    }
    modulesHelp(info)
    ListInternalFunctions()
}

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
        Description:   "This module will help you in obfuscating any powershell script in order to bypass any AV durring execution.",
    }
    modulesHelp(info)
}

func MainOptions() {
    options := []string{
        "  MODULE         none             yes       Select a module or function to interact with.",
        "  RUN            none             yes       To execute the function. Alias to -> (start, execute, exec, launch).",
        "  show modules   To see modules you can interact with.",
    }
    generateOptions("main/africana_run.fn", options)
    info := ModuleHelpInfo{
        Example:          "    set module verses\n    run\n",
    }
    modulesUsage(info)
}

func SetupsOptions() {
    fmt.Printf(`
%sModule options %s(setups/setups_launcher.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  DISTRO         none             yes       Distro to install africana on.    -> supported distros. (arch, ubuntu, macos, android, windows).
  FUNCTION       none             yes       The function to execute.          -> ex. (Install, update, repair or uninstall).
  RUN            none             yes       To execute the function. Alias to -> (start, execute, exec, launch).

%sex. %s%susage%s:
--  -----
  set distro kali
  set function install
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
}

func TorsocksOptions() {
    fmt.Printf(`
%sModule options %s(src/securities/torrsocks_setup.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  function       none             yes       The function to execute. ex.      -> (setups, vanish, exitnode, status, ipaddress, restore, reload, chains, stop)
  run            none             yes       To execute the function. Alias to -> (start, execute, exec, launch).

%sSupported Distros%s:
--------- --------
   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  -----
  set function vanish
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
}

func NetworksOption() {
    fmt.Printf(`
%sModule options %s(src/networks/networks_pentest.fn):

  %sName      Current Setting  Required  Description%s
  ----      ---------------  --------  -----------
  MODE      single           yes       Kind of attack to perform (single or all) single will attack single rhost, all for entire subnetmask.
  IFACE     eth0             yes       Interface to use for penetration testing.
  RHOST     none             yes       Alias to (RHOSTS, TARGET, TARGETS) The target to perform functions on.
  PASSWD    Jesus            yes       The password to set for beef-xss login page with the default user:beef
  LHOST     ->               yes       %sDefault%s: %s%s%s. Mainly needed if you need to use (responder, beefninja and xsshooker to handle reverse calls.
  GATEWAY   ->               yes       %sDefault%s: %s%s%s. The default roouter ipaddres. Will be needed when running functions like (beefninja).
  SPOOFER   ettercap         yes       The tool to be used for spoofing target host to our domain. the other (bettercap).
  PROXIES    none             no        Just incase you want to run your traffic through proxies. -> (discover, portscan, vulnscan, enumscan, smbexpl, psniffer, responder, beefninja, xsshooker).
  FAKEDNS   *                yes       Dns to be resolved when performing dns spoof. Mainly needed when beefninja function is at hand.
  FUNCTION  none             yes       The function you want network module to perform. ex. (portscan, vulnscan, enumscan, smbexpl, psniffer, responder, beefninja).

%sSupported Distros%s:
--------- --------
   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  -----
  set function discover
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc,       bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc,           bcolors.Bold, bcolors.Endc, bcolors.Yellow, Gateway, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
}

func OptionsResponder() {
    fmt.Printf(`
%sModule Options %s(src/networks/responder.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  RHOST          none             yes       Target to attack. Host's ipadress.
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  PROXIES        none             no        Just incase you want to run your traffic through proxies.

%sSupported Distros%s:
--------- --------
   Id  Name
   --  ----
   0   All Browsers

%sex. %s%susage%s:
--  -----
  set LHOST 127.0.0.1
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
}

func OptionsBeefNinja() {
    fmt.Printf(`
%sModule Options %s(src/networks/beefninja.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  RHOST          none             yes       Target to attack. Host's ipadress.
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  PROXIES        none             no        Just incase you want to run your traffic through proxies.
  SPOOFER        ettercap         yes       Tool to be used to spoof dns and repond to them. ex. (bettercap)

%sSupported Distros%s:
--------- --------
   Id  Name
   --  ----
   0   All Browsers

%sex. %s%susage%s:
--  -----
  set LHOST 127.0.0.1
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
}

func OptionsXssHoocker() {
    fmt.Printf(`
%sModule Options %s(src/networks/xsshooker.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  RHOST          none             yes       Target to attack. Host's ipadress.
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  PROXIES        none             no        Just incase you want to run your traffic through proxies.

%sSupported Distros%s:
--------- --------
   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  -----
  set LHOST 127.0.0.1
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
}

func NetworksOptions(Mode, IFace, RHost, Passwd, Lhost, Gateway, Spoofer, Proxies, FakeDns, Function string) {
    rows := [][]string{
        {"MODE", Mode, "yes", "Kind of attack to perform (single or all) single will attack single rhost, all for entire subnetmask."},
        {"IFACE", IFace, "yes", "Interface to use for penetration testing."},
        {"RHOST", RHost, "yes", "Alias to (RHOSTS, TARGET, TARGETS) The target to perform functions on."},
        {"PASSWD", Passwd, "yes", "The password to set for beef-xss login page with the default user:beef"},
        {"LHOST", Lhost, "yes", "Mainly needed if you need to use (responder, beefninja and xsshooker to handle reverse calls."},
        {"GATEWAY", Gateway, "yes", "The default router ipaddress. Will be needed when running functions like (beefninja)."},
        {"SPOOFER", Spoofer, "yes", "The tool to be used for spoofing target host to our domain. The other (bettercap)."},
        {"PROXIES", Proxies, "no", "Just in case you want to run your traffic through proxies. -> (discover, portscan, vulnscan, enumscan, smbexpl, psniffer, responder, beefninja, xsshooker)."},
        {"FAKEDNS", FakeDns, "yes", "DNS to be resolved when performing DNS spoof. Mainly needed when beefninja function is at hand."},
        {"FUNCTION", Function, "yes", "The function you want network module to perform. ex. (portscan, vulnscan, enumscan, smbexpl, psniffer, responder, beefninja)."},
    }

    fmt.Println(FormatModuleOptions(
        "src/networks/networks_pentest.fn",
        DefaultTableConfig,
        rows,
    ))
}

func ExploitsOptions(Icon, Lhost, Lport, Hport, Script, Malware, Function, Proxy, OutPutDir, Listener, Protocol string) {
    rows := [][]string{
        {"ICON", Icon, "yes", "Icon to be used while generating backdoors"},
        {"LHOST", Lhost, "yes", "Mainly needed when generating backdoors and launching Listeners"},
        {"LPORT", Lport, "yes", "Listener port to handle beacons"},
        {"HPORT", Hport, "yes", "The port to handle file smaglers in blacjack function"},
        {"BUILD", Malware, "yes", "Output name of the backdoor to be generated"},
        {"SCRIPT", Script, "yes", "Your powershell script location to be opfsicated"},
        {"OUTPUT", OutPutDir, "yes", "Location for generated backdoor"},
        {"PROXIES", Proxy, "no", "Run traffic through proxies"},
        {"PROTOCOL", Protocol, "yes", "Protocol for host communication (tcp, http, https)"},
        {"LISTENER", Listener, "yes", "Listener for call back connections"},
        {"FUNCTION", Function, "yes", "Function to perform (ghost, shellz, etc)"},
    }

    fmt.Println(FormatModuleOptions(
        "src/exploits/backdoor_pentest.fn",
        DefaultTableConfig,
        rows,
    ))
}

func BlackJackOptions(Lhost string) {
    rows := [][]string{
        {"LHOST", Lhost, "yes", "Listener host address"},
        {"LPORT", "9999", "yes", "Listener port to handle beacons"},
        {"HPORT", "3333", "yes", "Port for file smaglers in blacjack"},
        {"SCRIPT", "none", "yes", "Powershell script location"},
        {"PROXIES", "none", "no", "Traffic through proxies"},
        {"PROTOCOL", "tcp", "yes", "Communication protocol"},
    }

    fmt.Println(FormatModuleOptions(
        "src/exploits/blackjack_listener.fn",
        DefaultTableConfig,
        rows,
    ))
}


func AndroRatOptions(Lhost, Lport, Hport, Malware, OutPutDir, Proxies, Protocol string) {
    rows := [][]string{
        {"LHOST", Lhost, "yes", "Mainly needed when generating backdoors."},
        {"LPORT", Lport, "yes", "Listener port to handle beacons."},
        {"HPORT", Hport, "yes", "The port to handle file smaglers in blacjack function."},
        {"BUILD", Malware, "yes", "Output name of the backdoor to be generated."},
        {"OUTPUT", OutPutDir, "yes", "Output location."},
        {"PROXIES", Proxies, "no", "Run traffic through proxies."},
        {"PROTOCOL", Protocol, "yes", "Protocol for host communication (tcp, http, https)."},
    }

    config := TableConfig{
        NameWidth:    12,
        SettingWidth: 18, 
        ReqWidth:     9,
    }

    fmt.Println(FormatModuleOptions(
        "src/exploits/androrat_listener.fn",
        config,
        rows,
    ))
}

func HelpInfoDiscover() {
    fmt.Printf(`
       %sName%s: discover
   %sFunction%s: src/securities/discover.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: Yes
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024-03-14

%sProvided by%s: <rojahsmontari@gmail.com>
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sBasic options%s:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  function     none             yes       The function to execute. ex. -> (setups, vanish, status, torip, chains, reload, exitnode, restore and stop)

%sex. %s%susage%s:
--  -----
  set function discover
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will scan for all connected devices in the network given using bettercap then arrange the targets
  in a table for you to select one to attack further.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoInPortScan() {
    fmt.Printf(`
       %sName%s: portscan
   %sFunction%s: src/securities/portscan.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: Yes
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024-03-14

%sProvided by%s: <rojahsmontari@gmail.com>
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sBasic options%s:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  function     none             yes       The function to execute. ex. -> (setups, vanish, status, torip, chains, reload, exitnode, restore and stop)

%sex. %s%susage%s:
--  -----
  set function portscan
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will scan all open ports of the target to reveal open ports.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoInVulnScan() {
    fmt.Printf(`
       %sName%s: vulnscan
   %sFunction%s: src/securities/vulnscan.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: Yes
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024-03-14

%sProvided by%s: <rojahsmontari@gmail.com>
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sBasic options%s:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  function     none             yes       The function to execute. ex. -> (setups, vanish, status, torip, chains, reload, exitnode, restore and stop)

%sex. %s%susage%s:
--  -----
  set function vulnscan
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Perform vulnerbility scan on open ports of the target you have set.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoInEnumScan() {
    fmt.Printf(`
       %sName%s: enumscan
   %sFunction%s: src/securities/enumscan.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: Yes
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024-03-14

%sProvided by%s: <rojahsmontari@gmail.com>
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sBasic options%s:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  function     none             yes       The function to execute. ex. -> (setups, vanish, status, torip, chains, reload, exitnode, restore and stop)

%sex. %s%susage%s:
--  -----
  set function enumscan
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Digg for S.M.B deep information on the target set.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoSmbExplo() {
    fmt.Printf(`
       %sName%s: smbexplo
   %sFunction%s: src/securities/smbexplo.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: Yes
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024-03-14

%sProvided by%s: <rojahsmontari@gmail.com>
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sBasic options%s:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  function     none             yes       The function to execute. ex. -> (setups, vanish, status, torip, chains, reload, exitnode, restore and stop)

%sex. %s%susage%s:
--  -----
  set function smbexplo
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Launch known vulnerbility exploits on the target's S.M.B services.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoPSniffer() {
    fmt.Printf(`
       %sName%s: psniffer
   %sFunction%s: src/securities/psniffer.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: Yes
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024-03-14

%sProvided by%s: <rojahsmontari@gmail.com>
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sBasic options%s:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  function     none             yes       The function to execute. ex. -> (setups, vanish, status, torip, chains, reload, exitnode, restore and stop)

%sex. %s%susage%s:
--  -----
  set function psniffer
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Sniff all Packets from connected devices to the router(Perform M.I.T.M).

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoResponder() {
    fmt.Printf(`
       %sName%s: responder
   %sFunction%s: src/securities/responder.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: Yes
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024-03-14

%sProvided by%s: <rojahsmontari@gmail.com>
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sBasic options%s:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  function     none             yes       The function to execute. ex. -> (setups, vanish, status, torip, chains, reload, exitnode, restore and stop)

%sex. %s%susage%s:
--  -----
  set function responder
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Launch reponder asking for your LHOST, Configuring Wpadscript and weponizing it self. Attack supports alot of windows recent version

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoBeefNinja() {
    fmt.Printf(`
       %sName%s: beefninja
   %sFunction%s: src/securities/beefninja.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: Yes
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024-03-14

%sProvided by%s: <rojahsmontari@gmail.com>
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sBasic options%s:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  function     none             yes       The function to execute. ex. -> (setups, vanish, status, torip, chains, reload, exitnode, restore and stop)

%sex. %s%susage%s:
--  -----
  set function beefninja
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Launch a Combination of both beef-xss and bettercap in a unique way to inject hook.js in either one or all targets. All settings are done for you.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoXssHoocker() {
    fmt.Printf(`
       %sName%s: xsshooker
   %sFunction%s: src/securities/xsshooker.fn
   %sPlatform%s: All
       %sArch%s: x64, x86, amd_64, android
 %sPrivileged%s: Yes
    %sLicense%s: Africana Framework License(BSD)
       %sRank%s: Normal
  %sDisclosed%s: 2024-03-14

%sProvided by%s: <rojahsmontari@gmail.com>
 %sCreated by%s: r0jahsm0ntar1

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sBasic options%s:
  Name           Current Setting  Required  Description
  ----           ---------------  --------  -----------
  function     none             yes       The function to execute. ex. -> (setups, vanish, status, torip, chains, reload, exitnode, restore and stop)

%sex. %s%susage%s:
--  -----
  set function xsshooker
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will try to Get you a revers Shell through XSS Injection. Still Working on this Option.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sDescription%s:
-----------
  This is the Exploits module that contains all Listener, backdoors and obfsicatorsfunctions.
`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
    ListExploitsFunctions()
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
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool derived from villain framework. It supports both tcp, http and https reverse shells. It has inbuild evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool like villein or blackjack that supports both tcp, http and https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoDllBackDoor() {
    fmt.Printf(`
       %sName%s: dllbackdoor
   %sFunction%s: src/exploits/dllbackdoor.fn
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

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoRegBackDoor() {
    fmt.Printf(`
       %sName%s: regbackdoor
   %sFunction%s: src/exploits/dllbackdoor.fn
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

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, Lhost, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, OutPutDir, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}




func ListWirelessFunctions() {
    fmt.Printf(`
%sNetworks Functions%s:
-------- ---------

  # %sName       Description%s
  - ----       -----------
  %s1. %s   wifite %s> Wifite is a tool to audit WEP or WPA encrypted wireless networks.
  %s2. %s  fluxion %s> Fluxion is a tool to audit WEP or WPA encrypted wireless networks. Only manual option is supported by now.
  %s3. %sbettercap %s> Bettercap is a tool to audit Internal network and wirekless network like, WEP or WPA encrypted wireless networks.
  %s4. %sairgeddon %s> Airgeddon Fluxion is a tool to audit WEP or WPA encrypted wireless networks. Only manual option is supported by now.
  %s5. %swifipumpk %s> wifipumpkin - Is a Powerful framework for rogue access point attack. This option run automated mode directly.
  %s6. %sKillerpuk %s> This option runs wifipumpkin3 in manual mode where africana sets everything for you.
  %s7. %s   To add %s> Not yet implimented.
  %s8. %s   To add %s> Not yet implimented.
  %s9. %s   To add %s> Not yet implimented.

%sex. %s%susage%s:
--  -----

    set function wifite
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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

%sSupported Distros%s:
--------- --------
      Id  Name
      --  ----
   -> 0   All Distros

%sDescription%s:
-----------
  This is the wireless module containing all wireless pentesting functions.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
    ListWirelessFunctions()
}

func WirelessOptions() {
    fmt.Printf(`
%sModule Options %s(src/wirelss/wireless_pentest.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  IFACE          ->               yes       %sDefault%s: %swlan0%s. Mainly needed when generating backdoors and launching Listeners.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  SCRIPT         none             yes       Your powershell script location to be opfsicated. Mostly neede when using (ghos).
  PROXIES        none             no        Just incase you want to run your traffic through proxies.
  PROTOCOL       tcp              yes       The kind of protocol to be use while communicating to your host machine. (tcp, http or https).

%sSupported Distros%s:
--------- --------
   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  -----
  set IFACE wlan0
  set FUNCTION wifite
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Yellow, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
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
  WORDLISTS      rockyou.txt      yes       Wordlist location for cracking captured handshakes.

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
  WORDLISTS      rockyou.txt      yes       Wordlist location for cracking captured handshakes.

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
  WORDLISTS      rockyou.txt      yes       Wordlist location for cracking captured handshakes.

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
  WORDLISTS      rockyou.txt      yes       Wordlist location for cracking captured handshakes.

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
  WORDLISTS      rockyou.txt      yes       Wordlist location for cracking captured handshakes.

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

func HelpInfoCrackersModules(Module string) {
    info := ModuleHelpInfo{
        Name:          Module,
        Function:      "src/crackers",
        Platform:      "All",
        Arch:          "x64, x86, amd_64, android",
        Privileged:    "No",
        License:       "Africana Framework License(BSD)",
        Rank:          "Normal",
        Disclosed:     "2024",
        CreatedBy:     "r0jahsm0ntar1",
        TestedDistros: "All Distros",
        Description:   fmt.Sprintf("%s module for credential bruteforcing using large wordlists", Module),
        Example:       "set function " + Module + "\nrun",
    }
    modulesHelp(info)
}

func CrackersOptions() {
    options := []string{
        "\n  MODE           none             yes       Attack mode (online/offline)\n",
        "  RHOST          none             yes       Target host\n",
        "  WORDLIST       rockyou.txt      yes       Path to wordlist\n",
        "  USERNAME       root             yes       Single username to test\n",
        "  PASSWORD       password         yes       Single password to test\n\n",
    }
    generateOptions("src/crackers/passwords_pentest.fn", options)
}

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
        Description:   "Crackers is a module enriched with creative attacking faces to help redtemers in successfully cracking or brutforce passwords from services online or local encripted files.",
    }
    modulesHelp(info)
    ListCrackersFunctions()
}


func ListCrackersFunctions() {
    fmt.Printf(`
%sCrackers Functions%s:
-------- ---------

  # %sName          Description%s
  - ----          -----------

  %sOnline Crackers%s:
  ------ --------
  %s1. %s         ssh %s> Automated Bruteforcer for SSH using rockyou.txt wordlists.
  %s2. %s         ftp %s> Bruteforcer for FTP using rockyou.txt wordlists.
  %s3. %s         smb %s> Hydra Bruteforcer for SMB using rockyou.txt wordlists.
  %s4. %s         rdp %s> Bruteforcer for RDP using rockyou.txt wordlists.
  %s5. %s        ldap %s> Hydra Bruteforcer for LDAP using rockyou.txt wordlists.
  %s6. %s        smtp %s> Bruteforcer for SMTP using rockyou.txt wordlists.
  %s7. %s      telnet %s> Bruteforcer for TELNET using rockyou.txt wordlists.
  %s8. %s  http/https %s> Hydra Bruteforcer for HTTP/ HTTPS using rockyou.txt wordlists. Forms needed.
  %s9. %sall services %s> The Automatic Bruteforce Tool for all opened services. Works nice and automatic.

  %sOffline Crackers%s:
  ------- --------
  %s1. %s        pcap %s> Crack captured .pcap files. Full location is needed.
  %s2. %s        ntlm %s> Crack ntlm file using default wordlists.
  %s3. %s        hash %s> Auto identify hash and start bruteforcing for passwords.

%sex. %s%susage%s:
--  -----
  set function ssh
  set mode online
  set wordlist -> (Give full path)
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}


func PhishersOptions() {
    fmt.Printf(`
%sModule Options %s(src/phishers/phishers_pentest.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------

  MODE           none             yes       Mode to use. modes are(online or offline) Online attack for remote services. Offline attack for local files.
  RHOST          none             yes       Target host to be attacked.
  PROXIES        none             no        Just incase you want to run your traffic through proxies.
  FUNCTION       none             yes       The module or function to run. (ex. ssh, ftp, smb, rdp, ldap, smtp, telnet, http, https, auto)
  USERNAME       root             yes       Single user name to attack on a give service.
  PASSWORD       password         yes       Single password to use while attacking agiven name or password
  WORDLIST       rockyou.txt      yes       Alist of user names or passwords to be used.

%sSupported Distros%s:
--------- --------
   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  -----
  set function ssh
  set mode online
  set wordlist -> (Give full path)
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
}

func ListPhishersFunctions() {
    fmt.Printf(`
%sPhishers Functions%s:
-------- ---------

  # %sName          Description%s
  - ----          ----------- `, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)

    items := []struct {
        name string
        desc string
    }{
        {"gophish   ", "It is a phishing framework with a Web UI https://127.0.0.1:3333. Africana will launch it for you. Default  is: admin, Default password is: kali-gophish."},
        {"goodginx  ", "Goodginx is an advanced phishing framework with insane configurations.Default name evilginx2. Bypasses alot of security features like OTP."},
        {"zphisher  ", "A nice framework with alot of templets. Also bypasses OTP with ngrock support."},
        {"blackeye  ", "Writen in bash and full of phishing templets. Just check it out."},
        {"advphisher", "Wide range of phishing templets."},
        {"darkphish ", "Bypasses OTP with Wide range of phishing templets."},
        {"shellphish", "Start Killer Responder that configs all required fields to get you a reverse shell on windows. Supports IPv6."},
        {"setoolkit ", "This tool is equiped with alot of social engeneering. Supports cloning of actual websites."},
        {"thc       ", "This tool creates a templete of your interest imidietly but needs you to start your server and generate a link for phishing."},
    }

    for i, item := range items {
        fmt.Printf("\n %s%2d. %s%-8s%s > %s", bcolors.BrightBlue, i + 1, bcolors.Yellow, item.name, bcolors.Endc, item.desc)
        }
    info := ModuleHelpInfo{
        Example:          "    set module gophish\n    run\n",
    }
    modulesUsage(info)
}

func HelpInfoPhishers() {
    fmt.Printf(`
       %sName%s: phishers
   %sFunction%s: src/phishers/phishers.fn
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
  Phishers is a module enriched with creative attacking faces to help redtemers in successfully Perform phishing attacks with ease.
`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
    ListPhishersFunctions()
}

func HelpInfoGoPhish() {
    fmt.Printf(`
       %sName%s: gophish
   %sFunction%s: src/phishers/gophish.fn
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
  It is a Function that enables the redteamers to perform phising attacks on various bases.

%sex. %s%susage%s:
--  -----

  set function gophish
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoGoodGinx() { 
    fmt.Printf(`
       %sName%s: goodginx
   %sFunction%s: src/phishers/goodginx.fn
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
  It is a Function that enables the redteamers to perform phising attacks on various bases.

%sex. %s%susage%s:
--  -----

  set function goodginx
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoZPhisher() {
    fmt.Printf(`
       %sName%s: zphisher
   %sFunction%s: src/phishers/zphisher.fn
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
  It is a Function that enables the redteamers to perform phising attacks on various bases.

%sex. %s%susage%s:
--  -----

  set function zphisher
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoBlackEye() {
    fmt.Printf(`
       %sName%s: blackeye
   %sFunction%s: src/phishers/blackeye.fn
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
  It is a Function that enables the redteamers to perform phising attacks on various bases.

%sex. %s%susage%s:
--  -----

  set function blackeye
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}


func HelpInfoAdvPhisher() {
    fmt.Printf(`
       %sName%s: advphisher
   %sFunction%s: src/phishers/advphisher.fn
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
  It is a Function that enables the redteamers to perform phising attacks on various bases.

%sex. %s%susage%s:
--  -----

  set function advphisher
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}


func HelpInfoDarkPhish() { 
    fmt.Printf(`
       %sName%s: darkphish
   %sFunction%s: src/phishers/darkphish.fn
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
  It is a Function that enables the redteamers to perform phising attacks on various bases.

%sex. %s%susage%s:
--  -----

  set function darkphish
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoShellPhish() {
    fmt.Printf(`
       %sName%s: shellphish
   %sFunction%s: src/phishers/shellphish.fn
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
  It is a Function that enables the redteamers to perform phising attacks on various bases.

%sex. %s%susage%s:
--  -----

  set function shellphish
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoSetoolKit() {
    fmt.Printf(`
       %sName%s: setoolkit
   %sFunction%s: src/phishers/setoolkit.fn
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
  It is a Function that enables the redteamers to perform phising attacks on various bases.

%sex. %s%susage%s:
--  -----

  set function setoolkit
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoThc() {
    fmt.Printf(`
       %sName%s: thc
   %sFunction%s: src/phishers/thehacker_choice.fn
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
  It is a Function that enables the redteamers to perform phising attacks on various bases.

%sex. %s%susage%s:
--  -----

  set function thc
  run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func ListWebsitesFunctions() {
    fmt.Printf(`
%sWebsites Functions%s:
-------- ---------

  # %sName      Description%s
  - ----      -----------

  %s1. %s  netmap %s> Network Mapping and portscaning.
  %s2. %senumscan %s> Subdomain enumeration.
  %s3. %sdnsrecon %s> Dns and asn lookup.
  %s4. %stechscan %s> Web technology detection.
  %s5. %sasetscan %s> Wayback and asset discovery.
  %s6. %sfuzzscan %s> Digg for root files from the server.
  %s7. %sleakscan %s> Secret and leak detection.
  %s8. %svulnscan %s> Vulnerability scanning (XSS, SQLi, CSRF, SSRF, IDOR).
  %s9. %s  bounty %s> Bug bounty hunting techniques.

%sex. %s%susage%s:
--  -----

    set module enumscan or -> use enumscan
    set rhosts https://example.com
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Yellow, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoWebsites() {
    fmt.Printf(`
       %sName%s: websites
   %sFunction%s: src/websites/bug_bounty.fn
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
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane web attacks with ease.
  It consists off recons, vulners, ddos among others.
`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
    ListWebsitesFunctions()
}

func HelpInfoEnumScan() {
    fmt.Printf(`
       %sName%s: enumscan
   %sFunction%s: src/websites/enum_scan.fn
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
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane enumscan with ease.
  It consists off recons, vulners, ddos among others.

%sex. %s%susage%s:
--  -----

    set module enumscan or -> use enumscan
    set rhosts https://example.com
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}


func HelpInfoDnsRecon() {
    fmt.Printf(`
       %sName%s: dnsrecon
   %sFunction%s: src/websites/dns_recon.fn
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
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane enumscan with ease.
  It consists off recons, vulners, ddos among others.

%sex. %s%susage%s:
--  -----

    set module dnsrecon or -> use dnsrecon
    set rhosts https://example.com
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}


func HelpInfoPortScan() {
    fmt.Printf(`
       %sName%s: portscan
   %sFunction%s: src/websites/port_scan.fn
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
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane enumscan with ease.
  It consists off recons, vulners, ddos among others.

%sex. %s%susage%s:
--  -----

    set module portscan or -> use portscan
    set rhosts https://example.com
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoTechScan() {
    fmt.Printf(`
       %sName%s: techscan
   %sFunction%s: src/websites/tech_scan.fn
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
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane enumscan with ease.
  It consists off recons, vulners, ddos among others.

%sex. %s%susage%s:
--  -----

    set module techscan or -> use techscan
    set rhosts https://example.com
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}


func HelpInfoFuzzScan() {
    fmt.Printf(`
       %sName%s: fuzzscan
   %sFunction%s: src/websites/fuzz_scan.fn
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
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane enumscan with ease.
  It consists off recons, vulners, ddos among others.

%sex. %s%susage%s:
--  -----

    set module fuzzscan or -> use fuzzscan
    set rhosts https://example.com
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoLeakScan() {
    fmt.Printf(`
       %sName%s: leakscan
   %sFunction%s: src/websites/leak_scan.fn
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
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane enumscan with ease.
  It consists off recons, vulners, ddos among others.

%sex. %s%susage%s:
--  -----

    set module leakscan or -> use leakscan
    set rhosts https://example.com
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}


func HelpInfoVulnScan() {
    fmt.Printf(`
       %sName%s: vulnscan
   %sFunction%s: src/websites/vuln_scan.fn
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
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane enumscan with ease.
  It consists off recons, vulners, ddos among others.

%sex. %s%susage%s:
--  -----

    set module vulnscan or -> use vulnscan
    set rhosts https://example.com
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoAsetScan() {
    fmt.Printf(`
       %sName%s: asetscan
   %sFunction%s: src/websites/aset_scan.fn
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
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane enumscan with ease.
  It consists off recons, vulners, ddos among others.

%sex. %s%susage%s:
--  -----

    set module asetscan or -> use asetscan
    set rhosts https://example.com
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}


func HelpInfoAutoScan() {
    fmt.Printf(`
       %sName%s: autoscan
   %sFunction%s: src/websites/auto_scan.fn
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
  Websites is a module enriched with creative attacking faces to help redtemers in successfully Perform insane enumscan with ease.
  It consists off recons, vulners, ddos among others.

%sex. %s%susage%s:
--  -----

    set module autoscan or -> use autoscan
    set rhosts https://example.com
    run

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func WebsitesOptions() {
    fmt.Printf(`
%sModule Options %s(src/websites/bugbounty_pentest.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------

  RHOST          none             yes       Target host to be attacked.
  PROXIES        none             no        Just incase you want to run your traffic through proxies.
  FUNCTION       none             yes       The module or function to run. (ex. netmap, dnsrecon, techscan, asetscan, fuzzscan, leakscan, vulnscan, bounty)
  USERNAME       root             yes       Single user name to attack on a give service.
  PASSWORD       password         yes       Single password to use while attacking agiven name or password
  WORDLIST       rockyou.txt      yes       Alist of user names or passwords to be used. -> (Give full path)

%sSupported Distros%s:
--------- --------
   Id  Name
   --  ----
   0   Websites

%sex. %s%susage%s:
--  -----
  set function netmap
  set rhost example.com
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.BrightBlue, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Green, bcolors.Endc, bcolors.Green, bcolors.Endc)
}

func HelpInfoCredits() {
    fmt.Printf(`
       %sName%s: credits
   %sFunction%s: src/core/credits
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
  Credits is a module to print on third party tools with their authors. It enables africana to acknowledge each developer for his/her hard work.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
}

func HelpInfoVerses() {
    fmt.Printf(`
       %sName%s: verses
   %sFunction%s: src/core/scriptures
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
  Verses is a module to narate the Bible scriptures one by one. It enables africana developer to acknowledge our LORD GOD JESUS CHRIST for Creating everything including you, me and everything.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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
print options that are currently set.

If run from a Function context, this will set the value in the Function's datastore: -> Use -g to operate on the global datastore.

If setting a PAYLOAD, this command can take an index from 'show payloads'.

%sOPTIONS%s:
-------

    -c, --clear   Clear the values, explicitly setting to nil(default)
    -g, --global  Operate on global datastore variables
    -h, --help    Help banner.

`, bcolors.Bold, bcolors.Endc, bcolors.Bold, bcolors.Endc)
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
    -g, --guide         Watch tutarials on %sYouTube %s: %s%shttps://youtube.com/@r0jahsm0ntar1%s.
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
    guide               Watch tutarials on %sYouTube %s: %s%shttps://youtube.com/@RojahsMontari%s.
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
    fmt.Println(bcolors.Yellow + "\n[!] " + bcolors.Endc + "Choice selected not implemented yet!, but coming soon!" + bcolors.Endc)
}
