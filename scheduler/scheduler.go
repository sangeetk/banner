package banner

import (
	"errors"
	"fmt"
	"sync"
)

// Storage for timeslots ??

// Timeslot is valid for time t when  T1 <= t < T2
type Timeslot struct {
	T1      int64
	T2      int64
	Banners []*Banner
	Next    *Timeslot
}

func (t *Timeslot) String() string {
	str := fmt.Sprintf("{%v - %v, [ ", t.T1, t.T2)
	for _, b := range t.Banners {
		str += b.ID + " "
	}
	str += "]}"
	return str
}

type Scheduler struct {
	Lock sync.Mutex
	Head *Timeslot
}

func New() *Scheduler {
	return &Scheduler{}
}

// Add an item to the scheduler
func (s *Scheduler) Schedule(b *Banner) error {
	if b.ActiveAt >= b.ExpireAt {
		return errors.New("Invalid schedule")
	}
	//	if b.ExpireAt <= time.Now().Unix() {
	//		return errors.New("Expired schedule")
	//	}

	if s.Head == nil {
		// Case 1: Empty schedule
		// Time : +----+----+----+----+----+----+----+----+
		// Head -> nil
		// New ->     AAAAAAAAAA
		// Result:
		// Head ->    AAAAAAAAAA -> nil
		s.Head = &Timeslot{b.ActiveAt, b.ExpireAt, []*Banner{b}, nil}
		fmt.Println("Case 1:")
		return nil
	}

	// Find the right position
	var prev *Timeslot = nil
	for p := s.Head; p != nil; prev, p = p, p.Next {

		// Case 2: Schedule in the empty timeslot
		// Time : +----+----+----+----+----+----+----+----+
		// Head ->               ########## -> ##### -> nil
		// New ->  AAAAAAAAAA
		// Result:
		// Head -> AAAAAAAAAA -> ########## -> ##### -> nil
		if b.ExpireAt <= p.T1 {
			q := &Timeslot{b.ActiveAt, b.ExpireAt, []*Banner{b}, p}
			if s.Head == p {
				s.Head = q
			} else {
				prev.Next = q
			}
			fmt.Println("Case 2:")
			return nil
		}

		// Case 3: Schedule in the end
		// Time : +----+----+----+----+----+----+----+----+
		// Head -> ########## -> ##### -> nil
		// New ->                           AAAAAAAAAA
		// Result:
		// Head -> ########## -> ##### ->   AAAAAAAAAA -> nil
		if p.T2 <= b.ActiveAt && p.Next == nil {
			q := &Timeslot{b.ActiveAt, b.ExpireAt, []*Banner{b}, nil}
			p.Next = q
			fmt.Println("Case 3:")
			return nil
		}

		// Case 4: Left Overlap 1
		// Time : +----+----+----+----+----+----+----+----+
		// Head ->    ########## -> ##### -> nil
		// New ->  AAAAAAAAA
		// Result:
		// Head -> AAAAAAAAA#### -> ##### -> nil
		// Head -> AAAAAAAAA -> #### -> ##### -> nil
		if b.ActiveAt <= p.T1 && b.ExpireAt > p.T1 {
			// Split the current timeslot into two slots
			q := &Timeslot{b.ActiveAt, b.ExpireAt, []*Banner{b}, p}
			// Shrink the timeslot
			p.T1 = q.T2
			if prev == nil {
				s.Head = q
			} else {
				prev.Next = q
			}
			fmt.Println("Case 4:")
			return nil
		}

		// Case 5: Right Overlap
		// Time : +----+----+----+----+----+----+----+----+
		// Head -> ########## -> ##### -> nil
		// New ->     AAAAAAAAA
		// Result:
		// Head -> ##########AA -> ##### -> nil
		if b.ActiveAt >= p.T1 && b.ExpireAt > p.T1 {
			// Split the current timeslot into two slots
			q := &Timeslot{b.ActiveAt, b.ExpireAt, []*Banner{b}, p}
			// Shrink the timeslot
			p.T1 = q.T2
			if prev == nil {
				s.Head = q
			} else {
				prev.Next = q
			}
			fmt.Println("Case 4:")
			return nil
		}

		// Case 7:
		// Time : +----+----+----+----+----+----+----+----+
		// Head -> ########## -> ##### -> nil
		// New ->                           ##########
		// Result:

		// Time : +----+----+----+----+----+----+----+----+
		// Head -> ########## -> ##### -> nil
		// New ->                           ##########
		// Result:

		// Time : +----+----+----+----+----+----+----+----+
		// Head -> ########## -> ##### -> nil
		// New ->                           ##########
		// Result:

		// Time : +----+----+----+----+----+----+----+----+
		// Head -> ########## -> ##### -> nil
		// New ->                           ##########
		// Result:

		// Time : +----+----+----+----+----+----+----+----+
		// Head -> ########## -> ##### -> nil
		// New ->                           ##########
		// Result:

	}
	return nil
}

// Get an active item
func (s *Scheduler) Get() (*Banner, error) {
	return nil, nil
}

// Get an active item
func (s *Scheduler) GetByTime(t int64) (*Banner, error) {
	for p := s.Head; p != nil; p = p.Next {
		if t >= p.T1 && t < p.T2 {
			return p.Banners[0], nil
		}
	}
	return nil, errors.New("Not Found")
}

// Cleanup unused memory
func (s *Scheduler) Cleanup() {

}

// Debug scheduler
func (s *Scheduler) Debug() {
	// Display linked list
	fmt.Print("Head -> ")
	for t := s.Head; t != nil; t = t.Next {
		for i := t.T1; i < t.T2; i++ {
			fmt.Print(t.Banners[0].ID)
		}
		fmt.Print(" -> ")
	}
	fmt.Println("nil")

	// Display timeline
	var start, end int64
	if s.Head != nil {
		start = s.Head.T1
	}
	fmt.Printf("Timeline: (%02d)", start)
	for t := s.Head; t != nil; t = t.Next {
		end = t.T2
	}
	for i := start; i <= end; i++ {
		if i%10 == 0 {
			fmt.Print("|")
		} else if i%5 == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Printf("(%02d)\n", end)

	// Display scheduled banners
	fmt.Printf("Scheduling:   ")
	for i := start; i <= end; i++ {
		b, err := s.GetByTime(i)
		if err != nil {
			fmt.Print(" ")
		} else {
			fmt.Print(b.ID)
		}
	}
	fmt.Println()
	fmt.Println()

}
