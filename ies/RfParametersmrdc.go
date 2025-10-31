package ies

import "rrc/utils"

// RF-ParametersMRDC ::= SEQUENCE
// Extensible
type RfParametersmrdc struct {
	Supportedbandcombinationlist                    *Bandcombinationlist
	Appliedfreqbandlistfilter                       *Freqbandlist
	SrsSwitchingtimerequested                       *utils.BOOLEAN                                             `ext`
	SupportedbandcombinationlistV1540               *BandcombinationlistV1540                                  `ext`
	SupportedbandcombinationlistV1550               *BandcombinationlistV1550                                  `ext`
	SupportedbandcombinationlistV1560               *BandcombinationlistV1560                                  `ext`
	SupportedbandcombinationlistnedcOnly            *Bandcombinationlist                                       `ext`
	SupportedbandcombinationlistV1570               *BandcombinationlistV1570                                  `ext`
	SupportedbandcombinationlistV1580               *BandcombinationlistV1580                                  `ext`
	SupportedbandcombinationlistV1590               *BandcombinationlistV1590                                  `ext`
	SupportedbandcombinationlistnedcOnlyV15a0       *RfParametersmrdcSupportedbandcombinationlistnedcOnlyV15a0 `ext`
	SupportedbandcombinationlistV1610               *BandcombinationlistV1610                                  `ext`
	SupportedbandcombinationlistnedcOnlyV1610       *BandcombinationlistV1610                                  `ext`
	SupportedbandcombinationlistUplinktxswitchR16   *BandcombinationlistUplinktxswitchR16                      `ext`
	SupportedbandcombinationlistV1630               *BandcombinationlistV1630                                  `ext`
	SupportedbandcombinationlistnedcOnlyV1630       *BandcombinationlistV1630                                  `ext`
	SupportedbandcombinationlistUplinktxswitchV1630 *BandcombinationlistUplinktxswitchV1630                    `ext`
	SupportedbandcombinationlistV1640               *BandcombinationlistV1640                                  `ext`
	SupportedbandcombinationlistnedcOnlyV1640       *BandcombinationlistV1640                                  `ext`
	SupportedbandcombinationlistUplinktxswitchV1640 *BandcombinationlistUplinktxswitchV1640                    `ext`
	SupportedbandcombinationlistUplinktxswitchV1670 *BandcombinationlistUplinktxswitchV1670                    `ext`
	SupportedbandcombinationlistV1700               *BandcombinationlistV1700                                  `ext`
	SupportedbandcombinationlistUplinktxswitchV1700 *BandcombinationlistUplinktxswitchV1700                    `ext`
	SupportedbandcombinationlistV1720               *BandcombinationlistV1720                                  `ext`
	SupportedbandcombinationlistnedcOnlyV1720       *RfParametersmrdcSupportedbandcombinationlistnedcOnlyV1720 `ext`
	SupportedbandcombinationlistUplinktxswitchV1720 *BandcombinationlistUplinktxswitchV1720                    `ext`
	SupportedbandcombinationlistV1730               *BandcombinationlistV1730                                  `ext`
	SupportedbandcombinationlistnedcOnlyV1730       *BandcombinationlistV1730                                  `ext`
	SupportedbandcombinationlistUplinktxswitchV1730 *BandcombinationlistUplinktxswitchV1730                    `ext`
}
