package ies

// UplinkDataCompression-r17 ::= CHOICE
const (
	UplinkdatacompressionR17ChoiceNothing = iota
	UplinkdatacompressionR17ChoiceNewsetup
	UplinkdatacompressionR17ChoiceDrbContinueudc
)

type UplinkdatacompressionR17 struct {
	Choice         uint64
	Newsetup       *UplinkdatacompressionR17Newsetup
	DrbContinueudc *struct{}
}
