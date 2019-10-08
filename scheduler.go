package banner

import (
	"errors"
	"sort"
	"time"
)

// Timeslot is valid for time t when  T1 <= t < T2
type Timeslot struct {
	Lock    bool
	T1      int64
	T2      int64
	Banners []*Banner
	Next    *Timeslot
}

// Scheduler is a linked list of timeslots
type Scheduler struct {
	Head *Timeslot
}

// Schedule a banner
func (s *Scheduler) Schedule(b *Banner) error {

	if b.ActiveAt >= b.ExpireAt {
		return errors.New(ErrorInvalidSchedule)
	}

	if b.ExpireAt <= time.Now().Unix() {
		return errors.New(ErrorExpired)
	}

	if s.Head == nil {
		// Case 1: Empty schedule
		// Time : +----+----+----+----+----+----+----+----+
		// Head -> nil
		// New ->     AAAAAAAAAA
		// Result:
		// Head ->    AAAAAAAAAA -> nil
		s.Head = &Timeslot{false, b.ActiveAt, b.ExpireAt, []*Banner{b}, nil}
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
			// fmt.Println("Case 2: p,b = ", p, b)
			q := &Timeslot{false, b.ActiveAt, b.ExpireAt, []*Banner{b}, p}
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
			// fmt.Println("Case 3: p,b = ", p, b)
			q := &Timeslot{false, b.ActiveAt, b.ExpireAt, []*Banner{b}, nil}
			p.Next = q
			return nil
		}

		// Case 4: Overlap 1 - Split new node
		// Time : +----+----+----+----+----+----+----+----+
		// Head ->    ########## -> ##### -> nil
		// New ->  AAAAAAAAA
		if b.ActiveAt < p.T1 {
			// fmt.Println("Case 4: p,b = ", p, b)
			b1 := Banner{b.ID, b.ActiveAt, b.ExpireAt, b.Image}
			b2 := Banner{b.ID, p.T1, b.ExpireAt, b.Image}

			q := &Timeslot{false, b.ActiveAt, p.T1, []*Banner{b1}, p}

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
			// fmt.Println("Case 5: p,b = ", p, b)

			banners := make([]*Banner, len(p.Banners))
			copy(banners, p.Banners)

			q := &Timeslot{false, p.T1, b.ActiveAt, banners, p}
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
				// fmt.Println("Case 6.1: p,b = ", p, b)
				banners := make([]*Banner, len(p.Banners))
				copy(banners, p.Banners)

				q := &Timeslot{false, b.ExpireAt, p.T2, banners, p.Next}
				p.T2 = b.ExpireAt
				p.Banners = append(p.Banners, b)
				sort.Sort(SortByExpiry(p.Banners))
				p.Next = q
				return nil
			}

			// Case 6.2: Longer new banner
			// Time : +----+----+----+----+----+----+----+----+
			// Head -> ########## -> ##### -> nil
			// New ->  AAAAAAAAAAAA
			if p.T2 < b.ExpireAt {
				// fmt.Println("Case 6.2: p,b = ", p, b)
				p.Banners = append(p.Banners, b)
				sort.Sort(SortByExpiry(p.Banners))

				if p.Next == nil {
					q := &Timeslot{false, p.T2, b.ExpireAt, []*Banner{b}, nil}
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
				// fmt.Println("Case 6.3: p,b = ", p, b)
				p.Banners = append(p.Banners, b)
				sort.Sort(SortByExpiry(p.Banners))
				return nil
			}

		}

	}
	return nil
}

// Get an active banner
func (s *Scheduler) Get() (*Banner, error) {
	return GetByTime(time.Now().Unix())
}

// GetAll active banners
func (s *Scheduler) GetAll() ([]*Banner, error) {
	return GetAllByTime(time.Now().Unix())
}

// GetByTime returns a banner, scheduled in future
func (s *Scheduler) GetByTime(t int64) (*Banner, error) {
	for p := s.Head; p != nil; p = p.Next {
		if t >= p.T1 && t < p.T2 {
			if len(p.Banners) == 0 {
				return nil, errors.New(ErrorNotFound)
			}
			return p.Banners[0], nil
		}
	}
	return nil, errors.New(ErrorNotFound)
}

// GetAllByTime returns a banner, scheduled in future
func (s *Scheduler) GetAllByTime(t int64) ([]*Banner, error) {
	for p := s.Head; p != nil; p = p.Next {
		if t >= p.T1 && t < p.T2 {
			if len(p.Banners) == 0 {
				return []*Banners{}, errors.New(ErrorNotFound)
			}
			return p.Banners, nil
		}
	}
	return []*Banners{}, errors.New(ErrorNotFound)
}

// Unschedule a banner
func (s *Scheduler) Unschedule(id string) (*Banner, error) {
	for p := s.Head; p != nil; p = p.Next {
		for i, b := range p.Banners {
			if b.ID == id {
				// Remove the banner
				p.Banners = append(p.Banners[:i], p.Banners[i+1:]...)
				return b, nil
			}
		}
	}
	return nil, errors.New(ErrorNotFound)
}
