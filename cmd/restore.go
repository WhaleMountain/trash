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
	restoreCmd := &cobra.Command{
		Use:  "restore",
		Args: cobra.MinimumNArgs(1),
		Run:  restoreData,
	}
	restoreCmd.Flags().StringVarP(&aFlag.RestorePath, "restore-path", "p", aConf.RestorePath, "Restore path")

	rootCmd.AddCommand(restoreCmd)
}

func restoreData(cmd *cobra.Command, args []string) {
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
			// 同じ名前のファイル・フォルダが存在する
			if _, err := os.Stat(filepath.Join(aFlag.RestorePath, fInfo["Name"])); err == nil {
				fmt.Println(aFlag.RestorePath + "/" + fInfo["Name"] + " is already exists.")
				return
			}

			// 指定場所に復元する
			target, _ := filepath.Abs(filepath.Join(aConf.TrashPath, file.Name()))
			if err := os.Rename(target, filepath.Join(aFlag.RestorePath, fInfo["Name"])); err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	<-checkDone
}
