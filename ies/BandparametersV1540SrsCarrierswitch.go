package ies

// BandParameters-v1540-srs-CarrierSwitch ::= CHOICE
const (
	BandparametersV1540SrsCarrierswitchChoiceNothing = iota
	BandparametersV1540SrsCarrierswitchChoiceNr
	BandparametersV1540SrsCarrierswitchChoiceEutra
)

type BandparametersV1540SrsCarrierswitch struct {
	Choice uint64
	Nr     *BandparametersV1540SrsCarrierswitchNr
	Eutra  *BandparametersV1540SrsCarrierswitchEutra
}
