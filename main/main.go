package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func loadContentFromFile(filePath string) (*bytes.Buffer, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	var content []byte
	buffer := make([]byte, 1024)
	for {
		n, _ := file.Read(buffer)
		if 0 == n {
			break
		}
		content = append(content, buffer[:n]...)

	}
	buf := bytes.NewBuffer(content)
	return buf, nil
}

const (
	endpoint = `http://cd.lianjia.com`
)

var (
	blockNames = []string{"gaoxin7", "tianfuxinqu"}
	resultFile = "%s_result_%s.txt"
)

func main() {
	for _, blockName := range blockNames {
		var allItems []HouseInfo
		var preItems []HouseInfo
		for i := 1; i < 100; i++ {
			requestURL := fmt.Sprintf("%s/ershoufang/%s/pg%d/", endpoint, blockName, i)
			time.Sleep(time.Second * 5)
			res, _ := GetItemFromUrl(requestURL)
			if NeedContinue(preItems, res) {
				preItems = res
				allItems = append(allItems, res...)
			} else {
				fmt.Println("发现重复的,停止")
				break
			}

		}
		resultFile := fmt.Sprintf(resultFile, time.Now(), blockName)
		WriteHouseInfoToFile(resultFile, allItems)
	}
}
