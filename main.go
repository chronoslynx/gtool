package main

import "github.com/chronoslynx/gtool/exe"

func main() {
	e, err := exe.LoadMachHdr("/bin/ls")
	if err != nil {
		panic(err)
	}
	e.Print()
}
