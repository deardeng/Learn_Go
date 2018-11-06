package main

import (
	"day7/example1/balance"
	"fmt"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
}

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}

	crcTable := crc32.MakeTable(crc32.Castagnoli)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]

	return

}
