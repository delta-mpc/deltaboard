package tests

import (
	"fmt"
	"strings"
	"testing"
)

func TestCardNo(t *testing.T) {
	cardNo := "453026199302018425"

	replaceStr := cardNo[4 : len(cardNo)-3]
	cardNo = strings.Replace(cardNo, replaceStr, strings.Repeat("*", len(replaceStr)), 1)

	fmt.Println("card info: ", cardNo)
}
