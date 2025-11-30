package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReestablishmentCause_Enum_reconfigurationFailure aper.Enumerated = 0
	ReestablishmentCause_Enum_handoverFailure        aper.Enumerated = 1
	ReestablishmentCause_Enum_otherFailure           aper.Enumerated = 2
	ReestablishmentCause_Enum_spare1                 aper.Enumerated = 3
)

type ReestablishmentCause struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ReestablishmentCause) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ReestablishmentCause", err)
	}
	return nil
}

func (ie *ReestablishmentCause) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ReestablishmentCause", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
