// Code generated by mksizeclasses_mcu64.go; DO NOT EDIT.
//go:generate go run mksizeclasses_mcu64.go

// +build noos,riscv64

package runtime

// class  bytes/obj  bytes/span  objects  tail waste  max waste
//     1          8        2048      256           0     87.50%
//     2         16        2048      128           0     43.75%
//     3         32        2048       64           0     46.88%
//     4         48        2048       42          32     32.32%
//     5         64        2048       32           0     23.44%
//     6         80        2048       25          48     20.65%
//     7         96        2048       21          32     16.94%
//     8        112        2048       18          32     14.75%
//     9        128        2048       16           0     11.72%
//    10        144        2048       14          32     11.82%
//    11        160        2048       12         128     15.04%
//    12        176        2048       11         112     13.53%
//    13        192        2048       10         128     13.57%
//    14        224        2048        9          32     15.19%
//    15        256        2048        8           0     12.11%
//    16        288        2048        7          32     12.16%
//    17        320        2048        6         128     15.33%
//    18        352        4096       11         224     13.79%
//    19        384        2048        5         128     13.82%
//    20        416        4096        9         352     15.41%
//    21        512        2048        4           0     18.55%
//    22        576        4096        7          64     12.33%
//    23        640        2048        3         128     15.48%
//    24        768        6144        8           0     16.54%
//    25        768        4096        5         256      6.13%
//    26        832        6144        7         320     12.39%
//    27       1024        2048        2           0     18.65%
//    28       1152        6144        5         384     16.59%
//    29       1280        4096        3         256     15.55%
//    30       1536        6144        4           0     16.60%
//    31       1664       10240        6         256      9.94%
//    32       2048        2048        1           0     18.70%
//    33       2560       10240        4           0     19.96%
//    34       2688        8192        3         128      6.21%
//    35       3072        6144        2           0     12.47%
//    36       3328       10240        3         256      9.97%
//    37       4096        4096        1           0     18.73%

const (
	_MaxSmallSize   = 4096
	smallSizeDiv    = 8
	smallSizeMax    = 256
	largeSizeDiv    = 128
	_NumSizeClasses = 38
	_PageShift      = 11
)

var class_to_size = [_NumSizeClasses]uint16{0, 8, 16, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 224, 256, 288, 320, 352, 384, 416, 512, 576, 640, 768, 768, 832, 1024, 1152, 1280, 1536, 1664, 2048, 2560, 2688, 3072, 3328, 4096}
var class_to_allocnpages = [_NumSizeClasses]uint8{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 2, 3, 1, 3, 2, 3, 5, 1, 5, 4, 3, 5, 2}

type divMagic struct {
	shift    uint8
	shift2   uint8
	mul      uint16
	baseMask uint16
}

var class_to_divmagic = [_NumSizeClasses]divMagic{{0, 0, 0, 0}, {3, 0, 1, 65528}, {4, 0, 1, 65520}, {5, 0, 1, 65504}, {4, 9, 171, 0}, {6, 0, 1, 65472}, {4, 9, 103, 0}, {5, 7, 43, 0}, {4, 10, 147, 0}, {7, 0, 1, 65408}, {4, 9, 57, 0}, {5, 9, 103, 0}, {4, 11, 187, 0}, {6, 7, 43, 0}, {5, 8, 37, 0}, {8, 0, 1, 65280}, {5, 9, 57, 0}, {6, 6, 13, 0}, {5, 11, 187, 0}, {7, 5, 11, 0}, {5, 10, 79, 0}, {9, 0, 1, 65024}, {6, 9, 57, 0}, {7, 6, 13, 0}, {8, 5, 11, 0}, {8, 5, 11, 0}, {6, 10, 79, 0}, {10, 0, 1, 64512}, {7, 8, 29, 0}, {8, 6, 13, 0}, {9, 5, 11, 0}, {7, 10, 79, 0}, {11, 0, 1, 63488}, {9, 6, 13, 0}, {7, 10, 49, 0}, {10, 3, 3, 0}, {8, 6, 5, 0}, {12, 0, 1, 61440}}
var size_to_class8 = [smallSizeMax/smallSizeDiv + 1]uint8{0, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 14, 14, 15, 15, 15, 15}
var size_to_class128 = [(_MaxSmallSize-smallSizeMax)/largeSizeDiv + 1]uint8{15, 19, 21, 23, 24, 27, 27, 28, 29, 30, 30, 31, 32, 32, 32, 33, 33, 33, 33, 34, 35, 35, 35, 36, 36, 37, 37, 37, 37, 37, 37}
