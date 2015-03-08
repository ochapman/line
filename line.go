/*
* File Name:	line.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2015-03-08
 */

package line

import (
	"bufio"
	"os"
)

type Line map[int][]byte

func GetFileLine(file string) (Line, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	lines := make(Line)
	for i := 0; ; i++ {
		line, _, err := r.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}
		lines[i] = line
	}
	return lines, nil
}
