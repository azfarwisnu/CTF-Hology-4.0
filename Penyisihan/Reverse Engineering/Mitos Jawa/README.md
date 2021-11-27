# Mitos Jawa (498 Pts)

## Deskripsi
Ada mitos yang mengatakan bahwa "itu" tidak reversible..

Author : **Yesver**

## Hint
- Cari tahu bagaimana "itu" bekerja

## Proof of Concept
String asal dapat ditemukan dengan melakukan brute-force pada flag, dengan optimisasi sebagai berikut:

1. Beri padding '0' di akhir bruteforced string, misal 'bru00000', lalu hitung nilai hash-nya. Jika lebih besar dari nilai hashcode yang diketahui, maka bruteforced string itu bukan bagian dari flagnya.
2. Beri padding 'z' di akhir bruteforced string, misal 'bruzzzzz', lalu hitung nilai hash-nya. Jika lebih kecil dari nilai hashcode yang diketahui, maka bruteforced string itu bukan bagian dari flagnya.
3. Jika nilai hashCode bruteforced string == nilai hashCode dari soal, maka cek apakah nilai hashCode uppercase(bruteforced string) == nilai hashCode uppercase dari soal. Jika iya, maka itu adalah flagnya (kalo readable, soalnya ada string lain yang nilai hashCodenya sama persis).
4. Jika tidak ada bruteforced string yang memenuhi, maka nilai hashCode dari soal bisa ditambah 2**32 dan lakukan bruteforce dari awal, karena hashCode sudah terkena integer overflow, sehingga kita perlu melakukan bruteforce pada hashCode juga.

Solver dapat diunduh di: https://mega.nz/file/raA3lCyA#rdw8N3cEMPHMdviMFac0Iwg77Ibhuwz2Wb5oNL5y3dI

<details>
	<summary>Flag : </summary>
	hology4{brut3_1t}
</details>

## Referensi
- https://ctflearn.com/challenge/197 (inspirasi soal)
- https://hg.openjdk.java.net/jdk8/jdk8/jdk/file/687fd7c7986d/src/share/classes/java/lang/String.java
