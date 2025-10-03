//John 3:16

package utils

import(
    "os"
    "io"
    "net"
    "log"
    "fmt"
    "sort"
    "sync"
    "time"
    "bytes"
    "bufio"
    "regexp"
    "unsafe"
    "syscall"
    "unicode"
    "strconv"
    "os/exec"
    "bcolors"
    "strings"
    "runtime"
    "net/url"
    "net/http"
    "math/big"
    "io/ioutil"
    "subprocess"
    "crypto/rand"
    "crypto/x509"
    "unicode/utf8"
    "crypto/ecdsa"
    "encoding/pem"
    "path/filepath"
    "crypto/elliptic"
    "encoding/base64"
    "crypto/x509/pkix"
)

var (
    cmd *exec.Cmd

    Function   = ""
    Protocol   = "tcp"
    OuterIcon  = "pdf"
    LPort      = "9999"
    HPort      = "3333"
    Obfuscator = "ghost"
    VenvName   = "afr_venv"
    InnerIcon  = "kaspersky"
    Listener   = "blackjack"
    BuildName  = "africana_fud"
    RawBuild   = "africana_raw"

    FakeDns    = "*"
    NeMode     = "all"
    IFace      = "eth0"
    Spoofer    = "ettercap"

    WlanIFace  = "wlan0"
    UserName   = "admin"
    PassWord   = "password"
    Ssid       = "The End Times Ministries."

    WiMode     = "auto"
    ExMode     = "auto"
    SeMode     = "auto"
    PhMode     = "all"
    CrMode     = "online"
    WeMode     = "recon"
    DDosMode   = "gui"

    BeefName   = "beef"
    BeefPass   = "Jesus"
    PhFakeDns  = "Jesus Christ is coming soon!"

    CHAR            = "├──"
    CHAR1           = "└──"
    SPACE_PREFIX    = "   "

    LHost, _ = GetDefaultIP()
    Distro, _ = GetLinuxDistroID()
    GateWay, _ = GetDefaultGatewayIP()
    Scanner = bufio.NewScanner(os.Stdin)

    Date = time.Now().Format("2006-01-02.15.04.05")

    HomeDir = os.Getenv("HOME")
    GoPath = filepath.Join(HomeDir, ".go")

    VenvPath = filepath.Join(BaseDir, VenvName)
    VenvPip = filepath.Join(VenvPath, "bin", "pip")
    VenvPython = filepath.Join(VenvPath, "bin", "python")

    WirelessTools = filepath.Join(ToolsDir, "wireless")
    ExploitsTools = filepath.Join(ToolsDir, "exploits")
    PhishersTools = filepath.Join(ToolsDir, "phishers")
    CrackersTools = filepath.Join(ToolsDir, "crackers")
    WeBountyTools = filepath.Join(ToolsDir, "websites")
    NetworkTools  = filepath.Join(ToolsDir, "networks")

    ReconDir = filepath.Join(OutPutDir, "recondir")
    SetupsLogs = filepath.Join(OutPutDir, "setups")
    SecLogs = filepath.Join(OutPutDir, "securities")
    NetworkLogs = filepath.Join(OutPutDir, "networks")
    ExploitsLogs = filepath.Join(OutPutDir, "exploits")
    PhishersLogs = filepath.Join(OutPutDir, "phishers")
    CrackersLogs = filepath.Join(OutPutDir, "crackers")
    WirelessLogs = filepath.Join(OutPutDir, "wireless")
    WebCrackersLogs = filepath.Join(OutPutDir, "websites")

    IconsDir = filepath.Join(TemplateDir, "icons")
    WinresDir = filepath.Join(TemplateDir, "rcedits")
    GeneratorDir = filepath.Join(ExploitsLogs, "generator")
    WinresInit = filepath.Join(TemplateDir, "rcedits", "winres")
    TemplateDir = filepath.Join(ExploitsTools, "payload_templates")

    OutPutBuild = filepath.Join(GeneratorDir, BuildName)
    WListeners = filepath.Join(ExploitsTools, "windows", "listeners")
    AListeners = filepath.Join(ExploitsTools, "androids", "listeners")
    ObfusCatorsDir = filepath.Join(ExploitsTools, "windows", "obfuscators")

    WordsList = filepath.Join(WordsListDir, "everything.txt")
    ResolversFile = filepath.Join(WordsListDir, "dns_all.txt")

    BaseDir, ToolsDir, CertDir, CertPath, KeyPath, OutPutDir, WordsListDir, RokyPath = subprocess.DirLocations()
    Flag, Shell, Input, Command, AgreementDir, AgreementPath, Proxies, ProxyURL, RHost, Script, Hashes, Pcap string
)

var DefaultValues = map[string]string{
    "build":       BuildName,
    "buildname":   BuildName,
    "ddosmode":    DDosMode,
    "distro":      Distro,
    "fakedns":     FakeDns,
    "func":        Function,
    "function":    Function,
    "functions":   Function,
    "funcs":       Function,
    "gateway":     GateWay,
    "hport":       HPort,
    "iface":       IFace,
    "innericon":   InnerIcon,
    "interface":   IFace,
    "lhost":       LHost,
    "listener":    Listener,
    "listeners":   Listener,
    "lport":       LPort,
    "mode":        NeMode,
    "module":      Function,
    "name":        BeefName,
    "obfuscator":  Obfuscator,
    "outericon":   OuterIcon,
    "output":      OutPutDir,
    "outputlog":   OutPutDir,
    "outputlogs":  OutPutDir,
    "passwd":      BeefPass,
    "password":    PassWord,
    "proxy":       Proxies,
    "proxies":     Proxies,
    "protocol":    Protocol,
    "venvname":    VenvName,
    "recondir":    ReconDir,
    "rhost":       RHost,
    "rhosts":      RHost,
    "script":      Script,
    "spoofer":     Spoofer,
    "ssid":        Ssid,
    "target":      RHost,
    "targets":     RHost,
    "toolsdir":    ToolsDir,
    "wordlist":    WordsList,
}

var AvailableCommands = []string{
    "help", "exit", "version", "status", "start", "stop", "restart", 
    "config", "log", "update", "install", "uninstall", "list",
    "set", "get", "show", "run", "scan", "analyze", "build", "use",
}

var CommandWords = []string{
    "set", "show", "run", "scan", "use", "config", "start", "stop", 
    "help", "version", "status", "install", "uninstall", "update",
    "exit", "restart", "log", "get", "analyze", "build", "list",

    "module", "modules", "function", "functions", "target", "targets",
    "lhost", "lport", "rhost", "rport", "proxy", "proxies", 
    "mode", "protocol", "output", "input", "config", "settings",
    "options", "option", "value", "values", "port", "ports",
    "host", "hosts", "network", "networks", "scan", "scans",
    "attack", "attacks", "exploit", "exploits", "payload", "payloads",
    "ssid", "iface", "interface", "password", "passwd", "wordlist",
    "gateway", "fakedns", "spoofer", "venvname", "toolsdir", "ddosmode",
    "recondir", "listener", "listeners", "outputlog", "innericon", "outericon",
    "buildname", "outputlogs", "obfuscator",
}

var CommandPredictions = map[string][]string{
    "show": {
        "config", "settings", "options", "status", "version", "modules",
        "plugins", "targets", "sessions", "history", "help", "commands",
    },
    "set": {
        "func", "funcs", "module", "ssid", "iface", "mode", "function", 
        "lhost", "lport", "hport", "rhost", "rhosts", "functions", 
        "target", "distro", "targets", "proxy", "script", "name", 
        "interface", "build", "proxies", "passwd", "gateway", "fakedns", 
        "spoofer", "output", "venvname", "toolsdir", "ddosmode", 
        "recondir", "password", "protocol", "listener", "wordlist", 
        "listeners", "outputlog", "innericon", "outericon", "buildname", 
        "outputlogs", "obfuscator", "option", "value",
    },
    "run": {
        "scan", "attack", "exploit", "recon", "bruteforce", "module",
    },
    "scan": {
        "network", "ports", "vulnerabilities", "targets", "local",
    },
    "use": {
        "module", "exploit", "payload", "auxiliary", "post",
    },
    "config": {
        "show", "set", "reset", "save", "load",
    },
    "start": {
        "service", "module", "attack", "scan", "listener",
    },
    "stop": {
        "service", "module", "attack", "scan", "listener",
    },
}

var CommandSuggestions = map[string]string{
    "set":    " <option> <value>",
    "show":   " <category>",
    "run":    " <module/scan>",
    "scan":   " <target>",
    "config": " <action>",
    "help":   " <command>",
    "use":    " <module>",
    "start":  " <service>",
    "stop":   " <service>",

    "se":     "t <option> <value>",
    "sh":     "ow <category>",
    "r":      "un <module/scan>",
    "sc":     "an <target>",
    "co":     "nfig <action>",
    "h":      "elp <command>",
    "u":      "se <module>",
    "st":     "art <service>",
}

var PartialSuggestions = map[string]map[string]string{
    "set": {
        "mo":  "dule <name>",
        "mod": "ule <name>",
        "fu":  "nction <name>",
        "func": "tion <name>",
        "ta":  "rget <host>",
        "tar": "get <host>",
        "lh":  "ost <ip>",
        "lho": "st <ip>",
        "lp":  "ort <number>",
        "lpo": "rt <number>",
        "rh":  "ost <ip>",
        "rho": "st <ip>",
        "pro": "xy <url>",
        "prox": "y <url>",
    },
    "show": {
        "o":   "ptions",
        "op":  "tions",
        "opt": "ions",
        "c":   "onfig",
        "co":  "nfig",
        "conf": "ig",
        "s":   "tatus",
        "st":  "atus",
        "stat": "us",
        "v":   "ersion",
        "ve":  "rsion",
        "ver": "sion",
        "m":   "odules",
        "mo":  "dules",
        "mod": "ules",
    },
    "run": {
        "s":  "can",
        "sc": "an",
        "a":  "ttack",
        "at": "tack",
        "e":  "xploit",
        "ex": "ploit",
        "b":  "ruteforce",
        "br": "uteforce",
    },
    "use": {
        "m":  "odule",
        "mo": "dule",
        "mod": "ule",
        "e":  "xploit",
        "ex": "ploit",
        "p":  "ayload",
        "pa": "yload",
    },
    "scan": {
        "n":  "etwork",
        "ne": "twork",
        "p":  "orts",
        "po": "rts",
        "v":  "ulnerabilities",
        "vu": "lnerabilities",
    },
}

var ValueSuggestions = map[string]string{
    "mode":      "<listen|attack|scan>",
    "protocol":  "<tcp|udp|http>",
    "lhost":     "<your_ip_address>",
    "lport":     "<port_number>",
    "rhost":     "<target_ip>",
    "threads":   "<number>",
    "timeout":   "<seconds>",
    "verbose":   "<true|false>",
    "output":    "<filename>",
}

var CommonValues = []string{"true", "false", "on", "off", "enable", "disable", "yes", "no"}

var CommandCategories = map[string][]string{
    "Configuration":  {"set", "config", "get"},
    "Information":    {"help", "show", "version", "status", "list"},
    "Execution":      {"run", "scan", "start", "stop", "restart", "analyze", "build", "use"},
    "System":         {"exit", "update", "install", "uninstall", "log"},
}

type PromptContext struct {
    Version      string
    Context      string
    Value        string
    ShowBrackets bool
    Color        string
}

type LineEditor struct {
    history          []string
    historyPos       int
    currentLine      string
    cursorPos        int
    prompt           string
    promptLen        int
    historyFile      string
    predictions      []string
    predictionPos    int
    showPredictions  bool
    initialized      bool
    predictionDrawn  bool
    tabPressed       bool
    predictionLines  int
    inPredictionMode bool
    originalLine     string
    originalCursor   int
    // Smart features
    suggestionMode    bool
    currentSuggestion string
    lastKeyTime       time.Time
    keyBuffer         string
    emojiEnabled      bool
    soundEnabled      bool
    tabPressTime      time.Time
}

type InterfaceInfo struct {
    Name        string
    HardwareAddr string
    Flags       net.Flags
    Addresses   []string
}

type Option func(*Spinner)

type Spinner struct {
    mu          sync.Mutex
    active      bool
    stopChan    chan struct{}
    wg          sync.WaitGroup
    spinChars   []string
    baseFormat  string
    baseArgs    []interface{}
    currentText string
    textEffect  func(string, int, int) string
    speed       time.Duration
    step        int
    writer      io.Writer
    taskWg      sync.WaitGroup
    bufWriter   *bufio.Writer
    clearOnStop bool
    noCursor    bool
    isWindows   bool
}

func CheckErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

var SpinnerStyles = map[string][]string{
    "classic":    {"|", "/", "-", "\\"},
    "triangle":   {"◢", "◣", "◤", "◥"},
    "circle":     {"◐", "◓", "◑", "◒"},
    "square":     {"◰", "◳", "◲", "◱"},
    "horizontal": {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
    "dots":       {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
    "fancy":      {"✦", "✧", "★", "✪", "✯", "✵", "✸", "✹"},
    "arrow":      {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
    "moon":       {"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"},
    "bar":        {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"},
    "vertical":   {"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂"},
    "clock":      {"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"},
    "bouncing":   {"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]", "[    ]"},
}

func ClearScreen() {
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }

    cmd.Stdout = os.Stdout
    if err := cmd.Run(); err != nil {
        fmt.Printf("%s[!] %sError clearing screen: %s\n", bcolors.BrightRed, bcolors.Endc, err)
    }
}

func toUpper(r rune) rune {
    if r >= 'a' && r <= 'z' {
        return r - 32
    }
    return r
}

func toLower(r rune) rune {
    if r >= 'A' && r <= 'Z' {
        return r + 32
    }
    return r
}

var TextEffects = map[string]func(string, int, int) string{
    "wave": func(text string, _, step int) string {
        runes := []rune(text)
        for i := range runes {
            if (runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= 'A' && runes[i] <= 'Z') {
                if (i+step)%4 == 0 {
                    if runes[i] >= 'a' && runes[i] <= 'z' {
                        runes[i] = toUpper(runes[i])
                    }
                } else {
                    if runes[i] >= 'A' && runes[i] <= 'Z' {
                        runes[i] = toLower(runes[i])
                    }
                }
            }
        }
        return string(runes)
    },
    "bounce": func(text string, pos, _ int) string {
        runes := []rune(text)
        if pos < len(runes) {
            if (runes[pos] >= 'a' && runes[pos] <= 'z') || (runes[pos] >= 'A' && runes[pos] <= 'Z') {
                if runes[pos] >= 'a' && runes[pos] <= 'z' {
                    runes[pos] = toUpper(runes[pos])
                } else {
                    runes[pos] = toLower(runes[pos])
                }
            }
        }
        return string(runes)
    },
    "randomcase": func(text string, _, step int) string {
        runes := []rune(text)
        for i := range runes {
            if (i+step)%3 == 0 {
                runes[i] = toUpper(runes[i])
            } else {
                runes[i] = toLower(runes[i])
            }
        }
        return string(runes)
    },
    "fadein": func(text string, _, step int) string {
        visibleChars := (step % (len(text) + 3)) - 2
        if visibleChars < 0 {
            return ""
        }
        if visibleChars > len(text) {
            return text
        }
        return text[:visibleChars]
    },
    "typewriter": func(text string, pos, _ int) string {
        if pos >= len(text) {
            return text
        }
        return text[:pos] + "█"
    },
    "neon": func(text string, pos, _ int) string {
        runes := []rune(text)
        if pos < len(runes) {
            return string(runes[:pos]) + "✨" + string(runes[pos:]) + "✨"
        }
        return text
    },
    "plain": func(text string, _, _ int) string {
        return text
    },
}

func New(options ...Option) *Spinner {
    s := &Spinner{
        spinChars:   SpinnerStyles["classic"],
        baseFormat:  "[+] Starting the africana framework console ...",
        baseArgs:    nil,
        textEffect:  TextEffects["plain"],
        speed:       100 * time.Millisecond,
        stopChan:    make(chan struct{}),
        writer:      os.Stdout,
        clearOnStop: true,
        noCursor:    true,
        isWindows:   runtime.GOOS == "windows",
    }
    for _, option := range options {
        option(s)
    }
    s.currentText = fmt.Sprintf(s.baseFormat, s.baseArgs...)
    if _, ok := s.writer.(*bufio.Writer); !ok {
        s.bufWriter = bufio.NewWriter(s.writer)
        s.writer = s.bufWriter
    }
    return s
}

func WithStyle(styleName string) Option {
    return func(s *Spinner) {
        if style, ok := SpinnerStyles[styleName]; ok {
            s.spinChars = style
        }
    }
}

func WithText(text string, a ...interface{}) Option {
    return func(s *Spinner) {
        s.baseFormat = text
        s.baseArgs = a
        s.currentText = fmt.Sprintf(text, a...)
    }
}

func WithEffect(effectName string) Option {
    return func(s *Spinner) {
        if effect, ok := TextEffects[effectName]; ok {
            s.textEffect = effect
        }
    }
}

func WithSpeed(speed time.Duration) Option {
    return func(s *Spinner) {
        s.speed = speed
    }
}

func WithWriter(w io.Writer) Option {
    return func(s *Spinner) {
        s.writer = w
    }
}

func WithClearOnStop(clear bool) Option {
    return func(s *Spinner) {
        s.clearOnStop = clear
    }
}

func WithCursorHidden(hide bool) Option {
    return func(s *Spinner) {
        s.noCursor = hide
    }
}

func (s *Spinner) hideCursor() {
    if s.noCursor {
        if s.isWindows {
            fmt.Fprint(s.writer, "\x1b[?25l")
        } else {
            fmt.Fprint(s.writer, "\033[?25l")
        }
        s.flush()
    }
}

func (s *Spinner) showCursor() {
    if s.noCursor {
        if s.isWindows {
            fmt.Fprint(s.writer, "\x1b[?25h")
        } else {
            fmt.Fprint(s.writer, "\033[?25h")
        }
        s.flush()
    }
}

func (s *Spinner) Start() {
    s.mu.Lock()
    defer s.mu.Unlock()

    if s.active {
        return
    }

    s.active = true
    s.stopChan = make(chan struct{})
    //s.hideCursor()

    s.wg.Add(1)
    go s.spin()
}

func (s *Spinner) spin() {
    defer s.wg.Done()
    //defer s.showCursor()

    charIdx := 0
    letterPos := 0
    ticker := time.NewTicker(s.speed)
    defer ticker.Stop()

    for {
        select {
        case <-s.stopChan:
            if s.clearOnStop {
                s.clearLine()
            } else {
                fmt.Fprint(s.writer, "\r" + s.currentText + "  \r" + s.currentText)
            }
            return
        case <-ticker.C:
            s.mu.Lock()
            s.currentText = fmt.Sprintf(s.baseFormat, s.baseArgs...)

            spinChar := s.spinChars[charIdx%len(s.spinChars)]
            displayText := s.textEffect(s.currentText, letterPos, s.step)
            displayMsg := fmt.Sprintf("\r%s %s", displayText, spinChar)
            s.mu.Unlock()

            s.clearLine()
            fmt.Fprint(s.writer, displayMsg)
            s.flush()

            charIdx++
            s.step++

            if charIdx%1 == 0 {
                letterPos = (letterPos + 1) % utf8.RuneCountInString(s.currentText)
            }
        }
    }
}

func (s *Spinner) StartWithTask(task func()) {
    s.mu.Lock()
    if s.active {
        s.mu.Unlock()
        return
    }

    s.active = true
    s.stopChan = make(chan struct{})
    //s.hideCursor()
    s.mu.Unlock()

    s.wg.Add(1)
    go s.spin()

    s.taskWg.Add(1)
    go func() {
        defer s.taskWg.Done()
        task()
        s.Stop()
    }()
}

func (s *Spinner) Stop() {
    s.mu.Lock()
    defer s.mu.Unlock()

    if !s.active {
        return
    }

    close(s.stopChan)
    s.wg.Wait()
    
    if s.clearOnStop {
        s.clearLine()
    } else {
        fmt.Fprint(s.writer, "\r" + s.currentText + "  \r" + s.currentText)
    }
    
    //s.showCursor()
    s.flush()
    s.active = false
}

func (s *Spinner) clearLine() {
    if s.isWindows {
        width := 80
        fmt.Fprint(s.writer, "\r" + strings.Repeat(" ", width) + "\r")
    } else {
        fmt.Fprint(s.writer, "\033[2K\r")
    }
}

func (s *Spinner) flush() {
    if s.bufWriter != nil {
        s.bufWriter.Flush()
    }
    if f, ok := s.writer.(interface{ Flush() error }); ok {
        f.Flush()
    }
}

func (s *Spinner) UpdateText(format string, a ...interface{}) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.baseFormat = format
    s.baseArgs = a
    s.currentText = fmt.Sprintf(format, a...)
}

func GenerateSelfSignedCert(certPath, keyPath string) {
    priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        fmt.Printf("%s[!] %sError generating key: ", bcolors.BrightRed, bcolors.Endc, err)
        return
    }

    notBefore := time.Now()
    notAfter := notBefore.Add(365 * 24 * time.Hour)
    serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
    if err != nil {
        fmt.Printf("%s[!] %sError generating serial number: ", bcolors.BrightRed, bcolors.Endc, err)
        return
    }

    template := x509.Certificate{
        SerialNumber: serialNumber,
        Subject: pkix.Name{
            Organization: []string{"The End Times Ministries."},
        },
        NotBefore:             notBefore,
        NotAfter:              notAfter,
        KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
        ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
        BasicConstraintsValid: true,
    }

    certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
    if err != nil {
        fmt.Printf("%s[!] %sError creating certificate: ", bcolors.BrightRed, bcolors.Endc, err)
        return
    }

    keyFile, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
        fmt.Printf("%s[!] %sError creating key file: ", bcolors.BrightRed, bcolors.Endc, err)
        return
    }
    defer keyFile.Close()

    privBytes, err := x509.MarshalECPrivateKey(priv)
    if err != nil {
        fmt.Printf("%s[!] %sError marshaling private key: ", bcolors.BrightRed, bcolors.Endc, err)
        return
    }
    pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes})

    certFile, err := os.Create(certPath)
    if err != nil {
        fmt.Printf("%s[!] %sError creating certificate file: ", bcolors.BrightRed, bcolors.Endc, err)
        return
    }
    defer certFile.Close()
    pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certDER})
}

func GetDefaultIP() (string, error) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }
    for _, iface := range interfaces {
        if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
            continue
        }
        addrs, err := iface.Addrs()
        if err != nil {
            return "", err
        }
        for _, addr := range addrs {
            switch v := addr.(type) {
            case *net.IPNet:
                if ip := v.IP.To4(); ip != nil {
                    return ip.String(), nil
                }
            case *net.IPAddr:
                if ip := v.IP.To4(); ip != nil {
                    return ip.String(), nil
                }
            }
        }
    }
    return "", fmt.Errorf("no active network interface found")
}

func IFaces() ([]InterfaceInfo, error) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return nil, fmt.Errorf("error getting interfaces: %w", err)
    }

    var result []InterfaceInfo

    for _, iface := range interfaces {
        info := InterfaceInfo{
            Name:        iface.Name,
            HardwareAddr: iface.HardwareAddr.String(),
            Flags:       iface.Flags,
        }

        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }

        for _, addr := range addrs {
            info.Addresses = append(info.Addresses, addr.String())
        }

        result = append(result, info)
    }

    return result, nil
}

func SystemShell(Command string, args ...interface{}) {
    fmt.Printf("%s[*] %sexec: %s\n\n", bcolors.BrightBlue, bcolors.Endc, Input)
    subprocess.Run(Command)
    return
}

func GetDefaultGatewayIP() (string, error) {
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/C", "route print 0.0.0.0")
    } else {
        cmd = exec.Command("sh", "-c", "ip route show default | awk '{print $3}'")
    }
    output, err := cmd.Output()
    if err != nil {
        return "", err
    }
    gatewayIP := strings.TrimSpace(string(output))
    if gatewayIP == "" {
        return "", fmt.Errorf("default gateway not found")
    }
    return gatewayIP, nil
}

func ValidateIPAddress(ip string) error {
    if net.ParseIP(ip) == nil {
        return fmt.Errorf("invalid IP address format")
    }
    return nil
}

func ValidatePort(port string) error {
    p, err := strconv.Atoi(port)
    if err != nil || p < 1 || p > 65535 {
        return fmt.Errorf("port must be between 1-65535")
    }
    return nil
}

func AskForProxy(proxyStr string) (*url.URL, error) {
    proxyStr = strings.TrimSpace(proxyStr)
    if proxyStr == "" {
        return nil, fmt.Errorf("proxy string is empty")
    }

    proxyURL, err := url.Parse(proxyStr)
    if err != nil {
        return nil, fmt.Errorf("failed to parse proxy URL: %w", err)
    }
    if proxyURL.Scheme == "" || proxyURL.Host == "" {
        return nil, fmt.Errorf("%s%s[!] %sInvalid proxy format (eg. http://localhost:8080, socks5://127.0.0.1:9050)", bcolors.Bold, bcolors.Yellow, bcolors.Endc)
    }

    //validSchemes := map[string]bool{"http": true, "https": true}
    //if !validSchemes[proxyURL.Scheme] {
    //    if strings.HasPrefix(proxyURL.Scheme, "socks") {
    //        return nil, fmt.Errorf("[!] SOCKS proxy support requires external modules (golang.org/x/net/proxy). Please use an HTTP proxy")
    //    }
    //    return nil, fmt.Errorf("[!] Invalid proxy scheme %q (use http or https)", proxyURL.Scheme)
    //}

    return proxyURL, nil
}

func SetProxy(proxyStr string) error {
    proxyURL, err := AskForProxy(proxyStr)
    if err != nil {
        return fmt.Errorf("proxy validation failed: %w", err)
    }
    envVars := map[string]string{
        "HTTP_PROXY":  proxyURL.String(),
        "HTTPS_PROXY": proxyURL.String(),
        "ALL_PROXY":   proxyURL.String(),
        "http_proxy":  proxyURL.String(),
        "https_proxy": proxyURL.String(),
        "all_proxy":   proxyURL.String(),
    }
    for key, value := range envVars {
        if err := os.Setenv(key, value); err != nil {
            return fmt.Errorf("failed to set %s: %w", key, err)
        }
    }

    validHTTPSchemes := map[string]bool{"http": true, "https": true}
    if validHTTPSchemes[proxyURL.Scheme] {
        http.DefaultTransport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
        //fmt.Printf("%s%s[>] %sGo HTTP transport also configured for: %s\n", bcolors.Bold, bcolors.Blue, bcolors.Endc, proxyURL.String())
    } else {
        fmt.Printf("%s%s[>] %sProxy set in environment for external tools: %s\n", bcolors.Bold, bcolors.Yellow, bcolors.Endc, proxyURL.String())
        fmt.Printf("%s%s[!] %sNote: Native Go HTTP calls will not use the %s proxy (requires golang.org/x/net/proxy).\n", bcolors.Bold, bcolors.Red, bcolors.Endc, proxyURL.Scheme)
    }
    return nil
}

func GetProxiedHTTPClient(proxyStr string) (*http.Client, error) {
    proxyURL, err := AskForProxy(proxyStr)
    if err != nil {
        return nil, err
    }

    customTransport := &http.Transport{
        Proxy: http.ProxyURL(proxyURL),
        DialContext: (&net.Dialer{
            Timeout:   30 * time.Second,
            KeepAlive: 30 * time.Second,
        }).DialContext,
        TLSHandshakeTimeout: 10 * time.Second,
        IdleConnTimeout:     30 * time.Second,
    }

    return &http.Client{
        Transport: customTransport,
        Timeout:   60 * time.Second,
    }, nil
}

func replaceStringsInFile(fileName string, replacements map[string]string) error {
    content, err := ioutil.ReadFile(fileName)
    if err != nil {
        return err
    }
    textContent := string(content)
    for old, new := range replacements {
        textContent = strings.Replace(textContent, old, new, -1)
    }
    return ioutil.WriteFile(fileName, []byte(textContent), 0644)
}

func Editors(filesToReplacements map[string]map[string]string) {
    for fileName, replacements := range filesToReplacements {
        err := replaceStringsInFile(fileName, replacements)
        if err != nil {
            fmt.Printf("%s[!] %sError Configuring: %s%v", bcolors.BrightRed, bcolors.Endc, fileName, err)
        } //else {
            //fmt.Printf("%s[*] %sDone configuring: %s%s%s", bcolors.Green, bcolors.Endc, bcolors.BrightBlue, fileName, bcolors.Endc)
        //}
    }
}

func ReadFileLetterByLetter(color, filename string, delay time.Duration) {
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        return
    }

    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("%s[!] %sError opening file: ", bcolors.BrightRed, bcolors.Endc, err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        for _, letter := range line {
            fmt.Print(color + string(letter) + bcolors.Endc)
            time.Sleep(delay)
        }
        fmt.Printf("%s", "\n")
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("%s[!] %sError reading file: ", bcolors.BrightRed, bcolors.Endc, err)
    }
}

func FileExists(filename string) bool {
    _, err := os.Stat(filename)
    return !os.IsNotExist(err)
}

func StripANSI(text string) string {
    ansi := regexp.MustCompile(`\x1B(?:[@-Z\\-_]]|\[[0-?]*[ -/]*[@-~])`)
    return ansi.ReplaceAllString(text, "")
}

func StripBrackets(text string) string {
    return strings.Replace(strings.Replace(text, "[", "", -1), "]", "", -1)
}

func GetAgreementPath() string {
    agreementDir := filepath.Join("/root/.afr" + subprocess.Version, "agreements")
    if !subprocess.CheckRoot(){
        homeDir := subprocess.GetHomeDir()
        if homeDir == "" {
            return ""
        }
        agreementDir = filepath.Join(homeDir, ".afr" + subprocess.Version, "agreements")
    }
    return filepath.Join(agreementDir, "covenant.txt")
}

func Agreements(filePath string) {
    dirPath := filepath.Dir(filePath)

    if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
        fmt.Printf("[!] Error creating directory: %v\n", err)
        return
    }

    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        if err := ioutil.WriteFile(filePath, []byte("user accepted to the agreement."), os.ModePerm); err != nil {
            fmt.Printf("[!] Error writing file: %v\n", err)
            return
        }
    }
}

func Sealing() {
    filePath := GetAgreementPath()
    if filePath == "" {
        fmt.Printf("%s%s[!] %sError: Could not determine agreement file path.", bcolors.Bold, bcolors.Red, bcolors.Endc)
        return
    }
    Agreements(filePath)
}

func Copy(src, dst string) error {
    if src == "" || dst == "" || src == dst {
        return fmt.Errorf("invalid path: src=%q, dst=%q", src, dst)
    }

    srcInfo, err := os.Stat(src)
    if err != nil {
        return fmt.Errorf("stat source failed: %w", err)
    }

    if srcInfo.IsDir() {
        return copyDir(src, dst, srcInfo)
    }
    return copyFile(src, dst, srcInfo)
}

func copyDir(src, dst string, srcInfo os.FileInfo) error {
    if empty, err := isDirEmpty(dst); err == nil && !empty {
        return fmt.Errorf("destination directory exists and is not empty")
    }

    if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
        return fmt.Errorf("create directory failed: %w", err)
    }

    entries, err := os.ReadDir(src)
    if err != nil {
        return err
    }

    for _, entry := range entries {
        srcPath := filepath.Join(src, entry.Name())
        dstPath := filepath.Join(dst, entry.Name())
        
        if err := Copy(srcPath, dstPath); err != nil {
            return err
        }
    }
    return nil
}

func copyFile(src, dst string, srcInfo os.FileInfo) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    if _, err := io.Copy(dstFile, srcFile); err != nil {
        return err
    }

    return os.Chmod(dst, srcInfo.Mode())
}

func isDirEmpty(dir string) (bool, error) {
    f, err := os.Open(dir)
    if err != nil {
        if os.IsNotExist(err) {
            return true, nil
        }
        return false, err
    }

    defer f.Close()

    _, err = f.Readdirnames(1)
    return err == io.EOF, nil
}

func Move(src, dst string) error {
    err := os.Rename(src, dst)
    if err == nil {
        return nil
    }

    if err = Copy(src, dst); err != nil {
        return err
    }
    return os.Remove(src)
}

func EncodeFileToPowerShellEncodedCommand(filePath string) (string, error) {
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return "", err
    }

    cleaned := string(bytes.Trim(content, "\x00"))

    buf := new(bytes.Buffer)
    for _, r := range cleaned {
        buf.WriteByte(byte(r))
        buf.WriteByte(byte(r >> 8))
    }

    encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
    return encoded, nil
}

func GetLinuxDistroID() (string, error) {
    if _, err := os.Stat("/etc/arch-release"); err == nil {
        return "arch", nil
    }

    file, err := os.Open("/etc/os-release")
    if err != nil {
        return "", err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "ID=") {
            distroID := strings.Trim(strings.TrimPrefix(line, "ID="), `"`)
            if distroID == "manjaro" || distroID == "endeavouros" || distroID == "garuda" {
                return "arch", nil
            }
            return distroID, nil
        }
    }

    return "", fmt.Errorf("could not determine Linux distro")
}

func IsArchLinux() bool {
    distroID, err := GetLinuxDistroID()
    if err != nil {
        if _, err := exec.LookPath("pacman"); err == nil {
            if _, err := os.Stat("/etc/debian_version"); os.IsNotExist(err) {
                return true
            }
        }
        return false
    }

    archDistros := map[string]bool{
        "arch":        true,
        "manjaro":     true,
        "endeavouros": true,
        "garuda":      true,
        "artix":       true,
        "arcolinux":   true,
    }

    return archDistros[distroID]
}

func DetectAndroid() bool {
    if IsNethunter() {
        return false
    }

    if DetectTermux() && !IsNethunter() {
        return true
    }

    if _, err := os.Stat("/system/build.prop"); err == nil {
        return true
    }

    if os.Getenv("ANDROID_ROOT") != "" || os.Getenv("ANDROID_DATA") != "" {
        return true
    }

    if _, err := os.Stat("/data/data/com.termux/files/home"); err == nil {
        return true
    }

    if runtime.GOOS == "linux" {
        if _, err := os.Stat("/proc/version"); err == nil {
            if data, err := os.ReadFile("/proc/version"); err == nil {
                if strings.Contains(strings.ToLower(string(data)), "android") {
                    return true
                }
            }
        }
    }

    return false
}

func DetectTermux() bool {
    if os.Getenv("TERMUX_VERSION") != "" {
        return true
    }

    if _, err := os.Stat("/data/data/com.termux/files/usr"); err == nil {
        return true
    }

    if strings.Contains(os.Getenv("PREFIX"), "com.termux") {
        return true
    }
    
    return false
}

func IsNethunter() bool {
    if _, err := exec.LookPath("nethunter"); err == nil {
        return true
    }

    if _, err := exec.LookPath("nh"); err == nil {
        return true
    }

    if _, err := os.Stat("/data/local/nhsystem"); err == nil {
        return true
    }

    if _, err := os.Stat("/etc/nethunter-release"); err == nil {
        return true
    }
    
    if _, err := os.Stat("/etc/kali-release"); err == nil {
        return true
    }
    
    if _, err := os.Stat("/etc/debian_chroot"); err == nil {
        return true
    }

    if os.Getenv("DEBIAN_CHROOT") != "" {
        return true
    }

    return false
}

func AppendToShellProfile(profilePath, content string) error {
    if _, err := os.Stat(profilePath); os.IsNotExist(err) {
        if err := os.WriteFile(profilePath, []byte(content), 0644); err != nil {
            return err
        }
        return nil
    }

    f, err := os.OpenFile(profilePath, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.WriteString(content)
    return err
}

func DisplayPrompt(version string, args ...interface{}) (string, error) {
    var prompt string

    if len(args) == 0 {
        prompt = fmt.Sprintf("%s%s%safr%s%s%s > %s",
            bcolors.Endc,
            bcolors.Underline,
            bcolors.Bold,
            version,
            bcolors.Endc,
            bcolors.BrightGreen,
            bcolors.Endc)
    } else {
        context := ""
        value := Function
        showBrackets := true
        color := "brightred"
        
        if len(args) >= 1 {
            if ctx, ok := args[0].(string); ok {
                context = ctx
            }
        }

        if len(args) >= 2 {
            if val, ok := args[1].(string); ok {
                value = val
            }
        }

        if len(args) >= 3 {
            if brackets, ok := args[2].(bool); ok {
                showBrackets = brackets
            }
        }

        if showBrackets {
            prompt = fmt.Sprintf("%s%s%safr%s%s %s(%s%ssrc/pentest_%s.fn%s)%s > %s",
                bcolors.Endc,
                bcolors.Underline,
                bcolors.Bold,
                version,
                bcolors.Endc,
                context,
                bcolors.Bold,
                getContextColor(color),
                value,
                bcolors.Endc,
                bcolors.BrightGreen,
                bcolors.Endc)
        } else {
            prompt = fmt.Sprintf("%s%s%safr%s%s %s%s%ssrc/pentest_%s.fn%s%s > %s",
                bcolors.Endc,
                bcolors.Underline,
                bcolors.Bold,
                version,
                bcolors.Endc,
                context,
                bcolors.Bold,
                getContextColor(color),
                value,
                bcolors.Endc,
                bcolors.BrightGreen,
                bcolors.Endc)
        }
    }

    editor := NewLineEditor(prompt)
    return editor.ReadLine()
}

func getContextColor(color string) string {
    colorMap := map[string]string{
        "red":        bcolors.BrightRed,
        "blue":       bcolors.Blue,
        "green":      bcolors.Green,
        "yellow":     bcolors.Yellow,
        "magenta":    bcolors.Magenta,
        "cyan":       bcolors.Cyan,
        "brightred":  bcolors.BrightRed,
        "brightblue": bcolors.BrightBlue,
        "default":    bcolors.BrightRed,
    }
    
    if selected, exists := colorMap[color]; exists {
        return selected
    }
    return bcolors.BrightRed
}

func SetAvailableCommands(commands []string) {
    AvailableCommands = commands
}

func SetCommandPredictions(predictions map[string][]string) {
    CommandPredictions = predictions
}

func SetCommandWords(words []string) {
    CommandWords = words
}

func NewLineEditor(prompt string) *LineEditor {
    visualLength := calculateVisualLength(prompt)
    historyFile := subprocess.HistoryFile
    
    editor := &LineEditor{
        history:          []string{},
        historyPos:       0,
        currentLine:      "",
        cursorPos:        0,
        prompt:           prompt,
        promptLen:        visualLength,
        historyFile:      historyFile,
        predictions:      []string{},
        predictionPos:    -1,
        showPredictions:  false,
        initialized:      false,
        predictionDrawn:  false,
        tabPressed:       false,
        predictionLines:  0,
        inPredictionMode: false,
        originalLine:     "",
        originalCursor:   0,
        suggestionMode:   false,
        currentSuggestion: "",
        lastKeyTime:      time.Now(),
        keyBuffer:        "",
        emojiEnabled:     false,
        soundEnabled:     false,
        tabPressTime:     time.Time{},
    }

    editor.loadHistory()
    editor.initialized = true
    return editor
}

func calculateVisualLength(prompt string) int {
    length := 0
    inEscape := false
    for _, r := range prompt {
        if r == '\033' {
            inEscape = true
            continue
        }
        if inEscape {
            if r == 'm' {
                inEscape = false
            }
            continue
        }
        length++
    }
    return length
}

func Levenshtein(a, b string) int {
    a, b = strings.ToLower(a), strings.ToLower(b)
    ar, br := []rune(a), []rune(b)
    la, lb := len(ar), len(br)

    if la == 0 {
        return lb
    }
    if lb == 0 {
        return la
    }

    prevRow := make([]int, lb+1)
    currRow := make([]int, lb+1)

    for j := 0; j <= lb; j++ {
        prevRow[j] = j
    }

    for i := 1; i <= la; i++ {
        currRow[0] = i
        for j := 1; j <= lb; j++ {
            cost := 0
            if ar[i-1] != br[j-1] {
                cost = 1
            }
            currRow[j] = min(
                currRow[j-1] + 1,
                prevRow[j] + 1,
                prevRow[j-1] + cost,
            )
        }
        prevRow, currRow = currRow, prevRow
    }
    return prevRow[lb]
}

func min(nums ...int) int {
    m := nums[0]
    for _, n := range nums {
        if n < m {
            m = n
        }
    }
    return m
}

func (le *LineEditor) getTerminalWidth() int {

    type winsize struct {
        Row    uint16
        Col    uint16
        Xpixel uint16
        Ypixel uint16
    }

    ws := &winsize{}
    retCode, _, _ := syscall.Syscall(
        syscall.SYS_IOCTL,
        uintptr(syscall.Stdout),
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(ws)),
    )

    if int(retCode) == -1 {
        return 80
    }
    
    width := int(ws.Col)
    if width <= 0 {
        return 80
    }
    return width
}

func (le *LineEditor) generatePredictions() {
    le.predictions = []string{}
    le.predictionPos = -1

    currentText := strings.TrimSpace(le.currentLine)
    if currentText == "" {

        le.predictions = AvailableCommands
    } else {
        parts := strings.Fields(currentText)
        if len(parts) == 0 {
            return
        }

        if len(parts) == 1 {

            cmd := strings.ToLower(parts[0])
            if subcommands, exists := CommandPredictions[cmd]; exists {
                le.predictions = subcommands
            } else {

                for _, availableCmd := range AvailableCommands {
                    lowerCmd := strings.ToLower(availableCmd)
                    if le.isPartialMatch(cmd, lowerCmd) {
                        le.predictions = append(le.predictions, availableCmd)
                    }
                }
            }
        } else if len(parts) >= 2 {

            cmd := strings.ToLower(parts[0])
            currentSub := strings.ToLower(parts[len(parts)-1])
            
            if subcommands, exists := CommandPredictions[cmd]; exists {
                for _, subcmd := range subcommands {
                    lowerSub := strings.ToLower(subcmd)
                    if le.isPartialMatch(currentSub, lowerSub) {
                        le.predictions = append(le.predictions, subcmd)
                    }
                }
            }

            if len(le.predictions) == 0 {
                for _, val := range CommonValues {
                    if le.isPartialMatch(currentSub, val) {
                        le.predictions = append(le.predictions, val)
                    }
                }
            }
        }
    }

    uniquePredictions := make(map[string]bool)
    var finalPredictions []string
    for _, p := range le.predictions {
        if !uniquePredictions[p] {
            uniquePredictions[p] = true
            finalPredictions = append(finalPredictions, p)
        }
    }
    sort.Strings(finalPredictions)
    le.predictions = finalPredictions

    if len(le.predictions) > 0 {
        le.showPredictions = true
        le.predictionPos = 0
    }
}

func (le *LineEditor) showInlinePredictions() {
    if len(le.predictions) == 0 || !le.showPredictions {
        return
    }

    if !le.inPredictionMode {
        le.originalLine = le.currentLine
        le.originalCursor = le.cursorPos
        le.inPredictionMode = true
    }

    le.redrawWithPredictions()
}

func (le *LineEditor) redrawWithPredictions() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)

    le.showSmartSuggestion()

    if le.showPredictions && len(le.predictions) > 0 {
        terminalWidth := le.getTerminalWidth()
        currentLineLength := le.promptLen + len(le.currentLine)
        availableWidth := terminalWidth - currentLineLength - 10

        predictionText := " " + bcolors.Yellow + "[" + bcolors.Endc

        startIdx := 0
        maxPredictions := len(le.predictions)

        if len(le.predictions) > 5 {
            visibleCount := 5
            if availableWidth < 50 {
                visibleCount = 3
            }

            if le.predictionPos >= visibleCount {
                startIdx = le.predictionPos - visibleCount + 1
                if startIdx + visibleCount > len(le.predictions) {
                    startIdx = len(le.predictions) - visibleCount
                }
            }

            endIdx := startIdx + visibleCount
            if endIdx > len(le.predictions) {
                endIdx = len(le.predictions)
            }

            maxPredictions = endIdx - startIdx

            if startIdx > 0 {
                predictionText += bcolors.Magenta + "←" + bcolors.Endc
            }
        }

        for i := 0; i < maxPredictions; i++ {
            actualIdx := startIdx + i
            if actualIdx >= len(le.predictions) {
                break
            }
            
            pred := le.predictions[actualIdx]
            if actualIdx == le.predictionPos {
                predictionText += bcolors.Green + "›" + pred + bcolors.Endc
            } else {
                predictionText += bcolors.Cyan + pred + bcolors.Endc
            }

            if i < maxPredictions-1 {
                predictionText += bcolors.Yellow + "|" + bcolors.Endc
            }
        }

        if startIdx + maxPredictions < len(le.predictions) {
            remaining := len(le.predictions) - (startIdx + maxPredictions)
            predictionText += bcolors.Yellow + "|" + bcolors.Endc
            predictionText += bcolors.Magenta + "+" + fmt.Sprintf("%d", remaining) + bcolors.Endc
        }
        
        predictionText += bcolors.Yellow + "]" + bcolors.Endc
        fmt.Print(predictionText)
    }

    fmt.Printf("\r\033[%dC", le.promptLen+le.cursorPos)
}

func (le *LineEditor) applyPrediction() {
    if le.predictionPos >= 0 && le.predictionPos < len(le.predictions) {
        selected := le.predictions[le.predictionPos]
        
        parts := strings.Fields(le.originalLine)
        
        if len(parts) == 0 {
            le.currentLine = selected + " "
            le.cursorPos = len(le.currentLine)
        } else if len(parts) == 1 {
            currentWord := parts[0]
            if le.isPartialMatch(currentWord, selected) {
                le.currentLine = selected
                if le.isCommandThatTakesArgs(selected) {
                    le.currentLine += " "
                }
            } else {
                le.currentLine = le.originalLine + " " + selected
                if le.isCommandThatTakesArgs(selected) {
                    le.currentLine += " "
                }
            }
            le.cursorPos = len(le.currentLine)
        } else {
            lastWord := parts[len(parts)-1]
            if le.isPartialMatch(lastWord, selected) {
                newParts := parts[:len(parts)-1]
                newParts = append(newParts, selected)
                le.currentLine = strings.Join(newParts, " ")
                if le.isCommandThatTakesArgs(selected) {
                    le.currentLine += " "
                }
            } else {
                le.currentLine = le.originalLine + " " + selected
                if le.isCommandThatTakesArgs(selected) {
                    le.currentLine += " "
                }
            }
            le.cursorPos = len(le.currentLine)
        }
    }
    
    le.exitPredictionMode()
}

func (le *LineEditor) isPartialMatch(partial, full string) bool {
    partialLower := strings.ToLower(partial)
    fullLower := strings.ToLower(full)

    if strings.HasPrefix(fullLower, partialLower) {
        return true
    }

    if Levenshtein(partialLower, fullLower) <= 2 {
        return true
    }

    abbreviations := map[string]string{
        "mo": "module", "mod": "module", "conf": "config", 
        "opt": "options", "set": "settings", "stat": "status",
        "ver": "version", "inst": "install", "uninst": "uninstall",
    }
    
    if expanded, exists := abbreviations[partialLower]; exists {
        return expanded == fullLower
    }
    
    return false
}

func (le *LineEditor) isCommandThatTakesArgs(cmd string) bool {
    commandsWithArgs := map[string]bool{
        "set": true, "show": true, "run": true, "scan": true, 
        "use": true, "config": true, "start": true, "stop": true,
    }

    return commandsWithArgs[strings.ToLower(cmd)]
}

func (le *LineEditor) exitPredictionMode() {
    le.showPredictions = false
    le.inPredictionMode = false
    le.predictions = []string{}
    le.predictionPos = -1
    le.predictionDrawn = false
    le.redraw()
}

func (le *LineEditor) nextPrediction() {
    if len(le.predictions) == 0 {
        return
    }
    
    le.predictionPos = (le.predictionPos + 1) % len(le.predictions)
    le.showInlinePredictions()
}

func (le *LineEditor) prevPrediction() {
    if len(le.predictions) == 0 {
        return
    }
    
    le.predictionPos--
    if le.predictionPos < 0 {
        le.predictionPos = len(le.predictions) - 1
    }
    le.showInlinePredictions()
}

func (le *LineEditor) firstPrediction() {
    if len(le.predictions) == 0 {
        return
    }
    
    le.predictionPos = 0
    le.showInlinePredictions()
}

func (le *LineEditor) lastPrediction() {
    if len(le.predictions) == 0 {
        return
    }
    
    le.predictionPos = len(le.predictions) - 1
    le.showInlinePredictions()
}

func (le *LineEditor) getSmartSuggestion() string {
    text := strings.TrimSpace(le.currentLine)
    if text == "" {
        return ""
    }

    parts := strings.Fields(text)
    if len(parts) == 0 {
        return ""
    }

    lastWord := parts[len(parts)-1]
    lowerLastWord := strings.ToLower(lastWord)

    if len(parts) == 1 {
        return le.getCommandSuggestion(lowerLastWord)
    } else if len(parts) >= 2 {
        return le.getContextSuggestion(parts)
    }

    return ""
}

func (le *LineEditor) getCommandSuggestion(cmd string) string {
    if suggestion, exists := CommandSuggestions[cmd]; exists {
        return suggestion
    }

    for fullCmd, suggestion := range CommandSuggestions {
        if strings.HasPrefix(fullCmd, cmd) && cmd != fullCmd {
            return fullCmd[len(cmd):] + suggestion
        }
    }

    return ""
}

func (le *LineEditor) getContextSuggestion(parts []string) string {
    cmd := strings.ToLower(parts[0])
    lastWord := strings.ToLower(parts[len(parts)-1])

    if le.looksLikeUserIsTypingValue(lastWord) {
        return ""
    }

    switch cmd {
    case "set":
        return le.getSetSuggestion(parts, lastWord)
    case "show":
        return le.getShowSuggestion(parts, lastWord)
    case "run":
        return le.getRunSuggestion(parts, lastWord)
    case "use":
        return le.getUseSuggestion(parts, lastWord)
    case "scan":
        return le.getScanSuggestion(parts, lastWord)
    }

    return ""
}

func (le *LineEditor) getSetSuggestion(parts []string, lastWord string) string {
    if len(parts) == 2 {
        if suggestions, exists := PartialSuggestions["set"]; exists {
            if suggestion, exists := suggestions[lastWord]; exists {
                return suggestion
            }
        }

        return " <value>"
    } else if len(parts) == 3 {
        option := strings.ToLower(parts[1])
        if suggestion, exists := ValueSuggestions[option]; exists {
            if !le.looksLikeUserIsTypingValue(lastWord) {
                return suggestion
            }
        }

        return " <value>"
    }

    return ""
}

func (le *LineEditor) getShowSuggestion(parts []string, lastWord string) string {
    if len(parts) == 2 {
        if suggestions, exists := PartialSuggestions["show"]; exists {
            if suggestion, exists := suggestions[lastWord]; exists {
                return suggestion
            }
        }
    }

    return ""
}

func (le *LineEditor) getRunSuggestion(parts []string, lastWord string) string {
    if len(parts) == 2 {
        if suggestions, exists := PartialSuggestions["run"]; exists {
            if suggestion, exists := suggestions[lastWord]; exists {
                return suggestion
            }
        }
    }

    return ""
}

func (le *LineEditor) getUseSuggestion(parts []string, lastWord string) string {
    if len(parts) == 2 {
        if suggestions, exists := PartialSuggestions["use"]; exists {
            if suggestion, exists := suggestions[lastWord]; exists {
                return suggestion
            }
        }
    }

    return ""
}

func (le *LineEditor) getScanSuggestion(parts []string, lastWord string) string {
    if len(parts) == 2 {
        if suggestions, exists := PartialSuggestions["scan"]; exists {
            if suggestion, exists := suggestions[lastWord]; exists {
                return suggestion
            }
        }
    }

    return ""
}

func (le *LineEditor) looksLikeUserIsTypingValue(word string) bool {
    if word == "" {
        return false
    }

    for _, r := range word {
        if unicode.IsDigit(r) {
            return true
        }
    }

    commonValuePrefixes := []string{"tru", "fals", "en", "dis", "on", "of", "ye", "no"}
    for _, prefix := range commonValuePrefixes {
        if strings.HasPrefix(strings.ToLower(word), prefix) {
            return true
        }
    }
    
    return false
}

func (le *LineEditor) showSmartSuggestion() {
    suggestion := le.getSmartSuggestion()
    if suggestion == "" {
        return
    }

    currentText := le.currentLine
    if len(currentText) > 0 {
        parts := strings.Fields(currentText)
        if len(parts) > 0 {
            lastWord := parts[len(parts)-1]
            if strings.Contains(suggestion, "<") && strings.Contains(suggestion, ">") {
                if strings.HasPrefix(suggestion, " "+lastWord) || le.currentLineEndsWithPartialTemplate(suggestion) {
                    return
                }
            } else if strings.HasPrefix(suggestion, " "+lastWord) {
                return
            }
        }
    }

    fmt.Printf("%s%s%s", bcolors.Background, bcolors.BrightBlue, suggestion)
    fmt.Printf("%s", bcolors.Endc)

    fmt.Printf("\r\033[%dC", le.promptLen+le.cursorPos)
}

func (le *LineEditor) currentLineEndsWithPartialTemplate(suggestion string) bool {
    if !strings.Contains(suggestion, "<") || !strings.Contains(suggestion, ">") {
        return false
    }

    templateStart := strings.Index(suggestion, "<")
    templateEnd := strings.Index(suggestion, ">")
    if templateStart == -1 || templateEnd == -1 {
        return false
    }

    template := suggestion[templateStart:templateEnd+1]
    currentText := le.currentLine
    if len(currentText) > 0 {
        return strings.Contains(currentText, "<") || 
               (len(template) > 2 && strings.Contains(currentText, template[1:2]))
    }

    return false
}

func (le *LineEditor) handleTabCompletion() {
    le.generatePredictions()

    now := time.Now()
    isDoubleTab := !le.tabPressTime.IsZero() && now.Sub(le.tabPressTime) < 500*time.Millisecond

    if len(le.predictions) == 1 && !isDoubleTab {
        selected := le.predictions[0]
        parts := strings.Fields(le.currentLine)
        
        if len(parts) > 0 {
            lastWord := parts[len(parts)-1]
            if le.isPartialMatch(lastWord, selected) {
                if len(parts) == 1 {
                    le.currentLine = selected
                } else {
                    newParts := parts[:len(parts)-1]
                    newParts = append(newParts, selected)
                    le.currentLine = strings.Join(newParts, " ")
                }
                if le.isCommandThatTakesArgs(selected) {
                    le.currentLine += " "
                }
                le.cursorPos = len(le.currentLine)
                le.redraw()
                le.tabPressTime = now
                return
            }
        }
    }

    if len(le.predictions) > 0 && !isDoubleTab {
        if le.inPredictionMode {
            le.nextPrediction()
        } else {
            le.enterPredictionMode()
        }
    } else {
        if isDoubleTab || len(le.predictions) == 0 {
            le.showAllCommands()
        } else {
            le.showDoubleTabHint()
        }
    }

    le.tabPressTime = now
}

func (le *LineEditor) enterPredictionMode() {
    le.showPredictions = true
    le.predictionPos = 0
    le.inPredictionMode = true
    le.originalLine = le.currentLine
    le.originalCursor = le.cursorPos
    le.showInlinePredictions()
}

func (le *LineEditor) showDoubleTabHint() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)
    fmt.Printf("%s%s(Tab again to show all commands)%s", bcolors.Yellow, bcolors.Bold, bcolors.Endc)
    fmt.Printf("\r\033[%dC", le.promptLen+le.cursorPos)
}

func (le *LineEditor) showAllCommands() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)
    fmt.Printf("\n%s%s            Available commands%s\n", bcolors.Green, bcolors.Bold, bcolors.Endc)

    for category, commands := range CommandCategories {
        fmt.Printf("  %s%s%s\n", bcolors.Yellow, category, bcolors.Endc)
        fmt.Printf("  ")
        for i, cmd := range commands {
            color := bcolors.Green
            if i%2 == 1 {
                color = bcolors.BrightGreen
            }
            fmt.Printf("%s%-12s%s ", color, cmd, bcolors.Endc)
            if (i+1)%4 == 0 && i != len(commands)-1 {
                fmt.Printf("\n  ")
            }
        }
        fmt.Printf("%s", "\n")
        fmt.Printf("%s", "\n")
    }

    fmt.Printf("%s%sPro tip: Start typing and press Tab for smart completions!%s\n", bcolors.Cyan, bcolors.Bold, bcolors.Endc)
    fmt.Printf("%s", "\n")

    fmt.Print(le.prompt + le.currentLine)
    if le.cursorPos < len(le.currentLine) {
        fmt.Printf("\r\033[%dC", le.promptLen+le.cursorPos)
    }
}

func (le *LineEditor) insertChar(c rune) {
    if le.inPredictionMode {
        le.exitPredictionMode()
    }
    originalLine := le.currentLine
    originalCursor := le.cursorPos

    defer func() {
        if r := recover(); r != nil {
            le.currentLine = originalLine
            le.cursorPos = originalCursor
            le.redraw()
            fmt.Printf("\n%s[Recovered from error]%s", bcolors.BrightRed, bcolors.Endc)
        }
    }()

    if le.cursorPos == len(le.currentLine) {
        le.currentLine += string(c)
        fmt.Print(string(c))
        le.cursorPos++
        le.autoInsertSpace()
    } else {
        if le.cursorPos >= 0 && le.cursorPos <= len(le.currentLine) {
            if le.cursorPos == len(le.currentLine) {
                le.currentLine += string(c)
            } else {
                le.currentLine = le.currentLine[:le.cursorPos] + string(c) + le.currentLine[le.cursorPos:]
            }
            le.cursorPos++
            le.redraw()
            le.autoInsertSpace()
        }
    }
}

func (le *LineEditor) autoInsertSpace() {
    if le.cursorPos != len(le.currentLine) {
        return
    }

    currentText := le.currentLine
    if len(currentText) < 3 {
        return
    }

    lastSpaceIndex := strings.LastIndex(currentText, " ")
    if lastSpaceIndex == -1 {
        return
    }

    if lastSpaceIndex+1 >= len(currentText) {
        return
    }

    lastPart := currentText[lastSpaceIndex+1:]
    if len(lastPart) < 2 {
        return
    }

    parts := strings.Fields(currentText)
    if len(parts) < 2 {
        return
    }

    command := strings.ToLower(parts[0])

    for _, commandWord := range CommandWords {
        commandLen := len(commandWord)

        if commandLen < 1 || commandLen >= len(lastPart) {
            continue
        }

        if !strings.HasPrefix(strings.ToLower(lastPart), strings.ToLower(commandWord)) {
            continue
        }

        remaining := lastPart[commandLen:]
        if len(remaining) == 0 {
            continue
        }

        if le.shouldInsertSpaceSimple(command, commandWord, remaining) {
            textBeforeLastPart := currentText[:lastSpaceIndex+1]

            if len(textBeforeLastPart)+len(commandWord)+1+len(remaining) <= len(currentText)+10 {
                newText := textBeforeLastPart + commandWord + " " + remaining

                if len(newText) > len(currentText) && len(newText) <= len(currentText)+20 {
                    le.currentLine = newText
                    le.cursorPos = len(le.currentLine)
                    le.redraw()
                    return
                }
            }
        }
    }
}

func (le *LineEditor) shouldInsertSpaceSimple(command, commandWord, remaining string) bool {

    if command == "set" {
        blacklist := map[string]bool{
            // Function names
            "func": true, "function": true, "funcs": true, "functions": true,
            // Module names
            "module": true, "modules": true,
            // Target names
            "target": true, "targets": true,
            // Option names
            "option": true, "options": true, "value": true, "values": true,
            // Script names
            "script": true, "name": true,
        }

        if blacklist[commandWord] {
            return false
        }

        whitelist := map[string]bool{
            "lhost": true, "lport": true, "rhost": true, "rport": true,
            "proxy": true, "mode": true, "protocol": true, "threads": true,
            "timeout": true, "output": true, "gateway": true, "verbose": true,
        }
        
        return whitelist[commandWord]
    }

    return false
}

func (le *LineEditor) autoFixCommonPatternsSafe(lastPart string, lastSpaceIndex int) {
    currentText := le.currentLine
    textBeforeLastPart := currentText[:lastSpaceIndex+1]

    for _, commandWord := range CommandWords {
        commandLen := len(commandWord)
        
        if commandLen < 1 || commandLen >= len(lastPart) {
            continue
        }

        if !strings.HasPrefix(strings.ToLower(lastPart), strings.ToLower(commandWord)) {
            continue
        }

        remaining := lastPart[commandLen:]
        if len(remaining) == 0 {
            continue
        }

        if le.looksLikeIPAddress(remaining) || le.looksLikePort(remaining) || le.looksLikeCommonValue(remaining) {
            if len(textBeforeLastPart)+len(commandWord)+1+len(remaining) <= len(currentText)+10 {
                newText := textBeforeLastPart + commandWord + " " + remaining
                if len(newText) > len(currentText) {
                    le.currentLine = newText
                    le.cursorPos = len(le.currentLine)
                    le.redraw()
                    return
                }
            }
        }
    }
}

func (le *LineEditor) isCommandWord(word string) bool {
    lowerWord := strings.ToLower(word)
    for _, cmd := range CommandWords {
        if lowerWord == cmd {
            return true
        }
    }
    return false
}

func (le *LineEditor) autoFixCommonPatterns(lastPart string) {
    for _, commandWord := range CommandWords {
        if len(lastPart) > len(commandWord) && strings.HasPrefix(strings.ToLower(lastPart), commandWord) {
            remaining := lastPart[len(commandWord):]

            if le.looksLikeIPAddress(remaining) || le.looksLikePort(remaining) || le.looksLikeCommonValue(remaining) {
                newLastPart := commandWord + " " + remaining
                le.currentLine = le.currentLine[:len(le.currentLine)-len(lastPart)] + newLastPart
                le.cursorPos = len(le.currentLine)
                le.redraw()
                return
            }
        }
    }
}

func (le *LineEditor) looksLikeIPAddress(s string) bool {
    if len(s) < 7 {
        return false
    }
    
    parts := strings.Split(s, ".")
    if len(parts) != 4 {
        return false
    }
    
    for _, part := range parts {
        if len(part) == 0 || len(part) > 3 {
            return false
        }
        for _, r := range part {
            if !unicode.IsDigit(r) {
                return false
            }
        }
    }
    return true
}

func (le *LineEditor) looksLikePort(s string) bool {
    if len(s) == 0 || len(s) > 5 {
        return false
    }

    for _, r := range s {
        if !unicode.IsDigit(r) {
            return false
        }
    }

    if port, err := strconv.Atoi(s); err == nil {
        return port > 0 && port <= 65535
    }
    return false
}

func (le *LineEditor) looksLikeCommonValue(s string) bool {
    if len(s) == 0 {
        return false
    }
    
    commonValues := []string{"true", "false", "on", "off", "enable", "disable", "yes", "no"}
    for _, val := range commonValues {
        if len(s) >= len(val) && strings.HasPrefix(strings.ToLower(s), val) {
            return true
        }
    }
    return false
}

func (le *LineEditor) backspace() {
    if le.inPredictionMode {
        le.exitPredictionMode()
    }

    originalLine := le.currentLine
    originalCursor := le.cursorPos

    defer func() {
        if r := recover(); r != nil {
            le.currentLine = originalLine
            le.cursorPos = originalCursor
            le.redraw()
            fmt.Printf("\n%s[Recovered from backspace error]%s", bcolors.BrightRed, bcolors.Endc)
        }
    }()
    
    if le.cursorPos > 0 && len(le.currentLine) > 0 {
        if le.cursorPos == len(le.currentLine) {
            if len(le.currentLine) > 0 {
                le.currentLine = le.currentLine[:len(le.currentLine)-1]
                fmt.Print("\b \b")
                le.cursorPos--
            }
        } else {
            if le.cursorPos-1 >= 0 && le.cursorPos-1 < len(le.currentLine) && le.cursorPos <= len(le.currentLine) {
                le.currentLine = le.currentLine[:le.cursorPos-1] + le.currentLine[le.cursorPos:]
                le.cursorPos--
                le.redraw()
            }
        }
    }
}

func (le *LineEditor) deleteChar() {
    if le.cursorPos < len(le.currentLine) {
        le.currentLine = le.currentLine[:le.cursorPos] + le.currentLine[le.cursorPos+1:]
        le.redraw()
    }
}

func (le *LineEditor) moveLeft() {
    if le.cursorPos > 0 {
        le.cursorPos--
        le.redraw()
    }
}

func (le *LineEditor) moveRight() {
    if le.cursorPos < len(le.currentLine) {
        le.cursorPos++
        le.redraw()
    }
}

func (le *LineEditor) moveHome() {
    if le.cursorPos != 0 {
        le.cursorPos = 0
        le.redraw()
    }
}

func (le *LineEditor) moveEnd() {
    if le.cursorPos != len(le.currentLine) {
        le.cursorPos = len(le.currentLine)
        le.redraw()
    }
}

func (le *LineEditor) loadHistory() {
    dir := filepath.Dir(le.historyFile)
    os.MkdirAll(dir, 0755)
    
    file, err := os.Open(le.historyFile)
    if err != nil {
        return
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" {
            le.history = append(le.history, line)
        }
    }
    le.historyPos = len(le.history)
}

func (le *LineEditor) saveHistory() {
    dir := filepath.Dir(le.historyFile)
    os.MkdirAll(dir, 0755)
    
    file, err := os.OpenFile(le.historyFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return
    }
    defer file.Close()
    
    if len(le.history) > 0 {
        lastCommand := le.history[len(le.history)-1]
        file.WriteString(lastCommand + "\n")
    }
}

func (le *LineEditor) addToHistory(line string) {
    if strings.TrimSpace(line) != "" {
        le.history = append(le.history, line)
        le.saveHistory()
    }
    le.historyPos = len(le.history)
}

func (le *LineEditor) clearLine() {
    fmt.Print("\r\033[K")
}

func (le *LineEditor) redraw() {
    if !le.initialized {
        return
    }
    
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)

    le.showSmartSuggestion()

    if le.cursorPos < len(le.currentLine) {
        fmt.Printf("\r\033[%dC", le.promptLen+le.cursorPos)
    }
}

func (le *LineEditor) historyUp() {
    if len(le.history) == 0 {
        return
    }

    if le.historyPos > 0 {
        le.historyPos--
        le.currentLine = le.history[le.historyPos]
    } else if le.historyPos == 0 {
        return
    } else {
        le.historyPos = 0
        le.currentLine = le.history[0]
    }

    le.cursorPos = len(le.currentLine)
    le.redraw()
}

func (le *LineEditor) historyDown() {
    if len(le.history) == 0 {
        return
    }

    if le.historyPos < len(le.history)-1 {
        le.historyPos++
        le.currentLine = le.history[le.historyPos]
    } else {
        le.historyPos = len(le.history)
        le.currentLine = ""
    }

    le.cursorPos = len(le.currentLine)
    le.redraw()
}

func (le *LineEditor) handleEscapeSequence(scanner *bufio.Scanner) bool {
    if !scanner.Scan() {
        return false
    }
    first := scanner.Text()
    
    if first == "[" {
        if !scanner.Scan() {
            return false
        }
        key := scanner.Text()
        
        switch key {
        case "A":
            if le.inPredictionMode {
                le.prevPrediction()
            } else {
                le.historyUp()
            }
        case "B":
            if le.inPredictionMode {
                le.nextPrediction()
            } else {
                le.historyDown()
            }
        case "C":
            if le.inPredictionMode {
                le.nextPrediction()
            } else {
                le.moveRight()
            }
        case "D":
            if le.inPredictionMode {
                le.prevPrediction()
            } else {
                le.moveLeft()
            }
        case "H":
            if le.inPredictionMode {
                le.firstPrediction()
            } else {
                le.moveHome()
            }
        case "F":
            if le.inPredictionMode {
                le.lastPrediction()
            } else {
                le.moveEnd()
            }
        case "5~":
            if le.inPredictionMode {
                le.firstPrediction()
            }
        case "6~":
            if le.inPredictionMode {
                le.lastPrediction()
            }
        }
        return true
    } else if first == "O" {
        if !scanner.Scan() {
            return false
        }
        key := scanner.Text()
        switch key {
        case "H":
            if le.inPredictionMode {
                le.firstPrediction()
            } else {
                le.moveHome()
            }
        case "F":
            if le.inPredictionMode {
                le.lastPrediction()
            } else {
                le.moveEnd()
            }
        }
        return true
    }
    return false
}

type terminalState struct {
    original syscall.Termios
}

func makeRaw(fd uintptr) (*terminalState, error) {
    var termios syscall.Termios
    _, _, err := syscall.Syscall6(
        syscall.SYS_IOCTL,
        fd,
        uintptr(syscall.TCGETS),
        uintptr(unsafe.Pointer(&termios)),
        0, 0, 0,
    )
    if err != 0 {
        return nil, err
    }

    original := termios

    termios.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.ISIG
    termios.Cc[syscall.VMIN] = 1
    termios.Cc[syscall.VTIME] = 0

    _, _, err = syscall.Syscall6(
        syscall.SYS_IOCTL,
        fd,
        uintptr(syscall.TCSETS),
        uintptr(unsafe.Pointer(&termios)),
        0, 0, 0,
    )
    if err != 0 {
        return nil, err
    }

    return &terminalState{original: original}, nil
}

func restoreTerminal(fd uintptr, state *terminalState) error {
    _, _, err := syscall.Syscall6(
        syscall.SYS_IOCTL,
        fd,
        uintptr(syscall.TCSETS),
        uintptr(unsafe.Pointer(&state.original)),
        0, 0, 0,
    )
    return err
}

func (le *LineEditor) ReadLine() (string, error) {
    le.currentLine = ""
    le.cursorPos = 0
    le.historyPos = len(le.history)
    le.predictions = []string{}
    le.predictionPos = -1
    le.showPredictions = false
    le.inPredictionMode = false
    le.originalLine = ""
    le.originalCursor = 0
    le.tabPressTime = time.Time{}
    
    fmt.Print(le.prompt)
    
    state, err := makeRaw(uintptr(syscall.Stdin))
    if err != nil {
        return le.fallbackReadLine()
    }
    defer restoreTerminal(uintptr(syscall.Stdin), state)
    
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanBytes)
    
    for scanner.Scan() {
        char := scanner.Text()
        switch char {
        case "\r", "\n":
            if le.inPredictionMode {
                le.applyPrediction()
            } else {
                fmt.Printf("%s", "\n")
                line := le.currentLine
                if line != "" {
                    le.addToHistory(line)
                }
                return line, nil
            }
        case "\x7f", "\x08":
            le.backspace()
        case "\x1b":
            le.handleEscapeSequence(scanner)
        case "\x03":
            return "", fmt.Errorf("interrupted")
        case "\x04":
            if le.currentLine == "" {
                return "", fmt.Errorf("EOF")
            }
        case "\t":
            le.handleTabCompletion()
        default:
            if len(char) > 0 {
                r := rune(char[0])
                if unicode.IsPrint(r) && r != '\r' && r != '\n' {
                    le.insertChar(r)
                }
            }
        }
    }
    return "", scanner.Err()
}

func (le *LineEditor) fallbackReadLine() (string, error) {
    fmt.Print(le.prompt)
    if Scanner.Scan() {
        line := Scanner.Text()
        if line != "" {
            le.addToHistory(line)
        }
        return line, nil
    }
    return "", Scanner.Err()
}

func (le *LineEditor) enableEmoji() {
    le.emojiEnabled = true
}

func (le *LineEditor) enableSound() {
    le.soundEnabled = true
}

func ListJunks() {
    fmt.Printf("%s[*] %sList output\n\n", bcolors.BrightBlue, bcolors.Endc)
    subprocess.Run("tree %s*", OutPutDir)
    return
}

func ClearJunks() {
    fmt.Printf("%s[*] %sClear output\n\n", bcolors.BrightBlue, bcolors.Endc)
    subprocess.Run("tree %s*; rm -rf %s/*", OutPutDir, OutPutDir)
    fmt.Printf("%s[*] %sOutput cleared ...\n", bcolors.BrightGreen, bcolors.Endc)
    return
}

func LocateTool(Tool string, args ...interface{}) error {
    if _, err := exec.LookPath(Tool); err != nil {
        fmt.Printf("%s[!] %s%s%s %sisn't installed. Try %s'apt install %s'%s\n", bcolors.BrightRed, bcolors.Endc, bcolors.Yellow, Tool, bcolors.Endc, bcolors.Green, Tool, bcolors.Endc)
        return err
    }
    return nil
}

func Info(colors, art, message string) {
    fmt.Printf("\n%s%s%s[%s] %s%s ...", bcolors.Endc, bcolors.Bold, colors, art, bcolors.Endc, message)
}

func Sleep() {
    subprocess.Run("sleep")
}

func GSleep(args []string) {
    if len(args) < 2 {
        return
    }

    seconds, err := strconv.Atoi(args[1])
    if err != nil || seconds <= 0 {
        return
    }

    duration := time.Duration(seconds) * time.Second
    time.Sleep(duration)
}

func BrowseTutorials() {
    switch runtime.GOOS {
    case "linux", "darwin":
        Command = `xdg-open "https://youtube.com/@r0jahsm0ntar1/?sub_confirmation=1" 2>/dev/null`
    case "windows":
        Command = `start "" "https://youtube.com/@r0jahsm0ntar1/?sub_confirmation=1"`
    default:
        fmt.Printf("%s[!] %sUnsupported operating system.\n", bcolors.BrightRed, bcolors.Endc)
        return
    }
    fmt.Printf("%s[*] %sLaunched youtube tutarials ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Run(Command)
}

func BrowserLogs() {
    subprocess.Run("mkdir -p /var/www/html/.old/; mv /var/www/html/* /var/www/html/.old/; cd /root/.afr%s/logs/; cat *.log | aha --black > /var/www/html/index.html; xdg-open 'http://0.0.0.0:80/index.html' 2>/dev/null; php -S 0.0.0:80; rm -rf /var/www/html/*; mv /var/www/html/.old/* /var/www/html/; rm -rf /var/www/html/.old/", subprocess.Version)
    return
}

func InitiLize() {
    for _, BaseDirs := range []string{ CertDir, OutPutDir, SetupsLogs, SecLogs, NetworkLogs, ExploitsLogs, PhishersLogs, CrackersLogs, WirelessLogs, WebCrackersLogs } {
        if err := os.MkdirAll(BaseDirs, 0755); err != nil {
            fmt.Printf("%s[!] %sError creating %s: %v%s\n", bcolors.BrightRed, bcolors.Endc, BaseDirs, err, bcolors.Endc)
            return
        }
    }

    if _, err := os.Stat(CertPath); os.IsNotExist(err) {
        GenerateSelfSignedCert(CertPath, KeyPath)
    }
    if err := subprocess.ValidateToolsDir(); err != nil {
        fmt.Printf("%s%s[!] %sError: %s\n", bcolors.Bold, bcolors.Red, bcolors.Endc, err)
        fmt.Printf("%s%s[>] %sPlease launch the installation process to avoid errors ...\n", bcolors.Bold, bcolors.Blue, bcolors.Endc)
    }
}

