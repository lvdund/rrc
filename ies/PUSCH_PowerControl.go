package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PowerControl struct {
	Tpc_Accumulation                 *PUSCH_PowerControl_tpc_Accumulation             `optional`
	Msg3_Alpha                       *Alpha                                           `optional`
	P0_NominalWithoutGrant           *int64                                           `lb:-202,ub:24,optional`
	P0_AlphaSets                     []P0_PUSCH_AlphaSet                              `lb:1,ub:maxNrofP0_PUSCH_AlphaSets,optional`
	PathlossReferenceRSToAddModList  []PUSCH_PathlossReferenceRS                      `lb:1,ub:maxNrofPUSCH_PathlossReferenceRSs,optional`
	PathlossReferenceRSToReleaseList []PUSCH_PathlossReferenceRS_Id                   `lb:1,ub:maxNrofPUSCH_PathlossReferenceRSs,optional`
	TwoPUSCH_PC_AdjustmentStates     *PUSCH_PowerControl_twoPUSCH_PC_AdjustmentStates `optional`
	DeltaMCS                         *PUSCH_PowerControl_deltaMCS                     `optional`
	Sri_PUSCH_MappingToAddModList    []SRI_PUSCH_PowerControl                         `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional`
	Sri_PUSCH_MappingToReleaseList   []SRI_PUSCH_PowerControlId                       `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional`
}

func (ie *PUSCH_PowerControl) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Tpc_Accumulation != nil, ie.Msg3_Alpha != nil, ie.P0_NominalWithoutGrant != nil, len(ie.P0_AlphaSets) > 0, len(ie.PathlossReferenceRSToAddModList) > 0, len(ie.PathlossReferenceRSToReleaseList) > 0, ie.TwoPUSCH_PC_AdjustmentStates != nil, ie.DeltaMCS != nil, len(ie.Sri_PUSCH_MappingToAddModList) > 0, len(ie.Sri_PUSCH_MappingToReleaseList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Tpc_Accumulation != nil {
		if err = ie.Tpc_Accumulation.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_Accumulation", err)
		}
	}
	if ie.Msg3_Alpha != nil {
		if err = ie.Msg3_Alpha.Encode(w); err != nil {
			return utils.WrapError("Encode Msg3_Alpha", err)
		}
	}
	if ie.P0_NominalWithoutGrant != nil {
		if err = w.WriteInteger(*ie.P0_NominalWithoutGrant, &aper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Encode P0_NominalWithoutGrant", err)
		}
	}
	if len(ie.P0_AlphaSets) > 0 {
		tmp_P0_AlphaSets := utils.NewSequence[*P0_PUSCH_AlphaSet]([]*P0_PUSCH_AlphaSet{}, aper.Constraint{Lb: 1, Ub: maxNrofP0_PUSCH_AlphaSets}, false)
		for _, i := range ie.P0_AlphaSets {
			tmp_P0_AlphaSets.Value = append(tmp_P0_AlphaSets.Value, &i)
		}
		if err = tmp_P0_AlphaSets.Encode(w); err != nil {
			return utils.WrapError("Encode P0_AlphaSets", err)
		}
	}
	if len(ie.PathlossReferenceRSToAddModList) > 0 {
		tmp_PathlossReferenceRSToAddModList := utils.NewSequence[*PUSCH_PathlossReferenceRS]([]*PUSCH_PathlossReferenceRS{}, aper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs}, false)
		for _, i := range ie.PathlossReferenceRSToAddModList {
			tmp_PathlossReferenceRSToAddModList.Value = append(tmp_PathlossReferenceRSToAddModList.Value, &i)
		}
		if err = tmp_PathlossReferenceRSToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode PathlossReferenceRSToAddModList", err)
		}
	}
	if len(ie.PathlossReferenceRSToReleaseList) > 0 {
		tmp_PathlossReferenceRSToReleaseList := utils.NewSequence[*PUSCH_PathlossReferenceRS_Id]([]*PUSCH_PathlossReferenceRS_Id{}, aper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs}, false)
		for _, i := range ie.PathlossReferenceRSToReleaseList {
			tmp_PathlossReferenceRSToReleaseList.Value = append(tmp_PathlossReferenceRSToReleaseList.Value, &i)
		}
		if err = tmp_PathlossReferenceRSToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode PathlossReferenceRSToReleaseList", err)
		}
	}
	if ie.TwoPUSCH_PC_AdjustmentStates != nil {
		if err = ie.TwoPUSCH_PC_AdjustmentStates.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUSCH_PC_AdjustmentStates", err)
		}
	}
	if ie.DeltaMCS != nil {
		if err = ie.DeltaMCS.Encode(w); err != nil {
			return utils.WrapError("Encode DeltaMCS", err)
		}
	}
	if len(ie.Sri_PUSCH_MappingToAddModList) > 0 {
		tmp_Sri_PUSCH_MappingToAddModList := utils.NewSequence[*SRI_PUSCH_PowerControl]([]*SRI_PUSCH_PowerControl{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		for _, i := range ie.Sri_PUSCH_MappingToAddModList {
			tmp_Sri_PUSCH_MappingToAddModList.Value = append(tmp_Sri_PUSCH_MappingToAddModList.Value, &i)
		}
		if err = tmp_Sri_PUSCH_MappingToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Sri_PUSCH_MappingToAddModList", err)
		}
	}
	if len(ie.Sri_PUSCH_MappingToReleaseList) > 0 {
		tmp_Sri_PUSCH_MappingToReleaseList := utils.NewSequence[*SRI_PUSCH_PowerControlId]([]*SRI_PUSCH_PowerControlId{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		for _, i := range ie.Sri_PUSCH_MappingToReleaseList {
			tmp_Sri_PUSCH_MappingToReleaseList.Value = append(tmp_Sri_PUSCH_MappingToReleaseList.Value, &i)
		}
		if err = tmp_Sri_PUSCH_MappingToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Sri_PUSCH_MappingToReleaseList", err)
		}
	}
	return nil
}

func (ie *PUSCH_PowerControl) Decode(r *aper.AperReader) error {
	var err error
	var Tpc_AccumulationPresent bool
	if Tpc_AccumulationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg3_AlphaPresent bool
	if Msg3_AlphaPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var P0_NominalWithoutGrantPresent bool
	if P0_NominalWithoutGrantPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var P0_AlphaSetsPresent bool
	if P0_AlphaSetsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PathlossReferenceRSToAddModListPresent bool
	if PathlossReferenceRSToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PathlossReferenceRSToReleaseListPresent bool
	if PathlossReferenceRSToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUSCH_PC_AdjustmentStatesPresent bool
	if TwoPUSCH_PC_AdjustmentStatesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DeltaMCSPresent bool
	if DeltaMCSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sri_PUSCH_MappingToAddModListPresent bool
	if Sri_PUSCH_MappingToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sri_PUSCH_MappingToReleaseListPresent bool
	if Sri_PUSCH_MappingToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Tpc_AccumulationPresent {
		ie.Tpc_Accumulation = new(PUSCH_PowerControl_tpc_Accumulation)
		if err = ie.Tpc_Accumulation.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_Accumulation", err)
		}
	}
	if Msg3_AlphaPresent {
		ie.Msg3_Alpha = new(Alpha)
		if err = ie.Msg3_Alpha.Decode(r); err != nil {
			return utils.WrapError("Decode Msg3_Alpha", err)
		}
	}
	if P0_NominalWithoutGrantPresent {
		var tmp_int_P0_NominalWithoutGrant int64
		if tmp_int_P0_NominalWithoutGrant, err = r.ReadInteger(&aper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode P0_NominalWithoutGrant", err)
		}
		ie.P0_NominalWithoutGrant = &tmp_int_P0_NominalWithoutGrant
	}
	if P0_AlphaSetsPresent {
		tmp_P0_AlphaSets := utils.NewSequence[*P0_PUSCH_AlphaSet]([]*P0_PUSCH_AlphaSet{}, aper.Constraint{Lb: 1, Ub: maxNrofP0_PUSCH_AlphaSets}, false)
		fn_P0_AlphaSets := func() *P0_PUSCH_AlphaSet {
			return new(P0_PUSCH_AlphaSet)
		}
		if err = tmp_P0_AlphaSets.Decode(r, fn_P0_AlphaSets); err != nil {
			return utils.WrapError("Decode P0_AlphaSets", err)
		}
		ie.P0_AlphaSets = []P0_PUSCH_AlphaSet{}
		for _, i := range tmp_P0_AlphaSets.Value {
			ie.P0_AlphaSets = append(ie.P0_AlphaSets, *i)
		}
	}
	if PathlossReferenceRSToAddModListPresent {
		tmp_PathlossReferenceRSToAddModList := utils.NewSequence[*PUSCH_PathlossReferenceRS]([]*PUSCH_PathlossReferenceRS{}, aper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs}, false)
		fn_PathlossReferenceRSToAddModList := func() *PUSCH_PathlossReferenceRS {
			return new(PUSCH_PathlossReferenceRS)
		}
		if err = tmp_PathlossReferenceRSToAddModList.Decode(r, fn_PathlossReferenceRSToAddModList); err != nil {
			return utils.WrapError("Decode PathlossReferenceRSToAddModList", err)
		}
		ie.PathlossReferenceRSToAddModList = []PUSCH_PathlossReferenceRS{}
		for _, i := range tmp_PathlossReferenceRSToAddModList.Value {
			ie.PathlossReferenceRSToAddModList = append(ie.PathlossReferenceRSToAddModList, *i)
		}
	}
	if PathlossReferenceRSToReleaseListPresent {
		tmp_PathlossReferenceRSToReleaseList := utils.NewSequence[*PUSCH_PathlossReferenceRS_Id]([]*PUSCH_PathlossReferenceRS_Id{}, aper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs}, false)
		fn_PathlossReferenceRSToReleaseList := func() *PUSCH_PathlossReferenceRS_Id {
			return new(PUSCH_PathlossReferenceRS_Id)
		}
		if err = tmp_PathlossReferenceRSToReleaseList.Decode(r, fn_PathlossReferenceRSToReleaseList); err != nil {
			return utils.WrapError("Decode PathlossReferenceRSToReleaseList", err)
		}
		ie.PathlossReferenceRSToReleaseList = []PUSCH_PathlossReferenceRS_Id{}
		for _, i := range tmp_PathlossReferenceRSToReleaseList.Value {
			ie.PathlossReferenceRSToReleaseList = append(ie.PathlossReferenceRSToReleaseList, *i)
		}
	}
	if TwoPUSCH_PC_AdjustmentStatesPresent {
		ie.TwoPUSCH_PC_AdjustmentStates = new(PUSCH_PowerControl_twoPUSCH_PC_AdjustmentStates)
		if err = ie.TwoPUSCH_PC_AdjustmentStates.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUSCH_PC_AdjustmentStates", err)
		}
	}
	if DeltaMCSPresent {
		ie.DeltaMCS = new(PUSCH_PowerControl_deltaMCS)
		if err = ie.DeltaMCS.Decode(r); err != nil {
			return utils.WrapError("Decode DeltaMCS", err)
		}
	}
	if Sri_PUSCH_MappingToAddModListPresent {
		tmp_Sri_PUSCH_MappingToAddModList := utils.NewSequence[*SRI_PUSCH_PowerControl]([]*SRI_PUSCH_PowerControl{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		fn_Sri_PUSCH_MappingToAddModList := func() *SRI_PUSCH_PowerControl {
			return new(SRI_PUSCH_PowerControl)
		}
		if err = tmp_Sri_PUSCH_MappingToAddModList.Decode(r, fn_Sri_PUSCH_MappingToAddModList); err != nil {
			return utils.WrapError("Decode Sri_PUSCH_MappingToAddModList", err)
		}
		ie.Sri_PUSCH_MappingToAddModList = []SRI_PUSCH_PowerControl{}
		for _, i := range tmp_Sri_PUSCH_MappingToAddModList.Value {
			ie.Sri_PUSCH_MappingToAddModList = append(ie.Sri_PUSCH_MappingToAddModList, *i)
		}
	}
	if Sri_PUSCH_MappingToReleaseListPresent {
		tmp_Sri_PUSCH_MappingToReleaseList := utils.NewSequence[*SRI_PUSCH_PowerControlId]([]*SRI_PUSCH_PowerControlId{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		fn_Sri_PUSCH_MappingToReleaseList := func() *SRI_PUSCH_PowerControlId {
			return new(SRI_PUSCH_PowerControlId)
		}
		if err = tmp_Sri_PUSCH_MappingToReleaseList.Decode(r, fn_Sri_PUSCH_MappingToReleaseList); err != nil {
			return utils.WrapError("Decode Sri_PUSCH_MappingToReleaseList", err)
		}
		ie.Sri_PUSCH_MappingToReleaseList = []SRI_PUSCH_PowerControlId{}
		for _, i := range tmp_Sri_PUSCH_MappingToReleaseList.Value {
			ie.Sri_PUSCH_MappingToReleaseList = append(ie.Sri_PUSCH_MappingToReleaseList, *i)
		}
	}
	return nil
}
