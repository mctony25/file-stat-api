package controller

import (
	"file-stat/response"
	"file-stat/stat"
	"net/http"
	"time"
)

type FileController struct{}

func (fc FileController) GetFileList(w http.ResponseWriter, r *http.Request) {

	directoryName := "/tmp/"
	dirPath := r.FormValue("dirPath")
	if "" != dirPath {
		directoryName = dirPath
	}
	stats := response.FileStat{
		TotalSize:    0,
		TotalElement: 0,
		TotalFile:    0,
	}
	fileLister := stat.FileLister{}
	files, err := fileLister.List(directoryName)
	if err != nil {
		errorResponse := response.Error{ErrorCode: response.NoSuchDirectory, Message: err.Error()}
		(response.JsonResponse{}).Send(w, errorResponse, http.StatusBadRequest)

		return
	}

	var fileList []response.FileInfo
	for _, file := range files {
		isDir := file.IsDir()
		fileInfo := response.FileInfo{
			Name:             file.Name(),
			Size:             file.Size(),
			IsDirectory:      isDir,
			LastModification: file.ModTime().Format(time.RFC3339),
		}

		fileList = append(fileList, fileInfo)
	}

	statAgg := stat.FileAggregate{}
	stats = statAgg.AggregateFromList(files)

	fileListResponse := response.FileList{
		Files: fileList,
		Stats: stats,
	}

	(response.JsonResponse{}).Send(w, fileListResponse, http.StatusOK)

	return
}
