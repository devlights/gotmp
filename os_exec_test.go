package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestOsExec(t *testing.T) {
	// Arrange
	buf := new(bytes.Buffer)

	c := exec.Command("go", "env", "GOPATH")
	c.Stdout = buf

	// Act
	err := c.Run()

	// Assert
	if err != nil {
		t.Error(err)
	}

	// Output
	// t.Log() の出力は、失敗するか -v オプションをつけた場合しか出力されない
	t.Log(buf.String())
}
