// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-ConfigDCI-1-3-v1900 (line 11462).

var pDSCHConfigDCI13V1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "harq-ProcessNumberSizeDCI-1-3-Ext-r19", Optional: true},
	},
}

var pDSCHConfigDCI13V1900HarqProcessNumberSizeDCI13ExtR19Constraints = per.Constrained(0, 5)

type PDSCH_ConfigDCI_1_3_v1900 struct {
	Harq_ProcessNumberSizeDCI_1_3_Ext_r19 *int64
}

func (ie *PDSCH_ConfigDCI_1_3_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHConfigDCI13V1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Harq_ProcessNumberSizeDCI_1_3_Ext_r19 != nil}); err != nil {
		return err
	}

	// 2. harq-ProcessNumberSizeDCI-1-3-Ext-r19: integer
	{
		if ie.Harq_ProcessNumberSizeDCI_1_3_Ext_r19 != nil {
			if err := e.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_1_3_Ext_r19, pDSCHConfigDCI13V1900HarqProcessNumberSizeDCI13ExtR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDSCH_ConfigDCI_1_3_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHConfigDCI13V1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. harq-ProcessNumberSizeDCI-1-3-Ext-r19: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pDSCHConfigDCI13V1900HarqProcessNumberSizeDCI13ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_1_3_Ext_r19 = &v
		}
	}

	return nil
}
