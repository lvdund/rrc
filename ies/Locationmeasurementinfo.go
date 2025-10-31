package ies

// LocationMeasurementInfo ::= CHOICE
// Extensible
const (
	LocationmeasurementinfoChoiceNothing = iota
	LocationmeasurementinfoChoiceEutraRstd
	LocationmeasurementinfoChoiceEutraFinetimingdetection
	LocationmeasurementinfoChoiceNrPrsMeasurementR16
)

type Locationmeasurementinfo struct {
	Choice                   uint64
	EutraRstd                *EutraRstdInfolist
	EutraFinetimingdetection *struct{}
	NrPrsMeasurementR16      *NrPrsMeasurementinfolistR16
}
