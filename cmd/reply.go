package cmd

import (
	"net/url"

	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newReplyCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "reply",
		Aliases: []string{"rp"},
		Func: func(c *ishell.Context) {
			// 引数をチェック
			if len(c.Args) < 1 {
				cmd.drawWrongArgMessage(c.Cmd.Name)
				return
			}
			// リプライ先のツイートIDを取得
			tweetID, err := cmd.view.GetDataFromTweetNum(c.Args[0], "tweetID")
			if err != nil {
				cmd.drawErrorMessage(err.Error())
				return
			}
			// 引数をパース
			status, files := cmd.parseTweetCmdArgs(c.Args[1:])
			// リプライ先を設定
			val := url.Values{}
			val.Add("in_reply_to_status_id", tweetID)
			val.Add("auto_populate_reply_metadata", "true")
			// 画像をアップロード
			if len(files) != 0 {
				mediaIDs, err := cmd.upload(files)
				if err != nil {
					cmd.drawErrorMessage(err.Error())
					return
				}
				val.Add("media_ids", mediaIDs)
			}
			// リプライ
			tweetStr, err := cmd.api.PostTweet(val, status)
			if err != nil {
				cmd.drawErrorMessage(err.Error())
				return
			}
			cmd.drawMessage("REPLYED", tweetStr, cmd.cfg.Color.Reply)
		},
		Help: "post a reply",
		LongHelp: createLongHelp(
			"Post a reply.\nIf there is no tweet text, 'にゃーん' will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"rp",
			"reply [<tweetnumber>] [text] [image]...",
			"reply 2 meow cat.jpg",
		),
	})
}
