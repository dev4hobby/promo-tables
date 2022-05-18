package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Env struct {
	AppEnv                 string      `json:"APP_ENV"`
	MySQLUser              string      `json:"MYSQL_USER"`
	MySQLPassword          string      `json:"MYSQL_PASSWORD"`
	MySQLHost              string      `json:"MYSQL_HOST"`
	MySQLPort              string      `json:"MYSQL_PORT"`
	MySQLDatabase          string      `json:"MYSQL_DATABASE"`
	MySQLCharset           string      `json:"MYSQL_CHARSET"`
	MockUserCount          json.Number `json:"MOCK_USER_COUNT"`
	MockPromoCategoryCount json.Number `json:"MOCK_PROMO_CATEGORY_COUNT"`
}

var env Env
var projectRootDir = "promo-tables"

func init() {
	re := regexp.MustCompile(`^(.*` + projectRootDir + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	envPath := string(rootPath) + "/.env"

	data, err := ioutil.ReadFile(envPath)
	CheckError(err)

	_env := map[string]string{}
	for _, line := range strings.Split(string(data), "\n") {
		kv := strings.SplitN(line, "=", 2)
		if len(kv) != 2 {
			continue
		}

		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		_env[key] = value
	}

	serialized, err := json.Marshal(_env)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(serialized, &env)
	if err != nil {
		panic(err)
	}
}

func GetEnv() Env {
	return env
}

func (o Env) IsLocal() bool {
	return o.AppEnv == "local"
}

func (o Env) IsGithubAction() bool {
	return o.AppEnv == "github-action"
}
