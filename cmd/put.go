package cmd

import (
	"areus/internal/rename"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "put",
		Run: putData,
	})
}

func putData(cmd *cobra.Command, args []string) {

	// ゴミ箱が存在するか
	checkDir()
	checkDate()

	target, _ := filepath.Abs(args[0])
	fileName := filepath.Base(args[0])

	// 指定されたファイル・ディレクトリが存在するか
	if _, err := os.Stat(target); err != nil {
		fmt.Println(fileName + " is not found.")
		return
	}

	// すでにゴミ箱に存在する
	if err := checkExists(fileName); err != nil {
		ext := filepath.Ext(fileName)
		fileName = strings.Join([]string{fileName[:len(fileName)-len(ext)], time.Now().Format("_150405_"), ext}, "")
	}

	fInfo := map[string]string{
		"Name":    fileName,
		"PutDate": time.Now().Format("2006/01/02 15:04:05"),
	}
	encodeFileName := rename.Encode(fInfo)

	// ゴミ箱に移動する
	if err := os.Rename(target, filepath.Join(aConf.TrashPath, encodeFileName)); err != nil {
		fmt.Println(err)
		return
	}
}

func checkExists(fileName string) error {
	files, _ := ioutil.ReadDir(aConf.TrashPath)

	var fInfo = make(map[string]string, 2)
	for _, file := range files {
		if file.Name() == "config.toml" {
			continue
		}
		if err := rename.Decode(file.Name(), &fInfo); err != nil {
			return err
		}

		if fInfo["Name"] == fileName {
			return fmt.Errorf("already exists")
		}
	}

	return nil
}
