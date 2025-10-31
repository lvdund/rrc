package ies

// UplinkTxDirectCurrentTwoCarrier-r16 ::= SEQUENCE
type UplinktxdirectcurrenttwocarrierR16 struct {
	CarrieroneinfoR16          UplinktxdirectcurrentcarrierinfoR16
	CarriertwoinfoR16          UplinktxdirectcurrentcarrierinfoR16
	SinglepaTxdirectcurrentR16 UplinktxdirectcurrenttwocarrierinfoR16
	SecondpaTxdirectcurrentR16 *UplinktxdirectcurrenttwocarrierinfoR16
}
