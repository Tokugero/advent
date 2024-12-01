from solution import *

def test_test() -> None:
    assert solution1(parseInput("test.txt")) == 11

def test_prompt() -> None:
    assert solution2(parseInput("test.txt")) == 31