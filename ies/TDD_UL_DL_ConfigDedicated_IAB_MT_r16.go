package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_ConfigDedicated_IAB_MT_r16 struct {
	SlotSpecificConfigurationsToAddModList_IAB_MT_r16  []TDD_UL_DL_SlotConfig_IAB_MT_r16 `lb:1,ub:maxNrofSlots,optional`
	SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 []TDD_UL_DL_SlotIndex             `lb:1,ub:maxNrofSlots,optional`
}

func (ie *TDD_UL_DL_ConfigDedicated_IAB_MT_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16) > 0, len(ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16) > 0 {
		tmp_SlotSpecificConfigurationsToAddModList_IAB_MT_r16 := utils.NewSequence[*TDD_UL_DL_SlotConfig_IAB_MT_r16]([]*TDD_UL_DL_SlotConfig_IAB_MT_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		for _, i := range ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16 {
			tmp_SlotSpecificConfigurationsToAddModList_IAB_MT_r16.Value = append(tmp_SlotSpecificConfigurationsToAddModList_IAB_MT_r16.Value, &i)
		}
		if err = tmp_SlotSpecificConfigurationsToAddModList_IAB_MT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SlotSpecificConfigurationsToAddModList_IAB_MT_r16", err)
		}
	}
	if len(ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16) > 0 {
		tmp_SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 := utils.NewSequence[*TDD_UL_DL_SlotIndex]([]*TDD_UL_DL_SlotIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		for _, i := range ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 {
			tmp_SlotSpecificConfigurationsToReleaseList_IAB_MT_r16.Value = append(tmp_SlotSpecificConfigurationsToReleaseList_IAB_MT_r16.Value, &i)
		}
		if err = tmp_SlotSpecificConfigurationsToReleaseList_IAB_MT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SlotSpecificConfigurationsToReleaseList_IAB_MT_r16", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_ConfigDedicated_IAB_MT_r16) Decode(r *uper.UperReader) error {
	var err error
	var SlotSpecificConfigurationsToAddModList_IAB_MT_r16Present bool
	if SlotSpecificConfigurationsToAddModList_IAB_MT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SlotSpecificConfigurationsToReleaseList_IAB_MT_r16Present bool
	if SlotSpecificConfigurationsToReleaseList_IAB_MT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SlotSpecificConfigurationsToAddModList_IAB_MT_r16Present {
		tmp_SlotSpecificConfigurationsToAddModList_IAB_MT_r16 := utils.NewSequence[*TDD_UL_DL_SlotConfig_IAB_MT_r16]([]*TDD_UL_DL_SlotConfig_IAB_MT_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		fn_SlotSpecificConfigurationsToAddModList_IAB_MT_r16 := func() *TDD_UL_DL_SlotConfig_IAB_MT_r16 {
			return new(TDD_UL_DL_SlotConfig_IAB_MT_r16)
		}
		if err = tmp_SlotSpecificConfigurationsToAddModList_IAB_MT_r16.Decode(r, fn_SlotSpecificConfigurationsToAddModList_IAB_MT_r16); err != nil {
			return utils.WrapError("Decode SlotSpecificConfigurationsToAddModList_IAB_MT_r16", err)
		}
		ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16 = []TDD_UL_DL_SlotConfig_IAB_MT_r16{}
		for _, i := range tmp_SlotSpecificConfigurationsToAddModList_IAB_MT_r16.Value {
			ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16 = append(ie.SlotSpecificConfigurationsToAddModList_IAB_MT_r16, *i)
		}
	}
	if SlotSpecificConfigurationsToReleaseList_IAB_MT_r16Present {
		tmp_SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 := utils.NewSequence[*TDD_UL_DL_SlotIndex]([]*TDD_UL_DL_SlotIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		fn_SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 := func() *TDD_UL_DL_SlotIndex {
			return new(TDD_UL_DL_SlotIndex)
		}
		if err = tmp_SlotSpecificConfigurationsToReleaseList_IAB_MT_r16.Decode(r, fn_SlotSpecificConfigurationsToReleaseList_IAB_MT_r16); err != nil {
			return utils.WrapError("Decode SlotSpecificConfigurationsToReleaseList_IAB_MT_r16", err)
		}
		ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 = []TDD_UL_DL_SlotIndex{}
		for _, i := range tmp_SlotSpecificConfigurationsToReleaseList_IAB_MT_r16.Value {
			ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16 = append(ie.SlotSpecificConfigurationsToReleaseList_IAB_MT_r16, *i)
		}
	}
	return nil
}
