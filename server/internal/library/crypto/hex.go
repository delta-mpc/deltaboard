package crypto

import (
	"strings"
)

func Strip0x(hexStr string) string {
	if strings.HasPrefix(hexStr, "0x") {
		hexStr = hexStr[2:]
	}
	return hexStr
}
