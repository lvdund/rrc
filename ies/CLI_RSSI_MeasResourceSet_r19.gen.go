// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CLI-RSSI-MeasResourceSet-r19 (line 6006).

var cLIRSSIMeasResourceSetR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "aperiodicTriggeringOffset-r19", Optional: true},
		{Name: "cli-RSSI-MeasResourceSetId-r19", Optional: true},
		{Name: "cli-RSSI-MeasResourceIdList-r19", Optional: true},
	},
}

var cLIRSSIMeasResourceSetR19AperiodicTriggeringOffsetR19Constraints = per.Constrained(0, 31)

var cLIRSSIMeasResourceSetR19CliRSSIMeasResourceIdListR19Constraints = per.SizeRange(1, common.MaxNrofCLI_RSSI_MeasResourcesPerSet_r19)

type CLI_RSSI_MeasResourceSet_r19 struct {
	AperiodicTriggeringOffset_r19   *int64
	Cli_RSSI_MeasResourceSetId_r19  *CLI_RSSI_MeasResourceSetId_r19
	Cli_RSSI_MeasResourceIdList_r19 []CLI_RSSI_MeasResourceId_r19
}

func (ie *CLI_RSSI_MeasResourceSet_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cLIRSSIMeasResourceSetR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AperiodicTriggeringOffset_r19 != nil, ie.Cli_RSSI_MeasResourceSetId_r19 != nil, ie.Cli_RSSI_MeasResourceIdList_r19 != nil}); err != nil {
		return err
	}

	// 3. aperiodicTriggeringOffset-r19: integer
	{
		if ie.AperiodicTriggeringOffset_r19 != nil {
			if err := e.EncodeInteger(*ie.AperiodicTriggeringOffset_r19, cLIRSSIMeasResourceSetR19AperiodicTriggeringOffsetR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. cli-RSSI-MeasResourceSetId-r19: ref
	{
		if ie.Cli_RSSI_MeasResourceSetId_r19 != nil {
			if err := ie.Cli_RSSI_MeasResourceSetId_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. cli-RSSI-MeasResourceIdList-r19: sequence-of
	{
		if ie.Cli_RSSI_MeasResourceIdList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(cLIRSSIMeasResourceSetR19CliRSSIMeasResourceIdListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Cli_RSSI_MeasResourceIdList_r19))); err != nil {
				return err
			}
			for i := range ie.Cli_RSSI_MeasResourceIdList_r19 {
				if err := ie.Cli_RSSI_MeasResourceIdList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CLI_RSSI_MeasResourceSet_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cLIRSSIMeasResourceSetR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. aperiodicTriggeringOffset-r19: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(cLIRSSIMeasResourceSetR19AperiodicTriggeringOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.AperiodicTriggeringOffset_r19 = &v
		}
	}

	// 4. cli-RSSI-MeasResourceSetId-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Cli_RSSI_MeasResourceSetId_r19 = new(CLI_RSSI_MeasResourceSetId_r19)
			if err := ie.Cli_RSSI_MeasResourceSetId_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. cli-RSSI-MeasResourceIdList-r19: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(cLIRSSIMeasResourceSetR19CliRSSIMeasResourceIdListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Cli_RSSI_MeasResourceIdList_r19 = make([]CLI_RSSI_MeasResourceId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Cli_RSSI_MeasResourceIdList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
