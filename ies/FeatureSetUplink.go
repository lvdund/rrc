package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink struct {
	FeatureSetListPerUplinkCC                 []FeatureSetUplinkPerCC_Id                                  `lb:1,ub:maxNrofServingCells,madatory`
	ScalingFactor                             *FeatureSetUplink_scalingFactor                             `optional`
	Dummy3                                    *FeatureSetUplink_dummy3                                    `optional`
	IntraBandFreqSeparationUL                 *FreqSeparationClass                                        `optional`
	SearchSpaceSharingCA_UL                   *FeatureSetUplink_searchSpaceSharingCA_UL                   `optional`
	Dummy1                                    *DummyI                                                     `optional`
	SupportedSRS_Resources                    *SRS_Resources                                              `optional`
	TwoPUCCH_Group                            *FeatureSetUplink_twoPUCCH_Group                            `optional`
	DynamicSwitchSUL                          *FeatureSetUplink_dynamicSwitchSUL                          `optional`
	SimultaneousTxSUL_NonSUL                  *FeatureSetUplink_simultaneousTxSUL_NonSUL                  `optional`
	Pusch_ProcessingType1_DifferentTB_PerSlot *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot `optional`
	Dummy2                                    *DummyF                                                     `optional`
}

func (ie *FeatureSetUplink) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ScalingFactor != nil, ie.Dummy3 != nil, ie.IntraBandFreqSeparationUL != nil, ie.SearchSpaceSharingCA_UL != nil, ie.Dummy1 != nil, ie.SupportedSRS_Resources != nil, ie.TwoPUCCH_Group != nil, ie.DynamicSwitchSUL != nil, ie.SimultaneousTxSUL_NonSUL != nil, ie.Pusch_ProcessingType1_DifferentTB_PerSlot != nil, ie.Dummy2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_FeatureSetListPerUplinkCC := utils.NewSequence[*FeatureSetUplinkPerCC_Id]([]*FeatureSetUplinkPerCC_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.FeatureSetListPerUplinkCC {
		tmp_FeatureSetListPerUplinkCC.Value = append(tmp_FeatureSetListPerUplinkCC.Value, &i)
	}
	if err = tmp_FeatureSetListPerUplinkCC.Encode(w); err != nil {
		return utils.WrapError("Encode FeatureSetListPerUplinkCC", err)
	}
	if ie.ScalingFactor != nil {
		if err = ie.ScalingFactor.Encode(w); err != nil {
			return utils.WrapError("Encode ScalingFactor", err)
		}
	}
	if ie.Dummy3 != nil {
		if err = ie.Dummy3.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy3", err)
		}
	}
	if ie.IntraBandFreqSeparationUL != nil {
		if err = ie.IntraBandFreqSeparationUL.Encode(w); err != nil {
			return utils.WrapError("Encode IntraBandFreqSeparationUL", err)
		}
	}
	if ie.SearchSpaceSharingCA_UL != nil {
		if err = ie.SearchSpaceSharingCA_UL.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceSharingCA_UL", err)
		}
	}
	if ie.Dummy1 != nil {
		if err = ie.Dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy1", err)
		}
	}
	if ie.SupportedSRS_Resources != nil {
		if err = ie.SupportedSRS_Resources.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedSRS_Resources", err)
		}
	}
	if ie.TwoPUCCH_Group != nil {
		if err = ie.TwoPUCCH_Group.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_Group", err)
		}
	}
	if ie.DynamicSwitchSUL != nil {
		if err = ie.DynamicSwitchSUL.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicSwitchSUL", err)
		}
	}
	if ie.SimultaneousTxSUL_NonSUL != nil {
		if err = ie.SimultaneousTxSUL_NonSUL.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousTxSUL_NonSUL", err)
		}
	}
	if ie.Pusch_ProcessingType1_DifferentTB_PerSlot != nil {
		if err = ie.Pusch_ProcessingType1_DifferentTB_PerSlot.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_ProcessingType1_DifferentTB_PerSlot", err)
		}
	}
	if ie.Dummy2 != nil {
		if err = ie.Dummy2.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy2", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink) Decode(r *uper.UperReader) error {
	var err error
	var ScalingFactorPresent bool
	if ScalingFactorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy3Present bool
	if Dummy3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraBandFreqSeparationULPresent bool
	if IntraBandFreqSeparationULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceSharingCA_ULPresent bool
	if SearchSpaceSharingCA_ULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy1Present bool
	if Dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedSRS_ResourcesPresent bool
	if SupportedSRS_ResourcesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_GroupPresent bool
	if TwoPUCCH_GroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicSwitchSULPresent bool
	if DynamicSwitchSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SimultaneousTxSUL_NonSULPresent bool
	if SimultaneousTxSUL_NonSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_ProcessingType1_DifferentTB_PerSlotPresent bool
	if Pusch_ProcessingType1_DifferentTB_PerSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy2Present bool
	if Dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_FeatureSetListPerUplinkCC := utils.NewSequence[*FeatureSetUplinkPerCC_Id]([]*FeatureSetUplinkPerCC_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_FeatureSetListPerUplinkCC := func() *FeatureSetUplinkPerCC_Id {
		return new(FeatureSetUplinkPerCC_Id)
	}
	if err = tmp_FeatureSetListPerUplinkCC.Decode(r, fn_FeatureSetListPerUplinkCC); err != nil {
		return utils.WrapError("Decode FeatureSetListPerUplinkCC", err)
	}
	ie.FeatureSetListPerUplinkCC = []FeatureSetUplinkPerCC_Id{}
	for _, i := range tmp_FeatureSetListPerUplinkCC.Value {
		ie.FeatureSetListPerUplinkCC = append(ie.FeatureSetListPerUplinkCC, *i)
	}
	if ScalingFactorPresent {
		ie.ScalingFactor = new(FeatureSetUplink_scalingFactor)
		if err = ie.ScalingFactor.Decode(r); err != nil {
			return utils.WrapError("Decode ScalingFactor", err)
		}
	}
	if Dummy3Present {
		ie.Dummy3 = new(FeatureSetUplink_dummy3)
		if err = ie.Dummy3.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy3", err)
		}
	}
	if IntraBandFreqSeparationULPresent {
		ie.IntraBandFreqSeparationUL = new(FreqSeparationClass)
		if err = ie.IntraBandFreqSeparationUL.Decode(r); err != nil {
			return utils.WrapError("Decode IntraBandFreqSeparationUL", err)
		}
	}
	if SearchSpaceSharingCA_ULPresent {
		ie.SearchSpaceSharingCA_UL = new(FeatureSetUplink_searchSpaceSharingCA_UL)
		if err = ie.SearchSpaceSharingCA_UL.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceSharingCA_UL", err)
		}
	}
	if Dummy1Present {
		ie.Dummy1 = new(DummyI)
		if err = ie.Dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy1", err)
		}
	}
	if SupportedSRS_ResourcesPresent {
		ie.SupportedSRS_Resources = new(SRS_Resources)
		if err = ie.SupportedSRS_Resources.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedSRS_Resources", err)
		}
	}
	if TwoPUCCH_GroupPresent {
		ie.TwoPUCCH_Group = new(FeatureSetUplink_twoPUCCH_Group)
		if err = ie.TwoPUCCH_Group.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_Group", err)
		}
	}
	if DynamicSwitchSULPresent {
		ie.DynamicSwitchSUL = new(FeatureSetUplink_dynamicSwitchSUL)
		if err = ie.DynamicSwitchSUL.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicSwitchSUL", err)
		}
	}
	if SimultaneousTxSUL_NonSULPresent {
		ie.SimultaneousTxSUL_NonSUL = new(FeatureSetUplink_simultaneousTxSUL_NonSUL)
		if err = ie.SimultaneousTxSUL_NonSUL.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousTxSUL_NonSUL", err)
		}
	}
	if Pusch_ProcessingType1_DifferentTB_PerSlotPresent {
		ie.Pusch_ProcessingType1_DifferentTB_PerSlot = new(FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot)
		if err = ie.Pusch_ProcessingType1_DifferentTB_PerSlot.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_ProcessingType1_DifferentTB_PerSlot", err)
		}
	}
	if Dummy2Present {
		ie.Dummy2 = new(DummyF)
		if err = ie.Dummy2.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy2", err)
		}
	}
	return nil
}
