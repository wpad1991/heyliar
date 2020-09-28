package heyliar

// Stackstring is Stack for string
type Stackstring struct {
	items []string
}

// Stackfloat64 is Stack for float64
type Stackfloat64 struct {
	items []float64
}

// Queuestring is Queue for string
type Queuestring struct {
	items []string
}

// Push is Push
func (s *Stackstring) Push(item string) {
	s.items = append(s.items, item)
}

// Pop is Pop
func (s *Stackstring) Pop() string {

	if s.items == nil {
		panic("Stack Item is nil")
	}

	sSize := len(s.items)

	if sSize == 0 {
		panic("Stack Item Is Empty")
	}

	rs := s.items[sSize-1]
	s.items = s.items[0 : len(s.items)-1]

	return rs
}

// Top is Top
func (s *Stackstring) Top() string {
	if s.items == nil {
		panic("Stack Item is nil")
	}

	sSize := len(s.items)

	if sSize == 0 {
		panic("Stack Item Is Empty")
	}

	rs := s.items[sSize-1]

	return rs
}

// GetSize is GetSize
func (s *Stackstring) GetSize() int {

	if s.items == nil {
		return 0
	}

	sSize := len(s.items)

	return sSize
}

// Push is Push
func (s *Stackfloat64) Push(item float64) {
	s.items = append(s.items, item)
}

// Pop is Pop
func (s *Stackfloat64) Pop() float64 {

	if s.items == nil {
		panic("Stack Item is nil")
	}

	sSize := len(s.items)

	if sSize == 0 {
		panic("Stack Item Is Empty")
	}

	rs := s.items[sSize-1]
	s.items = s.items[0 : len(s.items)-1]

	return rs
}

// Top is Top
func (s *Stackfloat64) Top() float64 {
	if s.items == nil {
		panic("Stack Item is nil")
	}

	sSize := len(s.items)

	if sSize == 0 {
		panic("Stack Item Is Empty")
	}

	rs := s.items[sSize-1]

	return rs
}

// GetSize is GetSize
func (s *Stackfloat64) GetSize() int {

	if s.items == nil {
		return 0
	}

	sSize := len(s.items)

	return sSize
}

// Push is Push
func (q *Queuestring) Push(item string) {
	q.items = append(q.items, item)
}

// Pop is Pop
func (q *Queuestring) Pop() string {

	if q.items == nil {
		panic("Queue Item is nil")
	}

	qSize := len(q.items)

	if qSize == 0 {
		panic("Queue Item Is Empty")
	}

	rs := q.items[0]
	q.items = q.items[1:len(q.items)]

	return rs
}

// Top is Top
func (q *Queuestring) Top() string {
	if q.items == nil {
		panic("Queue Item is nil")
	}

	qSize := len(q.items)

	if qSize == 0 {
		panic("Queue Item Is Empty")
	}

	rs := q.items[0]

	return rs
}

// GetSize is GetSize
func (q *Queuestring) GetSize() int {

	if q.items == nil {
		return 0
	}

	qSize := len(q.items)

	return qSize
}
