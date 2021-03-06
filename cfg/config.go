package cfg

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const middle = "=========\n"

type Cfg struct {
	ConfMap map[string]string
	strSet  string
}

func (c *Cfg) InitCfg(path string) {
	c.ConfMap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1 + 1 {
			c.strSet = strings.TrimSpace(s[n1 + 1 : n2])
			continue
		}

		if len(c.strSet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index + 1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strSet + middle + frist
		c.ConfMap[key] = strings.TrimSpace(second)
	}
}

func (c Cfg) Read(node, key string) string {
	key = node + middle + key
	v, found := c.ConfMap[key]
	if !found {
		return ""
	}
	return v
}