package utils

import (
    "os"
    "net"
    "fmt"
    "time"
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

func GetDefaultIP() (string, error) {
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

func generateSelfSignedCert(certPath, keyPath string) {
    filePath := "/root/.africana/certs/"
    priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
        fmt.Println(bcolors.BLUE + "[+] " + bcolors.GREEN + "Error creating file: %s\n" + bcolors.ENDC, err)
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
            fmt.Printf(bcolors.RED + "\nError:" + bcolors.DARKCYAN + "Replacing " + bcolors.ENDC + "strings in file %s: %v" + bcolors.ENDC, fileName, err)
        } else {
            fmt.Printf(bcolors.DARKCYAN + "\nSucces:" + bcolors.ENDC + "Replacing file: %s%s!\n" + bcolors.ENDC, bcolors.RED, fileName)
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

func ClearScreen() {
    if runtime.GOOS == "linux" {
        subprocess.Popen("clear")
    } else if runtime.GOOS == "windows" {
        subprocess.Popen("cls")
    } else {
        fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Unsupported operating system.\n" + bcolors.ENDC)
    }
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

func WordLists() {
    if runtime.GOOS == "linux" {
        filePath := "/usr/share/wordlists/rockyou.txt"
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            subprocess.Popen(`gunzip /usr/share/wordlists/rockyou.txt.gz`)
        }
    }
}

func Foundations() {
    fileLogs := "/root/.africana/logs/"
    if err := os.MkdirAll(fileLogs, os.ModePerm); err != nil {
        fmt.Println(bcolors.BLUE + "[+] " + bcolors.GREEN + "Error creating file: %s\n" + bcolors.ENDC, err)
        return
    }
    fileOuts := "/root/.africana/output"
    if err := os.MkdirAll(fileOuts, os.ModePerm); err != nil {
        fmt.Println(bcolors.BLUE + "[+] " + bcolors.GREEN + "Error creating file: %s\n" + bcolors.ENDC, err)
        return
    }
}

func UpsentTools() {
    fmt.Println(bcolors.ENDC + "\n" + `¯\_(ツ)_/¯` + bcolors.RED + "Choice selected not yet implimented, " + bcolors.YELLOW + "comming soon!\n" + bcolors.ENDC)
}

func InitiLize() {
    Certs(); Foundations(); WordLists()
}
