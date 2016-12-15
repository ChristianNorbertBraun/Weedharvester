package weedharvester

import (
	"fmt"
	"testing"
)

func TestAssign(t *testing.T) {
	master := master{url: "http://docker:9333"}
	assignment := master.Assign()

	if len(assignment.Fid) == 0 {
		t.Error("Returned assignment doesn't have a fileId")
	}
	fmt.Println(assignment)
}
