package fastmath

import "unsafe"

// https://stackoverflow.com/questions/32042673/optimized-low-accuracy-approximation-to-rootnx-n

func Sqrt(x float32) float32 {
	i := *(*uint32)(unsafe.Pointer(&x))
	// i = 0x1fc00000 + (i >> 1)
	i = 0x1fbb4f2e + (i >> 1)
	return *(*float32)(unsafe.Pointer(&i))
}

func Cbrt(x float32) float32 {
	i := *(*uint32)(unsafe.Pointer(&x))
	// i = 0x2a555556 + i/3
	i = 0x2a51067f + i/3
	return *(*float32)(unsafe.Pointer(&i))
}

func Rcpr(x float32) float32 {
	i := *(*uint32)(unsafe.Pointer(&x))
	// i = 0x7f000000 - i
	i = 0x7ef311c2 - i
	return *(*float32)(unsafe.Pointer(&i))
}

func RSqrt(x float32) float32 {
	i := *(*uint32)(unsafe.Pointer(&x))
	// i = 0x5f400000 - (i >> 1)
	i = 0x5f37642f - (i >> 1)
	return *(*float32)(unsafe.Pointer(&i))
}

func RCbrt(x float32) float32 {
	i := *(*uint32)(unsafe.Pointer(&x))
	// i = 0x54aaaaaa - i/3
	i = 0x54a232a3 - i/3
	return *(*float32)(unsafe.Pointer(&i))
}

func Log(x, b float32) float32 {
	i := *(*int32)(unsafe.Pointer(&x))
	i = i - 0x3f800000
	j := *(*int32)(unsafe.Pointer(&b))
	j = j - 0x3f800000
	return float32(i) / float32(j)
}

func Log2(x float32) float32 {
	i := *(*int32)(unsafe.Pointer(&x))
	x = float32(i - 0x3f800000)
	i = *(*int32)(unsafe.Pointer(&x))
	i = i - 0xb800000
	return *(*float32)(unsafe.Pointer(&i))
}

func LogE(x float32) float32 {
	i := *(*int32)(unsafe.Pointer(&x))
	i = i - 0x3f800000
	return float32(i) * 8.262958e-8
}

func Log10(x float32) float32 {
	i := *(*int32)(unsafe.Pointer(&x))
	i = i - 0x3f800000
	return float32(i) * 3.5885572e-8
}

func Pow(x, p float32) float32 {
	i := *(*int32)(unsafe.Pointer(&x))
	i = i - 0x3f800000
	i = int32(float32(i) * p)
	i = i + 0x3f800000
	return *(*float32)(unsafe.Pointer(&i))
}

func Pow2(x float32) float32 {
	i := *(*int32)(unsafe.Pointer(&x))
	i = i + 0xb800000
	i = int32(*(*float32)(unsafe.Pointer(&i)))
	i = i + 0x3f800000
	return *(*float32)(unsafe.Pointer(&i))
}

func PowE(x float32) float32 {
	i := (int32)(0xb8aa3b * x)
	i = i + 0x3f800000
	return *(*float32)(unsafe.Pointer(&i))
}

func Pow10(x float32) float32 {
	i := (int32)(0x1a934f0 * x)
	i = i + 0x3f800000
	return *(*float32)(unsafe.Pointer(&i))
}
