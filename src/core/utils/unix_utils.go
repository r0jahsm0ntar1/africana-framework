//go:build !windows
// +build !windows

package utils

import (
    "os"
    "fmt"
    "sort"
    "time"
    "bufio"
    "unsafe"
    "syscall"
    "unicode"
    "strconv"
    "bcolors"
    "strings"
    "runtime"
    "subprocess"
    "path/filepath"
)

// Platform-specific ioctl constants
const (
    // Linux and most Unix-like systems
    ioctl_TCGETS = 0x5401
    ioctl_TCSETS = 0x5402
    
    // BSD systems (these are the actual values)
    ioctl_TIOCGETA = 0x40287408
    ioctl_TIOCSETA = 0x80287409
)

func DisplayPrompt(version string, args ...interface{}) (string, error) {
    var prompt string

    if len(args) == 0 {
        prompt = fmt.Sprintf("%s%s%safr%s%s%s > %s",
            bcolors.Endc,
            bcolors.Underline,
            bcolors.Bold,
            version,
            bcolors.Endc,
            bcolors.BrightGreen,
            bcolors.Endc)
    } else {
        context := ""
        value := Function
        showBrackets := true
        color := "brightred"

        if len(args) >= 1 {
            if ctx, ok := args[0].(string); ok {
                context = ctx
            }
        }

        if len(args) >= 2 {
            if val, ok := args[1].(string); ok {
                value = val
            }
        }

        if len(args) >= 3 {
            if brackets, ok := args[2].(bool); ok {
                showBrackets = brackets
            }
        }

        if showBrackets {
            prompt = fmt.Sprintf("%s%s%safr%s%s %s(%s%ssrc/pentest_%s.fn%s)%s > %s",
                bcolors.Endc,
                bcolors.Underline,
                bcolors.Bold,
                version,
                bcolors.Endc,
                context,
                bcolors.Bold,
                getContextColor(color),
                value,
                bcolors.Endc,
                bcolors.BrightGreen,
                bcolors.Endc)
        } else {
            prompt = fmt.Sprintf("%s%s%safr%s%s %s%s%ssrc/pentest_%s.fn%s%s > %s",
                bcolors.Endc,
                bcolors.Underline,
                bcolors.Bold,
                version,
                bcolors.Endc,
                context,
                bcolors.Bold,
                getContextColor(color),
                value,
                bcolors.Endc,
                bcolors.BrightGreen,
                bcolors.Endc)
        }
    }

    editor := NewLineEditor(prompt)
    return editor.ReadLine()
}

func getContextColor(color string) string {
    colorMap := map[string]string{
        "red":        bcolors.BrightRed,
        "blue":       bcolors.Blue,
        "green":      bcolors.Green,
        "yellow":     bcolors.Yellow,
        "magenta":    bcolors.Magenta,
        "cyan":       bcolors.Cyan,
        "brightred":  bcolors.BrightRed,
        "brightblue": bcolors.BrightBlue,
        "default":    bcolors.BrightRed,
    }

    if selected, exists := colorMap[color]; exists {
        return selected
    }
    return bcolors.BrightRed
}

func SetAvailableCommands(commands []string) {
    AvailableCommands = commands
}

func SetCommandPredictions(predictions map[string][]string) {
    CommandPredictions = predictions
}

func SetCommandWords(words []string) {
    CommandWords = words
}

func NewLineEditor(prompt string) *LineEditor {
    historyFile := subprocess.HistoryFile
    visualLength := calculateVisualLength(prompt)

    editor := &LineEditor{
        history:          []string{},
        historyPos:       0,
        currentLine:      "",
        cursorPos:        0,
        prompt:           prompt,
        promptLen:        visualLength,
        historyFile:      historyFile,
        predictions:      []string{},
        predictionPos:    -1,
        showPredictions:  false,
        initialized:      false,
        predictionDrawn:  false,
        tabPressed:       false,
        predictionLines:  0,
        inPredictionMode: false,
        originalLine:     "",
        originalCursor:   0,
        suggestionMode:   false,
        currentSuggestion: "",
        lastKeyTime:      time.Now(),
        keyBuffer:        "",
        emojiEnabled:     false,
        soundEnabled:     false,
        tabPressTime:     time.Time{},
    }

    editor.loadHistory()
    editor.initialized = true
    return editor
}

func calculateVisualLength(prompt string) int {
    length := 0
    inEscape := false
    for _, r := range prompt {
        if r == '\033' {
            inEscape = true
            continue
        }
        if inEscape {
            if r == 'm' {
                inEscape = false
            }
            continue
        }
        length++
    }
    return length
}

func (le *LineEditor) getTerminalWidth() int {
    if runtime.GOOS == "windows" {
        return 80
    }
    
    type winsize struct {
        Row    uint16
        Col    uint16
        Xpixel uint16
        Ypixel uint16
    }

    ws := &winsize{}
    retCode, _, _ := syscall.Syscall(
        syscall.SYS_IOCTL,
        uintptr(syscall.Stdout),
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(ws)),
    )

    if int(retCode) == -1 {
        return 80
    }

    width := int(ws.Col)
    if width <= 0 {
        return 80
    }
    return width
}

func (le *LineEditor) generatePredictions() {
    le.predictions = []string{}
    le.predictionPos = -1

    currentText := strings.TrimSpace(le.currentLine)
    if currentText == "" {
        le.predictions = AvailableCommands
        return
    }

    parts := strings.Fields(currentText)
    if len(parts) == 0 {
        return
    }

    currentWord := parts[len(parts)-1]
    lowerCurrentWord := strings.ToLower(currentWord)

    var predictions []string
    
    if len(parts) == 1 {
        for _, availableCmd := range AvailableCommands {
            lowerCmd := strings.ToLower(availableCmd)
            if le.isPartialMatch(lowerCurrentWord, lowerCmd) {
                predictions = append(predictions, availableCmd)
            }
        }
    } else {
        cmd := strings.ToLower(parts[0])
        
        switch cmd {
        case "set":
            predictions = le.getSetPredictions(parts, currentWord)
        case "show":
            predictions = le.getShowPredictions(currentWord)
        case "run":
            predictions = le.getRunPredictions(currentWord)
        case "use":
            predictions = le.getUsePredictions(currentWord)
        case "scan":
            predictions = le.getScanPredictions(currentWord)
        default:
            if subcommands, exists := CommandPredictions[cmd]; exists {
                for _, subcmd := range subcommands {
                    if le.isPartialMatch(currentWord, subcmd) {
                        predictions = append(predictions, subcmd)
                    }
                }
            } else {
                for _, val := range CommonValues {
                    if le.isPartialMatch(currentWord, val) {
                        predictions = append(predictions, val)
                    }
                }
            }
        }
    }

    le.predictions = le.sortPredictionsWithPriority(predictions, lowerCurrentWord)

    if len(le.predictions) > 0 {
        le.showPredictions = true
        le.predictionPos = 0
    }
}

func (le *LineEditor) sortPredictionsWithPriority(predictions []string, currentWord string) []string {
    if currentWord == "" {
        sort.Strings(predictions)
        return predictions
    }
    
    var prefixMatches []string
    var otherMatches []string
    
    for _, pred := range predictions {
        lowerPred := strings.ToLower(pred)
        if strings.HasPrefix(lowerPred, currentWord) {
            prefixMatches = append(prefixMatches, pred)
        } else {
            otherMatches = append(otherMatches, pred)
        }
    }

    sort.Strings(prefixMatches)
    sort.Strings(otherMatches)

    return append(prefixMatches, otherMatches...)
}

func (le *LineEditor) getSetPredictions(parts []string, currentWord string) []string {
    var predictions []string
    lowerCurrentWord := strings.ToLower(currentWord)

    if len(parts) == 2 {
        setOptions := []string{
            "lhost", "lport", "rhost", "target", "protocol", "function",
            "module", "build", "output", "iface", "gateway", "mode",
            "listener", "innericon", "outericon", "obfuscator", "proxy",
        }

        for _, option := range setOptions {
            if le.isPartialMatch(currentWord, option) {
                predictions = append(predictions, option)
            }
        }
    } else if len(parts) == 3 {
        option := strings.ToLower(parts[1])
        predictions = le.getValuePredictions(option, currentWord)
    }
    
    return le.sortPredictionsWithPriority(predictions, lowerCurrentWord)
}

func (le *LineEditor) getShowPredictions(currentWord string) []string {
    showOptions := []string{
        "options", "config", "settings", "status", "version",
        "modules", "functions", "targets", "sessions", "history",
    }

    var predictions []string
    lowerCurrentWord := strings.ToLower(currentWord)
    
    for _, option := range showOptions {
        if le.isPartialMatch(currentWord, option) {
            predictions = append(predictions, option)
        }
    }
    return le.sortPredictionsWithPriority(predictions, lowerCurrentWord)
}

func (le *LineEditor) getRunPredictions(currentWord string) []string {
    runOptions := []string{
        "scan", "attack", "exploit", "recon", "bruteforce", "module",
        "network-scan", "port-scan", "vulnerability-scan", "dos-attack",
    }

    var predictions []string
    lowerCurrentWord := strings.ToLower(currentWord)
    
    for _, option := range runOptions {
        if le.isPartialMatch(currentWord, option) {
            predictions = append(predictions, option)
        }
    }
    return le.sortPredictionsWithPriority(predictions, lowerCurrentWord)
}

func (le *LineEditor) getUsePredictions(currentWord string) []string {
    useOptions := []string{
        "scanner/network", "scanner/ports", "scanner/vulnerabilities",
        "exploit/windows", "exploit/linux", "exploit/web",
        "payload/generator", "listener/http", "listener/tcp",
        "post/multi", "auxiliary/scanner", "auxiliary/gather",
    }

    var predictions []string
    for _, option := range useOptions {
        if le.isPartialMatch(currentWord, option) {
            predictions = append(predictions, option)
        }
    }
    return predictions
}

func (le *LineEditor) getScanPredictions(currentWord string) []string {
    scanOptions := []string{
        "network", "ports", "vulnerabilities", "services", "os",
        "top-ports", "full", "quick", "comprehensive",
    }

    var predictions []string
    for _, option := range scanOptions {
        if le.isPartialMatch(currentWord, option) {
            predictions = append(predictions, option)
        }
    }
    return predictions
}

func (le *LineEditor) getValuePredictions(option, currentWord string) []string {
    valueMap := map[string][]string{
        "protocol":  {"tcp", "udp", "http", "https"},
        "mode":      {"listen", "attack", "scan", "advanced", "simple"},
        "function":  {"scanner", "exploit", "payload", "listener"},
        "listener":  {"http", "https", "tcp", "udp", "multi"},
    }

    if values, exists := valueMap[option]; exists {
        var predictions []string
        for _, value := range values {
            if le.isPartialMatch(currentWord, value) {
                predictions = append(predictions, value)
            }
        }
        return predictions
    }

    return []string{}
}

func (le *LineEditor) showInlinePredictions() {
    if len(le.predictions) == 0 || !le.showPredictions {
        return
    }

    if !le.inPredictionMode {
        le.originalLine = le.currentLine
        le.originalCursor = le.cursorPos
        le.inPredictionMode = true
    }

    le.redrawWithPredictions()
}

func (le *LineEditor) redrawWithPredictions() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)

    le.showSmartSuggestion()

    if le.showPredictions && len(le.predictions) > 0 {
        terminalWidth := le.getTerminalWidth()
        currentLineLength := le.promptLen + len(le.currentLine)
        availableWidth := terminalWidth - currentLineLength - 10

        predictionText := " " + bcolors.Yellow + "[" + bcolors.Endc

        startIdx := 0
        maxPredictions := len(le.predictions)

        if len(le.predictions) > 5 {
            visibleCount := 5
            if availableWidth < 50 {
                visibleCount = 3
            }

            if le.predictionPos >= visibleCount {
                startIdx = le.predictionPos - visibleCount + 1
                if startIdx+visibleCount > len(le.predictions) {
                    startIdx = len(le.predictions) - visibleCount
                }
            }

            endIdx := startIdx + visibleCount
            if endIdx > len(le.predictions) {
                endIdx = len(le.predictions)
            }

            maxPredictions = endIdx - startIdx

            if startIdx > 0 {
                predictionText += bcolors.Magenta + "←" + bcolors.Endc
            }
        }

        for i := 0; i < maxPredictions; i++ {
            actualIdx := startIdx + i
            if actualIdx >= len(le.predictions) {
                break
            }

            pred := le.predictions[actualIdx]
            if actualIdx == le.predictionPos {
                predictionText += bcolors.Green + "›" + pred + bcolors.Endc
            } else {
                predictionText += bcolors.Cyan + pred + bcolors.Endc
            }

            if i < maxPredictions-1 {
                predictionText += bcolors.Yellow + "|" + bcolors.Endc
            }
        }

        if startIdx+maxPredictions < len(le.predictions) {
            remaining := len(le.predictions) - (startIdx + maxPredictions)
            predictionText += bcolors.Yellow + "|" + bcolors.Endc
            predictionText += bcolors.Magenta + "+" + fmt.Sprintf("%d", remaining) + bcolors.Endc
        }

        predictionText += bcolors.Yellow + "]" + bcolors.Endc
        fmt.Print(predictionText)
    }

    fmt.Printf("\r\033[%dC", le.promptLen+le.cursorPos)
}

func (le *LineEditor) applyPrediction() {
    if le.predictionPos >= 0 && le.predictionPos < len(le.predictions) {
        selected := le.predictions[le.predictionPos]

        parts := strings.Fields(le.originalLine)

        if len(parts) == 0 {
            le.currentLine = selected + " "
            le.cursorPos = len(le.currentLine)
        } else if len(parts) == 1 {
            currentWord := parts[0]
            if le.isPartialMatch(currentWord, selected) {
                le.currentLine = selected
                if le.isCommandThatTakesArgs(selected) {
                    le.currentLine += " "
                }
            } else {
                le.currentLine = le.originalLine + " " + selected
                if le.isCommandThatTakesArgs(selected) {
                    le.currentLine += " "
                }
            }
            le.cursorPos = len(le.currentLine)
        } else {
            lastWord := parts[len(parts)-1]
            if le.isPartialMatch(lastWord, selected) {
                newParts := parts[:len(parts)-1]
                newParts = append(newParts, selected)
                le.currentLine = strings.Join(newParts, " ")
                if le.isCommandThatTakesArgs(selected) {
                    le.currentLine += " "
                }
            } else {
                le.currentLine = le.originalLine + " " + selected
                if le.isCommandThatTakesArgs(selected) {
                    le.currentLine += " "
                }
            }
            le.cursorPos = len(le.currentLine)
        }
    }

    le.exitPredictionMode()
}

func (le *LineEditor) isPartialMatch(partial, full string) bool {
    partialLower := strings.ToLower(partial)
    fullLower := strings.ToLower(full)

    if strings.HasPrefix(fullLower, partialLower) {
        return true
    }

    if Levenshtein(partialLower, fullLower) <= 2 {
        return true
    }

    abbreviations := map[string]string{
        "mo": "module", "mod": "module", "conf": "config",
        "opt": "options", "set": "settings", "stat": "status",
        "ver": "version", "inst": "install", "uninst": "uninstall",
    }

    if expanded, exists := abbreviations[partialLower]; exists {
        return expanded == fullLower
    }

    return false
}

func (le *LineEditor) isCommandThatTakesArgs(cmd string) bool {
    commandsWithArgs := map[string]bool{
        "set": true, "show": true, "run": true, "scan": true,
        "use": true, "config": true, "start": true, "stop": true,
    }

    return commandsWithArgs[strings.ToLower(cmd)]
}

func (le *LineEditor) exitPredictionMode() {
    le.showPredictions = false
    le.inPredictionMode = false
    le.predictions = []string{}
    le.predictionPos = -1
    le.predictionDrawn = false
    le.redraw()
}

func (le *LineEditor) nextPrediction() {
    if len(le.predictions) == 0 {
        return
    }

    le.predictionPos = (le.predictionPos + 1) % len(le.predictions)
    le.showInlinePredictions()
}

func (le *LineEditor) prevPrediction() {
    if len(le.predictions) == 0 {
        return
    }

    le.predictionPos--
    if le.predictionPos < 0 {
        le.predictionPos = len(le.predictions) - 1
    }
    le.showInlinePredictions()
}

func (le *LineEditor) firstPrediction() {
    if len(le.predictions) == 0 {
        return
    }

    le.predictionPos = 0
    le.showInlinePredictions()
}

func (le *LineEditor) lastPrediction() {
    if len(le.predictions) == 0 {
        return
    }

    le.predictionPos = len(le.predictions) - 1
    le.showInlinePredictions()
}

func (le *LineEditor) getSmartSuggestion() string {
    text := strings.TrimSpace(le.currentLine)
    if text == "" {
        return ""
    }

    parts := strings.Fields(text)
    if len(parts) == 0 {
        return ""
    }

    lastWord := parts[len(parts)-1]
    lowerLastWord := strings.ToLower(lastWord)

    if len(parts) == 1 {
        return le.getCommandSuggestion(lowerLastWord)
    } else if len(parts) >= 2 {
        return le.getContextSuggestion(parts)
    }

    return ""
}

func (le *LineEditor) getCommandSuggestion(cmd string) string {
    if suggestion, exists := CommandSuggestions[cmd]; exists {
        return suggestion
    }

    for fullCmd, suggestion := range CommandSuggestions {
        if strings.HasPrefix(fullCmd, cmd) && cmd != fullCmd {
            return fullCmd[len(cmd):] + suggestion
        }
    }

    return ""
}

func (le *LineEditor) getContextSuggestion(parts []string) string {
    cmd := strings.ToLower(parts[0])
    lastWord := strings.ToLower(parts[len(parts)-1])

    if le.looksLikeUserIsTypingValue(lastWord) {
        return ""
    }

    switch cmd {
    case "set":
        return le.getSetSuggestion(parts, lastWord)
    case "show":
        return le.getShowSuggestion(parts, lastWord)
    case "run":
        return le.getRunSuggestion(parts, lastWord)
    case "use":
        return le.getUseSuggestion(parts, lastWord)
    case "scan":
        return le.getScanSuggestion(parts, lastWord)
    }

    return ""
}

func (le *LineEditor) getSetSuggestion(parts []string, lastWord string) string {
    if len(parts) == 2 {
        if suggestions, exists := PartialSuggestions["set"]; exists {
            if suggestion, exists := suggestions[lastWord]; exists {
                return suggestion
            }
        }

        return " <value>"
    } else if len(parts) == 3 {
        option := strings.ToLower(parts[1])
        if suggestion, exists := ValueSuggestions[option]; exists {
            if !le.looksLikeUserIsTypingValue(lastWord) {
                return suggestion
            }
        }

        return " <value>"
    }

    return ""
}

func (le *LineEditor) getShowSuggestion(parts []string, lastWord string) string {
    if len(parts) == 2 {
        if suggestions, exists := PartialSuggestions["show"]; exists {
            if suggestion, exists := suggestions[lastWord]; exists {
                return suggestion
            }
        }
    }

    return ""
}

func (le *LineEditor) getRunSuggestion(parts []string, lastWord string) string {
    if len(parts) == 2 {
        if suggestions, exists := PartialSuggestions["run"]; exists {
            if suggestion, exists := suggestions[lastWord]; exists {
                return suggestion
            }
        }
    }

    return ""
}

func (le *LineEditor) getUseSuggestion(parts []string, lastWord string) string {
    if len(parts) == 2 {
        if suggestions, exists := PartialSuggestions["use"]; exists {
            if suggestion, exists := suggestions[lastWord]; exists {
                return suggestion
            }
        }
    }

    return ""
}

func (le *LineEditor) getScanSuggestion(parts []string, lastWord string) string {
    if len(parts) == 2 {
        if suggestions, exists := PartialSuggestions["scan"]; exists {
            if suggestion, exists := suggestions[lastWord]; exists {
                return suggestion
            }
        }
    }

    return ""
}

func (le *LineEditor) looksLikeUserIsTypingValue(word string) bool {
    if word == "" {
        return false
    }

    for _, r := range word {
        if unicode.IsDigit(r) {
            return true
        }
    }

    commonValuePrefixes := []string{"tru", "fals", "en", "dis", "on", "of", "ye", "no"}
    for _, prefix := range commonValuePrefixes {
        if strings.HasPrefix(strings.ToLower(word), prefix) {
            return true
        }
    }

    return false
}

func (le *LineEditor) showSmartSuggestion() {
    suggestion := le.getSmartSuggestion()
    if suggestion == "" {
        return
    }

    currentText := le.currentLine
    if len(currentText) > 0 {
        parts := strings.Fields(currentText)
        if len(parts) > 0 {
            lastWord := parts[len(parts)-1]
            if strings.Contains(suggestion, "<") && strings.Contains(suggestion, ">") {
                if strings.HasPrefix(suggestion, " "+lastWord) || le.currentLineEndsWithPartialTemplate(suggestion) {
                    return
                }
            } else if strings.HasPrefix(suggestion, " "+lastWord) {
                return
            }
        }
    }

    fmt.Printf("%s%s%s", bcolors.Background, bcolors.BrightBlue, suggestion)
    fmt.Printf("%s", bcolors.Endc)

    fmt.Printf("\r\033[%dC", le.promptLen+le.cursorPos)
}

func (le *LineEditor) currentLineEndsWithPartialTemplate(suggestion string) bool {
    if !strings.Contains(suggestion, "<") || !strings.Contains(suggestion, ">") {
        return false
    }

    templateStart := strings.Index(suggestion, "<")
    templateEnd := strings.Index(suggestion, ">")
    if templateStart == -1 || templateEnd == -1 {
        return false
    }

    template := suggestion[templateStart : templateEnd+1]
    currentText := le.currentLine
    if len(currentText) > 0 {
        return strings.Contains(currentText, "<") ||
            (len(template) > 2 && strings.Contains(currentText, template[1:2]))
    }

    return false
}

func (le *LineEditor) handleTabCompletion() {
    le.generatePredictions()

    now := time.Now()
    isDoubleTab := !le.tabPressTime.IsZero() && now.Sub(le.tabPressTime) < 500*time.Millisecond

    if isDoubleTab {
        le.showContextHelp()
    } else if len(le.predictions) > 0 {
        if le.inPredictionMode {
            le.nextPrediction()
        } else {
            le.enterPredictionMode()
        }
    }

    le.tabPressTime = now
}

func (le *LineEditor) showContextHelp() {
    parts := strings.Fields(le.currentLine)

    if len(parts) == 0 {
        le.showAllCommands()
        return
    }

    command := strings.ToLower(parts[0])

    switch command {
    case "set":
        le.showSetOptions()
    case "show":
        le.showShowOptions()
    case "run":
        le.showRunOptions()
    case "use":
        le.showUseOptions()
    case "scan":
        le.showScanOptions()
    case "start", "stop", "config", "help":
        if subcommands, exists := CommandPredictions[command]; exists {
            le.showCommandSpecificHelp(command, subcommands)
        } else {
            le.showAllCommands()
        }
    default:
        le.showAllCommands()
    }
}

func (le *LineEditor) showSetOptions() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)
    fmt.Printf("\n%s%s            SET Command Options%s\n", bcolors.Green, bcolors.Bold, bcolors.Endc)
    fmt.Printf("%s%sUsage: set <option> <value>%s\n\n", bcolors.Yellow, bcolors.Bold, bcolors.Endc)

    categories := map[string][]string{
        "Target Configuration": {
            "lhost", "lport", "rhost", "target", "protocol", "gateway",
        },
        "Build Options": {
            "build", "buildname", "innericon", "outericon", "obfuscator",
        },
        "Module Settings": {
            "function", "module", "listener", "mode", "spoofer",
        },
        "Network Settings": {
            "iface", "proxy", "fakedns", "venvname",
        },
        "Output Settings": {
            "output", "recondir", "toolsdir",
        },
    }

    for category, options := range categories {
        fmt.Printf("  %s%s%s\n", bcolors.Cyan, category, bcolors.Endc)
        fmt.Printf("  ")
        for i, option := range options {
            color := bcolors.Green
            if i%2 == 1 {
                color = bcolors.BrightGreen
            }
            fmt.Printf("%s%-15s%s ", color, option, bcolors.Endc)
            if (i+1)%3 == 0 && i != len(options)-1 {
                fmt.Printf("\n  ")
            }
        }
        fmt.Printf("\n\n")
    }

    fmt.Printf("%s%sExamples:%s\n", bcolors.Magenta, bcolors.Bold, bcolors.Endc)
    fmt.Printf("  set lhost 192.168.1.100    set lport 4444    set function scanner\n")
    fmt.Printf("  set protocol tcp           set output /path  set mode advanced\n\n")

    le.redrawPrompt()
}

func (le *LineEditor) showShowOptions() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)
    fmt.Printf("\n%s%s            SHOW Command Options%s\n", bcolors.Green, bcolors.Bold, bcolors.Endc)
    fmt.Printf("%s%sUsage: show <category>%s\n\n", bcolors.Yellow, bcolors.Bold, bcolors.Endc)

    options := []string{
        "options", "config", "settings", "status", "version",
        "modules", "functions", "targets", "sessions", "history",
    }

    fmt.Printf("  Available categories:\n")
    fmt.Printf("  ")
    for i, option := range options {
        color := bcolors.Blue
        if i%2 == 1 {
            color = bcolors.BrightBlue
        }
        fmt.Printf("%s%-12s%s ", color, option, bcolors.Endc)
        if (i+1)%4 == 0 && i != len(options)-1 {
            fmt.Printf("\n  ")
        }
    }
    fmt.Printf("\n\n")

    fmt.Printf("%s%sExamples:%s\n", bcolors.Magenta, bcolors.Bold, bcolors.Endc)
    fmt.Printf("  show options     show config     show status\n")
    fmt.Printf("  show modules     show version    show targets\n\n")

    le.redrawPrompt()
}

func (le *LineEditor) showRunOptions() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)
    fmt.Printf("\n%s%s            RUN Command Options%s\n", bcolors.Green, bcolors.Bold, bcolors.Endc)
    fmt.Printf("%s%sUsage: run <module/scan>%s\n\n", bcolors.Yellow, bcolors.Bold, bcolors.Endc)

    categories := map[string][]string{
        "Scans": {
            "scan", "network-scan", "port-scan", "vulnerability-scan",
        },
        "Attacks": {
            "attack", "exploit", "bruteforce", "dos-attack",
        },
        "Modules": {
            "module", "function", "tool",
        },
    }

    for category, options := range categories {
        fmt.Printf("  %s%s%s\n", bcolors.Cyan, category, bcolors.Endc)
        fmt.Printf("  ")
        for i, option := range options {
            color := bcolors.Green
            if i%2 == 1 {
                color = bcolors.BrightGreen
            }
            fmt.Printf("%s%-20s%s ", color, option, bcolors.Endc)
            if (i+1)%2 == 0 && i != len(options)-1 {
                fmt.Printf("\n  ")
            }
        }
        fmt.Printf("\n\n")
    }

    le.redrawPrompt()
}

func (le *LineEditor) showUseOptions() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)
    fmt.Printf("\n%s%s            USE Command Options%s\n", bcolors.Green, bcolors.Bold, bcolors.Endc)
    fmt.Printf("%s%sUsage: use <module>%s\n\n", bcolors.Yellow, bcolors.Bold, bcolors.Endc)

    moduleCategories := map[string][]string{
        "Reconnaissance": {
            "scanner/network", "scanner/ports", "scanner/vulnerabilities",
        },
        "Exploitation": {
            "exploit/windows", "exploit/linux", "exploit/web",
        },
        "Post-Exploitation": {
            "post/multi", "post/windows", "post/linux",
        },
        "Auxiliary": {
            "auxiliary/scanner", "auxiliary/admin", "auxiliary/gather",
        },
    }

    for category, modules := range moduleCategories {
        fmt.Printf("  %s%s%s\n", bcolors.Cyan, category, bcolors.Endc)
        for _, module := range modules {
            fmt.Printf("  %s%s%s\n", bcolors.BrightGreen, module, bcolors.Endc)
        }
        fmt.Printf("\n")
    }

    le.redrawPrompt()
}

func (le *LineEditor) showScanOptions() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)
    fmt.Printf("\n%s%s            SCAN Command Options%s\n", bcolors.Green, bcolors.Bold, bcolors.Endc)
    fmt.Printf("%s%sUsage: scan <type> [target]%s\n\n", bcolors.Yellow, bcolors.Bold, bcolors.Endc)

    scanTypes := map[string]string{
        "network":         "Scan entire network range",
        "ports":           "Scan ports on specific target",
        "vulnerabilities": "Scan for known vulnerabilities",
        "services":        "Identify running services",
        "os":              "Operating system detection",
        "top-ports":       "Scan most common ports",
        "full":            "Comprehensive scan",
    }

    fmt.Printf("  %s%-20s %s%s\n", bcolors.Cyan, "Scan Type", "Description", bcolors.Endc)
    fmt.Printf("  %s%s%s\n", bcolors.Yellow, strings.Repeat("-", 50), bcolors.Endc)

    for scanType, description := range scanTypes {
        fmt.Printf("  %s%-20s%s %s\n",
            bcolors.Green, scanType, bcolors.Endc,
            bcolors.BrightBlue+description+bcolors.Endc)
    }

    fmt.Printf("\n%s%sExamples:%s\n", bcolors.Magenta, bcolors.Bold, bcolors.Endc)
    fmt.Printf("  scan network 192.168.1.0/24\n")
    fmt.Printf("  scan ports 192.168.1.100\n")
    fmt.Printf("  scan vulnerabilities 192.168.1.100\n\n")

    le.redrawPrompt()
}

func (le *LineEditor) showCommandSpecificHelp(command string, subcommands []string) {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)
    fmt.Printf("\n%s%s            %s Command Options%s\n", bcolors.Green, bcolors.Bold, strings.ToUpper(command), bcolors.Endc)

    usage := ""
    if sugg, exists := CommandSuggestions[command]; exists {
        usage = "Usage: " + command + sugg
    } else {
        usage = "Usage: " + command + " <subcommand>"
    }

    fmt.Printf("%s%s%s\n\n", bcolors.Yellow, bcolors.Bold, usage)

    fmt.Printf("  Available subcommands:\n")
    fmt.Printf("  ")
    for i, subcmd := range subcommands {
        color := bcolors.Blue
        if i%2 == 1 {
            color = bcolors.BrightBlue
        }
        fmt.Printf("%s%-15s%s ", color, subcmd, bcolors.Endc)
        if (i+1)%4 == 0 && i != len(subcommands)-1 {
            fmt.Printf("\n  ")
        }
    }
    fmt.Printf("\n\n")

    le.redrawPrompt()
}

func (le *LineEditor) redrawPrompt() {
    fmt.Print(le.prompt + le.currentLine)
    if le.cursorPos < len(le.currentLine) {
        fmt.Printf("\r\033[%dC", le.promptLen+le.cursorPos)
    }
}

func (le *LineEditor) enterPredictionMode() {
    le.showPredictions = true
    le.predictionPos = 0
    le.inPredictionMode = true
    le.originalLine = le.currentLine
    le.originalCursor = le.cursorPos
    le.showInlinePredictions()
}

func (le *LineEditor) showDoubleTabHint() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)
    fmt.Printf("%s%s(Tab again to show all commands)%s", bcolors.Yellow, bcolors.Bold, bcolors.Endc)
    fmt.Printf("\r\033[%dC", le.promptLen+le.cursorPos)
}

func (le *LineEditor) showAllCommands() {
    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)

    fmt.Printf("\n\n%s%sCommands%s ", bcolors.Blue, bcolors.Bold, bcolors.Endc)
    fmt.Printf("%s(%d categories, %d total)%s\n\n", bcolors.DarkGray, len(CommandCategories), le.countTotalCommands(), bcolors.Endc)

    categories := []struct {
        name     string
        indent   int
        commands []string
        spaces   []int
    }{
        {
            "Configuration", 0,
            []string{"set", "config", "get"},
            []int{3, 1, 0},
        },
        {
            "Information", 2,
            []string{"help", "show", "version", "status", "list"},
            []int{2, 3, 1, 4, 0},
        },
        {
            "Execution", 4,
            []string{"run", "scan", "start", "stop", "restart", "analyze", "build", "use"},
            []int{3, 3, 3, 6, 1, 1, 1, 0},
        },
        {
            "System", 7,
            []string{"exit", "update", "install", "uninstall", "log"},
            []int{2, 1, 1, 1, 0},
        },
    }

    maxCommandStart := 0
    for _, cat := range categories {
        startPos := cat.indent + len(cat.name) + len(fmt.Sprintf(" (%d): ", len(cat.commands)))
        if startPos > maxCommandStart {
            maxCommandStart = startPos
        }
    }

    for _, cat := range categories {
        currentStart := cat.indent + len(cat.name) + len(fmt.Sprintf(" (%d): ", len(cat.commands)))
        fmt.Printf("%s", strings.Repeat(" ", cat.indent))

        fmt.Printf("%s%s%s %s(%d):%s ", bcolors.Yellow, bcolors.Bold, cat.name, bcolors.DarkGray, len(cat.commands), bcolors.Endc)

        padding := maxCommandStart - currentStart
        if padding > 0 {
            fmt.Printf("%s", strings.Repeat(" ", padding))
        }

        for i, cmd := range cat.commands {
            fmt.Printf("%s%s%s", bcolors.BrightGreen, cmd, bcolors.Endc)

            if cat.spaces[i] > 0 {
                fmt.Printf("%s", strings.Repeat(" ", cat.spaces[i]))
            }

            if i < len(cat.commands)-1 {
                fmt.Printf("%s•%s ", bcolors.DarkGray, bcolors.Endc)
            }
        }
        fmt.Printf("\n")
    }

    fmt.Printf("\n%s%sExamples%s\n", bcolors.Blue, bcolors.Bold, bcolors.Endc)
    fmt.Printf("  ")

    examples := []string{
        "set lhost " + LHost,
        "show options", 
        "run port-scan",
        "use scanner/network",
    }

    for i, example := range examples {
        if i > 0 {
            fmt.Printf("%s  %s", bcolors.DarkGray, bcolors.Endc)
        }
        fmt.Printf("%s↳%s %s%s%s", bcolors.DarkGray, bcolors.Endc, bcolors.BrightGreen, example, bcolors.Endc)
    }

    fmt.Printf("\n\n")
    le.redrawPrompt()
}

func (le *LineEditor) countTotalCommands() int {
    total := 0
    for _, commands := range CommandCategories {
        total += len(commands)
    }
    return total
}

func (le *LineEditor) insertChar(c rune) {
    if le.inPredictionMode {
        le.exitPredictionMode()
    }
    originalLine := le.currentLine
    originalCursor := le.cursorPos

    defer func() {
        if r := recover(); r != nil {
            le.currentLine = originalLine
            le.cursorPos = originalCursor
            le.redraw()
            fmt.Printf("\n%s[Recovered from error]%s", bcolors.BrightRed, bcolors.Endc)
        }
    }()

    if le.cursorPos == len(le.currentLine) {
        le.currentLine += string(c)
        fmt.Print(string(c))
        le.cursorPos++
        le.autoInsertSpace()
    } else {
        if le.cursorPos >= 0 && le.cursorPos <= len(le.currentLine) {
            if le.cursorPos == len(le.currentLine) {
                le.currentLine += string(c)
            } else {
                le.currentLine = le.currentLine[:le.cursorPos] + string(c) + le.currentLine[le.cursorPos:]
            }
            le.cursorPos++
            le.redraw()
            le.autoInsertSpace()
        }
    }
}

func (le *LineEditor) autoInsertSpace() {
    if le.cursorPos != len(le.currentLine) {
        return
    }

    currentText := le.currentLine
    if len(currentText) < 3 {
        return
    }

    lastSpaceIndex := strings.LastIndex(currentText, " ")
    if lastSpaceIndex == -1 {
        return
    }

    if lastSpaceIndex+1 >= len(currentText) {
        return
    }

    lastPart := currentText[lastSpaceIndex+1:]
    if len(lastPart) < 2 {
        return
    }

    parts := strings.Fields(currentText)
    if len(parts) < 2 {
        return
    }

    command := strings.ToLower(parts[0])

    for _, commandWord := range CommandWords {
        commandLen := len(commandWord)

        if commandLen < 1 || commandLen >= len(lastPart) {
            continue
        }

        if !strings.HasPrefix(strings.ToLower(lastPart), strings.ToLower(commandWord)) {
            continue
        }

        remaining := lastPart[commandLen:]
        if len(remaining) == 0 {
            continue
        }

        if le.shouldInsertSpaceSimple(command, commandWord, remaining) {
            textBeforeLastPart := currentText[:lastSpaceIndex+1]

            if len(textBeforeLastPart)+len(commandWord)+1+len(remaining) <= len(currentText)+10 {
                newText := textBeforeLastPart + commandWord + " " + remaining

                if len(newText) > len(currentText) && len(newText) <= len(currentText)+20 {
                    le.currentLine = newText
                    le.cursorPos = len(le.currentLine)
                    le.redraw()
                    return
                }
            }
        }
    }
}

func (le *LineEditor) shouldInsertSpaceSimple(command, commandWord, remaining string) bool {
    if command == "set" {
        blacklist := map[string]bool{

        }

        if blacklist[commandWord] {
            return false
        }

        whitelist := map[string]bool{
            "script": true, "name": true, "module": true, "modules": true, "target": true, "targets": true, "lhost": true, "lport": true, "rhost": true, "rport": true, "option": true, "options": true, "value": true, "values": true, "proxy": true, "mode": true, "protocol": true, "threads": true, "function": true, "func": true, "funcs": true, "functions": true, "timeout": true, "output": true, "gateway": true, "verbose": true,
        }

        return whitelist[commandWord]
    }

    return false
}

func (le *LineEditor) autoFixCommonPatternsSafe(lastPart string, lastSpaceIndex int) {
    currentText := le.currentLine
    textBeforeLastPart := currentText[:lastSpaceIndex+1]

    for _, commandWord := range CommandWords {
        commandLen := len(commandWord)

        if commandLen < 1 || commandLen >= len(lastPart) {
            continue
        }

        if !strings.HasPrefix(strings.ToLower(lastPart), strings.ToLower(commandWord)) {
            continue
        }

        remaining := lastPart[commandLen:]
        if len(remaining) == 0 {
            continue
        }

        if le.looksLikeIPAddress(remaining) || le.looksLikePort(remaining) || le.looksLikeCommonValue(remaining) {
            if len(textBeforeLastPart)+len(commandWord)+1+len(remaining) <= len(currentText)+10 {
                newText := textBeforeLastPart + commandWord + " " + remaining
                if len(newText) > len(currentText) {
                    le.currentLine = newText
                    le.cursorPos = len(le.currentLine)
                    le.redraw()
                    return
                }
            }
        }
    }
}

func (le *LineEditor) isCommandWord(word string) bool {
    lowerWord := strings.ToLower(word)
    for _, cmd := range CommandWords {
        if lowerWord == cmd {
            return true
        }
    }
    return false
}

func (le *LineEditor) autoFixCommonPatterns(lastPart string) {
    for _, commandWord := range CommandWords {
        if len(lastPart) > len(commandWord) && strings.HasPrefix(strings.ToLower(lastPart), commandWord) {
            remaining := lastPart[len(commandWord):]

            if le.looksLikeIPAddress(remaining) || le.looksLikePort(remaining) || le.looksLikeCommonValue(remaining) {
                newLastPart := commandWord + " " + remaining
                le.currentLine = le.currentLine[:len(le.currentLine)-len(lastPart)] + newLastPart
                le.cursorPos = len(le.currentLine)
                le.redraw()
                return
            }
        }
    }
}

func (le *LineEditor) looksLikeIPAddress(s string) bool {
    if len(s) < 7 {
        return false
    }

    parts := strings.Split(s, ".")
    if len(parts) != 4 {
        return false
    }

    for _, part := range parts {
        if len(part) == 0 || len(part) > 3 {
            return false
        }
        for _, r := range part {
            if !unicode.IsDigit(r) {
                return false
            }
        }
    }
    return true
}

func (le *LineEditor) looksLikePort(s string) bool {
    if len(s) == 0 || len(s) > 5 {
        return false
    }

    for _, r := range s {
        if !unicode.IsDigit(r) {
            return false
        }
    }

    if port, err := strconv.Atoi(s); err == nil {
        return port > 0 && port <= 65535
    }
    return false
}

func (le *LineEditor) looksLikeCommonValue(s string) bool {
    if len(s) == 0 {
        return false
    }

    commonValues := []string{"true", "false", "on", "off", "enable", "disable", "yes", "no"}
    for _, val := range commonValues {
        if len(s) >= len(val) && strings.HasPrefix(strings.ToLower(s), val) {
            return true
        }
    }
    return false
}

func (le *LineEditor) backspace() {
    if le.inPredictionMode {
        le.exitPredictionMode()
    }

    originalLine := le.currentLine
    originalCursor := le.cursorPos

    defer func() {
        if r := recover(); r != nil {
            le.currentLine = originalLine
            le.cursorPos = originalCursor
            le.redraw()
            fmt.Printf("\n%s[Recovered from backspace error]%s", bcolors.BrightRed, bcolors.Endc)
        }
    }()

    if le.cursorPos > 0 && len(le.currentLine) > 0 {
        if le.cursorPos == len(le.currentLine) {
            if len(le.currentLine) > 0 {
                le.currentLine = le.currentLine[:len(le.currentLine)-1]
                fmt.Print("\b \b")
                le.cursorPos--
            }
        } else {
            if le.cursorPos-1 >= 0 && le.cursorPos-1 < len(le.currentLine) && le.cursorPos <= len(le.currentLine) {
                le.currentLine = le.currentLine[:le.cursorPos-1] + le.currentLine[le.cursorPos:]
                le.cursorPos--
                le.redraw()
            }
        }
    }
}

func (le *LineEditor) deleteChar() {
    if le.cursorPos < len(le.currentLine) {
        le.currentLine = le.currentLine[:le.cursorPos] + le.currentLine[le.cursorPos+1:]
        le.redraw()
    }
}

func (le *LineEditor) moveLeft() {
    if le.cursorPos > 0 {
        le.cursorPos--
        le.redraw()
    }
}

func (le *LineEditor) moveRight() {
    if le.cursorPos < len(le.currentLine) {
        le.cursorPos++
        le.redraw()
    }
}

func (le *LineEditor) moveHome() {
    if le.cursorPos != 0 {
        le.cursorPos = 0
        le.redraw()
    }
}

func (le *LineEditor) moveEnd() {
    if le.cursorPos != len(le.currentLine) {
        le.cursorPos = len(le.currentLine)
        le.redraw()
    }
}

func (le *LineEditor) loadHistory() {
    dir := filepath.Dir(le.historyFile)
    os.MkdirAll(dir, 0755)

    file, err := os.Open(le.historyFile)
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" {
            le.history = append(le.history, line)
        }
    }
    le.historyPos = len(le.history)
}

func (le *LineEditor) saveHistory() {
    dir := filepath.Dir(le.historyFile)
    os.MkdirAll(dir, 0755)

    file, err := os.OpenFile(le.historyFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return
    }
    defer file.Close()

    if len(le.history) > 0 {
        lastCommand := le.history[len(le.history)-1]
        file.WriteString(lastCommand + "\n")
    }
}

func (le *LineEditor) addToHistory(line string) {
    if strings.TrimSpace(line) != "" {
        le.history = append(le.history, line)
        le.saveHistory()
    }
    le.historyPos = len(le.history)
}

func (le *LineEditor) clearLine() {
    fmt.Print("\r\033[K")
}

func (le *LineEditor) redraw() {
    if !le.initialized {
        return
    }

    le.clearLine()
    fmt.Print(le.prompt + le.currentLine)

    le.showSmartSuggestion()

    if le.cursorPos < len(le.currentLine) {
        fmt.Printf("\r\033[%dC", le.promptLen+le.cursorPos)
    }
}

func (le *LineEditor) historyUp() {
    if len(le.history) == 0 {
        return
    }

    if le.historyPos > 0 {
        le.historyPos--
        le.currentLine = le.history[le.historyPos]
    } else if le.historyPos == 0 {
        return
    } else {
        le.historyPos = 0
        le.currentLine = le.history[0]
    }

    le.cursorPos = len(le.currentLine)
    le.redraw()
}

func (le *LineEditor) historyDown() {
    if len(le.history) == 0 {
        return
    }

    if le.historyPos < len(le.history)-1 {
        le.historyPos++
        le.currentLine = le.history[le.historyPos]
    } else {
        le.historyPos = len(le.history)
        le.currentLine = ""
    }

    le.cursorPos = len(le.currentLine)
    le.redraw()
}

func (le *LineEditor) handleEscapeSequence(scanner *bufio.Scanner) bool {
    if !scanner.Scan() {
        return false
    }
    first := scanner.Text()

    if first == "[" {
        if !scanner.Scan() {
            return false
        }
        key := scanner.Text()

        switch key {
        case "A":
            if le.inPredictionMode {
                le.prevPrediction()
            } else {
                le.historyUp()
            }
        case "B":
            if le.inPredictionMode {
                le.nextPrediction()
            } else {
                le.historyDown()
            }
        case "C":
            if le.inPredictionMode {
                le.nextPrediction()
            } else {
                le.moveRight()
            }
        case "D":
            if le.inPredictionMode {
                le.prevPrediction()
            } else {
                le.moveLeft()
            }
        case "H":
            if le.inPredictionMode {
                le.firstPrediction()
            } else {
                le.moveHome()
            }
        case "F":
            if le.inPredictionMode {
                le.lastPrediction()
            } else {
                le.moveEnd()
            }
        case "5~":
            if le.inPredictionMode {
                le.firstPrediction()
            }
        case "6~":
            if le.inPredictionMode {
                le.lastPrediction()
            }
        }
        return true
    } else if first == "O" {
        if !scanner.Scan() {
            return false
        }
        key := scanner.Text()
        switch key {
        case "H":
            if le.inPredictionMode {
                le.firstPrediction()
            } else {
                le.moveHome()
            }
        case "F":
            if le.inPredictionMode {
                le.lastPrediction()
            } else {
                le.moveEnd()
            }
        }
        return true
    }
    return false
}

type terminalState struct {
    original syscall.Termios
}

// getIoctlConstants returns the appropriate ioctl constants for the current platform
func getIoctlConstants() (get uintptr, set uintptr) {
    switch runtime.GOOS {
    case "darwin", "freebsd", "openbsd", "netbsd":
        return ioctl_TIOCGETA, ioctl_TIOCSETA
    case "linux", "android":
        return ioctl_TCGETS, ioctl_TCSETS
    default:
        // Fallback for other Unix-like systems
        return ioctl_TCGETS, ioctl_TCSETS
    }
}

// makeRaw puts the terminal into raw mode
func makeRaw(fd uintptr) (*terminalState, error) {
    if runtime.GOOS == "windows" {
        return &terminalState{}, nil
    }
    
    var termios syscall.Termios
    
    ioctlGet, ioctlSet := getIoctlConstants()
    
    // Get current terminal settings
    _, _, err := syscall.Syscall6(
        syscall.SYS_IOCTL,
        fd,
        ioctlGet,
        uintptr(unsafe.Pointer(&termios)),
        0, 0, 0,
    )
    if err != 0 {
        return nil, fmt.Errorf("failed to get terminal settings: %v", err)
    }

    original := termios

    // Set terminal to raw mode
    termios.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.ISIG
    termios.Cc[syscall.VMIN] = 1
    termios.Cc[syscall.VTIME] = 0

    // Apply new terminal settings
    _, _, err = syscall.Syscall6(
        syscall.SYS_IOCTL,
        fd,
        ioctlSet,
        uintptr(unsafe.Pointer(&termios)),
        0, 0, 0,
    )
    if err != 0 {
        return nil, fmt.Errorf("failed to set terminal to raw mode: %v", err)
    }

    return &terminalState{original: original}, nil
}

// restoreTerminal restores the terminal to its original state
func restoreTerminal(fd uintptr, state *terminalState) error {
    if runtime.GOOS == "windows" {
        return nil
    }
    
    _, ioctlSet := getIoctlConstants()
    
    _, _, err := syscall.Syscall6(
        syscall.SYS_IOCTL,
        fd,
        ioctlSet,
        uintptr(unsafe.Pointer(&state.original)),
        0, 0, 0,
    )
    if err != 0 {
        return fmt.Errorf("failed to restore terminal settings: %v", err)
    }
    return nil
}

func (le *LineEditor) ReadLine() (string, error) {
    le.currentLine = ""
    le.cursorPos = 0
    le.historyPos = len(le.history)
    le.predictions = []string{}
    le.predictionPos = -1
    le.showPredictions = false
    le.inPredictionMode = false
    le.originalLine = ""
    le.originalCursor = 0
    le.tabPressTime = time.Time{}

    fmt.Print(le.prompt)

    if runtime.GOOS != "windows" {
        state, err := makeRaw(uintptr(syscall.Stdin))
        if err != nil {
            return le.fallbackReadLine()
        }
        defer restoreTerminal(uintptr(syscall.Stdin), state)
    }

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanBytes)

    for scanner.Scan() {
        char := scanner.Text()
        switch char {
        case "\r", "\n":
            if le.inPredictionMode {
                le.applyPrediction()
            } else {
                fmt.Printf("%s", "\n")
                line := le.currentLine
                if line != "" {
                    le.addToHistory(line)
                }
                return line, nil
            }
        case "\x7f", "\x08":
            le.backspace()
        case "\x1b":
            le.handleEscapeSequence(scanner)
        case "\x03":
            return "", fmt.Errorf("interrupted")
        case "\x04":
            if le.currentLine == "" {
                return "", fmt.Errorf("EOF")
            }
        case "\t":
            le.handleTabCompletion()
        default:
            if len(char) > 0 {
                r := rune(char[0])
                if unicode.IsPrint(r) && r != '\r' && r != '\n' {
                    le.insertChar(r)
                }
            }
        }
    }
    return "", scanner.Err()
}

func (le *LineEditor) fallbackReadLine() (string, error) {
    fmt.Print(le.prompt)
    if Scanner.Scan() {
        line := Scanner.Text()
        if line != "" {
            le.addToHistory(line)
        }
        return line, nil
    }
    return "", Scanner.Err()
}

func (le *LineEditor) enableEmoji() {
    le.emojiEnabled = true
}

func (le *LineEditor) enableSound() {
    le.soundEnabled = true
}
