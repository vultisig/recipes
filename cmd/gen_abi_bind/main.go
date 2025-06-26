package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	abiPath       = "abi"
	codegenFolder = "codegen"
	outPath       = path.Join("sdk", "evm", codegenFolder)
)

func main() {
	list, err := os.ReadDir(abiPath)
	if err != nil {
		panic(fmt.Errorf("os.ReadDir(abi): %w", err))
	}

	err = os.RemoveAll(outPath)
	if err != nil {
		panic(fmt.Errorf("os.RemoveAll(outPath): %w", err))
	}

	err = os.Mkdir(outPath, fs.ModePerm)
	if err != nil {
		panic(fmt.Errorf("os.Mkdir(outPath, fs.ModePerm): %w", err))
	}

	for _, entry := range list {
		noExt := strings.TrimSuffix(entry.Name(), ".json")
		outSubfolder := path.Join(outPath, strings.ToLower(noExt))

		err = os.Mkdir(outSubfolder, fs.ModePerm)
		if err != nil {
			panic(fmt.Errorf("os.Mkdir(outSubfolder, fs.ModePerm): %w", err))
		}

		var (
			abiFile = path.Join(abiPath, entry.Name())
			pkgName = strings.ToLower(noExt)
			outFile = path.Join(outSubfolder, pkgName+".go")
		)

		output, e := exec.Command(
			"abigen",
			"--v2",
			"--abi", abiFile,
			"--pkg", pkgName,
			"--out", outFile,
		).CombinedOutput()
		if e != nil {
			panic(fmt.Errorf("exec.Command: %s: %w", string(output), e))
		}
	}

	fmt.Println("codegen done to", outPath)
}
