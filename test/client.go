// +build ignore

package main

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-10-27

import (
	"fmt"
	"strconv"
	"time"

	"github.com/belfinor/uport"
)

func main() {

	c, err := uport.NewClient("127.0.0.1:6000")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	for {

		<-time.After(time.Second)

		txt := strconv.FormatInt(time.Now().UnixNano(), 10)
		fmt.Println("> " + txt)

		c.Send([]byte(txt))

		data, err := c.Read()
		if err != nil {
			continue
		}

		fmt.Println("< " + string(data))

	}
}
