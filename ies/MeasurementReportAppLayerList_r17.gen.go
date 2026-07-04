// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasurementReportAppLayerList-r17 (line 764).
// MeasurementReportAppLayerList-r17 ::= SEQUENCE (SIZE (1..maxNrofAppLayerMeas-r17)) OF MeasReportAppLayer-r17

var measurementReportAppLayerListR17SizeConstraints = per.SizeRange(1, common.MaxNrofAppLayerMeas_r17)

type MeasurementReportAppLayerList_r17 []MeasReportAppLayer_r17

func (ie *MeasurementReportAppLayerList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measurementReportAppLayerListR17SizeConstraints)
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

func (ie *MeasurementReportAppLayerList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measurementReportAppLayerListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasurementReportAppLayerList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
