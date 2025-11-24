package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ChannelAccessConfig_r16 struct {
	EnergyDetectionConfig_r16           *ChannelAccessConfig_r16_energyDetectionConfig_r16       `lb:-85,ub:-52,optional`
	Ul_toDL_COT_SharingED_Threshold_r16 *int64                                                   `lb:-85,ub:-52,optional`
	AbsenceOfAnyOtherTechnology_r16     *ChannelAccessConfig_r16_absenceOfAnyOtherTechnology_r16 `optional`
}

func (ie *ChannelAccessConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.EnergyDetectionConfig_r16 != nil, ie.Ul_toDL_COT_SharingED_Threshold_r16 != nil, ie.AbsenceOfAnyOtherTechnology_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.EnergyDetectionConfig_r16 != nil {
		if err = ie.EnergyDetectionConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode EnergyDetectionConfig_r16", err)
		}
	}
	if ie.Ul_toDL_COT_SharingED_Threshold_r16 != nil {
		if err = w.WriteInteger(*ie.Ul_toDL_COT_SharingED_Threshold_r16, &uper.Constraint{Lb: -85, Ub: -52}, false); err != nil {
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

func (ie *ChannelAccessConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var EnergyDetectionConfig_r16Present bool
	if EnergyDetectionConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_toDL_COT_SharingED_Threshold_r16Present bool
	if Ul_toDL_COT_SharingED_Threshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AbsenceOfAnyOtherTechnology_r16Present bool
	if AbsenceOfAnyOtherTechnology_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if EnergyDetectionConfig_r16Present {
		ie.EnergyDetectionConfig_r16 = new(ChannelAccessConfig_r16_energyDetectionConfig_r16)
		if err = ie.EnergyDetectionConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode EnergyDetectionConfig_r16", err)
		}
	}
	if Ul_toDL_COT_SharingED_Threshold_r16Present {
		var tmp_int_Ul_toDL_COT_SharingED_Threshold_r16 int64
		if tmp_int_Ul_toDL_COT_SharingED_Threshold_r16, err = r.ReadInteger(&uper.Constraint{Lb: -85, Ub: -52}, false); err != nil {
			return utils.WrapError("Decode Ul_toDL_COT_SharingED_Threshold_r16", err)
		}
		ie.Ul_toDL_COT_SharingED_Threshold_r16 = &tmp_int_Ul_toDL_COT_SharingED_Threshold_r16
	}
	if AbsenceOfAnyOtherTechnology_r16Present {
		ie.AbsenceOfAnyOtherTechnology_r16 = new(ChannelAccessConfig_r16_absenceOfAnyOtherTechnology_r16)
		if err = ie.AbsenceOfAnyOtherTechnology_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AbsenceOfAnyOtherTechnology_r16", err)
		}
	}
	return nil
}
