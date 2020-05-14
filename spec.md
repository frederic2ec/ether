# Ether Spec

## Grammar

```
expr   : DOLLAR IDENTIFIER EQ expr
       : comp ((AND|OR) comp)*

comp   : NOT comp
       : arith ((EE|LT|GT|LTE|GTE) arith)*

arith  : term ((PLUS|MINUS) term)*

term   : factor ((MUL|DIV) factor)*

factor : (PLUS|MINUS) factor
       : modulo

power  : atom ((POW|MOD) factor)*

atom   : INT|FLOAT|IDENTIFIER
       : LPAREN expr RPAREN
```

## Math

### Arithmetic

```
1 + 1 # Addition
2 - 1 # Subtraction
3 * 5 # Multiplication
40 / 8 # Division
2 ^ 5 # Exponention
12 % 5 # Modulo
```

### Bitwise

```
3 & 5 # and
3 ~ 5 # or
!0 # not
```

## Basic

### Variables

```
$ integer = 1 or $interger = 1
$ float = 1.45 or $float = 1.45
$ string = "string" or $string = "string"
```

### Comparators

```
1 == 1 #=> 1
2 == 1 #=> 0
2 != 1 #=> 1
1 < 10 #=> 1
1 > 10 #=> 0
2 <= 2 #=> 1
2 >= 2 #=> 1
```

### Logic

```
1 & 0 #=> 0
1 ~ 0 #=> 1
!1 & 0 #=> 1
```
