// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SidelinkParametersNR-r16 (line 25020).

var sidelinkParametersNRR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rlc-ParametersSidelink-r16", Optional: true},
		{Name: "mac-ParametersSidelink-r16", Optional: true},
		{Name: "fdd-Add-UE-Sidelink-Capabilities-r16", Optional: true},
		{Name: "tdd-Add-UE-Sidelink-Capabilities-r16", Optional: true},
		{Name: "supportedBandListSidelink-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var sidelinkParametersNRR16SupportedBandListSidelinkR16Constraints = per.SizeRange(1, common.MaxBands)

const (
	SidelinkParametersNR_r16_Ext_P0_OLPC_Sidelink_r17_Supported = 0
)

var sidelinkParametersNRR16ExtP0OLPCSidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SidelinkParametersNR_r16_Ext_P0_OLPC_Sidelink_r17_Supported},
}

var sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr1-r18", Optional: true},
		{Name: "fr2-r18", Optional: true},
	},
}

const (
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N1  = 0
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N2  = 1
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N4  = 2
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N6  = 3
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N8  = 4
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N12 = 5
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N16 = 6
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N24 = 7
)

var sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Fr1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N1, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N2, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N4, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N6, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N8, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N12, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N16, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr1_r18_N24},
}

const (
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N1   = 0
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N2   = 1
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N4   = 2
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N6   = 3
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N8   = 4
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N12  = 5
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N16  = 6
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N24  = 7
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N32  = 8
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N48  = 9
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N64  = 10
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N128 = 11
)

var sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Fr2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N1, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N2, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N4, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N6, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N8, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N12, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N16, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N24, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N32, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N48, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N64, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfActiveSL_PRS_Resources_r18_Fr2_r18_N128},
}

var sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr1-r18", Optional: true},
		{Name: "fr2-r18", Optional: true},
	},
}

const (
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N1 = 0
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N2 = 1
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N3 = 2
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N4 = 3
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N6 = 4
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N8 = 5
)

var sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Fr1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N1, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N2, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N3, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N4, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N6, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr1_r18_N8},
}

const (
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N1  = 0
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N2  = 1
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N4  = 2
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N8  = 3
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N12 = 4
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N16 = 5
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N24 = 6
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N32 = 7
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N48 = 8
	SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N64 = 9
)

var sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Fr2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N1, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N2, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N4, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N8, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N12, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N16, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N24, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N32, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N48, SidelinkParametersNR_r16_Ext_Sl_PRS_CommonProcCapabilityPerUE_r18_MaxNumOfSlotswithActiveSL_PRS_Resources_r18_Fr2_r18_N64},
}

type SidelinkParametersNR_r16 struct {
	Rlc_ParametersSidelink_r16           *RLC_ParametersSidelink_r16
	Mac_ParametersSidelink_r16           *MAC_ParametersSidelink_r16
	Fdd_Add_UE_Sidelink_Capabilities_r16 *UE_SidelinkCapabilityAddXDD_Mode_r16
	Tdd_Add_UE_Sidelink_Capabilities_r16 *UE_SidelinkCapabilityAddXDD_Mode_r16
	SupportedBandListSidelink_r16        []BandSidelink_r16
	RelayParameters_r17                  *RelayParameters_r17
	P0_OLPC_Sidelink_r17                 *int64
	Pdcp_ParametersSidelink_r18          *PDCP_ParametersSidelink_r18
	Sl_PRS_CommonProcCapabilityPerUE_r18 *struct {
		MaxNumOfActiveSL_PRS_Resources_r18 struct {
			Fr1_r18 *int64
			Fr2_r18 *int64
		}
		MaxNumOfSlotswithActiveSL_PRS_Resources_r18 struct {
			Fr1_r18 *int64
			Fr2_r18 *int64
		}
	}
}

func (ie *SidelinkParametersNR_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sidelinkParametersNRR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.RelayParameters_r17 != nil
	hasExtGroup1 := ie.P0_OLPC_Sidelink_r17 != nil
	hasExtGroup2 := ie.Pdcp_ParametersSidelink_r18 != nil || ie.Sl_PRS_CommonProcCapabilityPerUE_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rlc_ParametersSidelink_r16 != nil, ie.Mac_ParametersSidelink_r16 != nil, ie.Fdd_Add_UE_Sidelink_Capabilities_r16 != nil, ie.Tdd_Add_UE_Sidelink_Capabilities_r16 != nil, ie.SupportedBandListSidelink_r16 != nil}); err != nil {
		return err
	}

	// 3. rlc-ParametersSidelink-r16: ref
	{
		if ie.Rlc_ParametersSidelink_r16 != nil {
			if err := ie.Rlc_ParametersSidelink_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. mac-ParametersSidelink-r16: ref
	{
		if ie.Mac_ParametersSidelink_r16 != nil {
			if err := ie.Mac_ParametersSidelink_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. fdd-Add-UE-Sidelink-Capabilities-r16: ref
	{
		if ie.Fdd_Add_UE_Sidelink_Capabilities_r16 != nil {
			if err := ie.Fdd_Add_UE_Sidelink_Capabilities_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. tdd-Add-UE-Sidelink-Capabilities-r16: ref
	{
		if ie.Tdd_Add_UE_Sidelink_Capabilities_r16 != nil {
			if err := ie.Tdd_Add_UE_Sidelink_Capabilities_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. supportedBandListSidelink-r16: sequence-of
	{
		if ie.SupportedBandListSidelink_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sidelinkParametersNRR16SupportedBandListSidelinkR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SupportedBandListSidelink_r16))); err != nil {
				return err
			}
			for i := range ie.SupportedBandListSidelink_r16 {
				if err := ie.SupportedBandListSidelink_r16[i].Encode(e); err != nil {
					return err
				}
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
					{Name: "relayParameters-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RelayParameters_r17 != nil}); err != nil {
				return err
			}

			if ie.RelayParameters_r17 != nil {
				if err := ie.RelayParameters_r17.Encode(ex); err != nil {
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
					{Name: "p0-OLPC-Sidelink-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.P0_OLPC_Sidelink_r17 != nil}); err != nil {
				return err
			}

			if ie.P0_OLPC_Sidelink_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.P0_OLPC_Sidelink_r17, sidelinkParametersNRR16ExtP0OLPCSidelinkR17Constraints); err != nil {
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
					{Name: "pdcp-ParametersSidelink-r18", Optional: true},
					{Name: "sl-PRS-CommonProcCapabilityPerUE-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdcp_ParametersSidelink_r18 != nil, ie.Sl_PRS_CommonProcCapabilityPerUE_r18 != nil}); err != nil {
				return err
			}

			if ie.Pdcp_ParametersSidelink_r18 != nil {
				if err := ie.Pdcp_ParametersSidelink_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_CommonProcCapabilityPerUE_r18 != nil {
				c := ie.Sl_PRS_CommonProcCapabilityPerUE_r18
				{
					c := &c.MaxNumOfActiveSL_PRS_Resources_r18
					sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Seq := ex.NewSequenceEncoder(sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Constraints)
					if err := sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Seq.EncodePreamble([]bool{c.Fr1_r18 != nil, c.Fr2_r18 != nil}); err != nil {
						return err
					}
					if c.Fr1_r18 != nil {
						if err := ex.EncodeEnumerated((*c.Fr1_r18), sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Fr1R18Constraints); err != nil {
							return err
						}
					}
					if c.Fr2_r18 != nil {
						if err := ex.EncodeEnumerated((*c.Fr2_r18), sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Fr2R18Constraints); err != nil {
							return err
						}
					}
				}
				{
					c := &c.MaxNumOfSlotswithActiveSL_PRS_Resources_r18
					sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Seq := ex.NewSequenceEncoder(sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Constraints)
					if err := sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Seq.EncodePreamble([]bool{c.Fr1_r18 != nil, c.Fr2_r18 != nil}); err != nil {
						return err
					}
					if c.Fr1_r18 != nil {
						if err := ex.EncodeEnumerated((*c.Fr1_r18), sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Fr1R18Constraints); err != nil {
							return err
						}
					}
					if c.Fr2_r18 != nil {
						if err := ex.EncodeEnumerated((*c.Fr2_r18), sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Fr2R18Constraints); err != nil {
							return err
						}
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

func (ie *SidelinkParametersNR_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sidelinkParametersNRR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rlc-ParametersSidelink-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Rlc_ParametersSidelink_r16 = new(RLC_ParametersSidelink_r16)
			if err := ie.Rlc_ParametersSidelink_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. mac-ParametersSidelink-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mac_ParametersSidelink_r16 = new(MAC_ParametersSidelink_r16)
			if err := ie.Mac_ParametersSidelink_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. fdd-Add-UE-Sidelink-Capabilities-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Fdd_Add_UE_Sidelink_Capabilities_r16 = new(UE_SidelinkCapabilityAddXDD_Mode_r16)
			if err := ie.Fdd_Add_UE_Sidelink_Capabilities_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. tdd-Add-UE-Sidelink-Capabilities-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Tdd_Add_UE_Sidelink_Capabilities_r16 = new(UE_SidelinkCapabilityAddXDD_Mode_r16)
			if err := ie.Tdd_Add_UE_Sidelink_Capabilities_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. supportedBandListSidelink-r16: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(sidelinkParametersNRR16SupportedBandListSidelinkR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SupportedBandListSidelink_r16 = make([]BandSidelink_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SupportedBandListSidelink_r16[i].Decode(d); err != nil {
					return err
				}
			}
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
				{Name: "relayParameters-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.RelayParameters_r17 = new(RelayParameters_r17)
			if err := ie.RelayParameters_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "p0-OLPC-Sidelink-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sidelinkParametersNRR16ExtP0OLPCSidelinkR17Constraints)
			if err != nil {
				return err
			}
			ie.P0_OLPC_Sidelink_r17 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdcp-ParametersSidelink-r18", Optional: true},
				{Name: "sl-PRS-CommonProcCapabilityPerUE-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Pdcp_ParametersSidelink_r18 = new(PDCP_ParametersSidelink_r18)
			if err := ie.Pdcp_ParametersSidelink_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_PRS_CommonProcCapabilityPerUE_r18 = &struct {
				MaxNumOfActiveSL_PRS_Resources_r18 struct {
					Fr1_r18 *int64
					Fr2_r18 *int64
				}
				MaxNumOfSlotswithActiveSL_PRS_Resources_r18 struct {
					Fr1_r18 *int64
					Fr2_r18 *int64
				}
			}{}
			c := ie.Sl_PRS_CommonProcCapabilityPerUE_r18
			{
				c := &c.MaxNumOfActiveSL_PRS_Resources_r18
				sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Seq := dx.NewSequenceDecoder(sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Constraints)
				if err := sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Seq.DecodePreamble(); err != nil {
					return err
				}
				if sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Seq.IsComponentPresent(0) {
					c.Fr1_r18 = new(int64)
					v, err := dx.DecodeEnumerated(sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Fr1R18Constraints)
					if err != nil {
						return err
					}
					(*c.Fr1_r18) = v
				}
				if sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Seq.IsComponentPresent(1) {
					c.Fr2_r18 = new(int64)
					v, err := dx.DecodeEnumerated(sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfActiveSLPRSResourcesR18Fr2R18Constraints)
					if err != nil {
						return err
					}
					(*c.Fr2_r18) = v
				}
			}
			{
				c := &c.MaxNumOfSlotswithActiveSL_PRS_Resources_r18
				sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Seq := dx.NewSequenceDecoder(sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Constraints)
				if err := sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Seq.DecodePreamble(); err != nil {
					return err
				}
				if sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Seq.IsComponentPresent(0) {
					c.Fr1_r18 = new(int64)
					v, err := dx.DecodeEnumerated(sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Fr1R18Constraints)
					if err != nil {
						return err
					}
					(*c.Fr1_r18) = v
				}
				if sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Seq.IsComponentPresent(1) {
					c.Fr2_r18 = new(int64)
					v, err := dx.DecodeEnumerated(sidelinkParametersNRR16ExtSlPRSCommonProcCapabilityPerUER18MaxNumOfSlotswithActiveSLPRSResourcesR18Fr2R18Constraints)
					if err != nil {
						return err
					}
					(*c.Fr2_r18) = v
				}
			}
		}
	}

	return nil
}
