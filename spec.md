# Ether Spec

## Grammar

```
expr   : term ((PLUS|MINUS) term)*

term   : factor ((MUL|DIV) factor)*

factor : (PLUS|MINUS) factor
       : modulo

power  : atom ((POW|MOD) factor)*

atom   : INT|FLOAT
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
