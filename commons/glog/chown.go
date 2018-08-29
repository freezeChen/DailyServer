// +build !linux

package glog

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
