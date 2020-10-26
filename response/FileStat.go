package response

type FileStat struct {
	TotalSize    int64 `json:"total_size"`
	TotalElement int16 `json:"total_elements"`
	TotalFile    int16 `json:"total_files"`
}
