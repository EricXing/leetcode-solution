package problems

import "sort"

// In an exam room, there are N seats in a single row, numbered 0, 1, 2, ..., N-1.
//
// When a student enters the room, they must sit in the seat that maximizes t
// he distance to the closest person.  If there are multiple such seats,
// they sit in the seat with the lowest number.  (Also, if no one
// is in the room, then the student sits at seat number 0.)
//
// Return a class ExamRoom(int N) that exposes two functions: E
// xamRoom.seat() returning an int representing what seat the student sat
// in, and ExamRoom.leave(int p) representing that the student in seat
// number p now leaves the room.  It is guaranteed that any calls to
// ExamRoom.leave(p) have a student sitting in seat p.
//
// Example 1:
//
//	Input: ["ExamRoom","seat","seat","seat","seat","leave","seat"], [[10],[],[],[],[],[4],[]]
//	Output: [null,0,9,4,2,null,5]
//	Explanation:
//	ExamRoom(10) -> null
//	seat() -> 0, no one is in the room, then the student sits at seat number 0.
//	seat() -> 9, the student sits at the last seat number 9.
//	seat() -> 4, the student sits at the last seat number 4.
//	seat() -> 2, the student sits at the last seat number 2.
//	leave(4) -> null
//	seat() -> 5, the student sits at the last seat number 5.
//​​​​​​​
// Note:
//
//	1 <= N <= 10^9
//	ExamRoom.seat() and ExamRoom.leave() will be called at most 10^4 times across all test cases.
//	Calls to ExamRoom.leave(p) are guaranteed to have a student currently sitting in seat number p.
type ExamRoom struct {
	size     int
	occupied int
	seats    map[int]bool
}

func Constructor(N int) ExamRoom {
	return ExamRoom{size: N, occupied: 0, seats: make(map[int]bool)}
}

func (room *ExamRoom) Seat() int {
	if room.occupied == 0 {
		room.seat(0)
		return 0
	}
	seats := make([]int, 0, len(room.seats))
	for s := range room.seats {
		seats = append(seats, s)
	}
	sort.Sort(sort.IntSlice(seats))
	// loop over the list and find the pos between every two element,
	// find the farthest
	space := 0
	pos := -1
	if seats[0] > 0 {
		space = seats[0]
		pos = 0
	}
	for i := 0; i < len(seats); i++ {
		if i == len(seats)-1 {
			if seats[i] != room.size-1 {
				sp := room.size - 1 - seats[i]
				if sp > space {
					space = sp
					pos = room.size - 1
				}
			}
		} else {
			sp := (seats[i+1] - seats[i]) / 2
			if sp > space {
				space = sp
				pos = seats[i] + sp
			}
		}
	}

	room.seat(pos)
	return pos
}

func (room *ExamRoom) seat(pos int) {
	room.seats[pos] = true
	room.occupied += 1
}

func (room *ExamRoom) Leave(p int) {
	delete(room.seats, p)
	room.occupied -= 1
}
