#!/usr/bin/python3

def java_hash(s, overflow=False):
	def int_overflow(val):
		maxint = 2**31 - 1
		if not -maxint-1 <= val <= maxint:
			val = (val + (maxint + 1)) % (2 * (maxint + 1)) - maxint - 1
		return val

	h = 0
	for i in s:
		h *= 31
		h += ord(i)
	return h if not overflow else int_overflow(h)


def pad(text, char, length):
	return text + char*(length - len(text))


hc  = 551276894
hcu = 1334888766
found = False


def brute(text, idx, length, key='abcedfghijklmnopqrstuvwxyz0123456789{}_?'):
	global hc, hcu, found
	if len(text) == length:
		if java_hash(text) == hc and java_hash(text.upper(), overflow=True) == hcu:
			print('hology4{' + text + '}')
			found = True
	else:
		for i in key:
			text_brute = text + i
			if java_hash(pad(text_brute, '0', 8)) <= hc and java_hash(pad(text_brute, 'z', 8)) >= hc:
				brute(text_brute, idx + 1, length, key)


def main():
	global found, hc
	while True:
		brute('', 0, 8)
		hc += 2**32
		if found:
			break


if __name__ == '__main__':
	main()
