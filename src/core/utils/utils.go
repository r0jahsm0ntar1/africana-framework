package utils

import(
    "os"
    "io"
    "net"
    "fmt"
    "sync"
    "time"
    "bytes"
    "os/exec"
    "bcolors"
    "bufio"
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
    flag, shell, command, agreementDir, agreementPath, CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, WordList, RokyPath string

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
}

var SpinnerStyles = map[string][]string{
    "classic":    {"|", "/", "-", "\\"},
    "dots":       {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
    "bar":        {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"},
    "arrow":      {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
    "bouncing":   {"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]", "[    ]"},
    "vertical":   {"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂"},
    "horizontal": {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
    "circle":     {"◐", "◓", "◑", "◒"},
    "clock":      {"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"},
    "moon":       {"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"},
    "triangle":   {"◢", "◣", "◤", "◥"},
    "square":     {"◰", "◳", "◲", "◱"},
    "fancy":      {"✦", "✧", "★", "✪", "✯", "✵", "✸", "✹"},
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
        spinChars:  SpinnerStyles["classic"],
        baseFormat: "[+] Starting the africana framework console ...",
        baseArgs:   nil,
        textEffect: TextEffects["plain"],
        speed:      90 * time.Millisecond,
        stopChan:   make(chan struct{}),
        writer:     os.Stdout,
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

func (s *Spinner) Start() {
    s.mu.Lock()
    defer s.mu.Unlock()

    if s.active {
        return
    }

    s.active = true
    s.stopChan = make(chan struct{})

    s.wg.Add(1)
    go s.spin()
}

func (s *Spinner) spin() {
    defer s.wg.Done()

    charIdx := 0
    letterPos := 0
    ticker := time.NewTicker(s.speed)
    defer ticker.Stop()

    for {
        select {
        case <-s.stopChan:
            s.clearLine()
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
    s.clearLine()
    s.flush()
    s.active = false
}

func (s *Spinner) clearLine() {
    fmt.Fprint(s.writer, "\r\033[K")
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

// ClearScreen clears the terminal screen
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

// SystemShell executes a shell command
func SystemShell(Input string) {
    fmt.Printf("%s[*] %sexec: %s\n\n", bcolors.BrightBlue, bcolors.Endc, Input)
    subprocess.Popen(Input)
}

// GenerateSelfSignedCert generates a self-signed certificate
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
            Organization: []string{"The End Time Ministries"},
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

// GetDefaultIP returns the default IP address of the system
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

// Ifaces lists all network interfaces and their addresses
func Ifaces() ([]InterfaceInfo, error) {
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
            // You can choose to continue or return the error
            continue
        }

        for _, addr := range addrs {
            info.Addresses = append(info.Addresses, addr.String())
        }

        result = append(result, info)
    }

    return result, nil
}

// GetDefaultGatewayIP returns the default gateway IP address
func GetDefaultGatewayIP() (string, error) {
    var cmd *exec.Cmd
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

// AskForProxy validates and parses a proxy URL
func AskForProxy(Proxy string) (*url.URL, error) {
    proxyStr := strings.TrimSpace(Proxy)
    proxyURL, err := url.Parse(proxyStr)
    if err != nil || proxyURL.Scheme == "" || proxyURL.Host == "" {
        return nil, fmt.Errorf("invalid PROXY format (eg. http://localhost:80)")
    }

    validSchemes := map[string]bool{"http": true, "https": true, "socks5": true, "socks4": true}
    if !validSchemes[proxyURL.Scheme] {
        return nil, fmt.Errorf("invalid scheme (use http(s), socks4(5))")
    }
    return proxyURL, nil
}

// SetProxy sets the system proxy
func SetProxy(Proxy string) error {
    proxyURL, err := AskForProxy(Proxy)
    if err != nil {
        fmt.Printf("\n%s[!] %s%s\n", bcolors.BrightRed, bcolors.Endc, err)
        return err
    }
    
    if err := os.Setenv("HTTP_PROXY", proxyURL.String()); err != nil {
        return fmt.Errorf("error setting HTTP_PROXY: %w", err)
    }
    if err := os.Setenv("HTTPS_PROXY", proxyURL.String()); err != nil {
        return fmt.Errorf("error setting HTTPS_PROXY: %w", err)
    }
    return nil
}

// replaceStringsInFile replaces strings in a file
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

// Editors replaces strings in multiple files
func Editors(filesToReplacements map[string]map[string]string) {
    for fileName, replacements := range filesToReplacements {
        err := replaceStringsInFile(fileName, replacements)
        if err != nil {
            fmt.Printf("%s[!] %sError Configuring: %s%v", bcolors.BrightRed, bcolors.Endc, fileName, err)
        } else {
            fmt.Printf("%s[*] %sDone configuring: %s%s%s", bcolors.Green, bcolors.Endc, bcolors.BrightBlue, fileName, bcolors.Endc)
        }
    }
}

// BrowseTutarilas opens a YouTube tutorial link in the default browser
func BrowseTutarilas() {
    switch runtime.GOOS {
    case "linux", "darwin":
        command = `xdg-open "https://youtube.com/@RojahsMontari" 2>/dev/null`
    case "windows":
        command = `start "" "https://youtube.com/@RojahsMontari"`
    default:
        fmt.Printf("%s[!] %sUnsupported operating system.\n", bcolors.BrightRed, bcolors.Endc)
        return
    }
    fmt.Printf("%s[*] %sLaunched youtube tutarials ...\n", bcolors.Green, bcolors.Endc)
    subprocess.Popen(command)
}

// BrowserLogs moves logs to a web directory and opens them in a browser
func BrowserLogs() {
    subprocess.Popen(`mkdir -p /var/www/html/.old/; mv /var/www/html/* /var/www/html/.old/; cd /root/.afr3/logs/; cat *.log | aha --black > /var/www/html/index.html`)
    subprocess.Popen(`xdg-open "http://0.0.0.0:80/index.html" 2>/dev/null; php -S 0.0.0:80`)
    subprocess.Popen(`rm -rf /var/www/html/*; mv /var/www/html/.old/* /var/www/html/; rm -rf /var/www/html/.old/`)
    return
}


// ListJunks lists all junk files in the output directory
func ListJunks() {
    fmt.Printf("%s[*] %sList output\n\n", bcolors.BrightBlue, bcolors.Endc)
    subprocess.Popen("ls -a /root/.afr3/output")
    return
}

// ClearJunks clears all junk files in the output directory
func ClearJunks() {
    fmt.Printf("%s[*] %sClear output\n\n", bcolors.BrightBlue, bcolors.Endc)
    subprocess.Popen(`rm -rf /root/.afr3/output/*`)
    fmt.Printf("%s[*] %sOutput cleared ...\n", bcolors.BrightGreen, bcolors.Endc)
    return
}

// Sleep pauses the execution for a specified duration (in seconds)
func Sleep() {
    subprocess.Popen("sleep")
}

// GetAgreementPath returns the path to the agreement file based on the 's privilege level
func GetAgreementPath() string {
    baseDir := "/root/.afr3/agreements/"
    if !subprocess.IsRoot() {
        homeDir := subprocess.GetHomeDir()
        if homeDir == "" {
            return ""
        }
        baseDir = filepath.Join(homeDir, ".afr3", "agreements")
    }
    return filepath.Join(baseDir, "covenant.txt")
}

// Agreements creates the agreement file if it doesn't exist
func Agreements(filePath string) {
    dirPath := filepath.Dir(filePath)

    // Create the directory if it doesn't exist
    if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
        fmt.Printf("[!] Error creating directory: %v\n", err)
        return
    }

    // Create the agreement file if it doesn't exist
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        if err := ioutil.WriteFile(filePath, []byte("user accepted to the agreement."), os.ModePerm); err != nil {
            fmt.Printf("[!] Error writing file: %v\n", err)
            return
        }
    }
}

// Sealing ensures the agreement file exists
func Sealing() {
    filePath := GetAgreementPath()
    if filePath == "" {
        fmt.Println("[!] Error: Could not determine agreement file path.")
        return
    }
    Agreements(filePath)
}

// EnhancedCopy copies files/directories with overwrite support
func Copy(src, dst string) error {
    srcInfo, err := os.Stat(src)
    if err != nil {
        return fmt.Errorf("could not stat source: %w", err)
    }

    if srcInfo.IsDir() {
        return enhancedCopyDir(src, dst)
    }
    return enhancedCopyFile(src, dst, true) // true = allow overwrite
}

// enhancedCopyFile handles file copying with overwrite control
func enhancedCopyFile(src, dst string, overwrite bool) error {
    // If destination exists and is a directory, append source filename
    dstInfo, err := os.Stat(dst)
    if err == nil && dstInfo.IsDir() {
        dst = filepath.Join(dst, filepath.Base(src))
    }

    // Check if destination exists and overwrite is disabled
    if _, err := os.Stat(dst); err == nil && !overwrite {
        return fmt.Errorf("destination exists (overwrite disabled): %s", dst)
    }

    // Ensure parent directory exists
    if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
        return fmt.Errorf("failed to create parent directory: %w", err)
    }

    // Open source file
    in, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("failed to open source: %w", err)
    }
    defer in.Close()

    // Create destination file (truncates if exists)
    out, err := os.Create(dst)
    if err != nil {
        return fmt.Errorf("failed to create destination: %w", err)
    }
    defer out.Close()

    // Copy with buffer for better performance
    if _, err := io.Copy(out, in); err != nil {
        return fmt.Errorf("copy failed: %w", err)
    }

    // Preserve permissions and timestamps
    srcInfo, err := os.Stat(src)
    if err != nil {
        return fmt.Errorf("failed to get source info: %w", err)
    }
    if err := os.Chmod(dst, srcInfo.Mode()); err != nil {
        return fmt.Errorf("failed to set permissions: %w", err)
    }
    if err := os.Chtimes(dst, srcInfo.ModTime(), srcInfo.ModTime()); err != nil {
        return fmt.Errorf("failed to set timestamps: %w", err)
    }
    return nil
}

// enhancedCopyDir handles directory copying
func enhancedCopyDir(srcDir, dstDir string) error {
    srcInfo, err := os.Stat(srcDir)
    if err != nil {
        return fmt.Errorf("could not stat source dir: %w", err)
    }

    // Create destination directory
    if err := os.MkdirAll(dstDir, srcInfo.Mode()); err != nil {
        return fmt.Errorf("failed to create destination dir: %w", err)
    }

    // Read source directory
    entries, err := os.ReadDir(srcDir)
    if err != nil {
        return fmt.Errorf("failed to read source dir: %w", err)
    }

    // Copy each entry
    for _, entry := range entries {
        srcPath := filepath.Join(srcDir, entry.Name())
        dstPath := filepath.Join(dstDir, entry.Name())

        if entry.IsDir() {
            if err := enhancedCopyDir(srcPath, dstPath); err != nil {
                return err
            }
        } else {
            if err := enhancedCopyFile(srcPath, dstPath, true); err != nil {
                return err
            }
        }
    }

    // Preserve directory timestamps
    return os.Chtimes(dstDir, srcInfo.ModTime(), srcInfo.ModTime())
}

func EncodeFileToPowerShellEncodedCommand(filePath string) (string, error) {
    // Read the file content
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return "", err
    }

    // Clean any potential BOM or null characters (optional but good practice)
    cleaned := string(bytes.Trim(content, "\x00"))

    // Convert to UTF-16LE
    buf := new(bytes.Buffer)
    for _, r := range cleaned {
        buf.WriteByte(byte(r))        // Low byte
        buf.WriteByte(byte(r >> 8))   // High byte
    }

    // Encode to Base64 (same as `base64 -w 0`)
    encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
    return encoded, nil
}

// getLinuxDistroID attempts to parse /etc/os-release and get the ID
func GetLinuxDistroID() (string, error) {
    file, err := os.Open("/etc/os-release")
    if err != nil {
        return "", err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "ID=") {
            return strings.Trim(strings.TrimPrefix(line, "ID="), `"`), nil
        }
    }

    return "", fmt.Errorf("could not determine Linux distro")
}

func DetectAndroid() bool {
    // Android typically has GOOS=linux, but some clues can help
    // like the presence of `/system/build.prop`
    if _, err := os.Stat("/system/build.prop"); err == nil {
        return true
    }
    return false
}

// DirLocations returns the paths for certificates, output, and tools directories
func DirLocations() (string, string, string, string, string, string, string) {
    if subprocess.IsRoot() {
        CertDir = "/root/.afr3/certs/"
        OutPutDir = "/root/.afr3/output/"
        ToolsDir = "/root/.afr3/africana-base/"
        RokyPath  = "/usr/share/wordlists/rockyou.txt"
        KeyPath = filepath.Join(CertDir, "afr_key.pem")
        CertPath = filepath.Join(CertDir, "afr_cert.pem")
        WordList = filepath.Join(ToolsDir, "wordlists", "everything.txt")

    } else {
        homeDir := subprocess.GetHomeDir()
        if homeDir == "" {
            return "", "", "", "", "", "", ""
        }
        RokyPath  = "/usr/share/wordlists/rockyou.txt"
        CertDir = filepath.Join(homeDir, ".afr3", "certs")
        KeyPath = filepath.Join(CertDir, "afr_key.pem")
        OutPutDir = filepath.Join(homeDir, ".afr3", "output")
        CertPath = filepath.Join(CertDir, "afr_cert.pem")
        ToolsDir = filepath.Join(homeDir, ".afr3", "africana-base")
        WordList = filepath.Join(homeDir, "wordlists", "everything.txt")
    }
    return CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, RokyPath, WordList
}

// Levenshtein computes the edit distance between two strings (case-insensitive).
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
                currRow[j-1] + 1,      // Insert
                prevRow[j] + 1,        // Delete
                prevRow[j-1] + cost,   // Replace
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

// InitializePaths sets the correct paths based on whether the  is root or not
func InitiLize() {
    CertDir, OutPutDir, _, _, _, _, _ := DirLocations()

    // Create directories if they don't exist
    for _, dir := range []string{CertDir, OutPutDir} {
        if err := os.MkdirAll(dir, os.ModePerm); err != nil {
            fmt.Printf("%s[!] %sError creating directory %s: %s\n", bcolors.BrightRed, bcolors.Endc, dir, err)
            return
        }
    }

    // Generate self-signed certificate if it doesn't exist
    if _, err := os.Stat(CertPath); os.IsNotExist(err) {
        GenerateSelfSignedCert(CertPath, KeyPath)
    }
}
