# sortinputs
A short project to consolidate and sort input lines from any number of files

## My Approach
This project is constrained to newline-delimited text files, and it creates an output file at the given path containing a full set of unique text lines across all input files, lexographically sorted (like a dictionary!) without blank lines.

The following assumptions/creative liberties were taken:
* Blank lines are omitted, and so are blank spaces before text. Leading whitespace is trimmed.
* "Comments" in the text file are treated as text, and recognized as valid input.
* The sorting is case-sensitive. For example, Apple" comes before "Banana" comes before "apple".
* Only .txt files are considered. Binary files (and others) are ignored.
* Special characters, outside of [A-Za-z] are included in the text lines. I assume that anything in `*.txt` is meant to be read as plain text.

Since I favored scalability & performance of my program, I chose to implement this assignment with multiple thread processes. I wanted to ensure efficient processing of any number of text files in parallel. My provided solution is in Python, in the `python_solution` directory (but I also included my initial Golang implementation in the `*.go `files).

The Python solution uses Python's "threading" library to process each individual input file concurrently, as the program walks through the input directory. These threads read the file and add non-blank lines to a shared queue (a thread-safe data structure). From this queue, I add (and lexographically sort) unique lines into a SortedSet in a single step, since SortedSet maintains order during insert. For error handling, the Python code returns exceptions.

## Notes to a code reviewer
I chose to include my Golang solution here because I immediately thought of the solution in Golang before I translated it over to Python. If the reviewer wants insight into how I solved the problem, I hope to be transparent by providing my first draft `*.go` code. 

As my first iteration of this assignment, the Golang solution takes advantage of Golang's built-in "goroutines" and "channels" for its efficient concurrent file processing to handle a large number of files in parallel without significant performance bottlenecks. The code uses built-in libraries for the file I/O, sorting, and concurrency concepts. Since each goroutine maps 1-to-1 to a file to process, this code will scale well & use more computation resources as the number of files & size of files increase. For error handling, the Golang code returns errors.

Additionally, to play devil's advocate, if the highest priority for this project is its scalability, performance, and quick implementation, Golang generally has both better metrics than Python when it comes to concurrency/threading on CPUs, and it's easier to read/write than C++ for the sake of a simple take-home assignment.

### Test cases
I included three sets of test cases: 
* a "simple" case that matched the example provided in the assignment write-up
* a "complex" case with multiple sub-directories and files
* a "chaos" case where I tried to intentionally introduce edge cases to trip up my program

## Instructions
### Python
To run this program, have Python installed on your environment, download the directory, and run `pip install -r python_solution/requirements.txt` from the root directory. Then, run `python python_solution/main.py <input-directory-path> <output-file-path>` on your input files and given output file path. To test against all provided test cases, run `pytest` from the root directory.

### Golang
To use this program, download the directory and run `go build` from the `go_solution` directory. (You may have to download Golang & any other number of pre-requisites.) Run `./sortinputs <input-directory-path> <output-file-path>` from the `go_solution` directory to consolidate & sort the data into the given output file path. To run all tests, run `go test` from the `go_solution` directory. To run against any of my provided test files, say, the "chaos" test, run `./sortinputs tests/chaos tests/chaos-output.txt` from the `go_solution` directory. Then, compare output against `tests/chaos.txt`.

## Analysis
The program's space complexity is O(*n*), where *n* is the total number of lines, to account for memory usage in the Python queue/Golang channel & Python SortedSet/Golang map.
The program's time complexity is O(*n*log*n*), where *n* is the total number of lines processed across all files. The file processing & unique line collection steps dominate the combined time complexity between the file traversal, file processing, line collection, and writing the output. (Assuming that *m*, the # of unique lines, is less than *n*, the total # of lines, Golang has slightly better time complexity in that it only sorts the map of unique lines, not total lines. Thus, its time complexity would be O(*n*log*m*))

