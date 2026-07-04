// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-SpatialRelationInfoExt-r16 (line 12300).

var pUCCHSpatialRelationInfoExtR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-SpatialRelationInfoId-v1610", Optional: true},
		{Name: "pucch-PathlossReferenceRS-Id-v1610", Optional: true},
	},
}

type PUCCH_SpatialRelationInfoExt_r16 struct {
	Pucch_SpatialRelationInfoId_v1610  *PUCCH_SpatialRelationInfoId_v1610
	Pucch_PathlossReferenceRS_Id_v1610 *PUCCH_PathlossReferenceRS_Id_v1610
}

func (ie *PUCCH_SpatialRelationInfoExt_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHSpatialRelationInfoExtR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pucch_SpatialRelationInfoId_v1610 != nil, ie.Pucch_PathlossReferenceRS_Id_v1610 != nil}); err != nil {
		return err
	}

	// 3. pucch-SpatialRelationInfoId-v1610: ref
	{
		if ie.Pucch_SpatialRelationInfoId_v1610 != nil {
			if err := ie.Pucch_SpatialRelationInfoId_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. pucch-PathlossReferenceRS-Id-v1610: ref
	{
		if ie.Pucch_PathlossReferenceRS_Id_v1610 != nil {
			if err := ie.Pucch_PathlossReferenceRS_Id_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUCCH_SpatialRelationInfoExt_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHSpatialRelationInfoExtR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pucch-SpatialRelationInfoId-v1610: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Pucch_SpatialRelationInfoId_v1610 = new(PUCCH_SpatialRelationInfoId_v1610)
			if err := ie.Pucch_SpatialRelationInfoId_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. pucch-PathlossReferenceRS-Id-v1610: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Pucch_PathlossReferenceRS_Id_v1610 = new(PUCCH_PathlossReferenceRS_Id_v1610)
			if err := ie.Pucch_PathlossReferenceRS_Id_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
