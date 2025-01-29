package utils

import "strconv"

const (
	aInANCII = 97
	notAFile uint64 = 0xfefefefefefefefe
	notHFile uint64 = 0x7f7f7f7f7f7f7f7f
)

func CountSetBits(n uint64) int {
	count := 0
	curr := n

	for curr != 0 {
		if curr & 1 != 0 {
			count++
		}
		curr = curr >> 1
	}

	return count
}

func FindFirstSetBit(n uint64) uint64 {
	var currBit uint64 = 1
	for {
		if currBit & n != 0 {
			return currBit
		}
		currBit <<= 1
	}
}

func ShiftN(bit uint64) uint64 {
	return bit << 8
}

func ShiftNe(bit uint64) uint64 {
	return bit << 9 & notAFile
}

func ShiftE(bit uint64) uint64 {
	return bit << 1 & notAFile
}

func ShiftSe(bit uint64) uint64 {
	return bit >> 7 & notAFile
}

func ShiftS(bit uint64) uint64 {
	return bit >> 8
}

func ShiftSw(bit uint64) uint64 {
	return bit >> 9 & notHFile
}

func ShiftW(bit uint64) uint64 {
	return bit >> 1 & notHFile
}

func ShiftNw(bit uint64) uint64 {
	return bit << 7 & notHFile
}

func AlgToBit(alg string) uint64 {
	var result uint64 = 1
	file := int(alg[0]) - aInANCII
	rank, _ := strconv.Atoi(string(alg[1]))

	bitsToMove := (rank-1) * 8 + file
	return result << uint64(bitsToMove)
}

func BitToAlg(move uint64) string {
	var curr uint64 = 1
	ind := 0

	for {
		if curr == move {
			break
		}
		ind += 1
		curr <<= 1
	}

	fileI := ind % 8
	rankI := ind / 8

	rank := strconv.Itoa(int(rankI) + 1)
	file :=  string(rune(aInANCII + fileI))

	return file+rank
}