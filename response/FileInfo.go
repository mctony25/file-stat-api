package response

type FileInfo struct {
	Name             string `json:"name"`
	Size             int64  `json:"size"`
	IsDirectory      bool   `json:"is_directory"`
	LastModification string `json:"last_modification"`
}
