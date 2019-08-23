package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func falaLoro(input string) string {
	input = strings.TrimRight(input, "\n")
	lang := map[string]string{
		"esquerda": "ingles",
		"as duas":  "caiu",
		"direita":  "frances",
		"nenhuma":  "portugues",
	}
	return lang[input]

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for true {
		perna, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Println(falaLoro(perna))
	}
}
