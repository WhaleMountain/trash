package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"trash/internal/rename"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "remove",
		Args: cobra.ExactArgs(1),
		Run:  removeData,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use: "removeall",
		Run: removeAll,
	})
}

func removeData(cmd *cobra.Command, args []string) {
	checkDone := make(chan struct{})

	go func() {
		defer func() { checkDone <- struct{}{} }()
		checkDate()
	}()

	checkDir()

	files, err := ioutil.ReadDir(aConf.TrashPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var fInfo = make(map[string]string, 2)

	for _, file := range files {
		if file.Name() == "config.toml" {
			continue
		}

		if err := rename.Decode(file.Name(), &fInfo); err != nil {
			fmt.Println(err)
			continue
		}

		if fInfo["Name"] == args[0] {
			if err := os.RemoveAll(filepath.Join(aConf.TrashPath, file.Name())); err != nil {
				fmt.Println(err)
			} else {
				return
			}
		}
	}

	<-checkDone
}

func removeAll(cmd *cobra.Command, args []string) {

	checkDir()

	files, err := ioutil.ReadDir(aConf.TrashPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.Name() == "config.toml" {
			continue
		}

		if err := os.RemoveAll(filepath.Join(aConf.TrashPath, file.Name())); err != nil {
			fmt.Println(err)
		}
	}
}
