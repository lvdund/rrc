// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AdditionalRACH-ConfigList-r17 (line 5371).
// AdditionalRACH-ConfigList-r17 ::=       SEQUENCE (SIZE(1..maxAdditionalRACH-r17)) OF AdditionalRACH-Config-r17

var additionalRACHConfigListR17SizeConstraints = per.SizeRange(1, common.MaxAdditionalRACH_r17)

type AdditionalRACH_ConfigList_r17 []AdditionalRACH_Config_r17

func (ie *AdditionalRACH_ConfigList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(additionalRACHConfigListR17SizeConstraints)
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

func (ie *AdditionalRACH_ConfigList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(additionalRACHConfigListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(AdditionalRACH_ConfigList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
