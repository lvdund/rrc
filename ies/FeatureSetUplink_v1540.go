package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1540 struct {
	ZeroSlotOffsetAperiodicSRS        *FeatureSetUplink_v1540_zeroSlotOffsetAperiodicSRS        `optional`
	Pa_PhaseDiscontinuityImpacts      *FeatureSetUplink_v1540_pa_PhaseDiscontinuityImpacts      `optional`
	Pusch_SeparationWithGap           *FeatureSetUplink_v1540_pusch_SeparationWithGap           `optional`
	Pusch_ProcessingType2             *FeatureSetUplink_v1540_pusch_ProcessingType2             `optional`
	Ul_MCS_TableAlt_DynamicIndication *FeatureSetUplink_v1540_ul_MCS_TableAlt_DynamicIndication `optional`
}

func (ie *FeatureSetUplink_v1540) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ZeroSlotOffsetAperiodicSRS != nil, ie.Pa_PhaseDiscontinuityImpacts != nil, ie.Pusch_SeparationWithGap != nil, ie.Pusch_ProcessingType2 != nil, ie.Ul_MCS_TableAlt_DynamicIndication != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ZeroSlotOffsetAperiodicSRS != nil {
		if err = ie.ZeroSlotOffsetAperiodicSRS.Encode(w); err != nil {
			return utils.WrapError("Encode ZeroSlotOffsetAperiodicSRS", err)
		}
	}
	if ie.Pa_PhaseDiscontinuityImpacts != nil {
		if err = ie.Pa_PhaseDiscontinuityImpacts.Encode(w); err != nil {
			return utils.WrapError("Encode Pa_PhaseDiscontinuityImpacts", err)
		}
	}
	if ie.Pusch_SeparationWithGap != nil {
		if err = ie.Pusch_SeparationWithGap.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_SeparationWithGap", err)
		}
	}
	if ie.Pusch_ProcessingType2 != nil {
		if err = ie.Pusch_ProcessingType2.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_ProcessingType2", err)
		}
	}
	if ie.Ul_MCS_TableAlt_DynamicIndication != nil {
		if err = ie.Ul_MCS_TableAlt_DynamicIndication.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_MCS_TableAlt_DynamicIndication", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1540) Decode(r *aper.AperReader) error {
	var err error
	var ZeroSlotOffsetAperiodicSRSPresent bool
	if ZeroSlotOffsetAperiodicSRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pa_PhaseDiscontinuityImpactsPresent bool
	if Pa_PhaseDiscontinuityImpactsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_SeparationWithGapPresent bool
	if Pusch_SeparationWithGapPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_ProcessingType2Present bool
	if Pusch_ProcessingType2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_MCS_TableAlt_DynamicIndicationPresent bool
	if Ul_MCS_TableAlt_DynamicIndicationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ZeroSlotOffsetAperiodicSRSPresent {
		ie.ZeroSlotOffsetAperiodicSRS = new(FeatureSetUplink_v1540_zeroSlotOffsetAperiodicSRS)
		if err = ie.ZeroSlotOffsetAperiodicSRS.Decode(r); err != nil {
			return utils.WrapError("Decode ZeroSlotOffsetAperiodicSRS", err)
		}
	}
	if Pa_PhaseDiscontinuityImpactsPresent {
		ie.Pa_PhaseDiscontinuityImpacts = new(FeatureSetUplink_v1540_pa_PhaseDiscontinuityImpacts)
		if err = ie.Pa_PhaseDiscontinuityImpacts.Decode(r); err != nil {
			return utils.WrapError("Decode Pa_PhaseDiscontinuityImpacts", err)
		}
	}
	if Pusch_SeparationWithGapPresent {
		ie.Pusch_SeparationWithGap = new(FeatureSetUplink_v1540_pusch_SeparationWithGap)
		if err = ie.Pusch_SeparationWithGap.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_SeparationWithGap", err)
		}
	}
	if Pusch_ProcessingType2Present {
		ie.Pusch_ProcessingType2 = new(FeatureSetUplink_v1540_pusch_ProcessingType2)
		if err = ie.Pusch_ProcessingType2.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_ProcessingType2", err)
		}
	}
	if Ul_MCS_TableAlt_DynamicIndicationPresent {
		ie.Ul_MCS_TableAlt_DynamicIndication = new(FeatureSetUplink_v1540_ul_MCS_TableAlt_DynamicIndication)
		if err = ie.Ul_MCS_TableAlt_DynamicIndication.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_MCS_TableAlt_DynamicIndication", err)
		}
	}
	return nil
}
