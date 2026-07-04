// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CG-SDT-TA-ValidationConfig-r17 (line 1400).

var cGSDTTAValidationConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cg-SDT-RSRP-ChangeThreshold-r17"},
	},
}

const (
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB2    = 0
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB4    = 1
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB6    = 2
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB8    = 3
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB10   = 4
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB14   = 5
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB18   = 6
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB22   = 7
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB26   = 8
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB30   = 9
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB34   = 10
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_Spare5 = 11
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_Spare4 = 12
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_Spare3 = 13
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_Spare2 = 14
	CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_Spare1 = 15
)

var cGSDTTAValidationConfigR17CgSDTRSRPChangeThresholdR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB2, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB4, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB6, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB8, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB10, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB14, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB18, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB22, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB26, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB30, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_DB34, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_Spare5, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_Spare4, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_Spare3, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_Spare2, CG_SDT_TA_ValidationConfig_r17_Cg_SDT_RSRP_ChangeThreshold_r17_Spare1},
}

type CG_SDT_TA_ValidationConfig_r17 struct {
	Cg_SDT_RSRP_ChangeThreshold_r17 int64
}

func (ie *CG_SDT_TA_ValidationConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cGSDTTAValidationConfigR17Constraints)
	_ = seq

	// 1. cg-SDT-RSRP-ChangeThreshold-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Cg_SDT_RSRP_ChangeThreshold_r17, cGSDTTAValidationConfigR17CgSDTRSRPChangeThresholdR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CG_SDT_TA_ValidationConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cGSDTTAValidationConfigR17Constraints)
	_ = seq

	// 1. cg-SDT-RSRP-ChangeThreshold-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(cGSDTTAValidationConfigR17CgSDTRSRPChangeThresholdR17Constraints)
		if err != nil {
			return err
		}
		ie.Cg_SDT_RSRP_ChangeThreshold_r17 = v0
	}

	return nil
}
