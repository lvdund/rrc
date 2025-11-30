package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_s5     aper.Enumerated = 0
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_s10    aper.Enumerated = 1
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_s20    aper.Enumerated = 2
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_s30    aper.Enumerated = 3
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_s60    aper.Enumerated = 4
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_s120   aper.Enumerated = 5
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_s180   aper.Enumerated = 6
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_s240   aper.Enumerated = 7
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_s300   aper.Enumerated = 8
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_spare7 aper.Enumerated = 9
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_spare6 aper.Enumerated = 10
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_spare5 aper.Enumerated = 11
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_spare4 aper.Enumerated = 12
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_spare3 aper.Enumerated = 13
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_spare2 aper.Enumerated = 14
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16_Enum_spare1 aper.Enumerated = 15
)

type SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16", err)
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
