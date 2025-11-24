package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IMS_ParametersFR2_2_r17 struct {
	VoiceOverNR_r17 *IMS_ParametersFR2_2_r17_voiceOverNR_r17 `optional`
}

func (ie *IMS_ParametersFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.VoiceOverNR_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.VoiceOverNR_r17 != nil {
		if err = ie.VoiceOverNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode VoiceOverNR_r17", err)
		}
	}
	return nil
}

func (ie *IMS_ParametersFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var VoiceOverNR_r17Present bool
	if VoiceOverNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if VoiceOverNR_r17Present {
		ie.VoiceOverNR_r17 = new(IMS_ParametersFR2_2_r17_voiceOverNR_r17)
		if err = ie.VoiceOverNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode VoiceOverNR_r17", err)
		}
	}
	return nil
}
