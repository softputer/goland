package gauth

import (
	"crypto/des"
)

var parity = []byte{
	0x01, 0x02, 0x04, 0x07, 0x08, 0x0b, 0x0d, 0x0e,
	0x10, 0x13, 0x15, 0x16, 0x19, 0x1a, 0x1c, 0x1f,
	0x20, 0x23, 0x25, 0x26, 0x29, 0x2a, 0x2c, 0x2f,
	0x31, 0x32, 0x34, 0x37, 0x38, 0x3b, 0x3d, 0x3e,
	0x40, 0x43, 0x45, 0x46, 0x49, 0x4a, 0x4c, 0x4f,
	0x51, 0x52, 0x54, 0x57, 0x58, 0x5b, 0x5d, 0x5e,
	0x61, 0x62, 0x64, 0x67, 0x68, 0x6b, 0x6d, 0x6e,
	0x70, 0x73, 0x75, 0x76, 0x79, 0x7a, 0x7c, 0x7f,
	0x80, 0x83, 0x85, 0x86, 0x89, 0x8a, 0x8c, 0x8f,
	0x91, 0x92, 0x94, 0x97, 0x98, 0x9b, 0x9d, 0x9e,
	0xa1, 0xa2, 0xa4, 0xa7, 0xa8, 0xab, 0xad, 0xae,
	0xb0, 0xb3, 0xb5, 0xb6, 0xb9, 0xba, 0xbc, 0xbf,
	0xc1, 0xc2, 0xc4, 0xc7, 0xc8, 0xcb, 0xcd, 0xce,
	0xd0, 0xd3, 0xd5, 0xd6, 0xd9, 0xda, 0xdc, 0xdf,
	0xe0, 0xe3, 0xe5, 0xe6, 0xe9, 0xea, 0xec, 0xef,
	0xf1, 0xf2, 0xf4, 0xf7, 0xf8, 0xfb, 0xfd, 0xfe,
}

func des56to64(k56 []byte) []byte {
	var hi, lo uint32

	k64 := make([]byte, 8)

	hi = uint32(k56[0])<<24 | uint32(k56[1])<<16 | uint32(k56[2])<<8 | uint32(k56[3])
	lo = uint32(k56[4])<<24 | uint32(k56[5])<<16 | uint32(k56[6])<<8

	k64[0] = parity[(hi>>25)&0x7f]
	k64[1] = parity[(hi>>18)&0x7f]
	k64[2] = parity[(hi>>11)&0x7f]
	k64[3] = parity[(hi>>4)&0x7f]
	k64[4] = parity[((hi<<3)|(lo>>29))&0x7f]
	k64[5] = parity[(lo>>22)&0x7f]
	k64[6] = parity[(lo>>15)&0x7f]
	k64[7] = parity[(lo>>8)&0x7f]

	return k64
}

// key is 7 in length, convert to 64 bit key
func DesEncrypt(key, buf []byte) {
	var last []byte

	if len(buf) < 8 {
		return
	}

	key64 := des56to64(key)
	ci, err := des.NewCipher(key64)
	if err != nil {
		panic(err)
	}

	n := len(buf)
	n--
	r := n % 7
	n /= 7

	for i := 0; i < n; i++ {
		ci.Encrypt(buf, buf)
		last = buf
		buf = buf[7:]
	}

	if r > 0 {
		ci.Encrypt(last[r:], last[r:])
	}
}
