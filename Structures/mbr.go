package structures

import (
	"encoding/binary"
	"fmt"
	"os"
)

type MBR struct {
	Mbr_size           int32
	Mbr_creation_date  float32
	Mbr_disk_signature int32
	Mbr_partition      [4]Partition
	Mbr_fit            [1]byte
}

func (mbr *MBR) GetSize() int32 {
	return mbr.Mbr_size
}

func (mbr *MBR) SerializeMBR(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	err = binary.Write(file, binary.BigEndian, mbr)
	if err != nil {
		return fmt.Errorf("Error writing MBR: %v", err)
	}

	return nil
}
