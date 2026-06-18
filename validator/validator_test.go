package validator

import "testing"

// TestValidEmails comprueba direcciones que deben considerarse válidas.
func TestValidEmails(t *testing.T) {
	valid := []string{
		"simple@example.com",
		"very.common@example.com",
		"disposable.style.email.with+symbol@example.com",
		"other.email-with-hyphen@example.com",
		"fully-qualified-domain@example.com",
		"user.name+tag+sorting@example.com",
		"x@example.com",
		"example-indeed@strange-example.com",
		"test/test@test.com",
		"!#$%&'*+-/=?^_`{|}~@example.com",
		"user@mail.subdomain.example.com",
		"user@123.example.com",
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

// TestInvalidEmails comprueba direcciones que deben rechazarse.
func TestInvalidEmails(t *testing.T) {
	invalid := []string{
		"",
		"plainaddress",
		"@no-local-part.com",
		"no-at-sign.com",
		"user@",
		"user@@example.com",
		"user@.com",
		"user@example",
		".user@example.com",
		"user.@example.com",
		"us..er@example.com",
		"user name@example.com",
		"user@exam ple.com",
		"user@example..com",
		`"unterminated@example.com`,
		`"@example.com`,
		"John Doe <john@example.com>",
	}
	for _, e := range invalid {
		t.Run(e, func(t *testing.T) {
			if IsValid(e) {
				t.Errorf("se esperaba INVÁLIDO: %q", e)
			}
		})
	}
}

