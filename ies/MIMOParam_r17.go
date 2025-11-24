package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMOParam_r17 struct {
	AdditionalPCI_ToAddModList_r17       []SSB_MTC_AdditionalPCI_r17             `lb:1,ub:maxNrofAdditionalPCI_r17,optional`
	AdditionalPCI_ToReleaseList_r17      []AdditionalPCIIndex_r17                `lb:1,ub:maxNrofAdditionalPCI_r17,optional`
	UnifiedTCI_StateType_r17             *MIMOParam_r17_unifiedTCI_StateType_r17 `optional`
	Uplink_PowerControlToAddModList_r17  []Uplink_powerControl_r17               `lb:1,ub:maxUL_TCI_r17,optional`
	Uplink_PowerControlToReleaseList_r17 []Uplink_powerControlId_r17             `lb:1,ub:maxUL_TCI_r17,optional`
	SfnSchemePDCCH_r17                   *MIMOParam_r17_sfnSchemePDCCH_r17       `optional`
	SfnSchemePDSCH_r17                   *MIMOParam_r17_sfnSchemePDSCH_r17       `optional`
}

func (ie *MIMOParam_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.AdditionalPCI_ToAddModList_r17) > 0, len(ie.AdditionalPCI_ToReleaseList_r17) > 0, ie.UnifiedTCI_StateType_r17 != nil, len(ie.Uplink_PowerControlToAddModList_r17) > 0, len(ie.Uplink_PowerControlToReleaseList_r17) > 0, ie.SfnSchemePDCCH_r17 != nil, ie.SfnSchemePDSCH_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.AdditionalPCI_ToAddModList_r17) > 0 {
		tmp_AdditionalPCI_ToAddModList_r17 := utils.NewSequence[*SSB_MTC_AdditionalPCI_r17]([]*SSB_MTC_AdditionalPCI_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAdditionalPCI_r17}, false)
		for _, i := range ie.AdditionalPCI_ToAddModList_r17 {
			tmp_AdditionalPCI_ToAddModList_r17.Value = append(tmp_AdditionalPCI_ToAddModList_r17.Value, &i)
		}
		if err = tmp_AdditionalPCI_ToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalPCI_ToAddModList_r17", err)
		}
	}
	if len(ie.AdditionalPCI_ToReleaseList_r17) > 0 {
		tmp_AdditionalPCI_ToReleaseList_r17 := utils.NewSequence[*AdditionalPCIIndex_r17]([]*AdditionalPCIIndex_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAdditionalPCI_r17}, false)
		for _, i := range ie.AdditionalPCI_ToReleaseList_r17 {
			tmp_AdditionalPCI_ToReleaseList_r17.Value = append(tmp_AdditionalPCI_ToReleaseList_r17.Value, &i)
		}
		if err = tmp_AdditionalPCI_ToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalPCI_ToReleaseList_r17", err)
		}
	}
	if ie.UnifiedTCI_StateType_r17 != nil {
		if err = ie.UnifiedTCI_StateType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode UnifiedTCI_StateType_r17", err)
		}
	}
	if len(ie.Uplink_PowerControlToAddModList_r17) > 0 {
		tmp_Uplink_PowerControlToAddModList_r17 := utils.NewSequence[*Uplink_powerControl_r17]([]*Uplink_powerControl_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		for _, i := range ie.Uplink_PowerControlToAddModList_r17 {
			tmp_Uplink_PowerControlToAddModList_r17.Value = append(tmp_Uplink_PowerControlToAddModList_r17.Value, &i)
		}
		if err = tmp_Uplink_PowerControlToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Uplink_PowerControlToAddModList_r17", err)
		}
	}
	if len(ie.Uplink_PowerControlToReleaseList_r17) > 0 {
		tmp_Uplink_PowerControlToReleaseList_r17 := utils.NewSequence[*Uplink_powerControlId_r17]([]*Uplink_powerControlId_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		for _, i := range ie.Uplink_PowerControlToReleaseList_r17 {
			tmp_Uplink_PowerControlToReleaseList_r17.Value = append(tmp_Uplink_PowerControlToReleaseList_r17.Value, &i)
		}
		if err = tmp_Uplink_PowerControlToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Uplink_PowerControlToReleaseList_r17", err)
		}
	}
	if ie.SfnSchemePDCCH_r17 != nil {
		if err = ie.SfnSchemePDCCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SfnSchemePDCCH_r17", err)
		}
	}
	if ie.SfnSchemePDSCH_r17 != nil {
		if err = ie.SfnSchemePDSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SfnSchemePDSCH_r17", err)
		}
	}
	return nil
}

func (ie *MIMOParam_r17) Decode(r *uper.UperReader) error {
	var err error
	var AdditionalPCI_ToAddModList_r17Present bool
	if AdditionalPCI_ToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AdditionalPCI_ToReleaseList_r17Present bool
	if AdditionalPCI_ToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var UnifiedTCI_StateType_r17Present bool
	if UnifiedTCI_StateType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Uplink_PowerControlToAddModList_r17Present bool
	if Uplink_PowerControlToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Uplink_PowerControlToReleaseList_r17Present bool
	if Uplink_PowerControlToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SfnSchemePDCCH_r17Present bool
	if SfnSchemePDCCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SfnSchemePDSCH_r17Present bool
	if SfnSchemePDSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if AdditionalPCI_ToAddModList_r17Present {
		tmp_AdditionalPCI_ToAddModList_r17 := utils.NewSequence[*SSB_MTC_AdditionalPCI_r17]([]*SSB_MTC_AdditionalPCI_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAdditionalPCI_r17}, false)
		fn_AdditionalPCI_ToAddModList_r17 := func() *SSB_MTC_AdditionalPCI_r17 {
			return new(SSB_MTC_AdditionalPCI_r17)
		}
		if err = tmp_AdditionalPCI_ToAddModList_r17.Decode(r, fn_AdditionalPCI_ToAddModList_r17); err != nil {
			return utils.WrapError("Decode AdditionalPCI_ToAddModList_r17", err)
		}
		ie.AdditionalPCI_ToAddModList_r17 = []SSB_MTC_AdditionalPCI_r17{}
		for _, i := range tmp_AdditionalPCI_ToAddModList_r17.Value {
			ie.AdditionalPCI_ToAddModList_r17 = append(ie.AdditionalPCI_ToAddModList_r17, *i)
		}
	}
	if AdditionalPCI_ToReleaseList_r17Present {
		tmp_AdditionalPCI_ToReleaseList_r17 := utils.NewSequence[*AdditionalPCIIndex_r17]([]*AdditionalPCIIndex_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAdditionalPCI_r17}, false)
		fn_AdditionalPCI_ToReleaseList_r17 := func() *AdditionalPCIIndex_r17 {
			return new(AdditionalPCIIndex_r17)
		}
		if err = tmp_AdditionalPCI_ToReleaseList_r17.Decode(r, fn_AdditionalPCI_ToReleaseList_r17); err != nil {
			return utils.WrapError("Decode AdditionalPCI_ToReleaseList_r17", err)
		}
		ie.AdditionalPCI_ToReleaseList_r17 = []AdditionalPCIIndex_r17{}
		for _, i := range tmp_AdditionalPCI_ToReleaseList_r17.Value {
			ie.AdditionalPCI_ToReleaseList_r17 = append(ie.AdditionalPCI_ToReleaseList_r17, *i)
		}
	}
	if UnifiedTCI_StateType_r17Present {
		ie.UnifiedTCI_StateType_r17 = new(MIMOParam_r17_unifiedTCI_StateType_r17)
		if err = ie.UnifiedTCI_StateType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UnifiedTCI_StateType_r17", err)
		}
	}
	if Uplink_PowerControlToAddModList_r17Present {
		tmp_Uplink_PowerControlToAddModList_r17 := utils.NewSequence[*Uplink_powerControl_r17]([]*Uplink_powerControl_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		fn_Uplink_PowerControlToAddModList_r17 := func() *Uplink_powerControl_r17 {
			return new(Uplink_powerControl_r17)
		}
		if err = tmp_Uplink_PowerControlToAddModList_r17.Decode(r, fn_Uplink_PowerControlToAddModList_r17); err != nil {
			return utils.WrapError("Decode Uplink_PowerControlToAddModList_r17", err)
		}
		ie.Uplink_PowerControlToAddModList_r17 = []Uplink_powerControl_r17{}
		for _, i := range tmp_Uplink_PowerControlToAddModList_r17.Value {
			ie.Uplink_PowerControlToAddModList_r17 = append(ie.Uplink_PowerControlToAddModList_r17, *i)
		}
	}
	if Uplink_PowerControlToReleaseList_r17Present {
		tmp_Uplink_PowerControlToReleaseList_r17 := utils.NewSequence[*Uplink_powerControlId_r17]([]*Uplink_powerControlId_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		fn_Uplink_PowerControlToReleaseList_r17 := func() *Uplink_powerControlId_r17 {
			return new(Uplink_powerControlId_r17)
		}
		if err = tmp_Uplink_PowerControlToReleaseList_r17.Decode(r, fn_Uplink_PowerControlToReleaseList_r17); err != nil {
			return utils.WrapError("Decode Uplink_PowerControlToReleaseList_r17", err)
		}
		ie.Uplink_PowerControlToReleaseList_r17 = []Uplink_powerControlId_r17{}
		for _, i := range tmp_Uplink_PowerControlToReleaseList_r17.Value {
			ie.Uplink_PowerControlToReleaseList_r17 = append(ie.Uplink_PowerControlToReleaseList_r17, *i)
		}
	}
	if SfnSchemePDCCH_r17Present {
		ie.SfnSchemePDCCH_r17 = new(MIMOParam_r17_sfnSchemePDCCH_r17)
		if err = ie.SfnSchemePDCCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SfnSchemePDCCH_r17", err)
		}
	}
	if SfnSchemePDSCH_r17Present {
		ie.SfnSchemePDSCH_r17 = new(MIMOParam_r17_sfnSchemePDSCH_r17)
		if err = ie.SfnSchemePDSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SfnSchemePDSCH_r17", err)
		}
	}
	return nil
}
