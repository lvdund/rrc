// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-TPC-CommandConfig (line 15831).

var sRSTPCCommandConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "startingBitOfFormat2-3", Optional: true},
		{Name: "fieldTypeFormat2-3", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var sRSTPCCommandConfigStartingBitOfFormat23Constraints = per.Constrained(1, 31)

var sRSTPCCommandConfigFieldTypeFormat23Constraints = per.Constrained(0, 1)

var sRSTPCCommandConfigStartingBitOfFormat23SULConstraints = per.Constrained(1, 31)

var sRSTPCCommandConfigStartingBitOfFormat23R19Constraints = per.Constrained(1, 45)

type SRS_TPC_CommandConfig struct {
	StartingBitOfFormat2_3     *int64
	FieldTypeFormat2_3         *int64
	StartingBitOfFormat2_3SUL  *int64
	StartingBitOfFormat2_3_r19 *int64
}

func (ie *SRS_TPC_CommandConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSTPCCommandConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.StartingBitOfFormat2_3SUL != nil
	hasExtGroup1 := ie.StartingBitOfFormat2_3_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.StartingBitOfFormat2_3 != nil, ie.FieldTypeFormat2_3 != nil}); err != nil {
		return err
	}

	// 3. startingBitOfFormat2-3: integer
	{
		if ie.StartingBitOfFormat2_3 != nil {
			if err := e.EncodeInteger(*ie.StartingBitOfFormat2_3, sRSTPCCommandConfigStartingBitOfFormat23Constraints); err != nil {
				return err
			}
		}
	}

	// 4. fieldTypeFormat2-3: integer
	{
		if ie.FieldTypeFormat2_3 != nil {
			if err := e.EncodeInteger(*ie.FieldTypeFormat2_3, sRSTPCCommandConfigFieldTypeFormat23Constraints); err != nil {
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
					{Name: "startingBitOfFormat2-3SUL", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.StartingBitOfFormat2_3SUL != nil}); err != nil {
				return err
			}

			if ie.StartingBitOfFormat2_3SUL != nil {
				if err := ex.EncodeInteger(*ie.StartingBitOfFormat2_3SUL, sRSTPCCommandConfigStartingBitOfFormat23SULConstraints); err != nil {
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
					{Name: "startingBitOfFormat2-3-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.StartingBitOfFormat2_3_r19 != nil}); err != nil {
				return err
			}

			if ie.StartingBitOfFormat2_3_r19 != nil {
				if err := ex.EncodeInteger(*ie.StartingBitOfFormat2_3_r19, sRSTPCCommandConfigStartingBitOfFormat23R19Constraints); err != nil {
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

func (ie *SRS_TPC_CommandConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSTPCCommandConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. startingBitOfFormat2-3: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sRSTPCCommandConfigStartingBitOfFormat23Constraints)
			if err != nil {
				return err
			}
			ie.StartingBitOfFormat2_3 = &v
		}
	}

	// 4. fieldTypeFormat2-3: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sRSTPCCommandConfigFieldTypeFormat23Constraints)
			if err != nil {
				return err
			}
			ie.FieldTypeFormat2_3 = &v
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
				{Name: "startingBitOfFormat2-3SUL", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sRSTPCCommandConfigStartingBitOfFormat23SULConstraints)
			if err != nil {
				return err
			}
			ie.StartingBitOfFormat2_3SUL = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "startingBitOfFormat2-3-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sRSTPCCommandConfigStartingBitOfFormat23R19Constraints)
			if err != nil {
				return err
			}
			ie.StartingBitOfFormat2_3_r19 = &v
		}
	}

	return nil
}
