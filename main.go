package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
)

func main() {

	var (
		word      string
		key       int
		isEncrypt bool
	)

	flag.StringVar(&word, "word", "", "word to encrypt or decrypt")
	flag.IntVar(&key, "key", 0, "depth key")
	flag.BoolVar(&isEncrypt, "encrypt", false, "true to encrypt default set to false (decrypt)")

	flag.Parse()

	if key != 0 && word != "" {
		switch isEncrypt {
		case false:
			fmt.Println(decrypt(word, key))
			return
		case true:
			fmt.Println(encrypt(key, word))
			return
		}
	}

	flag.PrintDefaults()
	return

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

	var result string
	for _, n := range arr {
		result += strings.Join(n, "")
	}

	return result
}

func decrypt(encWord string, key int) string {

	//idx := len(encWord) / key
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
				//fmt.Println("NOT", accum, lastAccum)
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
			//fmt.Println(encWord[x:y])
			arr[count] = strings.Split(encWord[x:y], "")
		} else {
			//fmt.Println(encWord[x:])
			arr[count] = strings.Split(encWord[x:], "")
		}

		count++
	}

	xIndex := 0
	yIndex := 0

	var result string
	for _, n := range arr {
		for _ = range n {
			if xIndex == key {
				xIndex = 0
				yIndex++
			}
			result += arr[xIndex][yIndex]
			xIndex++
		}
	}

	return result
}
