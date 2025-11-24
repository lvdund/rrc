package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkConfigCommon struct {
	FrequencyInfoUL  *FrequencyInfoUL   `optional`
	InitialUplinkBWP *BWP_UplinkCommon  `optional`
	Dummy            TimeAlignmentTimer `madatory`
}

func (ie *UplinkConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.FrequencyInfoUL != nil, ie.InitialUplinkBWP != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FrequencyInfoUL != nil {
		if err = ie.FrequencyInfoUL.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyInfoUL", err)
		}
	}
	if ie.InitialUplinkBWP != nil {
		if err = ie.InitialUplinkBWP.Encode(w); err != nil {
			return utils.WrapError("Encode InitialUplinkBWP", err)
		}
	}
	if err = ie.Dummy.Encode(w); err != nil {
		return utils.WrapError("Encode Dummy", err)
	}
	return nil
}

func (ie *UplinkConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var FrequencyInfoULPresent bool
	if FrequencyInfoULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var InitialUplinkBWPPresent bool
	if InitialUplinkBWPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FrequencyInfoULPresent {
		ie.FrequencyInfoUL = new(FrequencyInfoUL)
		if err = ie.FrequencyInfoUL.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyInfoUL", err)
		}
	}
	if InitialUplinkBWPPresent {
		ie.InitialUplinkBWP = new(BWP_UplinkCommon)
		if err = ie.InitialUplinkBWP.Decode(r); err != nil {
			return utils.WrapError("Decode InitialUplinkBWP", err)
		}
	}
	if err = ie.Dummy.Decode(r); err != nil {
		return utils.WrapError("Decode Dummy", err)
	}
	return nil
}
