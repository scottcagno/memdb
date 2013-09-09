/*
 *	Copyright (c) 2013, Scott Cagno
 *	All rights reserved. License at
 *  sites.google.com/site/bsdc3license
 *
 * 	data.go -- memdb database
 */

package memdb

import (
	"sync"
	"time"
)

// database structure
type DataBase struct {
	sets	map[string]*Set
	mu		sync.Mutex
}

// initialize and return a pointer to a new DataBase struct
func InitDB() *DataBase {
	self := &DataBase {
		sets: 		map[string]*Set{},
	}
	return self
}

// return a set if one exists, else create new one
func (self *DataBase) GetSet(name string, rate int64) *Set {
	self.mu.Lock()
	if foundSet, ok := self.sets[name]; ok {
		self.mu.Unlock()
		return foundSet
	}
	self.mu.Unlock()
	return InitSet(name, rate)
}

// set structure
type Set struct {
	Records map[string]map[string][]string
	exprecs map[string]int64
	mu 		sync.Mutex								
}

// init and return Set instance
func InitSet(name string, rate int64) *Set {
	self := &Set {
		Records : map[string]map[string][]string{},
		exprecs : map[string]int64{},
	}
	if rate > 0 {
		go self.RunGC(rate)
	}
	return self
}

// run garbage collector (for records)
func (self *Set) RunGC(rate int64) {
	if rate > 0 {
		self.GC()
		time.AfterFunc(time.Duration(rate)*time.Second, func() { self.RunGC(rate) })
	}
}

// garbage collector (for records)
func (self *Set) GC() {
	self.mu.Lock()
	for recid, ttl := range self.exprecs {
		if ttl < time.Now().Unix() {
			delete(self.Records, recid)
			delete(self.exprecs, recid)
		}
	}
	self.mu.Unlock()
}

// add record
func (self *Set) Add(m map[string][]string) int {
	self.mu.Lock()
	self.Records[Random()]=m
	self.mu.Unlock()
	return 1
}

// del record
func (self *Set) Del(m map[string][]string) int {
	self.mu.Lock()
	match := map[string]int{}
	for k, v := range m {
		for id, record := range self.Records {
			if SliceContains(record[k], v) {
				match[id]++
			}
		}	
	}
	n := 0
	for uid, mlen := range match {
		if mlen == len(m) {
			delete(self.Records, uid)
			n++
		}
	}
	self.mu.Unlock()
	return n
}
