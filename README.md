# Text Processor – README

## Overview

This Go program processes a text file and applies a series of transformations, including:

* Capitalization modifications via markup commands
* Lowercasing and uppercasing
* Binary and hexadecimal number conversions
* Grammar correction (`a` ↔ `an`)
* Merging punctuation with preceding words
* Handling quotes and special punctuation (`...`, `!?`)

The program reads from an input file, applies all transformations, and writes the processed result to an output file.

---

## Usage

```bash
go run . <input_file> <output_file>
```

### Example

```bash
go run . input.txt output.txt
```

If the argument count is not exactly 2, the program will print:

```
Error !
```

---

## Features

### 1. **Markup Commands**

The program recognizes inline markup instructions and applies them to previous words:

| Command               | Meaning                      | Examples                |
| --------------------- | ---------------------------- | ----------------------- |
| `(cap)`               | Capitalize the previous word | `hello (cap)` → `Hello` |
| `(cap,3)` or `(cap3)` | Capitalize previous 3 words  |                         |
| `(low)`               | Lowercase previous word      |                         |
| `(up)`                | Uppercase previous word      |                         |

All three markup types support:

* `(cmd)` (affects one previous word))
* `(cmd3)` (number appended)
* `(cmd,3)` (number separated by comma)
* `(cmd ,3)` variants where the number is in the next word

### 2. **Hexadecimal and Binary Conversion**

| Command | Effect                                           |
| ------- | ------------------------------------------------ |
| `(hex)` | Converts the previous word from hex → decimal    |
| `(bin)` | Converts the previous word from binary → decimal |

Example:

```
FF (hex) → 255
1010 (bin) → 10
```

### 3. **Punctuation Handling**

The program attaches punctuation to the previous word when appropriate:

* `, ; ! ? : .`
* Ellipses (`...`)
* Combined punctuation (`!?`)

Example:

```
Hello , world ! → Hello, world!
```

### 4. **Quote Handling**

The program merges `'` into proper positions to fix orphaned or misplaced apostrophes or quote marks.

### 5. **Grammar Fix: a ↔ an**

Corrects the article based on the vowel sound of the following word:

* `a` → `an` before vowel sounds
* `an` → `a` before consonant sounds

Example:

```
a apple → an apple
an dog → a dog
```

### 6. **Final Cleanup**

The program removes empty tokens, merges leftover punctuation, and ensures words are spaced correctly.

---

## Input / Output Behavior

### Input

* A text file with words separated by spaces.
* May include markup commands, punctuation, quotes, binary/hex numbers.

### Output

* A processed, cleaned, grammatically adjusted text file.
* All transformations are applied in the order defined in the code.

### Console Output

The program prints the original and final processed text:

```
Initially: <contents of input file>

Result: <processed output>
```

---

## Code Structure

### Key Functions

* `Cap()`, `Low()`, `Up()` — modify case for previous words
* `HexToDec()`, `BinToDec()` — number conversions
* `CleanStr()` — remove empty tokens
* `punc()` and `fend()` — advanced punctuation cleanup
* `isWord()` / `notWord()` — determine word contents
* `TakeNumFromString()` — extract numeric arguments from markup

---

## Requirements

* Go 1.18+ recommended
* Read/write permissions for input/output files

---

## Notes

* This program does not currently handle errors from `os.ReadFile` or `strconv` conversions.
* Behavior assumes space-separated tokens—punctuation and quotes are cleaned manually by the algorithm.
