// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParametersSidelinkEUTRA-NR-v1630 (line 17161).

var bandParametersSidelinkEUTRANRV1630Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eutra"},
		{Name: "nr"},
	},
}

const (
	BandParametersSidelinkEUTRA_NR_v1630_Eutra = 0
	BandParametersSidelinkEUTRA_NR_v1630_Nr    = 1
)

var bandParametersSidelinkEUTRANRV1630NrConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tx-Sidelink-r16", Optional: true},
		{Name: "rx-Sidelink-r16", Optional: true},
		{Name: "sl-CrossCarrierScheduling-r16", Optional: true},
	},
}

const (
	BandParametersSidelinkEUTRA_NR_v1630_Nr_Tx_Sidelink_r16_Supported = 0
)

var bandParametersSidelinkEUTRANRV1630NrTxSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkEUTRA_NR_v1630_Nr_Tx_Sidelink_r16_Supported},
}

const (
	BandParametersSidelinkEUTRA_NR_v1630_Nr_Rx_Sidelink_r16_Supported = 0
)

var bandParametersSidelinkEUTRANRV1630NrRxSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkEUTRA_NR_v1630_Nr_Rx_Sidelink_r16_Supported},
}

const (
	BandParametersSidelinkEUTRA_NR_v1630_Nr_Sl_CrossCarrierScheduling_r16_Supported = 0
)

var bandParametersSidelinkEUTRANRV1630NrSlCrossCarrierSchedulingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParametersSidelinkEUTRA_NR_v1630_Nr_Sl_CrossCarrierScheduling_r16_Supported},
}

type BandParametersSidelinkEUTRA_NR_v1630 struct {
	Choice int
	Eutra  *struct{}
	Nr     *struct {
		Tx_Sidelink_r16               *int64
		Rx_Sidelink_r16               *int64
		Sl_CrossCarrierScheduling_r16 *int64
	}
}

func (ie *BandParametersSidelinkEUTRA_NR_v1630) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(bandParametersSidelinkEUTRANRV1630Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_v1630_Eutra:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case BandParametersSidelinkEUTRA_NR_v1630_Nr:
		c := (*ie.Nr)
		bandParametersSidelinkEUTRANRV1630NrSeq := e.NewSequenceEncoder(bandParametersSidelinkEUTRANRV1630NrConstraints)
		if err := bandParametersSidelinkEUTRANRV1630NrSeq.EncodePreamble([]bool{c.Tx_Sidelink_r16 != nil, c.Rx_Sidelink_r16 != nil, c.Sl_CrossCarrierScheduling_r16 != nil}); err != nil {
			return err
		}
		if c.Tx_Sidelink_r16 != nil {
			if err := e.EncodeEnumerated((*c.Tx_Sidelink_r16), bandParametersSidelinkEUTRANRV1630NrTxSidelinkR16Constraints); err != nil {
				return err
			}
		}
		if c.Rx_Sidelink_r16 != nil {
			if err := e.EncodeEnumerated((*c.Rx_Sidelink_r16), bandParametersSidelinkEUTRANRV1630NrRxSidelinkR16Constraints); err != nil {
				return err
			}
		}
		if c.Sl_CrossCarrierScheduling_r16 != nil {
			if err := e.EncodeEnumerated((*c.Sl_CrossCarrierScheduling_r16), bandParametersSidelinkEUTRANRV1630NrSlCrossCarrierSchedulingR16Constraints); err != nil {
				return err
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "BandParametersSidelinkEUTRA-NR-v1630",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_v1630) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(bandParametersSidelinkEUTRANRV1630Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "BandParametersSidelinkEUTRA-NR-v1630", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case BandParametersSidelinkEUTRA_NR_v1630_Eutra:
		ie.Eutra = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case BandParametersSidelinkEUTRA_NR_v1630_Nr:
		ie.Nr = &struct {
			Tx_Sidelink_r16               *int64
			Rx_Sidelink_r16               *int64
			Sl_CrossCarrierScheduling_r16 *int64
		}{}
		c := (*ie.Nr)
		bandParametersSidelinkEUTRANRV1630NrSeq := d.NewSequenceDecoder(bandParametersSidelinkEUTRANRV1630NrConstraints)
		if err := bandParametersSidelinkEUTRANRV1630NrSeq.DecodePreamble(); err != nil {
			return err
		}
		if bandParametersSidelinkEUTRANRV1630NrSeq.IsComponentPresent(0) {
			c.Tx_Sidelink_r16 = new(int64)
			v, err := d.DecodeEnumerated(bandParametersSidelinkEUTRANRV1630NrTxSidelinkR16Constraints)
			if err != nil {
				return err
			}
			(*c.Tx_Sidelink_r16) = v
		}
		if bandParametersSidelinkEUTRANRV1630NrSeq.IsComponentPresent(1) {
			c.Rx_Sidelink_r16 = new(int64)
			v, err := d.DecodeEnumerated(bandParametersSidelinkEUTRANRV1630NrRxSidelinkR16Constraints)
			if err != nil {
				return err
			}
			(*c.Rx_Sidelink_r16) = v
		}
		if bandParametersSidelinkEUTRANRV1630NrSeq.IsComponentPresent(2) {
			c.Sl_CrossCarrierScheduling_r16 = new(int64)
			v, err := d.DecodeEnumerated(bandParametersSidelinkEUTRANRV1630NrSlCrossCarrierSchedulingR16Constraints)
			if err != nil {
				return err
			}
			(*c.Sl_CrossCarrierScheduling_r16) = v
		}
	default:
		return &per.DecodeError{TypeName: "BandParametersSidelinkEUTRA-NR-v1630", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
