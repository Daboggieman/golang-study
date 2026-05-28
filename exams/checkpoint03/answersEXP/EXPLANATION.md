# Checkpoint 03 Answers — Detailed Explanation

This document describes each file in `exams/checkpoint03/answersEXP`, including the exact actions taken, the key algorithms/keywords used, and the logic behind the implementation.

---

## 01_split.go

### Action taken
- Implemented the `Split(s, sep string) []string` function from scratch.
- Did not use the `strings` package.
- Used manual string scanning and substring slicing.

### Keywords and logic
- `len(sep)`: checked for the empty separator edge case.
- `for i := 0; i+sepLen <= len(s); i++`: iterated over possible separator start positions.
- `s[i:i+sepLen] == sep`: compared each candidate substring to the separator.
- `append(result, s[start:i])`: saved the substring before each separator.
- `start = i + sepLen`: moved the start index past the separator.
- `i += sepLen - 1`: skipped the separator bytes to avoid overlapping matches.
- Final `append(result, s[start:])`: appended the last remaining segment.

### Why this matches Piscine style
- Uses only built-in operations and indexing.
- Avoids forbidden `strings` package.
- Returns a slice of substrings exactly as required.

---

## 02_activebits.go

### Action taken
- Implemented `ActiveBits(n int) int` using bitwise operations.
- Converted `n` into `uint` to support safe shifting and avoid negative bitwise behavior.

### Keywords and logic
- `uint(n)`: safely convert to unsigned for bit scanning.
- `for un > 0`: loop while there are bits left.
- `un&1 == 1`: tests the lowest bit using bitwise AND.
- `count++`: increments the active bit count.
- `un >>= 1`: shifts the number right to test the next bit.

### Why this matches Piscine style
- Uses bitwise operators only, as the exercise mandates.
- Does not use `math/bits` or any disallowed helper packages.

---

## 03_sortwordarr.go

### Action taken
- Implemented `SortWordArr(a []string)` using a simple in-place bubble sort.
- Sorted strings by ASCII order.

### Keywords and logic
- `for i := 0; i < len(a); i++`: outer loop over elements.
- `for j := 0; j < len(a)-1-i; j++`: inner loop to compare adjacent pairs.
- `if a[j] > a[j+1]`: lexicographic ASCII comparison.
- `a[j], a[j+1] = a[j+1], a[j]`: swap values in-place.

### Why this matches Piscine style
- Avoids `sort` package by implementing sorting manually.
- Works in-place and does not allocate a new slice beyond the swap operation.

---

## 04_advancedsortwordarr.go

### Action taken
- Implemented `AdvancedSortWordArr(a []string, f func(a, b string) int)`.
- Used the passed comparator function `f` instead of direct string operators.

### Keywords and logic
- `f(a[j], a[j+1]) == 1`: compare elements through the callback.
- Swap logic is identical to bubble sort, but condition is driven by `f`.
- The comparator returns `1` when the first element should be considered greater.

### Why this matches Piscine style
- Tests comprehension of function types as parameters.
- Uses the comparator exactly as described by the exercise.

---

## 05_listpushback.go

### Action taken
- Implemented `ListPushBack(l *List, data interface{})`.
- Added new nodes to the end of a linked list.

### Keywords and logic
- `&NodeL{Data: data}`: created a new node.
- `if l.Head == nil`: handled an empty list by setting both `Head` and `Tail`.
- `l.Tail.Next = newNode`: appended to the existing tail.
- `l.Tail = newNode`: updated the tail pointer.

### Why this matches Piscine style
- Uses standard linked-list node manipulation.
- Handles the empty-list special case correctly.

---

## 06_listsize.go

### Action taken
- Implemented `ListSize(l *List) int`.
- Traversed the list from `Head` to `nil`.

### Keywords and logic
- `current := l.Head`: starts traversal at the head.
- `for current != nil`: loop through list nodes.
- `count++`: increments for each node.
- `current = current.Next`: advances to the next node.

### Why this matches Piscine style
- Simple linked list traversal without helper packages.
- Returns `0` for empty lists.

---

## 07_listlast.go

### Action taken
- Implemented `ListLast(l *List) interface{}`.
- Returned tail data directly when available.

### Keywords and logic
- `if l.Tail != nil`: uses the tail pointer to access the last node quickly.
- Fallback traversal when `Tail` is unexpectedly nil.
- Returns `nil` for empty lists.

### Why this matches Piscine style
- Uses the `List` structure as intended.
- Works even if the tail pointer is missing.

---

## 08_listreverse.go

### Action taken
- Implemented `ListReverse(l *List)` to reverse pointers in place.
- Updated both `Head` and `Tail`.

### Keywords and logic
- `prev`, `current`, `next`: standard pointer reversal variables.
- `current.Next = prev`: reverses the pointer.
- `prev = current`: moves the previous pointer forward.
- `current = next`: continues traversal.
- `l.Tail = l.Head`: preserves original head as the new tail.
- `l.Head = prev`: updates the head after traversal finishes.

### Why this matches Piscine style
- Reverses the list by pointer manipulation, not by swapping node contents.
- Correctly handles empty and single-node lists.

---

## 09_listremoveif.go

### Action taken
- Implemented `ListRemoveIf(l *List, data_ref interface{})`.
- Removed all nodes whose `Data` matches `data_ref`.

### Keywords and logic
- `for l.Head != nil && l.Head.Data == data_ref`: removes matching head nodes.
- `current.Next = current.Next.Next`: removes interior nodes.
- `l.Tail = current`: updates the tail after removals.

### Why this matches Piscine style
- Handles consecutive removals and the empty-list case.
- Maintains proper list integrity.

---

## 10_btreeinsertdata.go

### Action taken
- Implemented `BTreeInsertData(root *TreeNode, data string) *TreeNode`.
- Built a binary search tree and preserved `Parent` pointers.

### Keywords and logic
- `if root == nil`: handles insertion into an empty tree.
- `if data < root.Data`: follows left-child insertion rules.
- `else`: sends equal and greater values to the right child.
- Recursive calls continue down the tree until insertion point is found.

### Why this matches Piscine style
- Uses the expected binary search tree ordering rules.
- Maintains the parent pointer for every inserted node.

---

## 11_doop/main.go

### Action taken
- Converted the `doop` program to use 01-edu printing conventions.
- Replaced `os.Stdout.WriteString` with `github.com/01-edu/z01` output.
- Added robust numeric parsing and 64-bit overflow detection.

### Keywords and logic
- `import "github.com/01-edu/z01"`: uses the Piscine-approved printing package.
- `PrintStr` / `PrintLn`: print strings character-by-character with `z01.PrintRune`.
- `Atoi(s string) (int64, bool)`: validates input and parses signed numbers.
- `maxPos` / `maxNeg` / `limit`: detect 64-bit overflow and fail safely.
- `Itoa(n int64) string`: converts integer values back into digits.
- `switch op`: selects `+`, `-`, `*`, `/`, `%` operations.
- Zero-division handling:
  - `No division by 0`
  - `No modulo by 0`

### Why this matches 01-edu protocols
- Uses `z01.PrintRune` instead of standard library output.
- Implements custom parsing and formatting without `fmt` or `strconv`.
- Prints results and errors in the expected Piscine style.

---

## Summary of protocol compliance

- All function solutions are in `package piscine` where appropriate.
- The `doop` program is in `package main` and uses only 01-edu-approved printing.
- Explanations are explicitly recorded in `EXPLANATION.md`.
- The answer folder now includes both code and detailed reasoning.
