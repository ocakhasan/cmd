package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {
	topics := GetEksiAgenda()

	prompt := promptui.Select{
		Label: fmt.Sprintf("Select Topic - %s", time.Now().Format("2006-01-02")),
		Items: topics.toStringArray(),
		Size:  20,
	}

	_, result, _ := prompt.Run()

	idx := -1
	for i, topic := range topics {
		if topic.String() == result {
			idx = i
			break
		}
	}

	if idx != -1 {
		if err := openbrowser(topics[idx].URL); err != nil {
			log.Fatal(err)
		}
	}

}

// got from https://gist.github.com/hyg/9c4afcd91fe24316cbf0
func openbrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
