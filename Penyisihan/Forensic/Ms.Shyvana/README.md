# Ms.Shyvana (100 Pts)

- [Deskripsi ](#deskripsi)
- [Konsep soal ](#konsep-soal)
- [Hint](#hint)
- [Proof of Concept](#proof-of-concept)

## Deskripsi
Ms.Shyvana menyadap server perusahan toktok dan mendapatkan akses login berupa username dan password namun ia tidak sempat membukanya karena tim fbii sudah menangkap dia. Sebelum tertangkap Ms.Shyvana mengirim sebuah file pcapng dan zip kepadamu dan meminta tolong untuk menemukan password zip dan menyelesaikan teka tekinya. Ms.Shyvana bilang kalau algoritma enkripsi passwordnya adalah rotasi 13.


Format Flag : Hology4{string_yang_ditemukan}

Author :  **fgm**

## Konsep soal
Network Analysis

## HINT
decode

## Proof of Concept
Didapatkan sebuah Capture msShyvana.pcapng  maka dibuka menggunakan wireshark untuk menganalisis

Di dalam deskripsi tersebut di jelaskan bahwasannya ms.shyvana mendapatkan akses login berupa username atau password , maka bisa disimpulkan Ms.Shyvana mendapatkannya dari seseorang yang ingin meremote server. Ketika meremote server umumnya orang menggunakan ssh atau telnet

maka cek protocol ssh atau telnet atau 
cek protocol TCP karena ssh atau telnet menggunakan protocol TCP pada layer transport maka bisa juga kita cek pada protocol TCP

Dan di dapatkan sebuah key FnQ4Cq00aT_Obm22mD_s66Z yang telah di encrypt oleh rot13 setelah menjadi plaintext maka akan menjadi 
SaD4Pd00nG_Boz22zQ_f66M 

key tersebut digunakan untuk membuka file zip yang berisi file txt dokumentoktok yang telah di decode dengan base64,base32,base16,caesar,rot47 

diawal kalimat string tersebut di decode oleh base64,base32 dan base16 lalu terdapat plaintext berupa  "APAKAH CAESAR ITU EFEKTIF" , yang merupakan hint untuk membuka decode selanjutnya menggunakan caesar dan pada di kalimat akhir ada plaintext berupa "AGAR RAHASIA ROT PERUSAHAAN TOK TOK TERJAMIN 47 KUNCI" yang merupakan Hint flag tersebut menggunakan ROT 47

dan didapatkan flag tersebut
T88rUUZZ7s_M9gJu_T5oOKto55Ook

Format flag : Hology4{T88rUUZZ7s_M9gJu_T5oOKto55Ook}

