package e0059

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
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
	println(text)

	return sumBytes(text)
}

func sumBytes(text []byte) int64 {

	var sum int64
	for _, t := range text {
		sum += int64(t)
	}

	return sum
}

const (
	HT = 0x09
	LF = 0x0a
	VT = 0x0b
	CR = 0x0d
)

func isAscii(b byte) bool {
	return b == HT || b == LF || b == VT || b == CR || (b >= 0x20 && b <= 0x7e)
}

func decrypt(data []byte) []byte {
	return nil
}

func scan(r io.Reader) ([]byte, error) {

	s := bufio.NewScanner(r)
	s.Split(func(data []byte, eof bool) (int, []byte, error) {
		if eof {
			return len(data), data, nil
		}
		for i, b := range data {
			if b != ',' {
				continue
			}

			return i, data[:i-1], nil
		}
		return -1, nil, bufio.ErrFinalToken
	})

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
