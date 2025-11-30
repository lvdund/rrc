package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EthernetHeaderCompression_r16_ehc_Common_r16_ehc_CID_Length_r16_Enum_bits7  aper.Enumerated = 0
	EthernetHeaderCompression_r16_ehc_Common_r16_ehc_CID_Length_r16_Enum_bits15 aper.Enumerated = 1
)

type EthernetHeaderCompression_r16_ehc_Common_r16_ehc_CID_Length_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *EthernetHeaderCompression_r16_ehc_Common_r16_ehc_CID_Length_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode EthernetHeaderCompression_r16_ehc_Common_r16_ehc_CID_Length_r16", err)
	}
	return nil
}

func (ie *EthernetHeaderCompression_r16_ehc_Common_r16_ehc_CID_Length_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode EthernetHeaderCompression_r16_ehc_Common_r16_ehc_CID_Length_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
