package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersFR2_2_r17 struct {
	DirectMCG_SCellActivation_r17       *MAC_ParametersFR2_2_r17_directMCG_SCellActivation_r17       `optional`
	DirectMCG_SCellActivationResume_r17 *MAC_ParametersFR2_2_r17_directMCG_SCellActivationResume_r17 `optional`
	DirectSCG_SCellActivation_r17       *MAC_ParametersFR2_2_r17_directSCG_SCellActivation_r17       `optional`
	DirectSCG_SCellActivationResume_r17 *MAC_ParametersFR2_2_r17_directSCG_SCellActivationResume_r17 `optional`
	Drx_Adaptation_r17                  *MAC_ParametersFR2_2_r17_drx_Adaptation_r17                  `optional`
}

func (ie *MAC_ParametersFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.DirectMCG_SCellActivation_r17 != nil, ie.DirectMCG_SCellActivationResume_r17 != nil, ie.DirectSCG_SCellActivation_r17 != nil, ie.DirectSCG_SCellActivationResume_r17 != nil, ie.Drx_Adaptation_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DirectMCG_SCellActivation_r17 != nil {
		if err = ie.DirectMCG_SCellActivation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DirectMCG_SCellActivation_r17", err)
		}
	}
	if ie.DirectMCG_SCellActivationResume_r17 != nil {
		if err = ie.DirectMCG_SCellActivationResume_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DirectMCG_SCellActivationResume_r17", err)
		}
	}
	if ie.DirectSCG_SCellActivation_r17 != nil {
		if err = ie.DirectSCG_SCellActivation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DirectSCG_SCellActivation_r17", err)
		}
	}
	if ie.DirectSCG_SCellActivationResume_r17 != nil {
		if err = ie.DirectSCG_SCellActivationResume_r17.Encode(w); err != nil {
			return utils.WrapError("Encode DirectSCG_SCellActivationResume_r17", err)
		}
	}
	if ie.Drx_Adaptation_r17 != nil {
		if err = ie.Drx_Adaptation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_Adaptation_r17", err)
		}
	}
	return nil
}

func (ie *MAC_ParametersFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var DirectMCG_SCellActivation_r17Present bool
	if DirectMCG_SCellActivation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DirectMCG_SCellActivationResume_r17Present bool
	if DirectMCG_SCellActivationResume_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DirectSCG_SCellActivation_r17Present bool
	if DirectSCG_SCellActivation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DirectSCG_SCellActivationResume_r17Present bool
	if DirectSCG_SCellActivationResume_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_Adaptation_r17Present bool
	if Drx_Adaptation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if DirectMCG_SCellActivation_r17Present {
		ie.DirectMCG_SCellActivation_r17 = new(MAC_ParametersFR2_2_r17_directMCG_SCellActivation_r17)
		if err = ie.DirectMCG_SCellActivation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DirectMCG_SCellActivation_r17", err)
		}
	}
	if DirectMCG_SCellActivationResume_r17Present {
		ie.DirectMCG_SCellActivationResume_r17 = new(MAC_ParametersFR2_2_r17_directMCG_SCellActivationResume_r17)
		if err = ie.DirectMCG_SCellActivationResume_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DirectMCG_SCellActivationResume_r17", err)
		}
	}
	if DirectSCG_SCellActivation_r17Present {
		ie.DirectSCG_SCellActivation_r17 = new(MAC_ParametersFR2_2_r17_directSCG_SCellActivation_r17)
		if err = ie.DirectSCG_SCellActivation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DirectSCG_SCellActivation_r17", err)
		}
	}
	if DirectSCG_SCellActivationResume_r17Present {
		ie.DirectSCG_SCellActivationResume_r17 = new(MAC_ParametersFR2_2_r17_directSCG_SCellActivationResume_r17)
		if err = ie.DirectSCG_SCellActivationResume_r17.Decode(r); err != nil {
			return utils.WrapError("Decode DirectSCG_SCellActivationResume_r17", err)
		}
	}
	if Drx_Adaptation_r17Present {
		ie.Drx_Adaptation_r17 = new(MAC_ParametersFR2_2_r17_drx_Adaptation_r17)
		if err = ie.Drx_Adaptation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_Adaptation_r17", err)
		}
	}
	return nil
}
