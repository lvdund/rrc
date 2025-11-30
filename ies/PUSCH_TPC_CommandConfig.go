package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_TPC_CommandConfig struct {
	Tpc_Index    *int64         `lb:1,ub:15,optional`
	Tpc_IndexSUL *int64         `lb:1,ub:15,optional`
	TargetCell   *ServCellIndex `optional`
}

func (ie *PUSCH_TPC_CommandConfig) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Tpc_Index != nil, ie.Tpc_IndexSUL != nil, ie.TargetCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Tpc_Index != nil {
		if err = w.WriteInteger(*ie.Tpc_Index, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Tpc_Index", err)
		}
	}
	if ie.Tpc_IndexSUL != nil {
		if err = w.WriteInteger(*ie.Tpc_IndexSUL, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Tpc_IndexSUL", err)
		}
	}
	if ie.TargetCell != nil {
		if err = ie.TargetCell.Encode(w); err != nil {
			return utils.WrapError("Encode TargetCell", err)
		}
	}
	return nil
}

func (ie *PUSCH_TPC_CommandConfig) Decode(r *aper.AperReader) error {
	var err error
	var Tpc_IndexPresent bool
	if Tpc_IndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_IndexSULPresent bool
	if Tpc_IndexSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TargetCellPresent bool
	if TargetCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Tpc_IndexPresent {
		var tmp_int_Tpc_Index int64
		if tmp_int_Tpc_Index, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Tpc_Index", err)
		}
		ie.Tpc_Index = &tmp_int_Tpc_Index
	}
	if Tpc_IndexSULPresent {
		var tmp_int_Tpc_IndexSUL int64
		if tmp_int_Tpc_IndexSUL, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Tpc_IndexSUL", err)
		}
		ie.Tpc_IndexSUL = &tmp_int_Tpc_IndexSUL
	}
	if TargetCellPresent {
		ie.TargetCell = new(ServCellIndex)
		if err = ie.TargetCell.Decode(r); err != nil {
			return utils.WrapError("Decode TargetCell", err)
		}
	}
	return nil
}
