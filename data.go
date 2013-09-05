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

// database struct
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

// add new collection
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

// add new record to collection
func (self **DataBase) AddRec(coll string, rec map[string][]string) int {}

func (self *DataBase) AddRecKey(coll string, rec map[string][]string) int {}

// delete methods
func (self *DataBase) DelCol(coll string) int {}

func (self *DataBase) DelRec(coll string, rec map[string][]string) int {}

func (self *DataBase) DelRecKey(coll string, rec map[string][]string) int {}

// update methods
func (self *DataBase) UpdCol(coll string) int {}

func (self *DataBase) UpdRec(coll string, rec map[string][]string) int {}

func (self *DataBase)	UpdRecKey(coll string, rec map[string][]string) int {}

// get methods
func (self *DataBase) GetCol(coll string) int {}

func (self *DataBase) GetRec(coll string, rec map[string][]string) int {}

func (self *DataBase) GetRecKey(coll string, rec map[string][]string) int {}

// collection struct
type Collection struct {
	name 	string
	ts 		time.Time
	Records []map[string][]string
}



// adding a new record
func (self *DataBase) Add(coll, key [32]byte, val [][]byte) int {
	return 0
}

// delete a record
func (self *DataBase) Del(coll, key [32]byte, val [][]byte) int {
	// db.Del("customer", nil, nil) 				-> delete collection
	// db.Del("customer", "name", "samuel jackson")	-> delete record

	// DELETE -> /orders/?name=scott&quantity=lt&quantity=500
	//	func	   coll			key, val pairs <-- ie multiple

	// [0]Record{
	//		"email" : ["greg@awesome.com", "scott@lame.com"], PUT /orders/?email=scott@lame.com
	//		"first" : ["Greg"], 
	//		"last"  : ["Pechiro"],
	//	}
	
	return 0
}

// HELPER -- add n number of records
func (self *DataBase) AddRandom(n int64) {}

// HELPER -- remove n number of records
func (self *DataBase) DelRandom(n int64) {}

// HELPER -- show memory stats
func (self *DataBase) ShowStats() {}
