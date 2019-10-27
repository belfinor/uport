package uport

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-10-27

import (
	"fmt"
	"net"
	"sync"

	"github.com/belfinor/log"
)

type Handler func(msg []byte) []byte

func Server(addr string, fn Handler) error {

	laddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	ln, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return err
	}

	pool := sync.Pool{New: func() interface{} { return make([]byte, 10240) }}

	for {

		buffer := pool.Get().([]byte)

		n, addr, err := ln.ReadFromUDP(buffer)
		if err != nil {
			pool.Put(buffer)
			continue
		}

		go func() {

			defer func() {

				if r := recover(); r != nil {
					log.Error(fmt.Sprint(r))
				}

			}()

			defer pool.Put(buffer)

			msg := buffer[:n]

			msg = fn(msg)

			if len(msg) > 0 {
				ln.WriteToUDP(msg, addr)
			}

		}()

	}

	return nil

}
