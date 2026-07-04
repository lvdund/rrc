// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: QuantityConfig (line 12775).

var quantityConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "quantityConfigNR-List", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var quantityConfigQuantityConfigNRListConstraints = per.SizeRange(1, common.MaxNrofQuantityConfig)

type QuantityConfig struct {
	QuantityConfigNR_List      []QuantityConfigNR
	QuantityConfigEUTRA        *FilterConfig
	QuantityConfigUTRA_FDD_r16 *QuantityConfigUTRA_FDD_r16
	QuantityConfigCLI_r16      *FilterConfigCLI_r16
}

func (ie *QuantityConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(quantityConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.QuantityConfigEUTRA != nil
	hasExtGroup1 := ie.QuantityConfigUTRA_FDD_r16 != nil || ie.QuantityConfigCLI_r16 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.QuantityConfigNR_List != nil}); err != nil {
		return err
	}

	// 3. quantityConfigNR-List: sequence-of
	{
		if ie.QuantityConfigNR_List != nil {
			seqOf := e.NewSequenceOfEncoder(quantityConfigQuantityConfigNRListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.QuantityConfigNR_List))); err != nil {
				return err
			}
			for i := range ie.QuantityConfigNR_List {
				if err := ie.QuantityConfigNR_List[i].Encode(e); err != nil {
					return err
				}
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
					{Name: "quantityConfigEUTRA", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.QuantityConfigEUTRA != nil}); err != nil {
				return err
			}

			if ie.QuantityConfigEUTRA != nil {
				if err := ie.QuantityConfigEUTRA.Encode(ex); err != nil {
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
					{Name: "quantityConfigUTRA-FDD-r16", Optional: true},
					{Name: "quantityConfigCLI-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.QuantityConfigUTRA_FDD_r16 != nil, ie.QuantityConfigCLI_r16 != nil}); err != nil {
				return err
			}

			if ie.QuantityConfigUTRA_FDD_r16 != nil {
				if err := ie.QuantityConfigUTRA_FDD_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.QuantityConfigCLI_r16 != nil {
				if err := ie.QuantityConfigCLI_r16.Encode(ex); err != nil {
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

func (ie *QuantityConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(quantityConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. quantityConfigNR-List: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(quantityConfigQuantityConfigNRListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.QuantityConfigNR_List = make([]QuantityConfigNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.QuantityConfigNR_List[i].Decode(d); err != nil {
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
				{Name: "quantityConfigEUTRA", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.QuantityConfigEUTRA = new(FilterConfig)
			if err := ie.QuantityConfigEUTRA.Decode(dx); err != nil {
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
				{Name: "quantityConfigUTRA-FDD-r16", Optional: true},
				{Name: "quantityConfigCLI-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.QuantityConfigUTRA_FDD_r16 = new(QuantityConfigUTRA_FDD_r16)
			if err := ie.QuantityConfigUTRA_FDD_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.QuantityConfigCLI_r16 = new(FilterConfigCLI_r16)
			if err := ie.QuantityConfigCLI_r16.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
