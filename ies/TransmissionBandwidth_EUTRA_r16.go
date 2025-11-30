package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TransmissionBandwidth_EUTRA_r16_Enum_rb6   aper.Enumerated = 0
	TransmissionBandwidth_EUTRA_r16_Enum_rb15  aper.Enumerated = 1
	TransmissionBandwidth_EUTRA_r16_Enum_rb25  aper.Enumerated = 2
	TransmissionBandwidth_EUTRA_r16_Enum_rb50  aper.Enumerated = 3
	TransmissionBandwidth_EUTRA_r16_Enum_rb75  aper.Enumerated = 4
	TransmissionBandwidth_EUTRA_r16_Enum_rb100 aper.Enumerated = 5
)

type TransmissionBandwidth_EUTRA_r16 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *TransmissionBandwidth_EUTRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode TransmissionBandwidth_EUTRA_r16", err)
	}
	return nil
}

func (ie *TransmissionBandwidth_EUTRA_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode TransmissionBandwidth_EUTRA_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
