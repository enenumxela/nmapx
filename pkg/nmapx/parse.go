package nmapx

import (
	"encoding/xml"
)

func Parse(input []byte) (output *NmapRun, err error) {
	output = &NmapRun{
		Raw: input,
	}

	if err = xml.Unmarshal(input, output); err != nil {
		return
	}

	return
}
