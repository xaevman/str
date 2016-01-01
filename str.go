//  ---------------------------------------------------------------------------
//
//  str.go
//
//  Copyright (c) 2015, Jared Chavez. 
//  All rights reserved.
//
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.
//
//  -----------

// Package str provides helper functions when working with strings.
package str

import (
    "bytes"
)

// ListToString converts a list of strings into a single, delimiter
// delimited string.
func ListToDelimString(items []string, delimiter string) string {
    var buffer bytes.Buffer

    for i := range items {
        buffer.WriteString(items[i])
        if i < len(items) - 1 {
            buffer.WriteString(delimiter)
        }
    }

    return buffer.String()
}
