package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_v1610_IEs struct {
	VoiceFallbackIndication_r16 *RRCRelease_v1610_IEs_voiceFallbackIndication_r16 `optional`
	MeasIdleConfig_r16          *MeasIdleConfigDedicated_r16                      `optional,setuprelease`
	NonCriticalExtension        *RRCRelease_v1650_IEs                             `optional`
}

func (ie *RRCRelease_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.VoiceFallbackIndication_r16 != nil, ie.MeasIdleConfig_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.VoiceFallbackIndication_r16 != nil {
		if err = ie.VoiceFallbackIndication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode VoiceFallbackIndication_r16", err)
		}
	}
	if ie.MeasIdleConfig_r16 != nil {
		tmp_MeasIdleConfig_r16 := utils.SetupRelease[*MeasIdleConfigDedicated_r16]{
			Setup: ie.MeasIdleConfig_r16,
		}
		if err = tmp_MeasIdleConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasIdleConfig_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCRelease_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var VoiceFallbackIndication_r16Present bool
	if VoiceFallbackIndication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasIdleConfig_r16Present bool
	if MeasIdleConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if VoiceFallbackIndication_r16Present {
		ie.VoiceFallbackIndication_r16 = new(RRCRelease_v1610_IEs_voiceFallbackIndication_r16)
		if err = ie.VoiceFallbackIndication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode VoiceFallbackIndication_r16", err)
		}
	}
	if MeasIdleConfig_r16Present {
		tmp_MeasIdleConfig_r16 := utils.SetupRelease[*MeasIdleConfigDedicated_r16]{}
		if err = tmp_MeasIdleConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasIdleConfig_r16", err)
		}
		ie.MeasIdleConfig_r16 = tmp_MeasIdleConfig_r16.Setup
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCRelease_v1650_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
