package problems

import (
	"log"
	"reflect"
	"testing"
)

func TestExamRoom_case1(t *testing.T) {
	seq := []string{"ExamRoom", "seat", "seat", "seat", "seat", "leave", "seat"}
	arg := [][]int{{10}, {}, {}, {}, {}, {4}, {}}
	exp := []int{-1, 0, 9, 4, 2, -1, 5}
	examRoomTestBase(seq, arg, exp, t, "case 1")
}

func TestExamRoom_case2(t *testing.T) {
	seq := []string{"ExamRoom", "seat", "seat", "seat", "seat", "leave", "leave", "seat"}
	arg := [][]int{{4}, {}, {}, {}, {}, {1}, {3}, {}}
	exp := []int{-1, 0, 3, 1, 2, -1, -1, 1}
	examRoomTestBase(seq, arg, exp, t, "case 2")
}

func TestExamRoom_case3(t *testing.T) {
	seq := []string{"ExamRoom", "seat", "seat", "seat", "leave", "leave", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "leave", "leave", "seat", "seat", "leave", "seat", "leave", "seat", "leave", "seat", "leave", "seat", "leave", "leave", "seat", "seat", "leave", "leave", "seat", "seat", "leave"}
	arg := [][]int{{10}, {}, {}, {}, {0}, {4}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {0}, {4}, {}, {}, {7}, {}, {3}, {}, {3}, {}, {9}, {}, {0}, {8}, {}, {}, {0}, {8}, {}, {}, {2}}
	exp := []int{-1, 0, 9, 4, -1, -1, 0, 4, 2, 6, 1, 3, 5, 7, 8, -1, -1, 0, 4, -1, 7, -1, 3, -1, 3, -1, 9, -1, -1, 0, 8, -1, -1, 0, 8, -1}
	examRoomTestBase(seq, arg, exp, t, "case 3")
}

func examRoomTestBase(seq []string, arg [][]int, exp []int, t *testing.T, caseName string) {
	result := make([]int, len(seq))
	room := Constructor(arg[0][0])
	result[0] = -1
	for i := 1; i < len(seq); i++ {
		switch seq[i] {
		case "seat":
			p := room.Seat()
			result[i] = p
		case "leave":
			room.Leave(arg[i][0])
			result[i] = -1
		}
	}
	log.Printf("expected:%v\n", exp)
	log.Printf("actual:%v\n", result)
	if reflect.DeepEqual(exp, result) {
		t.Log("855. Exam Room: ", caseName, " succeeded.")
	} else {
		t.Error("855. Exam Room: ", caseName, " failed.")
	}
}
