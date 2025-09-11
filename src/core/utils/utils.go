//John 3:16

package utils

import(
    "os"
    "io"
    "net"
    "log"
    "fmt"
    "sync"
    "time"
    "bytes"
    "bufio"
    "regexp"
    "strconv"
    "os/exec"
    "bcolors"
    "strings"
    "runtime"
    "net/url"
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
    PyEnvName  = "afr_venv"
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


    LHost, _ = GetDefaultIP()
    Distro, _ = GetLinuxDistroID()
    Gateway, _ = GetDefaultGatewayIP()
    Scanner = bufio.NewScanner(os.Stdin)

    Date = time.Now().Format("2006-01-02.15.04.05")

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

func AskForProxy(Proxies string) (*url.URL, error) {
    proxyStr := strings.TrimSpace(Proxies)
    proxyURL, err := url.Parse(proxyStr)
    if err != nil || proxyURL.Scheme == "" || proxyURL.Host == "" {
        return nil, fmt.Errorf("Invalid PROXY format (eg. http://localhost:80)")
    }

    validSchemes := map[string]bool{"http": true, "https": true, "socks5": true, "socks4": true}
    if !validSchemes[proxyURL.Scheme] {
        return nil, fmt.Errorf("Invalid scheme (use http(s), socks4(5))")
    }
    return proxyURL, nil
}

func SetProxy(Proxies string) error {
    proxyURL, err := AskForProxy(Proxies)
    if err != nil {
        fmt.Printf("\n%s[!] %s%s\n", bcolors.BrightRed, bcolors.Endc, err)
        return err
    }
    
    if err := os.Setenv("HTTP_PROXY", proxyURL.String()); err != nil {
        return fmt.Errorf("Error setting HTTP_PROXY: %w", err)
    }
    if err := os.Setenv("HTTPS_PROXY", proxyURL.String()); err != nil {
        return fmt.Errorf("Error setting HTTPS_PROXY: %w", err)
    }
    return nil
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
        fmt.Println()
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
        fmt.Println("[!] Error: Could not determine agreement file path.")
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
    if _, err := os.Stat("/system/build.prop"); err == nil {
        return true
    }
    return false
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
}
