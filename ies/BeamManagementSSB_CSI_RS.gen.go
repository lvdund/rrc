// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BeamManagementSSB-CSI-RS (line 22446).

var beamManagementSSBCSIRSConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberSSB-CSI-RS-ResourceOneTx"},
		{Name: "maxNumberCSI-RS-Resource"},
		{Name: "maxNumberCSI-RS-ResourceTwoTx"},
		{Name: "supportedCSI-RS-Density", Optional: true},
		{Name: "maxNumberAperiodicCSI-RS-Resource"},
	},
}

const (
	BeamManagementSSB_CSI_RS_MaxNumberSSB_CSI_RS_ResourceOneTx_N0  = 0
	BeamManagementSSB_CSI_RS_MaxNumberSSB_CSI_RS_ResourceOneTx_N8  = 1
	BeamManagementSSB_CSI_RS_MaxNumberSSB_CSI_RS_ResourceOneTx_N16 = 2
	BeamManagementSSB_CSI_RS_MaxNumberSSB_CSI_RS_ResourceOneTx_N32 = 3
	BeamManagementSSB_CSI_RS_MaxNumberSSB_CSI_RS_ResourceOneTx_N64 = 4
)

var beamManagementSSBCSIRSMaxNumberSSBCSIRSResourceOneTxConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamManagementSSB_CSI_RS_MaxNumberSSB_CSI_RS_ResourceOneTx_N0, BeamManagementSSB_CSI_RS_MaxNumberSSB_CSI_RS_ResourceOneTx_N8, BeamManagementSSB_CSI_RS_MaxNumberSSB_CSI_RS_ResourceOneTx_N16, BeamManagementSSB_CSI_RS_MaxNumberSSB_CSI_RS_ResourceOneTx_N32, BeamManagementSSB_CSI_RS_MaxNumberSSB_CSI_RS_ResourceOneTx_N64},
}

const (
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N0  = 0
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N4  = 1
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N8  = 2
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N16 = 3
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N32 = 4
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N64 = 5
)

var beamManagementSSBCSIRSMaxNumberCSIRSResourceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N0, BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N4, BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N8, BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N16, BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N32, BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_Resource_N64},
}

const (
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N0  = 0
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N4  = 1
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N8  = 2
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N16 = 3
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N32 = 4
	BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N64 = 5
)

var beamManagementSSBCSIRSMaxNumberCSIRSResourceTwoTxConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N0, BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N4, BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N8, BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N16, BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N32, BeamManagementSSB_CSI_RS_MaxNumberCSI_RS_ResourceTwoTx_N64},
}

const (
	BeamManagementSSB_CSI_RS_SupportedCSI_RS_Density_One         = 0
	BeamManagementSSB_CSI_RS_SupportedCSI_RS_Density_Three       = 1
	BeamManagementSSB_CSI_RS_SupportedCSI_RS_Density_OneAndThree = 2
)

var beamManagementSSBCSIRSSupportedCSIRSDensityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamManagementSSB_CSI_RS_SupportedCSI_RS_Density_One, BeamManagementSSB_CSI_RS_SupportedCSI_RS_Density_Three, BeamManagementSSB_CSI_RS_SupportedCSI_RS_Density_OneAndThree},
}

const (
	BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N0  = 0
	BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N1  = 1
	BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N4  = 2
	BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N8  = 3
	BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N16 = 4
	BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N32 = 5
	BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N64 = 6
)

var beamManagementSSBCSIRSMaxNumberAperiodicCSIRSResourceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N0, BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N1, BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N4, BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N8, BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N16, BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N32, BeamManagementSSB_CSI_RS_MaxNumberAperiodicCSI_RS_Resource_N64},
}

type BeamManagementSSB_CSI_RS struct {
	MaxNumberSSB_CSI_RS_ResourceOneTx int64
	MaxNumberCSI_RS_Resource          int64
	MaxNumberCSI_RS_ResourceTwoTx     int64
	SupportedCSI_RS_Density           *int64
	MaxNumberAperiodicCSI_RS_Resource int64
}

func (ie *BeamManagementSSB_CSI_RS) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(beamManagementSSBCSIRSConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedCSI_RS_Density != nil}); err != nil {
		return err
	}

	// 2. maxNumberSSB-CSI-RS-ResourceOneTx: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberSSB_CSI_RS_ResourceOneTx, beamManagementSSBCSIRSMaxNumberSSBCSIRSResourceOneTxConstraints); err != nil {
			return err
		}
	}

	// 3. maxNumberCSI-RS-Resource: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberCSI_RS_Resource, beamManagementSSBCSIRSMaxNumberCSIRSResourceConstraints); err != nil {
			return err
		}
	}

	// 4. maxNumberCSI-RS-ResourceTwoTx: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberCSI_RS_ResourceTwoTx, beamManagementSSBCSIRSMaxNumberCSIRSResourceTwoTxConstraints); err != nil {
			return err
		}
	}

	// 5. supportedCSI-RS-Density: enumerated
	{
		if ie.SupportedCSI_RS_Density != nil {
			if err := e.EncodeEnumerated(*ie.SupportedCSI_RS_Density, beamManagementSSBCSIRSSupportedCSIRSDensityConstraints); err != nil {
				return err
			}
		}
	}

	// 6. maxNumberAperiodicCSI-RS-Resource: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberAperiodicCSI_RS_Resource, beamManagementSSBCSIRSMaxNumberAperiodicCSIRSResourceConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *BeamManagementSSB_CSI_RS) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(beamManagementSSBCSIRSConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxNumberSSB-CSI-RS-ResourceOneTx: enumerated
	{
		v0, err := d.DecodeEnumerated(beamManagementSSBCSIRSMaxNumberSSBCSIRSResourceOneTxConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberSSB_CSI_RS_ResourceOneTx = v0
	}

	// 3. maxNumberCSI-RS-Resource: enumerated
	{
		v1, err := d.DecodeEnumerated(beamManagementSSBCSIRSMaxNumberCSIRSResourceConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberCSI_RS_Resource = v1
	}

	// 4. maxNumberCSI-RS-ResourceTwoTx: enumerated
	{
		v2, err := d.DecodeEnumerated(beamManagementSSBCSIRSMaxNumberCSIRSResourceTwoTxConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberCSI_RS_ResourceTwoTx = v2
	}

	// 5. supportedCSI-RS-Density: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(beamManagementSSBCSIRSSupportedCSIRSDensityConstraints)
			if err != nil {
				return err
			}
			ie.SupportedCSI_RS_Density = &idx
		}
	}

	// 6. maxNumberAperiodicCSI-RS-Resource: enumerated
	{
		v4, err := d.DecodeEnumerated(beamManagementSSBCSIRSMaxNumberAperiodicCSIRSResourceConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberAperiodicCSI_RS_Resource = v4
	}

	return nil
}
