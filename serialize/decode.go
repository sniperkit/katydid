//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package serialize

import "fmt"

type Decoder interface {
	Double() (float64, error)
	Int() (int64, error)
	Uint() (uint64, error)
	Bool() (bool, error)
	String() (string, error)
	Bytes() ([]byte, error)
}

type errValue struct{}

type Parser interface {
	Next() error
	IsLeaf() bool
	Up()
	Down()
	Decoder
}

func Sprint(value Decoder) string {
	return fmt.Sprintf("%#v", getValue(value))
}

func getValue(value Decoder) interface{} {
	var v interface{}
	var err error
	v, err = value.Bool()
	if err == nil {
		return v
	}
	v, err = value.Bytes()
	if err == nil {
		return v
	}
	v, err = value.String()
	if err == nil {
		return v
	}
	v, err = value.Int()
	if err == nil {
		return v
	}
	v, err = value.Uint()
	if err == nil {
		return v
	}
	v, err = value.Double()
	if err == nil {
		return v
	}
	return nil
}