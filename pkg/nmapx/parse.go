package nmapx

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

func Parse(input string) (output *NmapRun, err error) {
	output = &NmapRun{}

	xmlFile, err := os.Open(input)
	if err != nil {
		return
	}

	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	if err = xml.Unmarshal(xmlData, output); err != nil {
		return
	}

	return
}
