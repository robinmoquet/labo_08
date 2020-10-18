package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/robinmoquet/labo_08/zip"
)

/**
* Commande pour zipper des dossier
* 1er param => folder path
* 2eme param => output path
 */
func main() {
	// help param
	help := flag.Bool("help", false, "param help")
	flag.Parse()
	if *help {
		fmt.Println("First param is the folder path")
		fmt.Println("And second param is the output path file")
		fmt.Println("Ex : go-zip /your/path/to/compress /your/output/path/compress.zip")
	}

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("You must have two param, use \"go-zip -help\" for more infos")
		return
	}
	folderPath := args[0]
	outputFile := args[1]

	if err := zip.Zipper(folderPath, outputFile); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v has been created !\n", outputFile)
}
