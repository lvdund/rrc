package ies

import "rrc/utils"

// RedirectedCarrierInfo ::= CHOICE
// Extensible
type Redirectedcarrierinfo interface {
	isRedirectedcarrierinfo()
}

type RedirectedcarrierinfoEutra struct {
	Value ArfcnValueeutra
}

func (*RedirectedcarrierinfoEutra) isRedirectedcarrierinfo() {}

type RedirectedcarrierinfoGeran struct {
	Value Carrierfreqsgeran
}

func (*RedirectedcarrierinfoGeran) isRedirectedcarrierinfo() {}

type RedirectedcarrierinfoUtraFdd struct {
	Value ArfcnValueutra
}

func (*RedirectedcarrierinfoUtraFdd) isRedirectedcarrierinfo() {}

type RedirectedcarrierinfoUtraTdd struct {
	Value ArfcnValueutra
}

func (*RedirectedcarrierinfoUtraTdd) isRedirectedcarrierinfo() {}

type RedirectedcarrierinfoCdma2000Hrpd struct {
	Value Carrierfreqcdma2000
}

func (*RedirectedcarrierinfoCdma2000Hrpd) isRedirectedcarrierinfo() {}

type RedirectedcarrierinfoCdma20001xrtt struct {
	Value Carrierfreqcdma2000
}

func (*RedirectedcarrierinfoCdma20001xrtt) isRedirectedcarrierinfo() {}

type RedirectedcarrierinfoUtraTddR10 struct {
	Value CarrierfreqlistutraTddR10
}

func (*RedirectedcarrierinfoUtraTddR10) isRedirectedcarrierinfo() {}

type RedirectedcarrierinfoNrR15 struct {
	Value CarrierinfonrR15
}

func (*RedirectedcarrierinfoNrR15) isRedirectedcarrierinfo() {}
