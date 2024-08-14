package structures

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type MBR struct {
	Mbr_size           int32
	Mbr_creation_date  float32
	Mbr_disk_signature int32
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

func (mbr *MBR) DeserializeMBR(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	mbrSize := binary.Size(mbr)
	if mbrSize <= 0 {
		return fmt.Errorf("Error getting MBR size: %v", err)
	}

	buffer := make([]byte, mbrSize)
	_, err = file.Read(buffer)
	if err != nil {
		return fmt.Errorf("Error reading MBR: %v", err)
	}

	mbrReader := bytes.NewReader(buffer)
	err = binary.Read(mbrReader, binary.BigEndian, mbr)
	if err != nil {
		return fmt.Errorf("Error deserializing MBR: %v", err)
	}
	return nil

	return nil
}

func (mbr *MBR) Print() {

	creationTime := time.Unix(int64(mbr.Mbr_creation_date), 0)

	fmt.Printf("Size: %d\n", mbr.Mbr_size)
	fmt.Printf("Creation Date: %s\n", creationTime)
	fmt.Printf("Disk Signature: %d\n", mbr.Mbr_disk_signature)
}
