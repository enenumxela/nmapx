package main

import (
	"encoding/json"
	"fmt"

	"github.com/enenumxela/nmapx/pkg/nmapx"
)

func main() {
	scanData, err := nmapx.Parse("sample.xml")
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
