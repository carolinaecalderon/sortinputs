import os
import pytest
from python_solution.main import consolidate_data

@pytest.mark.parametrize("name", ["simple", "complex", "chaos"])
def test_consolidate_data(name):
    input_dir = "tests/" + name
    output_expected = input_dir + ".txt"
    output_file = input_dir + "-output.txt"

    # Run the function
    err = consolidate_data(input_dir, output_file)
    assert err is None

    # Read the contents of the files
    with open(output_file, "r") as f1, open(output_expected, "r") as f2:
        expected = f1.read()
        actual = f2.read()

    # Compare the contents
    assert expected == actual

    # Delete the test files
    try:
        os.remove(output_file)
    except OSError as e:
        print(f"error deleting file {output_file}:{e}")
