// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasurementReportAppLayerList-v1800 (line 766).
// MeasurementReportAppLayerList-v1800 ::= SEQUENCE (SIZE (1..maxNrofAppLayerMeas-r17)) OF MeasReportAppLayer-v1800

var measurementReportAppLayerListV1800SizeConstraints = per.SizeRange(1, common.MaxNrofAppLayerMeas_r17)

type MeasurementReportAppLayerList_v1800 []MeasReportAppLayer_v1800

func (ie *MeasurementReportAppLayerList_v1800) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measurementReportAppLayerListV1800SizeConstraints)
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

func (ie *MeasurementReportAppLayerList_v1800) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measurementReportAppLayerListV1800SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasurementReportAppLayerList_v1800, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
