// Copyright 2017 The go-xdgbasedir Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package home

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

var usrHome = os.Getenv("HOME")

// Dir detects and returns the user home directory.
func Dir() string {
	// At first, Check the $HOME environment variable
	if usrHome != "" {
		return usrHome
	}

	// Fallback with sh pwd commmand
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "cd && pwd")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return ""
	}
	usrHome = strings.TrimSpace(stdout.String())

	return usrHome
}
