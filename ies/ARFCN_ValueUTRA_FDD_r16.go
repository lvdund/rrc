package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ARFCN_ValueUTRA_FDD_r16 struct {
	Value uint64 `lb:0,ub:16383,madatory`
}

func (ie *ARFCN_ValueUTRA_FDD_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
		return utils.WrapError("Encode ARFCN_ValueUTRA_FDD_r16", err)
	}
	return nil
}

func (ie *ARFCN_ValueUTRA_FDD_r16) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
		return utils.WrapError("Decode ARFCN_ValueUTRA_FDD_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
