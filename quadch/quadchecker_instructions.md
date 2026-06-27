# QuadChecker Implementation Guide

This file explains exactly how to write the `quadchecker` program in Go.
It does not contain the finished code, but it tells you every step, the functions to use, and the logic needed.

## Goal

Write a Go program that:

- reads a text pattern from input
- determines whether the pattern matches one or more quad functions
- prints the matching quad names and their dimensions
- prints `Not a quad function` if there is no match
- prints multiple matches in alphabetical order, separated by ` || `
- always ends output with a newline

## Required project structure

Your folder should contain:

- `go.mod`
- `main.go`

Example structure:

```
quadchecker
├── go.mod
└── main.go
```

## `go.mod`

Create a `go.mod` file with a module name such as:

```go
module quadchecker

go 1.20
```

This tells Go the project root and the language version.

## `main.go` overview

The `main.go` file should contain:

- `package main`
- imports for:
  - `fmt`
  - `os`
  - `bufio`
  - `strings`
  - `sort`
- a `main()` function
- helper functions for reading input and checking quad patterns

## Input handling

The program must accept the quad text pattern as input. There are two safe ways to do this:

1. If the program is called with a single command-line argument, use that argument.
2. Otherwise, read the entire standard input.

### Functions to use

- `os.Args` to inspect command-line arguments
- `os.Stdin` with `bufio.NewScanner` or `io.ReadAll` to read piped input
- `strings.TrimSuffix` or `strings.TrimRight` to remove a trailing newline if needed
- `strings.Split(text, "\n")` to split the input into lines

### Recommended logic

1. In `main()`, check `len(os.Args)`:
   - if `len(os.Args) == 2`, use `os.Args[1]` as the pattern text
   - else read from `os.Stdin`
2. Normalize the input text:
   - remove any final empty line caused by a trailing newline
   - preserve line order exactly
3. Split input text into lines
4. If there are no lines or the pattern is empty, print `Not a quad function`

## Pattern analysis

A quad pattern is a rectangular drawing. You can treat the input as lines of text with rows and columns.

### Dimensions

- height = number of lines
- width = length of the longest line

If any line is shorter than the maximum width, you should treat missing characters as spaces when checking the pattern.

### Normalization helper

Write a helper that converts every line to the same width by padding shorter lines with spaces.
This makes pattern matching easier.

## Quad functions to detect

You need to detect these five patterns:

- `quadA`
- `quadB`
- `quadC`
- `quadD`
- `quadE`

These correspond to the classic 42-style rectangle generators.

### General rules for all quad patterns

For a rectangle with width `w` and height `h`:

- top row is row 0
- bottom row is row `h-1`
- left column is column 0
- right column is column `w-1`
- interior rows are rows 1 through `h-2`
- interior columns are columns 1 through `w-2`

If `w == 1` or `h == 1`, corners overlap and the pattern becomes a single column or row.

## Specific pattern definitions

### quadA

Expected characters:

- top-left corner: `o`
- top-right corner: `o`
- bottom-left corner: `o`
- bottom-right corner: `o`
- top and bottom edges: `-`
- left and right edges: `|`
- interior: space ` `

For `w >= 2` and `h >= 2`:

- top row = `o` + `-` repeated `w-2` + `o`
- bottom row = same as top row
- middle rows = `|` + spaces repeated `w-2` + `|`

For `w == 1` and `h == 1`:

- output is just `o`

For a single column (`w == 1`, `h > 1`):

- first line = `o`
- middle lines = `|`
- last line = `o`

For a single row (`h == 1`, `w > 1`):

- line = `o` + `-` repeated `w-2` + `o`

### quadB

Expected characters:

- top-left corner: `/`
- top-right corner: `\`
- bottom-left corner: `\`
- bottom-right corner: `/`
- top and bottom edges: `*`
- left and right edges: `*`
- interior: space ` `

For `w >= 2` and `h >= 2`:

- top row = `/` + `*` repeated `w-2` + `\`
- bottom row = `\` + `*` repeated `w-2` + `/`
- middle rows = `*` + spaces repeated `w-2` + `*`

For `w == 1` and `h == 1`:

- output is `/`

For a single column:

- first line = `/`
- middle lines = `*`
- last line = `\`

For a single row:

- line = `/` + `*` repeated `w-2` + `\`

### quadC

Expected characters:

- top-left corner: `A`
- top-right corner: `A`
- bottom-left corner: `C`
- bottom-right corner: `C`
- top and bottom edges: `B`
- left and right edges: `B`
- interior: space ` `

For `w >= 2` and `h >= 2`:

- top row = `A` + `B` repeated `w-2` + `A`
- bottom row = `C` + `B` repeated `w-2` + `C`
- middle rows = `B` + spaces repeated `w-2` + `B`

For a 1x1 rectangle:

- output is `A`

For a single column:

- top line = `A`
- middle lines = `B`
- bottom line = `C`

For a single row:

- line = `A` + `B` repeated `w-2` + `A`

### quadD

Expected characters:

- top-left corner: `A`
- top-right corner: `C`
- bottom-left corner: `A`
- bottom-right corner: `C`
- top and bottom edges: `B`
- left and right edges: `B`
- interior: space ` `

For `w >= 2` and `h >= 2`:

- top row = `A` + `B` repeated `w-2` + `C`
- bottom row = `A` + `B` repeated `w-2` + `C`
- middle rows = `B` + spaces repeated `w-2` + `B`

For a 1x1 rectangle:

- output is `A`

For a single column:

- top line = `A`
- middle lines = `B`
- bottom line = `C`

For a single row:

- line = `A` + `B` repeated `w-2` + `C`

### quadE

Expected characters:

- top-left corner: `A`
- top-right corner: `C`
- bottom-left corner: `C`
- bottom-right corner: `A`
- top and bottom edges: `B`
- left and right edges: `B`
- interior: space ` `

For `w >= 2` and `h >= 2`:

- top row = `A` + `B` repeated `w-2` + `C`
- bottom row = `C` + `B` repeated `w-2` + `A`
- middle rows = `B` + spaces repeated `w-2` + `B`

For a 1x1 rectangle:

- output is `A`

For a single column:

- top line = `A`
- middle lines = `B`
- bottom line = `C`

For a single row:

- line = `A` + `B` repeated `w-2` + `C`

## How to implement pattern checks

### Helper idea: `matchPattern`

Write a helper function that compares the input lines to an expected pattern.

It should:

- take normalized lines, width, height, and the expected characters for corners and edges
- iterate every row and every column
- compare the actual character to the expected character for that position
- return `false` as soon as any character differs

This helper lets you reuse the same checking logic for all five quadrants.

### What to check for each cell

For each position `(x, y)`:

- if `y == 0` and `x == 0`: expected top-left corner
- if `y == 0` and `x == w-1`: expected top-right corner
- if `y == h-1` and `x == 0`: expected bottom-left corner
- if `y == h-1` and `x == w-1`: expected bottom-right corner
- if `y == 0` or `y == h-1`: expected top/bottom edge character
- if `x == 0` or `x == w-1`: expected side edge character
- otherwise: expected interior space ` `

Special-case when `w == 1` or `h == 1`, because the same position may be both a corner and an edge. In that case, decide the expected character based on the combined rules above.

## Which matches are possible

Some inputs can match more than one pattern.

The examples show these cases:

- `[quadC] [1] [1]`, `quadD` and `quadE` all produce the same single-character output `A`
- `quadC` and `quadE` both produce the same 1x2 output:
  - `A`
  - `C`

That means your program must test all five quad patterns and collect every match.

## Result formatting

In `main()`:

1. Build a slice of matching names, for example:
   - `matches := []string{}`
2. For each quad pattern that returns `true`, append a formatted string:
   - `fmt.Sprintf("[quadA] [%d] [%d]", width, height)`
3. If there are matches:
   - sort them with `sort.Strings(matches)`
   - join them with `strings.Join(matches, " || ")`
   - print the result with `fmt.Println(...)`
4. If there are no matches:
   - print `Not a quad function`

Because `fmt.Println` adds a newline automatically, it guarantees the final newline.

## Hints for debugging

- When testing, use the same input style as the project examples.
- For a known quad function, pipe its output into your program:
  - `./quadA 3 3 | go run .`
- For ambiguous cases, the output should list multiple quad names.
- For invalid input, print exactly `Not a quad function`

## Examples to verify your implementation

Use these checks after writing and compiling `main.go`:

- `./quadA 3 3 | go run .`
  - expected: `[quadA] [3] [3]`

- `./quadC 1 1 | go run .`
  - expected: `[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]`

- `./quadE 1 2 | go run .`
  - expected: `[quadC] [1] [2] || [quadE] [1] [2]`

- `echo 0 0 | go run .`
  - expected: `Not a quad function`

## Important details

- The program should not print extra text beyond the required result.
- If more than one quad matches, print them in alphabetical order and separate with ` || `.
- Always end the output with a newline.
- `go run .` should work from the project directory.

## Suggested function breakdown

In `main.go`, you can use this structure:

- `func main()`
- `func readInput() (string, error)`
- `func normalizeLines(text string) ([]string, int, int)`
- `func isQuadA(lines []string, w int, h int) bool`
- `func isQuadB(lines []string, w int, h int) bool`
- `func isQuadC(lines []string, w int, h int) bool`
- `func isQuadD(lines []string, w int, h int) bool`
- `func isQuadE(lines []string, w int, h int) bool`
- `func matchRect(lines []string, w int, h int, chars patternChars) bool`

Where `patternChars` is a small helper type or struct describing:

- `topLeft`
- `topRight`
- `bottomLeft`
- `bottomRight`
- `horizontal`
- `vertical`

This is optional, but it helps keep the checker logic consistent.

## Conclusion

Follow these steps and you will have a complete `quadchecker` program:

1. create `go.mod`
2. create `main.go`
3. read input from `os.Args` or `os.Stdin`
4. split and normalize lines
5. compute width and height
6. implement the five quad pattern checks
7. collect matches, sort, join, and print
8. print `Not a quad function` when there are no matches

If you follow the exact character rules for each quad and handle width/height edge cases, the program will work correctly.
