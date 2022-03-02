package universal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Path        string
	Description string
	Annotation  map[int]string
	Map         map[string]string
	LineAt      map[string]int
	SortedLine  []*line
	IsOk        int
}
type line struct {
	key string
	at  int
}

func NewConfig(path string, description string) (config Config, err error) {
	config.Map = make(map[string]string)
	config.LineAt = make(map[string]int)
	config.Annotation = make(map[int]string)
	reader, err := os.Open(path)
	if !CheckErr(err, 0) {
		return Config{}, errors.New("路径不正确！")
	}
	defer func() {
		err = reader.Close()
		CheckErr(err, 0)
	}()
	BufferedReader := bufio.NewReader(reader)
	var tempBytes []byte
	lineCount := 0
	for {
		tempBytes, _, err = BufferedReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		lineCount++
		tempString := strings.TrimSpace(string(tempBytes))
		index := strings.Index(tempString, "=")
		if tempString == "" {
			continue
		} else if tempString[0] == '#' {
			config.Annotation[lineCount] = tempString[1:]
			continue
		}
		if index == 0 {
			return Config{}, errors.New("第" + strconv.Itoa(lineCount) + "行没有找到key")
		}
		key := tempString[:index]
		var value string
		if index == len(tempString)-1 {
			value = ""
		} else {
			value = tempString[index+1:]
		}
		key = strings.TrimSpace(key)
		key = strings.Trim(key, "\"")
		value = strings.TrimSpace(value)
		value = strings.Trim(value, "\"")
		config.Map[key] = value
		config.LineAt[key] = lineCount

	}
	config.Description = description
	config.Path = path
	config.IsOk = 666
	return
}
func (config *Config) SetSortedLine() {

	for k, v := range config.LineAt {
		config.SortedLine = append(config.SortedLine, &line{k, v})
	}
	length := len(config.SortedLine)
	for i := 0; i < length; i++ {
		min := i
		for j := i; j < length; j++ {
			if config.SortedLine[min].at > config.SortedLine[j].at {
				min = j
			}
		}
		if i != min {
			config.SortedLine[min], config.SortedLine[i] = config.SortedLine[i], config.SortedLine[min]
		}
	}
}

func (config *Config) Print() {
	SLIndex := 0
	flag := true
	for k, v := range config.Annotation {
		if flag {
			if SLIndex < len(config.SortedLine) {
				for config.SortedLine[SLIndex].at < k {
					fmt.Println(config.SortedLine[SLIndex].key, ":", config.Map[config.SortedLine[SLIndex].key])
				}
			} else {
				flag = false
			}
		}
		fmt.Printf("#%s\n", v)
	}

}
