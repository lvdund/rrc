package ies

import "rrc/utils"

// RedirectedCarrierInfo-r15-IEs ::= CHOICE
type RedirectedcarrierinfoR15Ies interface {
	isRedirectedcarrierinfoR15Ies()
}

type RedirectedcarrierinfoR15IesEutra struct {
	Value ArfcnValueeutraR9
}

func (*RedirectedcarrierinfoR15IesEutra) isRedirectedcarrierinfoR15Ies() {}

type RedirectedcarrierinfoR15IesGeran struct {
	Value Carrierfreqsgeran
}

func (*RedirectedcarrierinfoR15IesGeran) isRedirectedcarrierinfoR15Ies() {}

type RedirectedcarrierinfoR15IesUtraFdd struct {
	Value ArfcnValueutra
}

func (*RedirectedcarrierinfoR15IesUtraFdd) isRedirectedcarrierinfoR15Ies() {}

type RedirectedcarrierinfoR15IesCdma2000Hrpd struct {
	Value Carrierfreqcdma2000
}

func (*RedirectedcarrierinfoR15IesCdma2000Hrpd) isRedirectedcarrierinfoR15Ies() {}

type RedirectedcarrierinfoR15IesCdma20001xrtt struct {
	Value Carrierfreqcdma2000
}

func (*RedirectedcarrierinfoR15IesCdma20001xrtt) isRedirectedcarrierinfoR15Ies() {}

type RedirectedcarrierinfoR15IesUtraTdd struct {
	Value CarrierfreqlistutraTddR10
}

func (*RedirectedcarrierinfoR15IesUtraTdd) isRedirectedcarrierinfoR15Ies() {}
