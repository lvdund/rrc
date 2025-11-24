package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReleasePreferenceConfig_r16 struct {
	ReleasePreferenceProhibitTimer_r16 ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16 `madatory`
	ConnectedReporting                 *ReleasePreferenceConfig_r16_connectedReporting                `optional`
}

func (ie *ReleasePreferenceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ConnectedReporting != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ReleasePreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReleasePreferenceProhibitTimer_r16", err)
	}
	if ie.ConnectedReporting != nil {
		if err = ie.ConnectedReporting.Encode(w); err != nil {
			return utils.WrapError("Encode ConnectedReporting", err)
		}
	}
	return nil
}

func (ie *ReleasePreferenceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var ConnectedReportingPresent bool
	if ConnectedReportingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ReleasePreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReleasePreferenceProhibitTimer_r16", err)
	}
	if ConnectedReportingPresent {
		ie.ConnectedReporting = new(ReleasePreferenceConfig_r16_connectedReporting)
		if err = ie.ConnectedReporting.Decode(r); err != nil {
			return utils.WrapError("Decode ConnectedReporting", err)
		}
	}
	return nil
}
