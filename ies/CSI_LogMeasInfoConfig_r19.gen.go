// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-LogMeasInfoConfig-r19 (line 3502).

var cSILogMeasInfoConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "refCSI-LoggedMeasurementConfigId-r19"},
		{Name: "csi-LogMeasInfoList-r19"},
	},
}

var cSILogMeasInfoConfigR19CsiLogMeasInfoListR19Constraints = per.SizeRange(1, common.MaxLogCSI_MeasReport_r19)

type CSI_LogMeasInfoConfig_r19 struct {
	RefCSI_LoggedMeasurementConfigId_r19 CSI_LoggedMeasurementConfigId_r19
	Csi_LogMeasInfoList_r19              []CSI_LogMeasInfo_r19
}

func (ie *CSI_LogMeasInfoConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSILogMeasInfoConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. refCSI-LoggedMeasurementConfigId-r19: ref
	{
		if err := ie.RefCSI_LoggedMeasurementConfigId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 3. csi-LogMeasInfoList-r19: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cSILogMeasInfoConfigR19CsiLogMeasInfoListR19Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Csi_LogMeasInfoList_r19))); err != nil {
			return err
		}
		for i := range ie.Csi_LogMeasInfoList_r19 {
			if err := ie.Csi_LogMeasInfoList_r19[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CSI_LogMeasInfoConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSILogMeasInfoConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. refCSI-LoggedMeasurementConfigId-r19: ref
	{
		if err := ie.RefCSI_LoggedMeasurementConfigId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 3. csi-LogMeasInfoList-r19: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cSILogMeasInfoConfigR19CsiLogMeasInfoListR19Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Csi_LogMeasInfoList_r19 = make([]CSI_LogMeasInfo_r19, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Csi_LogMeasInfoList_r19[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
