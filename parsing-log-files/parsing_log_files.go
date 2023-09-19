package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
  re, err := regexp.Compile(`^(\[((TRC)|(DBG)|(INF)|(WRN)|(ERR)|(FTL))\])`)
  if err != nil {
    panic(err)
  }
  return re.MatchString(text)
}

func SplitLogLine(text string) []string {
  re, err := regexp.Compile(`<[~*=-]*>`)
  if err != nil {
    panic(err)
  }
  return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
  var s string;
  for _, line := range lines {
    s += line;
  }
  re, err := regexp.Compile(`(?i)"\w*password\w*"`)
  if err != nil {
    panic(err)
  }
  matches := re.FindAllStringIndex(s, -1)
  if matches == nil {
    return 0
  }
  return len(matches[0])
}

func RemoveEndOfLineText(text string) string {
  re, err := regexp.Compile(`end-of-line\d+`)
  if err != nil {
    panic(err)
  }
  return re.ReplaceAllString(text, "")


}

func TagWithUserName(lines []string) []string {
  userRe, err := regexp.Compile(`User\s+\w+`)
  if err != nil {
    panic(err)
  }
  var taggedLines []string
  for _, line := range lines {
    res := userRe.FindString(line)
    if res != "" {
      user := "[USR] " + strings.Fields(res)[1] + " "
      taggedLines = append(taggedLines, user + line)
    } else {
    taggedLines = append(taggedLines, line)
    }
  }
  return taggedLines
}
