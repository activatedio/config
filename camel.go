/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 * Copyright (c) 2018 Ma_124, <github.com/Ma124>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// From https://github.com/iancoleman/strcase/blob/master/camel.go

package cs

import (
	"strings"
	"sync"
)

var uppercaseAcronym = sync.Map{}

// "ID": "id",

// ConfigureAcronym allows you to add additional words which will be considered acronyms
/*
// Leaving this here for now in case we want to include it and better manage the acronyms for cs
func csureAcronym(key, val string) {
	uppercaseAcronym.Store(key, val)
}

*/

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string { //nolint:gocyclo // acceptable complexity for readability
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	a, hasAcronym := uppercaseAcronym.Load(s)
	if hasAcronym {
		s = a.(string)
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	prevIsCap := false
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		switch {
		case capNext:
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		case i == 0:
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		case prevIsCap && vIsCap && !hasAcronym:
			v += 'a'
			v -= 'A'
		}

		prevIsCap = vIsCap

		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}

// toLowerCamel converts a string to lowerCamelCase
func toLowerCamel(s string) string {
	return toCamelInitCase(s, false)
}
