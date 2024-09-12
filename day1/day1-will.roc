# Advent of Code 2023
# Day 1
# Will Harrington
# roc 9/9/24
#
app [main] { pf: platform "https://github.com/roc-lang/basic-cli/releases/download/0.15.0/SlwdbJ-3GR7uBWQo6zlmYWNYOxnvo8r6YABXD-45UOw.tar.br" }

import pf.Stdout
import "input.txt" as input : Str

main =
    Stdout.line! "Part One: $(Num.toStr partOne)"
# Stdout.line! "Part One: $(Num.toStr partOne2)"

# |> Str.joinWith ","
# |> List.map Num.toStr
# |> List.join

partOne =
    # this is so cringe
    input
    |> Str.split "\n"
    |> List.map \line -> Str.walkUtf8 line [] \state, byte -> List.appendIfOk state (Str.fromUtf8 [byte])
    |> List.map \line -> List.keepOks line \char -> Str.toU8 char
    |> List.map \line -> [List.first line, List.last line]
    |> List.map \line -> List.keepOks line \byte -> byte
    |> List.map \line -> Str.joinWith (List.map line \num -> Num.toStr num) ""
    |> List.keepOks \line -> Str.toU32 line
    |> List.sum

# partOne2 =
#     lineNumbers = List.map (Str.split input "\n") \line ->
#         Str.walkUtf8 line [] \state, byte -> List.appendIfOk state (Str.fromUtf8 [byte])
#         |> List.keepOks \char -> Str.toU8 char

#         [List.first line, List.last line]

# |> List.map \line -> 10 * (List.first line)? + (List.last line)?
