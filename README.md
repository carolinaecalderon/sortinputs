# sortinputs
a short project to consolidate and sort input lines from any number of files

## My Approach
For this assignment, I favored 

## Instructions
To use this program, download the Github directory and run ``go run main.go <input-directory-path> <output-file-path>`` from the root directory. (You may have to download Golang & any other number of pre-requisites.)
Run ``go build``, then run ``./sortinputs <input-directory-path> <output-file-path>`` to consolidate & sort the data into the given output file path.

For example, if you want to test against any of my provided test files, say, the "chaos" test, run ``./sortinputs tests/chaos tests/chaos-output.txt`` from the root directory. Then, compare output against ``tests/chaos.txt``.

This project is constrained to newline-delimited text files, and it createsn an output file at the given path containing a full set of unique text lines across all input files, lexographically sorted (like a dictionary!) without blank lines.

The following assumptions/creative liberties were taken:
(1) Blank lines are omitted, and so are blank spaces before a line. Leading whitespace is trimmed.
(2) "Comments" in the text file are treated as text, and recognized as valid input.
(3) The sorting is case-sensitive. For example, "Apple" comes before "apple".

## Notes to a code reviewer

## Analysis
The program's space complexity

Please consider the algorithmic complexity of your solution carefully. The
solution should scale both with respect to the number of input files and the
size of files.

