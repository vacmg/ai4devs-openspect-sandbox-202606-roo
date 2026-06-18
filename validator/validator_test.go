package validator

import "testing"

func TestValidEmails(t *testing.T) {
	valid := []string{
		// Casos típicos
		"simple@example.com",
		"very.common@example.com",
		"disposable.style.email.with+symbol@example.com",
		"other.email-with-hyphen@example.com",
		"fully-qualified-domain@example.com",
		"user.name+tag+sorting@example.com",
		"x@example.com",
		"example-indeed@strange-example.com",
		"test/test@test.com",
		// Caracteres especiales permitidos en atext
		"!#$%&'*+-/=?^_`{|}~@example.com",
		// Quoted local-part
		`"john..doe"@example.com`,
		`" "@example.com`,
		`"very.(),:;<>[]\".VERY.\"very@\\ \"very\".unusual"@strange.example.com`,
		// Subdominios
		"user@mail.subdomain.example.com",
		// Dígitos en dominio
		"user@123.example.com",
		// Local-part con dígitos al inicio
		"1234567890@example.com",
	}
	for _, e := range valid {
		t.Run(e, func(t *testing.T) {
			if !IsValid(e) {
				t.Errorf("se esperaba VÁLIDO: %q", e)
			}
		})
	}
}

func TestInvalidEmails(t *testing.T) {
	invalid := []string{
		"",
		"plainaddress",
		"@no-local-part.com",
		"no-at-sign.com",
		"user@",
		"user@@example.com",
		"user@.com",
		"user@example",         // sin TLD (sin punto en dominio)
		"user@-example.com",    // label empieza con guion
		"user@example-.com",    // label termina con guion
		"user@exa_mple.com",    // guion bajo no permitido en dominio
		".user@example.com",    // empieza con punto
		"user.@example.com",    // termina con punto
		"us..er@example.com",   // dos puntos consecutivos
		"user name@example.com", // espacio sin comillas
		"user@exam ple.com",
		"user@example..com",    // doble punto en dominio
		`"unterminated@example.com`,
		`"@example.com`,
		"a@b",                  // dominio sin punto
		// Demasiado largo (local > 64)
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com",
	}
	for _, e := range invalid {
		t.Run(e, func(t *testing.T) {
			if IsValid(e) {
				t.Errorf("se esperaba INVÁLIDO: %q", e)
			}
		})
	}
}

