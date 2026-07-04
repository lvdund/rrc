// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-AssociatedReportConfigInfo (line 6860).

var cSIAssociatedReportConfigInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportConfigId"},
		{Name: "resourcesForChannel"},
		{Name: "csi-IM-ResourcesForInterference", Optional: true},
		{Name: "nzp-CSI-RS-ResourcesForInterference", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var cSI_AssociatedReportConfigInfoResourcesForChannelConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nzp-CSI-RS"},
		{Name: "csi-SSB-ResourceSet"},
	},
}

const (
	CSI_AssociatedReportConfigInfo_ResourcesForChannel_Nzp_CSI_RS          = 0
	CSI_AssociatedReportConfigInfo_ResourcesForChannel_Csi_SSB_ResourceSet = 1
)

var cSIAssociatedReportConfigInfoCsiIMResourcesForInterferenceConstraints = per.Constrained(1, common.MaxNrofCSI_IM_ResourceSetsPerConfig)

var cSIAssociatedReportConfigInfoNzpCSIRSResourcesForInterferenceConstraints = per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig)

var cSIAssociatedReportConfigInfoExtResourcesForChannel2R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nzp-CSI-RS2-r17"},
		{Name: "csi-SSB-ResourceSet2-r17"},
	},
}

const (
	CSI_AssociatedReportConfigInfo_Ext_ResourcesForChannel2_r17_Nzp_CSI_RS2_r17          = 0
	CSI_AssociatedReportConfigInfo_Ext_ResourcesForChannel2_r17_Csi_SSB_ResourceSet2_r17 = 1
)

var cSIAssociatedReportConfigInfoCsiSSBResourceSetExtConstraints = per.Constrained(1, common.MaxNrofCSI_SSB_ResourceSetsPerConfigExt)

var cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "perSet-r18"},
		{Name: "perResource-r18"},
	},
}

const (
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerSet_r18      = 0
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerResource_r18 = 1
)

var cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "perSet-r18"},
		{Name: "perResource-r18"},
	},
}

const (
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerSet_r18      = 0
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerResource_r18 = 1
)

var cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceSet"},
		{Name: "qcl-info", Optional: true},
	},
}

var cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSQclInfoConstraints = per.SizeRange(1, common.MaxNrofAP_CSI_RS_ResourcesPerSet)

var cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceSet2-r17"},
		{Name: "qcl-info2-r17", Optional: true},
	},
}

var cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17QclInfo2R17Constraints = per.SizeRange(1, common.MaxNrofAP_CSI_RS_ResourcesPerSet)

var cSIAssociatedReportConfigInfoExtResourcesForChannelTDCPR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceSet2TDCP-r18"},
		{Name: "resourceSet3TDCP-r18", Optional: true},
	},
}

const (
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerSet_r18_First  = 0
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerSet_r18_Second = 1
)

var cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18PerSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerSet_r18_First, CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerSet_r18_Second},
}

var cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18PerResourceR18Constraints = per.SizeRange(1, common.MaxNrofAP_CSI_RS_ResourcesPerSet)

const (
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerResource_r18_Elem_First  = 0
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerResource_r18_Elem_Second = 1
)

var cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18PerResourceR18ElemConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerResource_r18_Elem_First, CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerResource_r18_Elem_Second},
}

const (
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerSet_r18_First  = 0
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerSet_r18_Second = 1
)

var cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18PerSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerSet_r18_First, CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerSet_r18_Second},
}

var cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18PerResourceR18Constraints = per.SizeRange(1, common.MaxNrofAP_CSI_RS_ResourcesPerSet)

const (
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerResource_r18_Elem_First  = 0
	CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerResource_r18_Elem_Second = 1
)

var cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18PerResourceR18ElemConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerResource_r18_Elem_First, CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerResource_r18_Elem_Second},
}

var cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceSetCLI-r19", Optional: true},
		{Name: "qcl-infoCLI-r19", Optional: true},
	},
}

var cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19QclInfoCLIR19Constraints = per.SizeRange(1, common.MaxNrofCLI_MeasResourceSetsPerConfig_r19)

var cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nzp-CSI-RS2-r19"},
		{Name: "nzp-CSI-RS3-r19", Optional: true},
		{Name: "nzp-CSI-RS4-r19", Optional: true},
	},
}

var cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceSet2CJTC-r19"},
		{Name: "qcl-info2CJTC-r19", Optional: true},
	},
}

var cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19QclInfo2CJTCR19Constraints = per.SizeRange(1, common.MaxNrofAP_CSI_RS_ResourcesPerSet)

var cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceSet3CJTC-r19"},
		{Name: "qcl-info3CJTC-r19", Optional: true},
	},
}

var cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19QclInfo3CJTCR19Constraints = per.SizeRange(1, common.MaxNrofAP_CSI_RS_ResourcesPerSet)

var cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceSet4CJTC-r19"},
		{Name: "qcl-info4CJTC-r19", Optional: true},
	},
}

var cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19QclInfo4CJTCR19Constraints = per.SizeRange(1, common.MaxNrofAP_CSI_RS_ResourcesPerSet)

var cSIAssociatedReportConfigInfoExtMrSelectedResourcesR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "firstSelectedResource-r19"},
		{Name: "secondSelectedResource-r19", Optional: true},
	},
}

type CSI_AssociatedReportConfigInfo struct {
	ReportConfigId      CSI_ReportConfigId
	ResourcesForChannel struct {
		Choice     int
		Nzp_CSI_RS *struct {
			ResourceSet int64
			Qcl_Info    []TCI_StateId
		}
		Csi_SSB_ResourceSet *int64
	}
	Csi_IM_ResourcesForInterference     *int64
	Nzp_CSI_RS_ResourcesForInterference *int64
	ResourcesForChannel2_r17            *struct {
		Choice          int
		Nzp_CSI_RS2_r17 *struct {
			ResourceSet2_r17 int64
			Qcl_Info2_r17    []TCI_StateId
		}
		Csi_SSB_ResourceSet2_r17 *int64
	}
	Csi_SSB_ResourceSetExt      *int64
	ResourcesForChannelTDCP_r18 *struct {
		ResourceSet2TDCP_r18 int64
		ResourceSet3TDCP_r18 *int64
	}
	ApplyIndicatedTCI_State_r18 *struct {
		Choice          int
		PerSet_r18      *int64
		PerResource_r18 *[]int64
	}
	ApplyIndicatedTCI_State2_r18 *struct {
		Choice          int
		PerSet_r18      *int64
		PerResource_r18 *[]int64
	}
	Csi_ReportSubConfigTriggerList_r18 *CSI_ReportSubConfigTriggerList_r18
	ResourcesForChannelCLI_r19         *struct {
		ResourceSetCLI_r19 *int64
		Qcl_InfoCLI_r19    []TCI_StateId
	}
	ResourcesForChannelCJTC_r19 *struct {
		Nzp_CSI_RS2_r19 struct {
			ResourceSet2CJTC_r19 int64
			Qcl_Info2CJTC_r19    []TCI_StateId
		}
		Nzp_CSI_RS3_r19 *struct {
			ResourceSet3CJTC_r19 int64
			Qcl_Info3CJTC_r19    []TCI_StateId
		}
		Nzp_CSI_RS4_r19 *struct {
			ResourceSet4CJTC_r19 int64
			Qcl_Info4CJTC_r19    []TCI_StateId
		}
	}
	Mr_SelectedResources_r19 *struct {
		FirstSelectedResource_r19  int64
		SecondSelectedResource_r19 *int64
	}
}

func (ie *CSI_AssociatedReportConfigInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIAssociatedReportConfigInfoConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ResourcesForChannel2_r17 != nil || ie.Csi_SSB_ResourceSetExt != nil
	hasExtGroup1 := ie.ResourcesForChannelTDCP_r18 != nil || ie.ApplyIndicatedTCI_State_r18 != nil || ie.ApplyIndicatedTCI_State2_r18 != nil || ie.Csi_ReportSubConfigTriggerList_r18 != nil
	hasExtGroup2 := ie.ResourcesForChannelCLI_r19 != nil || ie.ResourcesForChannelCJTC_r19 != nil || ie.Mr_SelectedResources_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_IM_ResourcesForInterference != nil, ie.Nzp_CSI_RS_ResourcesForInterference != nil}); err != nil {
		return err
	}

	// 3. reportConfigId: ref
	{
		if err := ie.ReportConfigId.Encode(e); err != nil {
			return err
		}
	}

	// 4. resourcesForChannel: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_AssociatedReportConfigInfoResourcesForChannelConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ResourcesForChannel.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ResourcesForChannel.Choice {
		case CSI_AssociatedReportConfigInfo_ResourcesForChannel_Nzp_CSI_RS:
			c := (*ie.ResourcesForChannel.Nzp_CSI_RS)
			cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSSeq := e.NewSequenceEncoder(cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSConstraints)
			if err := cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSSeq.EncodePreamble([]bool{c.Qcl_Info != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.ResourceSet, per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig)); err != nil {
				return err
			}
			if c.Qcl_Info != nil {
				seqOf := e.NewSequenceOfEncoder(cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSQclInfoConstraints)
				if err := seqOf.EncodeLength(int64(len(c.Qcl_Info))); err != nil {
					return err
				}
				for i := range c.Qcl_Info {
					if err := c.Qcl_Info[i].Encode(e); err != nil {
						return err
					}
				}
			}
		case CSI_AssociatedReportConfigInfo_ResourcesForChannel_Csi_SSB_ResourceSet:
			if err := e.EncodeInteger((*ie.ResourcesForChannel.Csi_SSB_ResourceSet), per.Constrained(1, common.MaxNrofCSI_SSB_ResourceSetsPerConfig)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ResourcesForChannel.Choice), Constraint: "index out of range"}
		}
	}

	// 5. csi-IM-ResourcesForInterference: integer
	{
		if ie.Csi_IM_ResourcesForInterference != nil {
			if err := e.EncodeInteger(*ie.Csi_IM_ResourcesForInterference, cSIAssociatedReportConfigInfoCsiIMResourcesForInterferenceConstraints); err != nil {
				return err
			}
		}
	}

	// 6. nzp-CSI-RS-ResourcesForInterference: integer
	{
		if ie.Nzp_CSI_RS_ResourcesForInterference != nil {
			if err := e.EncodeInteger(*ie.Nzp_CSI_RS_ResourcesForInterference, cSIAssociatedReportConfigInfoNzpCSIRSResourcesForInterferenceConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "resourcesForChannel2-r17", Optional: true},
					{Name: "csi-SSB-ResourceSetExt", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ResourcesForChannel2_r17 != nil, ie.Csi_SSB_ResourceSetExt != nil}); err != nil {
				return err
			}

			if ie.ResourcesForChannel2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannel2R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ResourcesForChannel2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ResourcesForChannel2_r17).Choice {
				case CSI_AssociatedReportConfigInfo_Ext_ResourcesForChannel2_r17_Nzp_CSI_RS2_r17:
					c := (*(*ie.ResourcesForChannel2_r17).Nzp_CSI_RS2_r17)
					cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17Seq := ex.NewSequenceEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17Constraints)
					if err := cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17Seq.EncodePreamble([]bool{c.Qcl_Info2_r17 != nil}); err != nil {
						return err
					}
					if err := ex.EncodeInteger(c.ResourceSet2_r17, per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig)); err != nil {
						return err
					}
					if c.Qcl_Info2_r17 != nil {
						seqOf := ex.NewSequenceOfEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17QclInfo2R17Constraints)
						if err := seqOf.EncodeLength(int64(len(c.Qcl_Info2_r17))); err != nil {
							return err
						}
						for i := range c.Qcl_Info2_r17 {
							if err := c.Qcl_Info2_r17[i].Encode(ex); err != nil {
								return err
							}
						}
					}
				case CSI_AssociatedReportConfigInfo_Ext_ResourcesForChannel2_r17_Csi_SSB_ResourceSet2_r17:
					if err := ex.EncodeInteger((*(*ie.ResourcesForChannel2_r17).Csi_SSB_ResourceSet2_r17), per.Constrained(1, common.MaxNrofCSI_SSB_ResourceSetsPerConfigExt)); err != nil {
						return err
					}
				}
			}

			if ie.Csi_SSB_ResourceSetExt != nil {
				if err := ex.EncodeInteger(*ie.Csi_SSB_ResourceSetExt, cSIAssociatedReportConfigInfoCsiSSBResourceSetExtConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "resourcesForChannelTDCP-r18", Optional: true},
					{Name: "applyIndicatedTCI-State-r18", Optional: true},
					{Name: "applyIndicatedTCI-State2-r18", Optional: true},
					{Name: "csi-ReportSubConfigTriggerList-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ResourcesForChannelTDCP_r18 != nil, ie.ApplyIndicatedTCI_State_r18 != nil, ie.ApplyIndicatedTCI_State2_r18 != nil, ie.Csi_ReportSubConfigTriggerList_r18 != nil}); err != nil {
				return err
			}

			if ie.ResourcesForChannelTDCP_r18 != nil {
				c := ie.ResourcesForChannelTDCP_r18
				cSIAssociatedReportConfigInfoExtResourcesForChannelTDCPR18Seq := ex.NewSequenceEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannelTDCPR18Constraints)
				if err := cSIAssociatedReportConfigInfoExtResourcesForChannelTDCPR18Seq.EncodePreamble([]bool{c.ResourceSet3TDCP_r18 != nil}); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ResourceSet2TDCP_r18, per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig)); err != nil {
					return err
				}
				if c.ResourceSet3TDCP_r18 != nil {
					if err := ex.EncodeInteger((*c.ResourceSet3TDCP_r18), per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig)); err != nil {
						return err
					}
				}
			}

			if ie.ApplyIndicatedTCI_State_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ApplyIndicatedTCI_State_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ApplyIndicatedTCI_State_r18).Choice {
				case CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerSet_r18:
					if err := ex.EncodeEnumerated((*(*ie.ApplyIndicatedTCI_State_r18).PerSet_r18), cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18PerSetR18Constraints); err != nil {
						return err
					}
				case CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerResource_r18:
					seqOf := ex.NewSequenceOfEncoder(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18PerResourceR18Constraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.ApplyIndicatedTCI_State_r18).PerResource_r18)))); err != nil {
						return err
					}
					for i := range *(*ie.ApplyIndicatedTCI_State_r18).PerResource_r18 {
						if err := ex.EncodeEnumerated((*(*ie.ApplyIndicatedTCI_State_r18).PerResource_r18)[i], cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18PerResourceR18ElemConstraints); err != nil {
							return err
						}
					}
				}
			}

			if ie.ApplyIndicatedTCI_State2_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ApplyIndicatedTCI_State2_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ApplyIndicatedTCI_State2_r18).Choice {
				case CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerSet_r18:
					if err := ex.EncodeEnumerated((*(*ie.ApplyIndicatedTCI_State2_r18).PerSet_r18), cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18PerSetR18Constraints); err != nil {
						return err
					}
				case CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerResource_r18:
					seqOf := ex.NewSequenceOfEncoder(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18PerResourceR18Constraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.ApplyIndicatedTCI_State2_r18).PerResource_r18)))); err != nil {
						return err
					}
					for i := range *(*ie.ApplyIndicatedTCI_State2_r18).PerResource_r18 {
						if err := ex.EncodeEnumerated((*(*ie.ApplyIndicatedTCI_State2_r18).PerResource_r18)[i], cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18PerResourceR18ElemConstraints); err != nil {
							return err
						}
					}
				}
			}

			if ie.Csi_ReportSubConfigTriggerList_r18 != nil {
				if err := ie.Csi_ReportSubConfigTriggerList_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "resourcesForChannelCLI-r19", Optional: true},
					{Name: "resourcesForChannelCJTC-r19", Optional: true},
					{Name: "mr-SelectedResources-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ResourcesForChannelCLI_r19 != nil, ie.ResourcesForChannelCJTC_r19 != nil, ie.Mr_SelectedResources_r19 != nil}); err != nil {
				return err
			}

			if ie.ResourcesForChannelCLI_r19 != nil {
				c := ie.ResourcesForChannelCLI_r19
				cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19Seq := ex.NewSequenceEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19Constraints)
				if err := cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19Seq.EncodePreamble([]bool{c.ResourceSetCLI_r19 != nil, c.Qcl_InfoCLI_r19 != nil}); err != nil {
					return err
				}
				if c.ResourceSetCLI_r19 != nil {
					if err := ex.EncodeInteger((*c.ResourceSetCLI_r19), per.Constrained(1, common.MaxNrofCLI_MeasResourceSetsPerConfig_r19)); err != nil {
						return err
					}
				}
				if c.Qcl_InfoCLI_r19 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19QclInfoCLIR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.Qcl_InfoCLI_r19))); err != nil {
						return err
					}
					for i := range c.Qcl_InfoCLI_r19 {
						if err := c.Qcl_InfoCLI_r19[i].Encode(ex); err != nil {
							return err
						}
					}
				}
			}

			if ie.ResourcesForChannelCJTC_r19 != nil {
				c := ie.ResourcesForChannelCJTC_r19
				cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19Seq := ex.NewSequenceEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19Constraints)
				if err := cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19Seq.EncodePreamble([]bool{c.Nzp_CSI_RS3_r19 != nil, c.Nzp_CSI_RS4_r19 != nil}); err != nil {
					return err
				}
				{
					c := &c.Nzp_CSI_RS2_r19
					cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19Seq := ex.NewSequenceEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19Constraints)
					if err := cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19Seq.EncodePreamble([]bool{c.Qcl_Info2CJTC_r19 != nil}); err != nil {
						return err
					}
					if err := ex.EncodeInteger(c.ResourceSet2CJTC_r19, per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig)); err != nil {
						return err
					}
					if c.Qcl_Info2CJTC_r19 != nil {
						seqOf := ex.NewSequenceOfEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19QclInfo2CJTCR19Constraints)
						if err := seqOf.EncodeLength(int64(len(c.Qcl_Info2CJTC_r19))); err != nil {
							return err
						}
						for i := range c.Qcl_Info2CJTC_r19 {
							if err := c.Qcl_Info2CJTC_r19[i].Encode(ex); err != nil {
								return err
							}
						}
					}
				}
				if c.Nzp_CSI_RS3_r19 != nil {
					c := (*c.Nzp_CSI_RS3_r19)
					cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19Seq := ex.NewSequenceEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19Constraints)
					if err := cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19Seq.EncodePreamble([]bool{c.Qcl_Info3CJTC_r19 != nil}); err != nil {
						return err
					}
					if err := ex.EncodeInteger(c.ResourceSet3CJTC_r19, per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig)); err != nil {
						return err
					}
					if c.Qcl_Info3CJTC_r19 != nil {
						seqOf := ex.NewSequenceOfEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19QclInfo3CJTCR19Constraints)
						if err := seqOf.EncodeLength(int64(len(c.Qcl_Info3CJTC_r19))); err != nil {
							return err
						}
						for i := range c.Qcl_Info3CJTC_r19 {
							if err := c.Qcl_Info3CJTC_r19[i].Encode(ex); err != nil {
								return err
							}
						}
					}
				}
				if c.Nzp_CSI_RS4_r19 != nil {
					c := (*c.Nzp_CSI_RS4_r19)
					cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19Seq := ex.NewSequenceEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19Constraints)
					if err := cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19Seq.EncodePreamble([]bool{c.Qcl_Info4CJTC_r19 != nil}); err != nil {
						return err
					}
					if err := ex.EncodeInteger(c.ResourceSet4CJTC_r19, per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig)); err != nil {
						return err
					}
					if c.Qcl_Info4CJTC_r19 != nil {
						seqOf := ex.NewSequenceOfEncoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19QclInfo4CJTCR19Constraints)
						if err := seqOf.EncodeLength(int64(len(c.Qcl_Info4CJTC_r19))); err != nil {
							return err
						}
						for i := range c.Qcl_Info4CJTC_r19 {
							if err := c.Qcl_Info4CJTC_r19[i].Encode(ex); err != nil {
								return err
							}
						}
					}
				}
			}

			if ie.Mr_SelectedResources_r19 != nil {
				c := ie.Mr_SelectedResources_r19
				cSIAssociatedReportConfigInfoExtMrSelectedResourcesR19Seq := ex.NewSequenceEncoder(cSIAssociatedReportConfigInfoExtMrSelectedResourcesR19Constraints)
				if err := cSIAssociatedReportConfigInfoExtMrSelectedResourcesR19Seq.EncodePreamble([]bool{c.SecondSelectedResource_r19 != nil}); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.FirstSelectedResource_r19, per.Constrained(1, 8)); err != nil {
					return err
				}
				if c.SecondSelectedResource_r19 != nil {
					if err := ex.EncodeInteger((*c.SecondSelectedResource_r19), per.Constrained(1, 8)); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CSI_AssociatedReportConfigInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIAssociatedReportConfigInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. reportConfigId: ref
	{
		if err := ie.ReportConfigId.Decode(d); err != nil {
			return err
		}
	}

	// 4. resourcesForChannel: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_AssociatedReportConfigInfoResourcesForChannelConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ResourcesForChannel.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_AssociatedReportConfigInfo_ResourcesForChannel_Nzp_CSI_RS:
			ie.ResourcesForChannel.Nzp_CSI_RS = &struct {
				ResourceSet int64
				Qcl_Info    []TCI_StateId
			}{}
			c := (*ie.ResourcesForChannel.Nzp_CSI_RS)
			cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSSeq := d.NewSequenceDecoder(cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSConstraints)
			if err := cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig))
				if err != nil {
					return err
				}
				c.ResourceSet = v
			}
			if cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSSeq.IsComponentPresent(1) {
				seqOf := d.NewSequenceOfDecoder(cSIAssociatedReportConfigInfoResourcesForChannelNzpCSIRSQclInfoConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Qcl_Info = make([]TCI_StateId, n)
				for i := int64(0); i < n; i++ {
					if err := c.Qcl_Info[i].Decode(d); err != nil {
						return err
					}
				}
			}
		case CSI_AssociatedReportConfigInfo_ResourcesForChannel_Csi_SSB_ResourceSet:
			ie.ResourcesForChannel.Csi_SSB_ResourceSet = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofCSI_SSB_ResourceSetsPerConfig))
			if err != nil {
				return err
			}
			(*ie.ResourcesForChannel.Csi_SSB_ResourceSet) = v
		}
	}

	// 5. csi-IM-ResourcesForInterference: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(cSIAssociatedReportConfigInfoCsiIMResourcesForInterferenceConstraints)
			if err != nil {
				return err
			}
			ie.Csi_IM_ResourcesForInterference = &v
		}
	}

	// 6. nzp-CSI-RS-ResourcesForInterference: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(cSIAssociatedReportConfigInfoNzpCSIRSResourcesForInterferenceConstraints)
			if err != nil {
				return err
			}
			ie.Nzp_CSI_RS_ResourcesForInterference = &v
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "resourcesForChannel2-r17", Optional: true},
				{Name: "csi-SSB-ResourceSetExt", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ResourcesForChannel2_r17 = &struct {
				Choice          int
				Nzp_CSI_RS2_r17 *struct {
					ResourceSet2_r17 int64
					Qcl_Info2_r17    []TCI_StateId
				}
				Csi_SSB_ResourceSet2_r17 *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannel2R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ResourcesForChannel2_r17).Choice = int(index)
			switch index {
			case CSI_AssociatedReportConfigInfo_Ext_ResourcesForChannel2_r17_Nzp_CSI_RS2_r17:
				(*ie.ResourcesForChannel2_r17).Nzp_CSI_RS2_r17 = &struct {
					ResourceSet2_r17 int64
					Qcl_Info2_r17    []TCI_StateId
				}{}
				c := (*(*ie.ResourcesForChannel2_r17).Nzp_CSI_RS2_r17)
				cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17Seq := dx.NewSequenceDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17Constraints)
				if err := cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig))
					if err != nil {
						return err
					}
					c.ResourceSet2_r17 = v
				}
				if cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17Seq.IsComponentPresent(1) {
					seqOf := dx.NewSequenceOfDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannel2R17NzpCSIRS2R17QclInfo2R17Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Qcl_Info2_r17 = make([]TCI_StateId, n)
					for i := int64(0); i < n; i++ {
						if err := c.Qcl_Info2_r17[i].Decode(dx); err != nil {
							return err
						}
					}
				}
			case CSI_AssociatedReportConfigInfo_Ext_ResourcesForChannel2_r17_Csi_SSB_ResourceSet2_r17:
				(*ie.ResourcesForChannel2_r17).Csi_SSB_ResourceSet2_r17 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofCSI_SSB_ResourceSetsPerConfigExt))
				if err != nil {
					return err
				}
				(*(*ie.ResourcesForChannel2_r17).Csi_SSB_ResourceSet2_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(cSIAssociatedReportConfigInfoCsiSSBResourceSetExtConstraints)
			if err != nil {
				return err
			}
			ie.Csi_SSB_ResourceSetExt = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "resourcesForChannelTDCP-r18", Optional: true},
				{Name: "applyIndicatedTCI-State-r18", Optional: true},
				{Name: "applyIndicatedTCI-State2-r18", Optional: true},
				{Name: "csi-ReportSubConfigTriggerList-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ResourcesForChannelTDCP_r18 = &struct {
				ResourceSet2TDCP_r18 int64
				ResourceSet3TDCP_r18 *int64
			}{}
			c := ie.ResourcesForChannelTDCP_r18
			cSIAssociatedReportConfigInfoExtResourcesForChannelTDCPR18Seq := dx.NewSequenceDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannelTDCPR18Constraints)
			if err := cSIAssociatedReportConfigInfoExtResourcesForChannelTDCPR18Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig))
				if err != nil {
					return err
				}
				c.ResourceSet2TDCP_r18 = v
			}
			if cSIAssociatedReportConfigInfoExtResourcesForChannelTDCPR18Seq.IsComponentPresent(1) {
				c.ResourceSet3TDCP_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig))
				if err != nil {
					return err
				}
				(*c.ResourceSet3TDCP_r18) = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ApplyIndicatedTCI_State_r18 = &struct {
				Choice          int
				PerSet_r18      *int64
				PerResource_r18 *[]int64
			}{}
			choiceDec := dx.NewChoiceDecoder(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ApplyIndicatedTCI_State_r18).Choice = int(index)
			switch index {
			case CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerSet_r18:
				(*ie.ApplyIndicatedTCI_State_r18).PerSet_r18 = new(int64)
				v, err := dx.DecodeEnumerated(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18PerSetR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.ApplyIndicatedTCI_State_r18).PerSet_r18) = v
			case CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State_r18_PerResource_r18:
				(*ie.ApplyIndicatedTCI_State_r18).PerResource_r18 = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18PerResourceR18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.ApplyIndicatedTCI_State_r18).PerResource_r18) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeEnumerated(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIStateR18PerResourceR18ElemConstraints)
					if err != nil {
						return err
					}
					(*(*ie.ApplyIndicatedTCI_State_r18).PerResource_r18)[i] = v
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.ApplyIndicatedTCI_State2_r18 = &struct {
				Choice          int
				PerSet_r18      *int64
				PerResource_r18 *[]int64
			}{}
			choiceDec := dx.NewChoiceDecoder(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ApplyIndicatedTCI_State2_r18).Choice = int(index)
			switch index {
			case CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerSet_r18:
				(*ie.ApplyIndicatedTCI_State2_r18).PerSet_r18 = new(int64)
				v, err := dx.DecodeEnumerated(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18PerSetR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.ApplyIndicatedTCI_State2_r18).PerSet_r18) = v
			case CSI_AssociatedReportConfigInfo_Ext_ApplyIndicatedTCI_State2_r18_PerResource_r18:
				(*ie.ApplyIndicatedTCI_State2_r18).PerResource_r18 = new([]int64)
				seqOf := dx.NewSequenceOfDecoder(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18PerResourceR18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.ApplyIndicatedTCI_State2_r18).PerResource_r18) = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeEnumerated(cSIAssociatedReportConfigInfoExtApplyIndicatedTCIState2R18PerResourceR18ElemConstraints)
					if err != nil {
						return err
					}
					(*(*ie.ApplyIndicatedTCI_State2_r18).PerResource_r18)[i] = v
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Csi_ReportSubConfigTriggerList_r18 = new(CSI_ReportSubConfigTriggerList_r18)
			if err := ie.Csi_ReportSubConfigTriggerList_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "resourcesForChannelCLI-r19", Optional: true},
				{Name: "resourcesForChannelCJTC-r19", Optional: true},
				{Name: "mr-SelectedResources-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ResourcesForChannelCLI_r19 = &struct {
				ResourceSetCLI_r19 *int64
				Qcl_InfoCLI_r19    []TCI_StateId
			}{}
			c := ie.ResourcesForChannelCLI_r19
			cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19Seq := dx.NewSequenceDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19Constraints)
			if err := cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19Seq.IsComponentPresent(0) {
				c.ResourceSetCLI_r19 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofCLI_MeasResourceSetsPerConfig_r19))
				if err != nil {
					return err
				}
				(*c.ResourceSetCLI_r19) = v
			}
			if cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19Seq.IsComponentPresent(1) {
				seqOf := dx.NewSequenceOfDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCLIR19QclInfoCLIR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Qcl_InfoCLI_r19 = make([]TCI_StateId, n)
				for i := int64(0); i < n; i++ {
					if err := c.Qcl_InfoCLI_r19[i].Decode(dx); err != nil {
						return err
					}
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ResourcesForChannelCJTC_r19 = &struct {
				Nzp_CSI_RS2_r19 struct {
					ResourceSet2CJTC_r19 int64
					Qcl_Info2CJTC_r19    []TCI_StateId
				}
				Nzp_CSI_RS3_r19 *struct {
					ResourceSet3CJTC_r19 int64
					Qcl_Info3CJTC_r19    []TCI_StateId
				}
				Nzp_CSI_RS4_r19 *struct {
					ResourceSet4CJTC_r19 int64
					Qcl_Info4CJTC_r19    []TCI_StateId
				}
			}{}
			c := ie.ResourcesForChannelCJTC_r19
			cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19Seq := dx.NewSequenceDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19Constraints)
			if err := cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				c := &c.Nzp_CSI_RS2_r19
				cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19Seq := dx.NewSequenceDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19Constraints)
				if err := cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig))
					if err != nil {
						return err
					}
					c.ResourceSet2CJTC_r19 = v
				}
				if cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19Seq.IsComponentPresent(1) {
					seqOf := dx.NewSequenceOfDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS2R19QclInfo2CJTCR19Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Qcl_Info2CJTC_r19 = make([]TCI_StateId, n)
					for i := int64(0); i < n; i++ {
						if err := c.Qcl_Info2CJTC_r19[i].Decode(dx); err != nil {
							return err
						}
					}
				}
			}
			if cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19Seq.IsComponentPresent(1) {
				c.Nzp_CSI_RS3_r19 = &struct {
					ResourceSet3CJTC_r19 int64
					Qcl_Info3CJTC_r19    []TCI_StateId
				}{}
				c := (*c.Nzp_CSI_RS3_r19)
				cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19Seq := dx.NewSequenceDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19Constraints)
				if err := cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig))
					if err != nil {
						return err
					}
					c.ResourceSet3CJTC_r19 = v
				}
				if cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19Seq.IsComponentPresent(1) {
					seqOf := dx.NewSequenceOfDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS3R19QclInfo3CJTCR19Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Qcl_Info3CJTC_r19 = make([]TCI_StateId, n)
					for i := int64(0); i < n; i++ {
						if err := c.Qcl_Info3CJTC_r19[i].Decode(dx); err != nil {
							return err
						}
					}
				}
			}
			if cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19Seq.IsComponentPresent(2) {
				c.Nzp_CSI_RS4_r19 = &struct {
					ResourceSet4CJTC_r19 int64
					Qcl_Info4CJTC_r19    []TCI_StateId
				}{}
				c := (*c.Nzp_CSI_RS4_r19)
				cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19Seq := dx.NewSequenceDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19Constraints)
				if err := cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig))
					if err != nil {
						return err
					}
					c.ResourceSet4CJTC_r19 = v
				}
				if cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19Seq.IsComponentPresent(1) {
					seqOf := dx.NewSequenceOfDecoder(cSIAssociatedReportConfigInfoExtResourcesForChannelCJTCR19NzpCSIRS4R19QclInfo4CJTCR19Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Qcl_Info4CJTC_r19 = make([]TCI_StateId, n)
					for i := int64(0); i < n; i++ {
						if err := c.Qcl_Info4CJTC_r19[i].Decode(dx); err != nil {
							return err
						}
					}
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Mr_SelectedResources_r19 = &struct {
				FirstSelectedResource_r19  int64
				SecondSelectedResource_r19 *int64
			}{}
			c := ie.Mr_SelectedResources_r19
			cSIAssociatedReportConfigInfoExtMrSelectedResourcesR19Seq := dx.NewSequenceDecoder(cSIAssociatedReportConfigInfoExtMrSelectedResourcesR19Constraints)
			if err := cSIAssociatedReportConfigInfoExtMrSelectedResourcesR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.FirstSelectedResource_r19 = v
			}
			if cSIAssociatedReportConfigInfoExtMrSelectedResourcesR19Seq.IsComponentPresent(1) {
				c.SecondSelectedResource_r19 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.SecondSelectedResource_r19) = v
			}
		}
	}

	return nil
}
