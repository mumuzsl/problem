package idea

import "strconv"

// 把多维数组转换为相应的一维数组坐标
// size: 多维数组各维度大小
// point: 多维数组的坐标
func MultiArrayPoint2LinearArrayPoint(size []uint64, point []uint64) string {
	sizeLen := len(size)
	result := "0"
	temp := make([]string, sizeLen)
	toStr := func(v uint64) string { return strconv.FormatUint(v, 10) }

	temp[0] = toStr(size[0])
	for i := 1; i < sizeLen-1; i++ {
		temp[i] = LargeMultiply(temp[i-1], toStr(size[i]))
	}
	for i := sizeLen - 1; i >= 1; i-- {
		tr := LargeMultiply(toStr(point[i]), temp[i])
		result = LargeAddition(result, tr)
	}

	return LargeAddition(result, toStr(point[0]))
}

// 大数乘法
var LargeMultiply func(a, b string) string

//v大数加法
var LargeAddition func(a, b string) string
