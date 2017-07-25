package e0054

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/shogg/euler54/poker"
)

func E0054() int64 {

	f, err := open("p054_poker.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			continue
		}

		h1 := poker.ParseHand(line[:14])
		h2 := poker.ParseHand(line[15:])

		if !h1.Less(h2) {
			sum++
		}
	}

	return int64(sum)
}

func open(fileName string) (*os.File, error) {

	f, err := os.Open(fileName)
	if os.IsNotExist(err) {

		f, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		resp, err := http.Get(
			"https://projecteuler.net/project/resources/p054_poker.txt")
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			return nil, err
		}

		f, err = os.Open("p054_poker.txt")
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	return f, nil
}
