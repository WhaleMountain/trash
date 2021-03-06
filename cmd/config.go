package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

// デフォルトの設定
type trashConfig struct {
	TrashPath   string `toml:"-"`
	RestorePath string
	DeleteTime  int
}

// ユーザが指定する
type trashFlags struct {
	DeleteTime  int
	RestorePath string
}

var aConf = trashConfig{}
var aFlag = trashFlags{}

func init() {
	userHome, _ := os.UserHomeDir()
	aConf.TrashPath = filepath.Join(userHome, ".Trash")
	if err := decodeConfig(); err != nil {
		aConf.RestorePath = userHome
		aConf.DeleteTime = 30
		encodeConfig()
	}

	confCmd := &cobra.Command{
		Use: "config",
		Run: showConfig,
	}

	confSetCmd := &cobra.Command{
		Use: "set",
		Run: setConfig,
	}
	confSetCmd.Flags().IntVarP(&aFlag.DeleteTime, "delete-time", "t", 30, "Set All delete time")
	confSetCmd.Flags().StringVarP(&aFlag.RestorePath, "restore-path", "p", aConf.RestorePath, "Set Restore default path")

	confCmd.AddCommand(confSetCmd)
	rootCmd.AddCommand(confCmd)
}

func showConfig(cmd *cobra.Command, args []string) {
	checkDone := make(chan struct{})

	go func() {
		defer func() { checkDone <- struct{}{} }()
		checkDir()
		checkDate()
	}()

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "TrashPath"+"\t", aConf.TrashPath)
	fmt.Fprintln(w, "RestorePath"+"\t", aConf.RestorePath)
	fmt.Fprintln(w, "DeleteTime"+"\t", aConf.DeleteTime)
	w.Flush()

	<-checkDone
}

func setConfig(cmd *cobra.Command, args []string) {
	checkDone := make(chan struct{})

	go func() {
		defer func() { checkDone <- struct{}{} }()
		checkDir()
		checkDate()
	}()

	aConf.RestorePath = aFlag.RestorePath
	aConf.DeleteTime = aFlag.DeleteTime

	encodeConfig()

	<-checkDone
}

func encodeConfig() {
	configToml, _ := os.Create(filepath.Join(aConf.TrashPath, "config.toml"))
	if err := toml.NewEncoder(configToml).Encode(aConf); err != nil {
		fmt.Println(err)
	}
}

func decodeConfig() error {
	if _, err := toml.DecodeFile(filepath.Join(aConf.TrashPath, "config.toml"), &aConf); err != nil {
		return err
	}

	return nil
}
