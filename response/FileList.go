package response

type FileList struct {
	Files []FileInfo `json:"files"`
	Stats FileStat `json:"stats"`
}
