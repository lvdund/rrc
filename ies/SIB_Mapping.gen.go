// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB-Mapping (line 15088).
// SIB-Mapping ::=                     SEQUENCE (SIZE (1..maxSIB)) OF SIB-TypeInfo

var sIBMappingSizeConstraints = per.SizeRange(1, common.MaxSIB)

type SIB_Mapping []SIB_TypeInfo

func (ie *SIB_Mapping) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sIBMappingSizeConstraints)
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

func (ie *SIB_Mapping) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sIBMappingSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SIB_Mapping, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
