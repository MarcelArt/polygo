package fiber

import (
	"fmt"
	"os"
)

const timeEnumFileTemplate = `// This file is auto generated by polygo
package enums

import "time"

const (
	Day   = time.Hour * 24
	Month = Day * 30
)
`

func (fp FiberProject) createTimeEnumFiles() error {
	if err := os.Mkdir(fmt.Sprintf("%s/enums", fp.Directory), 0755); err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("%s/enums/time.go", fp.Directory), []byte(timeEnumFileTemplate), 0644)
}
