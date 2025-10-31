package ies

import "rrc/utils"

// RF-Parameters ::= SEQUENCE
// Extensible
type RfParameters struct {
	Supportedbandlistnr                                  []Bandnr `lb:1,ub:maxBands`
	Supportedbandcombinationlist                         *Bandcombinationlist
	Appliedfreqbandlistfilter                            *Freqbandlist
	SupportedbandcombinationlistV1540                    *BandcombinationlistV1540                `ext`
	SrsSwitchingtimerequested                            *utils.BOOLEAN                           `ext`
	SupportedbandcombinationlistV1550                    *BandcombinationlistV1550                `ext`
	SupportedbandcombinationlistV1560                    *BandcombinationlistV1560                `ext`
	SupportedbandcombinationlistV1610                    *BandcombinationlistV1610                `ext`
	SupportedbandcombinationlistsidelinkeutraNrR16       *BandcombinationlistsidelinkeutraNrR16   `ext`
	SupportedbandcombinationlistUplinktxswitchR16        *BandcombinationlistUplinktxswitchR16    `ext`
	SupportedbandcombinationlistV1630                    *BandcombinationlistV1630                `ext`
	SupportedbandcombinationlistsidelinkeutraNrV1630     *BandcombinationlistsidelinkeutraNrV1630 `ext`
	SupportedbandcombinationlistUplinktxswitchV1630      *BandcombinationlistUplinktxswitchV1630  `ext`
	SupportedbandcombinationlistV1640                    *BandcombinationlistV1640                `ext`
	SupportedbandcombinationlistUplinktxswitchV1640      *BandcombinationlistUplinktxswitchV1640  `ext`
	SupportedbandcombinationlistV1650                    *BandcombinationlistV1650                `ext`
	SupportedbandcombinationlistUplinktxswitchV1650      *BandcombinationlistUplinktxswitchV1650  `ext`
	ExtendedbandN77R16                                   *RfParametersExtendedbandN77R16          `ext`
	SupportedbandcombinationlistUplinktxswitchV1670      *BandcombinationlistUplinktxswitchV1670  `ext`
	SupportedbandcombinationlistV1680                    *BandcombinationlistV1680                `ext`
	SupportedbandcombinationlistV1690                    *BandcombinationlistV1690                `ext`
	SupportedbandcombinationlistUplinktxswitchV1690      *BandcombinationlistUplinktxswitchV1690  `ext`
	SupportedbandcombinationlistV1700                    *BandcombinationlistV1700                `ext`
	SupportedbandcombinationlistUplinktxswitchV1700      *BandcombinationlistUplinktxswitchV1700  `ext`
	SupportedbandcombinationlistslRelaydiscoveryR17      *utils.OCTETSTRING                       `ext`
	SupportedbandcombinationlistslNonrelaydiscoveryR17   *utils.OCTETSTRING                       `ext`
	SupportedbandcombinationlistsidelinkeutraNrV1710     *BandcombinationlistsidelinkeutraNrV1710 `ext`
	SidelinkrequestedR17                                 *utils.BOOLEAN                           `ext`
	ExtendedbandN772R17                                  *RfParametersExtendedbandN772R17         `ext`
	SupportedbandcombinationlistV1720                    *BandcombinationlistV1720                `ext`
	SupportedbandcombinationlistUplinktxswitchV1720      *BandcombinationlistUplinktxswitchV1720  `ext`
	SupportedbandcombinationlistV1730                    *BandcombinationlistV1730                `ext`
	SupportedbandcombinationlistUplinktxswitchV1730      *BandcombinationlistUplinktxswitchV1730  `ext`
	SupportedbandcombinationlistslRelaydiscoveryV1730    *BandcombinationlistslDiscoveryR17       `ext`
	SupportedbandcombinationlistslNonrelaydiscoveryV1730 *BandcombinationlistslDiscoveryR17       `ext`
}
