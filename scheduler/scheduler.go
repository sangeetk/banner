package scheduler

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/sangeetk/banner"
)

// Storage for timeslots ??

// Timeslot is valid for time t when  T1 <= t < T2
type Timeslot struct {
	Lock    bool
	T1      int64
	T2      int64
	Banners []banner.Banner
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
func (s *Scheduler) Schedule(b banner.Banner) error {

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
		s.Head = &Timeslot{false, b.ActiveAt, b.ExpireAt, []banner.Banner{b}, nil}
		return nil
	}

	// Find the right position
	var prev *Timeslot = nil
	for p := s.Head; p != nil; prev, p = p, p.Next {

		if p.Lock {
			continue
		}

		// Case 2: Schedule in the empty timeslot
		// Time : +----+----+----+----+----+----+----+----+
		// Head ->               ########## -> ##### -> nil
		// New ->  AAAAAAAAAA
		// Result:
		// Head -> AAAAAAAAAA -> ########## -> ##### -> nil
		if b.ExpireAt <= p.T1 {
			fmt.Println("Case 2: p,b = ", p, b)
			q := &Timeslot{false, b.ActiveAt, b.ExpireAt, []banner.Banner{b}, p}
			if s.Head == p {
				s.Head = q
			} else {
				prev.Next = q
			}
			return nil
		}

		// Case 3: Schedule in the end
		// Time : +----+----+----+----+----+----+----+----+
		// Head -> ########## -> ##### -> nil
		// New ->                           AAAAAAAAAA
		if p.T2 <= b.ActiveAt && p.Next == nil {
			fmt.Println("Case 3: p,b = ", p, b)
			q := &Timeslot{false, b.ActiveAt, b.ExpireAt, []banner.Banner{b}, nil}
			p.Next = q
			return nil
		}

		// Case 4: Overlap 1 - Split new node
		// Time : +----+----+----+----+----+----+----+----+
		// Head ->    ########## -> ##### -> nil
		// New ->  AAAAAAAAA
		if b.ActiveAt < p.T1 {
			fmt.Println("Case 4: p,b = ", p, b)
			b1 := banner.Banner{b.ID, b.ActiveAt, b.ExpireAt, b.Image}
			b2 := banner.Banner{b.ID, p.T1, b.ExpireAt, b.Image}

			q := &Timeslot{false, b.ActiveAt, p.T1, []banner.Banner{b1}, p}

			if prev != nil {
				prev.Next = q
			} else {
				s.Head = q
			}

			q.Lock = true
			err := s.Schedule(b2)
			q.Lock = false
			return err
		}

		// Case 5: Overlap 2 - Split current node
		// Time : +----+----+----+----+----+----+----+----+
		// Head -> ########## -> ##### -> nil
		// New ->     AAAAAAAAA
		if p.T1 < b.ActiveAt && p.T2 > b.ActiveAt {
			fmt.Println("Case 5: p,b = ", p, b)
			q := &Timeslot{false, p.T1, b.ActiveAt, p.Banners, p}
			p.T1 = b.ActiveAt

			if prev != nil {
				prev.Next = q
			} else {
				s.Head = q
			}
			q.Lock = true
			err := s.Schedule(b)
			q.Lock = false
			return err
		}

		// Case 6: Start at same time
		if p.T1 == b.ActiveAt {
			// Case 6.1: Shorter new banner
			// Time : +----+----+----+----+----+----+----+----+
			// Head -> ########## -> ##### -> nil
			// New ->  AAAAAAA
			if b.ExpireAt < p.T2 {
				fmt.Println("Case 6.1: p,b = ", p, b)
				q := &Timeslot{false, b.ExpireAt, p.T2, p.Banners, p.Next}
				p.T2 = b.ExpireAt
				p.Banners = append(p.Banners, b)
				sort.Sort(banner.ByExpiry(p.Banners))
				p.Next = q
				return nil
			}

			// Case 6.2: Longer new banner
			// Time : +----+----+----+----+----+----+----+----+
			// Head -> ########## -> ##### -> nil
			// New ->  AAAAAAAAAAAA
			if p.T2 < b.ExpireAt {
				fmt.Println("Case 6.2: p,b = ", p, b)
				p.Banners = append(p.Banners, b)
				sort.Sort(banner.ByExpiry(p.Banners))

				if p.Next == nil {
					q := &Timeslot{false, p.T2, b.ExpireAt, []banner.Banner{b}, nil}
					p.Next = q
					return nil
				}

				p.Lock = true
				b.ActiveAt = p.T2
				err := s.Schedule(b)
				p.Lock = false
				return err
			}

			// case 6.3: Equal new banner
			// Time : +----+----+----+----+----+----+----+----+
			// Head -> ########## -> ##### -> nil
			// New ->  AAAAAAAAAA
			if p.T2 == b.ExpireAt {
				fmt.Println("Case 6.3: p,b = ", p, b)
				p.Banners = append(p.Banners, b)
				sort.Sort(banner.ByExpiry(p.Banners))
				return nil
			}

		}

	}
	return nil
}

// Get an active item
func (s *Scheduler) Get() (banner.Banner, error) {
	return banner.Banner{}, nil
}

// Get an active item
func (s *Scheduler) GetByTime(t int64) (banner.Banner, error) {
	for p := s.Head; p != nil; p = p.Next {
		if t >= p.T1 && t < p.T2 {
			return p.Banners[0], nil
		}
	}
	return banner.Banner{}, errors.New("Not Found")
}

// Cleanup unused memory
func (s *Scheduler) Cleanup() {

}

// Debug scheduler
func (s *Scheduler) Debug() {
	// Display linked list
	fmt.Print("Head -> ")
	for t := s.Head; t != nil; t = t.Next {
		// for i := t.T1; i < t.T2; i++ {
		//	fmt.Print(t.Banners[0].ID)
		// }
		fmt.Printf("{'%v',%v,%v}", t.Banners[0].ID, t.T1, t.T2)
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
