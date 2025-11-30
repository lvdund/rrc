package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MaxBW_Preference_r16 struct {
	ReducedMaxBW_FR1_r16 *ReducedMaxBW_FRx_r16 `optional`
	ReducedMaxBW_FR2_r16 *ReducedMaxBW_FRx_r16 `optional`
}

func (ie *MaxBW_Preference_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ReducedMaxBW_FR1_r16 != nil, ie.ReducedMaxBW_FR2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReducedMaxBW_FR1_r16 != nil {
		if err = ie.ReducedMaxBW_FR1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxBW_FR1_r16", err)
		}
	}
	if ie.ReducedMaxBW_FR2_r16 != nil {
		if err = ie.ReducedMaxBW_FR2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxBW_FR2_r16", err)
		}
	}
	return nil
}

func (ie *MaxBW_Preference_r16) Decode(r *aper.AperReader) error {
	var err error
	var ReducedMaxBW_FR1_r16Present bool
	if ReducedMaxBW_FR1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReducedMaxBW_FR2_r16Present bool
	if ReducedMaxBW_FR2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ReducedMaxBW_FR1_r16Present {
		ie.ReducedMaxBW_FR1_r16 = new(ReducedMaxBW_FRx_r16)
		if err = ie.ReducedMaxBW_FR1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxBW_FR1_r16", err)
		}
	}
	if ReducedMaxBW_FR2_r16Present {
		ie.ReducedMaxBW_FR2_r16 = new(ReducedMaxBW_FRx_r16)
		if err = ie.ReducedMaxBW_FR2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxBW_FR2_r16", err)
		}
	}
	return nil
}
