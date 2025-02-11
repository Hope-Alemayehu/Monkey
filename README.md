# Monkey Interpreter in Go

Monkey is a simple programming language interpreter implemented in Go, inspired by the language from the book ["Writing an Interpreter in Go" by Thorsten Ball](https://interpreterbook.com/).

## Features

- Variables and bindings
- Integer arithmetic
- Boolean expressions
- Conditionals (`if` statements)
- Functions and closures
- A REPL (Read-Eval-Print Loop)

## Getting Started

### Prerequisites

Ensure you have Go installed on your system.

- [Download Go](https://go.dev/dl/)

### Installation

1. Clone the repository:

```
   git clone https://github.com/Hope-Alemayehu/Monkey
   gitmonkey-interpreter.git
   cd monkey-interpreter
```
Build the project:

```
go build -o monkey
```
Run the REPL:
```
./monkey
```
Usage
Once inside the REPL, you can write Monkey code:

```
>> let five = 5;
>> let ten = 10;
>> five + ten;
15
>> if (5 < 10) { return true; } else { return false; }
true
```

## Project Structure
```
monkey-interpreter/
├── lexer/          # Lexical analysis
├── parser/         # Parsing expressions and statements
├── ast/            # Abstract Syntax Tree
├── repl/           # Interactive Read-Eval-Print Loop
├── token/          # Token definitions
├── main.go         # Entry point
├── README.md       # Project documentation
```

## Contributing
Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.