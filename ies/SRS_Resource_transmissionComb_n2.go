package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_transmissionComb_n2 struct {
	CombOffset_n2  int64 `lb:0,ub:1,madatory`
	CyclicShift_n2 int64 `lb:0,ub:7,madatory`
}

func (ie *SRS_Resource_transmissionComb_n2) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.CombOffset_n2, &aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteInteger CombOffset_n2", err)
	}
	if err = w.WriteInteger(ie.CyclicShift_n2, &aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger CyclicShift_n2", err)
	}
	return nil
}

func (ie *SRS_Resource_transmissionComb_n2) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_CombOffset_n2 int64
	if tmp_int_CombOffset_n2, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadInteger CombOffset_n2", err)
	}
	ie.CombOffset_n2 = tmp_int_CombOffset_n2
	var tmp_int_CyclicShift_n2 int64
	if tmp_int_CyclicShift_n2, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger CyclicShift_n2", err)
	}
	ie.CyclicShift_n2 = tmp_int_CyclicShift_n2
	return nil
}
