package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_ConfigDedicated struct {
	SlotSpecificConfigurationsToAddModList  []TDD_UL_DL_SlotConfig `lb:1,ub:maxNrofSlots,optional`
	SlotSpecificConfigurationsToReleaseList []TDD_UL_DL_SlotIndex  `lb:1,ub:maxNrofSlots,optional`
}

func (ie *TDD_UL_DL_ConfigDedicated) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.SlotSpecificConfigurationsToAddModList) > 0, len(ie.SlotSpecificConfigurationsToReleaseList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.SlotSpecificConfigurationsToAddModList) > 0 {
		tmp_SlotSpecificConfigurationsToAddModList := utils.NewSequence[*TDD_UL_DL_SlotConfig]([]*TDD_UL_DL_SlotConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		for _, i := range ie.SlotSpecificConfigurationsToAddModList {
			tmp_SlotSpecificConfigurationsToAddModList.Value = append(tmp_SlotSpecificConfigurationsToAddModList.Value, &i)
		}
		if err = tmp_SlotSpecificConfigurationsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode SlotSpecificConfigurationsToAddModList", err)
		}
	}
	if len(ie.SlotSpecificConfigurationsToReleaseList) > 0 {
		tmp_SlotSpecificConfigurationsToReleaseList := utils.NewSequence[*TDD_UL_DL_SlotIndex]([]*TDD_UL_DL_SlotIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		for _, i := range ie.SlotSpecificConfigurationsToReleaseList {
			tmp_SlotSpecificConfigurationsToReleaseList.Value = append(tmp_SlotSpecificConfigurationsToReleaseList.Value, &i)
		}
		if err = tmp_SlotSpecificConfigurationsToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode SlotSpecificConfigurationsToReleaseList", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_ConfigDedicated) Decode(r *uper.UperReader) error {
	var err error
	var SlotSpecificConfigurationsToAddModListPresent bool
	if SlotSpecificConfigurationsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SlotSpecificConfigurationsToReleaseListPresent bool
	if SlotSpecificConfigurationsToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SlotSpecificConfigurationsToAddModListPresent {
		tmp_SlotSpecificConfigurationsToAddModList := utils.NewSequence[*TDD_UL_DL_SlotConfig]([]*TDD_UL_DL_SlotConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		fn_SlotSpecificConfigurationsToAddModList := func() *TDD_UL_DL_SlotConfig {
			return new(TDD_UL_DL_SlotConfig)
		}
		if err = tmp_SlotSpecificConfigurationsToAddModList.Decode(r, fn_SlotSpecificConfigurationsToAddModList); err != nil {
			return utils.WrapError("Decode SlotSpecificConfigurationsToAddModList", err)
		}
		ie.SlotSpecificConfigurationsToAddModList = []TDD_UL_DL_SlotConfig{}
		for _, i := range tmp_SlotSpecificConfigurationsToAddModList.Value {
			ie.SlotSpecificConfigurationsToAddModList = append(ie.SlotSpecificConfigurationsToAddModList, *i)
		}
	}
	if SlotSpecificConfigurationsToReleaseListPresent {
		tmp_SlotSpecificConfigurationsToReleaseList := utils.NewSequence[*TDD_UL_DL_SlotIndex]([]*TDD_UL_DL_SlotIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		fn_SlotSpecificConfigurationsToReleaseList := func() *TDD_UL_DL_SlotIndex {
			return new(TDD_UL_DL_SlotIndex)
		}
		if err = tmp_SlotSpecificConfigurationsToReleaseList.Decode(r, fn_SlotSpecificConfigurationsToReleaseList); err != nil {
			return utils.WrapError("Decode SlotSpecificConfigurationsToReleaseList", err)
		}
		ie.SlotSpecificConfigurationsToReleaseList = []TDD_UL_DL_SlotIndex{}
		for _, i := range tmp_SlotSpecificConfigurationsToReleaseList.Value {
			ie.SlotSpecificConfigurationsToReleaseList = append(ie.SlotSpecificConfigurationsToReleaseList, *i)
		}
	}
	return nil
}
