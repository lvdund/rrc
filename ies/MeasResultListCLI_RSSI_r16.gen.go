// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultListCLI-RSSI-r16 (line 9865).
// MeasResultListCLI-RSSI-r16 ::=   SEQUENCE (SIZE (1.. maxCLI-Report-r16)) OF MeasResultCLI-RSSI-r16

var measResultListCLIRSSIR16SizeConstraints = per.SizeRange(1, common.MaxCLI_Report_r16)

type MeasResultListCLI_RSSI_r16 []MeasResultCLI_RSSI_r16

func (ie *MeasResultListCLI_RSSI_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultListCLIRSSIR16SizeConstraints)
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

func (ie *MeasResultListCLI_RSSI_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultListCLIRSSIR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultListCLI_RSSI_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
