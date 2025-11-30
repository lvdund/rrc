package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_transmissionComb_n4 struct {
	CombOffset_n4  int64 `lb:0,ub:3,madatory`
	CyclicShift_n4 int64 `lb:0,ub:11,madatory`
}

func (ie *SRS_Resource_transmissionComb_n4) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.CombOffset_n4, &aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger CombOffset_n4", err)
	}
	if err = w.WriteInteger(ie.CyclicShift_n4, &aper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("WriteInteger CyclicShift_n4", err)
	}
	return nil
}

func (ie *SRS_Resource_transmissionComb_n4) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_CombOffset_n4 int64
	if tmp_int_CombOffset_n4, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger CombOffset_n4", err)
	}
	ie.CombOffset_n4 = tmp_int_CombOffset_n4
	var tmp_int_CyclicShift_n4 int64
	if tmp_int_CyclicShift_n4, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("ReadInteger CyclicShift_n4", err)
	}
	ie.CyclicShift_n4 = tmp_int_CyclicShift_n4
	return nil
}
