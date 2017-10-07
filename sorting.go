package ga

// ByFitness is a receiver type that implements Sort for Genome []
type ByFitness []Genome

func (a ByFitness) Len() int           { return len(a) }
func (a ByFitness) Swap(x, y int)      { a[x], a[y] = a[y], a[x] }
func (a ByFitness) Less(x, y int) bool { return a[x].Fitness() < a[y].Fitness() }
