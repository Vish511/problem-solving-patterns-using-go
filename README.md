# Grokking Coding Patterns: Go Solutions for Common Problem-Solving Patterns

This repository contains Go implementations for solving coding problems from the course "Grokking the Coding Interview: Patterns for Coding Questions." Each folder represents a pattern from the course, and inside each folder, you will find the Go solutions for related problems. The course covers various common patterns that often appear in coding challenges.

## Table of Contents

- [Overview](#overview)
- [Repository Structure](#repository-structure)
- [Patterns Covered](#patterns-covered)
- [How to Run](#how-to-run)
- [Contributing](#contributing)
- [License](#license)

## Overview

The purpose of this repository is to provide clean, efficient, and easy-to-understand solutions for the common coding patterns covered in the *Grokking the Coding Interview* course. By implementing these solutions in Go, we aim to demonstrate problem-solving strategies and best practices in Go.

## Repository Structure

The repository is organized as follows:

```bash
├── main.go                 # Entry point for running the program
├── slidingwindow           # Folder for the first pattern (e.g., Sliding Window)
│   ├── longestsubstringwithsameletters.go
│   └── averageofsubarray.go
├── fastslowpointers        # Folder for the second pattern (e.g., Fast & Slow Pointers)
│   ├── happynumber.go
│   └── cycleinacirculararray.go
├── mergeintervals          # Folder for the third pattern (e.g., Merge Intervals)
│   ├── maximumcpuload.go
│   └── conflictingappointments.go
├── helpers                 # Folder for helper functions (e.g., Math functions, sorting functions)
│   ├── math.go
│   └── sort.go
└── README.md               # This README file
```

### `main.go`

The `main.go` file is the entry point for running and testing different pattern solutions. To run a specific problem, you need to **uncomment** the corresponding line in `main.go`. Once the desired problem is uncommented, run the program to see the output.

---

## Patterns Covered

### 1. **Sliding Window**

The sliding window pattern is useful for problems involving subarrays, substrings, or any contiguous segment of data. Common applications include finding maximum or minimum values, sums, or averages over a window of data.

- Problem 1: Find the maximum sum of a subarray of size K
- Problem 2: Longest substring with K distinct characters
- Problem 3: Substring with no repeating characters
- Problem 4: Smallest subarray with a sum greater than a given value
- …more

### 2. **Fast & Slow Pointers**

The fast & slow pointers pattern helps solve problems involving linked lists or arrays where we need to track two pointers moving at different speeds to detect cycles, find the middle of a list, or determine other important properties.

- Problem 1: Linked list cycle detection
- Problem 2: Middle of a linked list
- Problem 3: Happy Number
- Problem 4: Find the start of a cycle in a linked list
- …more

### 3. **Two Pointers**

The two pointers pattern is used when we need to process arrays or linked lists by having two pointers move towards each other or in tandem to solve problems efficiently.

- Problem 1: Pair with a given sum in a sorted array
- Problem 2: Remove duplicates from a sorted array
- Problem 3: Triplet sum to zero
- Problem 4: Container with most water
- …more

### 4. **Merge Intervals**

The merge intervals pattern helps solve problems involving merging, inserting, or comparing intervals in a sorted order, commonly found in problems related to time scheduling, event handling, or range merging.

- Problem 1: Merge overlapping intervals
- Problem 2: Insert an interval
- Problem 3: Interval intersection
- Problem 4: Find the minimum number of intervals to remove to make the rest non-overlapping
- …more

Each folder contains problem-specific implementations, typically with a `solution.go` or `example.go` file for each problem.

---

## How to Run

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/grokking-coding-patterns-go.git
    cd grokking-coding-patterns-go
    ```

2. **Open the `main.go` file** and **uncomment the line** for the specific problem you wish to run. For example, to run Problem 1 under the Sliding Window pattern, uncomment the corresponding line in `main.go`:

    ```go
    // Uncomment the line below to run Problem 1 from the Sliding Window pattern
    // slidingwindow.Problem1()
    ```

3. **After uncommenting the desired line, run the program:**

    ```bash
    go run main.go
    ```

4. If you want to run a problem from a different pattern, go back to `main.go`, uncomment the corresponding line, and rerun the program.

---

## Installing Go

If you don't have Go installed, follow these steps to install Go:

- [Go Installation Guide](https://golang.org/doc/install)

Once Go is installed, you can run the program as described above.

---

## Contributing

Contributions are welcome! If you would like to contribute solutions to other patterns or improve the existing solutions, feel free to fork the repository and create a pull request.

### Steps to contribute:
1. Fork the repository
2. Create a new branch
3. Add your solution or improvements
4. Submit a pull request

Make sure to follow Go's best practices and write clean, efficient code.

---

## License

This repository is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
```

### Changes:
- Code blocks are now properly separated into language-specific blocks (`bash` and `go`).
- The code under the `bash` block and `go` block now renders properly.
  
This version should now be correctly formatted for GitHub. Let me know if you need anything else!
