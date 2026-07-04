// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CLI-RSSI-MeasResource-r19 (line 5987).

var cLIRSSIMeasResourceR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cli-RSSI-MeasResourceId-r19", Optional: true},
		{Name: "startSymbol-r19", Optional: true},
		{Name: "nrofSymbols-r19", Optional: true},
		{Name: "startPRB-r19", Optional: true},
		{Name: "nrofPRBs-r19", Optional: true},
		{Name: "cli-RSSI-PeriodicityAndOffset-r19", Optional: true},
		{Name: "qcl-InfoPeriodic-CLI-RSSI-MeasResource-r19", Optional: true},
	},
}

var cLIRSSIMeasResourceR19StartSymbolR19Constraints = per.Constrained(0, 13)

var cLIRSSIMeasResourceR19NrofSymbolsR19Constraints = per.Constrained(1, 14)

var cLIRSSIMeasResourceR19StartPRBR19Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

var cLIRSSIMeasResourceR19NrofPRBsR19Constraints = per.Constrained(1, common.MaxNrofPhysicalResourceBlocks)

type CLI_RSSI_MeasResource_r19 struct {
	Cli_RSSI_MeasResourceId_r19                *CLI_RSSI_MeasResourceId_r19
	StartSymbol_r19                            *int64
	NrofSymbols_r19                            *int64
	StartPRB_r19                               *int64
	NrofPRBs_r19                               *int64
	Cli_RSSI_PeriodicityAndOffset_r19          *CSI_ReportPeriodicityAndOffset
	Qcl_InfoPeriodic_CLI_RSSI_MeasResource_r19 *TCI_StateId
}

func (ie *CLI_RSSI_MeasResource_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cLIRSSIMeasResourceR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cli_RSSI_MeasResourceId_r19 != nil, ie.StartSymbol_r19 != nil, ie.NrofSymbols_r19 != nil, ie.StartPRB_r19 != nil, ie.NrofPRBs_r19 != nil, ie.Cli_RSSI_PeriodicityAndOffset_r19 != nil, ie.Qcl_InfoPeriodic_CLI_RSSI_MeasResource_r19 != nil}); err != nil {
		return err
	}

	// 3. cli-RSSI-MeasResourceId-r19: ref
	{
		if ie.Cli_RSSI_MeasResourceId_r19 != nil {
			if err := ie.Cli_RSSI_MeasResourceId_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. startSymbol-r19: integer
	{
		if ie.StartSymbol_r19 != nil {
			if err := e.EncodeInteger(*ie.StartSymbol_r19, cLIRSSIMeasResourceR19StartSymbolR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. nrofSymbols-r19: integer
	{
		if ie.NrofSymbols_r19 != nil {
			if err := e.EncodeInteger(*ie.NrofSymbols_r19, cLIRSSIMeasResourceR19NrofSymbolsR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. startPRB-r19: integer
	{
		if ie.StartPRB_r19 != nil {
			if err := e.EncodeInteger(*ie.StartPRB_r19, cLIRSSIMeasResourceR19StartPRBR19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. nrofPRBs-r19: integer
	{
		if ie.NrofPRBs_r19 != nil {
			if err := e.EncodeInteger(*ie.NrofPRBs_r19, cLIRSSIMeasResourceR19NrofPRBsR19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. cli-RSSI-PeriodicityAndOffset-r19: ref
	{
		if ie.Cli_RSSI_PeriodicityAndOffset_r19 != nil {
			if err := ie.Cli_RSSI_PeriodicityAndOffset_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. qcl-InfoPeriodic-CLI-RSSI-MeasResource-r19: ref
	{
		if ie.Qcl_InfoPeriodic_CLI_RSSI_MeasResource_r19 != nil {
			if err := ie.Qcl_InfoPeriodic_CLI_RSSI_MeasResource_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CLI_RSSI_MeasResource_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cLIRSSIMeasResourceR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. cli-RSSI-MeasResourceId-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Cli_RSSI_MeasResourceId_r19 = new(CLI_RSSI_MeasResourceId_r19)
			if err := ie.Cli_RSSI_MeasResourceId_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. startSymbol-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(cLIRSSIMeasResourceR19StartSymbolR19Constraints)
			if err != nil {
				return err
			}
			ie.StartSymbol_r19 = &v
		}
	}

	// 5. nrofSymbols-r19: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(cLIRSSIMeasResourceR19NrofSymbolsR19Constraints)
			if err != nil {
				return err
			}
			ie.NrofSymbols_r19 = &v
		}
	}

	// 6. startPRB-r19: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(cLIRSSIMeasResourceR19StartPRBR19Constraints)
			if err != nil {
				return err
			}
			ie.StartPRB_r19 = &v
		}
	}

	// 7. nrofPRBs-r19: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(cLIRSSIMeasResourceR19NrofPRBsR19Constraints)
			if err != nil {
				return err
			}
			ie.NrofPRBs_r19 = &v
		}
	}

	// 8. cli-RSSI-PeriodicityAndOffset-r19: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Cli_RSSI_PeriodicityAndOffset_r19 = new(CSI_ReportPeriodicityAndOffset)
			if err := ie.Cli_RSSI_PeriodicityAndOffset_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. qcl-InfoPeriodic-CLI-RSSI-MeasResource-r19: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Qcl_InfoPeriodic_CLI_RSSI_MeasResource_r19 = new(TCI_StateId)
			if err := ie.Qcl_InfoPeriodic_CLI_RSSI_MeasResource_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
