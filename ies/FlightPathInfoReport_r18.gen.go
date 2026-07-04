// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FlightPathInfoReport-r18 (line 2982).
// FlightPathInfoReport-r18 ::=         SEQUENCE (SIZE (0..maxWayPoint-r18)) OF WayPoint-r18

var flightPathInfoReportR18SizeConstraints = per.SizeRange(0, common.MaxWayPoint_r18)

type FlightPathInfoReport_r18 []WayPoint_r18

func (ie *FlightPathInfoReport_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(flightPathInfoReportR18SizeConstraints)
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

func (ie *FlightPathInfoReport_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(flightPathInfoReportR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(FlightPathInfoReport_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
