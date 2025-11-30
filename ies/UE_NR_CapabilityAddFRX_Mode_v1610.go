package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_CapabilityAddFRX_Mode_v1610 struct {
	PowSav_ParametersFRX_Diff_r16 *PowSav_ParametersFRX_Diff_r16 `optional`
	Mac_ParametersFRX_Diff_r16    *MAC_ParametersFRX_Diff_r16    `optional`
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PowSav_ParametersFRX_Diff_r16 != nil, ie.Mac_ParametersFRX_Diff_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PowSav_ParametersFRX_Diff_r16 != nil {
		if err = ie.PowSav_ParametersFRX_Diff_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PowSav_ParametersFRX_Diff_r16", err)
		}
	}
	if ie.Mac_ParametersFRX_Diff_r16 != nil {
		if err = ie.Mac_ParametersFRX_Diff_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_ParametersFRX_Diff_r16", err)
		}
	}
	return nil
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1610) Decode(r *aper.AperReader) error {
	var err error
	var PowSav_ParametersFRX_Diff_r16Present bool
	if PowSav_ParametersFRX_Diff_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_ParametersFRX_Diff_r16Present bool
	if Mac_ParametersFRX_Diff_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PowSav_ParametersFRX_Diff_r16Present {
		ie.PowSav_ParametersFRX_Diff_r16 = new(PowSav_ParametersFRX_Diff_r16)
		if err = ie.PowSav_ParametersFRX_Diff_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PowSav_ParametersFRX_Diff_r16", err)
		}
	}
	if Mac_ParametersFRX_Diff_r16Present {
		ie.Mac_ParametersFRX_Diff_r16 = new(MAC_ParametersFRX_Diff_r16)
		if err = ie.Mac_ParametersFRX_Diff_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_ParametersFRX_Diff_r16", err)
		}
	}
	return nil
}
