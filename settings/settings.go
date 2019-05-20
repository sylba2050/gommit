package settings

import (
    "os"
    "path/filepath"
)

type prefix struct {
	Name     string
	Description string
}

func FileExists(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func isExitstSettingFile() bool {
    homepath := os.Getenv("HOME")
    return FileExists(filepath.Join(homepath, ".gommit.config"))
}


func GetSettings() []prefix{
	prefixes := []prefix{
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
    return prefixes
}
