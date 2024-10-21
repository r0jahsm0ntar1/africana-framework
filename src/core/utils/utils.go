package utils

import(
    "os"
    "net"
    "fmt"
    "time"
    "bytes"
    "bufio"
    "os/exec"
    "regexp"
    "syscall"
    "bcolors"
    "strings"
    "runtime"
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

var(
    cmd *exec.Cmd
    command string
    gatewayIP string
    userInput string
    scanner = bufio.NewScanner(os.Stdin)
)

func GetDefaultIP()(string, error) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }
    for _, iface := range interfaces {
        if iface.Flags&net.FlagUp == 0 {
            continue
        }
        if iface.Flags&net.FlagLoopback != 0 {
            continue
        }
        addrs, err := iface.Addrs()
        if err != nil {
            return "", err
        }
        for _, addr := range addrs {
            var ip net.IP
            switch neo := addr.(type) {
            case *net.IPNet:
                ip = neo.IP
            case *net.IPAddr:
                ip = neo.IP
            }
            if ip == nil || ip.IsLoopback() {
                continue
            }
            ip = ip.To4()
            if ip == nil {
                continue
            }
            return ip.String(), nil
        }
    }
    return "", fmt.Errorf(bcolors.BLUE + "[+] " + bcolors.GREEN + "no active network interface found\n" + bcolors.ENDC)
}

func GetDefaultGatewayIP()(string, error) {
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/C", "route", "print", "0.0.0.0")
    } else {
        cmd = exec.Command("sh", "-c", "ip route | grep default")
    }
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return "", err
    }
    output := out.String()
    if runtime.GOOS == "windows" {
        re := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
        matches := re.FindStringSubmatch(output)
        if len(matches) > 0 {
            gatewayIP = matches[0]
        }
    } else {
        fields := strings.Fields(output)
        if len(fields) >= 3 {
            gatewayIP = fields[2]
        }
    }
    return gatewayIP, nil
}

func generateSelfSignedCert(certPath, keyPath string) {
    filePath := "/root/.africana/certs/"
    priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
        fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Error creating file: %s\n" + bcolors.ENDC, err)
        return
    }
    notBefore := time.Now()
    notAfter := notBefore.Add(365 * 24 * time.Hour)
    serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
    if err != nil {
        panic(err)
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
        panic(err)
    }
    keyFile, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
        panic(err)
    }

    defer keyFile.Close()
    privBytes, err := x509.MarshalECPrivateKey(priv)
    if err != nil {
        panic(err)
    }
    err = pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes})
    if err != nil {
        panic(err)
    }
    certFile, err := os.Create(certPath)
    if err != nil {
        panic(err)
    }
    defer certFile.Close()
    err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certDER})
    if err != nil {
        panic(err)
    }
}

func Ifaces() {
    interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Error getting interfaces:", err)
        return
    }
    for _, iface := range interfaces {
        fmt.Println("Interface Name:", iface.Name)
        fmt.Println("Hardware Address:", iface.HardwareAddr)
        fmt.Println("Flags:", iface.Flags)
        addrs, err := iface.Addrs()
        if err != nil {
                fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Error getting addresses:", err)
                continue
        }
        for _, addr := range addrs {
                fmt.Println("Address:", addr.String())
        }
        fmt.Println()
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
            fmt.Printf(bcolors.RED + "Error: " + bcolors.DARKCYAN + "Configuring file: " + bcolors.ENDC + "%s %v" + bcolors.ENDC, fileName, err)
        } else {
            fmt.Printf(bcolors.DARKCYAN + "Succes: " + bcolors.ENDC + "Configuring file: %s%s!\n" + bcolors.ENDC, bcolors.RED, fileName)
        }
    }
}

func ClearScreen() {
    switch runtime.GOOS {
    case "linux", "darwin":
        command = "clear"
    case "windows":
        command = "cls"
    default:
        fmt.Println(bcolors.RED + "[*] " + bcolors.ENDC + "Unsupported operating system." + bcolors.ENDC)
        return
    }
    subprocess.Popen(command)
}

func BrowseTutarilas() {
    switch runtime.GOOS {
    case "linux", "darwin":
        command = `open https://youtube.com/@RojahsMontari`
    case "windows":
        command = `start "" "https://youtube.com/@RojahsMontari"`
    default:
        fmt.Println(bcolors.RED + "[*] " + bcolors.ENDC + "Unsupported operating system." + bcolors.ENDC)
        return
    }
    subprocess.Popen(command)
}

func TermLogs() {
    subprocess.Popen(`cat /root/.africana/logs/*.log`)
    for {
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Enter '0' 'exit' " + bcolors.DARKCYAN + "or" + bcolors.DARKGREY + bcolors.ITALIC + "'EXIT' " + bcolors.ENDC + "2 go back" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─✍🏼" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()
        switch strings.ToLower(userInput) {
        case "0", "e", "q", "exit", "quit":
            os.Exit(0)
        case "b", "r", "back", "return":
            return
        default:
            fmt.Println(bcolors.BLUE + "(" + bcolors.RED + "Poor choice of selection. Please select " + bcolors.YELLOW + "🦝00. or" + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 & Go back " + bcolors.BLUE + ")" + bcolors.ENDC)
        }
    }
}

func WordLists() {
    if runtime.GOOS == "linux" {
        filePath := "/usr/share/wordlists/rockyou.txt"
        gzFilePath := filePath + ".gz"

        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            if _, err := os.Stat(gzFilePath); os.IsNotExist(err) {
                return
            }
            command := "gunzip %s"
            subprocess.Popen(command, gzFilePath)
        }
    }
}

func Handler() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGINT)
    go func() {
        <-c
        fmt.Println(bcolors.Colors() + bcolors.BOLD + "\nCTRL-C detected. Exiting.... " + bcolors.ENDC)
        os.Exit(0)
    }()
}

func SystemShell(userInput string) {
    fmt.Printf(bcolors.BLUE + "[*] " + bcolors.ENDC + "exec: %s\n\n", userInput + bcolors.ENDC)
    subprocess.Popen(userInput)
}

func BrowserLogs() {
    subprocess.Popen(`mkdir -p /var/www/html/.old/; mv /var/www/html/* /var/www/html/.old/; cd /root/.africana/logs/; cat *.log | aha --black > /var/www/html/index.html`)
    subprocess.Popen(`xdg-open "http://0.0.0.0:80/index.html" 2>/dev/null; php -S 0.0.0:80`)
    subprocess.Popen(`rm -rf /var/www/html/*; mv /var/www/html/.old/* /var/www/html/; rm -rf /var/www/html/.old/`)
    return
}

func ClearLogs() {
    subprocess.Popen(`ls /root/.africana/logs/; ls /root/.africana/output/; rm -rf /root/.africana/logs/*; rm -rf /root/.africana/output/*`)
    return
}

func ListJunks() {
    subprocess.Popen(`ls /root/.africana/output`)
    return
}

func Certs() {
    certPath := "/root/.africana/certs/africana-cert.pem"
    keyPath := "/root/.africana/certs/africana-key.pem"
    if _, err := os.Stat(certPath); os.IsNotExist(err) {
        if _, err := os.Stat(keyPath); os.IsNotExist(err) {
            generateSelfSignedCert(certPath, keyPath)
        }
    }
}


func Foundations() {
    fileLogs := "/root/.africana/logs/"
    if err := os.MkdirAll(fileLogs, os.ModePerm); err != nil {
        fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Error creating file: %s\n" + bcolors.ENDC, err)
        return
    }
    fileOuts := "/root/.africana/output/"
    if err := os.MkdirAll(fileOuts, os.ModePerm); err != nil {
        fmt.Println(bcolors.BLUE + "[*] " + bcolors.ENDC + "Error creating file: %s\n" + bcolors.ENDC, err)
        return
    }
}

func Sleep() {
    scanner.Scan()
    userInput := scanner.Text()
    subprocess.Popen(`sleep %s`, userInput)
}

func UpsentTools() {
    fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯ ` + bcolors.RED + "Choice selected not implimented, " + bcolors.YELLOW + "comming soon!" + bcolors.ENDC)
}

func InitiLize() {
    Certs()
    Foundations()
    WordLists()
}
