package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DummyJ struct {
	MaxEnergyDetectionThreshold_r16     int64                                   `lb:-85,ub:-52,madatory`
	EnergyDetectionThresholdOffset_r16  int64                                   `lb:-20,ub:-13,madatory`
	Ul_toDL_COT_SharingED_Threshold_r16 *int64                                  `lb:-85,ub:-52,optional`
	AbsenceOfAnyOtherTechnology_r16     *DummyJ_absenceOfAnyOtherTechnology_r16 `optional`
}

func (ie *DummyJ) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ul_toDL_COT_SharingED_Threshold_r16 != nil, ie.AbsenceOfAnyOtherTechnology_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.MaxEnergyDetectionThreshold_r16, &aper.Constraint{Lb: -85, Ub: -52}, false); err != nil {
		return utils.WrapError("WriteInteger MaxEnergyDetectionThreshold_r16", err)
	}
	if err = w.WriteInteger(ie.EnergyDetectionThresholdOffset_r16, &aper.Constraint{Lb: -20, Ub: -13}, false); err != nil {
		return utils.WrapError("WriteInteger EnergyDetectionThresholdOffset_r16", err)
	}
	if ie.Ul_toDL_COT_SharingED_Threshold_r16 != nil {
		if err = w.WriteInteger(*ie.Ul_toDL_COT_SharingED_Threshold_r16, &aper.Constraint{Lb: -85, Ub: -52}, false); err != nil {
			return utils.WrapError("Encode Ul_toDL_COT_SharingED_Threshold_r16", err)
		}
	}
	if ie.AbsenceOfAnyOtherTechnology_r16 != nil {
		if err = ie.AbsenceOfAnyOtherTechnology_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AbsenceOfAnyOtherTechnology_r16", err)
		}
	}
	return nil
}

func (ie *DummyJ) Decode(r *aper.AperReader) error {
	var err error
	var Ul_toDL_COT_SharingED_Threshold_r16Present bool
	if Ul_toDL_COT_SharingED_Threshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AbsenceOfAnyOtherTechnology_r16Present bool
	if AbsenceOfAnyOtherTechnology_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_MaxEnergyDetectionThreshold_r16 int64
	if tmp_int_MaxEnergyDetectionThreshold_r16, err = r.ReadInteger(&aper.Constraint{Lb: -85, Ub: -52}, false); err != nil {
		return utils.WrapError("ReadInteger MaxEnergyDetectionThreshold_r16", err)
	}
	ie.MaxEnergyDetectionThreshold_r16 = tmp_int_MaxEnergyDetectionThreshold_r16
	var tmp_int_EnergyDetectionThresholdOffset_r16 int64
	if tmp_int_EnergyDetectionThresholdOffset_r16, err = r.ReadInteger(&aper.Constraint{Lb: -20, Ub: -13}, false); err != nil {
		return utils.WrapError("ReadInteger EnergyDetectionThresholdOffset_r16", err)
	}
	ie.EnergyDetectionThresholdOffset_r16 = tmp_int_EnergyDetectionThresholdOffset_r16
	if Ul_toDL_COT_SharingED_Threshold_r16Present {
		var tmp_int_Ul_toDL_COT_SharingED_Threshold_r16 int64
		if tmp_int_Ul_toDL_COT_SharingED_Threshold_r16, err = r.ReadInteger(&aper.Constraint{Lb: -85, Ub: -52}, false); err != nil {
			return utils.WrapError("Decode Ul_toDL_COT_SharingED_Threshold_r16", err)
		}
		ie.Ul_toDL_COT_SharingED_Threshold_r16 = &tmp_int_Ul_toDL_COT_SharingED_Threshold_r16
	}
	if AbsenceOfAnyOtherTechnology_r16Present {
		ie.AbsenceOfAnyOtherTechnology_r16 = new(DummyJ_absenceOfAnyOtherTechnology_r16)
		if err = ie.AbsenceOfAnyOtherTechnology_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AbsenceOfAnyOtherTechnology_r16", err)
		}
	}
	return nil
}
