package data

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

func Read(path string) ([]int, error) {
	fp, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	if !scanner.Scan() {
		err := errors.New("Invalid format")
		return nil, err
	}

	itemNum, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return nil, err
	}

	items := make([]int, itemNum)

	for i := 0; i < itemNum; i++ {
		if !scanner.Scan() {
			return nil, errors.New("Missing item")
		}

		item, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
