package git

import (
	"bytes"

	"github.com/magefile/mage/sh"
)

func SHA() (string, error) {

	b := new(bytes.Buffer)
	if _, err := sh.Exec(
		map[string]string{},
		b,
		nil,
		"git",
		"rev-parse",
		"--short",
		"HEAD",
	); err != nil {
		return "", err
	}
	return string(b.Bytes()), nil
}
