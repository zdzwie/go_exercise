package unionFind

type struct UnionFind {
	numbers []int
	id      []int
}

func (uf *UnionFind) Init(n int) {
	uf.numbers = make([]int, n)
	uf.id = make([]int, n)
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.id[p] == uf.id[q]
}

func (uf *UnionFind) Union(p, q int) {
	pid := uf.id[p]
	qid := uf.id[q]
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pid {
			uf.id[i] = qid
		}
	}
}
