package weedharvester

import "testing"

func TestAssign(t *testing.T) {
	master := master{url: *masterURL}
	assignment, err := master.Assign()

	if err != nil {
		t.Errorf("Error: Unable to assign fileid: %s", err)
	}

	if len(assignment.Fid) == 0 {
		t.Error("Returned assignment doesn't have a fileId")
	}
}

func TestFind(t *testing.T) {
	master := master{url: *masterURL}
	location, err := master.Find("1")
	if err != nil {
		t.Errorf("Erro: Unable to find ")
	}

	if len(location.PublicURL) == 0 {
		t.Error("Returned location which has no value")
	}
}
