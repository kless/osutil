// Copyright 2013 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package user

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kless/shutil/file"
)

const (
	USER     = "foo"
	SYS_USER = "bar"

	GROUP     = "g-foo"
	SYS_GROUP = "g-bar"
)

// Stores the ids at creating the groups.
var GID, SYS_GID int

// == Copy the system files before of be edited.

func init() {
	var err error

	if _USER_FILE, err = file.CopytoTemp(_USER_FILE, ""); err != nil {
		goto _error
	}
	if _GROUP_FILE, err = file.CopytoTemp(_GROUP_FILE, ""); err != nil {
		goto _error
	}
	if _SHADOW_FILE, err = file.CopytoTemp(_SHADOW_FILE, ""); err != nil {
		goto _error
	}
	if _GSHADOW_FILE, err = file.CopytoTemp(_GSHADOW_FILE, ""); err != nil {
		goto _error
	}

	return

_error:
	removeTempFiles()
	log.Fatal(err)
}

func removeTempFiles() {
	files, _ := filepath.Glob(filepath.Join(os.TempDir(), file.PREFIX_TEMP+"*"))

	for _, f := range files {
		if err := os.Remove(f); err != nil {
			log.Print(err)
		}
	}
}