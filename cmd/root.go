package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
	"trash/internal/rename"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "trash",
	Args: cobra.MinimumNArgs(1),
	Run:  putData,
}

func checkDir() {
	if _, err := os.Stat(aConf.TrashPath); err != nil {
		os.Mkdir(aConf.TrashPath, 0755)
	}
}

func checkDate() {
	today := time.Now().Format("2006/01/02 15:04:05")
	todayParse, _ := time.Parse("2006/01/02 15:04:05", today)
	timeDay := 24 * time.Hour

	files, _ := ioutil.ReadDir(aConf.TrashPath)

	// PutDate Asc (昇順)
	sort.Slice(files, func(i, j int) bool {
		var fInfoI = make(map[string]string, 2)
		var fInfoJ = make(map[string]string, 2)
		rename.Decode(files[i].Name(), &fInfoI)
		rename.Decode(files[j].Name(), &fInfoJ)

		putDateI, _ := time.Parse("2006/01/02 15:04:05", fInfoI["PutDate"])
		putDateJ, _ := time.Parse("2006/01/02 15:04:05", fInfoJ["PutDate"])

		return putDateI.Unix() < putDateJ.Unix()
	})

	var fInfo = make(map[string]string, 2)
	for _, file := range files {
		if file.Name() == "config.toml" {
			continue
		}
		if err := rename.Decode(file.Name(), &fInfo); err != nil {
			fmt.Println(err)
			continue
		}

		putDate, _ := time.Parse("2006/01/02 15:04:05", fInfo["PutDate"])
		if todayParse.Sub(putDate) >= time.Duration(aConf.DeleteTime)*timeDay {
			if err := os.RemoveAll(filepath.Join(aConf.TrashPath, file.Name())); err != nil {
				fmt.Println(err)
			}
		} else {
			break
		}
	}
}

// Execute aabbcc
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
