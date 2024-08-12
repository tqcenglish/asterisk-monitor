package utils

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/sirupsen/logrus"
)

// IsFlagPassed 判断字符串全是数字
func IsDigit(str string) bool {
	for _, x := range str {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

// IsFlagPassed 判断是否设置了指定 flag
func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// NumberOrStringToBool 字符串 1, yes 返回 true 否则 false
func NumberOrStringToBool(input string) bool {
	if input == "1" {
		return true
	}
	if input == "yes" {
		return true
	}
	return false
}

// NumberToBool 大于0 返回 true 否则 false
func NumberToBool(input int) bool {
	if input > 0 {
		return true
	} else {
		return false
	}
}

// Bool to int true 返回 1 否则 0
func BoolToInt(input bool) int {
	var output int
	if input {
		output = 1
	} else {
		output = 0
	}
	return output
}

// BoolToNumber true 返回 "1" false 返回  "0"
func BoolToNumber(input bool) string {
	if input {
		return "1"
	}
	return "0"
}

// BoolToString true 返回 "yes" false 返回  "no"
func BoolToString(input bool) string {
	if input {
		return "yes"
	}
	return "no"
}

// YesToOn  Yes 或 yes 返回 "on" 否则返回 "off"
// 会议 Muted 的值可能是 By admin
func YesToOn(input string) string {
	if input == "Yes" || input == "yes" || strings.Contains(input, "By") {
		return "on"
	}
	return "off"
}

// GetRandStr 生成随机密码
func GetRandStr(baseStr string, length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	bytes := make([]byte, length)
	l := len(baseStr)
	for i := 0; i < length; i++ {
		bytes[i] = baseStr[r.Intn(l)]
	}
	return string(bytes)
}

// SqlLike sql like 头尾添加 %%
func SqlLike(input string) string {
	return fmt.Sprintf("%%%s%%", input)
}

// IsChannel 判断输入是否是 Channel
func IsChannel(input string) bool {
	if strings.Contains(input, "SIP") {
		return true
	}

	if strings.Contains(input, "Local") {
		return true
	}

	if strings.Contains(input, "DAHDI") {
		return true
	}
	return false
}

// 判断文件是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 获取指定目录所有文件
func AllFiles(path string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		logrus.Error(err)
	}
	return files
}
