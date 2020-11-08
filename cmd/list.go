package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"
	"trash/internal/rename"
	"trash/internal/sort"

	"code.cloudfoundry.org/bytefmt"
	"github.com/spf13/cobra"
)

// Sort
type sortFlag struct {
	Sort string
	Desc bool
}

var sFlag = sortFlag{
	Sort: "name",
	Desc: false,
}

func init() {
	listCmd := &cobra.Command{
		Use: "list",
		Run: showList,
	}

	sortUsage := "Specify the column you want to sort. If you want to sort in descending order, add -d"
	listCmd.Flags().StringVarP(&sFlag.Sort, "sort", "s", sFlag.Sort, sortUsage)
	listCmd.Flags().BoolVarP(&sFlag.Desc, "desc", "d", sFlag.Desc, "Descending sort")

	rootCmd.AddCommand(listCmd)
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

	switch sFlag.Sort {
	case "Name", "name":
		if sFlag.Desc {
			sort.NameDesc(files)
		}
	case "Size", "size":
		if sFlag.Desc {
			sort.SizeDesc(files)
		} else {
			sort.SizeAsc(files)
		}
	case "Create", "create":
		if sFlag.Desc {
			sort.CreateDataDesc(files)
		} else {
			sort.CreateDataAsc(files)
		}
	case "Put", "put":
		if sFlag.Desc {
			sort.PutDateDesc(files)
		} else {
			sort.PutDateAsc(files)
		}
	default:
		fmt.Println("Column: Name, Size, Create, Put. Default is Name ascending sort.")
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
