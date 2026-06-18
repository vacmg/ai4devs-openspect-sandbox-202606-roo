// Comando emailcheck: lee una dirección de email desde STDIN y termina con
// código de salida 0 si es válida según RFC 5322 o 1 si no lo es.
//
// Uso:
//
//	echo "user@example.com" | emailcheck ; echo $?
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"emailcheck/validator"
)

func main() {
	os.Exit(run(os.Stdin, os.Stderr))
}

func run(in io.Reader, errOut io.Writer) int {
	data, err := io.ReadAll(bufio.NewReader(in))
	if err != nil {
		fmt.Fprintln(errOut, "error leyendo stdin:", err)
		return 1
	}
	// Aceptamos un email por entrada; recortamos espacios y saltos finales.
	email := strings.TrimRight(string(data), "\r\n")
	email = strings.TrimSpace(email)

	if validator.IsValid(email) {
		return 0
	}
	return 1
}

