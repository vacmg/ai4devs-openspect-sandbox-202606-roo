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

## Validación: librería estándar `net/mail`

En lugar de implementar la gramática de RFC 5322 desde cero, la validación
delega en la **librería estándar de Go
[`net/mail`](https://pkg.go.dev/net/mail)**, cuya función `mail.ParseAddress`
implementa el parser de direcciones de [RFC 5322
§3.4.1](https://datatracker.ietf.org/doc/html/rfc5322#section-3.4.1).

Sobre eso aplicamos dos reglas adicionales:

1. **Solo addr-spec puro**: rechazamos formas con display-name como
   `John Doe <john@example.com>` comprobando que `addr.Name == ""` y que
   `addr.Address` coincide con la entrada literal (también descarta
   comentarios CFWS).
2. **Dominio con al menos un punto**: descartamos `user@host` (que el parser
   acepta) ya que no es una dirección enrutable en la práctica.

> Nota: `net/mail` no aplica las reglas LDH estrictas del RFC 1035 al
> dominio (acepta guion bajo, etiquetas que empiezan/terminan con guion,
> etc.) ni el límite de 64 octetos en local-part. Si necesitas esa
> estrictez, conviene añadir una capa adicional o una librería externa
> como `github.com/go-playground/validator`.

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

