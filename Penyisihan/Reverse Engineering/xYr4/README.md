# xYr4 (100 Pts)

- [Deskripsi ](#deskripsi)
- [Konsep soal ](#konsep-soal)
- [Hint](#hint)
- [Proof of Concept](#proof-of-concept)

## Deskripsi
Can you find the flag ?

Format Flag : hology4{.*}

Author : **fgm**

## Konsep soal
Reverse engineering

## HINT
What do you like ? ghidra or ida 

## Proof of Concept
Didapatkan sebuah file elf64 bit

ketika file tersebut dijalankan dimintai masukkan  kunci yang digunakan untuk mendapatkan flag tersebut
lalu saya mereverse file tersebut menggunakan IDA Pro
saat file tersebut dibuka dengan IDA Pro terlihat bahwasannya inputan masukkan akan dibandingkan sebuah kunci yang jika benar akan memanggil fungsi  decode  
sebelum masuk fungsi decode saya mengecek string yang ada dalam Source Code tersebut dan di temukan sebuah kalimat "JPEKK8XX8QIR44XIW77GYVM" yang dicurigai sebagai flag tersebut yang sudah di decode oleh fungsi decode
untuk mendapatkan source code dengan Ida pro saya gunakan FN+F5 
dan pada decode tertera source tersebut
```
    int __usercall decode@<eax>(char *a1@<rdi>, __int64 a2@<rbp>)
    {
      char v2; // al@3
      __int64 v3; // rdx@4
      __int64 v4; // rcx@4
      __int64 v5; // rax@4
      signed int i; // [sp-3Ch] [bp-3Ch]@1
      __int64 v8; // [sp-38h] [bp-38h]@3
      __int64 v9; // [sp-10h] [bp-10h]@1
      __int64 v10; // [sp-8h] [bp-8h]@1

      __asm { rep nop edx }
      v10 = a2;
      v9 = *MK_FP(__FS__, 40LL);
      for ( i = 0; i <= 22; ++i )
      {
        v2 = sub_11E0((unsigned int)a1[i]);
        std::operator+<char,std::char_traits<char>,std::allocator<char>>(
              &v8,
          &cipher,
      (unsigned int)(char)(((char)(v2 - 65) - 4) % 26 + 65));
        sub_11A0(&cipher, &v8);
        sub_1110(&v8);
      }
      sub_1170(&std::cout, &unk_2005);
      sub_1160(&std::cout, &cipher);
      v5 = *MK_FP(__FS__, 40LL) ^ v9;
      if ( *MK_FP(__FS__, 40LL) != v9 )
            LODWORD(v5) = sub_1190(&std::cout, &cipher, v3, v4);
      return v5;
    }
```

Inti dari Program tersebut adalah kalimat "JPEKK8XX8QIR44XIW77GYVM" akan bergeser dan membentuk sebuah string baru yang akan menjadi sebuah flagnya, namun karena terlalu kompleks saya cek bagian fungsi benar yang digunakan untuk membandingkan inputan dan key yang dimana jika benar akan mendapatkan flagnya langsung. 
Source Code yang di peroleh 
```
    int __usercall benar@<eax>(__int64 a1@<rbp>, int a2@<esi>)

    {

    int v2; // ST18_4@3
    int result; // eax@5
      signed int v4; // [sp-14h] [bp-14h]@1
      int i; // [sp-10h] [bp-10h]@1
      __int64 v6; // [sp-8h] [bp-8h]@1

      __asm { rep nop edx }
      v6 = a1;
      v4 = 98723;
      for ( i = 0; i < 5; i = v2 + 1 )
      {
        v2 = i + 30;
        v4 += v2;
      }
      if ( v4 == a2 )
      {
        sub_1170(&std::cout, "TRUEE");
        result = decode(fl, (__int64)&v6);
          }
      else
      {
        result = sub_1170(&std::cout, "False");
      }
      return result;
    }
```
dari source code tersebut kita peroleh bahwasannya Inputan akan di bandingkan dengan V4 dimana v4 memiliki nilai integer 98723 dan akan di loop satu kali di tambah dengan 30 , sehingga bisa kita peroleh key tersebut adalah 98753

    Lalu jalankan program tersebut dengan key 98753 dan akan mendapatkan sebuah flagnya yaitu FLAGG4TT4MEN00TES33CUR

<details>
	<summary>Flag : </summary>
	hology4{FLAGG4TT4MEN00TES33CUR}
</details>

## Referensi
