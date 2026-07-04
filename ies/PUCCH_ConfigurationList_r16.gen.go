// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-ConfigurationList-r16 (line 12208).
// PUCCH-ConfigurationList-r16  ::=     SEQUENCE (SIZE (1..2)) OF PUCCH-Config

var pUCCHConfigurationListR16SizeConstraints = per.SizeRange(1, 2)

type PUCCH_ConfigurationList_r16 []PUCCH_Config

func (ie *PUCCH_ConfigurationList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(pUCCHConfigurationListR16SizeConstraints)
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

func (ie *PUCCH_ConfigurationList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(pUCCHConfigurationListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(PUCCH_ConfigurationList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
