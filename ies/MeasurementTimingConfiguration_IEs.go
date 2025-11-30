package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasurementTimingConfiguration_IEs struct {
	MeasTiming           *MeasTimingList                           `optional`
	NonCriticalExtension *MeasurementTimingConfiguration_v1550_IEs `optional`
}

func (ie *MeasurementTimingConfiguration_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasTiming != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasTiming != nil {
		if err = ie.MeasTiming.Encode(w); err != nil {
			return utils.WrapError("Encode MeasTiming", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *MeasurementTimingConfiguration_IEs) Decode(r *aper.AperReader) error {
	var err error
	var MeasTimingPresent bool
	if MeasTimingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasTimingPresent {
		ie.MeasTiming = new(MeasTimingList)
		if err = ie.MeasTiming.Decode(r); err != nil {
			return utils.WrapError("Decode MeasTiming", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(MeasurementTimingConfiguration_v1550_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
