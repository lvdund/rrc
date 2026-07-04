// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DownlinkPreemption (line 8089).

var downlinkPreemptionConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "int-RNTI"},
		{Name: "timeFrequencySet"},
		{Name: "dci-PayloadSize"},
		{Name: "int-ConfigurationPerServingCell"},
	},
}

const (
	DownlinkPreemption_TimeFrequencySet_Set0 = 0
	DownlinkPreemption_TimeFrequencySet_Set1 = 1
)

var downlinkPreemptionTimeFrequencySetConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DownlinkPreemption_TimeFrequencySet_Set0, DownlinkPreemption_TimeFrequencySet_Set1},
}

var downlinkPreemptionDciPayloadSizeConstraints = per.Constrained(0, common.MaxINT_DCI_PayloadSize)

var downlinkPreemptionIntConfigurationPerServingCellConstraints = per.SizeRange(1, common.MaxNrofServingCells)

type DownlinkPreemption struct {
	Int_RNTI                        RNTI_Value
	TimeFrequencySet                int64
	Dci_PayloadSize                 int64
	Int_ConfigurationPerServingCell []INT_ConfigurationPerServingCell
}

func (ie *DownlinkPreemption) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(downlinkPreemptionConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. int-RNTI: ref
	{
		if err := ie.Int_RNTI.Encode(e); err != nil {
			return err
		}
	}

	// 3. timeFrequencySet: enumerated
	{
		if err := e.EncodeEnumerated(ie.TimeFrequencySet, downlinkPreemptionTimeFrequencySetConstraints); err != nil {
			return err
		}
	}

	// 4. dci-PayloadSize: integer
	{
		if err := e.EncodeInteger(ie.Dci_PayloadSize, downlinkPreemptionDciPayloadSizeConstraints); err != nil {
			return err
		}
	}

	// 5. int-ConfigurationPerServingCell: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(downlinkPreemptionIntConfigurationPerServingCellConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Int_ConfigurationPerServingCell))); err != nil {
			return err
		}
		for i := range ie.Int_ConfigurationPerServingCell {
			if err := ie.Int_ConfigurationPerServingCell[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DownlinkPreemption) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(downlinkPreemptionConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. int-RNTI: ref
	{
		if err := ie.Int_RNTI.Decode(d); err != nil {
			return err
		}
	}

	// 3. timeFrequencySet: enumerated
	{
		v1, err := d.DecodeEnumerated(downlinkPreemptionTimeFrequencySetConstraints)
		if err != nil {
			return err
		}
		ie.TimeFrequencySet = v1
	}

	// 4. dci-PayloadSize: integer
	{
		v2, err := d.DecodeInteger(downlinkPreemptionDciPayloadSizeConstraints)
		if err != nil {
			return err
		}
		ie.Dci_PayloadSize = v2
	}

	// 5. int-ConfigurationPerServingCell: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(downlinkPreemptionIntConfigurationPerServingCellConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Int_ConfigurationPerServingCell = make([]INT_ConfigurationPerServingCell, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Int_ConfigurationPerServingCell[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
