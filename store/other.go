package store

import (
	"github.com/snksoft/crc"
)

func CalcHash(data []byte) (hash string, err error) {
	ccittCrc := crc.CalculateCRC(crc.CCITT, data)
	return string(ccittCrc), nil
}
