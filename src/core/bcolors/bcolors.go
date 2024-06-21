package bcolors

import (
    "runtime"
    "math/rand"
)

var (
    ENDC        = "\033[0m"
    BOLD        = "\033[1m"
    ITALIC      = "\033[3m"
    UNDERL      = "\033[4m"
    BLINK       = "\033[5m"
    ORANGE      = "\033[33m"
    MAGENTA     = "\033[35m"
    DARKCYAN    = "\033[36m"
    backBlack   = "\033[40m"
    backRed     = "\033[41m"
    backGreen   = "\033[42m"
    backYellow  = "\033[43m"
    backBlue    = "\033[44m"
    backMagenta = "\033[45m"
    backCyan    = "\033[46m"
    backWhite   = "\033[47m"
    RED         = "\033[91m"
    GREEN       = "\033[92m"
    YELLOW      = "\033[93m"
    BLUE        = "\033[94m"
    PURPLE      = "\033[95m"
    CYAN        = "\033[96m"
    colorTaken  = []string{}
)

func init() {
    if runtime.GOOS == "windows" {
        ENDC        = ""
        BOLD        = ""
        ITALIC      = ""
        UNDERL      = ""
        BLINK       = ""
        ORANGE      = ""
        MAGENTA     = ""
        DARKCYAN    = ""
        backBlack   = ""
        backRed     = ""
        backGreen   = ""
        backYellow  = ""
        backBlue    = ""
        backMagenta = ""
        backCyan    = ""
        backWhite   = ""
        RED         = ""
        GREEN       = ""
        YELLOW      = ""
        BLUE        = ""
        PURPLE      = ""
        CYAN        = ""
    }
}

func Colors(args ...string) string {
    colors := []string{MAGENTA, BLUE, PURPLE, CYAN, DARKCYAN, GREEN, YELLOW, RED, ORANGE,}
    if len(args) > 0 {
        return args[0]
    } else {
        if len(colorTaken) == 0 {
            return colors[rand.Intn(len(colors))]
        } else {
            availableColors := []string{}
            for _, c := range colors {
                unique := true
                for _, t := range colorTaken {
                    if c == t {
                        unique = false
                        break
                    }
                }
                if unique {
                    availableColors = append(availableColors, c)
                }
            }
            return availableColors[rand.Intn(len(availableColors))]
        }
    }
}
