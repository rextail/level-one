package t12

import (
	"fmt"
	"testing"
)

func TestSeqSet_AddStrings(t *testing.T) {
	set := New()

	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	
	if !set.AddStrings(seq...) {
		t.Errorf("sequence must be added to the set")
	}

	if set.AddStrings(seq...) { //не идемпотентно, результат должен быть другим
		t.Errorf("sequence must NOT be added to the set")
	}
	fmt.Println(set)
}
