// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CFRA-SSB-Resource (line 12973).

var cFRASSBResourceConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb"},
		{Name: "ra-PreambleIndex"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var cFRASSBResourceRaPreambleIndexConstraints = per.Constrained(0, 63)

var cFRASSBResourceMsgAPUSCHResourceIndexR16Constraints = per.Constrained(0, 3071)

type CFRA_SSB_Resource struct {
	Ssb                           SSB_Index
	Ra_PreambleIndex              int64
	MsgA_PUSCH_Resource_Index_r16 *int64
}

func (ie *CFRA_SSB_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cFRASSBResourceConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MsgA_PUSCH_Resource_Index_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. ssb: ref
	{
		if err := ie.Ssb.Encode(e); err != nil {
			return err
		}
	}

	// 3. ra-PreambleIndex: integer
	{
		if err := e.EncodeInteger(ie.Ra_PreambleIndex, cFRASSBResourceRaPreambleIndexConstraints); err != nil {
			return err
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
					{Name: "msgA-PUSCH-Resource-Index-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MsgA_PUSCH_Resource_Index_r16 != nil}); err != nil {
				return err
			}

			if ie.MsgA_PUSCH_Resource_Index_r16 != nil {
				if err := ex.EncodeInteger(*ie.MsgA_PUSCH_Resource_Index_r16, cFRASSBResourceMsgAPUSCHResourceIndexR16Constraints); err != nil {
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

func (ie *CFRA_SSB_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cFRASSBResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. ssb: ref
	{
		if err := ie.Ssb.Decode(d); err != nil {
			return err
		}
	}

	// 3. ra-PreambleIndex: integer
	{
		v1, err := d.DecodeInteger(cFRASSBResourceRaPreambleIndexConstraints)
		if err != nil {
			return err
		}
		ie.Ra_PreambleIndex = v1
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
				{Name: "msgA-PUSCH-Resource-Index-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(cFRASSBResourceMsgAPUSCHResourceIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_PUSCH_Resource_Index_r16 = &v
		}
	}

	return nil
}
