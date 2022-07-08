package utils

import "testing"

func TestGetAccountsFromCrossbell(t *testing.T) {
	if accounts, err := GetAccountsFromCrossbell("19"); err != nil {
		t.Fatal(err)
	} else {
		t.Log(accounts)
	}
}
