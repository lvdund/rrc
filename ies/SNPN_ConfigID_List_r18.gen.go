// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SNPN-ConfigID-List-r18 (line 26112).
// SNPN-ConfigID-List-r18 ::=       SEQUENCE (SIZE (1..maxSNPN-ConfigID-r18)) OF SNPN-ConfigID-r18

var sNPNConfigIDListR18SizeConstraints = per.SizeRange(1, common.MaxSNPN_ConfigID_r18)

type SNPN_ConfigID_List_r18 []SNPN_ConfigID_r18

func (ie *SNPN_ConfigID_List_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sNPNConfigIDListR18SizeConstraints)
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

func (ie *SNPN_ConfigID_List_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sNPNConfigIDListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SNPN_ConfigID_List_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
