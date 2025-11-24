package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_format1 struct {
	InitialCyclicShift  int64 `lb:0,ub:11,madatory`
	NrofSymbols         int64 `lb:4,ub:14,madatory`
	StartingSymbolIndex int64 `lb:0,ub:10,madatory`
	TimeDomainOCC       int64 `lb:0,ub:6,madatory`
}

func (ie *PUCCH_format1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.InitialCyclicShift, &uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("WriteInteger InitialCyclicShift", err)
	}
	if err = w.WriteInteger(ie.NrofSymbols, &uper.Constraint{Lb: 4, Ub: 14}, false); err != nil {
		return utils.WrapError("WriteInteger NrofSymbols", err)
	}
	if err = w.WriteInteger(ie.StartingSymbolIndex, &uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("WriteInteger StartingSymbolIndex", err)
	}
	if err = w.WriteInteger(ie.TimeDomainOCC, &uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteInteger TimeDomainOCC", err)
	}
	return nil
}

func (ie *PUCCH_format1) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_InitialCyclicShift int64
	if tmp_int_InitialCyclicShift, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("ReadInteger InitialCyclicShift", err)
	}
	ie.InitialCyclicShift = tmp_int_InitialCyclicShift
	var tmp_int_NrofSymbols int64
	if tmp_int_NrofSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 14}, false); err != nil {
		return utils.WrapError("ReadInteger NrofSymbols", err)
	}
	ie.NrofSymbols = tmp_int_NrofSymbols
	var tmp_int_StartingSymbolIndex int64
	if tmp_int_StartingSymbolIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("ReadInteger StartingSymbolIndex", err)
	}
	ie.StartingSymbolIndex = tmp_int_StartingSymbolIndex
	var tmp_int_TimeDomainOCC int64
	if tmp_int_TimeDomainOCC, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadInteger TimeDomainOCC", err)
	}
	ie.TimeDomainOCC = tmp_int_TimeDomainOCC
	return nil
}
