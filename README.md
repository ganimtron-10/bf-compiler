# bf-compiler

**bf-compiler** is a fast Brainfk interpreter and interactive REPL written in Golang. Unlike a naive interpreter that parses characters on the fly, this project uses a compilation step to convert Brainfk source into an optimized intermediate representation (IR), significantly boosting execution speed.

## How it Works

The compiler transforms the esoteric source code into a stream of instructions that the virtual machine can execute efficiently.

### Key Optimizations

* **Instruction Aggregation:** Consecutive identical commands like `++++++` or `>>>>` are collapsed into a single instruction with an operand (e.g., `OPInc` with `Operand: 6`), reducing the number of iterations in the execution loop.
* **Clear Loop Optimization:** The compiler identifies the common `[-]` and `[+]` patterns (which set the current cell to zero) and replaces the entire loop with a single `OPClear` operation.
* **Static Jump Mapping:** Bracket pairs `[` and `]` are pre-matched during the compilation phase. The `Operand` for jump instructions contains the exact index of the destination, allowing for $O(1)$ jumps during runtime.

### Technical Specifications

* **Memory:** 300,000 addressable cells (8-bit bytes).
* **Virtual Machine:** A bytecode-driven execution engine.
* **Error Handling:** Built-in detection for "Out of Bounds" memory access and "Unmatched Brackets."

## Supported Commands

| Symbol | Opcode | Description |
| --- | --- | --- |
| `>` | `OPIncPtr` | Move the data pointer to the right. |
| `<` | `OPDecPtr` | Move the data pointer to the left. |
| `+` | `OPInc` | Increment the byte at the data pointer. |
| `-` | `OPDec` | Decrement the byte at the data pointer. |
| `.` | `OPOut` | Output the character at the data pointer. |
| `,` | `OPIn` | Input a character to the data pointer. |
| `[` | `OPJmpFwd` | Jump past `]` if the current cell is 0. |
| `]` | `OPJmpBwd` | Jump back to `[` if the current cell is not 0. |

## Getting Started

### Prerequisites

* Go 1.21+ installed.

### Installation

```bash
git clone https://github.com/ganimtron-10/bf-compiler.git
cd bf-compiler

```

### Running the REPL

Launch the interactive shell to execute Brainfk code line-by-line:

```bash
go run .

```

### Example

Input a simple "Hello" (ASCII 72) in the REPL:

```text
>>> ++++++++[>+++++++++<-]>.
H

```

## Project Structure

* `main.go`: Entry point and REPL logic.
* `compiler.go`: Contains the `Compile` function and optimization logic.
* `vm.go`: The `Execute` engine and memory management.
