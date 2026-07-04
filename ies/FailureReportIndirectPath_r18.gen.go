// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FailureReportIndirectPath-r18 (line 506).

var failureReportIndirectPathR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "failureTypeIndirectPath-r18", Optional: true},
		{Name: "sl-MeasResultServingRelay-r18", Optional: true},
		{Name: "sl-MeasResultsCandRelay-r18", Optional: true},
		{Name: "n3c-RelayUE-InfoList-r18", Optional: true},
	},
}

const (
	FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_T421_Expiry                  = 0
	FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_Sl_Failure                   = 1
	FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_N3c_Failure                  = 2
	FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_RelayUE_Uu_RLF               = 3
	FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_RelayUE_Uu_RRC_Failure       = 4
	FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_IndirectPathAddChangeFailure = 5
	FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_Sl_PC5_Release               = 6
	FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_Spare1                       = 7
)

var failureReportIndirectPathR18FailureTypeIndirectPathR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_T421_Expiry, FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_Sl_Failure, FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_N3c_Failure, FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_RelayUE_Uu_RLF, FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_RelayUE_Uu_RRC_Failure, FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_IndirectPathAddChangeFailure, FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_Sl_PC5_Release, FailureReportIndirectPath_r18_FailureTypeIndirectPath_r18_Spare1},
}

var failureReportIndirectPathR18SlMeasResultServingRelayR18Constraints = per.SizeConstraints{}

var failureReportIndirectPathR18SlMeasResultsCandRelayR18Constraints = per.SizeConstraints{}

var failureReportIndirectPathR18N3cRelayUEInfoListR18Constraints = per.SizeRange(0, 8)

type FailureReportIndirectPath_r18 struct {
	FailureTypeIndirectPath_r18   *int64
	Sl_MeasResultServingRelay_r18 []byte
	Sl_MeasResultsCandRelay_r18   []byte
	N3c_RelayUE_InfoList_r18      []N3C_RelayUE_Info_r18
}

func (ie *FailureReportIndirectPath_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(failureReportIndirectPathR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FailureTypeIndirectPath_r18 != nil, ie.Sl_MeasResultServingRelay_r18 != nil, ie.Sl_MeasResultsCandRelay_r18 != nil, ie.N3c_RelayUE_InfoList_r18 != nil}); err != nil {
		return err
	}

	// 3. failureTypeIndirectPath-r18: enumerated
	{
		if ie.FailureTypeIndirectPath_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FailureTypeIndirectPath_r18, failureReportIndirectPathR18FailureTypeIndirectPathR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-MeasResultServingRelay-r18: octet-string
	{
		if ie.Sl_MeasResultServingRelay_r18 != nil {
			if err := e.EncodeOctetString(ie.Sl_MeasResultServingRelay_r18, failureReportIndirectPathR18SlMeasResultServingRelayR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-MeasResultsCandRelay-r18: octet-string
	{
		if ie.Sl_MeasResultsCandRelay_r18 != nil {
			if err := e.EncodeOctetString(ie.Sl_MeasResultsCandRelay_r18, failureReportIndirectPathR18SlMeasResultsCandRelayR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. n3c-RelayUE-InfoList-r18: sequence-of
	{
		if ie.N3c_RelayUE_InfoList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(failureReportIndirectPathR18N3cRelayUEInfoListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.N3c_RelayUE_InfoList_r18))); err != nil {
				return err
			}
			for i := range ie.N3c_RelayUE_InfoList_r18 {
				if err := ie.N3c_RelayUE_InfoList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FailureReportIndirectPath_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(failureReportIndirectPathR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. failureTypeIndirectPath-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(failureReportIndirectPathR18FailureTypeIndirectPathR18Constraints)
			if err != nil {
				return err
			}
			ie.FailureTypeIndirectPath_r18 = &idx
		}
	}

	// 4. sl-MeasResultServingRelay-r18: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(failureReportIndirectPathR18SlMeasResultServingRelayR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MeasResultServingRelay_r18 = v
		}
	}

	// 5. sl-MeasResultsCandRelay-r18: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(failureReportIndirectPathR18SlMeasResultsCandRelayR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MeasResultsCandRelay_r18 = v
		}
	}

	// 6. n3c-RelayUE-InfoList-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(failureReportIndirectPathR18N3cRelayUEInfoListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.N3c_RelayUE_InfoList_r18 = make([]N3C_RelayUE_Info_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.N3c_RelayUE_InfoList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
