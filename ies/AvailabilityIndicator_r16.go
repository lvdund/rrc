package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type AvailabilityIndicator_r16 struct {
	Ai_RNTI_r16                    AI_RNTI_r16                                `madatory`
	Dci_PayloadSizeAI_r16          int64                                      `lb:1,ub:maxAI_DCI_PayloadSize_r16,madatory`
	AvailableCombToAddModList_r16  []AvailabilityCombinationsPerCell_r16      `lb:1,ub:maxNrofDUCells_r16,optional`
	AvailableCombToReleaseList_r16 []AvailabilityCombinationsPerCellIndex_r16 `lb:1,ub:maxNrofDUCells_r16,optional`
}

func (ie *AvailabilityIndicator_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.AvailableCombToAddModList_r16) > 0, len(ie.AvailableCombToReleaseList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Ai_RNTI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ai_RNTI_r16", err)
	}
	if err = w.WriteInteger(ie.Dci_PayloadSizeAI_r16, &aper.Constraint{Lb: 1, Ub: maxAI_DCI_PayloadSize_r16}, false); err != nil {
		return utils.WrapError("WriteInteger Dci_PayloadSizeAI_r16", err)
	}
	if len(ie.AvailableCombToAddModList_r16) > 0 {
		tmp_AvailableCombToAddModList_r16 := utils.NewSequence[*AvailabilityCombinationsPerCell_r16]([]*AvailabilityCombinationsPerCell_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofDUCells_r16}, false)
		for _, i := range ie.AvailableCombToAddModList_r16 {
			tmp_AvailableCombToAddModList_r16.Value = append(tmp_AvailableCombToAddModList_r16.Value, &i)
		}
		if err = tmp_AvailableCombToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AvailableCombToAddModList_r16", err)
		}
	}
	if len(ie.AvailableCombToReleaseList_r16) > 0 {
		tmp_AvailableCombToReleaseList_r16 := utils.NewSequence[*AvailabilityCombinationsPerCellIndex_r16]([]*AvailabilityCombinationsPerCellIndex_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofDUCells_r16}, false)
		for _, i := range ie.AvailableCombToReleaseList_r16 {
			tmp_AvailableCombToReleaseList_r16.Value = append(tmp_AvailableCombToReleaseList_r16.Value, &i)
		}
		if err = tmp_AvailableCombToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AvailableCombToReleaseList_r16", err)
		}
	}
	return nil
}

func (ie *AvailabilityIndicator_r16) Decode(r *aper.AperReader) error {
	var err error
	var AvailableCombToAddModList_r16Present bool
	if AvailableCombToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AvailableCombToReleaseList_r16Present bool
	if AvailableCombToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Ai_RNTI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ai_RNTI_r16", err)
	}
	var tmp_int_Dci_PayloadSizeAI_r16 int64
	if tmp_int_Dci_PayloadSizeAI_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxAI_DCI_PayloadSize_r16}, false); err != nil {
		return utils.WrapError("ReadInteger Dci_PayloadSizeAI_r16", err)
	}
	ie.Dci_PayloadSizeAI_r16 = tmp_int_Dci_PayloadSizeAI_r16
	if AvailableCombToAddModList_r16Present {
		tmp_AvailableCombToAddModList_r16 := utils.NewSequence[*AvailabilityCombinationsPerCell_r16]([]*AvailabilityCombinationsPerCell_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofDUCells_r16}, false)
		fn_AvailableCombToAddModList_r16 := func() *AvailabilityCombinationsPerCell_r16 {
			return new(AvailabilityCombinationsPerCell_r16)
		}
		if err = tmp_AvailableCombToAddModList_r16.Decode(r, fn_AvailableCombToAddModList_r16); err != nil {
			return utils.WrapError("Decode AvailableCombToAddModList_r16", err)
		}
		ie.AvailableCombToAddModList_r16 = []AvailabilityCombinationsPerCell_r16{}
		for _, i := range tmp_AvailableCombToAddModList_r16.Value {
			ie.AvailableCombToAddModList_r16 = append(ie.AvailableCombToAddModList_r16, *i)
		}
	}
	if AvailableCombToReleaseList_r16Present {
		tmp_AvailableCombToReleaseList_r16 := utils.NewSequence[*AvailabilityCombinationsPerCellIndex_r16]([]*AvailabilityCombinationsPerCellIndex_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofDUCells_r16}, false)
		fn_AvailableCombToReleaseList_r16 := func() *AvailabilityCombinationsPerCellIndex_r16 {
			return new(AvailabilityCombinationsPerCellIndex_r16)
		}
		if err = tmp_AvailableCombToReleaseList_r16.Decode(r, fn_AvailableCombToReleaseList_r16); err != nil {
			return utils.WrapError("Decode AvailableCombToReleaseList_r16", err)
		}
		ie.AvailableCombToReleaseList_r16 = []AvailabilityCombinationsPerCellIndex_r16{}
		for _, i := range tmp_AvailableCombToReleaseList_r16.Value {
			ie.AvailableCombToReleaseList_r16 = append(ie.AvailableCombToReleaseList_r16, *i)
		}
	}
	return nil
}
