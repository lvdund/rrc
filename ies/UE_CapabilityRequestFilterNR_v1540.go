package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterNR_v1540 struct {
	Srs_SwitchingTimeRequest *UE_CapabilityRequestFilterNR_v1540_srs_SwitchingTimeRequest `optional`
	NonCriticalExtension     *UE_CapabilityRequestFilterNR_v1710                          `optional`
}

func (ie *UE_CapabilityRequestFilterNR_v1540) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Srs_SwitchingTimeRequest != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Srs_SwitchingTimeRequest != nil {
		if err = ie.Srs_SwitchingTimeRequest.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_SwitchingTimeRequest", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterNR_v1540) Decode(r *aper.AperReader) error {
	var err error
	var Srs_SwitchingTimeRequestPresent bool
	if Srs_SwitchingTimeRequestPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Srs_SwitchingTimeRequestPresent {
		ie.Srs_SwitchingTimeRequest = new(UE_CapabilityRequestFilterNR_v1540_srs_SwitchingTimeRequest)
		if err = ie.Srs_SwitchingTimeRequest.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_SwitchingTimeRequest", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_CapabilityRequestFilterNR_v1710)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
