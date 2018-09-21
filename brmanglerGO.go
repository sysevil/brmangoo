package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

func wordSpecial(word string, size int, specialArray []string, min int, max int) {

	for _, value := range specialArray {
		concatPassword := word + value
		reverseConcatPassword := value + word
		checkPasswdSize(concatPassword, min, max)
		checkPasswdSize(reverseConcatPassword, min, max)
		for _, value2 := range specialArray {
			concatPasswordDeep := concatPassword + word + value2
			checkPasswdSize(concatPasswordDeep, min, max)
		}
	}

}

func wordSpecialNum(word string, specialArray []string, min int, max int) {
	nums := 9999
	for _, value := range specialArray {
		for num := 1; num <= nums; num++ {
			numWordSpecial := strconv.Itoa(num) + word + value
			reverseNumWordSpecial := value + word + strconv.Itoa(num)
			middleWordSpecial := word + value + strconv.Itoa(num)
			checkPasswdSize(numWordSpecial, min, max)
			checkPasswdSize(reverseNumWordSpecial, min, max)
			checkPasswdSize(middleWordSpecial, min, max)
		}
	}
}

func main() {

	file := flag.String("file", "", "file with names --file /path/to/file")
	minimumPwdSize := flag.Int("min", 8, "minimum password size")
	maxPwdSize := flag.Int("max", 12, "max passoword size")
	special := flag.Bool("s", false, "DISABLE passwords with special characters")
	//leet := flag.Bool("leet", true, "disable leet speak (e.g. p@55w0rd)")
	//capitalize := flag.Bool("capitalize", true, "disable capitalize words like (eg Passwords)")
	//calcOutPut := flag.Bool("calc", false, "Calculate the quantity of outputs")
	//upcase := flag.Bool("upcase", true, "disable uppercase")
	//dates := flag.Bool("dates", true, "disable dates (password@022102009)")
	//twoLetters := flag.Bool("twoletters", false, "enable two letters passwords ( af212301031)")
	//insane := flag.Bool("insane", true, "Use *ALL* wordlists to create passwords")
	flag.Parse()

	SPC := []string{"!", "@", "$", "%", "*"}

	start := time.Now()
	fmt.Println(start.Format("3:04PM"))
	logo()

	if validadeInputFile(*file) == true {

		inFile, err := os.Open(*file)
		if err != nil {
			log.Fatal("[+] Error open a file")
		}
		defer inFile.Close()

		if *special {
			sizeSpecial := len(SPC)
			scanner := bufio.NewScanner(inFile)
			for scanner.Scan() {
				password := scanner.Text()
				wordSpecial(password, sizeSpecial, SPC, *minimumPwdSize, *maxPwdSize)
				wordSpecialNum(password, SPC, *minimumPwdSize, *maxPwdSize)
			}
		}
	}

}
