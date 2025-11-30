package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1650 struct {
	MpsPriorityIndication_r16 *UE_NR_Capability_v1650_mpsPriorityIndication_r16 `optional`
	HighSpeedParameters_v1650 *HighSpeedParameters_v1650                        `optional`
	NonCriticalExtension      *UE_NR_Capability_v1690                           `optional`
}

func (ie *UE_NR_Capability_v1650) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MpsPriorityIndication_r16 != nil, ie.HighSpeedParameters_v1650 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MpsPriorityIndication_r16 != nil {
		if err = ie.MpsPriorityIndication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MpsPriorityIndication_r16", err)
		}
	}
	if ie.HighSpeedParameters_v1650 != nil {
		if err = ie.HighSpeedParameters_v1650.Encode(w); err != nil {
			return utils.WrapError("Encode HighSpeedParameters_v1650", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1650) Decode(r *aper.AperReader) error {
	var err error
	var MpsPriorityIndication_r16Present bool
	if MpsPriorityIndication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HighSpeedParameters_v1650Present bool
	if HighSpeedParameters_v1650Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MpsPriorityIndication_r16Present {
		ie.MpsPriorityIndication_r16 = new(UE_NR_Capability_v1650_mpsPriorityIndication_r16)
		if err = ie.MpsPriorityIndication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MpsPriorityIndication_r16", err)
		}
	}
	if HighSpeedParameters_v1650Present {
		ie.HighSpeedParameters_v1650 = new(HighSpeedParameters_v1650)
		if err = ie.HighSpeedParameters_v1650.Decode(r); err != nil {
			return utils.WrapError("Decode HighSpeedParameters_v1650", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v1690)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
