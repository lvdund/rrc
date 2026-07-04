// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasReportAppLayer-v1800 (line 775).

var measReportAppLayerV1800Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "appLayerIdleInactiveConfig-r18", Optional: true},
		{Name: "measReportAppLayerContainerList-r18", Optional: true},
	},
}

var measReportAppLayerV1800MeasReportAppLayerContainerListR18Constraints = per.SizeRange(1, common.MaxNrofAppLayerReports_r18)

type MeasReportAppLayer_v1800 struct {
	AppLayerIdleInactiveConfig_r18      *AppLayerIdleInactiveConfig_r18
	MeasReportAppLayerContainerList_r18 [][]byte
}

func (ie *MeasReportAppLayer_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measReportAppLayerV1800Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AppLayerIdleInactiveConfig_r18 != nil, ie.MeasReportAppLayerContainerList_r18 != nil}); err != nil {
		return err
	}

	// 3. appLayerIdleInactiveConfig-r18: ref
	{
		if ie.AppLayerIdleInactiveConfig_r18 != nil {
			if err := ie.AppLayerIdleInactiveConfig_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. measReportAppLayerContainerList-r18: sequence-of
	{
		if ie.MeasReportAppLayerContainerList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(measReportAppLayerV1800MeasReportAppLayerContainerListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.MeasReportAppLayerContainerList_r18))); err != nil {
				return err
			}
			for i := range ie.MeasReportAppLayerContainerList_r18 {
				if err := e.EncodeOctetString(ie.MeasReportAppLayerContainerList_r18[i], per.SizeConstraints{}); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MeasReportAppLayer_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measReportAppLayerV1800Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. appLayerIdleInactiveConfig-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.AppLayerIdleInactiveConfig_r18 = new(AppLayerIdleInactiveConfig_r18)
			if err := ie.AppLayerIdleInactiveConfig_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. measReportAppLayerContainerList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(measReportAppLayerV1800MeasReportAppLayerContainerListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MeasReportAppLayerContainerList_r18 = make([][]byte, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeOctetString(per.SizeConstraints{})
				if err != nil {
					return err
				}
				ie.MeasReportAppLayerContainerList_r18[i] = v
			}
		}
	}

	return nil
}
