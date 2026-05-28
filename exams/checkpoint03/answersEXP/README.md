# Checkpoint 03 Answers (Expanded)

This folder contains complete solutions for the `checkpoint03` exam problems in `exams/checkpoint03`.

## Files

- `01_split.go` - `Split` implementation without using the `strings` package.
- `02_activebits.go` - `ActiveBits` implementation using bitwise operations.
- `03_sortwordarr.go` - `SortWordArr` in-place ASCII string sort.
- `04_advancedsortwordarr.go` - `AdvancedSortWordArr` using a comparison function.
- `05_listpushback.go` - `ListPushBack` linked-list tail insertion.
- `06_listsize.go` - `ListSize` linked-list size calculation.
- `07_listlast.go` - `ListLast` returns the last node's data.
- `08_listreverse.go` - `ListReverse` reverses a linked list in place.
- `09_listremoveif.go` - `ListRemoveIf` removes matching nodes from a linked list.
- `10_btreeinsertdata.go` - `BTreeInsertData` BST insertion with parent pointers.
- `11_doop/main.go` - `doop` CLI calculator with custom parsing and z01 output.
- `EXPLANATION.md` - Detailed explanation of logic, actions, and keywords used.

## How to verify

### Run tests for the package

```bash
cd /home/student/golang-study
go test ./exams/checkpoint03/answersEXP/...
```

### Run the `doop` program

```bash
go run ./exams/checkpoint03/answersEXP/11_doop 10 + 5
```

Expected output:
- `10 + 5` → `15`
- `10 / 0` → `No division by 0`
- `20 % 0` → `No modulo by 0`

## Notes

- Solutions follow Piscine-style constraints and 01-edu protocol where output uses `github.com/01-edu/z01` instead of standard library printing.
- The `doop` program uses a custom `Atoi` and `Itoa`, handles invalid arguments, and prints results/errors via `z01.PrintRune`.
- See `EXPLANATION.md` for a file-by-file breakdown of the implementation, keywords, logic, and actions taken.
