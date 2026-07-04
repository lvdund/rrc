// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MultiFrequencyBandListNR-Aerial-SIB-r18 (line 10183).
// MultiFrequencyBandListNR-Aerial-SIB-r18 ::= SEQUENCE (SIZE (1.. maxNrofMultiBands)) OF NR-MultiBandInfoAerial-r18

var multiFrequencyBandListNRAerialSIBR18SizeConstraints = per.SizeRange(1, common.MaxNrofMultiBands)

type MultiFrequencyBandListNR_Aerial_SIB_r18 []NR_MultiBandInfoAerial_r18

func (ie *MultiFrequencyBandListNR_Aerial_SIB_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(multiFrequencyBandListNRAerialSIBR18SizeConstraints)
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

func (ie *MultiFrequencyBandListNR_Aerial_SIB_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(multiFrequencyBandListNRAerialSIBR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MultiFrequencyBandListNR_Aerial_SIB_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
