package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type prefix struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func LoadJson(filepath string) []prefix {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("%v", err)
	}

	var prefixes []prefix
	if err := json.Unmarshal(bytes, &prefixes); err != nil {
		fmt.Printf("%v", err)
	}

	return prefixes
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func GetSettings() []prefix {
	homepath := os.Getenv("HOME")
	settingFilePath := filepath.Join(homepath, ".gommit.config")

	var prefixes []prefix
	if FileExists(settingFilePath) {
		prefixes = LoadJson(settingFilePath)
	} else {
		prefixes = []prefix{
			{Name: "feat", Description: "機能追加"},
			{Name: "fix", Description: "バグ修正"},
			{Name: "update", Description: "機能修正"},
			{Name: "style", Description: "機能に影響を与えない修正"},
			{Name: "doc", Description: "ドキュメントのみの修正"},
			{Name: "add", Description: "新規ファイル追加"},
			{Name: "delete", Description: "ファイル削除"},
			{Name: "refactor", Description: "リファクタリング"},
			{Name: "perf", Description: "性能向上"},
			{Name: "disable", Description: "機能削除"},
		}
	}

	return prefixes
}
