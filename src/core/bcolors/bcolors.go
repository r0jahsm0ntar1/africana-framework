//John 3:16

package bcolors

import (
    "fmt"
    "runtime"
    "math/rand"
)

var (
    Endc   = "\033[0m"
    Bold   = "\033[1m"
    Dim    = "\033[2m"
    Italic = "\033[3m"
    Underl = "\033[4m"
    Blink  = "\033[5m"
    Blink2 = "\033[6m"
    Invert = "\033[7m"
    Hidden = "\033[8m"
    Strike = "\033[9m"

    Black   = "\033[30m"
    Red     = "\033[31m"
    Green   = "\033[32m"
    Yellow  = "\033[33m"
    Blue    = "\033[34m"
    Magenta = "\033[35m"
    Cyan    = "\033[36m"
    White   = "\033[37m"

    Grey          = "\033[90m"
    BrightRed     = "\033[91m"
    BrightGreen   = "\033[92m"
    BrightYellow  = "\033[93m"
    BrightBlue    = "\033[94m"
    BrightMagenta = "\033[95m"
    BrightCyan    = "\033[96m"
    BrightWhite   = "\033[97m"

    BackBlack     = "\033[40m"
    BackRed       = "\033[41m"
    BackGreen     = "\033[42m"
    BackYellow    = "\033[43m"
    BackBlue      = "\033[44m"
    BackMagenta   = "\033[45m"
    BackCyan      = "\033[46m"
    BackWhite     = "\033[47m"

    BackGrey          = "\033[100m"
    BackBrightRed     = "\033[101m"
    BackBrightGreen   = "\033[102m"
    BackBrightYellow  = "\033[103m"
    BackBrightBlue    = "\033[104m"
    BackBrightMagenta = "\033[105m"
    BackBrightCyan    = "\033[106m"
    BackBrightWhite   = "\033[107m"

    XTERM_FG    []string
    XTERM_BG    []string
    colorTaken  []string
)

func init() {
    XTERM_FG = make([]string, 256)
    XTERM_BG = make([]string, 256)
    for i := 0; i < 256; i++ {
        XTERM_FG[i] = fmt.Sprintf("\033[38;5;%dm", i)
        XTERM_BG[i] = fmt.Sprintf("\033[48;5;%dm", i)
    }

    if runtime.GOOS == "windows" {
        Black, Red, Green, Yellow, Blue, Magenta, Cyan, White = "", "", "", "", "", "", "", ""
        Endc, Bold, Dim, Italic, Underl, Blink, Blink2, Invert, Hidden, Strike = "", "", "", "", "", "", "", "", "", ""
        BackBlack, BackRed, BackGreen, BackYellow, BackBlue, BackMagenta, BackCyan, BackWhite = "", "", "", "", "", "", "", ""
        Grey, BrightRed, BrightGreen, BrightYellow, BrightBlue, BrightMagenta, BrightCyan, BrightWhite = "", "", "", "", "", "", "", ""
        BackGrey, BackBrightRed, BackBrightGreen, BackBrightYellow, BackBrightBlue, BackBrightMagenta, BackBrightCyan, BackBrightWhite = "", "", "", "", "", "", "", ""

        for i := range XTERM_FG {
            XTERM_FG[i] = ""
            XTERM_BG[i] = ""
        }
    }
}

func Colors(args ...string) string {
    colors := []string{
        BrightRed, BrightGreen, BrightYellow, BrightBlue,
        BrightMagenta, BrightCyan, BrightWhite, Yellow,
        Magenta, Cyan, Green, Blue, Red,
        XTERM_FG[9],
        XTERM_FG[10],
        XTERM_FG[11],
        XTERM_FG[12],
        XTERM_FG[13],
        XTERM_FG[14],
        XTERM_FG[15],
        XTERM_FG[45],
        XTERM_FG[51],
        XTERM_FG[82],
        XTERM_FG[201],
    }

    if len(args) > 0 {
        return args[0]
    }

    availableColors := []string{}
    for _, c := range colors {
        used := false
        for _, t := range colorTaken {
            if c == t {
                used = true
                break
            }
        }
        if !used {
            availableColors = append(availableColors, c)
        }
    }

    if len(availableColors) == 0 {
        colorTaken = []string{}
        availableColors = colors
    }

    color := availableColors[rand.Intn(len(availableColors))]
    colorTaken = append(colorTaken, color)
    return color
}

func XtermColor(index int) string {
    if index < 0 || index > 255 {
        return Endc
    }
    return XTERM_FG[index]
}

func XtermBgColor(index int) string {
    if index < 0 || index > 255 {
        return Endc
    }
    return XTERM_BG[index]
}

func XtermColorRGB(r, g, b int) string {
    if r < 0 || r > 5 || g < 0 || g > 5 || b < 0 || b > 5 {
        return Endc
    }
    index := 16 + (36 * r) + (6 * g) + b
    return XTERM_FG[index]
}

func XtermBgColorRGB(r, g, b int) string {
    if r < 0 || r > 5 || g < 0 || g > 5 || b < 0 || b > 5 {
        return Endc
    }
    index := 16 + (36 * r) + (6 * g) + b
    return XTERM_BG[index]
}

func XtermGrayscale(level int) string {
    if level < 0 || level > 23 {
        return Endc
    }
    return XTERM_FG[232 + level]
}

func XtermBgGrayscale(level int) string {
    if level < 0 || level > 23 {
        return Endc
    }
    return XTERM_BG[232 + level]
}
