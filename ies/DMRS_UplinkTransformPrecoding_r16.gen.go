// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DMRS-UplinkTransformPrecoding-r16 (line 7809).

var dMRSUplinkTransformPrecodingR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pi2BPSK-ScramblingID0", Optional: true},
		{Name: "pi2BPSK-ScramblingID1", Optional: true},
	},
}

var dMRSUplinkTransformPrecodingR16Pi2BPSKScramblingID0Constraints = per.Constrained(0, 65535)

var dMRSUplinkTransformPrecodingR16Pi2BPSKScramblingID1Constraints = per.Constrained(0, 65535)

type DMRS_UplinkTransformPrecoding_r16 struct {
	Pi2BPSK_ScramblingID0 *int64
	Pi2BPSK_ScramblingID1 *int64
}

func (ie *DMRS_UplinkTransformPrecoding_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dMRSUplinkTransformPrecodingR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pi2BPSK_ScramblingID0 != nil, ie.Pi2BPSK_ScramblingID1 != nil}); err != nil {
		return err
	}

	// 2. pi2BPSK-ScramblingID0: integer
	{
		if ie.Pi2BPSK_ScramblingID0 != nil {
			if err := e.EncodeInteger(*ie.Pi2BPSK_ScramblingID0, dMRSUplinkTransformPrecodingR16Pi2BPSKScramblingID0Constraints); err != nil {
				return err
			}
		}
	}

	// 3. pi2BPSK-ScramblingID1: integer
	{
		if ie.Pi2BPSK_ScramblingID1 != nil {
			if err := e.EncodeInteger(*ie.Pi2BPSK_ScramblingID1, dMRSUplinkTransformPrecodingR16Pi2BPSKScramblingID1Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DMRS_UplinkTransformPrecoding_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dMRSUplinkTransformPrecodingR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pi2BPSK-ScramblingID0: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(dMRSUplinkTransformPrecodingR16Pi2BPSKScramblingID0Constraints)
			if err != nil {
				return err
			}
			ie.Pi2BPSK_ScramblingID0 = &v
		}
	}

	// 3. pi2BPSK-ScramblingID1: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(dMRSUplinkTransformPrecodingR16Pi2BPSKScramblingID1Constraints)
			if err != nil {
				return err
			}
			ie.Pi2BPSK_ScramblingID1 = &v
		}
	}

	return nil
}
