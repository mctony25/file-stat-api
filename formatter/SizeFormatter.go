package formatter

import (
	"fmt"
	"math"
)

type SizeFormatter struct{}

func (sf SizeFormatter) Format(size int64) string {
	convertedSize := float64(size)
	switch {
		// Terabytes
		case convertedSize > math.Pow(1024, 4):
			return fmt.Sprintf("%.2fTB", convertedSize / math.Pow(1024, 4))
		// Gigabytes
		case convertedSize > math.Pow(1024, 3):
			return fmt.Sprintf("%.2fGB", convertedSize / math.Pow(1024, 3))
		// Megabytes
		case convertedSize > math.Pow(1024, 2):
			return fmt.Sprintf("%.2fMB", convertedSize / math.Pow(1024, 2))
		// Kilobytes
		case convertedSize > 1024.0:
			return fmt.Sprintf("%.2fKB", convertedSize / 1024.0)
		// Bytes
		default:
			return fmt.Sprintf("%dB", size)
	}
}
