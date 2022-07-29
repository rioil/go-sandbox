package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main(){
	fp, err := os.Open("members.csv")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	/*
	 * trueの場合，
	 * hog"e (quote in an unquoted field)や
	 * "hog"e" (non-doubled quote in a quoted field)
	 * のような記述を容認する
	 * falseの場合，どちらも "hog""e" と書く必要がある
	 */
	reader.LazyQuotes = true
	reader.FieldsPerRecord = 2

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}