# Teorema de Bayes

## Explicacion del teorema de Bayes

El teorema de Bayes es una proposición lógica que relaciona la probabilidad condicional de un evento A dado un evento B con la probabilidad condicional de B dado A y la probabilidad marginal de B. Es decir, relaciona las probabilidades condicionales de dos eventos al revés de como se dan en la naturaleza.

$$ P(A|B) = \frac{P(B|A) \cdot P(A)}{P(B)} $$

## Ejemplo 1

Tenemos tres urnas: A con 3 bolas rojas y 5 negras, B con 2 bolas rojas y 1 negra y C con 2 bolas rojas y 3 negras. Escogemos una urna al azar y extraemos una bola.

La tabla de probabilidades obtenida es la siguiente:

```Bash
╭─────┬───────┬───────┬───────┬───────┬─────────────╮
│ URN │ PROB. │ COLOR │ COUNT │ PROB. │ TOTAL PROB. │
├─────┼───────┼───────┼───────┼───────┼─────────────┤
│ B   │ 0.33  │ Black │     1 │ 0.33  │ 0.21        │
│     │       │ Red   │     2 │ 0.67  │ 0.46        │
├─────┼───────┼───────┼───────┼───────┼─────────────┤
│ C   │ 0.33  │ Black │     3 │ 0.60  │ 0.39        │
│     │       │ Red   │     2 │ 0.40  │ 0.28        │
├─────┼───────┼───────┼───────┼───────┼─────────────┤
│ A   │ 0.33  │ Black │     5 │ 0.62  │ 0.40        │
│     │       │ Red   │     3 │ 0.38  │ 0.26        │
╰─────┴───────┴───────┴───────┴───────┴─────────────╯
```

Al obtener una bola roja, ¿cuál es la probabilidad de que la urna escogida sea la A?

$$ P(A|Red) = \frac{P(Red|A) \cdot P(A)}{P(Red)} =
\frac{0.38 \cdot 0.33}{.33 \cdot (.67 + .40 + .38)} = 0.26 $$

## Ejemplo 2

Tres máquinas, A, B y C, producen el 45%, 30% y 25%, respectivamente, del total de las piezas producidas en una fábrica. Los porcentajes de producción defectuosa de estas máquinas son del 3%, 4% y 5%.

Se obtienen los siguientes datos:

- Seleccionamos una pieza al azar; calcula la probabilidad de que sea defectuosa.

```Bash
Probabilidad defectuosa general:
3.8000000000000007 %
```

- Tomamos, al azar, una pieza y resulta ser defectuosa; calcula la probabilidad de haber sido producida por la máquina B.

```Bash
Probabilidad defectuosa producida por la máquina A:
35.52631578947368 %
Probabilidad defectuosa producida por la máquina B:
31.57894736842105 %
Probabilidad defectuosa producida por la máquina C:
32.89473684210526 %
```

- ¿Qué máquina tiene la mayor probabilidad de haber producido la citada pieza defectuosa?

```Bash
Probabilidad más alta de defectuosos:
35.52631578947368 % de la máquina A
```
