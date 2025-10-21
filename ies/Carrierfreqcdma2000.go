package ies

import "rrc/utils"

// CarrierFreqCDMA2000 ::= SEQUENCE
type Carrierfreqcdma2000 struct {
	Bandclass Bandclasscdma2000
	Arfcn     ArfcnValuecdma2000
}
