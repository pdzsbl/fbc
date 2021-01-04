package sw

import (
	"fmt"
	"testing"
)

func TestReadLineTxt(t *testing.T) {
	ntlist , _ := ReadLineTxt("/Users/jx/nt-config/nt.txt")

	for i:=0;i<len(ntlist);i++  {
		fmt.Println(ntlist[i])
	}
}
