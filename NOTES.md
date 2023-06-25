# Writing an Interpreter in Go

Notes from the [Writing an Interpreter in Go](https://interpreterbook.com) book.


* [Chapter 1 - Lexing](#chapter-1---lexing)
  * [Lexical Analysis](#lexical-analysis)

## Chapter 1 - Lexing

### Lexical Analysis

The source itself would be difficult to work with, inputting a programming language in order to then
have another programming language execute certain actions that are provided by the interpreted language.

Our source will move through 2 phases, before any "actions" are executed from our code.

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
factor in the Monkey language, as opposed to a language like Python.

A production-ready lexer may also attach line/column numbers and file names to tokens, as this can be very useful for later messages,
such as when errors occur - i.e. "error on line 5 near 'let' declaration" etc.

### REPL - Read, Eval, Print, Loop

This is what is commonly referred to as "interactive mode", in languages like Python and JavaScript you can use these extensively for debugging or
quick mock up of short snippets.

The concept is very simple and is in the name. Input is read, sent to the interpreter for evaluation, the result of that is printed and then the cycle is repeated.

## Chapter 2 - Parsing

### Parsers

#### What is a parser?

The Wikipedia entry provides a profoundly succinct definition:

> A parser is a software component that takes input data (frequently text) and builds
  a data structure – often some kind of parse tree, abstract syntax tree or other
  hierarchical structure – giving a structural representation of the input, checking for
  correct syntax in the process. […] The parser is often preceded by a separate lexical
  analyser, which creates tokens from the sequence of input characters;

That is to say that a parser is another transformation step of turning our source code (text) into something usable.
The key point here is that it can translate the various tokens it receives from our lexer into a hierarchical data structure which represents it.

A really great example here is that in JavaScript, there is a `JSON.parse` function, this turns some input string into a JSON structure which can be interacted with.

```javascript
var input = '{"name": "Jack", "age": 25}';
var output = JSON.parse(input);
output
{ name: "Jack", age: 25 }
```

The JSON parser in JavaScript is turning our `input` (text) into an `output` that is a JavaScript object which represents the initial input.
This is very obvious for JSON, but the same can actually be done for programming languages, these are more complex structures which they are parsed into.

### Abstract Syntax Tree (AST)

Most interpreters and compilers use a "syntax tree" to represent the source code that is input to them.

The "abstract" portion here is that there are some details omitted in the resulting structure, such as semicolons, whitespace, comments etc. are dependant on the language, but usually guide parser
when constructing the AST.

There is not a universal AST used by every parser, although they are conceptually very similar. There are varying implementation details.

As the process of parsing the input will ensure that it conforms to an expected structure, this is also sometimes referred to as *syntactic analysis*.

There are multiple strategies for parsing too, although the book does not go into much detail:

    - Top-down parsing
        - Recursive-decent parsing
        - Early parsing
    - Bottom-up parsing

The parsing strategy used for the Monkey programming language in this book is a top-down parser, specifically a recursive descent parser (top-down operator precedence) aka a "Pratt Parser" after its creator.

### Pratt Parsing

The main idea of this is to associate parsing functions (known as semantic code) with token types. Then, whenever a particular token type is encountered,
the function is called to return an AST node which represents the encountered type.

Each token type can have up to 2 parsing functions - depending on whether the token is found as either prefix or infix position.

- Infix would be `5 + 5;`, as the operator is "in-between" the integers.
- Prefix is like `--5`, the operator is _before_ its integer literal and will decrement the value.
