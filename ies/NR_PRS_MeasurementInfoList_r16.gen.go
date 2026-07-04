// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NR-PRS-MeasurementInfoList-r16 (line 8575).
// NR-PRS-MeasurementInfoList-r16 ::= SEQUENCE (SIZE (1..maxFreqLayers)) OF NR-PRS-MeasurementInfo-r16

var nRPRSMeasurementInfoListR16SizeConstraints = per.SizeRange(1, common.MaxFreqLayers)

type NR_PRS_MeasurementInfoList_r16 []NR_PRS_MeasurementInfo_r16

func (ie *NR_PRS_MeasurementInfoList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nRPRSMeasurementInfoListR16SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NR_PRS_MeasurementInfoList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nRPRSMeasurementInfoListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(NR_PRS_MeasurementInfoList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
