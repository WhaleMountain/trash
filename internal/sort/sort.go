package sort

import (
	"os"
	"sort"
	"time"
	"trash/internal/rename"
)

// PutDateAsc 昇順
func PutDateAsc(files []os.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		var fInfoI = make(map[string]string, 2)
		var fInfoJ = make(map[string]string, 2)
		rename.Decode(files[i].Name(), &fInfoI)
		rename.Decode(files[j].Name(), &fInfoJ)

		putDateI, _ := time.Parse("2006/01/02 15:04:05", fInfoI["PutDate"])
		putDateJ, _ := time.Parse("2006/01/02 15:04:05", fInfoJ["PutDate"])

		return putDateI.Unix() < putDateJ.Unix()
	})
}

// PutDateDesc 降順
func PutDateDesc(files []os.FileInfo) {
	// PutDate Asc (昇順)
	sort.Slice(files, func(i, j int) bool {
		var fInfoI = make(map[string]string, 2)
		var fInfoJ = make(map[string]string, 2)
		rename.Decode(files[i].Name(), &fInfoI)
		rename.Decode(files[j].Name(), &fInfoJ)

		putDateI, _ := time.Parse("2006/01/02 15:04:05", fInfoI["PutDate"])
		putDateJ, _ := time.Parse("2006/01/02 15:04:05", fInfoJ["PutDate"])

		return putDateI.Unix() > putDateJ.Unix()
	})
}

// CreateDataAsc 昇順
func CreateDataAsc(files []os.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Unix() < files[j].ModTime().Unix()
	})
}

// CreateDataDesc 降順
func CreateDataDesc(files []os.FileInfo) {
	// PutDate Asc (昇順)
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Unix() > files[j].ModTime().Unix()
	})
}

// SizeAsc 昇順
func SizeAsc(files []os.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() < files[j].Size()
	})
}

// SizeDesc 降順
func SizeDesc(files []os.FileInfo) {
	// PutDate Asc (昇順)
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() > files[j].Size()
	})
}

// NameDesc 降順
func NameDesc(files []os.FileInfo) {
	// PutDate Asc (昇順)
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() > files[j].Name()
	})
}
