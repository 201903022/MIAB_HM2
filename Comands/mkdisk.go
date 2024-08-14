package Commands

import (
	structures "MIAB_HM2/Structures"
	utils "MIAB_HM2/Utils"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type MKDISK struct {
	size int
	unit string
	path string
}

var (
	size int
	unit string
	path string
)

func CommandMkdisk() error {
	fmt.Println("Creating disk")
	size = 5
	unit = "M"
	path = "/home/jonathan/MIAB_2S/TAREA2/MIAB_HM2/Discos/disco.asdj"

	fmt.Println("Size: ", size)
	fmt.Println("Unit: ", unit)
	fmt.Println("Path: ", path)

	/*
		sizeBytes, err := utils.ConvertToBytes(size, unit)
		if err != nil {
			return err
		}
		fmt.Println("Size in bytes: ", sizeBytes)
	*/

	error := createDisk()
	if error != nil {
		fmt.Println("Error creating disk: ", error)
		return error
	}

	fmt.Println("Disk created")

	return nil
}

func createDisk() error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory: ", err)
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}
	defer file.Close()
	sizeBytes, err := utils.ConvertToBytes(size, unit)
	buffer := make([]byte, sizeBytes)
	if _, err := file.Write(buffer); err != nil {
		fmt.Println("Error writing to file: ", err)
		return fmt.Errorf("Error writing to file: %v", err)
	}

	createMBR()
	return nil
}

func createMBR() {
	sizeBytes, err := utils.ConvertToBytes(size, unit)

	if err != nil {
		fmt.Println("Error converting to bytes: ", err)
		return
	}
	mbr := &structures.MBR{
		Mbr_size:           int32(sizeBytes),
		Mbr_creation_date:  float32(time.Now().Unix()),
		Mbr_disk_signature: rand.Int31(),
	}

	//fmt.Println("MBR: ", mbr)

	error := mbr.SerializeMBR(path)
	if error != nil {
		fmt.Println("Error serializing MBR: ", error)
		return
	}

	//mbr.Print()

}
