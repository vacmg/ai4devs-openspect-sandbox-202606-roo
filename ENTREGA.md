*Micro-tarea:* Validador de email

*Pilar 1 — Herramienta:* ¿Cuál eliges?

Github copilot en modo Agente dentro de un IDE (Categoría B).

¿Por qué esta y no otra?

Porque el objetivo final es una función que valide emails, es una feature relativamente sencilla, y el agente deberia ser capaz de preparar el entorno de desarrollo, crear la estructura de carpetas y archivos, y generar la logica de validacion con TDD.

*Pilar 2 — Contexto:* ¿Qué información estás aportando? (lenguaje, framework, restricciones, ejemplos…)

Le he dado de contexto el lenguaje de programación (Go), que quiero que la salida sea una aplicacion de terminal que reciba por stdin un email y devuelva 0 si es correcto, o 1 si no; y que quiero hacer TDD.

¿Hay algo del contexto que has decidido omitir conscientemente?
Sí, he decidido omitir el framework de testing específico y las restricciones de validación de email (como la longitud máxima o los caracteres permitidos), para ver cómo el copiloto maneja la tarea con información mínima.

*Pilar 3 — Prompt:* ¿Cómo lo estructuras? (estilo, formato de salida, ejemplos…)

Quiero hacer una aplicacion para terminal en go que reciba por stdin un email y devuelva 0 si es correcto, o 1 si no. 
Quiero garantizar que la aplicacion funciona mediante tests usando TDD.
Quiero que compruebes que los emails siguen el estandar donde se define la sintaxis de los emails.

Second prompt:

Si existe alguna libreria estandar o muy conocida para validar emails, quiero que la uses en vez de escribir la logica de validacion desde cero.

*Resultado:* ¿Funcionó a la primera o tuviste que iterar?

Funcionó a la primera, pero quise iterar para utilizar una libreria en vez de hacer la logica de cero.

Una mejora que harías si volvieras a hacerlo.

Decidir de antemano si quieres usar una libreria o no, para no tener que iterar despues.
