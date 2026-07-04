// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SBFD-RACH-DualConfig-r19 (line 5384).

var sBFDRACHDualConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sbfd-AdditionalRACH-Config-r19", Optional: true},
		{Name: "sbfd-RACH-DualConfig-ValidRO-AcrossSymbolTypes-r19", Optional: true},
	},
}

const (
	SBFD_RACH_DualConfig_r19_Sbfd_RACH_DualConfig_ValidRO_AcrossSymbolTypes_r19_Enabled = 0
)

var sBFDRACHDualConfigR19SbfdRACHDualConfigValidROAcrossSymbolTypesR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SBFD_RACH_DualConfig_r19_Sbfd_RACH_DualConfig_ValidRO_AcrossSymbolTypes_r19_Enabled},
}

type SBFD_RACH_DualConfig_r19 struct {
	Sbfd_AdditionalRACH_Config_r19                     *RACH_ConfigCommon
	Sbfd_RACH_DualConfig_ValidRO_AcrossSymbolTypes_r19 *int64
}

func (ie *SBFD_RACH_DualConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sBFDRACHDualConfigR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sbfd_AdditionalRACH_Config_r19 != nil, ie.Sbfd_RACH_DualConfig_ValidRO_AcrossSymbolTypes_r19 != nil}); err != nil {
		return err
	}

	// 2. sbfd-AdditionalRACH-Config-r19: ref
	{
		if ie.Sbfd_AdditionalRACH_Config_r19 != nil {
			if err := ie.Sbfd_AdditionalRACH_Config_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sbfd-RACH-DualConfig-ValidRO-AcrossSymbolTypes-r19: enumerated
	{
		if ie.Sbfd_RACH_DualConfig_ValidRO_AcrossSymbolTypes_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Sbfd_RACH_DualConfig_ValidRO_AcrossSymbolTypes_r19, sBFDRACHDualConfigR19SbfdRACHDualConfigValidROAcrossSymbolTypesR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SBFD_RACH_DualConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sBFDRACHDualConfigR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sbfd-AdditionalRACH-Config-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sbfd_AdditionalRACH_Config_r19 = new(RACH_ConfigCommon)
			if err := ie.Sbfd_AdditionalRACH_Config_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sbfd-RACH-DualConfig-ValidRO-AcrossSymbolTypes-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sBFDRACHDualConfigR19SbfdRACHDualConfigValidROAcrossSymbolTypesR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_RACH_DualConfig_ValidRO_AcrossSymbolTypes_r19 = &idx
		}
	}

	return nil
}
