package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BetaOffsets struct {
	BetaOffsetACK_Index1       *int64 `lb:0,ub:31,optional`
	BetaOffsetACK_Index2       *int64 `lb:0,ub:31,optional`
	BetaOffsetACK_Index3       *int64 `lb:0,ub:31,optional`
	BetaOffsetCSI_Part1_Index1 *int64 `lb:0,ub:31,optional`
	BetaOffsetCSI_Part1_Index2 *int64 `lb:0,ub:31,optional`
	BetaOffsetCSI_Part2_Index1 *int64 `lb:0,ub:31,optional`
	BetaOffsetCSI_Part2_Index2 *int64 `lb:0,ub:31,optional`
}

func (ie *BetaOffsets) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.BetaOffsetACK_Index1 != nil, ie.BetaOffsetACK_Index2 != nil, ie.BetaOffsetACK_Index3 != nil, ie.BetaOffsetCSI_Part1_Index1 != nil, ie.BetaOffsetCSI_Part1_Index2 != nil, ie.BetaOffsetCSI_Part2_Index1 != nil, ie.BetaOffsetCSI_Part2_Index2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BetaOffsetACK_Index1 != nil {
		if err = w.WriteInteger(*ie.BetaOffsetACK_Index1, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode BetaOffsetACK_Index1", err)
		}
	}
	if ie.BetaOffsetACK_Index2 != nil {
		if err = w.WriteInteger(*ie.BetaOffsetACK_Index2, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode BetaOffsetACK_Index2", err)
		}
	}
	if ie.BetaOffsetACK_Index3 != nil {
		if err = w.WriteInteger(*ie.BetaOffsetACK_Index3, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode BetaOffsetACK_Index3", err)
		}
	}
	if ie.BetaOffsetCSI_Part1_Index1 != nil {
		if err = w.WriteInteger(*ie.BetaOffsetCSI_Part1_Index1, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode BetaOffsetCSI_Part1_Index1", err)
		}
	}
	if ie.BetaOffsetCSI_Part1_Index2 != nil {
		if err = w.WriteInteger(*ie.BetaOffsetCSI_Part1_Index2, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode BetaOffsetCSI_Part1_Index2", err)
		}
	}
	if ie.BetaOffsetCSI_Part2_Index1 != nil {
		if err = w.WriteInteger(*ie.BetaOffsetCSI_Part2_Index1, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode BetaOffsetCSI_Part2_Index1", err)
		}
	}
	if ie.BetaOffsetCSI_Part2_Index2 != nil {
		if err = w.WriteInteger(*ie.BetaOffsetCSI_Part2_Index2, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode BetaOffsetCSI_Part2_Index2", err)
		}
	}
	return nil
}

func (ie *BetaOffsets) Decode(r *aper.AperReader) error {
	var err error
	var BetaOffsetACK_Index1Present bool
	if BetaOffsetACK_Index1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BetaOffsetACK_Index2Present bool
	if BetaOffsetACK_Index2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BetaOffsetACK_Index3Present bool
	if BetaOffsetACK_Index3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BetaOffsetCSI_Part1_Index1Present bool
	if BetaOffsetCSI_Part1_Index1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BetaOffsetCSI_Part1_Index2Present bool
	if BetaOffsetCSI_Part1_Index2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BetaOffsetCSI_Part2_Index1Present bool
	if BetaOffsetCSI_Part2_Index1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BetaOffsetCSI_Part2_Index2Present bool
	if BetaOffsetCSI_Part2_Index2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BetaOffsetACK_Index1Present {
		var tmp_int_BetaOffsetACK_Index1 int64
		if tmp_int_BetaOffsetACK_Index1, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode BetaOffsetACK_Index1", err)
		}
		ie.BetaOffsetACK_Index1 = &tmp_int_BetaOffsetACK_Index1
	}
	if BetaOffsetACK_Index2Present {
		var tmp_int_BetaOffsetACK_Index2 int64
		if tmp_int_BetaOffsetACK_Index2, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode BetaOffsetACK_Index2", err)
		}
		ie.BetaOffsetACK_Index2 = &tmp_int_BetaOffsetACK_Index2
	}
	if BetaOffsetACK_Index3Present {
		var tmp_int_BetaOffsetACK_Index3 int64
		if tmp_int_BetaOffsetACK_Index3, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode BetaOffsetACK_Index3", err)
		}
		ie.BetaOffsetACK_Index3 = &tmp_int_BetaOffsetACK_Index3
	}
	if BetaOffsetCSI_Part1_Index1Present {
		var tmp_int_BetaOffsetCSI_Part1_Index1 int64
		if tmp_int_BetaOffsetCSI_Part1_Index1, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode BetaOffsetCSI_Part1_Index1", err)
		}
		ie.BetaOffsetCSI_Part1_Index1 = &tmp_int_BetaOffsetCSI_Part1_Index1
	}
	if BetaOffsetCSI_Part1_Index2Present {
		var tmp_int_BetaOffsetCSI_Part1_Index2 int64
		if tmp_int_BetaOffsetCSI_Part1_Index2, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode BetaOffsetCSI_Part1_Index2", err)
		}
		ie.BetaOffsetCSI_Part1_Index2 = &tmp_int_BetaOffsetCSI_Part1_Index2
	}
	if BetaOffsetCSI_Part2_Index1Present {
		var tmp_int_BetaOffsetCSI_Part2_Index1 int64
		if tmp_int_BetaOffsetCSI_Part2_Index1, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode BetaOffsetCSI_Part2_Index1", err)
		}
		ie.BetaOffsetCSI_Part2_Index1 = &tmp_int_BetaOffsetCSI_Part2_Index1
	}
	if BetaOffsetCSI_Part2_Index2Present {
		var tmp_int_BetaOffsetCSI_Part2_Index2 int64
		if tmp_int_BetaOffsetCSI_Part2_Index2, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode BetaOffsetCSI_Part2_Index2", err)
		}
		ie.BetaOffsetCSI_Part2_Index2 = &tmp_int_BetaOffsetCSI_Part2_Index2
	}
	return nil
}
