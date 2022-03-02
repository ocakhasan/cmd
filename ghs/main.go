package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var queryTypes = map[string]string{
	"c":       "commits",
	"commits": "commits",
	"u":       "users",
	"users":   "users",
	"code":    "code",
	"t":       "topics",
	"topics":  "topics",
	"w":       "wikis",
	"wikis":   "wikis",
	"i":       "issues",
	"issues":  "issues",
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

func main() {
	if len(os.Args) == 1 {
		log.Fatal("A query needs to be provided")
	}

	githubURL := "https://github.com/search?q="

	if len(os.Args) > 3 {
		log.Fatal("Only 1 query and 1 type is allowed")
	}

	query := os.Args[1]

	queryValues := strings.ReplaceAll(query, " ", "+")
	githubURL = githubURL + queryValues

	if len(os.Args) == 3 {
		queryType, ok := queryTypes[os.Args[2]]
		if !ok {
			log.Fatalf("query type %s is not supported", os.Args[2])
		}
		githubURL = fmt.Sprintf("%s&type=%s", githubURL, queryType)
	}

	if err := openbrowser(githubURL); err != nil {
		log.Fatal(err)
	}
}
