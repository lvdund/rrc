package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersFRX_Diff_r16_drx_Adaptation_r16 struct {
	Non_SharedSpectrumChAccess_r16 *MinTimeGap_r16 `optional`
	SharedSpectrumChAccess_r16     *MinTimeGap_r16 `optional`
}

func (ie *MAC_ParametersFRX_Diff_r16_drx_Adaptation_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Non_SharedSpectrumChAccess_r16 != nil, ie.SharedSpectrumChAccess_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Non_SharedSpectrumChAccess_r16 != nil {
		if err = ie.Non_SharedSpectrumChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Non_SharedSpectrumChAccess_r16", err)
		}
	}
	if ie.SharedSpectrumChAccess_r16 != nil {
		if err = ie.SharedSpectrumChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SharedSpectrumChAccess_r16", err)
		}
	}
	return nil
}

func (ie *MAC_ParametersFRX_Diff_r16_drx_Adaptation_r16) Decode(r *aper.AperReader) error {
	var err error
	var Non_SharedSpectrumChAccess_r16Present bool
	if Non_SharedSpectrumChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SharedSpectrumChAccess_r16Present bool
	if SharedSpectrumChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Non_SharedSpectrumChAccess_r16Present {
		ie.Non_SharedSpectrumChAccess_r16 = new(MinTimeGap_r16)
		if err = ie.Non_SharedSpectrumChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Non_SharedSpectrumChAccess_r16", err)
		}
	}
	if SharedSpectrumChAccess_r16Present {
		ie.SharedSpectrumChAccess_r16 = new(MinTimeGap_r16)
		if err = ie.SharedSpectrumChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SharedSpectrumChAccess_r16", err)
		}
	}
	return nil
}
