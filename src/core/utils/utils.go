package utils

import(
    "os"
    "net"
    "fmt"
    "time"
    "os/user"
    "os/exec"
    "bcolors"
    "strings"
    "runtime"
    "net/url"
    "math/big"
    "os/signal"
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
    flag       string
    shell      string
    command    string
    certDir   = "/root/.afr3/certs/"
    outputDir = "/root/.afr3/output/"
    agreementDir = "/root/.afr3/agreements/"
    agreement = "User accepted the terms."
    keyPath   = certDir + "africana-key.pem"
    certPath  = certDir + "africana-cert.pem"
    agreementPath = agreementDir + "covenant.txt"
    rkyPath   = "/usr/share/wordlists/rockyou.txt"
)

func isRoot() bool {
    return os.Geteuid() == 0
}

func ClearScreen() {
    switch runtime.GOOS {
    case "windows":
        shell = "cmd"
        flag = "/c"
        command = "cls"
    default:
        shell = "bash"
        flag = "-c"
        command = "clear"
    }
    cmd := exec.Command(shell, flag, command)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt)

    go func() {
        <-sigs
        if cmd.Process != nil {
            _ = cmd.Process.Signal(os.Interrupt)
        }
    }()

    if err := cmd.Start(); err != nil {
        msg, _ := fmt.Printf("%s[!] %sError starting command ... %s%s_", bcolors.RED, bcolors.BLUE, bcolors.ENDC, err)
        fmt.Fprintln(os.Stderr, msg)
        return
    }

    if err := cmd.Wait(); err != nil {
        msg, _ := fmt.Printf("%s[!] %sProcess is incomplete ... %s%s_", bcolors.RED, bcolors.ENDC, bcolors.ENDC, err)
        fmt.Fprintln(os.Stderr, msg)
    }

    signal.Stop(sigs)
    close(sigs)
}


func getUserHomeDir() string {
    usr, err := user.Current()
    if err != nil {
        fmt.Println("Error getting current user:", err)
        return ""
    }
    return usr.HomeDir
}

func GetAgreementPath() string {
    if isRoot() {
        return agreementPath
    } else {
        homeDir := getUserHomeDir()
        return homeDir + "/.afr3/agreements/covenant.txt"
    }
}

func userAgreements(filePath string) {
    dirPath := agreementDir
    if !isRoot() {
        dirPath = getUserHomeDir() + "/.afr3/agreements/"
    }

    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
            fmt.Printf("[!] Error creating directory: %v\n", err)
            return
        }
        err = ioutil.WriteFile(filePath, []byte("User agreement accepted."), os.ModePerm)
        if err != nil {
            fmt.Printf("[!] Error writing file: %v\n", err)
            return
        }
    }
}

func UserSealing() {
    filePath := GetAgreementPath()
    userAgreements(filePath)
}

func SystemShell(userInput string) {
    fmt.Printf("%s[*] %sexec: %s\n\n", bcolors.BLUE, bcolors.ENDC, userInput)
    subprocess.Popen(userInput)
}

func GenerateSelfSignedCert(certPath, keyPath string) {
    priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        fmt.Println("Error generating key:", err)
        return
    }

    notBefore := time.Now()
    notAfter := notBefore.Add(365 * 24 * time.Hour)
    serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
    if err != nil {
        fmt.Println("Error generating serial number:", err)
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
        fmt.Println("Error creating certificate:", err)
        return
    }

    keyFile, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
        fmt.Println("Error creating key file:", err)
        return
    }
    defer keyFile.Close()

    privBytes, err := x509.MarshalECPrivateKey(priv)
    if err != nil {
        fmt.Println("Error marshaling private key:", err)
        return
    }
    pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes})

    certFile, err := os.Create(certPath)
    if err != nil {
        fmt.Println("Error creating certificate file:", err)
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

func AskForProxy(userProxy string) *url.URL {
    for {
        proxyStr := strings.TrimSpace(userProxy)
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

func SetProxyEnv(proxyURL *url.URL) error {
    if err := os.Setenv("HTTP_PROXY", proxyURL.String()); err != nil {
        return err
    }
    if err := os.Setenv("HTTPS_PROXY", proxyURL.String()); err != nil {
        return err
    }
    return nil
}

func SetProxy(userProxy string) {
    proxyURL := AskForProxy(userProxy)
    if err := SetProxyEnv(proxyURL); err != nil {
        fmt.Printf("%s[!] %sError setting proxy environment Variables: %s", bcolors.RED, bcolors.ENDC, err)
        return
    }
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
            fmt.Printf("%s[!] %sError Configuring: %s%v", bcolors.RED, bcolors.ENDC, fileName, err)
        } else {
            fmt.Printf("%s[*] %sDone configuring: %s%s%s ...", bcolors.GREEN, bcolors.ENDC, bcolors.BLUE, fileName, bcolors.ENDC)
        }
    }
}

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

func BrowserLogs() {
    subprocess.Popen(`mkdir -p /var/www/html/.old/; mv /var/www/html/* /var/www/html/.old/; cd /root/.afr3/logs/; cat *.log | aha --black > /var/www/html/index.html`)
    subprocess.Popen(`xdg-open "http://0.0.0.0:80/index.html" 2>/dev/null; php -S 0.0.0:80`)
    subprocess.Popen(`rm -rf /var/www/html/*; mv /var/www/html/.old/* /var/www/html/; rm -rf /var/www/html/.old/`)
    return
}

func ClearLogs() {
    subprocess.Popen(`ls /root/.afr3/logs/; ls /root/.afr3/output/; rm -rf /root/.afr3/logs/*; rm -rf /root/.afr3/output/*`)
    return
}

func ListJunks() {
    subprocess.Popen(`ls /root/.afr3/output`)
    return
}

func ClearJunks() {
    subprocess.Popen(`rm -rf /root/.afr3/output/*`)
    fmt.Printf("%s[*] %s Succesfully cleared All junks.", bcolors.GREEN, bcolors.ENDC)
    return
}

func Sleep() {
    subprocess.Popen("sleep")
}

func InitiLize() {
    if !isRoot() {
        homeDir := getUserHomeDir()
        if homeDir == "" {
            return
        }
        certDir = homeDir + "/.afr3/certs/"
        outputDir = homeDir + "/.afr3/output/"
        keyPath = certDir + "africana-key.pem"
        certPath = certDir + "africana-cert.pem"
    }

    if err := os.MkdirAll(certDir, os.ModePerm); err != nil {
        msg, _ := fmt.Printf("%s[!] %sError creating cert directory ... %s%s_", bcolors.RED, bcolors.ENDC, bcolors.ENDC, err)
        fmt.Fprintln(os.Stderr, msg)
        return
    }
    if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
        msg, _ := fmt.Printf("%s[!] %sError creating output directory ... %s%s_", bcolors.RED, bcolors.ENDC, bcolors.ENDC, err)
        fmt.Fprintln(os.Stderr, msg)
        return
    }

    if _, err := os.Stat(certPath); os.IsNotExist(err) {
        GenerateSelfSignedCert(certPath, keyPath)
    }

    if runtime.GOOS == "linux" {
        gzFilePath := rkyPath + ".gz"
        if _, err := os.Stat(rkyPath); os.IsNotExist(err) {
            if _, err := os.Stat(gzFilePath); os.IsNotExist(err) {
                return
            }
            command := "gunzip %s"
            subprocess.Popen(command, gzFilePath)
         }
    }
}
