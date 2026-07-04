// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-ConfigDCI-1-3-v1860 (line 11458).

var pDSCHConfigDCI13V1860Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "enabledDefaultBeamForMultiCellScheduling-r18"},
	},
}

const (
	PDSCH_ConfigDCI_1_3_v1860_EnabledDefaultBeamForMultiCellScheduling_r18_Enabled = 0
)

var pDSCHConfigDCI13V1860EnabledDefaultBeamForMultiCellSchedulingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ConfigDCI_1_3_v1860_EnabledDefaultBeamForMultiCellScheduling_r18_Enabled},
}

type PDSCH_ConfigDCI_1_3_v1860 struct {
	EnabledDefaultBeamForMultiCellScheduling_r18 int64
}

func (ie *PDSCH_ConfigDCI_1_3_v1860) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHConfigDCI13V1860Constraints)
	_ = seq

	// 1. enabledDefaultBeamForMultiCellScheduling-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.EnabledDefaultBeamForMultiCellScheduling_r18, pDSCHConfigDCI13V1860EnabledDefaultBeamForMultiCellSchedulingR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDSCH_ConfigDCI_1_3_v1860) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHConfigDCI13V1860Constraints)
	_ = seq

	// 1. enabledDefaultBeamForMultiCellScheduling-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(pDSCHConfigDCI13V1860EnabledDefaultBeamForMultiCellSchedulingR18Constraints)
		if err != nil {
			return err
		}
		ie.EnabledDefaultBeamForMultiCellScheduling_r18 = v0
	}

	return nil
}
