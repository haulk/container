/*
 Package set will contain an implementation of a Set data structure. 
 It's little more than a map.

*/

package set


type Element interface {
	// Fingerprint is a unique constant associated to each individual
	// element.
	Fingerprint() int
}

type Set interface {
	// Contains returns true if the Element is in the set, false otherwise.
	Contains(Element) bool
	// ContainsFP returns true if the Element is in the set.
	ContainsFP(int) bool

	Delete(Element)
	DeleteFP(int)

	Add(Element) 
}


// set implements the Set interface.
type set map[int]Element

func (s set) Contains(e Element) bool {
	fp := e.Fingerprint()
	return s.ContainsFP(fp)	
}

func (s set) ContainsFP(i int) bool {
	_, out := s[i]
	return out
}

func (s set) Delete(e Element) {
	s.DeleteFP(e.Fingerprint())
}

func (s set) DeleteFP(i int) {
	delete(s, i)
}

func (s set) Add(e Element) {
	s[e.Fingerprint()] = e
}





