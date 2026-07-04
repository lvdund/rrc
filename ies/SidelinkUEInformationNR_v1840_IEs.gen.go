// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SidelinkUEInformationNR-v1840-IEs (line 2155).

var sidelinkUEInformationNRV1840IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PosRxInterestedFreqList2-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type SidelinkUEInformationNR_v1840_IEs struct {
	Sl_PosRxInterestedFreqList2_r18 *SL_InterestedFreqList_r16
	NonCriticalExtension            *struct{}
}

func (ie *SidelinkUEInformationNR_v1840_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sidelinkUEInformationNRV1840IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PosRxInterestedFreqList2_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sl-PosRxInterestedFreqList2-r18: ref
	{
		if ie.Sl_PosRxInterestedFreqList2_r18 != nil {
			if err := ie.Sl_PosRxInterestedFreqList2_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *SidelinkUEInformationNR_v1840_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sidelinkUEInformationNRV1840IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PosRxInterestedFreqList2-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_PosRxInterestedFreqList2_r18 = new(SL_InterestedFreqList_r16)
			if err := ie.Sl_PosRxInterestedFreqList2_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
