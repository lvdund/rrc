// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: WLAN-NameList-r16 (line 26740).
// WLAN-NameList-r16 ::= SEQUENCE (SIZE (1..maxWLAN-Name-r16)) OF WLAN-Name-r16

var wLANNameListR16SizeConstraints = per.SizeRange(1, common.MaxWLAN_Name_r16)

type WLAN_NameList_r16 []WLAN_Name_r16

func (ie *WLAN_NameList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(wLANNameListR16SizeConstraints)
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

func (ie *WLAN_NameList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(wLANNameListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(WLAN_NameList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
