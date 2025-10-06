package advent9

import (
	"fmt"
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

type Block struct{
	firstElement 	int
	lastElement		int
	len				int
}

func MapFiles(diskImg []rune)[]Block {
	b := []Block{}
	for i := 0 ; i < len(diskImg) ; i++ {
		if diskImg[i] == FREESPACE {
			continue
		}

		runner := i

		for runner < len(diskImg) && diskImg[runner] == diskImg[i] {
			runner++
		}
		runner--
		b = append(b, Block{firstElement: i, lastElement: runner, len: runner - i + 1})
		i = runner
	}

	return b
}

func FindNextFreeBlock(diskImg []rune, startSearchFrom int) (Block, bool) {
	ret := Block{}

	for i := startSearchFrom ; i < len(diskImg) ; i++ {
		if diskImg[i] != FREESPACE {
			continue
		}
		runner := i 
		for runner < len(diskImg) && diskImg[runner] == diskImg[i] {
			runner++
		}
		runner--
		
		b := Block{firstElement: i, lastElement: runner, len: runner - i + 1}
		return b, true
	}

	return ret, false
}

func CompressDiskImg(diskImg []rune) []rune {

	blocks := MapFiles(diskImg)
	blockNumber := len(blocks) - 1
	
	for blockNumber >= 0 {
		blockToMove := blocks[blockNumber]
		freeBlock := Block{0,-1,0}
		var ok bool
		for freeBlock.lastElement < blockToMove.firstElement{
			// fmt.Println("Im here")

			if freeBlock.len >= blockToMove.len {
				tmp := slices.Clone(diskImg[freeBlock.firstElement:freeBlock.lastElement + 1])
				copy(diskImg[freeBlock.firstElement:freeBlock.lastElement + 1], diskImg[blockToMove.firstElement : blockToMove.lastElement + 1])
				copy(diskImg[blockToMove.firstElement : blockToMove.lastElement + 1], tmp)
			}
			
			freeBlock, ok = FindNextFreeBlock(diskImg, freeBlock.lastElement + 1)
			if !ok {
				// fmt.Println("BREAKING")
				break
			}
		}
		
		blockNumber--
	}

	return diskImg
}


func CheckSum(diskImg []rune) int {
	ret := 0
	for i, v := range diskImg {
		if v == FREESPACE {
			continue
		}

		ret += i * ASCIIToInt(v)
	}

	return ret
}
