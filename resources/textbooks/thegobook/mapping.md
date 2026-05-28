# Introducing Go — Chapter-to-Repo Mapping

This document maps the textbook chapters to the existing repository folders and study topics.

## Book chapter → repository mapping

- `1. Getting Started` → `curriculum/01_hello`, `curriculum/02_variables`, `curriculum/03_control_flow`
  - Covers environment, terminal, first program, and reading Go code.

- `2. Types` → `curriculum/02_variables`, `curriculum/04_collections`
  - Maps numeric, string, boolean, and basic type concepts.

- `3. Variables` → `curriculum/02_variables`
  - Covers naming, scope, constants, and variable declarations.

- `4. Control Structures` → `curriculum/03_control_flow`
  - Covers `for`, `if`, and `switch`.

- `5. Arrays, Slices, and Maps` → `curriculum/04_collections`
  - Directly maps collection concepts and common slice/map operations.

- `6. Functions` → `curriculum/05_functions`
  - Includes normal functions, variadic functions, closure, recursion, and pointers.

- `7. Structs and Interfaces` → `curriculum/06_pointers_structs`, `curriculum/07_methods_interfaces`
  - Maps structs, methods, embedded types, and interfaces.

- `8. Packages` → `curriculum/13_packages`, `curriculum/11_standard_library`
  - Covers core packages, strings, IO, files, errors, containers, sort, and documentation.

- `9. Testing` → `curriculum/10_testing`
  - Maps to testing concepts and practical unit tests.

- `10. Concurrency` → `curriculum/09_concurrency`, `curriculum/14_context`
  - Maps goroutines, channels, select, and concurrent patterns.

- `11. Next Steps` → `curriculum/15_final_project`
  - Guidance for building projects, collaborating, and practical next steps.

## How to use this mapping

- Open `resources/textbooks/thegobook.pdf` and find a chapter.
- Use this document to jump to the repo folder with exercises and examples.
- Add your own chapter notes in `resources/textbooks/thegobook/` if you want deeper integration.

## Integration recommendation

- Keep the PDF as the canonical textbook.
- Use `toc.md` for quick reference.
- Use `mapping.md` to keep your study path aligned with the repo structure.
- Optionally, add a `chapter-notes` file for each textbook chapter if you want searchable summaries.
