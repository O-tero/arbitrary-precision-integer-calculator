package main

import (
	"fmt"
	"os"
	"bufio"
	"math"
	"math/big"
	"strings"
	"strconv"
	
)
// Helper function to print numbers in reverse (to make operations easier)
func reverseStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func add(x, y string) string {
	// Make the numbers the same length by padding with zeros
	maxLen := len(x)
	if len(y) > maxLen {
		maxLen = len(y)
	}
	x = strings.Repeat("0", maxLen-len(x)) + x
	y = strings.Repeat("0", maxLen-len(y)) + y

	// Adding each digit starting from the right
	var result string
	carry := 0

	for i := maxLen - 1; i >= 0; i-- {
		digitSum := int(x[i]-'0') + int(y[i]-'0') + carry
		result = string(digitSum%10+'0') + result
		carry = digitSum / 10
	}

	// If there is a carry left, add it
	if carry > 0 {
		result = string(carry+'0') + result
	}

	return result
}


func subtract(x, y string) string {
	// Make the numbers the same length by padding with zeros
	maxLen := len(x)
	if len(y) > maxLen {
		maxLen = len(y)
	}
	x = strings.Repeat("0", maxLen-len(x)) + x
	y = strings.Repeat("0", maxLen-len(y)) + y

	var result string
	borrow := 0

	for i := maxLen - 1; i >= 0; i-- {
		digitDiff := int(x[i]-'0') - int(y[i]-'0') - borrow
		if digitDiff < 0 {
			digitDiff += 10
			borrow = 1
		} else {
			borrow = 0
		}
		result = string(digitDiff+'0') + result
	}

	// Remove leading zeros
	result = strings.TrimLeft(result, "0")

	if result == "" {
		return "0"
	}
	return result
}

// Multiply two arbitrary precision integers represented as strings
func multiply(x, y string) string {
	bx, by := new(big.Int), new(big.Int)
	bx, _ = bx.SetString(x, 10)
	by, _ = by.SetString(y, 10)

	bz := new(big.Int)
	bz.Mul(bx, by)

	return bz.String()
}

// Divide two arbitrary precision integers represented as strings
func divide(x, y string) (string, string) {
	bx, by := new(big.Int), new(big.Int)
	bx, _ = bx.SetString(x, 10)
	by, _ = by.SetString(y, 10)

	quotient, remainder := new(big.Int), new(big.Int)
	quotient.QuoRem(bx, by, remainder)

	return quotient.String(), remainder.String()
}

// Exponentiation of an arbitrary precision integer represented as a string
func exponentiate(base, exp string) string {
	bx, be := new(big.Int), new(big.Int)
	bx, _ = bx.SetString(base, 10)
	be, _ = be.SetString(exp, 10)

	result := new(big.Int)
	result.Exp(bx, be, nil)

	return result.String()
}

// Factorial of an arbitrary precision integer represented as a string
func factorial(n string) string {
	bn, _ := new(big.Int).SetString(n, 10)
	result := new(big.Int).SetInt64(1)

	for i := new(big.Int).SetInt64(2); bn.Cmp(i) >= 0; i.Add(i, big.NewInt(1)) {
		result.Mul(result, i)
	}

	return result.String()
}

// Helper function to convert a string number to a specific base (e.g., "1011" to base 2)
func baseToDecimal(num string, base int) (*big.Int, error) {
	result := big.NewInt(0)
	for _, digit := range num {
		value := int(digit - '0')
		if value >= base {
			return nil, fmt.Errorf("invalid digit %c for base %d", digit, base)
		}
		result.Mul(result, big.NewInt(int64(base)))
		result.Add(result, big.NewInt(int64(value)))
	}
	return result, nil
}

// Convert a decimal number to a string in a given base
func decimalToBase(num *big.Int, base int) string {
	if num.Cmp(big.NewInt(0)) == 0 {
		return "0"
	}

	var result []rune
	for num.Cmp(big.NewInt(0)) > 0 {
		mod := new(big.Int)
		num.DivMod(num, big.NewInt(int64(base)), mod)
		result = append([]rune{rune(mod.Int64() + '0')}, result...)
	}
	return string(result)
}

// Logarithm base b of a number (log_b(n))
func logarithm(n, b string) (string, error) {
	// Parse inputs as big.Float
	bn, ok1 := new(big.Float).SetString(n)
	bb, ok2 := new(big.Float).SetString(b)

	if !ok1 || !ok2 {
		return "", fmt.Errorf("invalid input: n=%s, b=%s", n, b)
	}

	// Convert to float64 for approximation of ln
	num, _ := bn.Float64()
	base, _ := bb.Float64()

	if num <= 0 || base <= 0 || base == 1 {
		return "", fmt.Errorf("invalid values: n and b must be > 0 and b != 1")
	}

	// Compute natural logarithms
	logN := math.Log(num)   // ln(n)
	logB := math.Log(base)  // ln(b)

	// Calculate log_b(n) = ln(n) / ln(b)
	result := logN / logB

	// Return result as string
	return fmt.Sprintf("%.10f", result), nil
}


// Parse and handle fractions (represented as "a/b")
func addFractions(a, b, c, d string) string {
	// Convert a/b + c/d
	// (a*d + b*c) / (b*d)
	na, _ := new(big.Int).SetString(a, 10)
	nb, _ := new(big.Int).SetString(b, 10)
	nc, _ := new(big.Int).SetString(c, 10)
	nd, _ := new(big.Int).SetString(d, 10)

	num := new(big.Int).Add(new(big.Int).Mul(na, nd), new(big.Int).Mul(nb, nc))
	den := new(big.Int).Mul(nb, nd)

	return fmt.Sprintf("%s/%s", num.String(), den.String())
}

// Convert fraction to decimal
func fractionToDecimal(n, d string) string {
	num, _ := new(big.Int).SetString(n, 10)
	den, _ := new(big.Int).SetString(d, 10)
	result := new(big.Float).Quo(new(big.Float).SetInt(num), new(big.Float).SetInt(den))
	return result.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Arbitrary Precision Calculator!")
	fmt.Println("Enter expression (or 'exit' to quit):")

	for {
		// Read user input
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Exit condition
		if input == "exit" {
			break
		}

		// Parse input
		parts := strings.Fields(input)
		if len(parts) == 3 {
			// Binary operations: a op b
			a := parts[0]
			op := parts[1]
			b := parts[2]

			switch op {
			case "+":
				fmt.Println(add(a, b))
			case "-":
				fmt.Println(subtract(a, b))
			case "*":
				fmt.Println(multiply(a, b))
			case "/":
				if b == "0" {
					fmt.Println("Division by zero!")
				} else {
					quotient, remainder := divide(a, b)
					fmt.Println("Quotient:", quotient)
					fmt.Println("Remainder:", remainder)
				}
			case "%":
				ai, err1 := strconv.Atoi(a)
				bi, err2 := strconv.Atoi(b)
				if err1 != nil || err2 != nil {
					fmt.Println("Invalid number!")
				} else {
					fmt.Println(ai % bi)
				}
			default:
				fmt.Println("Unknown operator!")
			}
		} else if strings.HasSuffix(parts[0], "!") {
				// Factorial: a!
			number := strings.TrimSuffix(parts[0], "!")
			a, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Invalid number!")
				continue
			}
			fmt.Println(factorial(strconv.Itoa(a)))
		} else {
			fmt.Println("Unknown expression!")
		}
	}

	// Additional functionality
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter expression (or 'exit' to quit): ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		var x, y, op string
		var base, fracOp string
		var n, b string

		// Handle fractions
		if strings.Contains(input, "/") {
			fmt.Sscanf(input, "%s %s %s %s", &x, &op, &y, &fracOp)
			if fracOp == "/" {
				partsX := strings.Split(x, "/")
				partsY := strings.Split(y, "/")
				fmt.Println("Fraction result:", addFractions(partsX[0], partsX[1], partsY[0], partsY[1]))
			}
		}

		// Base conversion
		if strings.Contains(input, "base") {
			fmt.Sscanf(input, "%s base %s", &x, &base)
			converted, err := baseToDecimal(x, int(base[0]-'0'))
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Converted:", converted)
			}
		} else if strings.Contains(input, "^") {
			fmt.Sscanf(input, "%s ^ %s", &x, &y)
			fmt.Println("Exponentiation:", exponentiate(x, y))
		} else if strings.Contains(input, "log") {
			fmt.Sscanf(input, "log %s %s", &n, &b)
			result, err := logarithm(n, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Logarithm:", result)
			}
		} else if op == "+" {
			fmt.Println("Result:", add(x, y))
		} else if op == "-" {
			fmt.Println("Result:", subtract(x, y))
		} else if op == "*" {
			fmt.Println("Result:", multiply(x, y))
		} else if op == "/" {
			quotient, remainder := divide(x, y)
			fmt.Println("Quotient:", quotient)
			fmt.Println("Remainder:", remainder)
		} else if op == "!" {
			fmt.Println("Factorial:", factorial(x))
		} else {
			fmt.Println("Unknown operator!")
		}
	}
	fmt.Println("Goodbye!")
}