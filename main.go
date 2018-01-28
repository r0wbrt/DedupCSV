/* Copyright 2018 Robert Christian Taylor. All Rights Reserved
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Expected first argument to be the CSV column (Count starts at 0), the rest of the remaining arguments to be the path to the CSV file.")
		return
	}

	var recordMap map[string][]string = make(map[string][]string)
	csvFile := strings.Join(os.Args[2:], " ")
	file, err := os.Open(csvFile)
	if err != nil {
		return
	}

	var column int64

	column, err = strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}

	if column < 0 {
		panic("The column must be positive or zero.")
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}

		if len(record) < int(column) {
			continue
		}

		id := record[column]

		_, ok := recordMap[id]
		if !ok {
			recordMap[id] = record
		}
	}

	csvWriter := csv.NewWriter(os.Stdout)
	for _, v := range recordMap {
		csvWriter.Write(v)
	}

	csvWriter.Flush()

}
