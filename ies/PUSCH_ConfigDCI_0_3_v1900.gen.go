// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-ConfigDCI-0-3-v1900 (line 12540).

var pUSCHConfigDCI03V1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "harq-ProcessNumberSizeDCI-0-3-Ext-r19", Optional: true},
	},
}

var pUSCHConfigDCI03V1900HarqProcessNumberSizeDCI03ExtR19Constraints = per.Constrained(0, 5)

type PUSCH_ConfigDCI_0_3_v1900 struct {
	Harq_ProcessNumberSizeDCI_0_3_Ext_r19 *int64
}

func (ie *PUSCH_ConfigDCI_0_3_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHConfigDCI03V1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Harq_ProcessNumberSizeDCI_0_3_Ext_r19 != nil}); err != nil {
		return err
	}

	// 2. harq-ProcessNumberSizeDCI-0-3-Ext-r19: integer
	{
		if ie.Harq_ProcessNumberSizeDCI_0_3_Ext_r19 != nil {
			if err := e.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_0_3_Ext_r19, pUSCHConfigDCI03V1900HarqProcessNumberSizeDCI03ExtR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUSCH_ConfigDCI_0_3_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHConfigDCI03V1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. harq-ProcessNumberSizeDCI-0-3-Ext-r19: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pUSCHConfigDCI03V1900HarqProcessNumberSizeDCI03ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_0_3_Ext_r19 = &v
		}
	}

	return nil
}
