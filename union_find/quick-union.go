package unionFind

type QuickUnion struct {
	numbers []int
	id      []int
}

func (qu *QuickUnion) Init(n int) {
	qu.numbers = make([]int, n)
	qu.id = make([]int, n)
	for i := 0; i < n; i++ {
		qu.id[i] = i
	}
}

func (qu *QuickUnion) root(i int) int {
	for i != qu.id[i] {
		i = qu.id[i]
	}
	return i
}

func (qu *QuickUnion) Connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu *QuickUnion) Union(p, q int) {
	i := qu.root(p)
	j := qu.root(q)
	qu.id[i] = j
}
