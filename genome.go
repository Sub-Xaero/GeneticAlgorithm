package ga

import "fmt"

// Genome represents a bitstring and associated fitness value
type Genome struct {
	Sequence Bitstring
}

func (gene Genome) Copy() Genome {
	sequence := make(Bitstring, len(gene.Sequence))
	copy(sequence, gene.Sequence)
	return Genome{sequence}
}

func (gene Genome) String() string {
	return fmt.Sprintf("{%v}", gene.Sequence)
}
