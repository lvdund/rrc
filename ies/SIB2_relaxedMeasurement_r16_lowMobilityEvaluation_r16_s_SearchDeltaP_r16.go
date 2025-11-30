package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_dB3    aper.Enumerated = 0
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_dB6    aper.Enumerated = 1
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_dB9    aper.Enumerated = 2
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_dB12   aper.Enumerated = 3
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_dB15   aper.Enumerated = 4
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_spare3 aper.Enumerated = 5
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_spare2 aper.Enumerated = 6
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_spare1 aper.Enumerated = 7
)

type SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16", err)
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
