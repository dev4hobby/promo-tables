package test

import "testing"

func TestJoinTable(t *testing.T) {
	// need sequence
	want := true
	handler := GetServiceHandler(t)
	get := handler.JoinTableAsQueryBuilder()
	if get == want {
		t.Logf("[SYS] JoinTable() = %v, want %v", get, want)
	} else {
		t.Errorf("[SYS] JoinTable() = %v, want %v", get, want)
	}
}
