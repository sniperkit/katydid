//  Copyright 2016 Walter Schulze
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

package auto_test

import (
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/auto"
	"github.com/katydid/katydid/relapse/interp"
	"testing"
)

func test(t *testing.T, g *ast.Grammar, p parser.Interface, expected bool, desc string) {
	if interp.HasLeftRecursion(g) {
		t.Skipf("convert was not designed to handle recursion")
	}
	a := auto.Compile(g)
	match := auto.Interpret(a, p)
	if match != expected {
		t.Fatalf("Expected %v on given \n%s\n on \n%s", expected, g.String(), desc)
	}
}