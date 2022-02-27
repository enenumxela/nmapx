package commands

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/enenumxela/nmapx/pkg/nmapx"
	"github.com/spf13/cobra"
)

func Merge() (cmd *cobra.Command) {
	var (
		directory, output string
	)

	cmd = &cobra.Command{
		Use:   "merge",
		Short: "merge nmap XML output files into one.",
		Long:  "merge nmap XML output files into one.",
		Run: func(cmd *cobra.Command, args []string) {
			fileInfo, err := os.Stat(directory)
			if err != nil {
				log.Fatalln(err)
			}

			if !fileInfo.IsDir() {
				log.Fatalln("not a directory")
			}

			files, err := ioutil.ReadDir(directory)
			if err != nil {
				log.Fatalln(err)
			}

			data := &nmapx.NmapRun{}

			for index, file := range files {
				if !file.IsDir() && filepath.Ext(file.Name()) == ".xml" {
					xmlFilePath := filepath.Join(directory, file.Name())

					xmlFile, err := os.Open(xmlFilePath)
					if err != nil {
						fmt.Println(err)
						continue
					}

					defer xmlFile.Close()

					xmlData, err := ioutil.ReadAll(xmlFile)
					if err != nil {
						fmt.Println(err)
						continue
					}

					scanData, err := nmapx.Parse(xmlData)
					if err != nil {
						fmt.Println(err)
						continue
					}

					if index == 0 {
						data = scanData
					}

					data.Hosts = append(data.Hosts, scanData.Hosts...)
				}
			}

			file, err := xml.MarshalIndent(data, "", "\t")
			if err != nil {
				log.Fatalln(err)
			}

			if filepath.Ext(output) != ".xml" {
				output = output + ".xml"
			}

			if err = ioutil.WriteFile(output, file, 0644); err != nil {
				log.Fatalln(err)
			}
		},
	}

	cmd.Flags().StringVarP(&directory, "directory", "d", "./", "data to use")
	cmd.Flags().StringVarP(&output, "output", "o", "nmapx.xml", "merged XML file")

	return
}
