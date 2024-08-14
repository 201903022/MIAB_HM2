package Commands

import (
	structures "MIAB_HM2/Structures"
	"fmt"
)

//leer Archivo y mbr e imprimir
func CommandRep() error {
	path := "/home/jonathan/MIAB_2S/TAREA2/MIAB_HM2/Discos/disco.asdj"
	//verificar si existe el 	archivo
	fmt.Println("******************************************************************")
	fmt.Println("**********************REPORTE DE DISCO***************************")
	//fmt.Println("Path: ", path)
	mbr := &structures.MBR{}
	err := mbr.DeserializeMBR(path)
	if err != nil {
		return err
	}
	//fmt.Println("MBR: ", mbr)
	mbr.Print()
	fmt.Println("******************************************************************")

	return nil
}
