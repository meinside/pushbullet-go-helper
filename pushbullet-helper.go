package pbhelper

import (
	"fmt"
	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	ACCESS_TOKEN_FILENAME = ".pushbullet.token"
)

// read access token from file
func readAccessToken() (token string, err error) {
	confDir := filepath.Dir(os.Args[0])
	tokenFile := filepath.Join(confDir, ACCESS_TOKEN_FILENAME)

	if _, err := os.Stat(tokenFile); err != nil {
		fmt.Printf("Access token file does not exist: %s\n", tokenFile)
		return "", err
	}

	buf, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		fmt.Println("Failed to read access token file:", err)
		return "", err
	}

	return strings.TrimSpace(string(buf)), nil
}

// send push (note)
func SendNote(title string, message string) bool {
	token, err := readAccessToken()
	if err != nil {
		return false
	}

	pb := pushbullet.New(token)

	note := requests.NewNote()
	note.Title = title
	note.Body = message

	if _, err := pb.PostPushesNote(note); err != nil {
		fmt.Println("Push failed:", err)
		return false
	}
	return true
}
