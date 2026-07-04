// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RelayParameters-r17 (line 25286).

var relayParametersR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "relayUE-Operation-L2-r17", Optional: true},
		{Name: "remoteUE-Operation-L2-r17", Optional: true},
		{Name: "remoteUE-PathSwitchToIdleInactiveRelay-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	RelayParameters_r17_RelayUE_Operation_L2_r17_Supported = 0
)

var relayParametersR17RelayUEOperationL2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_RelayUE_Operation_L2_r17_Supported},
}

const (
	RelayParameters_r17_RemoteUE_Operation_L2_r17_Supported = 0
)

var relayParametersR17RemoteUEOperationL2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_RemoteUE_Operation_L2_r17_Supported},
}

const (
	RelayParameters_r17_RemoteUE_PathSwitchToIdleInactiveRelay_r17_Supported = 0
)

var relayParametersR17RemoteUEPathSwitchToIdleInactiveRelayR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_RemoteUE_PathSwitchToIdleInactiveRelay_r17_Supported},
}

const (
	RelayParameters_r17_Ext_RelayUE_U2U_OperationL2_r18_Supported = 0
)

var relayParametersR17ExtRelayUEU2UOperationL2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_RelayUE_U2U_OperationL2_r18_Supported},
}

const (
	RelayParameters_r17_Ext_RemoteUE_U2U_OperationL2_r18_Supported = 0
)

var relayParametersR17ExtRemoteUEU2UOperationL2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_RemoteUE_U2U_OperationL2_r18_Supported},
}

const (
	RelayParameters_r17_Ext_RemoteUE_U2N_PathSwitchOperationL2_r18_Supported = 0
)

var relayParametersR17ExtRemoteUEU2NPathSwitchOperationL2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_RemoteUE_U2N_PathSwitchOperationL2_r18_Supported},
}

const (
	RelayParameters_r17_Ext_MultipathRemoteUE_PC5L2_r18_Supported = 0
)

var relayParametersR17ExtMultipathRemoteUEPC5L2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_MultipathRemoteUE_PC5L2_r18_Supported},
}

const (
	RelayParameters_r17_Ext_MultipathRelayUE_N3C_r18_Supported = 0
)

var relayParametersR17ExtMultipathRelayUEN3CR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_MultipathRelayUE_N3C_r18_Supported},
}

const (
	RelayParameters_r17_Ext_MultipathRemoteUE_N3C_r18_Supported = 0
)

var relayParametersR17ExtMultipathRemoteUEN3CR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_MultipathRemoteUE_N3C_r18_Supported},
}

const (
	RelayParameters_r17_Ext_RemoteUE_IndirectPathAddChangeToIdleInactiveRelay_r18_Supported = 0
)

var relayParametersR17ExtRemoteUEIndirectPathAddChangeToIdleInactiveRelayR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_RemoteUE_IndirectPathAddChangeToIdleInactiveRelay_r18_Supported},
}

const (
	RelayParameters_r17_Ext_Pdcp_DuplicationMoreThanOneUuRLC_r18_Supported = 0
)

var relayParametersR17ExtPdcpDuplicationMoreThanOneUuRLCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_Pdcp_DuplicationMoreThanOneUuRLC_r18_Supported},
}

const (
	RelayParameters_r17_Ext_Pdcp_CADuplicationDirectpath_DRB_r18_Supported = 0
)

var relayParametersR17ExtPdcpCADuplicationDirectpathDRBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_Pdcp_CADuplicationDirectpath_DRB_r18_Supported},
}

const (
	RelayParameters_r17_Ext_Pdcp_CADuplicationDirectpath_SRB_r18_Supported = 0
)

var relayParametersR17ExtPdcpCADuplicationDirectpathSRBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_Pdcp_CADuplicationDirectpath_SRB_r18_Supported},
}

const (
	RelayParameters_r17_Ext_Pdcp_DuplicationMP_SplitDRB_r18_Supported = 0
)

var relayParametersR17ExtPdcpDuplicationMPSplitDRBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_Pdcp_DuplicationMP_SplitDRB_r18_Supported},
}

const (
	RelayParameters_r17_Ext_Pdcp_DuplicationMP_SplitSRB_r18_Supported = 0
)

var relayParametersR17ExtPdcpDuplicationMPSplitSRBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_Pdcp_DuplicationMP_SplitSRB_r18_Supported},
}

const (
	RelayParameters_r17_Ext_DirectpathRLF_RecoveryViaSRB1_r18_Supported = 0
)

var relayParametersR17ExtDirectpathRLFRecoveryViaSRB1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_DirectpathRLF_RecoveryViaSRB1_r18_Supported},
}

const (
	RelayParameters_r17_Ext_SplitDRB_WithUL_BothDirectIndirect_r18_Supported = 0
)

var relayParametersR17ExtSplitDRBWithULBothDirectIndirectR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_SplitDRB_WithUL_BothDirectIndirect_r18_Supported},
}

const (
	RelayParameters_r17_Ext_MultipathRemoteUE_ExtN3C_r19_Supported = 0
)

var relayParametersR17ExtMultipathRemoteUEExtN3CR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_MultipathRemoteUE_ExtN3C_r19_Supported},
}

const (
	RelayParameters_r17_Ext_RelayUE_MH_OperationL2_r19_Supported = 0
)

var relayParametersR17ExtRelayUEMHOperationL2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_RelayUE_MH_OperationL2_r19_Supported},
}

const (
	RelayParameters_r17_Ext_RemoteUE_MH_OperationL2_r19_Supported = 0
)

var relayParametersR17ExtRemoteUEMHOperationL2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RelayParameters_r17_Ext_RemoteUE_MH_OperationL2_r19_Supported},
}

type RelayParameters_r17 struct {
	RelayUE_Operation_L2_r17                              *int64
	RemoteUE_Operation_L2_r17                             *int64
	RemoteUE_PathSwitchToIdleInactiveRelay_r17            *int64
	RelayUE_U2U_OperationL2_r18                           *int64
	RemoteUE_U2U_OperationL2_r18                          *int64
	RemoteUE_U2N_PathSwitchOperationL2_r18                *int64
	MultipathRemoteUE_PC5L2_r18                           *int64
	MultipathRelayUE_N3C_r18                              *int64
	MultipathRemoteUE_N3C_r18                             *int64
	RemoteUE_IndirectPathAddChangeToIdleInactiveRelay_r18 *int64
	Pdcp_DuplicationMoreThanOneUuRLC_r18                  *int64
	Pdcp_CADuplicationDirectpath_DRB_r18                  *int64
	Pdcp_CADuplicationDirectpath_SRB_r18                  *int64
	Pdcp_DuplicationMP_SplitDRB_r18                       *int64
	Pdcp_DuplicationMP_SplitSRB_r18                       *int64
	DirectpathRLF_RecoveryViaSRB1_r18                     *int64
	SplitDRB_WithUL_BothDirectIndirect_r18                *int64
	MultipathRemoteUE_ExtN3C_r19                          *int64
	RelayUE_MH_OperationL2_r19                            *int64
	RemoteUE_MH_OperationL2_r19                           *int64
}

func (ie *RelayParameters_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(relayParametersR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.RelayUE_U2U_OperationL2_r18 != nil || ie.RemoteUE_U2U_OperationL2_r18 != nil || ie.RemoteUE_U2N_PathSwitchOperationL2_r18 != nil || ie.MultipathRemoteUE_PC5L2_r18 != nil || ie.MultipathRelayUE_N3C_r18 != nil || ie.MultipathRemoteUE_N3C_r18 != nil || ie.RemoteUE_IndirectPathAddChangeToIdleInactiveRelay_r18 != nil || ie.Pdcp_DuplicationMoreThanOneUuRLC_r18 != nil || ie.Pdcp_CADuplicationDirectpath_DRB_r18 != nil || ie.Pdcp_CADuplicationDirectpath_SRB_r18 != nil || ie.Pdcp_DuplicationMP_SplitDRB_r18 != nil || ie.Pdcp_DuplicationMP_SplitSRB_r18 != nil || ie.DirectpathRLF_RecoveryViaSRB1_r18 != nil || ie.SplitDRB_WithUL_BothDirectIndirect_r18 != nil
	hasExtGroup1 := ie.MultipathRemoteUE_ExtN3C_r19 != nil || ie.RelayUE_MH_OperationL2_r19 != nil || ie.RemoteUE_MH_OperationL2_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RelayUE_Operation_L2_r17 != nil, ie.RemoteUE_Operation_L2_r17 != nil, ie.RemoteUE_PathSwitchToIdleInactiveRelay_r17 != nil}); err != nil {
		return err
	}

	// 3. relayUE-Operation-L2-r17: enumerated
	{
		if ie.RelayUE_Operation_L2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.RelayUE_Operation_L2_r17, relayParametersR17RelayUEOperationL2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. remoteUE-Operation-L2-r17: enumerated
	{
		if ie.RemoteUE_Operation_L2_r17 != nil {
			if err := e.EncodeEnumerated(*ie.RemoteUE_Operation_L2_r17, relayParametersR17RemoteUEOperationL2R17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. remoteUE-PathSwitchToIdleInactiveRelay-r17: enumerated
	{
		if ie.RemoteUE_PathSwitchToIdleInactiveRelay_r17 != nil {
			if err := e.EncodeEnumerated(*ie.RemoteUE_PathSwitchToIdleInactiveRelay_r17, relayParametersR17RemoteUEPathSwitchToIdleInactiveRelayR17Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "relayUE-U2U-OperationL2-r18", Optional: true},
					{Name: "remoteUE-U2U-OperationL2-r18", Optional: true},
					{Name: "remoteUE-U2N-PathSwitchOperationL2-r18", Optional: true},
					{Name: "multipathRemoteUE-PC5L2-r18", Optional: true},
					{Name: "multipathRelayUE-N3C-r18", Optional: true},
					{Name: "multipathRemoteUE-N3C-r18", Optional: true},
					{Name: "remoteUE-IndirectPathAddChangeToIdleInactiveRelay-r18", Optional: true},
					{Name: "pdcp-DuplicationMoreThanOneUuRLC-r18", Optional: true},
					{Name: "pdcp-CADuplicationDirectpath-DRB-r18", Optional: true},
					{Name: "pdcp-CADuplicationDirectpath-SRB-r18", Optional: true},
					{Name: "pdcp-DuplicationMP-SplitDRB-r18", Optional: true},
					{Name: "pdcp-DuplicationMP-SplitSRB-r18", Optional: true},
					{Name: "directpathRLF-RecoveryViaSRB1-r18", Optional: true},
					{Name: "splitDRB-WithUL-BothDirectIndirect-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RelayUE_U2U_OperationL2_r18 != nil, ie.RemoteUE_U2U_OperationL2_r18 != nil, ie.RemoteUE_U2N_PathSwitchOperationL2_r18 != nil, ie.MultipathRemoteUE_PC5L2_r18 != nil, ie.MultipathRelayUE_N3C_r18 != nil, ie.MultipathRemoteUE_N3C_r18 != nil, ie.RemoteUE_IndirectPathAddChangeToIdleInactiveRelay_r18 != nil, ie.Pdcp_DuplicationMoreThanOneUuRLC_r18 != nil, ie.Pdcp_CADuplicationDirectpath_DRB_r18 != nil, ie.Pdcp_CADuplicationDirectpath_SRB_r18 != nil, ie.Pdcp_DuplicationMP_SplitDRB_r18 != nil, ie.Pdcp_DuplicationMP_SplitSRB_r18 != nil, ie.DirectpathRLF_RecoveryViaSRB1_r18 != nil, ie.SplitDRB_WithUL_BothDirectIndirect_r18 != nil}); err != nil {
				return err
			}

			if ie.RelayUE_U2U_OperationL2_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.RelayUE_U2U_OperationL2_r18, relayParametersR17ExtRelayUEU2UOperationL2R18Constraints); err != nil {
					return err
				}
			}

			if ie.RemoteUE_U2U_OperationL2_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.RemoteUE_U2U_OperationL2_r18, relayParametersR17ExtRemoteUEU2UOperationL2R18Constraints); err != nil {
					return err
				}
			}

			if ie.RemoteUE_U2N_PathSwitchOperationL2_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.RemoteUE_U2N_PathSwitchOperationL2_r18, relayParametersR17ExtRemoteUEU2NPathSwitchOperationL2R18Constraints); err != nil {
					return err
				}
			}

			if ie.MultipathRemoteUE_PC5L2_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MultipathRemoteUE_PC5L2_r18, relayParametersR17ExtMultipathRemoteUEPC5L2R18Constraints); err != nil {
					return err
				}
			}

			if ie.MultipathRelayUE_N3C_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MultipathRelayUE_N3C_r18, relayParametersR17ExtMultipathRelayUEN3CR18Constraints); err != nil {
					return err
				}
			}

			if ie.MultipathRemoteUE_N3C_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MultipathRemoteUE_N3C_r18, relayParametersR17ExtMultipathRemoteUEN3CR18Constraints); err != nil {
					return err
				}
			}

			if ie.RemoteUE_IndirectPathAddChangeToIdleInactiveRelay_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.RemoteUE_IndirectPathAddChangeToIdleInactiveRelay_r18, relayParametersR17ExtRemoteUEIndirectPathAddChangeToIdleInactiveRelayR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcp_DuplicationMoreThanOneUuRLC_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcp_DuplicationMoreThanOneUuRLC_r18, relayParametersR17ExtPdcpDuplicationMoreThanOneUuRLCR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcp_CADuplicationDirectpath_DRB_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcp_CADuplicationDirectpath_DRB_r18, relayParametersR17ExtPdcpCADuplicationDirectpathDRBR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcp_CADuplicationDirectpath_SRB_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcp_CADuplicationDirectpath_SRB_r18, relayParametersR17ExtPdcpCADuplicationDirectpathSRBR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcp_DuplicationMP_SplitDRB_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcp_DuplicationMP_SplitDRB_r18, relayParametersR17ExtPdcpDuplicationMPSplitDRBR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcp_DuplicationMP_SplitSRB_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcp_DuplicationMP_SplitSRB_r18, relayParametersR17ExtPdcpDuplicationMPSplitSRBR18Constraints); err != nil {
					return err
				}
			}

			if ie.DirectpathRLF_RecoveryViaSRB1_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.DirectpathRLF_RecoveryViaSRB1_r18, relayParametersR17ExtDirectpathRLFRecoveryViaSRB1R18Constraints); err != nil {
					return err
				}
			}

			if ie.SplitDRB_WithUL_BothDirectIndirect_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.SplitDRB_WithUL_BothDirectIndirect_r18, relayParametersR17ExtSplitDRBWithULBothDirectIndirectR18Constraints); err != nil {
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
					{Name: "multipathRemoteUE-ExtN3C-r19", Optional: true},
					{Name: "relayUE-MH-OperationL2-r19", Optional: true},
					{Name: "remoteUE-MH-OperationL2-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MultipathRemoteUE_ExtN3C_r19 != nil, ie.RelayUE_MH_OperationL2_r19 != nil, ie.RemoteUE_MH_OperationL2_r19 != nil}); err != nil {
				return err
			}

			if ie.MultipathRemoteUE_ExtN3C_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.MultipathRemoteUE_ExtN3C_r19, relayParametersR17ExtMultipathRemoteUEExtN3CR19Constraints); err != nil {
					return err
				}
			}

			if ie.RelayUE_MH_OperationL2_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.RelayUE_MH_OperationL2_r19, relayParametersR17ExtRelayUEMHOperationL2R19Constraints); err != nil {
					return err
				}
			}

			if ie.RemoteUE_MH_OperationL2_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.RemoteUE_MH_OperationL2_r19, relayParametersR17ExtRemoteUEMHOperationL2R19Constraints); err != nil {
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

func (ie *RelayParameters_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(relayParametersR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. relayUE-Operation-L2-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(relayParametersR17RelayUEOperationL2R17Constraints)
			if err != nil {
				return err
			}
			ie.RelayUE_Operation_L2_r17 = &idx
		}
	}

	// 4. remoteUE-Operation-L2-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(relayParametersR17RemoteUEOperationL2R17Constraints)
			if err != nil {
				return err
			}
			ie.RemoteUE_Operation_L2_r17 = &idx
		}
	}

	// 5. remoteUE-PathSwitchToIdleInactiveRelay-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(relayParametersR17RemoteUEPathSwitchToIdleInactiveRelayR17Constraints)
			if err != nil {
				return err
			}
			ie.RemoteUE_PathSwitchToIdleInactiveRelay_r17 = &idx
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
				{Name: "relayUE-U2U-OperationL2-r18", Optional: true},
				{Name: "remoteUE-U2U-OperationL2-r18", Optional: true},
				{Name: "remoteUE-U2N-PathSwitchOperationL2-r18", Optional: true},
				{Name: "multipathRemoteUE-PC5L2-r18", Optional: true},
				{Name: "multipathRelayUE-N3C-r18", Optional: true},
				{Name: "multipathRemoteUE-N3C-r18", Optional: true},
				{Name: "remoteUE-IndirectPathAddChangeToIdleInactiveRelay-r18", Optional: true},
				{Name: "pdcp-DuplicationMoreThanOneUuRLC-r18", Optional: true},
				{Name: "pdcp-CADuplicationDirectpath-DRB-r18", Optional: true},
				{Name: "pdcp-CADuplicationDirectpath-SRB-r18", Optional: true},
				{Name: "pdcp-DuplicationMP-SplitDRB-r18", Optional: true},
				{Name: "pdcp-DuplicationMP-SplitSRB-r18", Optional: true},
				{Name: "directpathRLF-RecoveryViaSRB1-r18", Optional: true},
				{Name: "splitDRB-WithUL-BothDirectIndirect-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtRelayUEU2UOperationL2R18Constraints)
			if err != nil {
				return err
			}
			ie.RelayUE_U2U_OperationL2_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtRemoteUEU2UOperationL2R18Constraints)
			if err != nil {
				return err
			}
			ie.RemoteUE_U2U_OperationL2_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtRemoteUEU2NPathSwitchOperationL2R18Constraints)
			if err != nil {
				return err
			}
			ie.RemoteUE_U2N_PathSwitchOperationL2_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtMultipathRemoteUEPC5L2R18Constraints)
			if err != nil {
				return err
			}
			ie.MultipathRemoteUE_PC5L2_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtMultipathRelayUEN3CR18Constraints)
			if err != nil {
				return err
			}
			ie.MultipathRelayUE_N3C_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtMultipathRemoteUEN3CR18Constraints)
			if err != nil {
				return err
			}
			ie.MultipathRemoteUE_N3C_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtRemoteUEIndirectPathAddChangeToIdleInactiveRelayR18Constraints)
			if err != nil {
				return err
			}
			ie.RemoteUE_IndirectPathAddChangeToIdleInactiveRelay_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtPdcpDuplicationMoreThanOneUuRLCR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationMoreThanOneUuRLC_r18 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtPdcpCADuplicationDirectpathDRBR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcp_CADuplicationDirectpath_DRB_r18 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtPdcpCADuplicationDirectpathSRBR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcp_CADuplicationDirectpath_SRB_r18 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtPdcpDuplicationMPSplitDRBR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationMP_SplitDRB_r18 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtPdcpDuplicationMPSplitSRBR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcp_DuplicationMP_SplitSRB_r18 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtDirectpathRLFRecoveryViaSRB1R18Constraints)
			if err != nil {
				return err
			}
			ie.DirectpathRLF_RecoveryViaSRB1_r18 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtSplitDRBWithULBothDirectIndirectR18Constraints)
			if err != nil {
				return err
			}
			ie.SplitDRB_WithUL_BothDirectIndirect_r18 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "multipathRemoteUE-ExtN3C-r19", Optional: true},
				{Name: "relayUE-MH-OperationL2-r19", Optional: true},
				{Name: "remoteUE-MH-OperationL2-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtMultipathRemoteUEExtN3CR19Constraints)
			if err != nil {
				return err
			}
			ie.MultipathRemoteUE_ExtN3C_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtRelayUEMHOperationL2R19Constraints)
			if err != nil {
				return err
			}
			ie.RelayUE_MH_OperationL2_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(relayParametersR17ExtRemoteUEMHOperationL2R19Constraints)
			if err != nil {
				return err
			}
			ie.RemoteUE_MH_OperationL2_r19 = &v
		}
	}

	return nil
}
