package slacklog

import (
	"os"
)

// EmojiTable : 絵文字データを保持する。
type EmojiTable struct {
	// NameToExtは絵文字名をキーとし、画像の拡張子が値である。
	// 絵文字は事前に全てダウンロードしている、という前提であり、そのため拡張子のみ
	// を保持している。
	NameToExt map[string]string
}

// NewEmojiTable : pathに指定したJSON形式の絵文字データを読み込み、EmojiTableを
// 生成する。
func NewEmojiTable(path string) (*EmojiTable, error) {
	emojis := &EmojiTable{
		NameToExt: map[string]string{},
	}

	if info, err := os.Stat(path); err != nil || info.IsDir() {
		// pathにディレクトリが存在しても、その場合は無視して、ファイル自体が存在し
		// なかったこととする。
		return nil, os.ErrNotExist
	}

	if err := ReadFileAsJSON(path, true, &emojis.NameToExt); err != nil {
		return nil, err
	}

	return emojis, nil
}
