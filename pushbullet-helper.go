/*
	My pushbullet-go helper class,
	for reading access token from a file
	and sending pushes easily.

	last update: 2015.07.29.
*/

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
	AccessTokenFilename = ".pushbullet.token"
)

// read access token from file
func readAccessToken() (token string, err error) {
	confDir := filepath.Dir(os.Args[0])
	tokenFile := filepath.Join(confDir, AccessTokenFilename)

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

// send push(note)
func SendNote(title string, message string) bool {
	return SendNoteToChannel("", title, message)
}

// send push(note) to channel
func SendNoteToChannel(channelTag string, title string, message string) bool {
	token, err := readAccessToken()
	if err != nil {
		return false
	}

	pb := pushbullet.New(token)

	note := requests.NewNote()
	if channelTag != "" {
		note.Push.ChannelTag = channelTag // set channel tag when provided
	}
	note.Title = title
	note.Body = message

	if _, err := pb.PostPushesNote(note); err != nil {
		fmt.Println("Push failed:", err)
		return false
	}
	return true
}

// send push(link)
func SendLink(title string, message string, url string) bool {
	return SendLinkToChannel("", title, message, url)
}

// send push(link) to channel
func SendLinkToChannel(channelTag string, title string, message string, url string) bool {
	token, err := readAccessToken()
	if err != nil {
		return false
	}

	pb := pushbullet.New(token)

	link := requests.NewLink()
	if channelTag != "" {
		link.Push.ChannelTag = channelTag // set channel tag when provided
	}
	link.Title = title
	link.Body = message
	link.Url = url

	if _, err := pb.PostPushesLink(link); err != nil {
		fmt.Println("Push failed:", err)
		return false
	}
	return true
}
