package configures

import (
    "fmt"
    "io/ioutil"
    "strings"
)

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

func replaceInMultipleFiles(filesToReplacements map[string]map[string]string) {
    for fileName, replacements := range filesToReplacements {
        err := replaceStringsInFile(fileName, replacements)
        if err != nil {
            fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error replacing strings in file %s: %v\n", fileName, err)
        } else {
            fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Replacements completed successfully in file %s!\n", fileName)
        }
    }
}

func main() {
    filesToReplacements := map[string]map[string]string{
        "/etc/tor/torrc": {
            "oldWord1": "newWord1",
            "oldWord2": "newWord2",
        },
        "/etc/privoxy/config": {
            "oldWord3": "newWord3",
            "oldWord4": "newWord4",
        },
        "/etc/squid/squid.conf": {
            "oldWord3": "newWord3",
            "oldWord4": "newWord4",
        },
        "/etc/dhcp/dhclient.conf": {
            "oldWord3": "newWord3",
            "oldWord4": "newWord4",
        },
        "/etc/dnsmasq.conf": {
            "#port=5353": "port=5353",
            "oldWord4": "newWord4",
        },
        "/etc/systemd/system/changemac@.service": {
            "oldWord3": "newWord3",
            "oldWord4": "newWord4",
        },
    }
    replaceInMultipleFiles(filesToReplacements)
}
