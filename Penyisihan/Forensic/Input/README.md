# input (500 Pts)

## Deskripsi
I get unknown file from my friend. He told me that he get that file from keyboard that use Linux FrameBuffer and US keyboard layout. Can you help me find the content of that file?

Format flag : hology4{spasi_diganti_underscore} (case insensitive)

Author : **RH01**

## Hint 

## Proof of Concept
Penjelasan teori dapat dilihat pada bagian (referensi)[#referensi]

Berikut script yang digunakan untuk membuat file
```
import struct
import signal
import atexit

f = open('/dev/input/by-path/platform-i8042-serio-0-event-kbd','rb')

o = open('/home/void/poros/hology4/Foren2/Chall','wb')

data = b''

def handle_exit():
        o.write(data)

atexit.register(handle_exit)
signal.signal(signal.SIGTERM, handle_exit)
signal.signal(signal.SIGINT, handle_exit)

while True :
        data += f.read(24)
        print(data)
```

Berikut script yang digunakan untuk mendapatkan key
```
import struct
# \s[0-9]+\s# 
keys = {
1:'Esc',
59:'F1',
60:'F2',
61:'F3',
62:'F4',
63:'F5',
64:'F6',
65:'F7',
66:'F8',
67:'F9',
68:'F10',
87:'F11',
88:'F12',
99:'PrintScrn',
70:'Scroll Lock',
119:'Pause',
41:'`',
2:'1',
3:'2',
4:'3',
5:'4',
6:'5',
7:'6',
8:'7',
9:'8',
10:'9',
11:'0',
12:'-',
13:'=',
14:'Backspace',
110:'Insert',
102:'Home',
104:'Page Up',
69:'Num Lock',
98:'KP /',
55:'KP *',
74:'KP -',
15:'Tab',
16:'Q',
17:'W',
18:'E',
19:'R',
20:'T',
21:'Y',
22:'U',
23:'I',
24:'O',
25:'P',
26:'[',
27:']',
28:'Return',
111:'Delete',
107:'End',
109:'Page Down',
71:'KP 7',
72:'KP 8',
73:'KP 9',
78:'KP +',
58:'Caps Lock',
30:'A',
31:'S',
32:'D',
33:'F',
34:'G',
35:'H',
36:'J',
37:'K',
38:'L',
39:';',
40:'\'',
75:'KP 4',
76:'KP 5',
77:'KP 6',
42:'Shift Left',
86:'International',
44:'Z',
45:'X',
46:'C',
47:'V',
48:'B',
49:'N',
50:'M',
51:',',
52:'.',
53:'/',
54:'Shift Right',
43:'\\',
103:'Cursor Up',
79:'KP 1',
80:'KP 2',
81:'KP 3',
96:'KP Enter',
29:'Ctrl Left',
125:'Logo Left (-> Option)',
56:'Alt Left (-> Command)',
57:'Space',
100:'Alt Right (-> Command)',
126:'Logo Right (-> Option)',
97:'Ctrl Right',
105:'Cursor Left',
108:'Cursor Down',
106:'Cursor Right',
82:'KP 0',
83:'KP .'
}

unpacked = []

with open('./Chall','rb') as f:
        while True :
                temp = f.read(24)
                if temp == b'':
                        break
                unpacked.append(str(struct.unpack('4IHHI',temp)))
                #print(struct.unpack('4IHHI',temp))

#print(unpacked)

def print_result(data):
        data=data[1:-1].split(', ')
        pressed = 'Ditekan' if (data[6] == '1') else 'Dilepas'
        print(f'Waktu : {data[0]}.{data[2]} , Key : {keys[int(data[5])]} , {pressed}')

for i in range(1,len(unpacked),3):
        print_result(unpacked[i])
        
'''
Output :
Waktu : 1634740738.254615 , Key : 0 , Ditekan
Waktu : 1634740738.332402 , Key : 0 , Dilepas
Waktu : 1634740739.602444 , Key : L , Ditekan
Waktu : 1634740739.688819 , Key : L , Dilepas
Waktu : 1634740740.54187 , Key : D , Ditekan
Waktu : 1634740740.140980 , Key : D , Dilepas
Waktu : 1634740741.621134 , Key : Space , Ditekan
Waktu : 1634740741.668217 , Key : Space , Dilepas
Waktu : 1634740742.797068 , Key : K , Ditekan
Waktu : 1634740742.879013 , Key : K , Dilepas
Waktu : 1634740743.279204 , Key : E , Ditekan
Waktu : 1634740743.361290 , Key : E , Dilepas
Waktu : 1634740743.670422 , Key : Y , Ditekan
Waktu : 1634740743.758114 , Key : Y , Dilepas
Waktu : 1634740744.519340 , Key : B , Ditekan
Waktu : 1634740744.591747 , Key : B , Dilepas
Waktu : 1634740744.911446 , Key : 0 , Ditekan
Waktu : 1634740744.978663 , Key : 0 , Dilepas
Waktu : 1634740747.147222 , Key : 4 , Ditekan
Waktu : 1634740747.199598 , Key : 4 , Dilepas
Waktu : 1634740747.509832 , Key : R , Ditekan
Waktu : 1634740747.576336 , Key : R , Dilepas
Waktu : 1634740747.791446 , Key : D , Ditekan
Waktu : 1634740747.882814 , Key : D , Dilepas
Waktu : 1634740748.533595 , Key : Space , Ditekan
Waktu : 1634740748.631352 , Key : Space , Dilepas
Waktu : 1634740749.172015 , Key : K , Ditekan
Waktu : 1634740749.249368 , Key : K , Dilepas
Waktu : 1634740749.633652 , Key : 3 , Ditekan
Waktu : 1634740749.711167 , Key : 3 , Dilepas
Waktu : 1634740750.744096 , Key : Y , Ditekan
Waktu : 1634740750.801714 , Key : Y , Dilepas
Waktu : 1634740751.989958 , Key : L , Ditekan
Waktu : 1634740752.67764 , Key : L , Dilepas
Waktu : 1634740752.317644 , Key : 0 , Ditekan
Waktu : 1634740752.375591 , Key : 0 , Dilepas
Waktu : 1634740752.949421 , Key : G , Ditekan
Waktu : 1634740753.32314 , Key : G , Dilepas
Waktu : 1634740753.121641 , Key : G , Ditekan
Waktu : 1634740753.193070 , Key : G , Dilepas
Waktu : 1634740754.54697 , Key : E , Ditekan
Waktu : 1634740754.142234 , Key : E , Dilepas
Waktu : 1634740754.325990 , Key : R , Ditekan
Waktu : 1634740754.413861 , Key : R , Dilepas
Waktu : 1634740755.280783 , Key : Ctrl Left , Ditekan
Waktu : 1634740755.411380 , Key : C , Ditekan
'''
```

<details>
<summary>Flag : </summary>
	hology4{0ld_keyb04rd_k3yl0gger} (case insensitive)
</details>

## Referensi
- https://thehackerdiary.wordpress.com/2017/04/21/exploring-devinput-1/
- https://gist.github.com/rickyzhang82/8581a762c9f9fc6ddb8390872552c250
