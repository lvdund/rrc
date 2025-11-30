package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UPInterruptionTimeAtHO_r17 struct {
	Value uint64 `lb:0,ub:1023,madatory`
}

func (ie *UPInterruptionTimeAtHO_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("Encode UPInterruptionTimeAtHO_r17", err)
	}
	return nil
}

func (ie *UPInterruptionTimeAtHO_r17) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("Decode UPInterruptionTimeAtHO_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
