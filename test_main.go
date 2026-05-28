package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
    "os/exec"
    "path/filepath"
    "sort"
    "strings"
)

func main() {
    rootDir, err := os.Getwd()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    input, err := readInput()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    input = normalizeText(input)
    if input == "" {
        fmt.Println("Not a quad function")
        return
    }

    width, height := dimensions(input)
    if width == 0 || height == 0 {
        fmt.Println("Not a quad function")
        return
    }

    quadFiles := map[string]string{
        "quadA": "quadA.go",
        "quadB": "quadB.go",
        "quadC": "quadC.go",
        "quadD": "quadD.go",
        "quadE": "quadE.go",
    }

    matches := []string{}
    for name, file := range quadFiles {
        absPath := filepath.Join(rootDir, file)
        if _, err := os.Stat(absPath); err != nil {
            continue
        }

        output, err := runQuad(rootDir, absPath, name, width, height)
        if err != nil {
            continue
        }
        if normalizeText(output) == input {
            matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", name, width, height))
        }
    }

    if len(matches) == 0 {
        fmt.Println("Not a quad function")
        return
    }

    sort.Strings(matches)
    fmt.Println(strings.Join(matches, " || "))
}

func readInput() (string, error) {
    if len(os.Args) > 1 {
        return strings.Join(os.Args[1:], " "), nil
    }

    data, err := io.ReadAll(os.Stdin)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func normalizeText(text string) string {
    return strings.TrimRight(text, "\n")
}

func dimensions(text string) (int, int) {
    lines := strings.Split(text, "\n")
    width := 0
    for _, line := range lines {
        if len(line) > width {
            width = len(line)
        }
    }
    return width, len(lines)
}

func runQuad(rootDir, absFile, quadName string, width, height int) (string, error) {
    wrapper, err := os.CreateTemp(rootDir, "quadrun-*.go")
    if err != nil {
        return "", err
    }
    defer os.Remove(wrapper.Name())

    funcName := strings.Title(strings.TrimPrefix(quadName, "quad"))
    wrapperCode := fmt.Sprintf(`package main

func main() {
    %s(%d, %d)
}
`, funcName, width, height)

    if _, err := wrapper.WriteString(wrapperCode); err != nil {
        wrapper.Close()
        return "", err
    }
    if err := wrapper.Close(); err != nil {
        return "", err
    }

    cmd := exec.Command("go", "run", wrapper.Name(), absFile)
    cmd.Dir = rootDir
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = nil
    err = cmd.Run()
    return out.String(), err
}
