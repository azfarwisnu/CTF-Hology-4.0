#include <unistd.h>
#include <stdio.h>
#include <string.h>

void vuln() {
	char buf[100];
	setbuf(stdin, buf);
	read(0, buf, 256);
}

int main() {
	setvbuf(stdin, 0, 2, 0);  //no buffering in stdin and stdout
        setvbuf(stdout, 0, 2, 0);
	char buf[100] = "Sibuk kerja, soalnya gampang aja. Good Luck\n";
	setbuf(stdout, buf);
	write(1, buf, strlen(buf));
	vuln();
	return 0;
}
