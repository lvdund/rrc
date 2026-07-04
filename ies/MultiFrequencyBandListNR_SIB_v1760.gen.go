// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MultiFrequencyBandListNR-SIB-v1760 (line 10177).
// MultiFrequencyBandListNR-SIB-v1760 ::=      SEQUENCE (SIZE (1.. maxNrofMultiBands)) OF NR-MultiBandInfo-v1760

var multiFrequencyBandListNRSIBV1760SizeConstraints = per.SizeRange(1, common.MaxNrofMultiBands)

type MultiFrequencyBandListNR_SIB_v1760 []NR_MultiBandInfo_v1760

func (ie *MultiFrequencyBandListNR_SIB_v1760) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(multiFrequencyBandListNRSIBV1760SizeConstraints)
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

func (ie *MultiFrequencyBandListNR_SIB_v1760) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(multiFrequencyBandListNRSIBV1760SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MultiFrequencyBandListNR_SIB_v1760, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
