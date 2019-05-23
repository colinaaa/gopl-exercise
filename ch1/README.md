# Chapter 1 Tutorial

> This chapter is a tour of the basic components of *Go*.

## Content

1. Hello, World
2. Command-Line Arguments
   1. Modify the `echo` program to also print `os.Args[0]`, the name of the command that invoked it.
   2. Modify the `echo` program to print the index and value of each of its arguments, on per line.
   3. Experiment to measure the difference in running time between out potentially inefficient versions and the on that uses `strings.Join`


## Highlights

### Semicolons

> Newlines following certain tokens are converted into semicolons, so where newlines are placed matters to proper parsing of *Go* code.
> For instance, the poening brace `{` of the function must be on the same line as the end of the `func` declaration, ont on a line by itself, and in the expression `x+y`, a newline is permitted after but not before the `+` operator.

### Assignment Operator

```go
s += sep + os.Args[i]
```

The statement is an *assignment operator*; it is equivalent to

```go
s = s + sep + os.Args[i]
```

The operator `+=` is an *assignment operator*. Each arithmetic and logical operator like `+` or `*` has a corresponding assignment operator.


