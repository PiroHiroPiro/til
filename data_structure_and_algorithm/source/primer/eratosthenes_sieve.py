# -*- coding: utf-8 -*-


def eratosthenes_sieve(number: int) -> list:
    m = 2
    data = [1] * (number + 1)
    data[0] = 0
    result = []

    while m <= number:
        i = 2 * m

        while i <= number:
            data[i] = 0
            i += m

        result.append(m)

        m += 1
        while len(data) > m and data[m] != 1:
            m += 1

    return result


if __name__ == "__main__":
    n = 1024
    print("n", n)

    primes = eratosthenes_sieve(n)
    print("primes", primes)
