# -*- coding: utf-8 -*-

from time import time

from eratosthenes_sieve import eratosthenes_sieve


def simple_factorize(number: int) -> int:
    start_time = time()

    primes = eratosthenes_sieve(number)
    i = 0
    pf = []

    while number != 1:
        if number % primes[i] == 0:
            quotient = int(number / primes[i])
            pf.append(primes[i])
            number = quotient
        else:
            i += 1

    print("simple factorized in %0.10fs." % (time() - start_time))

    return pf


def fast_factorize(number: int):
    start_time = time()

    m = number
    pf = []

    while number != 1:
        m -= 1

        if number % m != 0:
            continue

        d = int(number / m)
        pf.append(d)
        number /= d

        while number % d == 0:
            pf.append(d)
            number /= d

        m = number

    print("fast factorized in %0.10fs." % (time() - start_time))

    return pf


if __name__ == "__main__":
    n = 1024
    print("n", n)

    factors = simple_factorize(n)
    print("factors", factors)

    factors = fast_factorize(n)
    print("factors", factors)
