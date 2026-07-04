// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResumeRequest-IEs (line 1659).

var rRCResumeRequestIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resumeIdentity"},
		{Name: "resumeMAC-I"},
		{Name: "resumeCause"},
		{Name: "spare"},
	},
}

var rRCResumeRequestIEsResumeMACIConstraints = per.FixedSize(16)

var rRCResumeRequestIEsSpareConstraints = per.FixedSize(1)

type RRCResumeRequest_IEs struct {
	ResumeIdentity ShortI_RNTI_Value
	ResumeMAC_I    per.BitString
	ResumeCause    ResumeCause
	Spare          per.BitString
}

func (ie *RRCResumeRequest_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeRequestIEsConstraints)
	_ = seq

	// 1. resumeIdentity: ref
	{
		if err := ie.ResumeIdentity.Encode(e); err != nil {
			return err
		}
	}

	// 2. resumeMAC-I: bit-string
	{
		if err := e.EncodeBitString(ie.ResumeMAC_I, rRCResumeRequestIEsResumeMACIConstraints); err != nil {
			return err
		}
	}

	// 3. resumeCause: ref
	{
		if err := ie.ResumeCause.Encode(e); err != nil {
			return err
		}
	}

	// 4. spare: bit-string
	{
		if err := e.EncodeBitString(ie.Spare, rRCResumeRequestIEsSpareConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RRCResumeRequest_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeRequestIEsConstraints)
	_ = seq

	// 1. resumeIdentity: ref
	{
		if err := ie.ResumeIdentity.Decode(d); err != nil {
			return err
		}
	}

	// 2. resumeMAC-I: bit-string
	{
		v1, err := d.DecodeBitString(rRCResumeRequestIEsResumeMACIConstraints)
		if err != nil {
			return err
		}
		ie.ResumeMAC_I = v1
	}

	// 3. resumeCause: ref
	{
		if err := ie.ResumeCause.Decode(d); err != nil {
			return err
		}
	}

	// 4. spare: bit-string
	{
		v3, err := d.DecodeBitString(rRCResumeRequestIEsSpareConstraints)
		if err != nil {
			return err
		}
		ie.Spare = v3
	}

	return nil
}
