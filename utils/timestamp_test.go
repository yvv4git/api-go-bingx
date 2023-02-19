package utils_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	bingx "github.com/yvv4git/api-go-bingx/utils"
)

func Test_currentTimestamp(t *testing.T) {
	timeStamp := bingx.CurrentTimestamp()
	assert.Equal(t, 13, len(fmt.Sprintf("%d", timeStamp)))
}
