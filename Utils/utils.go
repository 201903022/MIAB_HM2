package utils

import (
	"encoding/binary"
	"fmt"
)

func Int32ToBytes(n int32) [4]byte {
	var buf [4]byte
	binary.LittleEndian.PutUint32((buf[:]), uint32(n))
	return buf
}

func Float64ToBytes(n float64) [4]byte {
	var buf [4]byte
	binary.LittleEndian.PutUint64((buf[:]), uint64(n))
	return buf
}

func ConvertToBytes(size int, unit string) (int, error) {
	switch unit {
	case "B":
		return size, nil
	case "K":
		return size * 1024, nil
	case "M":
		return size * 1024 * 1024, nil
	case "G":
		return size * 1024 * 1024 * 1024, nil
	default:
		return 0, fmt.Errorf("Invalid unit: %s", unit)
	}
}
