package ies

// RedirectedCarrierInfo-r15-IEs ::= CHOICE
const (
	RedirectedcarrierinfoR15ChoiceNothing = iota
	RedirectedcarrierinfoR15ChoiceEutra
	RedirectedcarrierinfoR15ChoiceGeran
	RedirectedcarrierinfoR15ChoiceUtraFdd
	RedirectedcarrierinfoR15ChoiceCdma2000Hrpd
	RedirectedcarrierinfoR15ChoiceCdma20001xrtt
	RedirectedcarrierinfoR15ChoiceUtraTdd
)

type RedirectedcarrierinfoR15 struct {
	Choice        uint64
	Eutra         *ArfcnValueeutraR9
	Geran         *Carrierfreqsgeran
	UtraFdd       *ArfcnValueutra
	Cdma2000Hrpd  *Carrierfreqcdma2000
	Cdma20001xrtt *Carrierfreqcdma2000
	UtraTdd       *CarrierfreqlistutraTddR10
}
