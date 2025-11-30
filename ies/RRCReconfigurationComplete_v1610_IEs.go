package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationComplete_v1610_IEs struct {
	Ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16         `optional`
	NeedForGapsInfoNR_r16        *NeedForGapsInfoNR_r16                `optional`
	NonCriticalExtension         *RRCReconfigurationComplete_v1640_IEs `optional`
}

func (ie *RRCReconfigurationComplete_v1610_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ue_MeasurementsAvailable_r16 != nil, ie.NeedForGapsInfoNR_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ue_MeasurementsAvailable_r16 != nil {
		if err = ie.Ue_MeasurementsAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_MeasurementsAvailable_r16", err)
		}
	}
	if ie.NeedForGapsInfoNR_r16 != nil {
		if err = ie.NeedForGapsInfoNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGapsInfoNR_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationComplete_v1610_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Ue_MeasurementsAvailable_r16Present bool
	if Ue_MeasurementsAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeedForGapsInfoNR_r16Present bool
	if NeedForGapsInfoNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ue_MeasurementsAvailable_r16Present {
		ie.Ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
		if err = ie.Ue_MeasurementsAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_MeasurementsAvailable_r16", err)
		}
	}
	if NeedForGapsInfoNR_r16Present {
		ie.NeedForGapsInfoNR_r16 = new(NeedForGapsInfoNR_r16)
		if err = ie.NeedForGapsInfoNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGapsInfoNR_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1640_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
