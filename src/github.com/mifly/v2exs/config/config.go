package config

import (
	"fmt"
	"errors"
	"os"
	"bufio"
	"strings"
	"log"
)

var configs map[string]string

func init() {
	fmt.Printf("config init\n")
	configs = make(map[string]string, 0)
	err := Load("conf/app_dev.conf")
	if err != nil {
		err = Load("conf/app.conf")
		if err != nil {
			log.Fatal("can not load app.conf")
		}
	}
}

func Load(filePath string) error {
	if len(filePath) == 0 {
		return errors.New("file path must not empty")
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open %s failed\n", filePath)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				configs[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scan file %s failed\n", filePath)
		return err
	}

	return nil
}

func Get(key string) string {
	if v, ok := configs[key]; ok {
		return v
	}
	return ""
}
