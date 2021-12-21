package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var labels = []string{"TODO", "FIXME"}

type Work struct {
	Prefix       string
	LineNumber   int
	FileName     string
	FullFileName string
	Content      string
}

func main() {
	var projectPath string
	flag.StringVar(&projectPath, "path", ".", "path to the project")
	flag.Parse()

	projectPath = strings.TrimRight(projectPath, "/")
	_, err := os.Stat(projectPath + "/.git")
	if os.IsNotExist(err) {
		log.Fatalf("look like this is not git folder=%s", projectPath)
	}

	var workNeedToBeDone []*Work
	for _, fname := range readAllFileName(projectPath) {
		fullPath := projectPath + "/" + fname
		file, err := os.Open(fullPath)
		if err != nil {
			log.Print(err)
			return
		}

		scanner := bufio.NewScanner(file)
		line := 1
		for scanner.Scan() {
			for _, label := range labels {
				regex := regexp.MustCompile(fmt.Sprintf(".*(%s|%s):?(.*)", label, strings.ToLower(label)))
				match := regex.FindStringSubmatch(scanner.Text())
				if len(match) == 0 {
					continue
				}
				content := ""
				if len(match) > 2 {
					content = match[2]
				}
				workNeedToBeDone = append(workNeedToBeDone, &Work{
					Prefix:       label,
					LineNumber:   line,
					Content:      content,
					FullFileName: fullPath,
					FileName:     fname,
				})
			}
			line = line + 1
		}
		file.Close()
	}

	for _, item := range workNeedToBeDone {
		log.Printf("%s:%d %s: %s\n", item.FileName, item.LineNumber, item.Prefix, item.Content)
	}

}

func readAllFileName(folder string) []string {
	lsFileCmd := exec.Command("git", "ls-files")
	var out bytes.Buffer
	lsFileCmd.Stdout = &out
	folderPath, err := filepath.Abs(folder)
	if err != nil {
		log.Fatal(err)
	}
	lsFileCmd.Dir = folderPath
	err = lsFileCmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	var files []string
	for _, line := range strings.Split(out.String(), "\n") {
		fullURL := folderPath + "/" + line
		if f, err := os.Stat(fullURL); err != nil || f.IsDir() {
			continue
		}
		files = append(files, line)
	}

	return files
}
