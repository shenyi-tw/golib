package log

import (
	"fmt"
	"testing"

	logger "github.com/shenyi-tw/golib/log"
)

func TestGetArticles2(t *testing.T) {
	logger.Log("INFO", fmt.Sprintf("time %d", 123))
}
