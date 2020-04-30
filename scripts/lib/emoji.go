package slacklog

import (
	"encoding/json"
	"io/ioutil"
)

// EmojiTable : 絵文字データを保持する。
// urlMapは絵文字名をキーとし、画像が置いてあるURLが値である。
type EmojiTable struct {
	urlMap map[string]string
}

// NewEmojiTable : pathに指定したJSON形式の絵文字データを読み込み、EmojiTableを
// 生成する。
func NewEmojiTable(path string) *EmojiTable {
	emojis := &EmojiTable{
		urlMap: map[string]string{},
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return emojis
	}

	json.Unmarshal(content, &emojis.urlMap)

	return emojis
}
