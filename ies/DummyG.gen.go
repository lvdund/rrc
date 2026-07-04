// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyG (line 22440).

var dummyGConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberSSB-CSI-RS-ResourceOneTx"},
		{Name: "maxNumberSSB-CSI-RS-ResourceTwoTx"},
		{Name: "supportedCSI-RS-Density"},
	},
}

const (
	DummyG_MaxNumberSSB_CSI_RS_ResourceOneTx_N8  = 0
	DummyG_MaxNumberSSB_CSI_RS_ResourceOneTx_N16 = 1
	DummyG_MaxNumberSSB_CSI_RS_ResourceOneTx_N32 = 2
	DummyG_MaxNumberSSB_CSI_RS_ResourceOneTx_N64 = 3
)

var dummyGMaxNumberSSBCSIRSResourceOneTxConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyG_MaxNumberSSB_CSI_RS_ResourceOneTx_N8, DummyG_MaxNumberSSB_CSI_RS_ResourceOneTx_N16, DummyG_MaxNumberSSB_CSI_RS_ResourceOneTx_N32, DummyG_MaxNumberSSB_CSI_RS_ResourceOneTx_N64},
}

const (
	DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N0  = 0
	DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N4  = 1
	DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N8  = 2
	DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N16 = 3
	DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N32 = 4
	DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N64 = 5
)

var dummyGMaxNumberSSBCSIRSResourceTwoTxConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N0, DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N4, DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N8, DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N16, DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N32, DummyG_MaxNumberSSB_CSI_RS_ResourceTwoTx_N64},
}

const (
	DummyG_SupportedCSI_RS_Density_One         = 0
	DummyG_SupportedCSI_RS_Density_Three       = 1
	DummyG_SupportedCSI_RS_Density_OneAndThree = 2
)

var dummyGSupportedCSIRSDensityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyG_SupportedCSI_RS_Density_One, DummyG_SupportedCSI_RS_Density_Three, DummyG_SupportedCSI_RS_Density_OneAndThree},
}

type DummyG struct {
	MaxNumberSSB_CSI_RS_ResourceOneTx int64
	MaxNumberSSB_CSI_RS_ResourceTwoTx int64
	SupportedCSI_RS_Density           int64
}

func (ie *DummyG) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyGConstraints)
	_ = seq

	// 1. maxNumberSSB-CSI-RS-ResourceOneTx: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSSB_CSI_RS_ResourceOneTx, dummyGMaxNumberSSBCSIRSResourceOneTxConstraints); err != nil {
			return err
		}
	}

	// 2. maxNumberSSB-CSI-RS-ResourceTwoTx: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSSB_CSI_RS_ResourceTwoTx, dummyGMaxNumberSSBCSIRSResourceTwoTxConstraints); err != nil {
			return err
		}
	}

	// 3. supportedCSI-RS-Density: enumerated
	{
		if err := e.EncodeEnumerated(ie.SupportedCSI_RS_Density, dummyGSupportedCSIRSDensityConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DummyG) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyGConstraints)
	_ = seq

	// 1. maxNumberSSB-CSI-RS-ResourceOneTx: enumerated
	{
		v0, err := d.DecodeEnumerated(dummyGMaxNumberSSBCSIRSResourceOneTxConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSSB_CSI_RS_ResourceOneTx = v0
	}

	// 2. maxNumberSSB-CSI-RS-ResourceTwoTx: enumerated
	{
		v1, err := d.DecodeEnumerated(dummyGMaxNumberSSBCSIRSResourceTwoTxConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSSB_CSI_RS_ResourceTwoTx = v1
	}

	// 3. supportedCSI-RS-Density: enumerated
	{
		v2, err := d.DecodeEnumerated(dummyGSupportedCSIRSDensityConstraints)
		if err != nil {
			return err
		}
		ie.SupportedCSI_RS_Density = v2
	}

	return nil
}
