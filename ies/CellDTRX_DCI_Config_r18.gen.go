// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CellDTRX-DCI-config-r18 (line 11762).

var cellDTRXDCIConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellDTRX-RNTI-r18"},
		{Name: "sizeDCI-2-9-r18"},
	},
}

var cellDTRXDCIConfigR18SizeDCI29R18Constraints = per.Constrained(1, common.MaxDCI_2_9_Size_r18)

type CellDTRX_DCI_Config_r18 struct {
	CellDTRX_RNTI_r18 RNTI_Value
	SizeDCI_2_9_r18   int64
}

func (ie *CellDTRX_DCI_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellDTRXDCIConfigR18Constraints)
	_ = seq

	// 1. cellDTRX-RNTI-r18: ref
	{
		if err := ie.CellDTRX_RNTI_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. sizeDCI-2-9-r18: integer
	{
		if err := e.EncodeInteger(ie.SizeDCI_2_9_r18, cellDTRXDCIConfigR18SizeDCI29R18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CellDTRX_DCI_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellDTRXDCIConfigR18Constraints)
	_ = seq

	// 1. cellDTRX-RNTI-r18: ref
	{
		if err := ie.CellDTRX_RNTI_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. sizeDCI-2-9-r18: integer
	{
		v1, err := d.DecodeInteger(cellDTRXDCIConfigR18SizeDCI29R18Constraints)
		if err != nil {
			return err
		}
		ie.SizeDCI_2_9_r18 = v1
	}

	return nil
}
