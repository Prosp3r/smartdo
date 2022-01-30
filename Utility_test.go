package main

import (
	"testing"

	"github.com/Prosp3r/smartdo/utility"
)

var SUsers = []struct {
	username    string
	password    string
	expresponse string
	gotresponse string
}{
	{"john", "doe123", "passed", ""},
	{"pros", "", "failed", ""},
	{"", "frigate123,", "failed", ""},
	{"", "", "failed", ""},
	{"", "frigate123,", "failed", ""},
}

func TestCreateCryptoWallet(t *testing.T) {
	//test for length, test for fake address,
	for _, v := range SUsers {
		res, _ := utility.CreateCryptoWallet(v.username, v.password)
		if res != nil {
			v.gotresponse = "passed"
		}
		v.gotresponse = "failed"
	}
	for _, v := range SUsers {
		if v.expresponse != v.gotresponse {
			t.Logf("Found %v on Username: '%v', Password: '%v' - - expected %v  --> %s\n", v.gotresponse, v.username, v.password, v.expresponse, failed)
		}
		t.Logf("Found %v on Username: '%v', Password: '%v' - - expected %v  --> %s\n", v.gotresponse, v.username, v.password, v.expresponse, succeed)
	}

}

var PreAddresses = []struct {
	Username     string
	Password     string
	ExpectedHash string
	GotHash      string
}{
	{"efemena", "password", "0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d", ""},
	{"omovie", "akomeno123,", "0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d", ""},
	{"ovovie", "akomeno123,", "", ""}, //non existent account
}

func TestGetUserAddress(t *testing.T) {
	for _, v := range PreAddresses {
		gu, _ := utility.GetUserAddress(v.Username, v.Password)
		if gu != nil {
			v.GotHash = gu.Hex()
		}
	}

	for _, v := range PreAddresses {
		if v.GotHash != v.ExpectedHash {
			t.Logf("Found %v on Username: '%v', Password: '%v' - - expected %v  --> %s\n", v.GotHash, v.Username, v.Password, v.ExpectedHash, failed)
		}
		t.Logf("Found %v on Username: '%v', Password: '%v' - - expected %v  --> %s\n", v.GotHash, v.Username, v.Password, v.ExpectedHash, succeed)
	}
}
