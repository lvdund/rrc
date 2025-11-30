package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RLC_ParametersSidelink_r16 struct {
	Am_WithLongSN_Sidelink_r16 *RLC_ParametersSidelink_r16_am_WithLongSN_Sidelink_r16 `optional`
	Um_WithLongSN_Sidelink_r16 *RLC_ParametersSidelink_r16_um_WithLongSN_Sidelink_r16 `optional`
}

func (ie *RLC_ParametersSidelink_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Am_WithLongSN_Sidelink_r16 != nil, ie.Um_WithLongSN_Sidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Am_WithLongSN_Sidelink_r16 != nil {
		if err = ie.Am_WithLongSN_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Am_WithLongSN_Sidelink_r16", err)
		}
	}
	if ie.Um_WithLongSN_Sidelink_r16 != nil {
		if err = ie.Um_WithLongSN_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Um_WithLongSN_Sidelink_r16", err)
		}
	}
	return nil
}

func (ie *RLC_ParametersSidelink_r16) Decode(r *aper.AperReader) error {
	var err error
	var Am_WithLongSN_Sidelink_r16Present bool
	if Am_WithLongSN_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Um_WithLongSN_Sidelink_r16Present bool
	if Um_WithLongSN_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Am_WithLongSN_Sidelink_r16Present {
		ie.Am_WithLongSN_Sidelink_r16 = new(RLC_ParametersSidelink_r16_am_WithLongSN_Sidelink_r16)
		if err = ie.Am_WithLongSN_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Am_WithLongSN_Sidelink_r16", err)
		}
	}
	if Um_WithLongSN_Sidelink_r16Present {
		ie.Um_WithLongSN_Sidelink_r16 = new(RLC_ParametersSidelink_r16_um_WithLongSN_Sidelink_r16)
		if err = ie.Um_WithLongSN_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Um_WithLongSN_Sidelink_r16", err)
		}
	}
	return nil
}
