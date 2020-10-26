package response

const (
	NoSuchDirectory  = 1001
	PermissionDenied = 1002
)

type Error struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}
