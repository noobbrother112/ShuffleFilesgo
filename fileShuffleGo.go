package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type Config struct {
	ListPath  string `json:"list_path"`
	GoalPath  string `json:"goal_path"`
	FileName  string `json:"file_name"`
	Extension string `json:"extension"`
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("파일을 읽을 수 없습니다:", err)
		return
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		fmt.Println("JSON 데이터를 파싱할 수 없습니다:", err)
		return
	}
	pickUpFile := config.ListPath + getFileList(config.ListPath)
	destPath := config.GoalPath + config.FileName + config.Extension
	fmt.Println(pickUpFile)
	fmt.Println(destPath)
	err = overwriteFile(pickUpFile, destPath)
	if err != nil {
		return
	}

}

func overwriteFile(srcPath string, destPath string) error {
	// Read the contents of the source file.
	srcBytes, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
	}

	// Write the contents of the source file to the destination file.
	err = ioutil.WriteFile(destPath, srcBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getFileList(dir string) string {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(20 * time.Millisecond)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln("Error reading directory:", err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	index := rand.Intn(len(fileNames))

	return fileNames[index]
}
