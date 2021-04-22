package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	key := 3
	word := "meet me after the party"
	fmt.Println("LENGTH WORD", len(word))

	result := encrypt(key, word)

	fmt.Println(result)

	//TODO: Automate this process
	//	fmt.Println(result[:idx+1])
	//	fmt.Println(result[idx+1 : (idx+1)*2])
	//	fmt.Println(result[(idx+1)*2 : ((idx+1)*2)+idx+1])
	//	fmt.Println(result[((idx+1)*2)+idx+1 : (idx+1)*2+idx+1+idx])
	decrypt(result, key)

}

func encrypt(key int, word string) string {
	arr := make([][]string, key)

	count := 0
	for _, l := range word {
		if count == key {
			count = 0
		}

		if string(l) == " " {
			l = '_'
		}
		arr[count] = append(arr[count], string(l))
		count++
	}

	fmt.Println("DEBUG", arr)

	var result string
	for _, n := range arr {
		result += strings.Join(n, "")
	}

	return result
}

func decrypt(encWord string, key int) {
	//fmt.Println(encWord)

	idx := len(encWord) / key
	fmt.Println("IDX", idx)
	count := 0
	arr := make([][]string, key)

	lastAccum := 0

	x, y := 0, 0
	for {
		if count == key {
			break
		}

		accum := int(math.Ceil(float64(len(encWord)-count) / float64(key)))

		//		fmt.Println(accum)
		//		fmt.Println(lastAccum)
		if count != 0 {
			//x += (idx + 1)
			if accum != lastAccum {
				fmt.Println("NOT", accum, lastAccum)
				x += (accum + 1)
			} else {
				x += accum
			}
		}
		lastAccum = accum

		//y += (idx + 1)
		y += accum
		//int(math.Ceil(float64(len(encWord)-count) / float64(key)))
		if count != key-1 {
			fmt.Println(encWord[x:y])
			arr[count] = strings.Split(encWord[x:y], "")
		} else {
			fmt.Println(encWord[x:])
			arr[count] = strings.Split(encWord[x:], "")
		}

		count++
	}

	fmt.Println("DEBUG", arr)
}
