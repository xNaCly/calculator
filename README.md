# Calculator

Simple educational calculator with lexer, parser and bytecode vm

This project is meant for explaining concepts used when interpreting programming languages, therefore containing
concepts such as:

- tokenizing characters to tokens
- parsing tokens to a abstract syntax tree
- tree walk interpreting
- transforming the abstract syntax tree to byte code
- interpreting byte code using a virtual machine

## Usage

```
$ calc "1+1*12.1/5"
3.4
```

```
$ cat calculations.txt
# operations
1+1
1-1
1*1
1/1

# chained operations
1*1+1-1/1
$ cat calculations.txt > calc
2
0
1
1
1
```

## How this project works

### Compiling the project

Compiling this project requires `go` with version 1.20:

```
$ go build .
```

Produces an executable for your architecture and operating system, which can be started:

```
$ ./calc
no input given
$ ./calc "1+1"
```

The last command supplies `calc` with `2+1*2` and promptly executes the expression:

```
$ ./calc "2+1*2"
index |            type |             raw

    0 |          NUMBER |               2
    1 |            PLUS |               +
    2 |          NUMBER |               1
    3 |        ASTERISK |               *
    4 |          NUMBER |               2
    5 |             EOF |             EOF
+
  *
    2
    1
  2
OP_LOAD    :: 2.000000
OP_STORE   :: 1.000000
OP_LOAD    :: 1.000000
OP_MULTIPY :: 1.000000
OP_STORE   :: 1.000000
OP_LOAD    :: 2.000000
OP_ADD     :: 1.000000
=> 4
```

The first output is the tokens generated with the lexical analysis, the second
output is the abstract syntax tree the parser builds, in the third part the
steps the virtual machine takes are traced to execute the expression. The last
output is the resulting number.
