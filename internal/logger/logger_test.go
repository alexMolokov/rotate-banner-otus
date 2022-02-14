package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	configApp "github.com/alexMolokov/rotate-banner-otus/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	cfg := &configApp.LoggerConf{Level: "error", Encoding: "json", Output: "stdout"}

	errorMsg := "error message"
	infoMsg := "info message"

	savedStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger, err := New(cfg)
	assert.Nil(t, err)

	logger.Error(errorMsg)
	logger.Info(infoMsg)

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)

	str := string(out)

	assert.Greater(t, len(out), 0)
	assert.False(t, strings.Contains(str, infoMsg))
	assert.True(t, strings.Contains(str, errorMsg))

	os.Stdout = savedStdout
	fmt.Println(str)
}
