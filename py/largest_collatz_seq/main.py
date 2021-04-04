'''
Longest Collatz squence

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
'''

import asyncio
import aiomultiprocess as aio

upper_bound = 1_000_000


async def main():
    longest_sequence = -1
    async with aio.Pool() as pool:
        res = await pool.map(collatz, list(range(2, upper_bound, 1)))
        results = {d['key']:d['terms'] for d in res}
        longest_sequence = max(results, key=results.get)
    
    print(longest_sequence)
    return longest_sequence, results


async def collatz(number):
    terms = 1
    x = number
    while x > 1:
        if x % 2 == 0:
            # even n/2
            x /= 2
        else:
            # odd number 3n + 1
            x = (3*x) + 1
        terms += 1
    return {'key': number, 'terms': terms}


if __name__ == '__main__':
    key, results = asyncio.run(main())
    print(key, results[key])