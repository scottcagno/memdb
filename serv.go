/*
 *	Copyright (c) 2013, Scott Cagno
 *	All rights reserved. License at
 *  sites.google.com/site/bsdc3license
 *
 * 	serv.go -- memdb network and server
 */

package memdb

import (
 	"encoding/json"
	"net/http"
	"strings"
	"time"
	"fmt"
	"log"
)

// server struct
type Server struct {
	http.Server
}

// initialize and return server
func InitServer() *Server {
	self := &Server{}
	self.ReadTimeout    = 10*time.Second
	self.WriteTimeout   = 10*time.Second
	self.MaxHeaderBytes = 1<<13
	self.TLSConfig      = nil
	self.TLSNextProto   = nil
	return self
}

// listener
func (self *Server) Listen(host string, muxer http.Handler) {
	self.Addr 		= host
	self.Handler 	= muxer
	err := self.ListenAndServe()
	if err != nil {
		log.Fatal("Listen failed: ", err)
	}
}

// muxer struct
type Muxer struct {
	//db 	DataBase
}

// initialize and return muxer
func InitMuxer() *Muxer {
	return &Muxer{}
}

// muxers http multiplexer/handler
func (self *Muxer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if _, ok := r.Form["debug"]; ok {
		debugMode(w, r)
		return
	}
	switch r.Method {
		case "GET":
			self.GetReq(w, r)
			break
		case "POST":
			self.PostReq(w, r)
			break
		case "PUT":
			self.PutReq(w, r)
			break
		case "DELETE":
			self.DelReq(w, r)
			break
		case "OPTIONS":
			self.OptReq(w, r)
			break
		default:
			fmt.Fprintf(w, "%s is an unknown request!", r.Method)
			break
	}
	return
}

// handle get request
func (self *Muxer) GetReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received GET request\n")
}

// handle post request
func (self *Muxer) PostReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received POST request\n")
}

// handle put request
func (self *Muxer) PutReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received PUT request\n")
}

// handle delete request
func (self *Muxer) DelReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received DELETE request\n")
}

// handle options request
func (self *Muxer) OptReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received OPTIONS request\n")
}

// debug mode helper
func debugMode(w http.ResponseWriter, r *http.Request) {
	clnt, meth, host := r.RemoteAddr, r.Method, r.Host
	coll := strings.Split(r.RequestURI[1:], "?")[0]
	delete(r.Form, "debug")
	b, _ := json.Marshal(r.Form)
	resp := []string{"clnt: %s","meth: %s","host: %s","coll: %s","data: %s"}
	fmt.Fprintf(w, strings.Join(resp, "\n"), clnt, meth, host, coll, b)
}
