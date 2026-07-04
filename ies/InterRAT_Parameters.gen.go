// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: InterRAT-Parameters (line 20892).

var interRATParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eutra", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

type InterRAT_Parameters struct {
	Eutra        *EUTRA_Parameters
	Utra_FDD_r16 *UTRA_FDD_Parameters_r16
}

func (ie *InterRAT_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interRATParametersConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Utra_FDD_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Eutra != nil}); err != nil {
		return err
	}

	// 3. eutra: ref
	{
		if ie.Eutra != nil {
			if err := ie.Eutra.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "utra-FDD-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Utra_FDD_r16 != nil}); err != nil {
				return err
			}

			if ie.Utra_FDD_r16 != nil {
				if err := ie.Utra_FDD_r16.Encode(ex); err != nil {
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

func (ie *InterRAT_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interRATParametersConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. eutra: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Eutra = new(EUTRA_Parameters)
			if err := ie.Eutra.Decode(d); err != nil {
				return err
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
				{Name: "utra-FDD-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Utra_FDD_r16 = new(UTRA_FDD_Parameters_r16)
			if err := ie.Utra_FDD_r16.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
