// Advent of Code 2023
// Day 1
// Will Harrington
// C
//

#include <stdio.h>
#include <stdlib.h>

int part_one_ansi(FILE *file)
{
    int sum = 0;
    char c;
    int num;
    int first;
    int last;
    while ((c = fgetc(file)) != EOF)
    {
        // WHY does this return 46 instead of 4 on the 2nd iteration?!
        num = atoi(&c);
        if (num != 0)
        {
            first = num;
            last = num;
        }
        while ((c = fgetc(file)) != '\n' && c != EOF)
        {
            num = atoi(&c);
            if (num != 0)
                last = num;
        }
        sum += 10 * first + last;
    }
    return sum;
}

int main()
{
    FILE *file = fopen("input.txt", "r");
    if (file == NULL)
        exit(1);

    // Part 1
    printf("Part 1 (ANSI): %d\n", part_one_ansi(file));
}
