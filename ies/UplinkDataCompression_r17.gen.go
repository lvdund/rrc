// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkDataCompression-r17 (line 11285).

var uplinkDataCompressionR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "newSetup"},
		{Name: "drb-ContinueUDC"},
	},
}

const (
	UplinkDataCompression_r17_NewSetup        = 0
	UplinkDataCompression_r17_Drb_ContinueUDC = 1
)

var uplinkDataCompressionR17NewSetupConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bufferSize-r17"},
		{Name: "dictionary-r17", Optional: true},
	},
}

const (
	UplinkDataCompression_r17_NewSetup_BufferSize_r17_Kbyte2 = 0
	UplinkDataCompression_r17_NewSetup_BufferSize_r17_Kbyte4 = 1
	UplinkDataCompression_r17_NewSetup_BufferSize_r17_Kbyte8 = 2
	UplinkDataCompression_r17_NewSetup_BufferSize_r17_Spare1 = 3
)

var uplinkDataCompressionR17NewSetupBufferSizeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkDataCompression_r17_NewSetup_BufferSize_r17_Kbyte2, UplinkDataCompression_r17_NewSetup_BufferSize_r17_Kbyte4, UplinkDataCompression_r17_NewSetup_BufferSize_r17_Kbyte8, UplinkDataCompression_r17_NewSetup_BufferSize_r17_Spare1},
}

const (
	UplinkDataCompression_r17_NewSetup_Dictionary_r17_Sip_SDP  = 0
	UplinkDataCompression_r17_NewSetup_Dictionary_r17_Operator = 1
)

var uplinkDataCompressionR17NewSetupDictionaryR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkDataCompression_r17_NewSetup_Dictionary_r17_Sip_SDP, UplinkDataCompression_r17_NewSetup_Dictionary_r17_Operator},
}

type UplinkDataCompression_r17 struct {
	Choice   int
	NewSetup *struct {
		BufferSize_r17 int64
		Dictionary_r17 *int64
	}
	Drb_ContinueUDC *struct{}
}

func (ie *UplinkDataCompression_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(uplinkDataCompressionR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case UplinkDataCompression_r17_NewSetup:
		c := (*ie.NewSetup)
		uplinkDataCompressionR17NewSetupSeq := e.NewSequenceEncoder(uplinkDataCompressionR17NewSetupConstraints)
		if err := uplinkDataCompressionR17NewSetupSeq.EncodePreamble([]bool{c.Dictionary_r17 != nil}); err != nil {
			return err
		}
		if err := e.EncodeEnumerated(c.BufferSize_r17, uplinkDataCompressionR17NewSetupBufferSizeR17Constraints); err != nil {
			return err
		}
		if c.Dictionary_r17 != nil {
			if err := e.EncodeEnumerated((*c.Dictionary_r17), uplinkDataCompressionR17NewSetupDictionaryR17Constraints); err != nil {
				return err
			}
		}
	case UplinkDataCompression_r17_Drb_ContinueUDC:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "UplinkDataCompression-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *UplinkDataCompression_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(uplinkDataCompressionR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "UplinkDataCompression-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case UplinkDataCompression_r17_NewSetup:
		ie.NewSetup = &struct {
			BufferSize_r17 int64
			Dictionary_r17 *int64
		}{}
		c := (*ie.NewSetup)
		uplinkDataCompressionR17NewSetupSeq := d.NewSequenceDecoder(uplinkDataCompressionR17NewSetupConstraints)
		if err := uplinkDataCompressionR17NewSetupSeq.DecodePreamble(); err != nil {
			return err
		}
		{
			v, err := d.DecodeEnumerated(uplinkDataCompressionR17NewSetupBufferSizeR17Constraints)
			if err != nil {
				return err
			}
			c.BufferSize_r17 = v
		}
		if uplinkDataCompressionR17NewSetupSeq.IsComponentPresent(1) {
			c.Dictionary_r17 = new(int64)
			v, err := d.DecodeEnumerated(uplinkDataCompressionR17NewSetupDictionaryR17Constraints)
			if err != nil {
				return err
			}
			(*c.Dictionary_r17) = v
		}
	case UplinkDataCompression_r17_Drb_ContinueUDC:
		ie.Drb_ContinueUDC = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "UplinkDataCompression-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
