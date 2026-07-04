// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkCancellation-r16 (line 16274).

var uplinkCancellationR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ci-RNTI-r16"},
		{Name: "dci-PayloadSizeForCI-r16"},
		{Name: "ci-ConfigurationPerServingCell-r16"},
	},
}

var uplinkCancellationR16DciPayloadSizeForCIR16Constraints = per.Constrained(0, common.MaxCI_DCI_PayloadSize_r16)

var uplinkCancellationR16CiConfigurationPerServingCellR16Constraints = per.SizeRange(1, common.MaxNrofServingCells)

type UplinkCancellation_r16 struct {
	Ci_RNTI_r16                        RNTI_Value
	Dci_PayloadSizeForCI_r16           int64
	Ci_ConfigurationPerServingCell_r16 []CI_ConfigurationPerServingCell_r16
}

func (ie *UplinkCancellation_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkCancellationR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. ci-RNTI-r16: ref
	{
		if err := ie.Ci_RNTI_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. dci-PayloadSizeForCI-r16: integer
	{
		if err := e.EncodeInteger(ie.Dci_PayloadSizeForCI_r16, uplinkCancellationR16DciPayloadSizeForCIR16Constraints); err != nil {
			return err
		}
	}

	// 4. ci-ConfigurationPerServingCell-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(uplinkCancellationR16CiConfigurationPerServingCellR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Ci_ConfigurationPerServingCell_r16))); err != nil {
			return err
		}
		for i := range ie.Ci_ConfigurationPerServingCell_r16 {
			if err := ie.Ci_ConfigurationPerServingCell_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UplinkCancellation_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkCancellationR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. ci-RNTI-r16: ref
	{
		if err := ie.Ci_RNTI_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. dci-PayloadSizeForCI-r16: integer
	{
		v1, err := d.DecodeInteger(uplinkCancellationR16DciPayloadSizeForCIR16Constraints)
		if err != nil {
			return err
		}
		ie.Dci_PayloadSizeForCI_r16 = v1
	}

	// 4. ci-ConfigurationPerServingCell-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(uplinkCancellationR16CiConfigurationPerServingCellR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Ci_ConfigurationPerServingCell_r16 = make([]CI_ConfigurationPerServingCell_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Ci_ConfigurationPerServingCell_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
