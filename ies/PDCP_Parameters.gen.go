// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCP-Parameters (line 22707).

var pDCPParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedROHC-Profiles"},
		{Name: "maxNumberROHC-ContextSessions"},
		{Name: "uplinkOnlyROHC-Profiles", Optional: true},
		{Name: "continueROHC-Context", Optional: true},
		{Name: "outOfOrderDelivery", Optional: true},
		{Name: "shortSN", Optional: true},
		{Name: "pdcp-DuplicationSRB", Optional: true},
		{Name: "pdcp-DuplicationMCG-OrSCG-DRB", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs2     = 0
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs4     = 1
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs8     = 2
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs12    = 3
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs16    = 4
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs24    = 5
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs32    = 6
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs48    = 7
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs64    = 8
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs128   = 9
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs256   = 10
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs512   = 11
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs1024  = 12
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs16384 = 13
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Spare2  = 14
	PDCP_Parameters_MaxNumberROHC_ContextSessions_Spare1  = 15
)

var pDCPParametersMaxNumberROHCContextSessionsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs2, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs4, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs8, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs12, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs16, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs24, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs32, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs48, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs64, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs128, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs256, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs512, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs1024, PDCP_Parameters_MaxNumberROHC_ContextSessions_Cs16384, PDCP_Parameters_MaxNumberROHC_ContextSessions_Spare2, PDCP_Parameters_MaxNumberROHC_ContextSessions_Spare1},
}

const (
	PDCP_Parameters_UplinkOnlyROHC_Profiles_Supported = 0
)

var pDCPParametersUplinkOnlyROHCProfilesConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_UplinkOnlyROHC_Profiles_Supported},
}

const (
	PDCP_Parameters_ContinueROHC_Context_Supported = 0
)

var pDCPParametersContinueROHCContextConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_ContinueROHC_Context_Supported},
}

const (
	PDCP_Parameters_OutOfOrderDelivery_Supported = 0
)

var pDCPParametersOutOfOrderDeliveryConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_OutOfOrderDelivery_Supported},
}

const (
	PDCP_Parameters_ShortSN_Supported = 0
)

var pDCPParametersShortSNConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_ShortSN_Supported},
}

const (
	PDCP_Parameters_Pdcp_DuplicationSRB_Supported = 0
)

var pDCPParametersPdcpDuplicationSRBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Pdcp_DuplicationSRB_Supported},
}

const (
	PDCP_Parameters_Pdcp_DuplicationMCG_OrSCG_DRB_Supported = 0
)

var pDCPParametersPdcpDuplicationMCGOrSCGDRBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Pdcp_DuplicationMCG_OrSCG_DRB_Supported},
}

const (
	PDCP_Parameters_Ext_Drb_IAB_r16_Supported = 0
)

var pDCPParametersExtDrbIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_Drb_IAB_r16_Supported},
}

const (
	PDCP_Parameters_Ext_Non_DRB_IAB_r16_Supported = 0
)

var pDCPParametersExtNonDRBIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_Non_DRB_IAB_r16_Supported},
}

const (
	PDCP_Parameters_Ext_ExtendedDiscardTimer_r16_Supported = 0
)

var pDCPParametersExtExtendedDiscardTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_ExtendedDiscardTimer_r16_Supported},
}

const (
	PDCP_Parameters_Ext_ContinueEHC_Context_r16_Supported = 0
)

var pDCPParametersExtContinueEHCContextR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_ContinueEHC_Context_r16_Supported},
}

const (
	PDCP_Parameters_Ext_Ehc_r16_Supported = 0
)

var pDCPParametersExtEhcR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_Ehc_r16_Supported},
}

const (
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs2     = 0
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs4     = 1
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs8     = 2
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs16    = 3
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs32    = 4
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs64    = 5
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs128   = 6
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs256   = 7
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs512   = 8
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs1024  = 9
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs2048  = 10
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs4096  = 11
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs8192  = 12
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs16384 = 13
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs32768 = 14
	PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs65536 = 15
)

var pDCPParametersExtMaxNumberEHCContextsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs2, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs4, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs8, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs16, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs32, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs64, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs128, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs256, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs512, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs1024, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs2048, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs4096, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs8192, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs16384, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs32768, PDCP_Parameters_Ext_MaxNumberEHC_Contexts_r16_Cs65536},
}

const (
	PDCP_Parameters_Ext_JointEHC_ROHC_Config_r16_Supported = 0
)

var pDCPParametersExtJointEHCROHCConfigR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_JointEHC_ROHC_Config_r16_Supported},
}

const (
	PDCP_Parameters_Ext_Pdcp_DuplicationMoreThanTwoRLC_r16_Supported = 0
)

var pDCPParametersExtPdcpDuplicationMoreThanTwoRLCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_Pdcp_DuplicationMoreThanTwoRLC_r16_Supported},
}

const (
	PDCP_Parameters_Ext_LongSN_RedCap_r17_Supported = 0
)

var pDCPParametersExtLongSNRedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_LongSN_RedCap_r17_Supported},
}

const (
	PDCP_Parameters_Ext_LongSN_NCR_r18_Supported = 0
)

var pDCPParametersExtLongSNNCRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_LongSN_NCR_r18_Supported},
}

const (
	PDCP_Parameters_Ext_SupportOfPDU_SetDiscard_r18_Supported = 0
)

var pDCPParametersExtSupportOfPDUSetDiscardR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_SupportOfPDU_SetDiscard_r18_Supported},
}

const (
	PDCP_Parameters_Ext_Psi_BasedDiscard_r18_Supported = 0
)

var pDCPParametersExtPsiBasedDiscardR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_Psi_BasedDiscard_r18_Supported},
}

const (
	PDCP_Parameters_Ext_SupportOfSN_GapReport_r18_Supported = 0
)

var pDCPParametersExtSupportOfSNGapReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_SupportOfSN_GapReport_r18_Supported},
}

var pDCPParametersExtUdcR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "standardDictionary-r17", Optional: true},
		{Name: "operatorDictionary-r17", Optional: true},
		{Name: "continueUDC-r17", Optional: true},
		{Name: "supportOfBufferSize-r17", Optional: true},
	},
}

const (
	PDCP_Parameters_Ext_Udc_r17_StandardDictionary_r17_Supported = 0
)

var pDCPParametersExtUdcR17StandardDictionaryR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_Udc_r17_StandardDictionary_r17_Supported},
}

const (
	PDCP_Parameters_Ext_Udc_r17_ContinueUDC_r17_Supported = 0
)

var pDCPParametersExtUdcR17ContinueUDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_Udc_r17_ContinueUDC_r17_Supported},
}

const (
	PDCP_Parameters_Ext_Udc_r17_SupportOfBufferSize_r17_Kbyte4 = 0
	PDCP_Parameters_Ext_Udc_r17_SupportOfBufferSize_r17_Kbyte8 = 1
)

var pDCPParametersExtUdcR17SupportOfBufferSizeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_Parameters_Ext_Udc_r17_SupportOfBufferSize_r17_Kbyte4, PDCP_Parameters_Ext_Udc_r17_SupportOfBufferSize_r17_Kbyte8},
}

type PDCP_Parameters struct {
	SupportedROHC_Profiles struct {
		Profile0x0000 bool
		Profile0x0001 bool
		Profile0x0002 bool
		Profile0x0003 bool
		Profile0x0004 bool
		Profile0x0006 bool
		Profile0x0101 bool
		Profile0x0102 bool
		Profile0x0103 bool
		Profile0x0104 bool
	}
	MaxNumberROHC_ContextSessions      int64
	UplinkOnlyROHC_Profiles            *int64
	ContinueROHC_Context               *int64
	OutOfOrderDelivery                 *int64
	ShortSN                            *int64
	Pdcp_DuplicationSRB                *int64
	Pdcp_DuplicationMCG_OrSCG_DRB      *int64
	Drb_IAB_r16                        *int64
	Non_DRB_IAB_r16                    *int64
	ExtendedDiscardTimer_r16           *int64
	ContinueEHC_Context_r16            *int64
	Ehc_r16                            *int64
	MaxNumberEHC_Contexts_r16          *int64
	JointEHC_ROHC_Config_r16           *int64
	Pdcp_DuplicationMoreThanTwoRLC_r16 *int64
	LongSN_RedCap_r17                  *int64
	Udc_r17                            *struct {
		StandardDictionary_r17 *int64
		OperatorDictionary_r17 *struct {
			VersionOfDictionary_r17 int64
			AssociatedPLMN_ID_r17   PLMN_Identity
		}
		ContinueUDC_r17         *int64
		SupportOfBufferSize_r17 *int64
	}
	LongSN_NCR_r18              *int64
	SupportOfPDU_SetDiscard_r18 *int64
	Psi_BasedDiscard_r18        *int64
	SupportOfSN_GapReport_r18   *int64
}

func (ie *PDCP_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCPParametersConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Drb_IAB_r16 != nil || ie.Non_DRB_IAB_r16 != nil || ie.ExtendedDiscardTimer_r16 != nil || ie.ContinueEHC_Context_r16 != nil || ie.Ehc_r16 != nil || ie.MaxNumberEHC_Contexts_r16 != nil || ie.JointEHC_ROHC_Config_r16 != nil || ie.Pdcp_DuplicationMoreThanTwoRLC_r16 != nil
	hasExtGroup1 := ie.LongSN_RedCap_r17 != nil || ie.Udc_r17 != nil
	hasExtGroup2 := ie.LongSN_NCR_r18 != nil || ie.SupportOfPDU_SetDiscard_r18 != nil || ie.Psi_BasedDiscard_r18 != nil || ie.SupportOfSN_GapReport_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkOnlyROHC_Profiles != nil, ie.ContinueROHC_Context != nil, ie.OutOfOrderDelivery != nil, ie.ShortSN != nil, ie.Pdcp_DuplicationSRB != nil, ie.Pdcp_DuplicationMCG_OrSCG_DRB != nil}); err != nil {
		return err
	}

	// 3. supportedROHC-Profiles: sequence
	{
		{
			c := &ie.SupportedROHC_Profiles
			if err := e.EncodeBoolean(c.Profile0x0000); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Profile0x0001); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Profile0x0002); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Profile0x0003); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Profile0x0004); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Profile0x0006); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Profile0x0101); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Profile0x0102); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Profile0x0103); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.Profile0x0104); err != nil {
				return err
			}
		}
	}

	// 4. maxNumberROHC-ContextSessions: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberROHC_ContextSessions, pDCPParametersMaxNumberROHCContextSessionsConstraints); err != nil {
			return err
		}
	}

	// 5. uplinkOnlyROHC-Profiles: enumerated
	{
		if ie.UplinkOnlyROHC_Profiles != nil {
			if err := e.EncodeEnumerated(*ie.UplinkOnlyROHC_Profiles, pDCPParametersUplinkOnlyROHCProfilesConstraints); err != nil {
				return err
			}
		}
	}

	// 6. continueROHC-Context: enumerated
	{
		if ie.ContinueROHC_Context != nil {
			if err := e.EncodeEnumerated(*ie.ContinueROHC_Context, pDCPParametersContinueROHCContextConstraints); err != nil {
				return err
			}
		}
	}

	// 7. outOfOrderDelivery: enumerated
	{
		if ie.OutOfOrderDelivery != nil {
			if err := e.EncodeEnumerated(*ie.OutOfOrderDelivery, pDCPParametersOutOfOrderDeliveryConstraints); err != nil {
				return err
			}
		}
	}

	// 8. shortSN: enumerated
	{
		if ie.ShortSN != nil {
			if err := e.EncodeEnumerated(*ie.ShortSN, pDCPParametersShortSNConstraints); err != nil {
				return err
			}
		}
	}

	// 9. pdcp-DuplicationSRB: enumerated
	{
		if ie.Pdcp_DuplicationSRB != nil {
			if err := e.EncodeEnumerated(*ie.Pdcp_DuplicationSRB, pDCPParametersPdcpDuplicationSRBConstraints); err != nil {
				return err
			}
		}
	}

	// 10. pdcp-DuplicationMCG-OrSCG-DRB: enumerated
	{
		if ie.Pdcp_DuplicationMCG_OrSCG_DRB != nil {
			if err := e.EncodeEnumerated(*ie.Pdcp_DuplicationMCG_OrSCG_DRB, pDCPParametersPdcpDuplicationMCGOrSCGDRBConstraints); err != nil {
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
					{Name: "drb-IAB-r16", Optional: true},
					{Name: "non-DRB-IAB-r16", Optional: true},
					{Name: "extendedDiscardTimer-r16", Optional: true},
					{Name: "continueEHC-Context-r16", Optional: true},
					{Name: "ehc-r16", Optional: true},
					{Name: "maxNumberEHC-Contexts-r16", Optional: true},
					{Name: "jointEHC-ROHC-Config-r16", Optional: true},
					{Name: "pdcp-DuplicationMoreThanTwoRLC-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Drb_IAB_r16 != nil, ie.Non_DRB_IAB_r16 != nil, ie.ExtendedDiscardTimer_r16 != nil, ie.ContinueEHC_Context_r16 != nil, ie.Ehc_r16 != nil, ie.MaxNumberEHC_Contexts_r16 != nil, ie.JointEHC_ROHC_Config_r16 != nil, ie.Pdcp_DuplicationMoreThanTwoRLC_r16 != nil}); err != nil {
				return err
			}

			if ie.Drb_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Drb_IAB_r16, pDCPParametersExtDrbIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.Non_DRB_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Non_DRB_IAB_r16, pDCPParametersExtNonDRBIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.ExtendedDiscardTimer_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ExtendedDiscardTimer_r16, pDCPParametersExtExtendedDiscardTimerR16Constraints); err != nil {
					return err
				}
			}

			if ie.ContinueEHC_Context_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ContinueEHC_Context_r16, pDCPParametersExtContinueEHCContextR16Constraints); err != nil {
					return err
				}
			}

			if ie.Ehc_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Ehc_r16, pDCPParametersExtEhcR16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberEHC_Contexts_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberEHC_Contexts_r16, pDCPParametersExtMaxNumberEHCContextsR16Constraints); err != nil {
					return err
				}
			}

			if ie.JointEHC_ROHC_Config_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.JointEHC_ROHC_Config_r16, pDCPParametersExtJointEHCROHCConfigR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcp_DuplicationMoreThanTwoRLC_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcp_DuplicationMoreThanTwoRLC_r16, pDCPParametersExtPdcpDuplicationMoreThanTwoRLCR16Constraints); err != nil {
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
					{Name: "longSN-RedCap-r17", Optional: true},
					{Name: "udc-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.LongSN_RedCap_r17 != nil, ie.Udc_r17 != nil}); err != nil {
				return err
			}

			if ie.LongSN_RedCap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.LongSN_RedCap_r17, pDCPParametersExtLongSNRedCapR17Constraints); err != nil {
					return err
				}
			}

			if ie.Udc_r17 != nil {
				c := ie.Udc_r17
				pDCPParametersExtUdcR17Seq := ex.NewSequenceEncoder(pDCPParametersExtUdcR17Constraints)
				if err := pDCPParametersExtUdcR17Seq.EncodePreamble([]bool{c.StandardDictionary_r17 != nil, c.OperatorDictionary_r17 != nil, c.ContinueUDC_r17 != nil, c.SupportOfBufferSize_r17 != nil}); err != nil {
					return err
				}
				if c.StandardDictionary_r17 != nil {
					if err := ex.EncodeEnumerated((*c.StandardDictionary_r17), pDCPParametersExtUdcR17StandardDictionaryR17Constraints); err != nil {
						return err
					}
				}
				if c.OperatorDictionary_r17 != nil {
					c := (*c.OperatorDictionary_r17)
					if err := ex.EncodeInteger(c.VersionOfDictionary_r17, per.Constrained(0, 15)); err != nil {
						return err
					}
					if err := c.AssociatedPLMN_ID_r17.Encode(ex); err != nil {
						return err
					}
				}
				if c.ContinueUDC_r17 != nil {
					if err := ex.EncodeEnumerated((*c.ContinueUDC_r17), pDCPParametersExtUdcR17ContinueUDCR17Constraints); err != nil {
						return err
					}
				}
				if c.SupportOfBufferSize_r17 != nil {
					if err := ex.EncodeEnumerated((*c.SupportOfBufferSize_r17), pDCPParametersExtUdcR17SupportOfBufferSizeR17Constraints); err != nil {
						return err
					}
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
					{Name: "longSN-NCR-r18", Optional: true},
					{Name: "supportOfPDU-SetDiscard-r18", Optional: true},
					{Name: "psi-BasedDiscard-r18", Optional: true},
					{Name: "supportOfSN-GapReport-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.LongSN_NCR_r18 != nil, ie.SupportOfPDU_SetDiscard_r18 != nil, ie.Psi_BasedDiscard_r18 != nil, ie.SupportOfSN_GapReport_r18 != nil}); err != nil {
				return err
			}

			if ie.LongSN_NCR_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.LongSN_NCR_r18, pDCPParametersExtLongSNNCRR18Constraints); err != nil {
					return err
				}
			}

			if ie.SupportOfPDU_SetDiscard_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportOfPDU_SetDiscard_r18, pDCPParametersExtSupportOfPDUSetDiscardR18Constraints); err != nil {
					return err
				}
			}

			if ie.Psi_BasedDiscard_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Psi_BasedDiscard_r18, pDCPParametersExtPsiBasedDiscardR18Constraints); err != nil {
					return err
				}
			}

			if ie.SupportOfSN_GapReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportOfSN_GapReport_r18, pDCPParametersExtSupportOfSNGapReportR18Constraints); err != nil {
					return err
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

func (ie *PDCP_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCPParametersConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. supportedROHC-Profiles: sequence
	{
		{
			c := &ie.SupportedROHC_Profiles
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Profile0x0000 = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Profile0x0001 = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Profile0x0002 = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Profile0x0003 = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Profile0x0004 = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Profile0x0006 = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Profile0x0101 = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Profile0x0102 = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Profile0x0103 = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.Profile0x0104 = v
			}
		}
	}

	// 4. maxNumberROHC-ContextSessions: enumerated
	{
		v1, err := d.DecodeEnumerated(pDCPParametersMaxNumberROHCContextSessionsConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberROHC_ContextSessions = v1
	}

	// 5. uplinkOnlyROHC-Profiles: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pDCPParametersUplinkOnlyROHCProfilesConstraints)
			if err != nil {
				return err
			}
			ie.UplinkOnlyROHC_Profiles = &idx
		}
	}

	// 6. continueROHC-Context: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(pDCPParametersContinueROHCContextConstraints)
			if err != nil {
				return err
			}
			ie.ContinueROHC_Context = &idx
		}
	}

	// 7. outOfOrderDelivery: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(pDCPParametersOutOfOrderDeliveryConstraints)
			if err != nil {
				return err
			}
			ie.OutOfOrderDelivery = &idx
		}
	}

	// 8. shortSN: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(pDCPParametersShortSNConstraints)
			if err != nil {
				return err
			}
			ie.ShortSN = &idx
		}
	}

	// 9. pdcp-DuplicationSRB: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(pDCPParametersPdcpDuplicationSRBConstraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationSRB = &idx
		}
	}

	// 10. pdcp-DuplicationMCG-OrSCG-DRB: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(pDCPParametersPdcpDuplicationMCGOrSCGDRBConstraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationMCG_OrSCG_DRB = &idx
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
				{Name: "drb-IAB-r16", Optional: true},
				{Name: "non-DRB-IAB-r16", Optional: true},
				{Name: "extendedDiscardTimer-r16", Optional: true},
				{Name: "continueEHC-Context-r16", Optional: true},
				{Name: "ehc-r16", Optional: true},
				{Name: "maxNumberEHC-Contexts-r16", Optional: true},
				{Name: "jointEHC-ROHC-Config-r16", Optional: true},
				{Name: "pdcp-DuplicationMoreThanTwoRLC-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtDrbIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Drb_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtNonDRBIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Non_DRB_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtExtendedDiscardTimerR16Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedDiscardTimer_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtContinueEHCContextR16Constraints)
			if err != nil {
				return err
			}
			ie.ContinueEHC_Context_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtEhcR16Constraints)
			if err != nil {
				return err
			}
			ie.Ehc_r16 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtMaxNumberEHCContextsR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberEHC_Contexts_r16 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtJointEHCROHCConfigR16Constraints)
			if err != nil {
				return err
			}
			ie.JointEHC_ROHC_Config_r16 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtPdcpDuplicationMoreThanTwoRLCR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationMoreThanTwoRLC_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "longSN-RedCap-r17", Optional: true},
				{Name: "udc-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtLongSNRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.LongSN_RedCap_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Udc_r17 = &struct {
				StandardDictionary_r17 *int64
				OperatorDictionary_r17 *struct {
					VersionOfDictionary_r17 int64
					AssociatedPLMN_ID_r17   PLMN_Identity
				}
				ContinueUDC_r17         *int64
				SupportOfBufferSize_r17 *int64
			}{}
			c := ie.Udc_r17
			pDCPParametersExtUdcR17Seq := dx.NewSequenceDecoder(pDCPParametersExtUdcR17Constraints)
			if err := pDCPParametersExtUdcR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if pDCPParametersExtUdcR17Seq.IsComponentPresent(0) {
				c.StandardDictionary_r17 = new(int64)
				v, err := dx.DecodeEnumerated(pDCPParametersExtUdcR17StandardDictionaryR17Constraints)
				if err != nil {
					return err
				}
				(*c.StandardDictionary_r17) = v
			}
			if pDCPParametersExtUdcR17Seq.IsComponentPresent(1) {
				c.OperatorDictionary_r17 = &struct {
					VersionOfDictionary_r17 int64
					AssociatedPLMN_ID_r17   PLMN_Identity
				}{}
				c := (*c.OperatorDictionary_r17)
				{
					v, err := dx.DecodeInteger(per.Constrained(0, 15))
					if err != nil {
						return err
					}
					c.VersionOfDictionary_r17 = v
				}
				{
					if err := c.AssociatedPLMN_ID_r17.Decode(dx); err != nil {
						return err
					}
				}
			}
			if pDCPParametersExtUdcR17Seq.IsComponentPresent(2) {
				c.ContinueUDC_r17 = new(int64)
				v, err := dx.DecodeEnumerated(pDCPParametersExtUdcR17ContinueUDCR17Constraints)
				if err != nil {
					return err
				}
				(*c.ContinueUDC_r17) = v
			}
			if pDCPParametersExtUdcR17Seq.IsComponentPresent(3) {
				c.SupportOfBufferSize_r17 = new(int64)
				v, err := dx.DecodeEnumerated(pDCPParametersExtUdcR17SupportOfBufferSizeR17Constraints)
				if err != nil {
					return err
				}
				(*c.SupportOfBufferSize_r17) = v
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
				{Name: "longSN-NCR-r18", Optional: true},
				{Name: "supportOfPDU-SetDiscard-r18", Optional: true},
				{Name: "psi-BasedDiscard-r18", Optional: true},
				{Name: "supportOfSN-GapReport-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtLongSNNCRR18Constraints)
			if err != nil {
				return err
			}
			ie.LongSN_NCR_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtSupportOfPDUSetDiscardR18Constraints)
			if err != nil {
				return err
			}
			ie.SupportOfPDU_SetDiscard_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtPsiBasedDiscardR18Constraints)
			if err != nil {
				return err
			}
			ie.Psi_BasedDiscard_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(pDCPParametersExtSupportOfSNGapReportR18Constraints)
			if err != nil {
				return err
			}
			ie.SupportOfSN_GapReport_r18 = &v
		}
	}

	return nil
}
