package parsinglogfiles

import "regexp"

var validLineRegex = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
var splitLineRegex = regexp.MustCompile(`<[~*=-]*>`)
var passwordRegex = regexp.MustCompile(`"(?i).*password.*"`)
var removeTextRegex = regexp.MustCompile(`end-of-line\d+`)
var userRegex = regexp.MustCompile(`User\s+(\S+)`)

func IsValidLine(text string) bool {
    return validLineRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
    return splitLineRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0

    for _, line := range lines {
        if passwordRegex.MatchString(line) {
            count++
        }           
    }

    return count
}

func RemoveEndOfLineText(text string) string {
	return removeTextRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    result := make([]string, len(lines))
    
    for i, line := range lines {
        match := userRegex.FindStringSubmatch(line)
        if match != nil {
            result[i] = "[USR] " + match[1] + " " + line
        } else {
            result[i] = line
        }
    }

    return result
}