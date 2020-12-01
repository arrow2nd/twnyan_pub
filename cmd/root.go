package cmd

import (
	"os"

	"github.com/arrow2nd/twnyan/config"
	"github.com/arrow2nd/twnyan/twitter"
	"gopkg.in/abiosoft/ishell.v2"
)

const version = "0.0.4"

var (
	shell    = ishell.New()
	cfg      config.ConfigData
	tweets   twitter.Tweets
	listName []string
	listID   []int64
)

func init() {
	// 設定読み込み
	if err := cfg.Load(); err != nil {
		config.Setup()
		cfg.Load()
	}

	// プロンプト
	shell.SetPrompt(cfg.Default.Prompt)

	// API認証処理
	twitter.SetConfig(&cfg)

	// リストを取得
	listName, listID = twitter.GetLists()
}

// Run 実行
func Run() {
	if len(os.Args) > 1 {
		shell.Process(os.Args[1:]...)
	} else {
		shell.Process("timeline")
		shell.Run()
	}
}
