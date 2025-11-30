package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRB_ToAddMod_discardOnPDCP_Enum_true aper.Enumerated = 0
)

type SRB_ToAddMod_discardOnPDCP struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SRB_ToAddMod_discardOnPDCP) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SRB_ToAddMod_discardOnPDCP", err)
	}
	return nil
}

func (ie *SRB_ToAddMod_discardOnPDCP) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SRB_ToAddMod_discardOnPDCP", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
