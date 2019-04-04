# -*- coding: utf-8 -*-

from random import sample, choice
from time import time


def linear_search(number: int, target: list) -> bool:
    start_time = time()

    for i, n in enumerate(target):
        if n == number:
            print("linear searched in %0.10fs." % (time() - start_time))
            return True

    print("linear searched in %0.10fs." % (time() - start_time))

    return False


def binary_search(number: int, target: list) -> int:
    target = sorted(target)

    start_time = time()

    while 1:
        if len(target) < 1:
            break

        center = 0 if len(target) == 1 else int(len(target) / 2)

        if target[center] == number:
            print("binary searched in %0.10fs." % (time() - start_time))
            return True
        elif target[center] > number:
            target = target[:center]
        elif target[center] < number:
            target = target[(center + 1):]

    print("binary searched in %0.10fs." % (time() - start_time))
    return False


if __name__ == "__main__":
    numbers = list(range(10000))
    shuffle_numbers = sample(numbers)
    target_number = choice(numbers)
    shuffle_numbers = shuffle_numbers[:(int(len(shuffle_numbers) / 2))]
    print(target_number, shuffle_numbers, target_number in shuffle_numbers)

    found = linear_search(target_number, shuffle_numbers)
    print(found)

    found = binary_search(target_number, shuffle_numbers)
    print(found)
