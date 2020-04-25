package slacklog

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/slack-go/slack"
)

func doDownloadEmoji() error {
	slackToken := os.Getenv("SLACK_TOKEN")
	if slackToken == "" {
		return fmt.Errorf("$SLACK_TOKEN required")
	}

	if len(os.Args) < 4 {
		fmt.Println("Usage: go run scripts/main.go emoji {emojis-dir}")
		return nil
	}

	emojisDir := filepath.Clean(os.Args[2])

	api := slack.New(slackToken)

	emojis, err := api.GetEmoji()
	if err != nil {
		return err
	}

	err = mkdir(emojisDir)
	if err != nil {
		return err
	}

	for name, url := range emojis {
		if url[:6] == "alias:" {
			continue
		}
		downloadEmojiToFile(url, name, emojisDir)
		emojis[name] = filepath.Ext(emojis[name])
	}

	emojisJson, err := json.Marshal(emojis)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(emojisDir, "emoji.json"), emojisJson, 0666)
	if err != nil {
		return err
	}

	return nil
}

func downloadEmojiToFile(url, name, basePath string) error {
	extension := filepath.Ext(url)
	destFile := filepath.Join(basePath, name+extension)
	if _, err := os.Stat(destFile); err == nil {
		return nil
	}

	fmt.Printf("Downloading: :%s:\n", name)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("[%s]: %s", resp.Status, url)
	}

	w, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, resp.Body)

	return err
}
