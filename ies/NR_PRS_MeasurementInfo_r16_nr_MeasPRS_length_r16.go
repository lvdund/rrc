package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms1dot5 aper.Enumerated = 0
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms3     aper.Enumerated = 1
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms3dot5 aper.Enumerated = 2
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms4     aper.Enumerated = 3
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms5dot5 aper.Enumerated = 4
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms6     aper.Enumerated = 5
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms10    aper.Enumerated = 6
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms20    aper.Enumerated = 7
)

type NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16", err)
	}
	return nil
}

func (ie *NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
