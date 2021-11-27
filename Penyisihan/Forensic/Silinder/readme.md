# silinder

## Deskripsi
ada hal penting di dalam file ini....

Author : Genzo

## Hint
- 

## Proof of Concept
BITMAP merupakan sebuah file yang memiliki cara kerja yan unik. pada file aslinya yang terdapat flag, memiliki ukuran 10 x 7 (dalam hex) dan memiliki "BITMAPINFOHEADER" yang sesuai. namun di dalam file yang diberikan ke peserta
merupakan file yang telah dimanipulasi "BITMAPINFOHEADER"-nya dan di ubah ukurannya menjadi 5 x 2. 
1. Manipulasi BITMAPINFOHEADER : langkah pertama yang harus di lakukan peserta adalah mengembalikan BITMAPINFOHEADER-nya ke semula, sehingga file tersebut bisa dibuka dengan ekstensi ".bmp". Sesuai dengan teori yang ada 
size dari BITMAPINFOHEADER tersebut harus memiliki size 40 (dalam ascii) dan 28 (dalam hex). file di soal sudah dimanipulasi sehingga size dari BITMAPINFOHEADER-nya tidak sesuai dengan size yang seharunya yaitu 40.
2. Mengambalikan ukuran : setelah file tersebut berhasil dikembalikan ke ekstensinya yaitu ".bmp", maka langkah selanjutnya, peserta diharuskan untuk mengembalikannya ke ukuran biasa.
	Sebagai contoh :
	Ukuran asli file tersebut adalah 10x7(dalam hex).
	Ukuran file yang diberikan adalah 5x3.
		Maka secara teori, misalkan gambaran hex dari file aslinya adalah seperti ini : 0A07ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQR
			maka jika direpresentasikan dalam file bitmap akan menjadi seperti ini :
				ABCDEFGHIJ
				KLMNOPQRST
				UVWXYZABCD
				EFGHIJKLMN
				OPQRSTUVWX
				YZABCDEFGH
				IJKLMNOPQR   

		Namun file yang diberikan ke peserta, maka representasinya akan menjadi seperti ini :
				(YD)(ZE)(AF)(BG)(CH)
				(IN)(JO)(KP)(LQ)(MR)
			sehingga gambarnya terlihat bertindihan.
	
	Peserta harus mengembalikan file tesebut ke ukuran semulanya karena flagnya terdapat di file dengan ukuran yang sesungguhnya maka akan di dapatkan sebuah flag
				
				

<details>
	<summary>Flag : </summary>
	hology4{Haloo_bos_bos_q_semuanyaa}
</details>

## Referensi
- https://rasyidmf.com/Challenge?category=Cryptography
- https://www.file-recovery.com/bmp-signature-format.htm
