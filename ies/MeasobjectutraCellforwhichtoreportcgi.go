package ies

// MeasObjectUTRA-cellForWhichToReportCGI ::= CHOICE
const (
	MeasobjectutraCellforwhichtoreportcgiChoiceNothing = iota
	MeasobjectutraCellforwhichtoreportcgiChoiceUtraFdd
	MeasobjectutraCellforwhichtoreportcgiChoiceUtraTdd
)

type MeasobjectutraCellforwhichtoreportcgi struct {
	Choice  uint64
	UtraFdd *PhyscellidutraFdd
	UtraTdd *PhyscellidutraTdd
}
