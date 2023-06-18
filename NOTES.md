# Writing an Interpreter in Go

Notes from the [Writing an Interpreter in Go](https://interpreterbook.com) book.


* [Chapter 1 - Lexing](#chapter-1---lexing)
  * [Lexical Analysis](#lexical-analysis)

## Chapter 1 - Lexing

### Lexical Analysis

The source itself would be difficult to work with, inputting a programming language in order to then
have another programming language execute certain actions.

Our source will move through 2 phases, before an "actions" are done from our code.

The first step is to turn our source into "tokens", this is the process known as *lexical analysis* and is done by
a lexer (sometimes known as a tokeniser or scanner). These tokens are very small, categorisable data structures which can be 
fed to a parser which will then create an abstract syntax tree (AST), which is our second transformation.


An example would be the following:

Input source code (text):

```
let x = 5;
```

And the resulting tokens would be:

```
[
     LET,
     IDENTIFIER("x"),
     EQUAL_SIGN,
     INTEGER(5),
     SEMICOLON
]
```

Something to note, is that whitespace does not show up as a token here. This is because whitespace is not an important
factor in the Monkey language, as opposed to an interpreted language like Python.

A production-ready lexer may also attach line/column numbers and file names to tokens, as this can be very useful for later messages,
such as when errors occur - i.e. "error on line 5 near 'let' declaration" etc.
