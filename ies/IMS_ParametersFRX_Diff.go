package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IMS_ParametersFRX_Diff struct {
	VoiceOverNR *IMS_ParametersFRX_Diff_voiceOverNR `optional`
}

func (ie *IMS_ParametersFRX_Diff) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.VoiceOverNR != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.VoiceOverNR != nil {
		if err = ie.VoiceOverNR.Encode(w); err != nil {
			return utils.WrapError("Encode VoiceOverNR", err)
		}
	}
	return nil
}

func (ie *IMS_ParametersFRX_Diff) Decode(r *uper.UperReader) error {
	var err error
	var VoiceOverNRPresent bool
	if VoiceOverNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if VoiceOverNRPresent {
		ie.VoiceOverNR = new(IMS_ParametersFRX_Diff_voiceOverNR)
		if err = ie.VoiceOverNR.Decode(r); err != nil {
			return utils.WrapError("Decode VoiceOverNR", err)
		}
	}
	return nil
}
