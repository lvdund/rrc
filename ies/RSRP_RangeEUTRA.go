package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RSRP_RangeEUTRA struct {
	Value uint64 `lb:0,ub:97,madatory`
}

func (ie *RSRP_RangeEUTRA) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 97}, false); err != nil {
		return utils.WrapError("Encode RSRP_RangeEUTRA", err)
	}
	return nil
}

func (ie *RSRP_RangeEUTRA) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 97}, false); err != nil {
		return utils.WrapError("Decode RSRP_RangeEUTRA", err)
	}
	ie.Value = uint64(v)
	return nil
}
