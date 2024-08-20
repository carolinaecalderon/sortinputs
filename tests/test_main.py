import pytest
from python_solution.main import consolidate_data

@pytest.mark.parametrize("name, input_dir, output_file, output_actual", [
    ("simple", "tests/simple", "tests/simple-output.txt", "tests/simple.txt"),
    ("complex", "tests/complex", "tests/complex-output.txt", "tests/complex.txt"),
    ("chaos", "tests/chaos", "tests/chaos-output.txt", "tests/chaos.txt"),
])
def test_consolidate_data(name, input_dir, output_file, output_actual):
    # Run the function
    err = consolidate_data(input_dir, output_file)
    assert err is None

    # Read the contents of the files
    with open(output_file, "r") as f1, open(output_actual, "r") as f2:
        expected = f1.read()
        actual = f2.read()

    # Compare the contents
    assert expected == actual
