// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-SRAP-Config-ToAddMod-r19 (line 27466).

var sLSRAPConfigToAddModR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-SRAP-ConfigId-r19"},
		{Name: "sl-SRAP-ConfigRelay-r19"},
	},
}

type SL_SRAP_Config_ToAddMod_r19 struct {
	Sl_SRAP_ConfigId_r19    SL_SRAP_ConfigId_r19
	Sl_SRAP_ConfigRelay_r19 SL_SRAP_Config_r17
}

func (ie *SL_SRAP_Config_ToAddMod_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSRAPConfigToAddModR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-SRAP-ConfigId-r19: ref
	{
		if err := ie.Sl_SRAP_ConfigId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-SRAP-ConfigRelay-r19: ref
	{
		if err := ie.Sl_SRAP_ConfigRelay_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_SRAP_Config_ToAddMod_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSRAPConfigToAddModR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-SRAP-ConfigId-r19: ref
	{
		if err := ie.Sl_SRAP_ConfigId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-SRAP-ConfigRelay-r19: ref
	{
		if err := ie.Sl_SRAP_ConfigRelay_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
