package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkConfigCommonSIB struct {
	FrequencyInfoUL          FrequencyInfoUL_SIB `madatory`
	InitialUplinkBWP         BWP_UplinkCommon    `madatory`
	TimeAlignmentTimerCommon TimeAlignmentTimer  `madatory`
}

func (ie *UplinkConfigCommonSIB) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.FrequencyInfoUL.Encode(w); err != nil {
		return utils.WrapError("Encode FrequencyInfoUL", err)
	}
	if err = ie.InitialUplinkBWP.Encode(w); err != nil {
		return utils.WrapError("Encode InitialUplinkBWP", err)
	}
	if err = ie.TimeAlignmentTimerCommon.Encode(w); err != nil {
		return utils.WrapError("Encode TimeAlignmentTimerCommon", err)
	}
	return nil
}

func (ie *UplinkConfigCommonSIB) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.FrequencyInfoUL.Decode(r); err != nil {
		return utils.WrapError("Decode FrequencyInfoUL", err)
	}
	if err = ie.InitialUplinkBWP.Decode(r); err != nil {
		return utils.WrapError("Decode InitialUplinkBWP", err)
	}
	if err = ie.TimeAlignmentTimerCommon.Decode(r); err != nil {
		return utils.WrapError("Decode TimeAlignmentTimerCommon", err)
	}
	return nil
}
