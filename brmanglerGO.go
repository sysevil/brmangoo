package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func logo() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("██████╗ ██████╗    ███╗   ███╗ █████╗ ███╗   ██╗ ██████╗ ██╗     ███████╗██████╗")
	fmt.Println("██╔══██╗██╔══██╗   ████╗ ████║██╔══██╗████╗  ██║██╔════╝ ██║     ██╔════╝██╔══██╗")
	fmt.Println("██████╔╝██████╔╝   ██╔████╔██║███████║██╔██╗ ██║██║  ███╗██║     █████╗  ██████╔╝")
	fmt.Println("██╔══██╗██╔══██╗   ██║╚██╔╝██║██╔══██║██║╚██╗██║██║   ██║██║     ██╔══╝  ██╔══██╗")
	fmt.Println("██████╔╝██║  ██║██╗██║ ╚═╝ ██║██║  ██║██║ ╚████║╚██████╔╝███████╗███████╗██║  ██║")
	fmt.Println("╚═════╝ ╚═╝  ╚═╝╚═╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝")
	fmt.Println("Brazilian wordlist generator hu3hu3hu3")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func validadeInputFile(file string) bool {
	if file == "" {
		fmt.Println("[+] File is mandatory")
		os.Exit(1)
	}
	return true
}

func checkPasswdSize(password string, min int, max int) {
	if len(password) >= min && len(password) <= max {
		fmt.Printf("%s \n", strings.TrimSuffix(password, "\n"))
	}
}

func main() {
	file := flag.String("file", "", "file with names --file /path/to/file")
	minimumPwdSize := flag.Int("min", 8, "minimum password size")
	maxPwdSize := flag.Int("max", 12, "max passoword size")
	//leet := flag.Bool("leet", true, "disable leet speak (e.g. p@55w0rd)")
	//capitalize := flag.Bool("capitalize", true, "disable capitalize words like (eg Passwords)")
	//calcOutPut := flag.Bool("calc", false, "Calculate the quantity of outputs")
	//upcase := flag.Bool("upcase", true, "disable uppercase")
	//dates := flag.Bool("dates", true, "disable dates (password@022102009)")
	//special := flag.Bool("special", true, "disable passwords with special chars")
	//twoLetters := flag.Bool("twoletters", false, "enable two letters passwords ( af212301031)")
	//insane := flag.Bool("insane", true, "Use *ALL* wordlists to create passwords")
	flag.Parse()

	start := time.Now()
	fmt.Println(start.Format("3:04PM"))
	logo()

	if validadeInputFile(*file) == true {

		inFile, err := os.Open(*file)
		if err != nil {
			log.Fatal("[+] Error open a file")
		}
		defer inFile.Close()
		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			checkPasswdSize(scanner.Text(), *minimumPwdSize, *maxPwdSize)
		}
	}

}
