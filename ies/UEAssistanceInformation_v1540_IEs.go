package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformation_v1540_IEs struct {
	OverheatingAssistance *OverheatingAssistance             `optional`
	NonCriticalExtension  *UEAssistanceInformation_v1610_IEs `optional`
}

func (ie *UEAssistanceInformation_v1540_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.OverheatingAssistance != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OverheatingAssistance != nil {
		if err = ie.OverheatingAssistance.Encode(w); err != nil {
			return utils.WrapError("Encode OverheatingAssistance", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformation_v1540_IEs) Decode(r *aper.AperReader) error {
	var err error
	var OverheatingAssistancePresent bool
	if OverheatingAssistancePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if OverheatingAssistancePresent {
		ie.OverheatingAssistance = new(OverheatingAssistance)
		if err = ie.OverheatingAssistance.Decode(r); err != nil {
			return utils.WrapError("Decode OverheatingAssistance", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UEAssistanceInformation_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
