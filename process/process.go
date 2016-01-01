// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func run(cmd *Cmd, wait_time int) (stdout []byte, stderr []byte, err error) {
	exited, sleeproutine_error := make(chan bool, 1), make(chan error, 1)
	ran_long_enough := false

	go func() {
		// This go routine waits for the process to exit, or `wait_time`,
		// whichever comes first. If `wait_time` expires, `cmd` is killed.
		select {
		case <-exited:
			break
		case <-time.After(wait_time):
			ran_long_enough = true
			err := cmd.Process.Kill()
			if err != nil && err.Error() != "os: process already finished" {
				sleeproutine_error <- err
			}
		}
		sleeproutine_error <- nil
	}()

	err = cmd.Run()
	exited <- true

	if !ran_long_enough && err != nil {
		t.Fatalf("Process died unexpectedly: %q", err)
	}

	// Sync with sleep routine
	if err = <-sleeproutine_error; err != nil {
		t.Fatalf("Unexpected error: %q", err)
	}
}
