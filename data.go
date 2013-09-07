/*
 *	Copyright (c) 2013, Scott Cagno
 *	All rights reserved. License at
 *  	sites.google.com/site/bsdc3license
 *
 * 	data.go -- memdb database
 */

package memdb

import (
	"sync"
	"time"
)


// expired sets (for GC)
type expSets struct {
	key string
	ttl int64
}

// database
type DataBase struct {
	sets	map[string]*Set
	mu	sync.Mutex
}

// init and return a DB instance
func InitDB() *DB {
	return &DB {
		sets: map[string]*Set{},
		expired: make([]*expSets, 0),	
	}
}

// return a set if one exists, else create new one
func (self *DB) GetSet(name string) *Set {
	self.mu.Lock()
	if foundSet, ok := self.sets[name]; ok {
		self.mu.Unlock()
		return foundSet
	}
	self.mu.Unlock()
	return InitSet(name)
}

// set struct
type  Set struct {
	Name	string
	Records map[string]map[string][]string
	mu 	sync.Mutex
}

// init and return Set instance
func InitSet(name string, rate int64) *Set {
	self := &Set {
		Name : name,
		Records : map[string]map[string][]string{},
	}
	if self.Rate > 0 { go self.GC(rate) }
	unixts := fmt.Sprintf("%v", time.Now.Unix())
	self.Records["conf"]=map[string][]string{ "ts", []string{ unixts } }
	return self
}

// set garbage collector
func (self *Set) GC(rate int64) {
	self.mu.Lock()
	if record, ok := self.Records["conf"]; ok {
		if time, ok := record["ts"]; ok {
			if ts, err := strconv.ParseInt(time[0], 10, 64); != nil {
				log.Println(err)	
			}
			if (ts + rate) < time.Now().Unix() {
				record = nil
				delete(record, )
			}
		}
	}
	for key, record := self.Records {
		if ((record["ts"][0].Unix() + rate) < time.Now().Unix()) {
			record = nil
			delete(record, key)
		} else { break }
	}
	self.mu.Unlock()
	time.AfterFunc(time.Duration(rate)*time.Second, func(){ self.GC(rate) })
}

// add collection
func (self *DataBase) Add(name string) int {
	self.mu.Lock()
	self.Collections[name]=&Collection{
		name	: name,
		ts		: time.Now(),
		Records	: make([]map[string][]string, 0),
	}
	// check to make sure collection exists
	if _, ok := self.Collections[name]; ok {
		self.mu.Unlock()
		return 1
	}
	self.mu.Unlock()
	return 0
}

// delete collection
func (self *DataBase) Del(name string) int {
	self.mu.Lock()
	// make sure collection exists
	if _, ok := self.Collections[name]; ok {
		delete(self.Collections, name)
	}
	// make sure collection doesn't exist
	if _, ok := self.Collections[name]; !ok {
		self.mu.Unlock()
		// successfully deleted
		return 1
	}
	// fail
	self.mu.Unlock()
	return 0
}

// collection struct
type Collection struct {
	name 	string
	ts 		time.Time
	Records []map[string][]string
}

/*
// add record(s)
func (self *DataBase) AddRec(name string, record map[string][]string) int {
	self.mu.Lock()
	// check to make sure collection exists
	if collection, ok := self.Collections[name]; ok {
		collection = append(collection, record)
	}
	self.mu.Unlock()	
}

// add record key(s)
func (self *DataBase) AddRecKey(name string, rec map[string][]string) int {
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
*/
