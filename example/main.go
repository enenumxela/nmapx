package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/enenumxela/nmapx/pkg/nmapx"
)

func main() {
	xmlFile, err := os.Open("sample.xml")
	if err != nil {
		return
	}

	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	scanData, err := nmapx.Parse(xmlData)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonData, err := json.Marshal(scanData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", string(jsonData))
}
