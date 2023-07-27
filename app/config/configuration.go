package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// 설정 파일 구조체
type Configuration struct {
	WebServer WebServer `json:"server"`
	Database  Database  `json:"database"`
	Swagger   Swagger   `json:"swagger"`
}

// 설정파일 내용 출력
func (c *Configuration) String() string {
	return "WebServer: {" + c.WebServer.String() + "}\nDatabase: {" + c.Database.String() + "}"
}

// 설정 파일 경로를 받아서 구조체에 저장
func (c *Configuration) LoadConfiguration(path string) error {
	// 파일을 읽어서 구조체에 저장
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed open config file [" + path + "]")
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		fmt.Println("Failed parsing config file Contents")
		return err
	}
	return nil
}
