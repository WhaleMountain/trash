package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"
	"trash/internal/rename"

	"code.cloudfoundry.org/bytefmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "list",
		Run: showList,
	})
}

func showList(cmd *cobra.Command, args []string) {
	checkDone := make(chan struct{})

	go func() {
		defer func() { checkDone <- struct{}{} }()
		checkDir()
		checkDate()
	}()

	files, err := ioutil.ReadDir(aConf.TrashPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	var fInfo = make(map[string]string, 2)

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "-- Name -- \t -- Size -- \t -- Create Date -- \t -- Put Date --")
	for _, file := range files {
		if file.Name() == "config.toml" {
			continue
		}
		if err := rename.Decode(file.Name(), &fInfo); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Fprintln(w, fInfo["Name"], "\t", bytefmt.ByteSize(uint64(file.Size())), "\t", file.ModTime().Format("2006/01/02 15:04:05"), "\t", fInfo["PutDate"])
	}
	w.Flush()

	<-checkDone
}
