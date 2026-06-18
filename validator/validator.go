// Package validator valida direcciones de email usando la librería estándar
// de Go `net/mail`, que implementa el parser de direcciones de RFC 5322.
//
// Se exige que la entrada sea una addr-spec pura (sin display-name ni
// comentarios CFWS), por lo que verificamos que el resultado de ParseAddress
// coincide con la entrada y no incluye nombre.
package validator

import (
	"net/mail"
	"strings"
)

// IsValid devuelve true si s es una dirección de email válida según RFC 5322
// usando el parser de net/mail. Adicionalmente:
//
//   - Rechaza formatos con display-name (`Nombre <addr@host>`).
//   - Exige al menos un punto en el dominio (descarta `user@host`, que el
//     parser acepta pero no es una dirección enrutable en la práctica).
func IsValid(s string) bool {
	if s == "" {
		return false
	}
	addr, err := mail.ParseAddress(s)
	if err != nil {
		return false
	}
	if addr.Name != "" || addr.Address != s {
		return false
	}
	at := strings.LastIndex(s, "@")
	if at < 0 {
		return false
	}
	domain := s[at+1:]
	if !strings.Contains(domain, ".") {
		return false
	}
	return true
}
