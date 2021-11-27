# Customization (349 Pts)

## Deskripsi
I made a customization on a very popular algorithm among CTF players . Can you get my flag?

Flag format : hology4{string}

Author : **RH01**

## Hint 

## Proof of Concept
Program yang dibuat berdasarkan algoritma encoding base64. bit oktet (8 bit) dari setiap karakter diubah menjadi sixtet (6 bit) dan diubah menjadi index untuk tabel yang berisi 64 karakter. Custom yang dilakukan pada isi tabel dan dilakukan xor terhadap nilai 0x5e setelah diencode 10x. 

Untuk mengembalikan nilainya maka decode xor nilai flag dengan nilai 0x5e lalu decode 10x.

Untuk melakukan reverse dapat menggunakan CFR pada http://www.javadecompilers.com
```
/*
 * Decompiled with CFR 0.150.
 */
import java.util.Scanner;

public class RE {
    static char[] x = new char[]{'\"', 'd', 's', 'V', 'U', '8', 'q', ';', '|', 'G', 'k', 'M', '9', 'K', '<', '2', 'C', 'a', 'c', ')', '5', 'v', 'D', ']', '7', '\'', '>', '$', 'y', 'E', 'l', 'W', 'A', 'F', '%', 'i', 'O', '&', '*', '[', 'Y', 'H', 'L', 'u', 'g', '(', 'B', ':', '?', '1', '!', '-', '0', 'w', 'J', 'R', 'm', 'f', 'N', '_', '{', 'b', 'p', '4'};

    public static void main(String[] arrstring) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Masukkan Flag : ");
        char[] arrc = scanner.nextLine().toCharArray();
        scanner.close();
        String string = RE.enc(arrc, 10);
        Integer[] arrinteger = new Integer[string.length()];
        for (int i = 0; i < string.length(); ++i) {
            arrinteger[i] = string.charAt(i) ^ 0x5E;
        }
        if (arrinteger[0] == 61 && arrinteger[1] == 41 && arrinteger[2] == 25 && arrinteger[3] == 102 && arrinteger[4] == 98 && arrinteger[5] == 115 && arrinteger[6] == 111 && arrinteger[7] == 26 && arrinteger[8] == 108 && arrinteger[9] == 127 && arrinteger[10] == 31 && arrinteger[11] == 56 && arrinteger[12] == 21 && arrinteger[13] == 107 && arrinteger[14] == 103 && arrinteger[15] == 24 && arrinteger[16] == 98 && arrinteger[17] == 3 && arrinteger[18] == 21 && arrinteger[19] == 102 && arrinteger[20] == 39 && arrinteger[21] == 119 && arrinteger[22] == 56 && arrinteger[23] == 12 && arrinteger[24] == 103 && arrinteger[25] == 8 && arrinteger[26] == 11 && arrinteger[27] == 12 && arrinteger[28] == 61 && arrinteger[29] == 55 && arrinteger[30] == 24 && arrinteger[31] == 40 && arrinteger[32] == 63 && arrinteger[33] == 110 && arrinteger[34] == 60 && arrinteger[35] == 34 && arrinteger[36] == 63 && arrinteger[37] == 26 && arrinteger[38] == 21 && arrinteger[39] == 22 && arrinteger[40] == 25 && arrinteger[41] == 5 && arrinteger[42] == 102 && arrinteger[43] == 24 && arrinteger[44] == 34 && arrinteger[45] == 3 && arrinteger[46] == 121 && arrinteger[47] == 53 && arrinteger[48] == 108 && arrinteger[49] == 102 && arrinteger[50] == 34 && arrinteger[51] == 5 && arrinteger[52] == 40 && arrinteger[53] == 26 && arrinteger[54] == 102 && arrinteger[55] == 20 && arrinteger[56] == 63 && arrinteger[57] == 110 && arrinteger[58] == 110 && arrinteger[59] == 56 && arrinteger[60] == 39 && arrinteger[61] == 107 && arrinteger[62] == 27 && arrinteger[63] == 1 && arrinteger[64] == 103 && arrinteger[65] == 116 && arrinteger[66] == 120 && arrinteger[67] == 20 && arrinteger[68] == 107 && arrinteger[69] == 116 && arrinteger[70] == 11 && arrinteger[71] == 41 && arrinteger[72] == 63 && arrinteger[73] == 61 && arrinteger[74] == 29 && arrinteger[75] == 127 && arrinteger[76] == 107 && arrinteger[77] == 116 && arrinteger[78] == 29 && arrinteger[79] == 51 && arrinteger[80] == 9 && arrinteger[81] == 8 && arrinteger[82] == 107 && arrinteger[83] == 41 && arrinteger[84] == 103 && arrinteger[85] == 26 && arrinteger[86] == 102 && arrinteger[87] == 55 && arrinteger[88] == 21 && arrinteger[89] == 115 && arrinteger[90] == 60 && arrinteger[91] == 19 && arrinteger[92] == 40 && arrinteger[93] == 119 && arrinteger[94] == 107 && arrinteger[95] == 7 && arrinteger[96] == 122 && arrinteger[97] == 47 && arrinteger[98] == 21 && arrinteger[99] == 34 && arrinteger[100] == 63 && arrinteger[101] == 55 && arrinteger[102] == 17 && arrinteger[103] == 1 && arrinteger[104] == 63 && arrinteger[105] == 107 && arrinteger[106] == 110 && arrinteger[107] == 46 && arrinteger[108] == 27 && arrinteger[109] == 115 && arrinteger[110] == 58 && arrinteger[111] == 122 && arrinteger[112] == 108 && arrinteger[113] == 55 && arrinteger[114] == 40 && arrinteger[115] == 102 && arrinteger[116] == 63 && arrinteger[117] == 47 && arrinteger[118] == 102 && arrinteger[119] == 122 && arrinteger[120] == 53 && arrinteger[121] == 102 && arrinteger[122] == 107 && arrinteger[123] == 56 && arrinteger[124] == 40 && arrinteger[125] == 119 && arrinteger[126] == 107 && arrinteger[127] == 7 && arrinteger[128] == 96 && arrinteger[129] == 20 && arrinteger[130] == 103 && arrinteger[131] == 51 && arrinteger[132] == 39 && arrinteger[133] == 3 && arrinteger[134] == 121 && arrinteger[135] == 22 && arrinteger[136] == 119 && arrinteger[137] == 41 && arrinteger[138] == 60 && arrinteger[139] == 43 && arrinteger[140] == 98 && arrinteger[141] == 47 && arrinteger[142] == 25 && arrinteger[143] == 21 && arrinteger[144] == 61 && arrinteger[145] == 110 && arrinteger[146] == 63 && arrinteger[147] == 37 && arrinteger[148] == 25 && arrinteger[149] == 55 && arrinteger[150] == 25 && arrinteger[151] == 115 && arrinteger[152] == 119 && arrinteger[153] == 127 && arrinteger[154] == 7 && arrinteger[155] == 56 && arrinteger[156] == 96 && arrinteger[157] == 3 && arrinteger[158] == 17 && arrinteger[159] == 18 && arrinteger[160] == 63 && arrinteger[161] == 12 && arrinteger[162] == 27 && arrinteger[163] == 20 && arrinteger[164] == 63 && arrinteger[165] == 5 && arrinteger[166] == 121 && arrinteger[167] == 55 && arrinteger[168] == 63 && arrinteger[169] == 5 && arrinteger[170] == 105 && arrinteger[171] == 127 && arrinteger[172] == 39 && arrinteger[173] == 119 && arrinteger[174] == 17 && arrinteger[175] == 18 && arrinteger[176] == 63 && arrinteger[177] == 107 && arrinteger[178] == 60 && arrinteger[179] == 8 && arrinteger[180] == 108 && arrinteger[181] == 110 && arrinteger[182] == 57 && arrinteger[183] == 22 && arrinteger[184] == 105 && arrinteger[185] == 17 && arrinteger[186] == 121 && arrinteger[187] == 43 && arrinteger[188] == 21 && arrinteger[189] == 61 && arrinteger[190] == 39 && arrinteger[191] == 51 && arrinteger[192] == 21 && arrinteger[193] == 12 && arrinteger[194] == 105 && arrinteger[195] == 12 && arrinteger[196] == 53 && arrinteger[197] == 17 && arrinteger[198] == 27 && arrinteger[199] == 26 && arrinteger[200] == 9 && arrinteger[201] == 45 && arrinteger[202] == 7 && arrinteger[203] == 56 && arrinteger[204] == 98 && arrinteger[205] == 11 && arrinteger[206] == 27 && arrinteger[207] == 41 && arrinteger[208] == 27 && arrinteger[209] == 123 && arrinteger[210] == 17 && arrinteger[211] == 116 && arrinteger[212] == 9 && arrinteger[213] == 45 && arrinteger[214] == 63 && arrinteger[215] == 40 && arrinteger[216] == 26 && arrinteger[217] == 40 && arrinteger[218] == 118 && arrinteger[219] == 24 && arrinteger[220] == 63 && arrinteger[221] == 47 && arrinteger[222] == 27 && arrinteger[223] == 12 && arrinteger[224] == 63 && arrinteger[225] == 40 && arrinteger[226] == 118 && arrinteger[227] == 5 && arrinteger[228] == 34 && arrinteger[229] == 3 && arrinteger[230] == 17 && arrinteger[231] == 41 && arrinteger[232] == 121 && arrinteger[233] == 127 && arrinteger[234] == 102 && arrinteger[235] == 101 && arrinteger[236] == 40 && arrinteger[237] == 5 && arrinteger[238] == 121 && arrinteger[239] == 53 && arrinteger[240] == 105 && arrinteger[241] == 119 && arrinteger[242] == 40 && arrinteger[243] == 56 && arrinteger[244] == 108 && arrinteger[245] == 110 && arrinteger[246] == 40 && arrinteger[247] == 108 && arrinteger[248] == 25 && arrinteger[249] == 115 && arrinteger[250] == 24 && arrinteger[251] == 43 && arrinteger[252] == 119 && arrinteger[253] == 12 && arrinteger[254] == 120 && arrinteger[255] == 118 && arrinteger[256] == 61 && arrinteger[257] == 41 && arrinteger[258] == 34 && arrinteger[259] == 5 && arrinteger[260] == 61 && arrinteger[261] == 17 && arrinteger[262] == 29 && arrinteger[263] == 24 && arrinteger[264] == 105 && arrinteger[265] == 40 && arrinteger[266] == 107 && arrinteger[267] == 37 && arrinteger[268] == 39 && arrinteger[269] == 61 && arrinteger[270] == 105 && arrinteger[271] == 1 && arrinteger[272] == 63 && arrinteger[273] == 12 && arrinteger[274] == 105 && arrinteger[275] == 116 && arrinteger[276] == 108 && arrinteger[277] == 115 && arrinteger[278] == 27 && arrinteger[279] == 53 && arrinteger[280] == 25 && arrinteger[281] == 55 && arrinteger[282] == 25 && arrinteger[283] == 11 && arrinteger[284] == 53 && arrinteger[285] == 3 && arrinteger[286] == 120 && arrinteger[287] == 34 && arrinteger[288] == 25 && arrinteger[289] == 115 && arrinteger[290] == 24 && arrinteger[291] == 101 && arrinteger[292] == 96 && arrinteger[293] == 61 && arrinteger[294] == 17 && arrinteger[295] == 97 && arrinteger[296] == 98 && arrinteger[297] == 40 && arrinteger[298] == 17 && arrinteger[299] == 127 && arrinteger[300] == 98 && arrinteger[301] == 8 && arrinteger[302] == 27 && arrinteger[303] == 118 && arrinteger[304] == 108 && arrinteger[305] == 45 && arrinteger[306] == 110 && arrinteger[307] == 111 && arrinteger[308] == 63 && arrinteger[309] == 47 && arrinteger[310] == 102 && arrinteger[311] == 20 && arrinteger[312] == 63 && arrinteger[313] == 110 && arrinteger[314] == 121 && arrinteger[315] == 56 && arrinteger[316] == 107 && arrinteger[317] == 5 && arrinteger[318] == 120 && arrinteger[319] == 103 && arrinteger[320] == 21 && arrinteger[321] == 127 && arrinteger[322] == 120 && arrinteger[323] == 20 && arrinteger[324] == 121 && arrinteger[325] == 123 && arrinteger[326] == 120 && arrinteger[327] == 40 && arrinteger[328] == 63 && arrinteger[329] == 107 && arrinteger[330] == 31 && arrinteger[331] == 5 && arrinteger[332] == 61 && arrinteger[333] == 123 && arrinteger[334] == 27 && arrinteger[335] == 53 && arrinteger[336] == 21 && arrinteger[337] == 115 && arrinteger[338] == 124 && arrinteger[339] == 41 && arrinteger[340] == 103 && arrinteger[341] == 26 && arrinteger[342] == 103 && arrinteger[343] == 24 && arrinteger[344] == 27 && arrinteger[345] == 110 && arrinteger[346] == 60 && arrinteger[347] == 101 && arrinteger[348] == 21 && arrinteger[349] == 107 && arrinteger[350] == 21 && arrinteger[351] == 47 && arrinteger[352] == 29 && arrinteger[353] == 115 && arrinteger[354] == 58 && arrinteger[355] == 58 && arrinteger[356] == 34 && arrinteger[357] == 119 && arrinteger[358] == 40 && arrinteger[359] == 108 && arrinteger[360] == 27 && arrinteger[361] == 17 && arrinteger[362] == 29 && arrinteger[363] == 12 && arrinteger[364] == 61 && arrinteger[365] == 123 && arrinteger[366] == 121 && arrinteger[367] == 108 && arrinteger[368] == 103 && arrinteger[369] == 17 && arrinteger[370] == 7 && arrinteger[371] == 56 && arrinteger[372] == 53 && arrinteger[373] == 11 && arrinteger[374] == 40 && arrinteger[375] == 108 && arrinteger[376] == 26 && arrinteger[377] == 3 && arrinteger[378] == 58 && arrinteger[379] == 19 && arrinteger[380] == 21 && arrinteger[381] == 3 && arrinteger[382] == 121 && arrinteger[383] == 55 && arrinteger[384] == 21 && arrinteger[385] == 20 && arrinteger[386] == 21 && arrinteger[387] == 12 && arrinteger[388] == 39 && arrinteger[389] == 119 && arrinteger[390] == 111 && arrinteger[391] == 40 && arrinteger[392] == 53 && arrinteger[393] == 11 && arrinteger[394] == 40 && arrinteger[395] == 101 && arrinteger[396] == 25 && arrinteger[397] == 17 && arrinteger[398] == 27 && arrinteger[399] == 47 && arrinteger[400] == 63 && arrinteger[401] == 40 && arrinteger[402] == 34 && arrinteger[403] == 111 && arrinteger[404] == 105 && arrinteger[405] == 41 && arrinteger[406] == 39 && arrinteger[407] == 51 && arrinteger[408] == 27 && arrinteger[409] == 115 && arrinteger[410] == 25 && arrinteger[411] == 24 && arrinteger[412] == 119 && arrinteger[413] == 110 && arrinteger[414] == 40 && arrinteger[415] == 43 && arrinteger[416] == 25 && arrinteger[417] == 12 && arrinteger[418] == 27 && arrinteger[419] == 37 && arrinteger[420] == 107 && arrinteger[421] == 17 && arrinteger[422] == 107 && arrinteger[423] == 97 && arrinteger[424] == 50 && arrinteger[425] == 115 && arrinteger[426] == 37 && arrinteger[427] == 37 && arrinteger[428] == 105 && arrinteger[429] == 12 && arrinteger[430] == 121 && arrinteger[431] == 61 && arrinteger[432] == 108 && arrinteger[433] == 8 && arrinteger[434] == 57 && arrinteger[435] == 5 && arrinteger[436] == 96 && arrinteger[437] == 26 && arrinteger[438] == 102 && arrinteger[439] == 20 && arrinteger[440] == 63 && arrinteger[441] == 110 && arrinteger[442] == 105 && arrinteger[443] == 56 && arrinteger[444] == 39 && arrinteger[445] == 110 && arrinteger[446] == 107 && arrinteger[447] == 18 && arrinteger[448] == 105 && arrinteger[449] == 127 && arrinteger[450] == 102 && arrinteger[451] == 108 && arrinteger[452] == 27 && arrinteger[453] == 115 && arrinteger[454] == 120 && arrinteger[455] == 40 && arrinteger[456] == 21 && arrinteger[457] == 61 && arrinteger[458] == 121 && arrinteger[459] == 102 && arrinteger[460] == 27 && arrinteger[461] == 20 && arrinteger[462] == 27 && arrinteger[463] == 12 && arrinteger[464] == 96 && arrinteger[465] == 127 && arrinteger[466] == 22 && arrinteger[467] == 102 && arrinteger[468] == 98 && arrinteger[469] == 119 && arrinteger[470] == 51 && arrinteger[471] == 118 && arrinteger[472] == 121 && arrinteger[473] == 41 && arrinteger[474] == 25 && arrinteger[475] == 11 && arrinteger[476] == 19 && arrinteger[477] == 61 && arrinteger[478] == 105 && arrinteger[479] == 118 && arrinteger[480] == 122 && arrinteger[481] == 11 && arrinteger[482] == 29 && arrinteger[483] == 97 && arrinteger[484] == 61 && arrinteger[485] == 47 && arrinteger[486] == 57 && arrinteger[487] == 22 && arrinteger[488] == 27 && arrinteger[489] == 123 && arrinteger[490] == 31 && arrinteger[491] == 56 && arrinteger[492] == 119 && arrinteger[493] == 41 && arrinteger[494] == 120 && arrinteger[495] == 116 && arrinteger[496] == 96 && arrinteger[497] == 41 && arrinteger[498] == 107 && arrinteger[499] == 7 && arrinteger[500] == 63 && arrinteger[501] == 26 && arrinteger[502] == 97 && arrinteger[503] == 41 && arrinteger[504] == 63 && arrinteger[505] == 17 && arrinteger[506] == 118 && arrinteger[507] == 27 && arrinteger[508] == 53 && arrinteger[509] == 40 && arrinteger[510] == 40 && arrinteger[511] == 121 && arrinteger[512] == 21 && arrinteger[513] == 12 && arrinteger[514] == 27 && arrinteger[515] == 101 && arrinteger[516] == 121 && arrinteger[517] == 116 && arrinteger[518] == 118 && arrinteger[519] == 22 && arrinteger[520] == 103 && arrinteger[521] == 3 && arrinteger[522] == 121 && arrinteger[523] == 102 && arrinteger[524] == 119 && arrinteger[525] == 110 && arrinteger[526] == 102 && arrinteger[527] == 122 && arrinteger[528] == 108 && arrinteger[529] == 45 && arrinteger[530] == 41 && arrinteger[531] == 123 && arrinteger[532] == 103 && arrinteger[533] == 116 && arrinteger[534] == 111 && arrinteger[535] == 55 && arrinteger[536] == 103 && arrinteger[537] == 45 && arrinteger[538] == 40 && arrinteger[539] == 101 && arrinteger[540] == 25 && arrinteger[541] == 17 && arrinteger[542] == 40 && arrinteger[543] == 11 && arrinteger[544] == 63 && arrinteger[545] == 107 && arrinteger[546] == 37 && arrinteger[547] == 12 && arrinteger[548] == 107 && arrinteger[549] == 55 && arrinteger[550] == 17 && arrinteger[551] == 1 && arrinteger[552] == 119 && arrinteger[553] == 127 && arrinteger[554] == 22 && arrinteger[555] == 19 && arrinteger[556] == 19 && arrinteger[557] == 119 && arrinteger[558] == 40 && arrinteger[559] == 122 && arrinteger[560] == 105 && arrinteger[561] == 40 && arrinteger[562] == 118 && arrinteger[563] == 108 && arrinteger[564] == 63 && arrinteger[565] == 116 && arrinteger[566] == 102 && arrinteger[567] == 27 && arrinteger[568] == 21 && arrinteger[569] == 61 && arrinteger[570] == 102 && arrinteger[571] == 56 && arrinteger[572] == 21 && arrinteger[573] == 26 && arrinteger[574] == 27 && arrinteger[575] == 122 && arrinteger[576] == 105 && arrinteger[577] == 40 && arrinteger[578] == 121 && arrinteger[579] == 20 && arrinteger[580] == 61 && arrinteger[581] == 55 && arrinteger[582] == 120 && arrinteger[583] == 108 && arrinteger[584] == 21 && arrinteger[585] == 61 && arrinteger[586] == 105 && arrinteger[587] == 12 && arrinteger[588] == 53 && arrinteger[589] == 107 && arrinteger[590] == 121 && arrinteger[591] == 41 && arrinteger[592] == 63 && arrinteger[593] == 41 && arrinteger[594] == 110 && arrinteger[595] == 116 && arrinteger[596] == 53 && arrinteger[597] == 101 && arrinteger[598] == 121 && arrinteger[599] == 61 && arrinteger[600] == 98 && arrinteger[601] == 11 && arrinteger[602] == 110 && arrinteger[603] == 37 && arrinteger[604] == 39 && arrinteger[605] == 119 && arrinteger[606] == 120 && arrinteger[607] == 118 && arrinteger[608] == 63 && arrinteger[609] == 119 && arrinteger[610] == 58 && arrinteger[611] == 20 && arrinteger[612] == 119 && arrinteger[613] == 107 && arrinteger[614] == 41 && arrinteger[615] == 11 && arrinteger[616] == 63 && arrinteger[617] == 12 && arrinteger[618] == 105 && arrinteger[619] == 12 && arrinteger[620] == 103 && arrinteger[621] == 47 && arrinteger[622] == 25 && arrinteger[623] == 122 && arrinteger[624] == 9 && arrinteger[625] == 45 && arrinteger[626] == 120 && arrinteger[627] == 102 && arrinteger[628] == 96 && arrinteger[629] == 61 && arrinteger[630] == 120 && arrinteger[631] == 53 && arrinteger[632] == 27 && arrinteger[633] == 20 && arrinteger[634] == 17 && arrinteger[635] == 22 && arrinteger[636] == 34 && arrinteger[637] == 107 && arrinteger[638] == 121 && arrinteger[639] == 34 && arrinteger[640] == 21 && arrinteger[641] == 127 && arrinteger[642] == 120 && arrinteger[643] == 20 && arrinteger[644] == 61 && arrinteger[645] == 120 && arrinteger[646] == 41 && arrinteger[647] == 27 && arrinteger[648] == 25 && arrinteger[649] == 110 && arrinteger[650] == 22 && arrinteger[651] == 19 && arrinteger[652] == 29 && arrinteger[653] == 97) {
            System.out.println("You Got the Flag, Congrats !!!");
        } else {
            System.out.println("Nope, Wrong Flag");
        }
    }

    public static String enc(char[] arrc, int n) {
        Object object = "";
        for (int i = 0; i < arrc.length; ++i) {
            object = (String)object + String.format("%08d", Integer.valueOf(Integer.toBinaryString(arrc[i])));
        }
        while (((String)object).length() % 6 != 0) {
            object = (String)object + "0";
        }
        Object object2 = "";
        for (int i = 0; i < ((String)object).length() / 6; ++i) {
            object2 = (String)object2 + x[Integer.parseInt(((String)object).substring(i * 6, i * 6 + 6), 2)];
        }
        if (n > 1) {
            return RE.enc(((String)object2).toCharArray(), n - 1);
        }
        return object2;
    }
}


```

Berikut script decryptornya

```
public class RE_Decryptor {
    
    static char x [] = {'"', 'd', 's', 'V', 'U', '8', 'q', ';', '|', 'G', 'k', 'M', '9', 'K', '<', '2', 'C', 'a', 'c', ')', '5', 'v', 'D', ']', '7', '\'', '>', '$', 'y', 'E', 'l', 'W', 'A', 'F', '%', 'i', 'O', '&', '*', '[', 'Y', 'H', 'L', 'u', 'g', '(', 'B', ':', '?', '1', '!', '-', '0', 'w', 'J', 'R', 'm', 'f', 'N', '_', '{', 'b', 'p', '4'};    

public static void main(String[] args) {        
        
        int encrypted [] = {61,41,25,102,98,115,111,26,108,127,31,56,21,107,103,24,98,3,21,102,39,119,56,12,103,8,11,12,61,55,24,40,63,110,60,34,63,26,21,22,25,5,102,24,34,3,121,53,108,102,34,5,40,26,102,20,63,110,110,56,39,107,27,1,103,116,120,20,107,116,11,41,63,61,29,127,107,116,29,51,9,8,107,41,103,26,102,55,21,115,60,19,40,119,107,7,122,47,21,34,63,55,17,1,63,107,110,46,27,115,58,122,108,55,40,102,63,47,102,122,53,102,107,56,40,119,107,7,96,20,103,51,39,3,121,22,119,41,60,43,98,47,25,21,61,110,63,37,25,55,25,115,119,127,7,56,96,3,17,18,63,12,27,20,63,5,121,55,63,5,105,127,39,119,17,18,63,107,60,8,108,110,57,22,105,17,121,43,21,61,39,51,21,12,105,12,53,17,27,26,9,45,7,56,98,11,27,41,27,123,17,116,9,45,63,40,26,40,118,24,63,47,27,12,63,40,118,5,34,3,17,41,121,127,102,101,40,5,121,53,105,119,40,56,108,110,40,108,25,115,24,43,119,12,120,118,61,41,34,5,61,17,29,24,105,40,107,37,39,61,105,1,63,12,105,116,108,115,27,53,25,55,25,11,53,3,120,34,25,115,24,101,96,61,17,97,98,40,17,127,98,8,27,118,108,45,110,111,63,47,102,20,63,110,121,56,107,5,120,103,21,127,120,20,121,123,120,40,63,107,31,5,61,123,27,53,21,115,124,41,103,26,103,24,27,110,60,101,21,107,21,47,29,115,58,58,34,119,40,108,27,17,29,12,61,123,121,108,103,17,7,56,53,11,40,108,26,3,58,19,21,3,121,55,21,20,21,12,39,119,111,40,53,11,40,101,25,17,27,47,63,40,34,111,105,41,39,51,27,115,25,24,119,110,40,43,25,12,27,37,107,17,107,97,50,115,37,37,105,12,121,61,108,8,57,5,96,26,102,20,63,110,105,56,39,110,107,18,105,127,102,108,27,115,120,40,21,61,121,102,27,20,27,12,96,127,22,102,98,119,51,118,121,41,25,11,19,61,105,118,122,11,29,97,61,47,57,22,27,123,31,56,119,41,120,116,96,41,107,7,63,26,97,41,63,17,118,27,53,40,40,121,21,12,27,101,121,116,118,22,103,3,121,102,119,110,102,122,108,45,41,123,103,116,111,55,103,45,40,101,25,17,40,11,63,107,37,12,107,55,17,1,119,127,22,19,19,119,40,122,105,40,118,108,63,116,102,27,21,61,102,56,21,26,27,122,105,40,121,20,61,55,120,108,21,61,105,12,53,107,121,41,63,41,110,116,53,101,121,61,98,11,110,37,39,119,120,118,63,119,58,20,119,107,41,11,63,12,105,12,103,47,25,122,9,45,120,102,96,61,120,53,27,20,17,22,34,107,121,34,21,127,120,20,61,120,41,27,25,110,22,19,29,97};        
        
        char out [] = new char[encrypted.length];

        for (int i = 0; i < encrypted.length; i++) {
            out[i] = (char) (encrypted[i] ^ 0x5e);
        }

        System.out.println(dec(out,10));        

    }

    public static String dec(char[] input,int n) {
        String combine = "";                        
                
        for (int i = 0; i < input.length; i++) {            
            combine += String.format("%06d", Integer.valueOf(Integer.toBinaryString(new String(x).indexOf(input[i]))));            
        }                                                 

        String output = "";
                                
        for (int j = 0; j < combine.length()/8; j++) {            

            output += (char)Integer.parseInt(combine.substring(j*8, j*8+8),2);
        }
        
        if(n > 1){
            return dec(output.toCharArray(),n-1);
        }else{
            return output;
        }
    }
}

/* Output
 JuSt_4_cUst0m_b4se64_W1th0ut_p4dd1ng
*/
```

<details>
<summary>Flag : </summary>
	hology4{JuSt_4_cUst0m_b4se64_W1th0ut_p4dd1ng}
</details>

## Referensi
- https://en.wikipedia.org/wiki/Base64
