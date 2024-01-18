package cfg

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	SqliteEnable            bool   `json:"sqlite_enable"`
	SqliteUrl               string `json:"sqlite_url"`
	MysqlEnable             bool   `json:"mysql_enable"`
	MysqlUrl                string `json:"mysql_url"`
	ProxyEnable             bool   `json:"proxy_enable"`
	ProxyUrl                string `json:"proxy_url"`
	TransmissionUrl         string `json:"transmission_url"`
	TransmissionUser        string `json:"transmission_user"`
	TransmissionPasswd      string `json:"transmission_passwd"`
	TransmissionDownloadDir string `json:"transmission_download_dir"`
	LogUrl                  string `json:"log_url"`
	Season                  string `json:"season"`
}

func NewConfig(name string) *Config {
	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	c := &Config{}

	configTypeOf := reflect.TypeOf(*c)
	elem := reflect.ValueOf(c).Elem()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "##") {
			continue
		}
		parts := strings.Split(line, "=")
		for i := 0; i < configTypeOf.NumField(); i++ {
			field := configTypeOf.Field(i)
			attributeJsonName := field.Tag.Get("json")
			if parts[0] == attributeJsonName {
				if elem.Field(i).Kind() == reflect.Bool {
					if parts[1] == "true" {
						elem.Field(i).SetBool(true)
						break
					}
				}
				if elem.Field(i).Kind() == reflect.String {
					elem.Field(i).SetString(parts[1])
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return c
}
