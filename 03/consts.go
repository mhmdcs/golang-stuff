package main

func main() {
	const = (
		a = 1 // 1
		b = 2 * 1024 // 2048, two kilobytes
		c = b << 3 // 16384, left shift bitwise operation
		// 2048 is 100000000000, shifting the bits of 2048 to the left by 3 positions results in 1000000000000000 which is 16384 in decimal

		g uint8 = 0x07 // 7
		h uint8 = g & 0x03 // 3, because performing AND bitwise on 00000111 AND 00000011 results in 00000011 (3)
		
		s = "a string"
		t = len(s) // 8
		u = s[2:] // syntax error, consts must be known at compile time, can't access index unless during runtime
	)
}
