package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type P0AlphaSet_r17 struct {
	P0_r17              *int64                             `lb:-16,ub:15,optional`
	Alpha_r17           *Alpha                             `optional`
	ClosedLoopIndex_r17 P0AlphaSet_r17_closedLoopIndex_r17 `madatory`
}

func (ie *P0AlphaSet_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.P0_r17 != nil, ie.Alpha_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.P0_r17 != nil {
		if err = w.WriteInteger(*ie.P0_r17, &aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode P0_r17", err)
		}
	}
	if ie.Alpha_r17 != nil {
		if err = ie.Alpha_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Alpha_r17", err)
		}
	}
	if err = ie.ClosedLoopIndex_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ClosedLoopIndex_r17", err)
	}
	return nil
}

func (ie *P0AlphaSet_r17) Decode(r *aper.AperReader) error {
	var err error
	var P0_r17Present bool
	if P0_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Alpha_r17Present bool
	if Alpha_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if P0_r17Present {
		var tmp_int_P0_r17 int64
		if tmp_int_P0_r17, err = r.ReadInteger(&aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode P0_r17", err)
		}
		ie.P0_r17 = &tmp_int_P0_r17
	}
	if Alpha_r17Present {
		ie.Alpha_r17 = new(Alpha)
		if err = ie.Alpha_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Alpha_r17", err)
		}
	}
	if err = ie.ClosedLoopIndex_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ClosedLoopIndex_r17", err)
	}
	return nil
}
