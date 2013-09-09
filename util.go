/*
 *	Copyright (c) 2013, Scott Cagno
 *	All rights reserved. License at
 *  sites.google.com/site/bsdc3license
 *
 * 	util.go -- memdb utilities
 */

package memdb

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"io"
	"fmt"
)

// check if slice s, contains slice sub
func SliceContains(s, sub []string) bool {
	var ok []bool
	for i := range s {
		for j := range sub {
			if s[i] == sub[j] {
				ok = append(ok, true)
			}
		}
	}
	return len(ok) == len(sub)
}

// return random hash (6*10^49)
func Random() string {
	e := make([]byte, 16)
	rand.Read(e)
	seed := make([]byte, base64.URLEncoding.EncodedLen(len(e)))
	base64.URLEncoding.Encode(seed, e)
	h := md5.New()
	i := 3
	for i > 0 {
		io.WriteString(h, string(seed))
		i--
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
