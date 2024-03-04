package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Convert(input, banner string) (string, error) {
	if input == "" {
		return "", nil
	}

	ok, err := ValidInput(input)
	if !ok {
		return "", err
	}

	alphabet, err := getAlphab(banner)
	if err != nil {
		return "", err
	}

	input = strings.ReplaceAll(input, "\\n", "\n")

	words := strings.Split(input, "\n")

	result := make([]string, 0, len(words))

	for indexWord, word := range words {
		if word == "" {
			result = append(result, word)
			continue
		}

		var middleResult string

		for indexHeight := 0; indexHeight < height; indexHeight++ {
			for _, letter := range word {
				middleResult += alphabet[letter][indexHeight]
			}

			if indexHeight != height-1 || indexWord == len(words)-1 {
				middleResult += "\n"
			}
		}

		result = append(result, middleResult)
	}

	result = adjustNewLines(result)

	res := strings.Join(result, "\n")

	return res, nil
}

func getAlphab(banner string) (map[rune][]string, error) {
	cwd, _ := os.Getwd()

	// if banner != "shadow.txt" || banner != "standard.txt" || banner != "thinkertoy.txt" {
	// 	fmt.Println("FIND")
	// 	return nil, fmt.Errorf("banner not found")
	// }

	if banner == "shadow.txt" {
		banner = "shadow.txt"
	} else if banner == "standard.txt" {
		banner = "standard.txt"
	} else if banner == "thinkertoy.txt" {
		banner = "thinkertoy.txt"
	} else {
		return nil, fmt.Errorf("error")
	}
	cwd = TrimCwd(cwd) + "/internal/banner/" + banner

	file, err := os.Open(cwd)
	if err != nil {
		return nil, fmt.Errorf("os error:%w", err)
	}

	scanner := bufio.NewScanner(file)
	textFromFile, err := os.ReadFile(cwd)
	if err != nil {
		return nil, fmt.Errorf("os error:%w", err)
	}

	hashStandard := uint64(8250112135784318067)
	hashShadow := uint64(8377067621923326644)
	hashTinkertoy := uint64(4863852022380994373)
	hashRead := strToHash(textFromFile)
	if (banner == "standard.txt" && hashStandard != hashRead) ||
		(banner == "shadow.txt" && hashShadow != hashRead) ||
		(banner == "thinkertoy.txt" && hashTinkertoy != hashRead) {
		return nil, errors.New("the banner was modified")
	}

	alphabet := make(map[rune][]string)

	skip := true

	var indexRune rune = 32

	for scanner.Scan() {
		if skip {
			skip = false

			continue
		}

		letter := make([]string, height)

		for i := 0; i < height; i++ {
			letter[i] = scanner.Text()

			if i != height-1 {
				scanner.Scan()
			}
		}

		alphabet[indexRune] = letter
		indexRune++

		skip = true
	}

	return alphabet, nil
}

func adjustNewLines(result []string) []string {
	onlyNewLine := true

	for _, v := range result {
		if v != "" {
			onlyNewLine = false
			break
		}
	}

	if onlyNewLine {
		return result
	}

	toAdd := false

	for _, v := range result {
		if v == "" {
			toAdd = true
			continue
		}
		toAdd = false
	}

	if toAdd {
		result = append(result, "")
	}

	return result
}
