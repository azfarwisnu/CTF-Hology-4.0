FROM ubuntu:16.04

#base install
RUN dpkg --add-architecture i386
RUN apt update
RUN apt install -y libc6:i386 gcc socat netcat cmake

# Service / Chall
COPY service /service
RUN chown -R root:root /service
RUN chmod -R 555 /service

# Flag
COPY flag.txt /service

# Finishing
WORKDIR /service
RUN useradd ctf
RUN echo 2 | sudo tee /proc/sys/kernel/randomize_va_space
RUN gcc pwnpwnclub.c -fno-stack-protector -m32 -no-pie -o pwnpwnclub.o
EXPOSE 8000
ENTRYPOINT socat TCP4-LISTEN:8000,reuseaddr,fork,su=ctf EXEC:'./pwnpwnclub.o'
