// Package validator implementa validación de direcciones de email según RFC 5322
// (sintaxis de addr-spec). Soporta:
//
//   - local-part como dot-atom o quoted-string
//   - domain como dot-atom con etiquetas LDH (RFC 1035) — al menos 2 etiquetas
//   - Límites: local-part <= 64 octetos, longitud total <= 254 octetos
//
// No soporta domain-literals ([IPv4/IPv6]) ni comentarios CFWS, que son partes
// poco comunes y desaconsejadas por RFC 5321/5322bis para uso práctico.
package validator

import "strings"

// atext son los caracteres permitidos sin comillas en local-part (RFC 5322 §3.2.3).
func isAtext(c byte) bool {
	switch {
	case c >= 'a' && c <= 'z':
		return true
	case c >= 'A' && c <= 'Z':
		return true
	case c >= '0' && c <= '9':
		return true
	}
	switch c {
	case '!', '#', '$', '%', '&', '\'', '*', '+', '-', '/', '=',
		'?', '^', '_', '`', '{', '|', '}', '~':
		return true
	}
	return false
}

// qtext: caracteres permitidos sin escape dentro de "..." (RFC 5322 §3.2.4).
// Excluye DQUOTE ("), backslash (\) y controles; permitimos imprimibles ASCII.
func isQtext(c byte) bool {
	if c == '"' || c == '\\' {
		return false
	}
	return c >= 0x20 && c < 0x7F
}

// IsValid devuelve true si s es una dirección de email sintácticamente válida.
func IsValid(s string) bool {
	if s == "" || len(s) > 254 {
		return false
	}

	local, domain, ok := splitAtLastUnquotedAt(s)
	if !ok {
		return false
	}
	if len(local) == 0 || len(local) > 64 {
		return false
	}
	if !validLocalPart(local) {
		return false
	}
	return validDomain(domain)
}

// splitAtLastUnquotedAt separa local@domain teniendo en cuenta que la local-part
// puede ir entre comillas y contener '@'. Devuelve false si no hay exactamente
// una @ que actúe como separador.
func splitAtLastUnquotedAt(s string) (local, domain string, ok bool) {
	inQuotes := false
	escaped := false
	atIdx := -1
	for i := 0; i < len(s); i++ {
		c := s[i]
		if escaped {
			escaped = false
			continue
		}
		if inQuotes {
			if c == '\\' {
				escaped = true
				continue
			}
			if c == '"' {
				inQuotes = false
			}
			continue
		}
		switch c {
		case '"':
			inQuotes = true
		case '@':
			if atIdx != -1 {
				// Más de una @ fuera de comillas → inválido
				return "", "", false
			}
			atIdx = i
		}
	}
	if inQuotes || escaped || atIdx <= 0 || atIdx == len(s)-1 {
		return "", "", false
	}
	return s[:atIdx], s[atIdx+1:], true
}

func validLocalPart(s string) bool {
	if s[0] == '"' {
		return validQuotedString(s)
	}
	return validDotAtom(s, isAtext)
}

// validDotAtom valida 1*char *("." 1*char) sin puntos al inicio/fin ni dobles.
func validDotAtom(s string, isChar func(byte) bool) bool {
	if len(s) == 0 || s[0] == '.' || s[len(s)-1] == '.' {
		return false
	}
	prevDot := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '.' {
			if prevDot {
				return false
			}
			prevDot = true
			continue
		}
		if !isChar(c) {
			return false
		}
		prevDot = false
	}
	return true
}

func validQuotedString(s string) bool {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return false
	}
	body := s[1 : len(s)-1]
	for i := 0; i < len(body); i++ {
		c := body[i]
		if c == '\\' {
			// quoted-pair: \X donde X es cualquier carácter visible o WSP
			if i+1 >= len(body) {
				return false
			}
			next := body[i+1]
			if next < 0x20 || next >= 0x7F {
				if next != '\t' {
					return false
				}
			}
			i++
			continue
		}
		if c == '"' {
			return false
		}
		if !isQtext(c) && c != ' ' && c != '\t' {
			return false
		}
	}
	return true
}

// validDomain exige al menos dos etiquetas separadas por '.', cada una LDH
// (letra/dígito/guion), sin guiones al inicio o al final, longitud 1..63,
// y longitud total del dominio <= 253.
func validDomain(s string) bool {
	if len(s) == 0 || len(s) > 253 {
		return false
	}
	labels := strings.Split(s, ".")
	if len(labels) < 2 {
		return false
	}
	for _, lbl := range labels {
		if !validLabel(lbl) {
			return false
		}
	}
	return true
}

func validLabel(lbl string) bool {
	if len(lbl) == 0 || len(lbl) > 63 {
		return false
	}
	if lbl[0] == '-' || lbl[len(lbl)-1] == '-' {
		return false
	}
	for i := 0; i < len(lbl); i++ {
		c := lbl[i]
		switch {
		case c >= 'a' && c <= 'z':
		case c >= 'A' && c <= 'Z':
		case c >= '0' && c <= '9':
		case c == '-':
		default:
			return false
		}
	}
	return true
}

