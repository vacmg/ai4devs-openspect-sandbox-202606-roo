# emailcheck

Aplicación CLI en Go que lee un email por **STDIN** y devuelve:

- `0` si el email es **válido** según la sintaxis de RFC 5322 (addr-spec).
- `1` si **no** es válido (o si la entrada está vacía / hay error de E/S).

Desarrollada con TDD: la lógica de validación está cubierta por tests con casos
válidos e inválidos antes de implementar.

## Estructura

```
cmd/emailcheck/   # binario CLI (lee stdin, devuelve exit code)
validator/        # paquete con la validación RFC 5322 + tests
```

## Estándar implementado

Se valida la sintaxis **addr-spec** definida en
[RFC 5322 §3.4.1](https://datatracker.ietf.org/doc/html/rfc5322#section-3.4.1):

- `local-part` puede ser:
  - `dot-atom`: secuencias de `atext` separadas por puntos (sin puntos al
    inicio/fin ni dobles). `atext` incluye letras, dígitos y los símbolos
    ``! # $ % & ' * + - / = ? ^ _ ` { | } ~``.
  - `quoted-string`: `"..."` con `qtext` y `quoted-pair` (`\X`).
- `domain` es un `dot-atom` con etiquetas LDH (RFC 1035): letras, dígitos y
  guiones, sin guiones al inicio/fin, longitud de etiqueta 1..63, y al menos
  dos etiquetas (TLD obligatorio).
- Límites prácticos: `local-part` ≤ 64, total ≤ 254, dominio ≤ 253.

No se soportan `domain-literal` (`[IPv4/IPv6]`) ni comentarios CFWS, por ser
inusuales y desaconsejados en la práctica (RFC 5321).

## Compilar

```bash
go build -o emailcheck ./cmd/emailcheck
```

## Uso

```bash
echo "user@example.com" | ./emailcheck ; echo $?   # 0
echo "no-arroba"        | ./emailcheck ; echo $?   # 1
```

Útil en scripts:

```bash
if echo "$EMAIL" | ./emailcheck; then
  echo "válido"
else
  echo "inválido"
fi
```

## Tests (TDD)

```bash
go test ./...
go test ./... -v          # detallado
go test ./validator/ -run TestValidEmails -v
```

Cobertura:

```bash
go test ./... -cover
```

