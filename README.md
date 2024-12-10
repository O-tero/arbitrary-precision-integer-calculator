# Arbitrary Precision Integer Calculator
* This project implements an arbitrary-precision integer calculator in Go. The calculator supports basic arithmetic operations, factorials, and exponentiation. Additionally, it features advanced functionality like base conversions, fractional arithmetic, and logarithmic calculations. What's unique about this implementation is that it doesn't rely on external libraries for its core features, instead leveraging Go's built-in capabilities (math/big) to handle arbitrary-precision arithmetic.

* The application is wrapped in a REPL (Read-Eval-Print Loop), allowing users to interact with the calculator via simple textual commands.

## Key Features
1. Arbitrary-Precision Arithmetic
- The calculator handles operations on integers of any size without overflow.
Operations supported:
Addition: e.g., `1000000000000000000 + 2000000000000000000`. Subtraction: e.g., `5000 - 3000`. Multiplication: e.g., `123456789 * 987654321`. Division with Remainder: e.g., `10 / 3` produces a quotient and remainder. Factorials: e.g., `100!`. Exponentiation: e.g., `2^1000`
2. Support for Non-Decimal Bases
Convert numbers between arbitrary bases and decimal.
Example: Convert `1011` from binary (base 2) to decimal, or vice versa.
Demonstrates understanding of base arithmetic and encoding, especially for binary, hexadecimal, and other non-standard bases.
3. Fractional Arithmetic
Perform operations on fractions (e.g., `1/2 + 3/4`).
Returns results in fraction form without premature conversion to decimals.
Ability to convert fractions to decimal representation.
4. Logarithmic Calculations
Compute logarithms in arbitrary bases using natural logarithms.
Example: Calculate `log_2(16`) or `log_10(1000`).
5. Interactive REPL Interface
User-friendly REPL loop to input commands and view results interactively.
Supports mixed functionality (e.g., fractions and logarithms) seamlessly.
Error handling for invalid inputs and unsupported operations.
Design Considerations
### Why Go?
- Go doesn't have native support for arbitrary-precision integers in its core language features but offers the `math/big` package, which provides low-level building blocks. By combining these primitives with custom logic, the calculator achieves a wide range of functionality.

- Modularity
The program is designed with modularity in mind:
Each feature (e.g., base conversion, fraction handling, logarithms) is implemented as a separate function.
The REPL acts as the orchestrator, parsing user input and calling the appropriate functionality.
### What's Most Interesting?
- Base Conversions:

 Implementing base conversion is computationally significant when working with extremely large numbers in non-decimal bases.
 This required thinking through digit-by-digit processing and efficient modular arithmetic.
- Fraction Handling:

Instead of converting fractions to decimals prematurely, operations are performed directly on the numerator and denominator. This ensures precision and minimizes rounding errors.
- Logarithmic Calculations:

Using natural logarithms and their properties to derive arbitrary base logarithms showcases an elegant use of mathematical principles.
- REPL Interface:

Building an intuitive REPL interface required balancing simplicity with flexibility. The REPL dynamically parses input and invokes the correct functions, demonstrating robust input parsing logic.
## Getting Started
### Prerequisites
* Install Go: [![Download Go](https://go.dev/doc/install])

### Running the Program

1. Clone the repository and navigate to the folder.

2. Save the program in a file called `calculator.go`.

3. Open a terminal and navigate to the folder where the file is saved.

4. Run the following command:


```golang
go run calculator.go
```

5 Start using the REPL.

## REPL Usage
The program will display a prompt:

```golang
Enter expression (or 'exit' to quit):
```

Example Commands:
* Basic Operations:
`100 + 200`
`500 * 300`

* Base Conversion:
`1011 base 2`
Convert a number to decimal from any base.

* Fraction Arithmetic:
`1/2 + 3/4`

* Logarithms:
`log 16 2`
Calculate the logarithm of a number with a specified base.

* Factorials:
`100!`

* Exponentiation:
`2^10`

* Type `exit` to quit the program.

## Challenges and Solutions
* Precision Handling - Ensuring precision for operations like factorials and logarithms required careful use of the `math/big` package.
* Error Handling - The program validates inputs to prevent operations like dividing by zero, invalid base conversions, or malformed fractions.
* Fraction Simplification - Fractions are stored as numerator and denominator. The addition and multiplication operations reduce them to their simplest form.
* Scalability - Efficient algorithms for base conversion and logarithmic calculations were chosen to handle large-scale inputs without significant performance degradation.
## Future Improvements
* Advanced Mathematical Functions - Add support for square roots, trigonometric functions, and more.
* Decimal Bases for Fractional Input - Extend base conversion to work for fractional bases (e.g., converting `1.101` in base 2).
* Custom Error Messages - Provide more descriptive feedback for invalid inputs.
* Configuration Options - Allow users to set precision for logarithms or choose rounding modes for fractions.
## Conclusion
This calculator highlights the versatility of Go's `math/big` package and the power of modular program design. It’s a practical demonstration of implementing advanced mathematical operations and a user-friendly REPL interface in a language without native support for arbitrary precision. This project serves as a great foundation for anyone looking to explore computational mathematics and algorithmic design in Go.
