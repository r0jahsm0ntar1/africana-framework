package bcolors

import (
    "fmt"
    "runtime"
    "math/rand"
    "time"
)

var PURPLE = "\033[95m"
var CYAN = "\033[96m"
var DARKCYAN = "\033[36m"
var BLUE = "\033[94m"
var GREEN = "\033[92m"
var YELLOW = "\033[93m"
var RED = "\033[91m"
var BOLD = "\033[1m"
var UNDERL = "\033[4m"
var ENDC = "\033[0m"
var backBlack = "\033[40m"
var backRed = "\033[41m"
var backGreen = "\033[42m"
var backYellow = "\033[43m"
var backBlue = "\033[44m"
var backMagenta = "\033[45m"
var backCyan = "\033[46m"
var backWhite = "\033[47m"

func init() {
    if runtime.GOOS == "windows" {
        PURPLE = ""
        CYAN = ""
        BLUE = ""
        GREEN = ""
        YELLOW = ""
        RED = ""
        ENDC = ""
        BOLD = ""
        UNDERL = ""
        backBlack = ""
        backRed = ""
        backGreen = ""
        backYellow = ""
        backBlue = ""
        backMagenta = ""
        backCyan = ""
        backWhite = ""
        DARKCYAN = ""
    }
}

func Colors(words string) {
    colors := []string{ BLUE, PURPLE, CYAN, DARKCYAN, GREEN, YELLOW, RED,
    }
    rand.Seed(time.Now().UnixNano())
    randomIndex := rand.Intn(len(colors))
    fmt.Printf("%s%s\n", colors[randomIndex], words)
}
