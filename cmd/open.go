package cmd

import (
	"github.com/gookit/color"
	"github.com/pkg/browser"
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newOpenCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "open",
		Aliases: []string{"op"},
		Func: func(c *ishell.Context) {
			// 引数をチェック
			if len(c.Args) != 1 {
				cmd.drawWrongArgMessage(c.Cmd.Name)
				return
			}
			// 該当ツイートのURLを取得
			uri, err := cmd.view.GetTweetURL(c.Args[0])
			if err != nil {
				color.Error.Prompt(err.Error())
				return
			}
			cmd.drawMessage("OPENED", uri, cmd.cfg.Color.Accent3)
			// ブラウザを開く
			browser.OpenURL(uri)
		},
		Help: "view the tweet in your browser",
		LongHelp: createLongHelp(
			"View the tweet in your browser.",
			"op",
			"open [<tweetnumber>]",
			"open 2",
		),
	})
}
