package main

import (
	"os"
	"bytes"
	"fmt"
	"strconv"
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

var BLOCK_NAME = []string{"gaoxin7", "tianfuxinqu"}
var BASE_STRING = "http://cd.lianjia.com/ershoufang/%s/pg%s/"
var RESULT_FILE = "result_%s.txt"

func main() {

	for _, blockName := range BLOCK_NAME {
		var allItems []HouseInfo
		var preItems []HouseInfo
		for i := 1; i < 50; i++ {
			requestUrl := fmt.Sprintf(BASE_STRING, blockName, strconv.Itoa(i))
			res, _ := GetItemFromUrl(requestUrl)
			if NeedContinue(preItems, res) {
				preItems = res
				allItems = append(allItems, res...)
			}

		}
		resultFile := fmt.Sprintf(RESULT_FILE, blockName)
		WriteHouseInfoToFile(resultFile, allItems)
	}

}
