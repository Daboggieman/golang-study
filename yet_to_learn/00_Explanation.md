s## 1. What is Atoi and Itoa?
These are historical terms that made their way into Go and Piscines.
- **A** stands for ASCII (which basically means Text / String).
- **I** stands for Integer (a Math Number).
- **to** simply means "convert to".

### Atoi (String to Integer)
If you have a string `"123"`, you **cannot** do math with it. In Go, doing `"123" + 1` results in a crash.
**Atoi** takes the text `"123"` and converts it into the actual usable math number `123`.
*How it works behind the scenes:* It loops through each character, converts it from text to a digit (by subtracting the ASCII value of '0'), and builds the number mathematically (e.g., `1 * 100 + 2 * 10 + 3 * 1`).

### Itoa (Integer to String / ASCII)
If you have the number `123` and want to print it character by character, or concatenate it to a string like `"Score: " + 123` (which is illegal in Go), you use **Itoa**. It turns the number `123` into the text `"123"`.
*How it works behind the scenes:* It repeatedly divides the number by 10, grabs the remainder (which is the last digit), converts it to text, and builds the string backwards.

---

## 2. Type Conversion
In Go, types are very strict. You cannot mix them.
If you have `var a int = 5` and `var b float64 = 2.5`, you cannot add them together.
**Type Conversion** is explicitly telling Go to change the shape of the data:
`float64(a) + b` turns the integer `5` into `5.0` so the math works perfectly.

For strings and characters (runes/bytes):
- `string('A')` converts a character to a string.
- `byte(65)` converts a number to an ASCII character (`'A'`).

---

## 3. Factorials
In math, the factorial of a number (written as `5!`) means multiplying that number by every number below it down to 1.
`5! = 5 * 4 * 3 * 2 * 1 = 120`.

In programming, there are **two different types** of ways to calculate a factorial:
1. **Iterative (Using a Loop):** Using a standard `for` loop to multiply the numbers one by one. This is exactly what you are used to doing.
2. **Recursive (Function calling itself):** You calculate the factorial by having the function call *itself* with a smaller number until it hits 1. (e.g., `Factorial(5)` returns `5 * Factorial(4)`). It loops without using a `for` keyword!

---