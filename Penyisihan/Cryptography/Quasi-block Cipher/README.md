# Quasi-block Cipher (500 Pts)

## Deskripsi

I made a new implementation for OTP that utilises blocks and hash functions. Oh boy, i sure do hope nobody can decipher this flag that i just encrypted with this implementation.

## Hint

What is a hash function's worst weakness?


## Proof of Concept

#!/usr/bin/python3

from math import ceil, floor
import hashlib

def byte_xor(ba1, ba2):
    return bytes([_a ^ _b for _a, _b in zip(ba1, ba2)])

with open("ciphertext",'rb') as f_in:
    p = f_in.read()

blocks = ceil(len(p)/32)

c_parts = []
curr = 0
for i in range(blocks):
    c_parts.append(p[curr:curr+32])
    curr += 32

with open('rockyou.txt', 'rb') as dic:
    for line in dic:
        key = line[:-1]
        if len(key) >= blocks:
            keyparts = []
            curr = 0
            amount = floor(len(key)/blocks)

            for i in range(blocks):
                m = hashlib.sha256(key[curr:curr+amount])
                keyparts.append(m.digest())
                curr += amount

            p_parts = []
            for i in range(blocks):
                p_parts.append(byte_xor(c_parts[i], keyparts[i]))

            p = b''.join(p_parts)
            if b'hology4{' in p:
                break

with open('decrypted', 'wb') as f_out:
    f_out.write(p)

<details>
<summary>Flag : </summary>
hology4{whAt_da_dict1oNary_do1n}
</details>

## Referensi

[Hash Function from Wikipedia](https://en.wikipedia.org/wiki/Hash_function)
