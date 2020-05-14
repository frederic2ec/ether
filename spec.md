# Ether Spec

## Grammar

```
expr   : KEYWORD:var IDENTIFIER EQ expr
       : term ((PLUS|MINUS) term)*

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

## Basic

### Variables

```
$ integer = 1
$ float = 1.45
$ string = "string"
```
