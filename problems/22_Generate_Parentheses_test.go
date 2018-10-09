package problems

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T){
	parenthesis := GenerateParenthesis(3)
	set := StringSet{make(map[string]bool)}
	set.Add([]string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	if set.Size() == len(parenthesis){
		for _, val := range parenthesis {
			if ! set.In(val) {
				t.Error("22.Generate Parentheses: case 1 failed.")
				return
			}
		}
		t.Log("22.Generate Parentheses: case 1 succeeded.")
	} else {
		t.Error("22.Generate Parentheses: case 1 failed.")
	}
}

 type StringSet struct {
	 values map[string]bool
 }

 func (set *StringSet) Add(val []string){
	for _, s := range val {
		set.values[s] = true
	}
 }

 func (set *StringSet) In(val string) bool{
	 _, ok := set.values[val]
	 return ok
 }

 func (set *StringSet) Size() int {
	 return len(set.values)
 }
