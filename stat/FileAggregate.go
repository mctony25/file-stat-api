package stat

import (
	"file-stat/response"
	"fmt"
	"os"
	"path/filepath"
)

type FileAggregate struct{}

func (fa FileAggregate) AggregateFromList(fileList []os.FileInfo) (response.FileStat) {

	var size int64
	var fileCount int16
	var elemCount int16
	for _, file := range fileList {
		if !file.IsDir() {
			size += file.Size()
			fileCount += 1
		}

		elemCount += 1
	}

	return response.FileStat{TotalElement: elemCount, TotalSize: size, TotalFile: fileCount}
}

func (fa FileAggregate) Aggregate(directoryName string, recursive bool) (response.FileStat, error) {
	var size int64
	var fileCount int16
	var elemCount int16

	if _, err := os.Stat(directoryName); err != nil {
		fmt.Println(fmt.Sprintf(" #Error : %s", err.Error()))
		return response.FileStat{TotalElement: 0, TotalSize: 0, TotalFile: 0}, err
	}
	err := filepath.Walk(directoryName, func(_ string, fileInfo os.FileInfo, err error) error {
		elemCount += 1
		fmt.Println(fmt.Sprintf("File name: %s", fileInfo.Name()))
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			if _, err := os.Stat(directoryName + "/" + fileInfo.Name()); err == nil {
				size += fileInfo.Size()
				fileCount += 1
			}
		} else {
			if recursive {
				fileStat, _ := fa.Aggregate(directoryName + "/" + fileInfo.Name(), recursive)
				size += fileStat.TotalSize
				fileCount += fileStat.TotalFile
				elemCount += fileStat.TotalElement
			}
		}

		return err
	})

	return response.FileStat{TotalElement: elemCount, TotalSize: size, TotalFile: fileCount}, err
}
