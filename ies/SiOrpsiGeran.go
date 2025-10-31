package ies

// SI-OrPSI-GERAN ::= CHOICE
const (
	SiOrpsiGeranChoiceNothing = iota
	SiOrpsiGeranChoiceSi
	SiOrpsiGeranChoicePsi
)

type SiOrpsiGeran struct {
	Choice uint64
	Si     *Systeminfolistgeran
	Psi    *Systeminfolistgeran
}
