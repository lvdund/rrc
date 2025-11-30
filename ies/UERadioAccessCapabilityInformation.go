package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UERadioAccessCapabilityInformation struct {
	CriticalExtensions UERadioAccessCapabilityInformation_CriticalExtensions `madatory`
}

func (ie *UERadioAccessCapabilityInformation) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CriticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode CriticalExtensions", err)
	}
	return nil
}

func (ie *UERadioAccessCapabilityInformation) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CriticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode CriticalExtensions", err)
	}
	return nil
}
