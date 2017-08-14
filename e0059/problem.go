package e0059

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func Problem() int64 {

	f, err := open(
		"p059_cipher.txt",
		"https://projecteuler.net/project/resources/p059_cipher.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cipher, err := scan(f)
	if err != nil {
		log.Fatal(err)
	}

	text := decrypt(cipher)
	println(string(text))

	return sumBytes(text)
}

func isASCII(b byte) bool {
	return b >= 0x0a && b <= 0x7e
}

var (
	illPunct  = regexp.MustCompile(`[;,?!]\S`)
	illBraces = regexp.MustCompile(`\S\(|\)\S`)
)

func illformed(text []byte) bool {

	if illPunct.Find(text) != nil {
		return true
	}

	if illBraces.Find(text) != nil {
		return true
	}

	return false
}

func keygen() func() []byte {

	var key = [3]byte{'a' - 1, 'a', 'a'}
	return func() []byte {

		for i := 0; i < 3; i++ {

			if key[i] < 'z' {
				key[i]++
				return key[:]
			}

			key[i] = 'a'
		}

		return nil
	}
}

func decrypt(data []byte) []byte {

	gen := keygen()

	var tmp = make([]byte, len(data))
	for key := gen(); key != nil; key = gen() {

		ascii := true
		copy(tmp, data)

		for i := 0; i < len(tmp); i++ {
			tmp[i] = data[i] ^ key[i%3]
			if !isASCII(tmp[i]) {
				ascii = false
				break
			}
		}

		if !ascii {
			continue
		}

		if illformed(tmp) {
			continue
		}

		return tmp
	}

	return nil
}

func sumBytes(text []byte) int64 {

	var sum int64
	for _, t := range text {
		sum += int64(t)
	}

	return sum
}

func atComma(data []byte, eof bool) (int, []byte, error) {
	if eof {
		return 0, nil, nil
	}
	for i, b := range data {
		if b != ',' && b != 0xa {
			continue
		}

		return i + 1, data[:i], nil
	}
	return -1, nil, fmt.Errorf("eof")
}

func scan(r io.Reader) ([]byte, error) {

	s := bufio.NewScanner(r)
	s.Split(atComma)

	var cipher []byte
	for s.Scan() {
		t := s.Text()
		v, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}

		cipher = append(cipher, byte(v))
	}
	if s.Err() != nil {
		return nil, s.Err()
	}

	return cipher, nil
}

func open(fileName, url string) (*os.File, error) {

	f, err := os.Open(fileName)
	if os.IsNotExist(err) {

		f, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			return nil, err
		}

		f, err = os.Open(fileName)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	return f, nil
}
