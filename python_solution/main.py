import os
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
        raise NotADirectoryError(f"Input directory {input_dir} is not a directory")

    # collect all the lines
    threads = []
    for root, _, files in os.walk(input_dir):
        for file in files:
            path = os.path.join(root, file)
            thread = threading.Thread(target=process_file, args=(path,lines))
            thread.start()
            threads.append(thread)

    # wait for all threads to finish
    for thread in threads:
        thread.join()

    # collect all unique lines
    while not lines.empty():
        line = lines.get()
        unique_lines.add(line)

    # write the output file & return error/nil
    try:
        with open(output_file, 'w') as file:
            for line in unique_lines:
                file.write(f"{line}\n")
    except Exception as e:
        raise IOError(f"Error creating output file {output_file}: {e}")
    
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
        print(f"Error opening file {fp}: {e}")