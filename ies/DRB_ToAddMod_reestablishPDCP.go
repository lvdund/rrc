package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRB_ToAddMod_reestablishPDCP_Enum_true aper.Enumerated = 0
)

type DRB_ToAddMod_reestablishPDCP struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *DRB_ToAddMod_reestablishPDCP) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode DRB_ToAddMod_reestablishPDCP", err)
	}
	return nil
}

func (ie *DRB_ToAddMod_reestablishPDCP) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode DRB_ToAddMod_reestablishPDCP", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
