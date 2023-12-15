package utils

import (
	"fmt"
)

// ConvertSizeToStr returns a human-readable size value.
func ConvertSizeToStr(size int64) string {
	// byte to kb = float(n)/(1<<10)
	// byte to mb = float(n)/(1<<20)
	// byte to gb = float(n)/(1<<30)

	switch {
	case size > (1 << 30):
		return fmt.Sprintf("%.2f Go", float64(size)/(1<<30))
	case size > (1 << 20):
		return fmt.Sprintf("%.2f Mb", float64(size)/(1<<20))
	case size > (1 << 10):
		return fmt.Sprintf("%.2f Kb", float64(size)/(1<<10))
	default:
		return fmt.Sprintf("%d Bytes", size)
	}
}
