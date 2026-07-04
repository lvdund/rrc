// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MIMOParam-r17 (line 14805).

var mIMOParamR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalPCI-ToAddModList-r17", Optional: true},
		{Name: "additionalPCI-ToReleaseList-r17", Optional: true},
		{Name: "unifiedTCI-StateType-r17", Optional: true},
		{Name: "uplink-PowerControlToAddModList-r17", Optional: true},
		{Name: "uplink-PowerControlToReleaseList-r17", Optional: true},
		{Name: "sfnSchemePDCCH-r17", Optional: true},
		{Name: "sfnSchemePDSCH-r17", Optional: true},
	},
}

var mIMOParamR17AdditionalPCIToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofAdditionalPCI_r17)

var mIMOParamR17AdditionalPCIToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofAdditionalPCI_r17)

const (
	MIMOParam_r17_UnifiedTCI_StateType_r17_Separate = 0
	MIMOParam_r17_UnifiedTCI_StateType_r17_Joint    = 1
)

var mIMOParamR17UnifiedTCIStateTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMOParam_r17_UnifiedTCI_StateType_r17_Separate, MIMOParam_r17_UnifiedTCI_StateType_r17_Joint},
}

var mIMOParamR17UplinkPowerControlToAddModListR17Constraints = per.SizeRange(1, common.MaxUL_TCI_r17)

var mIMOParamR17UplinkPowerControlToReleaseListR17Constraints = per.SizeRange(1, common.MaxUL_TCI_r17)

const (
	MIMOParam_r17_SfnSchemePDCCH_r17_SfnSchemeA = 0
	MIMOParam_r17_SfnSchemePDCCH_r17_SfnSchemeB = 1
)

var mIMOParamR17SfnSchemePDCCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMOParam_r17_SfnSchemePDCCH_r17_SfnSchemeA, MIMOParam_r17_SfnSchemePDCCH_r17_SfnSchemeB},
}

const (
	MIMOParam_r17_SfnSchemePDSCH_r17_SfnSchemeA = 0
	MIMOParam_r17_SfnSchemePDSCH_r17_SfnSchemeB = 1
)

var mIMOParamR17SfnSchemePDSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMOParam_r17_SfnSchemePDSCH_r17_SfnSchemeA, MIMOParam_r17_SfnSchemePDSCH_r17_SfnSchemeB},
}

type MIMOParam_r17 struct {
	AdditionalPCI_ToAddModList_r17       []SSB_MTC_AdditionalPCI_r17
	AdditionalPCI_ToReleaseList_r17      []AdditionalPCIIndex_r17
	UnifiedTCI_StateType_r17             *int64
	Uplink_PowerControlToAddModList_r17  []Uplink_PowerControl_r17
	Uplink_PowerControlToReleaseList_r17 []Uplink_PowerControlId_r17
	SfnSchemePDCCH_r17                   *int64
	SfnSchemePDSCH_r17                   *int64
}

func (ie *MIMOParam_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mIMOParamR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AdditionalPCI_ToAddModList_r17 != nil, ie.AdditionalPCI_ToReleaseList_r17 != nil, ie.UnifiedTCI_StateType_r17 != nil, ie.Uplink_PowerControlToAddModList_r17 != nil, ie.Uplink_PowerControlToReleaseList_r17 != nil, ie.SfnSchemePDCCH_r17 != nil, ie.SfnSchemePDSCH_r17 != nil}); err != nil {
		return err
	}

	// 2. additionalPCI-ToAddModList-r17: sequence-of
	{
		if ie.AdditionalPCI_ToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(mIMOParamR17AdditionalPCIToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AdditionalPCI_ToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.AdditionalPCI_ToAddModList_r17 {
				if err := ie.AdditionalPCI_ToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. additionalPCI-ToReleaseList-r17: sequence-of
	{
		if ie.AdditionalPCI_ToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(mIMOParamR17AdditionalPCIToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AdditionalPCI_ToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.AdditionalPCI_ToReleaseList_r17 {
				if err := ie.AdditionalPCI_ToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. unifiedTCI-StateType-r17: enumerated
	{
		if ie.UnifiedTCI_StateType_r17 != nil {
			if err := e.EncodeEnumerated(*ie.UnifiedTCI_StateType_r17, mIMOParamR17UnifiedTCIStateTypeR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. uplink-PowerControlToAddModList-r17: sequence-of
	{
		if ie.Uplink_PowerControlToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(mIMOParamR17UplinkPowerControlToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Uplink_PowerControlToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.Uplink_PowerControlToAddModList_r17 {
				if err := ie.Uplink_PowerControlToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. uplink-PowerControlToReleaseList-r17: sequence-of
	{
		if ie.Uplink_PowerControlToReleaseList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(mIMOParamR17UplinkPowerControlToReleaseListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Uplink_PowerControlToReleaseList_r17))); err != nil {
				return err
			}
			for i := range ie.Uplink_PowerControlToReleaseList_r17 {
				if err := ie.Uplink_PowerControlToReleaseList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. sfnSchemePDCCH-r17: enumerated
	{
		if ie.SfnSchemePDCCH_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SfnSchemePDCCH_r17, mIMOParamR17SfnSchemePDCCHR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sfnSchemePDSCH-r17: enumerated
	{
		if ie.SfnSchemePDSCH_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SfnSchemePDSCH_r17, mIMOParamR17SfnSchemePDSCHR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MIMOParam_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mIMOParamR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. additionalPCI-ToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(mIMOParamR17AdditionalPCIToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AdditionalPCI_ToAddModList_r17 = make([]SSB_MTC_AdditionalPCI_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AdditionalPCI_ToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. additionalPCI-ToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(mIMOParamR17AdditionalPCIToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AdditionalPCI_ToReleaseList_r17 = make([]AdditionalPCIIndex_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AdditionalPCI_ToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. unifiedTCI-StateType-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mIMOParamR17UnifiedTCIStateTypeR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedTCI_StateType_r17 = &idx
		}
	}

	// 5. uplink-PowerControlToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(mIMOParamR17UplinkPowerControlToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Uplink_PowerControlToAddModList_r17 = make([]Uplink_PowerControl_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Uplink_PowerControlToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. uplink-PowerControlToReleaseList-r17: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(mIMOParamR17UplinkPowerControlToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Uplink_PowerControlToReleaseList_r17 = make([]Uplink_PowerControlId_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Uplink_PowerControlToReleaseList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. sfnSchemePDCCH-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(mIMOParamR17SfnSchemePDCCHR17Constraints)
			if err != nil {
				return err
			}
			ie.SfnSchemePDCCH_r17 = &idx
		}
	}

	// 8. sfnSchemePDSCH-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(mIMOParamR17SfnSchemePDSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.SfnSchemePDSCH_r17 = &idx
		}
	}

	return nil
}
