from functools import cmp_to_key
from typing import Union


def compare(l: Union[list, int], r: Union[list, int]) -> int:
    if isinstance(l, int):
        l = [l]
    if isinstance(r, int):
        r = [r]

    for a, b in zip(l, r):
        if isinstance(a, list) or isinstance(b, list):
            if (cmp := compare(a, b)) != 0:
                return cmp
            continue
        if a < b:
            return 1
        if a > b:
            return -1

    if len(l) > len(r):
        return -1
    if len(l) < len(r):
        return 1
    return 0


def part1():
    with open("../../inputs/day13.txt") as f:
        i = 1
        su = 0
        while l1 := f.readline():
            l2 = f.readline()
            f.readline()
            ll1 = eval(l1)
            ll2 = eval(l2)
            if compare(ll1, ll2) >= 0:
                su += i
            i += 1
    print(su)


def cmp_func(l: Union[list, int], r: Union[list, int]) -> int:
    cmp = compare(l, r)
    return 1 if cmp < 0 else -1


def part2():
    packets = [[[2]], [[6]]]
    with open("../../inputs/day13.txt") as f:
        packets.extend(eval(l) for l in f.readlines() if l != '\n')
    packets.sort(key=cmp_to_key(cmp_func))
    print((packets.index([[2]]) + 1) * (packets.index([[6]]) + 1))


def main():
    part1()
    part2()


if __name__ == "__main__":
    main()
