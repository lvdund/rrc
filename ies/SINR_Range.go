package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SINR_Range struct {
	Value uint64 `lb:0,ub:127,madatory`
}

func (ie *SINR_Range) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("Encode SINR_Range", err)
	}
	return nil
}

func (ie *SINR_Range) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("Decode SINR_Range", err)
	}
	ie.Value = uint64(v)
	return nil
}
