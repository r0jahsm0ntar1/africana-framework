package utils

import(
    "os"
    "net"
    "fmt"
    "time"
    "os/exec"
    "bcolors"
    "strings"
    "runtime"
    "net/url"
    "math/big"
    "path/filepath"
    "io/ioutil"
    "subprocess"
    "crypto/rand"
    "crypto/x509"
    "crypto/ecdsa"
    "encoding/pem"
    "crypto/elliptic"
    "crypto/x509/pkix"
)

var (
    cmd *exec.Cmd
    flag, shell, command, agreementDir, agreementPath, CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, WordList, KyPath string

)

// ClearScreen clears the terminal screen
func ClearScreen() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }

    cmd.Stdout = os.Stdout
    if err := cmd.Run(); err != nil {
        fmt.Printf("%s[!] %sError clearing screen: %s\n", bcolors.RED, bcolors.ENDC, err)
    }
}

// SystemShell executes a shell command
func SystemShell(Input string) {
    fmt.Printf("%s[*] %sexec: %s\n\n", bcolors.BLUE, bcolors.ENDC, Input)
    subprocess.Popen(Input)
}

// GenerateSelfSignedCert generates a self-signed certificate
func GenerateSelfSignedCert(certPath, keyPath string) {
    priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        fmt.Printf("Error generating key:", err)
        return
    }

    notBefore := time.Now()
    notAfter := notBefore.Add(365 * 24 * time.Hour)
    serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
    if err != nil {
        fmt.Printf("Error generating serial number:", err)
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
        fmt.Printf("Error creating certificate:", err)
        return
    }

    keyFile, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
        fmt.Printf("Error creating key file:", err)
        return
    }
    defer keyFile.Close()

    privBytes, err := x509.MarshalECPrivateKey(priv)
    if err != nil {
        fmt.Printf("Error marshaling private key:", err)
        return
    }
    pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes})

    certFile, err := os.Create(certPath)
    if err != nil {
        fmt.Printf("Error creating certificate file:", err)
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
func Ifaces() {
    interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Printf("%s[!] %sError getting interfaces: %s", bcolors.RED, bcolors.ENDC, err)
        return
    }
    for _, iface := range interfaces {
        fmt.Printf("Interface Name: ", iface.Name)
        fmt.Printf("Hardware Address: ", iface.HardwareAddr)
        fmt.Printf("Flags: ", iface.Flags)
        addrs, err := iface.Addrs()
        if err != nil {
            fmt.Printf("%s[!] %sError getting addresses: %s", bcolors.RED, bcolors.ENDC, err)
            continue
        }
        for _, addr := range addrs {
            fmt.Printf("Address: ", addr.String())
        }
    }
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
func AskForProxy(Proxy string) *url.URL {
    for {
        proxyStr := strings.TrimSpace(Proxy)
        proxyURL, err := url.Parse(proxyStr)
        if err != nil || proxyURL.Scheme == "" || proxyURL.Host == "" {
            fmt.Printf("%s[!] %sInvalid URL format. eg. http://localhost:80).\n", bcolors.RED, bcolors.ENDC)
            return nil
        }

        validSchemes := map[string]bool{"http": true, "https": true, "socks5": true, "socks4": true}
        if !validSchemes[proxyURL.Scheme] {
            fmt.Printf("%s[!] %sInvalid scheme. Use, http(s), socks4(5).\n", bcolors.RED, bcolors.ENDC)
            return nil
        }
        return proxyURL
    }
}

// SetProxyEnv sets the proxy environment variables
func SetProxyEnv(proxyURL *url.URL) error {
    if err := os.Setenv("HTTP_PROXY", proxyURL.String()); err != nil {
        return err
    }
    if err := os.Setenv("HTTPS_PROXY", proxyURL.String()); err != nil {
        return err
    }
    return nil
}

// SetProxy sets the system proxy
func SetProxy(Proxy string) {
    proxyURL := AskForProxy(Proxy)
    if err := SetProxyEnv(proxyURL); err != nil {
        fmt.Printf("%s[!] %sError setting proxy environment Variables: %s", bcolors.RED, bcolors.ENDC, err)
        return
    }
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
            fmt.Printf("%s[!] %sError Configuring: %s%v", bcolors.RED, bcolors.ENDC, fileName, err)
        } else {
            fmt.Printf("%s[*] %sDone configuring: %s%s%s", bcolors.GREEN, bcolors.ENDC, bcolors.BLUE, fileName, bcolors.ENDC)
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
        fmt.Printf("%s[!] %sUnsupported operating system.\n", bcolors.RED, bcolors.ENDC)
        return
    }
    fmt.Printf("%s[*] %sLaunched youtube tutarials ...\n", bcolors.GREEN, bcolors.ENDC)
    subprocess.Popen(command)
}

// BrowserLogs moves logs to a web directory and opens them in a browser
func BrowserLogs() {
    subprocess.Popen(`mkdir -p /var/www/html/.old/; mv /var/www/html/* /var/www/html/.old/; cd /root/.afr3/logs/; cat *.log | aha --black > /var/www/html/index.html`)
    subprocess.Popen(`xdg-open "http://0.0.0.0:80/index.html" 2>/dev/null; php -S 0.0.0:80`)
    subprocess.Popen(`rm -rf /var/www/html/*; mv /var/www/html/.old/* /var/www/html/; rm -rf /var/www/html/.old/`)
    return
}

// ClearLogs clears all logs
func ClearLogs() {
    subprocess.Popen(`ls /root/.afr3/logs/; ls /root/.afr3/output/; rm -rf /root/.afr3/logs/*; rm -rf /root/.afr3/output/*`)
    return
}

// ListJunks lists all junk files in the output directory
func ListJunks() {
    subprocess.Popen(`ls /root/.afr3/output`)
    return
}

// ClearJunks clears all junk files in the output directory
func ClearJunks() {
    subprocess.Popen(`rm -rf /root/.afr3/output/*`)
    fmt.Printf("%s[*] %s Succesfully cleared All junks.", bcolors.GREEN, bcolors.ENDC)
    return
}

// Sleep pauses the execution for a specified duration (in seconds)
func Sleep(){ //(seconds int) {
    //if seconds < 0 {
    //    fmt.Println("Error: Duration cannot be negative.")
    //    return
    //}
    // Pause execution for the specified duration
    //time.Sleep(time.Duration(seconds) * time.Second)
    //fmt.Printf("Slept for %d seconds.\n", seconds)
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

// DirLocations returns the paths for certificates, output, and tools directories
func DirLocations() (string, string, string, string, string, string) {
    if subprocess.IsRoot() {
        CertDir = "/root/.afr3/certs"
        OutPutDir = "/root/.afr3/output"
        ToolsDir = "/root/.afr3/africana-base"
        KeyPath = filepath.Join(CertDir, "africana-key.pem")
        CertPath = filepath.Join(CertDir, "africana-cert.pem")
        WordList = filepath.Join(ToolsDir, "wordlists", "everything.txt")
    } else {
        homeDir := subprocess.GetHomeDir()
        if homeDir == "" {
            return "", "", "", "", "", ""
        }
        CertDir = filepath.Join(homeDir, ".afr3", "certs")
        OutPutDir = filepath.Join(homeDir, ".afr3", "output")
        ToolsDir = filepath.Join(homeDir, ".afr3", "africana-base")
        KeyPath = filepath.Join(CertDir, "africana-key.pem")
        CertPath = filepath.Join(CertDir, "africana-cert.pem")
        WordList = filepath.Join(homeDir, "wordlists", "everything.txt")
    }
    return CertDir, OutPutDir, KeyPath, CertPath, ToolsDir, WordList
}

// InitializePaths sets the correct paths based on whether the  is root or not
func InitiLize() {
    CertDir, OutPutDir, _, _, _, _ := DirLocations()
    KyPath  = "/usr/share/wordlist/rockyou.txt"

    // Create directories if they don't exist
    for _, dir := range []string{CertDir, OutPutDir} {
        if err := os.MkdirAll(dir, os.ModePerm); err != nil {
            fmt.Printf("%s[!] %sError creating directory %s: %s\n", bcolors.RED, bcolors.ENDC, dir, err)
            return
        }
    }

    // Generate self-signed certificate if it doesn't exist
    if _, err := os.Stat(CertPath); os.IsNotExist(err) {
        GenerateSelfSignedCert(CertPath, KeyPath)
    }

    // Handle rockyou.txt.gz for Linux
    if runtime.GOOS == "linux" {
        gzFilePath := KyPath + ".gz"
        if _, err := os.Stat(KyPath); os.IsNotExist(err) {
            if _, err := os.Stat(gzFilePath); os.IsNotExist(err) {
                return
            }
            command := fmt.Sprintf("gunzip %s", gzFilePath)
            subprocess.Popen(command)
        }
    }
}
