package main

import (
	"fmt"
	"io/ioutil"
)

type Jammu struct {
	korean string
	roman  string
}

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	input := string(dat)
	var jammus []Jammu
	for _, a := range input {
		//	readInput(string(a))

		jammus = append(jammus, findhan(a)...)
	}

	for _, jam := range jammus {
		fmt.Print(jam.roman)
	}
}
func readInput(decodeMe string) {
	fmt.Print(decodeMe, ": ")
	var input string
	fmt.Scanln(&input)

}

func findhan(hangul rune) (jammus []Jammu) {
	ans := int(hangul) - 44032
	initial := 0
	mid := 0
	last := 0
	if (ans >= 0) && ans <= 11184 {
		for init := 0; init < 20; init++ {
			if init*588 > ans {
				initial = init - 1
				break
			}
		}

		for midd := 0; midd < 22; midd++ {
			if initial*588+midd*28 > ans {
				mid = midd - 1
				break
			}
		}

		for lastt := 0; lastt < 29; lastt++ {
			if initial*588+mid*28+lastt == ans {
				last = lastt
				break
			}
		}
		jammus = append(jammus, GetFirstChar(initial))
		jammus = append(jammus, GetMidChar(mid))
		jammus = append(jammus, GetLastChar(last))
	} else {
		//	fmt.Println(string(hangul))
		jammus = append(jammus, Jammu{korean: string(hangul), roman: string(hangul)})
	}
	return
}

// Returns the first char from the Hangul
func GetFirstChar(init int) (jamm Jammu) {
	switch init {
	case 0:
		return Jammu{korean: "ㄱ", roman: "g"}
	case 1:
		return Jammu{korean: "ㄲ", roman: "kk"}
	case 2:
		return Jammu{korean: "ㄴ", roman: "n"}
	case 3:
		return Jammu{korean: "ㄷ", roman: "d"}
	case 4:
		return Jammu{korean: "ㄸ", roman: "tt"}
	case 5:
		return Jammu{korean: "ㄹ", roman: "r"}
	case 6:
		return Jammu{korean: "ㅁ", roman: "m"}
	case 7:
		return Jammu{korean: "ㅂ", roman: "b"}
	case 8:
		return Jammu{korean: "ㅃ", roman: "pp"}
	case 9:
		return Jammu{korean: "ㅅ", roman: "s"}
	case 10:
		return Jammu{korean: "ㅆ", roman: "ss"}
	case 11:
		return Jammu{korean: "ㅇ", roman: ""}
	case 12:
		return Jammu{korean: "ㅈ", roman: "j"}
	case 13:
		return Jammu{korean: "ㅉ", roman: "jj"}
	case 14:
		return Jammu{korean: "ㅊ", roman: "ch"}
	case 15:
		return Jammu{korean: "ㅋ", roman: "k"}
	case 16:
		return Jammu{korean: "ㅌ", roman: "t"}
	case 17:
		return Jammu{korean: "ㅍ", roman: "p"}
	case 18:
		return Jammu{korean: "ㅎ", roman: "h"}
	}
	return Jammu{korean: "", roman: ""}
}

func GetMidChar(mid int) (jamm Jammu) {
	switch mid {
	case 0:
		return Jammu{korean: "ㅏ", roman: "a"}
	case 1:
		return Jammu{korean: "ㅐ", roman: "ae"}
	case 2:
		return Jammu{korean: "ㅑ", roman: "ya"}
	case 3:
		return Jammu{korean: "ㅒ", roman: "yae"}
	case 4:
		return Jammu{korean: "ㅓ", roman: "eo"}
	case 5:
		return Jammu{korean: "ㅔ", roman: "e"}
	case 6:
		return Jammu{korean: "ㅕ", roman: "yeo"}
	case 7:
		return Jammu{korean: "ㅖ", roman: "ye"}
	case 8:
		return Jammu{korean: "ㅗ", roman: "o"}
	case 9:
		return Jammu{korean: "ㅘ", roman: "wa"}
	case 10:
		return Jammu{korean: "ㅙ", roman: "wae"}
	case 11:
		return Jammu{korean: "ㅚ", roman: "oe"}
	case 12:
		return Jammu{korean: "ㅛ", roman: "yo"}
	case 13:
		return Jammu{korean: "ㅜ", roman: "u"}
	case 14:
		return Jammu{korean: "ㅝ", roman: "wo"}
	case 15:
		return Jammu{korean: "ㅞ", roman: "we"}
	case 16:
		return Jammu{korean: "ㅟ", roman: "wi"}
	case 17:
		return Jammu{korean: "ㅠ", roman: "yu"}
	case 18:
		return Jammu{korean: "ㅡ", roman: "eu"}
	case 19:
		return Jammu{korean: "ㅢ", roman: "ui"}
	case 20:
		return Jammu{korean: "ㅣ", roman: "i"}

	}
	return Jammu{korean: "", roman: ""}
}

func GetLastChar(last int) (jamm Jammu) {
	switch last {
	case 0:
		return Jammu{korean: "", roman: ""}
	case 1:
		return Jammu{korean: "ㄱ", roman: "k"}
	case 2:
		return Jammu{korean: "ㄲ", roman: "k"}
	case 3:
		return Jammu{korean: "ㄳ", roman: "?"}
	case 4:
		return Jammu{korean: "ㄴ", roman: "n"}
	case 5:
		return Jammu{korean: "ㄵ", roman: "?"}
	case 6:
		return Jammu{korean: "ㄶ", roman: "?"}
	case 7:
		return Jammu{korean: "ㄷ", roman: "t"}
	case 8:
		return Jammu{korean: "ㄹ", roman: "l"}
	case 9:
		return Jammu{korean: "ㄺ", roman: "?"}
	case 10:
		return Jammu{korean: "ㄻ", roman: "?"}
	case 11:
		return Jammu{korean: "ㄼ", roman: "?"}
	case 12:
		return Jammu{korean: "ㄽ", roman: "?"}
	case 13:
		return Jammu{korean: "ㄾ", roman: "?"}
	case 14:
		return Jammu{korean: "ㄿ", roman: "?"}
	case 15:
		return Jammu{korean: "ㅀ", roman: "?"}
	case 16:
		return Jammu{korean: "ㅁ", roman: "m"}
	case 17:
		return Jammu{korean: "ㅂ", roman: "p"}
	case 18:
		return Jammu{korean: "ㅄ", roman: "?"}
	case 19:
		return Jammu{korean: "ㅅ", roman: "t"}
	case 20:
		return Jammu{korean: "ㅆ", roman: "t"}
	case 21:
		return Jammu{korean: "ㅇ", roman: "ng"}
	case 22:
		return Jammu{korean: "ㅈ", roman: "t"}
	case 23:
		return Jammu{korean: "ㅊ", roman: ""}
	case 24:
		return Jammu{korean: "ㅋ", roman: "k"}
	case 25:
		return Jammu{korean: "ㅌ", roman: "t"}
	case 26:
		return Jammu{korean: "ㅍ", roman: "p"}
	case 27:
		return Jammu{korean: "ㅎ", roman: "t"}

	}
	return Jammu{korean: "", roman: ""}
}
