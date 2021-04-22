package setutils

type Pmfset map[string]int

func (p Pmfset) AddRunes(s string) int {
	return len(s)
}

func (p Pmfset) AddWords(s string) int {
	return len(s)
}

func (p Pmfset) AddBytes(s string) int {
	return len(s)
}

func (p Pmfset) Cardinality() int {
	return len(p)
}

func Union(p1 Pmfset, p2 Pmfset) Pmfset {
	return p1
}

