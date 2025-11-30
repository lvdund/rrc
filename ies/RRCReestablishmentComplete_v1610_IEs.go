package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReestablishmentComplete_v1610_IEs struct {
	Ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16 `optional`
	NonCriticalExtension         interface{}                   `optional`
}

func (ie *RRCReestablishmentComplete_v1610_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ue_MeasurementsAvailable_r16 != nil}
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
	return nil
}

func (ie *RRCReestablishmentComplete_v1610_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Ue_MeasurementsAvailable_r16Present bool
	if Ue_MeasurementsAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ue_MeasurementsAvailable_r16Present {
		ie.Ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
		if err = ie.Ue_MeasurementsAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_MeasurementsAvailable_r16", err)
		}
	}
	return nil
}
