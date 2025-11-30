package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRB_ToAddMod_r17_recoverPDCP_r17_Enum_true aper.Enumerated = 0
)

type MRB_ToAddMod_r17_recoverPDCP_r17 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MRB_ToAddMod_r17_recoverPDCP_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MRB_ToAddMod_r17_recoverPDCP_r17", err)
	}
	return nil
}

func (ie *MRB_ToAddMod_r17_recoverPDCP_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MRB_ToAddMod_r17_recoverPDCP_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
