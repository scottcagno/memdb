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

// database struct yeah
type DataBase struct {
	Collections map[string]*Collection
	mu			sync.Mutex
}

// initialize database and return
func InitDataBase(gcrate int64) *DataBase {
	self := &DataBase{
		Collections: make(map[string]*Collection, 0),
	}
	if gcrate > 0 {
		go func(){ self.GC(gcrate) }()
	}
	return self
}

// garbage collector
func (self *DataBase) GC(rate int64) {
	self.mu.Lock()
	for name, collection := self.Collections {
		if ((collection.ts.Unix() + rate) < time.Now().Unix()) {
			delete(self.Collections, name)
		} else { break }
	}
	self.mu.Unlock()
	time.AfterFunc(time.Duration(rate)*time.Second, func(){ self.GC(rate) })
}

// add collection
func (self *DataBase) AddCol(name string) int {
	self.mu.Lock()
	self.Collections[name]=&Collection{
		name: name,
		Records: make([]map[string][]string, 0),
	}
	if _, ok := self.Collections[name]; ok {
		self.mu.Unlock()
		return 1
	}
	self.mu.Unlock()
	return 0
}

// add record(s)
func (self *DataBase) AddRec(coll string, rec map[string][]string) int {
	return 0
}

// add record key(s)
func (self *DataBase) AddRecKey(coll string, rec map[string][]string) int {
	return 0
}

// delete collection
func (self *DataBase) DelCol(coll string) int {
	return 0
}

// delete record(s)
func (self *DataBase) DelRec(coll string, rec map[string][]string) int {
	return 0
}

// delete record key(s)
func (self *DataBase) DelRecKey(coll string, rec map[string][]string) int {
	return 0
}

// update collection
func (self *DataBase) UpdCol(coll string) int {
	return 0
}

// update record(s)
func (self *DataBase) UpdRec(coll string, rec map[string][]string) int {
	return 0
}

// update record key(s)
func (self *DataBase) UpdRecKey(coll string, rec map[string][]string) int {
	return 0
}

// get collection
func (self *DataBase) GetCol(coll string) int {
	return 0
}

// get record(s)
func (self *DataBase) GetRec(coll string, rec map[string][]string) int {
	return 0
}

// get record key(s)
func (self *DataBase) GetRecKey(coll string, rec map[string][]string) int {
	return 0
}

// collection struct
type Collection struct {
	name 	string
	ts 		time.Time
	Records []map[string][]string
}

// HELPER -- add n number of records
func (self *DataBase) AddRandom(n int64) {
	return
}

// HELPER -- remove n number of records
func (self *DataBase) DelRandom(n int64) {
	return
}

// HELPER -- show memory stats
func (self *DataBase) ShowStats() {
	return
}
