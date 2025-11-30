package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasurementTimingConfiguration_v1550_IEs struct {
	CampOnFirstSSB       bool                                      `madatory`
	PsCellOnlyOnFirstSSB bool                                      `madatory`
	NonCriticalExtension *MeasurementTimingConfiguration_v1610_IEs `optional`
}

func (ie *MeasurementTimingConfiguration_v1550_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBoolean(ie.CampOnFirstSSB); err != nil {
		return utils.WrapError("WriteBoolean CampOnFirstSSB", err)
	}
	if err = w.WriteBoolean(ie.PsCellOnlyOnFirstSSB); err != nil {
		return utils.WrapError("WriteBoolean PsCellOnlyOnFirstSSB", err)
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *MeasurementTimingConfiguration_v1550_IEs) Decode(r *aper.AperReader) error {
	var err error
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bool_CampOnFirstSSB bool
	if tmp_bool_CampOnFirstSSB, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean CampOnFirstSSB", err)
	}
	ie.CampOnFirstSSB = tmp_bool_CampOnFirstSSB
	var tmp_bool_PsCellOnlyOnFirstSSB bool
	if tmp_bool_PsCellOnlyOnFirstSSB, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean PsCellOnlyOnFirstSSB", err)
	}
	ie.PsCellOnlyOnFirstSSB = tmp_bool_PsCellOnlyOnFirstSSB
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(MeasurementTimingConfiguration_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
