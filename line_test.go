/*
* File Name:	line_test.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gamil.com>
* Created:	2015-03-11
 */
package line_test

import (
	"fmt"
	"github.com/ochapman/line"
	"testing"
)

func TestMultiLine(t *testing.T) {
	lines, err := line.GetFileLine("testdata/file.txt")
	if err != nil {
		t.Errorf("GetFileLine() failed: %s\n", err)
	}
	for i := 1; i < len(lines); i++ {
		fmt.Printf("line: %d %s\n", i, lines[i])
	}
}
