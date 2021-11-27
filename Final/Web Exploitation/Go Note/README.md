## Go-Note (500 Pts)

### Desc
kano membuat sebuah project OSS karena tertarik ikut hacktoberfest,kano membuat sebuah web simple yang berfungsi untuk menyimpan note saver yang berbasis cloud

tetapi karena mau showoff kano langsung mempublish OSS buatannya ke public padahal kano masih belum aware terhadap secure code dan security awareness,disini bisakah kamu membantu kano untuk melakukan audit kepada OSS yang dibuat oleh kano?

Author : **Just_Humanz**

### Proof of Concept
jadi ada simple webapp yang ada vuln idor,nah didalam note saver tersebut ada file `CiCd` milik dev yang dibuat untuk deploy OSS langkap dengan private key untuk ssh ke servernya,nah didalam servernya udah ada cronjob buat ngebackup file hology ke minIO atau aws3 bucket sehabis itu harus ngedump semua filenya abis itu cek satu satu filenya kalau nanti ketemu row `bendera` tapi flagnya masih terenkripsi,untuk dekripsinya butuh dapetin KEY yang,KEY ada di env dan ada sedikit trik jadi butuh baca soure codenya

jadi secara garis besar bakal kayak begini  
IDOR > get CiCD file > ssh > cek cronjob > dump sqlite file > get env > baca code > decript flag 

<details>
    <summary>Flag : </summary>
    Hology4{k3ep_d1gg1ng}
</details>

### Requirement
- lxc 
- nginx
- npm+node
- golang
- minIo client
- nginx (optional)
- screen/tmux
- ssh server

### Deploy
Deploy ubuntu di lxc dulu
```
lxc launch ubuntu 20.04 go-note
lxc console go-note
```
Install Requirement
```
apt update && apt install bla bla bla bla
```
ganti ssh port dan enable ssh server
```
nano /etc/ssh/sshd_config
systemctl start sshd
```
buat user dev dengan no root terus masukan key di `ssh-key` ke dalam dev biar bisa ssh dan jangan lupa disable bash historynya,sehabis itu copy source codenya **untuk sqlitenya cukup bawa yang hology yang hology.xx jangan** 

lalu build api dan go-notenya

###### api
jangan lupa update ip address di cicd filenya
```
cd ~/Go-Note/api
go build -o note
su root
tmux 
source .env
./note &>> note.log
```
###### go-note
cara susah: https://thucnc.medium.com/deploy-a-vuejs-web-app-with-nginx-on-ubuntu-18-04-f93860219030

cara gampang: `npm i && HOST='0.0.0.0' npm run dev`

setelah semua ok tinggal bind biar app yang di dalam lxc bisa ke expose oleh public **untuk api portnya harus tetep sama yaitu 2525 sementara go-note bisa diubah semaunya,ssh nya juga sama**
untuk referensi bisa dilihat disini https://discuss.linuxcontainers.org/t/forward-port-80-and-443-from-wan-to-container/2042  

lalu buat bucket di asw3-bucket atau minio terus upload semua sqlite file ke bucket,sehabis itu config di chall biar sqlitenya auto backup ke bucket  
