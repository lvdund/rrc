// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB-Mapping-v1700 (line 15090).
// SIB-Mapping-v1700  ::=              SEQUENCE (SIZE (1..maxSIB)) OF SIB-TypeInfo-v1700

var sIBMappingV1700SizeConstraints = per.SizeRange(1, common.MaxSIB)

type SIB_Mapping_v1700 []SIB_TypeInfo_v1700

func (ie *SIB_Mapping_v1700) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sIBMappingV1700SizeConstraints)
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

func (ie *SIB_Mapping_v1700) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sIBMappingV1700SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SIB_Mapping_v1700, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
