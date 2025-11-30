package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersFRX_Diff_r16 struct {
	DirectMCG_SCellActivation_r16       *MAC_ParametersFRX_Diff_r16_directMCG_SCellActivation_r16       `optional`
	DirectMCG_SCellActivationResume_r16 *MAC_ParametersFRX_Diff_r16_directMCG_SCellActivationResume_r16 `optional`
	DirectSCG_SCellActivation_r16       *MAC_ParametersFRX_Diff_r16_directSCG_SCellActivation_r16       `optional`
	DirectSCG_SCellActivationResume_r16 *MAC_ParametersFRX_Diff_r16_directSCG_SCellActivationResume_r16 `optional`
	Drx_Adaptation_r16                  *MAC_ParametersFRX_Diff_r16_drx_Adaptation_r16                  `optional`
}

func (ie *MAC_ParametersFRX_Diff_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DirectMCG_SCellActivation_r16 != nil, ie.DirectMCG_SCellActivationResume_r16 != nil, ie.DirectSCG_SCellActivation_r16 != nil, ie.DirectSCG_SCellActivationResume_r16 != nil, ie.Drx_Adaptation_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DirectMCG_SCellActivation_r16 != nil {
		if err = ie.DirectMCG_SCellActivation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DirectMCG_SCellActivation_r16", err)
		}
	}
	if ie.DirectMCG_SCellActivationResume_r16 != nil {
		if err = ie.DirectMCG_SCellActivationResume_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DirectMCG_SCellActivationResume_r16", err)
		}
	}
	if ie.DirectSCG_SCellActivation_r16 != nil {
		if err = ie.DirectSCG_SCellActivation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DirectSCG_SCellActivation_r16", err)
		}
	}
	if ie.DirectSCG_SCellActivationResume_r16 != nil {
		if err = ie.DirectSCG_SCellActivationResume_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DirectSCG_SCellActivationResume_r16", err)
		}
	}
	if ie.Drx_Adaptation_r16 != nil {
		if err = ie.Drx_Adaptation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_Adaptation_r16", err)
		}
	}
	return nil
}

func (ie *MAC_ParametersFRX_Diff_r16) Decode(r *aper.AperReader) error {
	var err error
	var DirectMCG_SCellActivation_r16Present bool
	if DirectMCG_SCellActivation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DirectMCG_SCellActivationResume_r16Present bool
	if DirectMCG_SCellActivationResume_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DirectSCG_SCellActivation_r16Present bool
	if DirectSCG_SCellActivation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DirectSCG_SCellActivationResume_r16Present bool
	if DirectSCG_SCellActivationResume_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_Adaptation_r16Present bool
	if Drx_Adaptation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if DirectMCG_SCellActivation_r16Present {
		ie.DirectMCG_SCellActivation_r16 = new(MAC_ParametersFRX_Diff_r16_directMCG_SCellActivation_r16)
		if err = ie.DirectMCG_SCellActivation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DirectMCG_SCellActivation_r16", err)
		}
	}
	if DirectMCG_SCellActivationResume_r16Present {
		ie.DirectMCG_SCellActivationResume_r16 = new(MAC_ParametersFRX_Diff_r16_directMCG_SCellActivationResume_r16)
		if err = ie.DirectMCG_SCellActivationResume_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DirectMCG_SCellActivationResume_r16", err)
		}
	}
	if DirectSCG_SCellActivation_r16Present {
		ie.DirectSCG_SCellActivation_r16 = new(MAC_ParametersFRX_Diff_r16_directSCG_SCellActivation_r16)
		if err = ie.DirectSCG_SCellActivation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DirectSCG_SCellActivation_r16", err)
		}
	}
	if DirectSCG_SCellActivationResume_r16Present {
		ie.DirectSCG_SCellActivationResume_r16 = new(MAC_ParametersFRX_Diff_r16_directSCG_SCellActivationResume_r16)
		if err = ie.DirectSCG_SCellActivationResume_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DirectSCG_SCellActivationResume_r16", err)
		}
	}
	if Drx_Adaptation_r16Present {
		ie.Drx_Adaptation_r16 = new(MAC_ParametersFRX_Diff_r16_drx_Adaptation_r16)
		if err = ie.Drx_Adaptation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_Adaptation_r16", err)
		}
	}
	return nil
}
