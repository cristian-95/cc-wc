# cc-wc

My solution to the challenge [Build Your Own wc Tool](https://codingchallenges.fyi/challenges/challenge-wc/).

## Requirements

- [Go](https://golang.org/dl/) (version 1.16 or higher).

## Usage

1. Build the program:

   ```bash
   go build ccwc.go
   ```

2. Run the program with the desired options:

   ```bash
   ./ccwc [options] <file>
   ```
   or

   ```bash
   cat <file> | ./ccwc [options]
   ```


#### Option flags

- **-c** to count bytes
- **-l** to count lines
- **-w** to count words
- **-m** to count chars



