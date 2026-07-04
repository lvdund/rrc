// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB1 (line 1980).

var sIB1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellSelectionInfo", Optional: true},
		{Name: "cellAccessRelatedInfo"},
		{Name: "connEstFailureControl", Optional: true},
		{Name: "si-SchedulingInfo", Optional: true},
		{Name: "servingCellConfigCommon", Optional: true},
		{Name: "ims-EmergencySupport", Optional: true},
		{Name: "eCallOverIMS-Support", Optional: true},
		{Name: "ue-TimersAndConstants", Optional: true},
		{Name: "uac-BarringInfo", Optional: true},
		{Name: "useFullResumeID", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	SIB1_Ims_EmergencySupport_True = 0
)

var sIB1ImsEmergencySupportConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_Ims_EmergencySupport_True},
}

const (
	SIB1_ECallOverIMS_Support_True = 0
)

var sIB1ECallOverIMSSupportConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_ECallOverIMS_Support_True},
}

const (
	SIB1_UseFullResumeID_True = 0
)

var sIB1UseFullResumeIDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_UseFullResumeID_True},
}

var sIB1LateNonCriticalExtensionConstraints = per.SizeConstraints{}

var sIB1CellSelectionInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "q-RxLevMin"},
		{Name: "q-RxLevMinOffset", Optional: true},
		{Name: "q-RxLevMinSUL", Optional: true},
		{Name: "q-QualMin", Optional: true},
		{Name: "q-QualMinOffset", Optional: true},
	},
}

var sIB1UacBarringInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uac-BarringForCommon", Optional: true},
		{Name: "uac-BarringPerPLMN-List", Optional: true},
		{Name: "uac-BarringInfoSetList"},
		{Name: "uac-AccessCategory1-SelectionAssistanceInfo", Optional: true},
	},
}

var sIB1UacBarringInfoUacAccessCategory1SelectionAssistanceInfoConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "plmnCommon"},
		{Name: "individualPLMNList"},
	},
}

const (
	SIB1_Uac_BarringInfo_Uac_AccessCategory1_SelectionAssistanceInfo_PlmnCommon         = 0
	SIB1_Uac_BarringInfo_Uac_AccessCategory1_SelectionAssistanceInfo_IndividualPLMNList = 1
)

var sIB1UacBarringInfoUacAccessCategory1SelectionAssistanceInfoIndividualPLMNListConstraints = per.SizeRange(2, common.MaxPLMN)

type SIB1 struct {
	CellSelectionInfo *struct {
		Q_RxLevMin       Q_RxLevMin
		Q_RxLevMinOffset *int64
		Q_RxLevMinSUL    *Q_RxLevMin
		Q_QualMin        *Q_QualMin
		Q_QualMinOffset  *int64
	}
	CellAccessRelatedInfo   CellAccessRelatedInfo
	ConnEstFailureControl   *ConnEstFailureControl
	Si_SchedulingInfo       *SI_SchedulingInfo
	ServingCellConfigCommon *ServingCellConfigCommonSIB
	Ims_EmergencySupport    *int64
	ECallOverIMS_Support    *int64
	Ue_TimersAndConstants   *UE_TimersAndConstants
	Uac_BarringInfo         *struct {
		Uac_BarringForCommon                        *UAC_BarringPerCatList
		Uac_BarringPerPLMN_List                     *UAC_BarringPerPLMN_List
		Uac_BarringInfoSetList                      UAC_BarringInfoSetList
		Uac_AccessCategory1_SelectionAssistanceInfo *struct {
			Choice             int
			PlmnCommon         *UAC_AccessCategory1_SelectionAssistanceInfo
			IndividualPLMNList *[]UAC_AccessCategory1_SelectionAssistanceInfo
		}
	}
	UseFullResumeID          *int64
	LateNonCriticalExtension []byte
	NonCriticalExtension     *SIB1_v1610_IEs
}

func (ie *SIB1) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB1Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CellSelectionInfo != nil, ie.ConnEstFailureControl != nil, ie.Si_SchedulingInfo != nil, ie.ServingCellConfigCommon != nil, ie.Ims_EmergencySupport != nil, ie.ECallOverIMS_Support != nil, ie.Ue_TimersAndConstants != nil, ie.Uac_BarringInfo != nil, ie.UseFullResumeID != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. cellSelectionInfo: sequence
	{
		if ie.CellSelectionInfo != nil {
			c := ie.CellSelectionInfo
			sIB1CellSelectionInfoSeq := e.NewSequenceEncoder(sIB1CellSelectionInfoConstraints)
			if err := sIB1CellSelectionInfoSeq.EncodePreamble([]bool{c.Q_RxLevMinOffset != nil, c.Q_RxLevMinSUL != nil, c.Q_QualMin != nil, c.Q_QualMinOffset != nil}); err != nil {
				return err
			}
			if err := c.Q_RxLevMin.Encode(e); err != nil {
				return err
			}
			if c.Q_RxLevMinOffset != nil {
				if err := e.EncodeInteger((*c.Q_RxLevMinOffset), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
			if c.Q_RxLevMinSUL != nil {
				if err := c.Q_RxLevMinSUL.Encode(e); err != nil {
					return err
				}
			}
			if c.Q_QualMin != nil {
				if err := c.Q_QualMin.Encode(e); err != nil {
					return err
				}
			}
			if c.Q_QualMinOffset != nil {
				if err := e.EncodeInteger((*c.Q_QualMinOffset), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
		}
	}

	// 3. cellAccessRelatedInfo: ref
	{
		if err := ie.CellAccessRelatedInfo.Encode(e); err != nil {
			return err
		}
	}

	// 4. connEstFailureControl: ref
	{
		if ie.ConnEstFailureControl != nil {
			if err := ie.ConnEstFailureControl.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. si-SchedulingInfo: ref
	{
		if ie.Si_SchedulingInfo != nil {
			if err := ie.Si_SchedulingInfo.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. servingCellConfigCommon: ref
	{
		if ie.ServingCellConfigCommon != nil {
			if err := ie.ServingCellConfigCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. ims-EmergencySupport: enumerated
	{
		if ie.Ims_EmergencySupport != nil {
			if err := e.EncodeEnumerated(*ie.Ims_EmergencySupport, sIB1ImsEmergencySupportConstraints); err != nil {
				return err
			}
		}
	}

	// 8. eCallOverIMS-Support: enumerated
	{
		if ie.ECallOverIMS_Support != nil {
			if err := e.EncodeEnumerated(*ie.ECallOverIMS_Support, sIB1ECallOverIMSSupportConstraints); err != nil {
				return err
			}
		}
	}

	// 9. ue-TimersAndConstants: ref
	{
		if ie.Ue_TimersAndConstants != nil {
			if err := ie.Ue_TimersAndConstants.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. uac-BarringInfo: sequence
	{
		if ie.Uac_BarringInfo != nil {
			c := ie.Uac_BarringInfo
			sIB1UacBarringInfoSeq := e.NewSequenceEncoder(sIB1UacBarringInfoConstraints)
			if err := sIB1UacBarringInfoSeq.EncodePreamble([]bool{c.Uac_BarringForCommon != nil, c.Uac_BarringPerPLMN_List != nil, c.Uac_AccessCategory1_SelectionAssistanceInfo != nil}); err != nil {
				return err
			}
			if c.Uac_BarringForCommon != nil {
				if err := c.Uac_BarringForCommon.Encode(e); err != nil {
					return err
				}
			}
			if c.Uac_BarringPerPLMN_List != nil {
				if err := c.Uac_BarringPerPLMN_List.Encode(e); err != nil {
					return err
				}
			}
			if err := c.Uac_BarringInfoSetList.Encode(e); err != nil {
				return err
			}
			if c.Uac_AccessCategory1_SelectionAssistanceInfo != nil {
				choiceEnc := e.NewChoiceEncoder(sIB1UacBarringInfoUacAccessCategory1SelectionAssistanceInfoConstraints)
				if err := choiceEnc.EncodeChoice(int64((*c.Uac_AccessCategory1_SelectionAssistanceInfo).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.Uac_AccessCategory1_SelectionAssistanceInfo).Choice {
				case SIB1_Uac_BarringInfo_Uac_AccessCategory1_SelectionAssistanceInfo_PlmnCommon:
					if err := (*c.Uac_AccessCategory1_SelectionAssistanceInfo).PlmnCommon.Encode(e); err != nil {
						return err
					}
				case SIB1_Uac_BarringInfo_Uac_AccessCategory1_SelectionAssistanceInfo_IndividualPLMNList:
					seqOf := e.NewSequenceOfEncoder(sIB1UacBarringInfoUacAccessCategory1SelectionAssistanceInfoIndividualPLMNListConstraints)
					if err := seqOf.EncodeLength(int64(len((*(*c.Uac_AccessCategory1_SelectionAssistanceInfo).IndividualPLMNList)))); err != nil {
						return err
					}
					for i := range *(*c.Uac_AccessCategory1_SelectionAssistanceInfo).IndividualPLMNList {
						if err := (*(*c.Uac_AccessCategory1_SelectionAssistanceInfo).IndividualPLMNList)[i].Encode(e); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	// 11. useFullResumeID: enumerated
	{
		if ie.UseFullResumeID != nil {
			if err := e.EncodeEnumerated(*ie.UseFullResumeID, sIB1UseFullResumeIDConstraints); err != nil {
				return err
			}
		}
	}

	// 12. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB1LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 13. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB1) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB1Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cellSelectionInfo: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.CellSelectionInfo = &struct {
				Q_RxLevMin       Q_RxLevMin
				Q_RxLevMinOffset *int64
				Q_RxLevMinSUL    *Q_RxLevMin
				Q_QualMin        *Q_QualMin
				Q_QualMinOffset  *int64
			}{}
			c := ie.CellSelectionInfo
			sIB1CellSelectionInfoSeq := d.NewSequenceDecoder(sIB1CellSelectionInfoConstraints)
			if err := sIB1CellSelectionInfoSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.Q_RxLevMin.Decode(d); err != nil {
					return err
				}
			}
			if sIB1CellSelectionInfoSeq.IsComponentPresent(1) {
				c.Q_RxLevMinOffset = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.Q_RxLevMinOffset) = v
			}
			if sIB1CellSelectionInfoSeq.IsComponentPresent(2) {
				c.Q_RxLevMinSUL = new(Q_RxLevMin)
				if err := (*c.Q_RxLevMinSUL).Decode(d); err != nil {
					return err
				}
			}
			if sIB1CellSelectionInfoSeq.IsComponentPresent(3) {
				c.Q_QualMin = new(Q_QualMin)
				if err := (*c.Q_QualMin).Decode(d); err != nil {
					return err
				}
			}
			if sIB1CellSelectionInfoSeq.IsComponentPresent(4) {
				c.Q_QualMinOffset = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.Q_QualMinOffset) = v
			}
		}
	}

	// 3. cellAccessRelatedInfo: ref
	{
		if err := ie.CellAccessRelatedInfo.Decode(d); err != nil {
			return err
		}
	}

	// 4. connEstFailureControl: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ConnEstFailureControl = new(ConnEstFailureControl)
			if err := ie.ConnEstFailureControl.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. si-SchedulingInfo: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Si_SchedulingInfo = new(SI_SchedulingInfo)
			if err := ie.Si_SchedulingInfo.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. servingCellConfigCommon: ref
	{
		if seq.IsComponentPresent(4) {
			ie.ServingCellConfigCommon = new(ServingCellConfigCommonSIB)
			if err := ie.ServingCellConfigCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. ims-EmergencySupport: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(sIB1ImsEmergencySupportConstraints)
			if err != nil {
				return err
			}
			ie.Ims_EmergencySupport = &idx
		}
	}

	// 8. eCallOverIMS-Support: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(sIB1ECallOverIMSSupportConstraints)
			if err != nil {
				return err
			}
			ie.ECallOverIMS_Support = &idx
		}
	}

	// 9. ue-TimersAndConstants: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Ue_TimersAndConstants = new(UE_TimersAndConstants)
			if err := ie.Ue_TimersAndConstants.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. uac-BarringInfo: sequence
	{
		if seq.IsComponentPresent(8) {
			ie.Uac_BarringInfo = &struct {
				Uac_BarringForCommon                        *UAC_BarringPerCatList
				Uac_BarringPerPLMN_List                     *UAC_BarringPerPLMN_List
				Uac_BarringInfoSetList                      UAC_BarringInfoSetList
				Uac_AccessCategory1_SelectionAssistanceInfo *struct {
					Choice             int
					PlmnCommon         *UAC_AccessCategory1_SelectionAssistanceInfo
					IndividualPLMNList *[]UAC_AccessCategory1_SelectionAssistanceInfo
				}
			}{}
			c := ie.Uac_BarringInfo
			sIB1UacBarringInfoSeq := d.NewSequenceDecoder(sIB1UacBarringInfoConstraints)
			if err := sIB1UacBarringInfoSeq.DecodePreamble(); err != nil {
				return err
			}
			if sIB1UacBarringInfoSeq.IsComponentPresent(0) {
				c.Uac_BarringForCommon = new(UAC_BarringPerCatList)
				if err := (*c.Uac_BarringForCommon).Decode(d); err != nil {
					return err
				}
			}
			if sIB1UacBarringInfoSeq.IsComponentPresent(1) {
				c.Uac_BarringPerPLMN_List = new(UAC_BarringPerPLMN_List)
				if err := (*c.Uac_BarringPerPLMN_List).Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.Uac_BarringInfoSetList.Decode(d); err != nil {
					return err
				}
			}
			if sIB1UacBarringInfoSeq.IsComponentPresent(3) {
				c.Uac_AccessCategory1_SelectionAssistanceInfo = &struct {
					Choice             int
					PlmnCommon         *UAC_AccessCategory1_SelectionAssistanceInfo
					IndividualPLMNList *[]UAC_AccessCategory1_SelectionAssistanceInfo
				}{}
				choiceDec := d.NewChoiceDecoder(sIB1UacBarringInfoUacAccessCategory1SelectionAssistanceInfoConstraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.Uac_AccessCategory1_SelectionAssistanceInfo).Choice = int(index)
				switch index {
				case SIB1_Uac_BarringInfo_Uac_AccessCategory1_SelectionAssistanceInfo_PlmnCommon:
					(*c.Uac_AccessCategory1_SelectionAssistanceInfo).PlmnCommon = new(UAC_AccessCategory1_SelectionAssistanceInfo)
					if err := (*c.Uac_AccessCategory1_SelectionAssistanceInfo).PlmnCommon.Decode(d); err != nil {
						return err
					}
				case SIB1_Uac_BarringInfo_Uac_AccessCategory1_SelectionAssistanceInfo_IndividualPLMNList:
					(*c.Uac_AccessCategory1_SelectionAssistanceInfo).IndividualPLMNList = new([]UAC_AccessCategory1_SelectionAssistanceInfo)
					seqOf := d.NewSequenceOfDecoder(sIB1UacBarringInfoUacAccessCategory1SelectionAssistanceInfoIndividualPLMNListConstraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					(*(*c.Uac_AccessCategory1_SelectionAssistanceInfo).IndividualPLMNList) = make([]UAC_AccessCategory1_SelectionAssistanceInfo, n)
					for i := int64(0); i < n; i++ {
						if err := (*(*c.Uac_AccessCategory1_SelectionAssistanceInfo).IndividualPLMNList)[i].Decode(d); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	// 11. useFullResumeID: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(sIB1UseFullResumeIDConstraints)
			if err != nil {
				return err
			}
			ie.UseFullResumeID = &idx
		}
	}

	// 12. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(10) {
			v, err := d.DecodeOctetString(sIB1LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 13. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(11) {
			ie.NonCriticalExtension = new(SIB1_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
