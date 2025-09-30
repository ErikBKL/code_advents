package advent9

import (
	"os"
	"slices"
)

const (
	FREESPACE = '.'
)

func DiskChecksum(pathToFile string) (int, error) {
	// read file into diskMap (a slice of runes)
	diskMap, err := ReadFile(pathToFile)
	if err != nil {
		return 0, err
	}
	// expand the diskMap into a diskImg (id's and dots)
	diskImg := MapToImg(diskMap)
	// compress the diskImg
	diskImg = CompressDiskImg(diskImg)
	// calculate checksum
	checkSum := CheckSum(diskImg)
	return checkSum, nil
}

func ReadFile(pathToFile string) ([]rune, error) {
	bytes, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, err
	}

	return []rune(string(bytes)), nil
}

func ASCIIToInt(n rune) int {
	return int(n - '0')
}

func MapToImg(diskMap []rune) []rune {
	id := '0'
	diskImg := []rune{}

	for i, v := range diskMap {
		if i%2 == 0 {
			for j := 0; j < ASCIIToInt(v); j++ {
				diskImg = append(diskImg, id)
			}
			id++
		} else {
			for j := 0; j < ASCIIToInt(v); j++ {
				diskImg = append(diskImg, FREESPACE)
			}
		}
	}

	return diskImg
}

func CompressDiskImg(diskImg []rune) []rune {
	// set right to last element in slice
	right := len(diskImg) - 1
	// set left to first element
	left := 0

	// while left < right
	for left < right {
		// rightRunner = right
		for diskImg[right] != FREESPACE && diskImg[left] == FREESPACE && left < right {

			endFreeChunk, freeChunkLen := StatsFreeChunk(diskImg, left)
			startBlock, endBlock, isMatch := TryFindFittingBlock(diskImg, right, freeChunkLen, endFreeChunk)
			if isMatch {
				tmp := slices.Clone(diskImg[left:endFreeChunk])
				copy(diskImg[left:endFreeChunk], diskImg[startBlock:endBlock])
				copy(diskImg[startBlock:endBlock], tmp)
			} else {
				right = EndOfPrevBlock(diskImg, right)
			}

			left += endBlock - startBlock
		}

		for diskImg[left] != FREESPACE && left < right {
			left++
		}

		for diskImg[right] == FREESPACE && right > left {
			right--
		}
	}
	return diskImg
}

func EndOfPrevBlock(diskImg []rune, right int) int {
	curr := diskImg[right]
	for diskImg[right] == curr {
		right--
	}

	return right
}

func TryFindFittingBlock(diskImg []rune, right, freeChunkLen, freeSpaceLimit int) (int, int, bool) {
	leftOfBlock := right

	for leftOfBlock > freeSpaceLimit {
		for diskImg[leftOfBlock-1] == diskImg[right] {
			leftOfBlock--
		}

		if right-leftOfBlock+1 <= freeChunkLen {
			return leftOfBlock, right + 1, true
		} else {
			right = leftOfBlock - 1
			leftOfBlock--
			for diskImg[right] == FREESPACE {
				right--
				leftOfBlock--
			}
		}
	}

	return 0, 0, false
}

func StatsFreeChunk(diskImg []rune, left int) (int, int) {
	end := left
	for diskImg[end] == FREESPACE {
		end++
	}

	blockLen := end - left
	return end, blockLen
}

func CheckSum(diskImg []rune) int {
	ret := 0
	for i, v := range diskImg {
		if v == FREESPACE {
			break
		}

		ret += i * ASCIIToInt(v)
	}

	return ret
}
