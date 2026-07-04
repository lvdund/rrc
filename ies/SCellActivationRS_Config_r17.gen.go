// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SCellActivationRS-Config-r17 (line 14229).

var sCellActivationRSConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "scellActivationRS-Id-r17"},
		{Name: "resourceSet-r17"},
		{Name: "gapBetweenBursts-r17", Optional: true},
		{Name: "qcl-Info-r17"},
	},
}

var sCellActivationRSConfigR17GapBetweenBurstsR17Constraints = per.Constrained(2, 31)

type SCellActivationRS_Config_r17 struct {
	ScellActivationRS_Id_r17 SCellActivationRS_ConfigId_r17
	ResourceSet_r17          NZP_CSI_RS_ResourceSetId
	GapBetweenBursts_r17     *int64
	Qcl_Info_r17             TCI_StateId
}

func (ie *SCellActivationRS_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCellActivationRSConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.GapBetweenBursts_r17 != nil}); err != nil {
		return err
	}

	// 3. scellActivationRS-Id-r17: ref
	{
		if err := ie.ScellActivationRS_Id_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. resourceSet-r17: ref
	{
		if err := ie.ResourceSet_r17.Encode(e); err != nil {
			return err
		}
	}

	// 5. gapBetweenBursts-r17: integer
	{
		if ie.GapBetweenBursts_r17 != nil {
			if err := e.EncodeInteger(*ie.GapBetweenBursts_r17, sCellActivationRSConfigR17GapBetweenBurstsR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. qcl-Info-r17: ref
	{
		if err := ie.Qcl_Info_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SCellActivationRS_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCellActivationRSConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. scellActivationRS-Id-r17: ref
	{
		if err := ie.ScellActivationRS_Id_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. resourceSet-r17: ref
	{
		if err := ie.ResourceSet_r17.Decode(d); err != nil {
			return err
		}
	}

	// 5. gapBetweenBursts-r17: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sCellActivationRSConfigR17GapBetweenBurstsR17Constraints)
			if err != nil {
				return err
			}
			ie.GapBetweenBursts_r17 = &v
		}
	}

	// 6. qcl-Info-r17: ref
	{
		if err := ie.Qcl_Info_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
