// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-Grp-CarrierTypes-r16 (line 18336).

var pUCCHGrpCarrierTypesR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr1-NonSharedTDD-r16", Optional: true},
		{Name: "fr1-SharedTDD-r16", Optional: true},
		{Name: "fr1-NonSharedFDD-r16", Optional: true},
		{Name: "fr2-r16", Optional: true},
	},
}

const (
	PUCCH_Grp_CarrierTypes_r16_Fr1_NonSharedTDD_r16_Supported = 0
)

var pUCCHGrpCarrierTypesR16Fr1NonSharedTDDR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Grp_CarrierTypes_r16_Fr1_NonSharedTDD_r16_Supported},
}

const (
	PUCCH_Grp_CarrierTypes_r16_Fr1_SharedTDD_r16_Supported = 0
)

var pUCCHGrpCarrierTypesR16Fr1SharedTDDR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Grp_CarrierTypes_r16_Fr1_SharedTDD_r16_Supported},
}

const (
	PUCCH_Grp_CarrierTypes_r16_Fr1_NonSharedFDD_r16_Supported = 0
)

var pUCCHGrpCarrierTypesR16Fr1NonSharedFDDR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Grp_CarrierTypes_r16_Fr1_NonSharedFDD_r16_Supported},
}

const (
	PUCCH_Grp_CarrierTypes_r16_Fr2_r16_Supported = 0
)

var pUCCHGrpCarrierTypesR16Fr2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Grp_CarrierTypes_r16_Fr2_r16_Supported},
}

type PUCCH_Grp_CarrierTypes_r16 struct {
	Fr1_NonSharedTDD_r16 *int64
	Fr1_SharedTDD_r16    *int64
	Fr1_NonSharedFDD_r16 *int64
	Fr2_r16              *int64
}

func (ie *PUCCH_Grp_CarrierTypes_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHGrpCarrierTypesR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Fr1_NonSharedTDD_r16 != nil, ie.Fr1_SharedTDD_r16 != nil, ie.Fr1_NonSharedFDD_r16 != nil, ie.Fr2_r16 != nil}); err != nil {
		return err
	}

	// 2. fr1-NonSharedTDD-r16: enumerated
	{
		if ie.Fr1_NonSharedTDD_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Fr1_NonSharedTDD_r16, pUCCHGrpCarrierTypesR16Fr1NonSharedTDDR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. fr1-SharedTDD-r16: enumerated
	{
		if ie.Fr1_SharedTDD_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Fr1_SharedTDD_r16, pUCCHGrpCarrierTypesR16Fr1SharedTDDR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. fr1-NonSharedFDD-r16: enumerated
	{
		if ie.Fr1_NonSharedFDD_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Fr1_NonSharedFDD_r16, pUCCHGrpCarrierTypesR16Fr1NonSharedFDDR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. fr2-r16: enumerated
	{
		if ie.Fr2_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Fr2_r16, pUCCHGrpCarrierTypesR16Fr2R16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUCCH_Grp_CarrierTypes_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHGrpCarrierTypesR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. fr1-NonSharedTDD-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pUCCHGrpCarrierTypesR16Fr1NonSharedTDDR16Constraints)
			if err != nil {
				return err
			}
			ie.Fr1_NonSharedTDD_r16 = &idx
		}
	}

	// 3. fr1-SharedTDD-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pUCCHGrpCarrierTypesR16Fr1SharedTDDR16Constraints)
			if err != nil {
				return err
			}
			ie.Fr1_SharedTDD_r16 = &idx
		}
	}

	// 4. fr1-NonSharedFDD-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pUCCHGrpCarrierTypesR16Fr1NonSharedFDDR16Constraints)
			if err != nil {
				return err
			}
			ie.Fr1_NonSharedFDD_r16 = &idx
		}
	}

	// 5. fr2-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(pUCCHGrpCarrierTypesR16Fr2R16Constraints)
			if err != nil {
				return err
			}
			ie.Fr2_r16 = &idx
		}
	}

	return nil
}
