package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type P0_PUSCH_AlphaSet struct {
	P0_PUSCH_AlphaSetId P0_PUSCH_AlphaSetId `madatory`
	P0                  *int64              `lb:-16,ub:15,optional`
	Alpha               *Alpha              `optional`
}

func (ie *P0_PUSCH_AlphaSet) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.P0 != nil, ie.Alpha != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.P0_PUSCH_AlphaSetId.Encode(w); err != nil {
		return utils.WrapError("Encode P0_PUSCH_AlphaSetId", err)
	}
	if ie.P0 != nil {
		if err = w.WriteInteger(*ie.P0, &aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode P0", err)
		}
	}
	if ie.Alpha != nil {
		if err = ie.Alpha.Encode(w); err != nil {
			return utils.WrapError("Encode Alpha", err)
		}
	}
	return nil
}

func (ie *P0_PUSCH_AlphaSet) Decode(r *aper.AperReader) error {
	var err error
	var P0Present bool
	if P0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AlphaPresent bool
	if AlphaPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.P0_PUSCH_AlphaSetId.Decode(r); err != nil {
		return utils.WrapError("Decode P0_PUSCH_AlphaSetId", err)
	}
	if P0Present {
		var tmp_int_P0 int64
		if tmp_int_P0, err = r.ReadInteger(&aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode P0", err)
		}
		ie.P0 = &tmp_int_P0
	}
	if AlphaPresent {
		ie.Alpha = new(Alpha)
		if err = ie.Alpha.Decode(r); err != nil {
			return utils.WrapError("Decode Alpha", err)
		}
	}
	return nil
}
