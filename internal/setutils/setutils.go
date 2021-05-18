package setutils

type Pmfset map[string]int

func (p Pmfset) Stringcount(s string) int {
	_, exists := p[s]
	if !exists {
		return 0
	}
	return p[s]
}

func (p Pmfset) Stringexists(s string) bool {
	c := p.Stringcount(s)
	b := c != 0
	return b
}

func (p1 Pmfset) Subtract(p2 Pmfset) Pmfset {
	pr := make(map[string]int)
	for k := range p1 {
		pr[k] = p1.Stringcount(k) - p2.Stringcount(k)
	}
	return pr
}

func (p1 Pmfset) Intersection(p2 Pmfset) Pmfset {
	pr := make(map[string]int)
	for k := range p1 {
		if p2.Stringexists(k) {
			pr[k] = p1[k] + p2[k]
		}
	}
	return pr
}

func (p1 Pmfset) Union(p2 Pmfset) Pmfset {
	pr := make(map[string]int)
	pd := p2.Subtract(p1)
	for k := range p1 {
		pr[k] = p1.Stringcount(k) + p2.Stringcount(k)
	}
	for k := range pd {
		pr[k] = p2.Stringcount(k)
	}
	return pr
}

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
