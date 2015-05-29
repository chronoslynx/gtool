package exe

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

type MachHdr32 struct {
	Magic                              uint32
	CpuType                            uint32
	CpuSubType                         uint32
	Filetype, Ncmds, SizeOfCmds, Flags uint32
}

func (m MachHdr32) Print() {
	fmt.Println("Magic\t\tCpu Type\t\tCpu Subtype\t\tFile type\t\tNcmds\t\tSizeof cmds\t\tflags")
	fmt.Printf("0x%x\t\t%s\t\t%s\t\t%s\t\t%s\t\t%s\t\t0x%x\n",
		m.Magic,
		strconv.Itoa(int(m.CpuType)),
		strconv.Itoa(int(m.CpuSubType)),
		strconv.Itoa(int(m.Filetype)),
		strconv.Itoa(int(m.Ncmds)),
		strconv.Itoa(int(m.SizeOfCmds)),
		m.Flags)
}

type MachHdr64 struct {
	MachHdr32
	reserved uint32
}

func LoadMachHdr(file string) (e ExeFmt, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	m := MachHdr32{}
	err = binary.Read(f, binary.LittleEndian, &m)
	if err != nil {
		return
	}
	return m, nil
}
