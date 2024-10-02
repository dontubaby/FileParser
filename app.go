package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ElementConverter(el1, el2 string) (n1, n2 int, err error) {
	n1, err = strconv.Atoi(el1)
	if err != nil {
		log.Fatalf("Error! Cant convert first element %s", err)
		return -1, -1, err
	}
	n2, err = strconv.Atoi(el2)
	if err != nil {
		log.Fatalf("Error! Cant convert first element %s", err)
		return -1, -1, err
	}
	return n1, n2, nil
}

func FileParser(inputFile, outputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error! Cant open input file", err)
	}
	output, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	match := regexp.MustCompile(`[1-9]{1,}[\+\-\*\/][1-9]{1,}`)
	matchPlus := regexp.MustCompile(`([\+])`)
	matchMinus := regexp.MustCompile(`([\-])`)
	matchMultiply := regexp.MustCompile(`([\*])`)
	matchDiv := regexp.MustCompile(`([\/])`)

	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		var res1 int
		var result, res2 string
		if match.MatchString(scanner.Text()) {
			data := strings.Trim(scanner.Text(), "=?")
			switch {
			case matchPlus.MatchString(data):
				elem1, elem2, _ := strings.Cut(data, "+")
				n1, n2, err := ElementConverter(elem1, elem2)
				if err != nil {
					log.Fatalf("Error! Cant conver element %s", err)
				}
				res1 = n1 + n2
				res2 = strconv.Itoa(res1)
				result = elem1 + "+" + elem2 + "=" + res2 + "\n"
				writer.WriteString(result)
			case matchMinus.MatchString(data):
				elem1, elem2, _ := strings.Cut(data, "-")
				n1, n2, err := ElementConverter(elem1, elem2)
				if err != nil {
					log.Fatalf("Error! Cant convert element %s", err)
				}
				res1 = n1 - n2
				res2 = strconv.Itoa(res1)
				result = elem1 + "-" + elem2 + "=" + res2 + "\n"
				writer.WriteString(result)
				fmt.Println("result", result)
			case matchMultiply.MatchString(data):
				elem1, elem2, _ := strings.Cut(data, "*")
				n1, n2, err := ElementConverter(elem1, elem2)
				if err != nil {
					log.Fatalf("Error! Cant convert element %s", err)
				}
				res1 = n1 * n2
				res2 = strconv.Itoa(res1)
				result = elem1 + "*" + elem2 + "=" + res2 + "\n"
				writer.WriteString(result)
			case matchDiv.MatchString(data):
				elem1, elem2, _ := strings.Cut(data, "/")
				n1, n2, err := ElementConverter(elem1, elem2)
				if err != nil {
					log.Fatalf("Error! Cant convert element %s", err)
				}
				res1 = n1 / n2
				res2 = strconv.Itoa(res1)
				result = elem1 + "/" + elem2 + "=" + res2 + "\n"
				writer.WriteString(result)
			}
			writer.Flush()
		}
	}

}

func main() {
	FileParser("input.txt", "output.txt")
}
