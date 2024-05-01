package unionFind

type WeightedQuickUnion struct {
	numbers []int
	id      []int
	sz      []int
}

func (wqu *WeightedQuickUnion) Init(n int) {
	wqu.numbers = make([]int, n)
	wqu.id = make([]int, n)
	wqu.sz = make([]int, n)
	for i := 0; i < n; i++ {
		wqu.id[i] = i
		wqu.sz[i] = 1
	}
}

func (wqu *WeightedQuickUnion) root(i int) int {
	for i != wqu.id[i] {
		wqu.id[i] = wqu.id[wqu.id[i]]  // Path compression
		i = wqu.id[i]
	}

	return i
}

func (wqu *WeightedQuickUnion) Connected(p, q int) bool {
	return wqu.root(p) == wqu.root(q)
}

func (wqu *WeightedQuickUnion) Union(p, q int) {
	i := wqu.root(p)
	j := wqu.root(q)
	if i == j {
		return
	}

	if wqu.sz[i] < wqu.sz[j] {
		wqu.id[i] = j
		wqu.sz[j] += wqu.sz[i]
	} else {
		wqu.id[j] = i
		wqu.sz[i] += wqu.sz[j]
	}
}
