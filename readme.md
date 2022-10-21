# Coinspaid GO traning

## Homework

### Lesson 1

#### Exercise 1. Fibonacci sequence

Реализовать функцию, которая возвращает следующее в ряду число Фибоначчи.

#### Exercise 2. Arithmetic progression

Реализовать функцию, которая принимает на вход правило в виде функции. Правило определяет an от an-1.
Возвращает генератор который последовательно возвращает очередной член прогрессии.

#### Exercise 3. Check for duplicate characters in a string

Определить есть ли в строке повторяющиеся буквы латинского алфавита, игнорируя регистр.
Можно использовать только один цикл range либо for и только одну переменную int64 помимо итератора.
Дополнительные пакеты использовать нельзя.

#### Exercise 4. Reverse polish notation calculator

Написать функции eval, plus, minus так чтобы
eval(10, 20, plus, "45", minus) => -15
eval(10, 2.5, plus) => 12.5
В случае если вычислить не удается, предусмотреть возвращение 0
Для преобразования пакет strconv:
strconv.Atoi(“10”) - для целых
strconv.ParseFloat("1.23", 64) - для float

### Lesson 2

#### Exercise 1. Slice resizing

Определить как меняется размер слайса при добавлении элементов.

### Lesson 4

#### Exercise 1. Thread safe map

Сделать потокобезопасный map с помощью mutex
