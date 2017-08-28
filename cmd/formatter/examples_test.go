/*
 * This file is part of arduino-cli.
 *
 * arduino-cli is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2017 BCMI LABS SA (http://www.arduino.cc/)
 */

package formatter_test

import (
	"fmt"

	"github.com/bcmi-labs/arduino-cli/cmd/formatter"
)

type TestStruct struct {
	Value int `json:"value"`
}

func (ts TestStruct) String() string {
	return fmt.Sprint("VALUE = ", ts.Value)
}

func ExampleJSONFormatter_Print() {
	var example struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
		Field3 struct {
			Inner1 string  `json:"inner1"`
			Inner2 float32 `json:"inner2"`
		} `json:"field3"`
	}

	example.Field1 = "test"
	example.Field2 = 10
	example.Field3.Inner1 = "inner test"
	example.Field3.Inner2 = 10.432412

	var jf formatter.JSONFormatter
	jf.Print(example)

	var example2 float64 = 3.14
	jf.Print(example2)

	// Output:
	// {"field1":"test","field2":10,"field3":{"inner1":"inner test","inner2":10.432412}}
	//
}

func ExampleJSONFormatter_Format() {
	var example struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
		Field3 struct {
			Inner1 string  `json:"inner1"`
			Inner2 float32 `json:"inner2"`
		} `json:"field3"`
	}

	example.Field1 = "test"
	example.Field2 = 10
	example.Field3.Inner1 = "inner test"
	example.Field3.Inner2 = 10.432412

	var jf formatter.JSONFormatter
	var result string
	result, err := jf.Format(example)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("RESULT:", result)
	}

	var example2 float32 = 3.14
	result, err = jf.Format(example2)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else if result == "" {
		fmt.Println("RESULT: <empty string>")
	}

	// Output:
	// RESULT: {"field1":"test","field2":10,"field3":{"inner1":"inner test","inner2":10.432412}}
	// RESULT: <empty string>
}

func ExampleJSONFormatter_Print_debug() {
	valid := TestStruct{20}
	invalid := "invalid"
	jf := formatter.JSONFormatter{
		Debug: false,
	}
	//using struct
	jf.Print(valid)

	//using string (invalid sine it's not a struct or a map)
	jf.Print(invalid)

	jf.Debug = true
	jf.Print(valid)
	jf.Print(invalid)

	//using map
	newValue := make(map[string]int)
	newValue["value2"] = 10

	jf.Print(newValue)
	// Output:
	// {"value":20}
	//
	// {"value":20}
	// Only structs and maps values are accepted
	// {"value2":10}
}

func ExampleFormat() {
	formatter.SetFormatter("text")
	result, err := formatter.Format(TestStruct{5})
	if err == nil {
		fmt.Println("RESULT: ", result)
	} else {
		fmt.Println("ERROR: ", err)
	}

	formatter.SetFormatter("json")
	result, err = formatter.Format(TestStruct{10})
	if err == nil {
		fmt.Println("RESULT: ", result)
	} else {
		fmt.Println("ERROR: ", err)
	}
	result, err = formatter.Format(5)
	if err == nil {
		if result == "" {
			result = "<empty string>"
			fmt.Println("RESULT: ", result)
		} else {
			fmt.Println("unexpected RESULT: ", result)
		}
	} else {
		fmt.Sprintln("ERROR: ", err)
	}

	// Output:
	// RESULT:  VALUE = 5
	// RESULT:  {"value":10}
	// RESULT:  <empty string>
}
