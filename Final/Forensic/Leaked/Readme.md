# Leaked (500 Pts)

## Deskripsi
Seorang newbie forensic lupa email dan password akunnya di sebuah website rahasia . Di website tersebut ia mendownload sebuah file yang terkunci. Ketika sedang mencoba mencari kunci file tersebut , tiba2 ia mendapatkan tugas lain yang harus diselesaikan . Bantu dia membuka file tersebut & menemukan email dan passwordnya yang hilang. Sebagai tambahan cari tau tools apa yang sedang ia pelajari

Format Flag : `hology4{email:password_toolsyangdipelajari_namafileyangdidownload_isifile}`
```
Contoh : 
email : hello@gmail.com  
password : inipassword  
tools yang dipelajari : nmap (lowercase)  
file yang di download : flag.txt  
isi file : 5eb63bbbe01eeed093cb22bb8f5acdc3  
flag : hology4{hello@gmail.com:inipassword_nmap_flag.txt_5eb63bbbe01eeed093cb22bb8f5acdc3}
```

Notes : 
- Password file yang didownload bukanlah password user
- Password file tidak perlu di bruteforce
- Password file berbeda dengan password email

Author : **RH01**
https://drive.google.com/file/d/17X_EBqHkW8zuXUGrMt6Gu8zP4bt5I9-t/view?usp=sharing

## Hint 
- "Jangan abaikan proses"
- mungkin yang kamu cari tidak sesulit/sejauh yang kamu pikirkan

## Proof of Concept
Build profile linux dengan menambahkan leaked-89.zip ke dalam "volatility/volatility/plugins/overlay/linux/"

Analisa menggunakan volatility . Lakukan analisa proses dengan `linux_psaux`

```bash
$ python2 /opt/ctf-tools/volatility/vol.py -f memory.dmp --profile=Linuxleaked-89x64 linux_psaux
<snip>
2606   1000   1000   bash                                                            
2621   1000   1000   python3                                                         
2698   0      0      [loop12]                                                        
2737   1000   1000   bash                                                            
2754   0      0      /usr/lib/snapd/snapd                                            
2797   0      0      [kworker/u2:1]                                                  
2880   0      0      sudo ./avml memory.dmp                                          
2881   0      0      ./avml memory.dmp
```

terlihat program bash dan program python3 sedang berjalan

```bash
# Enumerasi File Linux disimpan di data/files
$ mkdir data
$ python2 /opt/ctf-tools/volatility/vol.py -f memory.dmp --profile=Linuxleaked-89x64 linux_enumerate_files > data/files.txt
```

Analisa folder home , diketahui usernamenya adalah `void` dan jika dianalisa lebih lanjut terdapat folder `.mozilla` yang berarti browser yang digunakan adalah mozilla. 

Selanjutnya cek ada tidaknya file `logins.json`. Jika ada maka ada data akun yang disimpan. Untuk mengekstrak data dari login.json harus dipasangkan dengan file `key4.db`. 

```bash
# alamat inode didapat dari data/files hasil enumerasi sebelumnya
$ python2 /opt/ctf-tools/volatility/vol.py -f memory.dmp --profile=Linuxleaked-89x64 linux_find_file -i 0xffff9e1329ef9208 -O key4.db
$ python2 /opt/ctf-tools/volatility/vol.py -f memory.dmp --profile=Linuxleaked-89x64 linux_find_file -i 0xffff9e12e193dad0 -O logins.json
```

Tukar sementara logins.json dan key4.db pada mozilla anda (`/home/<user>/.mozilla/firefox/<default>/`) dan buka informasi login mozilla pada `about:logins`. Disana akan terdapat email dan password yaitu 
```
email : bob@forensic.hology.com
password : ohnomypasswordisleaked
```

Selanjutnya dump `places.sqlite` untuk menganalisa lokasi yang pernah dikunjungi , namun tidak ada informasi penting disana. Dump `formhistory.sqlite` untuk mendapatkan informasi history dari form termasuk searchbar-history.

```bash
$ python2 /opt/ctf-tools/volatility/vol.py -f memory.dmp --profile=Linuxleaked-89x64 linux_find_file -i 0xffff9e12ef9e3000 -O places.sqlite
$ python2 /opt/ctf-tools/volatility/vol.py -f memory.dmp --profile=Linuxleaked-89x64 linux_find_file -i 0xffff9e12e180abb8 -O formhistory.sqlite
```

Untuk menemukan file yang didownload dapat dilihat pada folder `Downloads` disana terdapat file `secret.zip`. Nama file yang didownload adalah `secret.zip`. Lalu berdasarkan informasi dari proses yang berjalan , maka cari pid proses python dan analisa proses python menggunakan `linux_proc_maps`.

```
$ python2 /opt/ctf-tools/volatility/vol.py -f memory.dmp --profile=Linuxleaked-89x64 linux_proc_maps -p 2621
Offset             Pid      Name                 Start              End                Flags               Pgoff Major  Minor  Inode      File Path
------------------ -------- -------------------- ------------------ ------------------ ------ ------------------ ------ ------ ---------- ---------
0xffff9e12dce945c0     2621 python3              0x0000000000400000 0x00000000007b4000 r-x                   0x0      8      1     656495 /usr/bin/python3.6
0xffff9e12dce945c0     2621 python3              0x00000000009b3000 0x00000000009b4000 r--              0x3b3000      8      1     656495 /usr/bin/python3.6
0xffff9e12dce945c0     2621 python3              0x00000000009b4000 0x0000000000a51000 rw-              0x3b4000      8      1     656495 /usr/bin/python3.6
0xffff9e12dce945c0     2621 python3              0x0000000000a51000 0x0000000000a84000 rw-                   0x0      0      0          0 
0xffff9e12dce945c0     2621 python3              0x00000000016fa000 0x00000000017f3000 rw-                   0x0      0      0          0 [heap]
<snippet>
```

Kode yang dimasukkan disimpan pada heap. Dump pada bagian heap untuk mendapatkan kodenya menggunakan `linux_dump_map`
```bash
$ python2 /opt/ctf-tools/volatility/vol.py -f memory.dmp --profile=Linuxleaked-89x64 linux_dump_map -p 2621 -s 0x00000000016fa000 -D .
$ strings task.2477.0x21ab000.vma
<snippet>
>> x=''.join([chr(int(i,16)) for i in '79 30 75 2d 67 30 74 2d 74 68 33 2d 6b 33 79'.split(' ')])
>>>
```

Didapatkan x bernilai `y0u-g0t-th3-k3y`. Dump `secret.zip` lalu buka menggunakan kunci tersebut . Didapatkan isinya adalah `7eefda6cdd434e0f9af420a92f2dc7cb`


<details>
<summary>Flag : </summary>
hology4{bob@forensic.hology.com:ohnomypasswordisleaked_exiftool_secret.zip_7eefda6cdd434e0f9af420a92f2dc7cb}
</details>

## Referensi
- https://github.com/volatilityfoundation/volatility/wiki/Linux-Command-Reference
- https://docs.python.org/3.8/c-api/memory.html
- https://ctftime.org/writeup/5125
- https://support.mozilla.org/en-US/questions/1278039 
