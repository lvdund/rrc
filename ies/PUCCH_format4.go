package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_format4 struct {
	NrofSymbols         int64                    `lb:4,ub:14,madatory`
	Occ_Length          PUCCH_format4_occ_Length `madatory`
	Occ_Index           PUCCH_format4_occ_Index  `madatory`
	StartingSymbolIndex int64                    `lb:0,ub:10,madatory`
}

func (ie *PUCCH_format4) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.NrofSymbols, &uper.Constraint{Lb: 4, Ub: 14}, false); err != nil {
		return utils.WrapError("WriteInteger NrofSymbols", err)
	}
	if err = ie.Occ_Length.Encode(w); err != nil {
		return utils.WrapError("Encode Occ_Length", err)
	}
	if err = ie.Occ_Index.Encode(w); err != nil {
		return utils.WrapError("Encode Occ_Index", err)
	}
	if err = w.WriteInteger(ie.StartingSymbolIndex, &uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("WriteInteger StartingSymbolIndex", err)
	}
	return nil
}

func (ie *PUCCH_format4) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_NrofSymbols int64
	if tmp_int_NrofSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 14}, false); err != nil {
		return utils.WrapError("ReadInteger NrofSymbols", err)
	}
	ie.NrofSymbols = tmp_int_NrofSymbols
	if err = ie.Occ_Length.Decode(r); err != nil {
		return utils.WrapError("Decode Occ_Length", err)
	}
	if err = ie.Occ_Index.Decode(r); err != nil {
		return utils.WrapError("Decode Occ_Index", err)
	}
	var tmp_int_StartingSymbolIndex int64
	if tmp_int_StartingSymbolIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("ReadInteger StartingSymbolIndex", err)
	}
	ie.StartingSymbolIndex = tmp_int_StartingSymbolIndex
	return nil
}
