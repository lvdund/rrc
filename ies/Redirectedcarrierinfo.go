package ies

// RedirectedCarrierInfo ::= CHOICE
// Extensible
const (
	RedirectedcarrierinfoChoiceNothing = iota
	RedirectedcarrierinfoChoiceEutra
	RedirectedcarrierinfoChoiceGeran
	RedirectedcarrierinfoChoiceUtraFdd
	RedirectedcarrierinfoChoiceUtraTdd
	RedirectedcarrierinfoChoiceCdma2000Hrpd
	RedirectedcarrierinfoChoiceCdma20001xrtt
	RedirectedcarrierinfoChoiceUtraTddR10
	RedirectedcarrierinfoChoiceNrR15
)

type Redirectedcarrierinfo struct {
	Choice        uint64
	Eutra         *ArfcnValueeutra
	Geran         *Carrierfreqsgeran
	UtraFdd       *ArfcnValueutra
	UtraTdd       *ArfcnValueutra
	Cdma2000Hrpd  *Carrierfreqcdma2000
	Cdma20001xrtt *Carrierfreqcdma2000
	UtraTddR10    *CarrierfreqlistutraTddR10
	NrR15         *CarrierinfonrR15
}
