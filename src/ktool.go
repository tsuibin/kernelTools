package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	makeFile, err := os.Create("Makefile")

	if err != nil {
		fmt.Println("Create Makefile error")
		os.Exit(1)

	}

	makeFile.WriteString("\nifneq\t($(KERNELRELEASE),)")
	makeFile.WriteString("\nobj-m\t:= ")

	var filename string
	for i := 1; i < len(os.Args); i++ {

		filename = os.Args[i]
		filename = strings.Replace(filename, ".c", ".o ", 1)
		makeFile.WriteString(filename)
	}

	if len(os.Args) == 1 {
		dir, _ := os.Open(".")
		files, _ := dir.Readdir(0)
		for _, f := range files {
			if !f.IsDir() {
				filename = f.Name()
				if strings.Contains(filename, ".c") {
					filename = strings.Replace(filename, ".c", ".o ", 1)
					makeFile.WriteString(filename)
				}
			}
		}
	}

	makeFile.WriteString("\n")
	makeFile.WriteString("\nelse")
	makeFile.WriteString("\nKDIR\t:= /lib/modules/$(shell uname -r)/build")
	makeFile.WriteString("\nPWD\t:= $(shell pwd)")
	makeFile.WriteString("\n")
	makeFile.WriteString("default:\t")
	makeFile.WriteString("\n\t$(MAKE) -C $(KDIR) SUBDIRS=$(PWD) modules ")
	makeFile.WriteString("\n\trm -r -f .tmp_versions *.mod.c .*.cmd *.o *~ *.order ")
	makeFile.WriteString("Modules.symvers ")
	makeFile.WriteString("*.symvers ")
	makeFile.WriteString("\n")
	makeFile.WriteString("clean:\t")
	makeFile.WriteString("\n\trm -r -f *.ko mmake *.order ")
	makeFile.WriteString("\n")
	makeFile.WriteString("\nendif")
	makeFile.WriteString("\n\n")

	mk := exec.Command("ls -l")
	mk.Start()

}
