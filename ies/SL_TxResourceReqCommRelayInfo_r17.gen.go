// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-TxResourceReqCommRelayInfo-r17 (line 2218).

var sLTxResourceReqCommRelayInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RelayDRXConfig-r17", Optional: true},
		{Name: "sl-TxResourceReqCommRelay-r17"},
	},
}

type SL_TxResourceReqCommRelayInfo_r17 struct {
	Sl_RelayDRXConfig_r17         *SL_TxResourceReq_v1700
	Sl_TxResourceReqCommRelay_r17 SL_TxResourceReqCommRelay_r17
}

func (ie *SL_TxResourceReqCommRelayInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLTxResourceReqCommRelayInfoR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RelayDRXConfig_r17 != nil}); err != nil {
		return err
	}

	// 2. sl-RelayDRXConfig-r17: ref
	{
		if ie.Sl_RelayDRXConfig_r17 != nil {
			if err := ie.Sl_RelayDRXConfig_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sl-TxResourceReqCommRelay-r17: ref
	{
		if err := ie.Sl_TxResourceReqCommRelay_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_TxResourceReqCommRelayInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLTxResourceReqCommRelayInfoR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RelayDRXConfig-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_RelayDRXConfig_r17 = new(SL_TxResourceReq_v1700)
			if err := ie.Sl_RelayDRXConfig_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sl-TxResourceReqCommRelay-r17: ref
	{
		if err := ie.Sl_TxResourceReqCommRelay_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
