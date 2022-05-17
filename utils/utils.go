package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

type IntArray []int

type GitLog struct {
	CommitHash string
	Author     string
	DateStr    string
	Message    string
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (s *IntArray) UnmarshalJSON(b []byte) error {
	if b == nil {
		*s = []int{}
	} else {
		var arr []int
		err := json.Unmarshal([]byte(strings.ReplaceAll(string(b), "\"", "")), &arr)
		if err == nil {
			*s = arr
		}
	}

	return nil
}

func IsTesting() bool {
	return strings.Contains(os.Args[0], "test.test")
}

func ParseGit() GitLog {
	out, err := exec.Command("git", "log", "-1").CombinedOutput()

	if err != nil {
		panic(err)
	}

	var lines []string
	for _, line := range strings.Split(string(out), "\n") {
		l := strings.TrimSpace(line)
		if len(l) > 0 {
			lines = append(lines, l)
		}
	}

	return GitLog{
		CommitHash: strings.Split(lines[0], " ")[1],
		Author:     strings.SplitN(lines[1], " ", 2)[1],
		DateStr:    strings.TrimSpace(strings.SplitN(lines[2], " ", 2)[1]),
		Message:    lines[3],
	}
}

func Contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

func PlainTextToSHA256(plainText string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(plainText)))
}

func IntegerToString(x int) string {
	return fmt.Sprintf("%d", x)
}

func GetUTCNowWithDiff(diff int) time.Time {
	return time.Now().UTC().Add(time.Hour * time.Duration(diff))
}

func MustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	CheckError(err)
	return b
}

func PickRandomIndex(arr []int) int { // +1.18, we can refactor this to Generic function
	return rand.Intn(len(arr))
}

func GenerateRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}
