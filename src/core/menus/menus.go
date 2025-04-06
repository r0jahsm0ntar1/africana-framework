package menus

import(
    "fmt"
    "utils"
    "bcolors"
)
var (
    Interfaces, _ = utils.Ifaces()
    Lhost, _ = utils.GetDefaultIP()
    Gateway, _ = utils.GetDefaultGatewayIP()
    CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordList = utils.DirLocations()
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

%s 1. %sarch      -> Using blackarch repos.
%s 2. %skali      -> Using apt-get.
%s 3. %smacos     -> Install in Docker image.
%s 4. %subuntu    -> Install in docker due to package breackages.
%s 5. %sandroid   -> Will use termux and install in chroot environment.
%s 6. %swindows   -> Still working on it. Africana will use commando vm.
%s 7. %supdate    -> Git pull . will be used to get new version of africana.
%s 8. %srepair    -> Will check for issues and try to fix.
%s 9. %suninstall -> All installed africana packages will be uninstalled.

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuTwo() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %ssetups    -> Will install all necesary packages and configure them.
%s 2. %svanish    -> Start anonimity through tor.
%s 3. %sstatus    -> Check if all anonimity softwares are up and running.
%s 4. %storip     -> Check for external tor IP and status.
%s 5. %schains    -> Display all outgoing traffick through proxies.
%s 6. %sreload    -> Restart all tor services and reconnect afresh.
%s 7. %sexitnode  -> Change your current exit node.
%s 8. %srestore   -> Bring back all default Iptables and other configs files.
%s 9. %sstop      -> Disconect from tor network and Kill all anonimty services.

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

  %sOnline Crackers%s:%s    Offline Crackers%s:
  ----------------    -----------------
  %s1. %ssh               %s10. %spcap
  %s2. %sftp              %s11. %sntlm
  %s3. %ssmb              %s12. %shash
  %s4. %srdp
  %s5. %sldap
  %s6. %ssmtp
  %s7. %stelnet
  %s8. %shttp/https
  %s9. %sall services

%s h. %s%s%s%sGet help.%s %se. %s%s%s%sExit afr3%s %s0. %s%s%s%sGo back.%s

`, bcolors.ITALIC, bcolors.UNDERL, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.UNDERL, bcolors.BOLD, bcolors.ITALIC, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.ITALIC, bcolors.BOLD, bcolors.UNDERL, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ITALIC, bcolors.UNDERL, bcolors.ENDC)
}

func MenuSeven() {
    fmt.Printf(`
%s%s%sSelect a number from the table below.%s

%s 1. %sgophish
%s 2. %sgoodginx
%s 3. %szphisher
%s 4. %sblackeye
%s 5. %sadvphisher
%s 6. %sdarkphish
%s 7. %sshellphish
%s 8. %ssetoolkit
%s 9. %sthc -> or TheHackerChoice

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

func HelpInfoMenuMain() {
    fmt.Printf(`
Usage: ./africana [options]

Common options:
    -a, --auto          Start africana in automation mode 'start from main menu'

Setup options:
    -i, --install       Launch installation menu to install neede dependencies
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
    -h, --help          Show this help message and exit.
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
%sUsage%s: run [Function]

%sRun a Function%s:
--- - --------
  show modules to list modules or <show all>

%sCommand      Description%s
-------      -----------
run          Enables you to Interact with a Function by name or search term/index. Alias to "use", "exec", "start", "launch", "exploit", "execute".

%sex. %s%susage%s:
--  ------

  run setups or execute setups

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfOptions() {
    fmt.Printf(`
%sUsage%s: options -> to show list of options for a given function or module.
           Same as show options, when 'option' command is invoked does the same.

%sex. %s%susage%s:
--  ------

  show options

Just like show options command, options shows commands in which a sub Function can run in a given sub menu.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoShow() {
    fmt.Printf(`
%sUsage%s: show [all], [modules], [options], [functions]

%sex. %s%susage%s:
--  ------

  show modules

%s[*]%s Valid parameters for the %s"show" %scommand are: all, modules, options
%s[*]%s Additional module-specific parameters are: missing, advanced, evasion, targets, actions

`, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC)
}

func HelpInfoUse() {
    fmt.Printf(`
%sUsage%s: use <name|term>

Just like start & run, use enables you to Interact with a Function by name or search term/index. If a Function
name is not found, it will be treated as a search term. An index from the previous search results can be selected if desired.

%sex. %s%susage%s:
--  ------

  use setups, torsocks, networks, exploits, wireless, phishers, websites, credits, verses

  use setups
  use <name>

`, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoStart() {
    fmt.Printf(`
%sUsage%s: start <name|term>

Just like use & run, start enables you to Interact with a Function by name or search term/index.
If a Function name is not found, it will be treated as a search term. An index from the previous search
results can be selected if desired.

%sex. %s%susage%s:
--  ------

  start setups, torsocks, networks, exploits, wireless, phishers, websites, credits, verses

  start setups
  start <name>

`, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoList() {
    fmt.Printf(`
%sUsage%s: list <name|term>

Show all modules or sub-Functions in a specific face you are. If you are in modules setups,
list modules command, will list all sub-modules available in that Function.

%sex. %s%susage%s:
--  ------

  list modules

%s[*]%s Valid parameters for the "list" command are: modules, functions or all.
%s[*]%s Additional Function-specific parameters are: missing, advanced.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC)
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

%sex. %s%susage%s:
--  ------

  set module exploits
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfo() {
    fmt.Printf(`
%sUsage%s: info <Function name>

%sex. %s%susage%s:
--  ------

  info setups -> or info [ 1. setups   2. torsocks 3. networks 4. exploits 5. wireless ]
                         [ 5. crackers 6. phishers 7. websites 8. credits  9. verse.   ] -> Integer input support. keys are. (1 2 3 4 5 6 7 8 9 0 and 99)

%sDescription%s:
-----------

  Queries the supplied Function or modules for information. If no Function is given, show info for the currently active Function.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTips() {
    fmt.Printf(`
%sUsage%s: getips <Function name>

%sex. %s%susage%s:
--  ------

  info setups, torsocks, networks, exploits, wireless, phishers, webs, verses.

%sDescription%s:
-----------
  Gives you special information about a Function and how to successfully use it in.
real life scenario.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoFeatures() {
    fmt.Printf(`
%sUsage%s: features

%sex. %s%susage%s:
--  ------
  features -> or show features

%sDescription%s:
-----------

Prints feature plans for africana. Ex. modules to be added, what to be corrected.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  %s1. %s    setups%s: Install, Update, Repair or Uninstall africana-framework.
  %s2. %s  torsocks%s: Configure the system for strictly tight anonimity.(fix dns leak, changemac in auto boot)
  %s3. %s  networks%s: Start internal network attacks. From sniffing, injecting packets to hooking hosts.
  %s4. %s  exploits%s: Generate undetectebal R.A.Ts and (Launch next gen c2s to attack all systems. Evasions also included)
  %s5. %s  wireless%s: From wifies, bluetooths, cantools and other wireless networks attack vectors.
  %s6. %s  crackers%s: Crack (NTLMS, HASHES, PCAPS) & bruteforce (SSH, FTP, SMB, RPC etc.)
  %s7. %s  phishers%s: Perform almost all kinds of Phishing attacks. With OPT Bypass.
  %s8. %s  websites%s: Launch Web Penetration testing engines with full free bugbounty automation kit.
  %s9. %s   credits%s: Show who developes and mentains africana-framework with (third party tools developers)
  %s99.%s    verses%s: Launch Bible verses in an uniform way manner as used in the framework.

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
  module         none             yes       Select a module or function to interact with.

  %ssupported modules%s:
  --------- -------

  1. setups   2. torsocks 3. networks 4. exploits 5. wireless
  5. crackers 6. phishers 7. websites 8. credits  9. verse.  -> Integer input support. keys are. (1 2 3 4 5 6 7 8 9 0 and 99)

%sex. %s%susage%s:
--  ------
  set module setups  or ->  use 1, use setups, execute 1, execute setups etc. Type 'menu' to get advanced usages.

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
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

  %s1. %s      kali%s: Necessary tools will be installed in in Kali-Linux(Debian distros). (Use this it is stable)
  %s2. %s    ubuntu%s: This module will install africana-framework in Ubuntu-Linux.
  %s3. %s      arch%s: Tools will be installed in any Arch-Linux Distros using Blackarch repo.
  %s4. %s     macos%s: Under development. Install africana on MackingTosh systems.
  %s5. %s   android%s: Install africana-framework in Termux using chroot environment.
  %s6. %s   windows%s: Under development. But can run if tools well installed using commando vm.

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
  %s1. %s   install%s: Installs africana in selected distro.
  %s2. %s    update%s: Get new release of africana-framework from github and install it.
  %s3. %s    repair%s: Repair africana-framework if broken or with issues.
  %s4. %s      auto%s: Auto detect system and do the necessary.
  %s5. %s uninstall%s: Completely uninstall africana-framework from system.

%sex. %s%susage%s:
--  ------
  set distro kali
  set function install
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC    , bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
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

func HelpInfoKali() {
    fmt.Printf(`
       %sName%s: kali
   %sFunction%s: src/core/setups_kali.fn
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
   -> 0   Kali-Linux

%sDescription%s:
-----------
  It is a module to install africana-framework in kali-linux a stable debian based distro that has a wide comunity support to avoid package breaks and missing dependencies use kali for africana.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoArch() {
    fmt.Printf(`
       %sName%s: arch
   %sFunction%s: src/core/setups_arch.fn
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
   -> 0   Arch-Linux

%sDescription%s:
-----------
  It is a module to install africana-framework in arch based distros. Arch is well established and all tools could be installed with blackman an intergration of black-arch in any arch-linux distro. No errors reported. africana can run well in arch-linux distros.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoUbuntu() {
    fmt.Printf(`
       %sName%s: ubuntu
   %sFunction%s: src/core/setups_ubuntu.fn
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
   -> 0   Ubuntu-Linux

%sDescription%s:
-----------
  It is a module to install africana-framework ubuntu which is a good distor but has alot of problems while installing kali-linux packages. To avoid issues like dependencies problems, Pleas use docker image or install kali-linux in Ubuntu docker then install africana.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoMacos() {
    fmt.Printf(`
       %sName%s: macos
   %sFunction%s: src/core/setups_macos.fn
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
   -> 0   Macintosh

%sDescription%s:
-----------
  It is a module to install africana-framework Macos which is a good distor but has alot of problems while installing kali-linux packages. To avoid issues like dependencies problems, Pleas use docker image or install kali-linux in Macos docker then install africana.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoWindows() {
    fmt.Printf(`
       %sName%s: windows
   %sFunction%s: src/core/setups_windows.fn
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
   -> 0   Windows

%sDescription%s:
-----------
  It is a module to install africana-framework Windows which is a good distor but has alot of problems while installing kali-linux packages. To avoid issues like dependencies problems, Pleas use docker image or install kali-linux in Macos docker then install africana.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoAndroid() {
    fmt.Printf(`
       %sName%s: android
   %sFunction%s: src/core/setups_android.fn
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
   -> 0   Android

%sDescription%s:
-----------
  It is a Function to install africana-framework in android devices. Kali-linux will be installed in termux then kali-linux in chroot environment that will set all dependencies for africana-framework.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoUpdate() {
    fmt.Printf(`
       %sName%s: update
   %sFunction%s: src/core/setups_update.fn
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
   -> 0   All

%sDescription%s:
-----------
  It is a Function to update and mentain africana-framework.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoUninstaller() {
    fmt.Printf(`
       %sName%s: uninstall
   %sFunction%s: src/core/setups_uninstall.fn
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
   -> 0   All

%sDescription%s:
-----------
  It is a Function to Uninstall africana completelty from your system with all it's dependencies. Incase of a bug, email me at rojahsmontari@gmail.com

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoAuto() {
    fmt.Printf(`
       %sName%s: auto
   %sFunction%s: src/core/setups/auto_manage.fn
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
   -> 0   All

%sDescription%s:
-----------
  It is a Function to auto select distro and install africana with all it's dependencies.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoRepair() {
    fmt.Printf(`
       %sName%s: repair
   %sFunction%s: src/core/setups/repair.fn
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
   -> 0   All

%sDescription%s:
-----------
  It is a function repairs africana incase it is broken.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoClearLogs() {
    fmt.Printf(`
       %sName%s: clearlogs
   %sFunction%s: src/core/utils/clear_logs.fn
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
   -> 0   All

%sDescription%s:
-----------
  This module will clear all your logs that has been recorded from the last time you cleaned the log folder..

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  function       none             yes       The function to execute. ex.      -> (setups, vanish, exitnode, status, ipaddress, restore, reload, chains, stop)
  run            none             yes       To execute the function. Alias to -> (start, execute, exec, launch).

%sSupported Distros%s:
-----------------

   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  ------
  set function vanish
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func HelpInfoTorsocksSetups() {
    fmt.Printf(`
       %sName%s: setups
   %sFunction%s: src/securities/setups.fn
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
  set function setups
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will install dnsmasq, squid, privoxy and tor. It will (also set configs) so that all your local traffick will go through
  privoxy -> squid -> then tor network. It is done with great care and integrity for super securities.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTorsocksVanish() {
    fmt.Printf(`
       %sName%s: vanish
   %sFunction%s: src/securities/vanish.fn
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
  set function vanish
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will start services like changemacc to change maccadress in a random way then start dnsmasq, squid, privoxy and tor.
  It will (also set configs) so that all your local traffick will go through privoxy > squid > then tor network.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTorsocksStatus() {
    fmt.Printf(`
       %sName%s: status
   %sFunction%s: src/securities/status.fn
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
  set function status
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will query the system to see if macchanger, dnsmasq, squid, privoxy and tor are working correctly and if all traffic that 
  goes through privoxy > squid > then tor network. It is done with great care and integrity for super securities.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTorsocksTorIp() {
    fmt.Printf(`
       %sName%s: torip
   %sFunction%s: src/securities/torip.fn
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
  set function torip
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will check for your external IP. It querries tor website for your gateway IP. If your system's proxy is correctly
  configured, then you will get a congratulation mesage from tor website.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTorsocksChains() {
    fmt.Printf(`
       %sName%s: chains
   %sFunction%s: src/securities/chains.fn
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
  set function chains
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will querry  /var/log/privoxy/log to follow all logs living your system through squid, privoxy to tor.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTorsocksReload() {
    fmt.Printf(`
       %sName%s: reload
   %sFunction%s: src/securities/reload.fn
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
  set function reload
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will querry  /var/log/privoxy/log to follow all logs living your system through squid, privoxy to tor.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}


func HelpInfoTorsocksExitnode() {
    fmt.Printf(`
       %sName%s: exitnode
   %sFunction%s: src/securities/exitnode.fn
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
  set function exitnode
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will shufle the exit nodes to new ones. If you see your nrtwork is slow, This module can help to find a fast one.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTorsocksRestore() {
    fmt.Printf(`
       %sName%s: restore
   %sFunction%s: src/securities/restore.fn
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
  set function restore
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will restore your Iptables to default. If the Function was killed instantly and IPTABLES were not set as intended, This module
  will help you fix the lack off internet connection.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoTorsocksStop() {
    fmt.Printf(`
       %sName%s: stop
   %sFunction%s: src/securities/stop.fn
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
  set function stop
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will restore your Iptables to default. If the Function was killed instantly and IPTABLES were not set as intended, This module
  will help you fix the lack off internet connection.
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func ListInternalFunctions() {
    fmt.Printf(`
%setworks Functions%s:
-------- ---------

  # %sName       Description%s
  - ----       -----------
  %s1. %sdiscover %s: Discover all devices connected to the network. Lets you locate targets.
  %s2. %sportscan %s: Get open ports on the target you have set.
  %s3. %svulnscan %s: Perform vulnerbility scan on open ports of the target you have set.
  %s4. %senumscan %s: Recon for S.M.B deep information on the target set.
  %s5. %ssmbexplo %s: Launch known vulnerbility exploits on the target's S.M.B services.
  %s6. %spsniffer %s: Sniff all Packets from connected devices to the router(Perform M.I.T.M).
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
   %sFunction%s: src/networks
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
%sModule options %s(src/networks/networks_pentest.fn):

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
  set function discover
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC,       bcolors.BOLD, bcolors.ENDC, bcolors.YELLOW, Lhost, bcolors.ENDC,           bcolors.BOLD, bcolors.ENDC, bcolors.YELLOW, Gateway, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
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
  set function discover
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will scan for all connected devices in the network given using bettercap then arrange the targets
  in a table for you to select one to attack further.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  set function portscan
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will scan all open ports of the target to reveal open ports.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  set function vulnscan
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Perform vulnerbility scan on open ports of the target you have set.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  set function enumscan
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Digg for S.M.B deep information on the target set.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  set function smbexplo
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Launch known vulnerbility exploits on the target's S.M.B services.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  set function psniffer
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Sniff all Packets from connected devices to the router(Perform M.I.T.M).

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  set function responder
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Launch reponder asking for your LHOST, Configuring Wpadscript and weponizing it self. Attack supports alot of windows recent version

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func OptionsResponder() {
    fmt.Printf(`
%sModule Options %s(src/networks/responder.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  RHOST          none             yes       Target to attack. Host's ipadress.
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
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

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
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
  set function beefninja
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will Launch a Combination of both beef-xss and bettercap in a unique way to inject hook.js in either one or all targets. All settings are done for you.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func OptionsBeefNinja() {
    fmt.Printf(`
%sModule Options %s(src/networks/beefninja.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  RHOST          none             yes       Target to attack. Host's ipadress.
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
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

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
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
  set function xsshooker
  run                                     alias to -> (start, execute, exec, launch)

%sDescription%s:
-----------

  This module will try to Get you a revers Shell through XSS Injection. Still Working on this Option.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func OptionsXssHoocker() {
    fmt.Printf(`
%sModule Options %s(src/networks/xsshooker.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  RHOST          none             yes       Target to attack. Host's ipadress.
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
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

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  BUILD          africana_malware yes       Output name of the backdoor to be generated. Neede when using(androrat, noisemaker and codebreaker)
  SCRIPT         none             yes       Your powershell script location to be opfsicated. Mostly neede when using (ghos).
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
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

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool derived from villain framework. It supports both tcp, http & https reverse shells. It has inbuild evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool like villein or blackjack that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func BlackJackOptions() {
    fmt.Printf(`
%sModule Options %s(src/exploits/blackjack_c2.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
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

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  SCRIPT         none             yes       Full location of the powershell script to be obfsicated.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  SCRIPT         none             yes       Powershell script to obfsicate inorder to bypass AVs.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
  PROTOCOL       tcp              yes       Communication protocol to be used between blackjack and client. Supported are, (tcp, http and https).

%sDescription%s:
-----------
  It is a tool that supports both tcp, http & https reverse shells. It has in build evasions and bypasses almost all avs. It is the best for now.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}


func AndroRatOptions() {
    fmt.Printf(`
%sModule Options %s(src/exploits/androrat_c2.fn):

  %sName           Current Setting  Required  Description%s
  ----           ---------------  --------  -----------
  LHOST          ->               yes       %sDefault%s: %s%s%s. Mainly needed when generating backdoors and launching C2s.
  LPORT          9999             yes       Listener port to handle beacons.
  HPORT          3333             yes       The port to handle file smaglers in blacjack function.
  BUILD          africana_malware yes       Output name of the backdoor to be generated. Neede when using(androrat, noisemaker and codebreaker)
  SCRIPT         none             yes       Your powershell script location to be opfsicated. Mostly neede when using (ghos).
  OUTPUT         ->               yes       %sDefault%s: %s%s%s. Location you want your generated backdoor to be placed.
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

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, Lhost, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, OutPutDir, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func ListWirelessFunctions() {
    fmt.Printf(`
%setworks Functions%s:
-------- ---------

  # %sName       Description%s
  - ----       -----------
  %s1. %s    wifite%s: Wifite is a tool to audit WEP or WPA encrypted wireless networks.
  %s2. %s   fluxion%s: Fluxion is a tool to audit WEP or WPA encrypted wireless networks. Only manual option is supported by now.
  %s3. %s bettercap%s: Bettercap is a tool to audit Internal network & wirekless network like, WEP or WPA encrypted wireless networks.
  %s4. %s airgeddon%s: Airgeddon Fluxion is a tool to audit WEP or WPA encrypted wireless networks. Only manual option is supported by now.
  %s5. %s wifipumpk%s: wifipumpkin - Is a Powerful framework for rogue access point attack. This option run automated mode directly.
  %s6. %s Killerpuk%s: This option runs wifipumpkin3 in manual mode where africana sets everything for you.
  %s7. %s    To add%s: Not yet implimented.
  %s8. %s    To add%s: Not yet implimented.
  %s9. %s    To add%s: Not yet implimented.

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
  IFACE          ->               yes       %sDefault%s: %swlan0%s. Mainly needed when generating backdoors and launching C2s.
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

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.ORANGE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
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

%sSuported Distros%s:
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
--  ------

  set function wifite
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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

%sSuported Distros%s:
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
--  ------

  set function fluxion
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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

%sSuported Distros%s:
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
--  ------

  set function bettercap
  set mode auto
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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

%sSuported Distros%s:
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
--  ------

  set function airgeddon
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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

%sSuported Distros%s:
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
--  ------

  set function wifipumpkin
  set mode auto
  set ssid ----
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func HelpInfoCrackers() {
    fmt.Printf(`
       %sName%s: crackers
   %sFunction%s: src/core/crackers.fn
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
  Crackers is a module enriched with creative attacking faces to help redtemers in successfully cracking or brutforce passwords from services online or local encripted files.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
    ListCrackersFunctions()
}

func ListCrackersFunctions() {
    fmt.Printf(`
%sCrackers Functions%s:
-------- ---------

  # %sName       Description%s
  - ----       -----------

  %sOnline Crackers%s:
  ---------------
  %s1. %s         ssh%s: Automated Bruteforcer for SSH using rockyou.txt wordlists.
  %s2. %s         ftp%s: Bruteforcer for FTP using rockyou.txt wordlists.
  %s3. %s         smb%s: Hydra Bruteforcer for SMB using rockyou.txt wordlists.
  %s4. %s         rdp%s: Bruteforcer for RDP using rockyou.txt wordlists.
  %s5. %s        ldap%s: Hydra Bruteforcer for LDAP using rockyou.txt wordlists.
  %s6. %s        smtp%s: Bruteforcer for SMTP using rockyou.txt wordlists.
  %s7. %s      telnet%s: Bruteforcer for TELNET using rockyou.txt wordlists.
  %s8. %s  http/https%s: Hydra Bruteforcer for HTTP/ HTTPS using rockyou.txt wordlists. Forms needed.
  %s9. %sall services%s: The Automatic Bruteforce Tool for all opened services. Works nice and automatic.

  %sOffline Crackers%s:
  ------- --------
  %s1. %s        pcap%s: Crack captured .pcap files. Full location is needed.
  %s2. %s        ntlm%s: Crack ntlm file using default wordlists.
  %s3. %s        hash%s: Auto identify hash and start bruteforcing for passwords.

%sex. %s%susage%s:
--  ------

  set function ssh
  set mode online
  set wordlist -> (Give full path)
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}


func CrackersOptions() {
    fmt.Printf(`
%sModule Options %s(src/crackers/passwords_pentest.fn):

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
-----------------

   Id  Name
   --  ----
   0   All Distros

%sex. %s%susage%s:
--  ------
  set function ssh
  set mode online
  set wordlist -> (Give full path)
  run

View the full module info with the %s'info'%s, or %s'info -d'%s command.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC, bcolors.DARKGREEN, bcolors.ENDC)
}

func ListPhishersFunctions() {
    fmt.Printf(`
%setworks Functions%s:
-------- ---------

  # %sName       Description%s
  - ----       -----------
  %s1. %s   gophish%s: It is a phishing framework with a Web UI https://127.0.0.1:3333. Africana will launch it for you. Default  is: admin, Default password is: kali-gophish.
  %s2. %s  goodginx%s: Goodginx is an advanced phishing framework with insane configurations.Default name evilginx2. Bypasses alot of security features like OTP.
  %s3. %s  zphisher%s: A nice framework with alot of templets. Also bypasses OTP with ngrock support.
  %s4. %sadvphisher%s: Wide range of phishing templets.
  %s5. %sshellphish%s: Supports otp bypass. Wide range of phishing templets.
  %s6. %s darkphish%s: Bypasses OTP with Wide range of phishing templets.
  %s7. %s setoolkit%s: This tool is equiped with alot of social engeneering. Supports cloning of actual websites.
  %s8. %s  blackeye%s: Writen in bash and full of phishing templets. Just check it out.
  %s9. %s       thc%s: This tool creates a templete of your interest imidietly but needs you to start your server and generate a link for phishing.

%sex. %s%susage%s:
--  ------

    set function gophish
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

  set function gophish
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

  set function goodginx
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

  set function zphisher
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

  set function blackeye
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

  set function advphisher
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

  set function darkphish
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

  set function shellphish
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

  set function setoolkit
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

  set function thc
  run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}





func ListWebsitesFunctions() {
    fmt.Printf(`
%sWebsites Functions%s:
-------- ---------

  # %sName       Description%s
  - ----       -----------

  %s1. %s    netmap%s: Network Mapping and portscaning.
  %s2. %s  enumscan%s: Subdomain enumeration.
  %s3. %s  dnsrecon%s: Dns and asn lookup.
  %s4. %s  techscan%s: Web technology detection.
  %s5. %s  asetscan%s: Wayback and asset discovery.
  %s6. %s  fuzzscan%s: Digg for root files from the server.
  %s7. %s  leakscan%s: Secret and leak detection.
  %s8. %s  vulnscan%s: Vulnerability scanning (XSS, SQLi, CSRF, SSRF, IDOR).
  %s9. %s    bounty%s: Bug bounty hunting techniques.

%sex. %s%susage%s:
--  ------

    set module enumscan or -> use enumscan
    set rhosts https://example.com
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ORANGE, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

    set module enumscan or -> use enumscan
    set rhosts https://example.com
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

    set module dnsrecon or -> use dnsrecon
    set rhosts https://example.com
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

    set module portscan or -> use portscan
    set rhosts https://example.com
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

    set module techscan or -> use techscan
    set rhosts https://example.com
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

    set module fuzzscan or -> use fuzzscan
    set rhosts https://example.com
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

    set module leakscan or -> use leakscan
    set rhosts https://example.com
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

    set module vulnscan or -> use vulnscan
    set rhosts https://example.com
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

    set module asetscan or -> use asetscan
    set rhosts https://example.com
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
--  ------

    set module autoscan or -> use autoscan
    set rhosts https://example.com
    run

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BLUE, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
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
  Verses is a module to narate the Bible scriptures one by one. It enables africana developer to acknowledge our LORD GOD JESUS CHRIST for Creating everything including you, me & everything.

`, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC, bcolors.BOLD, bcolors.ENDC)
}

func UpsentTools() {
    fmt.Println(bcolors.YELLOW + "\n[!] " + bcolors.ENDC + "Choice selected not implimented yet!, but comming soon!" + bcolors.ENDC)
}
