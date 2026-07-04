// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NR-NS-PmaxListAerial-r18 (line 10757).
// NR-NS-PmaxListAerial-r18 ::=            SEQUENCE (SIZE (1..maxNR-NS-Pmax)) OF NR-NS-PmaxValueAerial-r18

var nRNSPmaxListAerialR18SizeConstraints = per.SizeRange(1, common.MaxNR_NS_Pmax)

type NR_NS_PmaxListAerial_r18 []NR_NS_PmaxValueAerial_r18

func (ie *NR_NS_PmaxListAerial_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nRNSPmaxListAerialR18SizeConstraints)
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

func (ie *NR_NS_PmaxListAerial_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nRNSPmaxListAerialR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(NR_NS_PmaxListAerial_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
