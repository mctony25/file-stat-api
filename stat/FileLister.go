package stat

import (
	"io/ioutil"
	"os"
	"sort"
)

type FileLister struct{}

func (fl FileLister) List(directoryName string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(directoryName)

	if err != nil {
		return nil, err
	}

	fl.sortBySize(files)

	return files, nil
}

func (fl FileLister) sortBySize(files []os.FileInfo) {

	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() > files[j].Size()
	})
}
