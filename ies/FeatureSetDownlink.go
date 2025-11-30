package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink struct {
	FeatureSetListPerDownlinkCC               []FeatureSetDownlinkPerCC_Id                                  `lb:1,ub:maxNrofServingCells,madatory`
	IntraBandFreqSeparationDL                 *FreqSeparationClass                                          `optional`
	ScalingFactor                             *FeatureSetDownlink_scalingFactor                             `optional`
	Dummy8                                    *FeatureSetDownlink_dummy8                                    `optional`
	ScellWithoutSSB                           *FeatureSetDownlink_scellWithoutSSB                           `optional`
	Csi_RS_MeasSCellWithoutSSB                *FeatureSetDownlink_csi_RS_MeasSCellWithoutSSB                `optional`
	Dummy1                                    *FeatureSetDownlink_dummy1                                    `optional`
	Type1_3_CSS                               *FeatureSetDownlink_type1_3_CSS                               `optional`
	Pdcch_MonitoringAnyOccasions              *FeatureSetDownlink_pdcch_MonitoringAnyOccasions              `optional`
	Dummy2                                    *FeatureSetDownlink_dummy2                                    `optional`
	Ue_SpecificUL_DL_Assignment               *FeatureSetDownlink_ue_SpecificUL_DL_Assignment               `optional`
	SearchSpaceSharingCA_DL                   *FeatureSetDownlink_searchSpaceSharingCA_DL                   `optional`
	TimeDurationForQCL                        *FeatureSetDownlink_timeDurationForQCL                        `optional`
	Pdsch_ProcessingType1_DifferentTB_PerSlot *FeatureSetDownlink_pdsch_ProcessingType1_DifferentTB_PerSlot `optional`
	Dummy3                                    *DummyA                                                       `optional`
	Dummy4                                    []DummyB                                                      `lb:1,ub:maxNrofCodebooks,optional`
	Dummy5                                    []DummyC                                                      `lb:1,ub:maxNrofCodebooks,optional`
	Dummy6                                    []DummyD                                                      `lb:1,ub:maxNrofCodebooks,optional`
	Dummy7                                    []DummyE                                                      `lb:1,ub:maxNrofCodebooks,optional`
}

func (ie *FeatureSetDownlink) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.IntraBandFreqSeparationDL != nil, ie.ScalingFactor != nil, ie.Dummy8 != nil, ie.ScellWithoutSSB != nil, ie.Csi_RS_MeasSCellWithoutSSB != nil, ie.Dummy1 != nil, ie.Type1_3_CSS != nil, ie.Pdcch_MonitoringAnyOccasions != nil, ie.Dummy2 != nil, ie.Ue_SpecificUL_DL_Assignment != nil, ie.SearchSpaceSharingCA_DL != nil, ie.TimeDurationForQCL != nil, ie.Pdsch_ProcessingType1_DifferentTB_PerSlot != nil, ie.Dummy3 != nil, len(ie.Dummy4) > 0, len(ie.Dummy5) > 0, len(ie.Dummy6) > 0, len(ie.Dummy7) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_FeatureSetListPerDownlinkCC := utils.NewSequence[*FeatureSetDownlinkPerCC_Id]([]*FeatureSetDownlinkPerCC_Id{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.FeatureSetListPerDownlinkCC {
		tmp_FeatureSetListPerDownlinkCC.Value = append(tmp_FeatureSetListPerDownlinkCC.Value, &i)
	}
	if err = tmp_FeatureSetListPerDownlinkCC.Encode(w); err != nil {
		return utils.WrapError("Encode FeatureSetListPerDownlinkCC", err)
	}
	if ie.IntraBandFreqSeparationDL != nil {
		if err = ie.IntraBandFreqSeparationDL.Encode(w); err != nil {
			return utils.WrapError("Encode IntraBandFreqSeparationDL", err)
		}
	}
	if ie.ScalingFactor != nil {
		if err = ie.ScalingFactor.Encode(w); err != nil {
			return utils.WrapError("Encode ScalingFactor", err)
		}
	}
	if ie.Dummy8 != nil {
		if err = ie.Dummy8.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy8", err)
		}
	}
	if ie.ScellWithoutSSB != nil {
		if err = ie.ScellWithoutSSB.Encode(w); err != nil {
			return utils.WrapError("Encode ScellWithoutSSB", err)
		}
	}
	if ie.Csi_RS_MeasSCellWithoutSSB != nil {
		if err = ie.Csi_RS_MeasSCellWithoutSSB.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RS_MeasSCellWithoutSSB", err)
		}
	}
	if ie.Dummy1 != nil {
		if err = ie.Dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy1", err)
		}
	}
	if ie.Type1_3_CSS != nil {
		if err = ie.Type1_3_CSS.Encode(w); err != nil {
			return utils.WrapError("Encode Type1_3_CSS", err)
		}
	}
	if ie.Pdcch_MonitoringAnyOccasions != nil {
		if err = ie.Pdcch_MonitoringAnyOccasions.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_MonitoringAnyOccasions", err)
		}
	}
	if ie.Dummy2 != nil {
		if err = ie.Dummy2.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy2", err)
		}
	}
	if ie.Ue_SpecificUL_DL_Assignment != nil {
		if err = ie.Ue_SpecificUL_DL_Assignment.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_SpecificUL_DL_Assignment", err)
		}
	}
	if ie.SearchSpaceSharingCA_DL != nil {
		if err = ie.SearchSpaceSharingCA_DL.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceSharingCA_DL", err)
		}
	}
	if ie.TimeDurationForQCL != nil {
		if err = ie.TimeDurationForQCL.Encode(w); err != nil {
			return utils.WrapError("Encode TimeDurationForQCL", err)
		}
	}
	if ie.Pdsch_ProcessingType1_DifferentTB_PerSlot != nil {
		if err = ie.Pdsch_ProcessingType1_DifferentTB_PerSlot.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ProcessingType1_DifferentTB_PerSlot", err)
		}
	}
	if ie.Dummy3 != nil {
		if err = ie.Dummy3.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy3", err)
		}
	}
	if len(ie.Dummy4) > 0 {
		tmp_Dummy4 := utils.NewSequence[*DummyB]([]*DummyB{}, aper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		for _, i := range ie.Dummy4 {
			tmp_Dummy4.Value = append(tmp_Dummy4.Value, &i)
		}
		if err = tmp_Dummy4.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy4", err)
		}
	}
	if len(ie.Dummy5) > 0 {
		tmp_Dummy5 := utils.NewSequence[*DummyC]([]*DummyC{}, aper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		for _, i := range ie.Dummy5 {
			tmp_Dummy5.Value = append(tmp_Dummy5.Value, &i)
		}
		if err = tmp_Dummy5.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy5", err)
		}
	}
	if len(ie.Dummy6) > 0 {
		tmp_Dummy6 := utils.NewSequence[*DummyD]([]*DummyD{}, aper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		for _, i := range ie.Dummy6 {
			tmp_Dummy6.Value = append(tmp_Dummy6.Value, &i)
		}
		if err = tmp_Dummy6.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy6", err)
		}
	}
	if len(ie.Dummy7) > 0 {
		tmp_Dummy7 := utils.NewSequence[*DummyE]([]*DummyE{}, aper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		for _, i := range ie.Dummy7 {
			tmp_Dummy7.Value = append(tmp_Dummy7.Value, &i)
		}
		if err = tmp_Dummy7.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy7", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink) Decode(r *aper.AperReader) error {
	var err error
	var IntraBandFreqSeparationDLPresent bool
	if IntraBandFreqSeparationDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ScalingFactorPresent bool
	if ScalingFactorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy8Present bool
	if Dummy8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ScellWithoutSSBPresent bool
	if ScellWithoutSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RS_MeasSCellWithoutSSBPresent bool
	if Csi_RS_MeasSCellWithoutSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy1Present bool
	if Dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1_3_CSSPresent bool
	if Type1_3_CSSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_MonitoringAnyOccasionsPresent bool
	if Pdcch_MonitoringAnyOccasionsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy2Present bool
	if Dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_SpecificUL_DL_AssignmentPresent bool
	if Ue_SpecificUL_DL_AssignmentPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceSharingCA_DLPresent bool
	if SearchSpaceSharingCA_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeDurationForQCLPresent bool
	if TimeDurationForQCLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ProcessingType1_DifferentTB_PerSlotPresent bool
	if Pdsch_ProcessingType1_DifferentTB_PerSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy3Present bool
	if Dummy3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy4Present bool
	if Dummy4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy5Present bool
	if Dummy5Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy6Present bool
	if Dummy6Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy7Present bool
	if Dummy7Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_FeatureSetListPerDownlinkCC := utils.NewSequence[*FeatureSetDownlinkPerCC_Id]([]*FeatureSetDownlinkPerCC_Id{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_FeatureSetListPerDownlinkCC := func() *FeatureSetDownlinkPerCC_Id {
		return new(FeatureSetDownlinkPerCC_Id)
	}
	if err = tmp_FeatureSetListPerDownlinkCC.Decode(r, fn_FeatureSetListPerDownlinkCC); err != nil {
		return utils.WrapError("Decode FeatureSetListPerDownlinkCC", err)
	}
	ie.FeatureSetListPerDownlinkCC = []FeatureSetDownlinkPerCC_Id{}
	for _, i := range tmp_FeatureSetListPerDownlinkCC.Value {
		ie.FeatureSetListPerDownlinkCC = append(ie.FeatureSetListPerDownlinkCC, *i)
	}
	if IntraBandFreqSeparationDLPresent {
		ie.IntraBandFreqSeparationDL = new(FreqSeparationClass)
		if err = ie.IntraBandFreqSeparationDL.Decode(r); err != nil {
			return utils.WrapError("Decode IntraBandFreqSeparationDL", err)
		}
	}
	if ScalingFactorPresent {
		ie.ScalingFactor = new(FeatureSetDownlink_scalingFactor)
		if err = ie.ScalingFactor.Decode(r); err != nil {
			return utils.WrapError("Decode ScalingFactor", err)
		}
	}
	if Dummy8Present {
		ie.Dummy8 = new(FeatureSetDownlink_dummy8)
		if err = ie.Dummy8.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy8", err)
		}
	}
	if ScellWithoutSSBPresent {
		ie.ScellWithoutSSB = new(FeatureSetDownlink_scellWithoutSSB)
		if err = ie.ScellWithoutSSB.Decode(r); err != nil {
			return utils.WrapError("Decode ScellWithoutSSB", err)
		}
	}
	if Csi_RS_MeasSCellWithoutSSBPresent {
		ie.Csi_RS_MeasSCellWithoutSSB = new(FeatureSetDownlink_csi_RS_MeasSCellWithoutSSB)
		if err = ie.Csi_RS_MeasSCellWithoutSSB.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_MeasSCellWithoutSSB", err)
		}
	}
	if Dummy1Present {
		ie.Dummy1 = new(FeatureSetDownlink_dummy1)
		if err = ie.Dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy1", err)
		}
	}
	if Type1_3_CSSPresent {
		ie.Type1_3_CSS = new(FeatureSetDownlink_type1_3_CSS)
		if err = ie.Type1_3_CSS.Decode(r); err != nil {
			return utils.WrapError("Decode Type1_3_CSS", err)
		}
	}
	if Pdcch_MonitoringAnyOccasionsPresent {
		ie.Pdcch_MonitoringAnyOccasions = new(FeatureSetDownlink_pdcch_MonitoringAnyOccasions)
		if err = ie.Pdcch_MonitoringAnyOccasions.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_MonitoringAnyOccasions", err)
		}
	}
	if Dummy2Present {
		ie.Dummy2 = new(FeatureSetDownlink_dummy2)
		if err = ie.Dummy2.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy2", err)
		}
	}
	if Ue_SpecificUL_DL_AssignmentPresent {
		ie.Ue_SpecificUL_DL_Assignment = new(FeatureSetDownlink_ue_SpecificUL_DL_Assignment)
		if err = ie.Ue_SpecificUL_DL_Assignment.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_SpecificUL_DL_Assignment", err)
		}
	}
	if SearchSpaceSharingCA_DLPresent {
		ie.SearchSpaceSharingCA_DL = new(FeatureSetDownlink_searchSpaceSharingCA_DL)
		if err = ie.SearchSpaceSharingCA_DL.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceSharingCA_DL", err)
		}
	}
	if TimeDurationForQCLPresent {
		ie.TimeDurationForQCL = new(FeatureSetDownlink_timeDurationForQCL)
		if err = ie.TimeDurationForQCL.Decode(r); err != nil {
			return utils.WrapError("Decode TimeDurationForQCL", err)
		}
	}
	if Pdsch_ProcessingType1_DifferentTB_PerSlotPresent {
		ie.Pdsch_ProcessingType1_DifferentTB_PerSlot = new(FeatureSetDownlink_pdsch_ProcessingType1_DifferentTB_PerSlot)
		if err = ie.Pdsch_ProcessingType1_DifferentTB_PerSlot.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ProcessingType1_DifferentTB_PerSlot", err)
		}
	}
	if Dummy3Present {
		ie.Dummy3 = new(DummyA)
		if err = ie.Dummy3.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy3", err)
		}
	}
	if Dummy4Present {
		tmp_Dummy4 := utils.NewSequence[*DummyB]([]*DummyB{}, aper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		fn_Dummy4 := func() *DummyB {
			return new(DummyB)
		}
		if err = tmp_Dummy4.Decode(r, fn_Dummy4); err != nil {
			return utils.WrapError("Decode Dummy4", err)
		}
		ie.Dummy4 = []DummyB{}
		for _, i := range tmp_Dummy4.Value {
			ie.Dummy4 = append(ie.Dummy4, *i)
		}
	}
	if Dummy5Present {
		tmp_Dummy5 := utils.NewSequence[*DummyC]([]*DummyC{}, aper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		fn_Dummy5 := func() *DummyC {
			return new(DummyC)
		}
		if err = tmp_Dummy5.Decode(r, fn_Dummy5); err != nil {
			return utils.WrapError("Decode Dummy5", err)
		}
		ie.Dummy5 = []DummyC{}
		for _, i := range tmp_Dummy5.Value {
			ie.Dummy5 = append(ie.Dummy5, *i)
		}
	}
	if Dummy6Present {
		tmp_Dummy6 := utils.NewSequence[*DummyD]([]*DummyD{}, aper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		fn_Dummy6 := func() *DummyD {
			return new(DummyD)
		}
		if err = tmp_Dummy6.Decode(r, fn_Dummy6); err != nil {
			return utils.WrapError("Decode Dummy6", err)
		}
		ie.Dummy6 = []DummyD{}
		for _, i := range tmp_Dummy6.Value {
			ie.Dummy6 = append(ie.Dummy6, *i)
		}
	}
	if Dummy7Present {
		tmp_Dummy7 := utils.NewSequence[*DummyE]([]*DummyE{}, aper.Constraint{Lb: 1, Ub: maxNrofCodebooks}, false)
		fn_Dummy7 := func() *DummyE {
			return new(DummyE)
		}
		if err = tmp_Dummy7.Decode(r, fn_Dummy7); err != nil {
			return utils.WrapError("Decode Dummy7", err)
		}
		ie.Dummy7 = []DummyE{}
		for _, i := range tmp_Dummy7.Value {
			ie.Dummy7 = append(ie.Dummy7, *i)
		}
	}
	return nil
}
