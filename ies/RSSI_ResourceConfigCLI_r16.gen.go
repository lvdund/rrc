// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RSSI-ResourceConfigCLI-r16 (line 9359).

var rSSIResourceConfigCLIR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rssi-ResourceId-r16"},
		{Name: "rssi-SCS-r16"},
		{Name: "startPRB-r16"},
		{Name: "nrofPRBs-r16"},
		{Name: "startPosition-r16"},
		{Name: "nrofSymbols-r16"},
		{Name: "rssi-PeriodicityAndOffset-r16"},
		{Name: "refServCellIndex-r16", Optional: true},
	},
}

var rSSIResourceConfigCLIR16StartPRBR16Constraints = per.Constrained(0, 2169)

var rSSIResourceConfigCLIR16NrofPRBsR16Constraints = per.Constrained(4, common.MaxNrofPhysicalResourceBlocksPlus1)

var rSSIResourceConfigCLIR16StartPositionR16Constraints = per.Constrained(0, 13)

var rSSIResourceConfigCLIR16NrofSymbolsR16Constraints = per.Constrained(1, 14)

type RSSI_ResourceConfigCLI_r16 struct {
	Rssi_ResourceId_r16           RSSI_ResourceId_r16
	Rssi_SCS_r16                  SubcarrierSpacing
	StartPRB_r16                  int64
	NrofPRBs_r16                  int64
	StartPosition_r16             int64
	NrofSymbols_r16               int64
	Rssi_PeriodicityAndOffset_r16 RSSI_PeriodicityAndOffset_r16
	RefServCellIndex_r16          *ServCellIndex
}

func (ie *RSSI_ResourceConfigCLI_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rSSIResourceConfigCLIR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RefServCellIndex_r16 != nil}); err != nil {
		return err
	}

	// 3. rssi-ResourceId-r16: ref
	{
		if err := ie.Rssi_ResourceId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. rssi-SCS-r16: ref
	{
		if err := ie.Rssi_SCS_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. startPRB-r16: integer
	{
		if err := e.EncodeInteger(ie.StartPRB_r16, rSSIResourceConfigCLIR16StartPRBR16Constraints); err != nil {
			return err
		}
	}

	// 6. nrofPRBs-r16: integer
	{
		if err := e.EncodeInteger(ie.NrofPRBs_r16, rSSIResourceConfigCLIR16NrofPRBsR16Constraints); err != nil {
			return err
		}
	}

	// 7. startPosition-r16: integer
	{
		if err := e.EncodeInteger(ie.StartPosition_r16, rSSIResourceConfigCLIR16StartPositionR16Constraints); err != nil {
			return err
		}
	}

	// 8. nrofSymbols-r16: integer
	{
		if err := e.EncodeInteger(ie.NrofSymbols_r16, rSSIResourceConfigCLIR16NrofSymbolsR16Constraints); err != nil {
			return err
		}
	}

	// 9. rssi-PeriodicityAndOffset-r16: ref
	{
		if err := ie.Rssi_PeriodicityAndOffset_r16.Encode(e); err != nil {
			return err
		}
	}

	// 10. refServCellIndex-r16: ref
	{
		if ie.RefServCellIndex_r16 != nil {
			if err := ie.RefServCellIndex_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RSSI_ResourceConfigCLI_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rSSIResourceConfigCLIR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rssi-ResourceId-r16: ref
	{
		if err := ie.Rssi_ResourceId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. rssi-SCS-r16: ref
	{
		if err := ie.Rssi_SCS_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. startPRB-r16: integer
	{
		v2, err := d.DecodeInteger(rSSIResourceConfigCLIR16StartPRBR16Constraints)
		if err != nil {
			return err
		}
		ie.StartPRB_r16 = v2
	}

	// 6. nrofPRBs-r16: integer
	{
		v3, err := d.DecodeInteger(rSSIResourceConfigCLIR16NrofPRBsR16Constraints)
		if err != nil {
			return err
		}
		ie.NrofPRBs_r16 = v3
	}

	// 7. startPosition-r16: integer
	{
		v4, err := d.DecodeInteger(rSSIResourceConfigCLIR16StartPositionR16Constraints)
		if err != nil {
			return err
		}
		ie.StartPosition_r16 = v4
	}

	// 8. nrofSymbols-r16: integer
	{
		v5, err := d.DecodeInteger(rSSIResourceConfigCLIR16NrofSymbolsR16Constraints)
		if err != nil {
			return err
		}
		ie.NrofSymbols_r16 = v5
	}

	// 9. rssi-PeriodicityAndOffset-r16: ref
	{
		if err := ie.Rssi_PeriodicityAndOffset_r16.Decode(d); err != nil {
			return err
		}
	}

	// 10. refServCellIndex-r16: ref
	{
		if seq.IsComponentPresent(7) {
			ie.RefServCellIndex_r16 = new(ServCellIndex)
			if err := ie.RefServCellIndex_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
