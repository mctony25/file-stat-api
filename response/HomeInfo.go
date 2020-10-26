package response

import "time"

type HomeInfo struct {
	AppVersion string    `json:"app_version"`
	Timestamp  time.Time `json:"timestamp"`
}
