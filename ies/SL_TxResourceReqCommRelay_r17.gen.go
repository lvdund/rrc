// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-TxResourceReqCommRelay-r17 (line 2223).

var sLTxResourceReqCommRelayR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-TxResourceReqL2U2N-Relay-r17"},
		{Name: "sl-TxResourceReqL3U2N-Relay-r17"},
	},
}

const (
	SL_TxResourceReqCommRelay_r17_Sl_TxResourceReqL2U2N_Relay_r17 = 0
	SL_TxResourceReqCommRelay_r17_Sl_TxResourceReqL3U2N_Relay_r17 = 1
)

type SL_TxResourceReqCommRelay_r17 struct {
	Choice                          int
	Sl_TxResourceReqL2U2N_Relay_r17 *SL_TxResourceReqL2U2N_Relay_r17
	Sl_TxResourceReqL3U2N_Relay_r17 *SL_TxResourceReq_r16
}

func (ie *SL_TxResourceReqCommRelay_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(sLTxResourceReqCommRelayR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_TxResourceReqCommRelay_r17_Sl_TxResourceReqL2U2N_Relay_r17:
		if err := ie.Sl_TxResourceReqL2U2N_Relay_r17.Encode(e); err != nil {
			return err
		}
	case SL_TxResourceReqCommRelay_r17_Sl_TxResourceReqL3U2N_Relay_r17:
		if err := ie.Sl_TxResourceReqL3U2N_Relay_r17.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "SL-TxResourceReqCommRelay-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *SL_TxResourceReqCommRelay_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(sLTxResourceReqCommRelayR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "SL-TxResourceReqCommRelay-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case SL_TxResourceReqCommRelay_r17_Sl_TxResourceReqL2U2N_Relay_r17:
		ie.Sl_TxResourceReqL2U2N_Relay_r17 = new(SL_TxResourceReqL2U2N_Relay_r17)
		if err := ie.Sl_TxResourceReqL2U2N_Relay_r17.Decode(d); err != nil {
			return err
		}
	case SL_TxResourceReqCommRelay_r17_Sl_TxResourceReqL3U2N_Relay_r17:
		ie.Sl_TxResourceReqL3U2N_Relay_r17 = new(SL_TxResourceReq_r16)
		if err := ie.Sl_TxResourceReqL3U2N_Relay_r17.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "SL-TxResourceReqCommRelay-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
