// +build ignore

package main

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-10-27

import (
	"github.com/belfinor/uport"
)

func main() {

	if err := uport.Server(":6000", func(in []byte) []byte {

		return in

	}); err != nil {
		panic(err)
	}

}
