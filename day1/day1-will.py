# Advent of Code 2023
# Day 1
# Will Harrington
# Python 3.12

with open('input.txt') as file:
    lines = file.readlines()

    # Part 1
    part_one = sum(
        int((digits := [char for char in line if char.isdigit()])[0] + digits[-1])
        for line in lines
    )

    # Part 2
    NUMBERS = {
        'one': 1,
        'two': 2,
        'three': 3,
        'four': 4,
        'five': 5,
        'six': 6,
        'seven': 7,
        'eight': 8,
        'nine': 9,
    }
    part_two = 0
    for line in lines:
        substring = ''
        for char in line:
            if char.isdigit():
                first = int(char)
                break
            else:
                substring += char
                for number, value in NUMBERS.items():
                    if number in substring:
                        first = value
                        break
                else:
                    continue
                break

        substring = ''
        for char in reversed(line):
            if char.isdigit():
                last = int(char)
                break
            else:
                substring = char + substring
                for number, value in NUMBERS.items():
                    if number in substring:
                        last = value
                        break
                else:
                    continue
                break
        calibration = 10 * first + last
        part_two += calibration

print(f'{part_one=}')
print(f'{part_two=}')
