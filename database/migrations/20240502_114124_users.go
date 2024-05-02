package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Users_20240502_114124 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20240502_114124{}
	m.Created = "20240502_114124"

	migration.Register("Users_20240502_114124", m)
}

// Run the migrations
func (m *Users_20240502_114124) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Users_20240502_114124) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
