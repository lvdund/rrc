package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16 struct {
	TimeDurationForCI_r16    *CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeDurationForCI_r16   `optional`
	TimeGranularityForCI_r16 CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16 `madatory`
	FrequencyRegionForCI_r16 int64                                                                               `lb:0,ub:37949,madatory`
	DeltaOffset_r16          int64                                                                               `lb:0,ub:2,madatory`
}

func (ie *CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.TimeDurationForCI_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.TimeDurationForCI_r16 != nil {
		if err = ie.TimeDurationForCI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TimeDurationForCI_r16", err)
		}
	}
	if err = ie.TimeGranularityForCI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TimeGranularityForCI_r16", err)
	}
	if err = w.WriteInteger(ie.FrequencyRegionForCI_r16, &aper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("WriteInteger FrequencyRegionForCI_r16", err)
	}
	if err = w.WriteInteger(ie.DeltaOffset_r16, &aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger DeltaOffset_r16", err)
	}
	return nil
}

func (ie *CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16) Decode(r *aper.AperReader) error {
	var err error
	var TimeDurationForCI_r16Present bool
	if TimeDurationForCI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if TimeDurationForCI_r16Present {
		ie.TimeDurationForCI_r16 = new(CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeDurationForCI_r16)
		if err = ie.TimeDurationForCI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TimeDurationForCI_r16", err)
		}
	}
	if err = ie.TimeGranularityForCI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TimeGranularityForCI_r16", err)
	}
	var tmp_int_FrequencyRegionForCI_r16 int64
	if tmp_int_FrequencyRegionForCI_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("ReadInteger FrequencyRegionForCI_r16", err)
	}
	ie.FrequencyRegionForCI_r16 = tmp_int_FrequencyRegionForCI_r16
	var tmp_int_DeltaOffset_r16 int64
	if tmp_int_DeltaOffset_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger DeltaOffset_r16", err)
	}
	ie.DeltaOffset_r16 = tmp_int_DeltaOffset_r16
	return nil
}
