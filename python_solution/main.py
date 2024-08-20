import os
import sys
import pathlib
import threading
from queue import Queue
from sortedcontainers import SortedSet

def consolidate_data(input_dir, output_file):
    # initialize variables
    lines = Queue()
    unique_lines = SortedSet()

    # check that the input_dir is not empty/can be walked
    if not os.path.isdir(input_dir):
        raise NotADirectoryError(f"input directory {input_dir} is not a directory")

    # collect all the lines
    threads = []
    for root, _, files in os.walk(input_dir):
        for file in files:
            path = os.path.join(root, file)
            # each thread processes a file in parallel
            t = threading.Thread(target=process_file, args=(path,lines))
            t.start()
            threads.append(t)

    # wait for all threads to finish
    for t in threads:
        t.join()

    # collect all unique lines
    while not lines.empty():
        unique_lines.add(lines.get())

    # write the output file & return error/nil
    try:
        with open(output_file, 'w') as file:
            for line in unique_lines:
                file.write(f"{line}\n")
    except Exception as e:
        raise IOError(f"error writing file {output_file}: {e}")
    
# function to process each file
def process_file(fp, lines):
    if pathlib.Path(fp).suffix != ".txt":
        return # Skip files that aren't .txt
    try:
        with open(fp, 'r') as file:
            for line in file:
                line = line.strip()
                if line: # Ignore blank lines
                    lines.put(line)
    except Exception as e:
        print(f"error opening file {fp}: {e}")

def main():
    if len(sys.argv) != 3:
        print(f"to run this program, `python main.py <input_directory> <output_file>`")
        sys.exit(1)

    consolidate_data(sys.argv[1], sys.argv[2])

if __name__ == "__main__":
    main()