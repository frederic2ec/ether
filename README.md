# Ether
[![Go Report Card](https://goreportcard.com/badge/github.com/frederic2ec/ether)](https://goreportcard.com/report/github.com/frederic2ec/ether) ![Test](https://github.com/frederic2ec/ether/workflows/Test/badge.svg)

Ether is a programming language written in Go. The language is inspired by the book [Writing An Interpreter In Go](https://interpreterbook.com/)

## Table of contents
---
* [Ether](#ether)
  * [Quick start](#quick-start)
  * [Usage](#usage)
  * [Variable bindings](#variable-bindings)
  * [Arithmetic expressions](#arithmetic-expressions)
  * [Conditional expressions](#conditional-expressions)
  * [License](#license)


### Quick start
---
```#!sh
$ go get github.com/frederic2ec/ether
$ ether
```

### Usage
---
```#!sh
$ ether
Usage: ether [options] [<filename>]
  -i    enable interactive mode
  -v    display version information
```

### Variable bindings
---
```#!sh
#=> $variable = "value"
```

### Arithmetic expressions
---
```#!sh
#=> $a = 10
#=> $b = 5
#=> (a + b) / 2
25
```

### Conditional expressions
---
```#!sh
#=> $a = 10
#=> $b = 20
#=> $c = if(b > a): 50 } else: 51 }
#=> c
50
```

### License
---
This work is licensed under the terms of the GPL-3.0 License.