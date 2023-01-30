package models

import (
	"context"
	"embed"
	"fmt"
	"os"
)

type ConfBox struct {
	Ctx context.Context
	Fs  embed.FS
	Dir string
}

func (c ConfBox) Get(filename string) (bs []byte) {
	if filename == "" {
		return
	}

	f := fmt.Sprintf("%s%s%s", c.Dir, string(os.PathSeparator), filename)
	var err error
	bs, err = os.ReadFile(f)
	if err != nil {
		fmt.Printf("read file %s from system failed \n", f)
		err = nil
	}
	if len(bs) == 0 {
		bs, err = c.Fs.ReadFile(f)
		if err != nil {
			fmt.Printf("read file %s from embed failed \n", f)
		}
	}
	return
}
