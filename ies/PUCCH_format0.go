package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_format0 struct {
	InitialCyclicShift  int64 `lb:0,ub:11,madatory`
	NrofSymbols         int64 `lb:1,ub:2,madatory`
	StartingSymbolIndex int64 `lb:0,ub:13,madatory`
}

func (ie *PUCCH_format0) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.InitialCyclicShift, &aper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("WriteInteger InitialCyclicShift", err)
	}
	if err = w.WriteInteger(ie.NrofSymbols, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger NrofSymbols", err)
	}
	if err = w.WriteInteger(ie.StartingSymbolIndex, &aper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger StartingSymbolIndex", err)
	}
	return nil
}

func (ie *PUCCH_format0) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_InitialCyclicShift int64
	if tmp_int_InitialCyclicShift, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("ReadInteger InitialCyclicShift", err)
	}
	ie.InitialCyclicShift = tmp_int_InitialCyclicShift
	var tmp_int_NrofSymbols int64
	if tmp_int_NrofSymbols, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger NrofSymbols", err)
	}
	ie.NrofSymbols = tmp_int_NrofSymbols
	var tmp_int_StartingSymbolIndex int64
	if tmp_int_StartingSymbolIndex, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger StartingSymbolIndex", err)
	}
	ie.StartingSymbolIndex = tmp_int_StartingSymbolIndex
	return nil
}
