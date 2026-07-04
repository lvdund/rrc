// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-CodeBlockGroupTransmissionList-r16 (line 11505).
// PDSCH-CodeBlockGroupTransmissionList-r16 ::=    SEQUENCE (SIZE (1..2)) OF PDSCH-CodeBlockGroupTransmission

var pDSCHCodeBlockGroupTransmissionListR16SizeConstraints = per.SizeRange(1, 2)

type PDSCH_CodeBlockGroupTransmissionList_r16 []PDSCH_CodeBlockGroupTransmission

func (ie *PDSCH_CodeBlockGroupTransmissionList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pDSCHCodeBlockGroupTransmissionListR16SizeConstraints)
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

func (ie *PDSCH_CodeBlockGroupTransmissionList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pDSCHCodeBlockGroupTransmissionListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PDSCH_CodeBlockGroupTransmissionList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
