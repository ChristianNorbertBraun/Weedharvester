package weedharvester

import "testing"

func TestAssign(t *testing.T) {
	master := master{url: "http://docker:9333"}
	assignment := master.Assign()

	if len(assignment.Fid) == 0 {
		t.Error("Returned assignment doesn't have a fileId")
	}
}

func TestFind(t *testing.T) {
	master := master{url: "http://docker:9333"}
	location := master.Find("1")

	if len(location.PublicURL) == 0 {
		t.Error("Returned location which has no value")
	}
}
