package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MobilityStateParameters struct {
	T_Evaluation       MobilityStateParameters_t_Evaluation `madatory`
	T_HystNormal       MobilityStateParameters_t_HystNormal `madatory`
	N_CellChangeMedium int64                                `lb:1,ub:16,madatory`
	N_CellChangeHigh   int64                                `lb:1,ub:16,madatory`
}

func (ie *MobilityStateParameters) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.T_Evaluation.Encode(w); err != nil {
		return utils.WrapError("Encode T_Evaluation", err)
	}
	if err = ie.T_HystNormal.Encode(w); err != nil {
		return utils.WrapError("Encode T_HystNormal", err)
	}
	if err = w.WriteInteger(ie.N_CellChangeMedium, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger N_CellChangeMedium", err)
	}
	if err = w.WriteInteger(ie.N_CellChangeHigh, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger N_CellChangeHigh", err)
	}
	return nil
}

func (ie *MobilityStateParameters) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.T_Evaluation.Decode(r); err != nil {
		return utils.WrapError("Decode T_Evaluation", err)
	}
	if err = ie.T_HystNormal.Decode(r); err != nil {
		return utils.WrapError("Decode T_HystNormal", err)
	}
	var tmp_int_N_CellChangeMedium int64
	if tmp_int_N_CellChangeMedium, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger N_CellChangeMedium", err)
	}
	ie.N_CellChangeMedium = tmp_int_N_CellChangeMedium
	var tmp_int_N_CellChangeHigh int64
	if tmp_int_N_CellChangeHigh, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger N_CellChangeHigh", err)
	}
	ie.N_CellChangeHigh = tmp_int_N_CellChangeHigh
	return nil
}
