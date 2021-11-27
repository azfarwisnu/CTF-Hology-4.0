# Wait, That's Illegal (463 Pts)

## Deskripsi
I don't know if that's allowed

Author : Yesver

## Hint
- You might need some brute force for this

## Proof of Concept
Untuk menyelesaikan soal ini, kita perlu melihat hubungan antarbilangan.
1. Karena n2 = n1*i, maka kita bisa melakukan modulo pada c2 dengan n1, dan mengubah n2 menjadi n1.
2. Karena gcd(e1, e2) = 10, kita bisa melakukan common modulus attack untuk mendapatkan plaintext^{10} mod n1.
3. Untuk mendapatkan plaintext, kita bisa melakukan bruteforce.

<details>
    <summary>Flag : </summary>
    hology4{n1c3_m4th_m4t3_xd}
</details>

## Referensi
-
