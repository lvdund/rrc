// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RSSI-ResourceListConfigCLI-r16 (line 9349).
// RSSI-ResourceListConfigCLI-r16 ::=  SEQUENCE (SIZE (1.. maxNrofCLI-RSSI-Resources-r16)) OF RSSI-ResourceConfigCLI-r16

var rSSIResourceListConfigCLIR16SizeConstraints = per.SizeRange(1, common.MaxNrofCLI_RSSI_Resources_r16)

type RSSI_ResourceListConfigCLI_r16 []RSSI_ResourceConfigCLI_r16

func (ie *RSSI_ResourceListConfigCLI_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(rSSIResourceListConfigCLIR16SizeConstraints)
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

func (ie *RSSI_ResourceListConfigCLI_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(rSSIResourceListConfigCLIR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(RSSI_ResourceListConfigCLI_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
