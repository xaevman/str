//  ---------------------------------------------------------------------------
//
//  all_test.go
//
//  Copyright (c) 2015, Jared Chavez. 
//  All rights reserved.
//
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.
//
//  -----------

package str

import (
    "fmt"
    "testing"
)

type t1Test struct {
    Vals   []string
    Result string
}

var (
    t1 = []*t1Test {
        &t1Test { []string{ "1" }, "1" },
        &t1Test { []string{ "1" ,"2" }, "1,2" },
        &t1Test { []string{ "1" ,"2", "3" }, "1,2,3" },
        &t1Test { []string{ "1", "2", "3", "4" }, "1,2,3,4" },
        &t1Test { []string{ "1", "2", "3", "4", "5" }, "1,2,3,4,5" },
    }
)

func TestListToDelimStr(t *testing.T) {
    for i := range t1 {
        res := ListToDelimString(t1[i].Vals, ",")
        if res != t1[i].Result {
            t.Fatal(fmt.Sprintf(
                "Index %d: Expected %s, got %s", 
                i, 
                t1[i].Result, 
                res,
            ))
        }
    }
}
