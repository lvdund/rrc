package ies

// BandCombination-UplinkTxSwitch-r16 ::= SEQUENCE
// Extensible
type BandcombinationUplinktxswitchR16 struct {
	BandcombinationR16                      Bandcombination
	BandcombinationV1540                    *BandcombinationV1540
	BandcombinationV1560                    *BandcombinationV1560
	BandcombinationV1570                    *BandcombinationV1570
	BandcombinationV1580                    *BandcombinationV1580
	BandcombinationV1590                    *BandcombinationV1590
	BandcombinationV1610                    *BandcombinationV1610
	SupportedbandpairlistnrR16              []UltxswitchingbandpairR16 `lb:1,ub:maxULTxSwitchingBandPairs`
	UplinktxswitchingOptionsupportR16       *BandcombinationUplinktxswitchR16UplinktxswitchingOptionsupportR16
	UplinktxswitchingPowerboostingR16       *BandcombinationUplinktxswitchR16UplinktxswitchingPowerboostingR16
	UplinktxswitchingPuschTranscoherenceR16 *BandcombinationUplinktxswitchR16UplinktxswitchingPuschTranscoherenceR16 `ext`
}
