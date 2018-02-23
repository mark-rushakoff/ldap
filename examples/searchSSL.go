// +build ignore

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/mark-rushakoff/ldapserver"
)

var (
	LdapServer string   = "localhost"
	LdapPort   uint16   = 636
	BaseDN     string   = "dc=enterprise,dc=org"
	Filter     string   = "(cn=kirkj)"
	Attributes []string = []string{"mail"}
)

func main() {
	l, err := ldapserver.DialSSL("tcp", fmt.Sprintf("%s:%d", LdapServer, LdapPort), nil)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.String())
	}
	defer l.Close()
	// l.Debug = true

	search := ldapserver.NewSearchRequest(
		BaseDN,
		ldapserver.ScopeWholeSubtree, ldapserver.NeverDerefAliases, 0, 0, false,
		Filter,
		Attributes,
		nil)

	sr, err := l.Search(search)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.String())
		return
	}

	log.Printf("Search: %s -> num of entries = %d\n", search.Filter, len(sr.Entries))
	sr.PrettyPrint(0)
}
