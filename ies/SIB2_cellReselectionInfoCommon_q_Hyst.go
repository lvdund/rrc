package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB0  aper.Enumerated = 0
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB1  aper.Enumerated = 1
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB2  aper.Enumerated = 2
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB3  aper.Enumerated = 3
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB4  aper.Enumerated = 4
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB5  aper.Enumerated = 5
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB6  aper.Enumerated = 6
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB8  aper.Enumerated = 7
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB10 aper.Enumerated = 8
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB12 aper.Enumerated = 9
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB14 aper.Enumerated = 10
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB16 aper.Enumerated = 11
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB18 aper.Enumerated = 12
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB20 aper.Enumerated = 13
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB22 aper.Enumerated = 14
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB24 aper.Enumerated = 15
)

type SIB2_cellReselectionInfoCommon_q_Hyst struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SIB2_cellReselectionInfoCommon_q_Hyst) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SIB2_cellReselectionInfoCommon_q_Hyst", err)
	}
	return nil
}

func (ie *SIB2_cellReselectionInfoCommon_q_Hyst) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SIB2_cellReselectionInfoCommon_q_Hyst", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
