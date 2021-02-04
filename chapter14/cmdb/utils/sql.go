package utils

import (
	"fmt"
	"strings"
)

func Like(q string) string {
	q = strings.TrimSpace(q)
	if q == "" {
		return ""
	}
	return fmt.Sprintf("%%%s%%", q)
}
