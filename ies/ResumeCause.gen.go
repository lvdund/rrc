// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ResumeCause (line 14003).

const (
	ResumeCause_Emergency                          = 0
	ResumeCause_HighPriorityAccess                 = 1
	ResumeCause_Mt_Access                          = 2
	ResumeCause_Mo_Signalling                      = 3
	ResumeCause_Mo_Data                            = 4
	ResumeCause_Mo_VoiceCall                       = 5
	ResumeCause_Mo_VideoCall                       = 6
	ResumeCause_Mo_SMS                             = 7
	ResumeCause_Rna_Update                         = 8
	ResumeCause_Mps_PriorityAccess                 = 9
	ResumeCause_Mcs_PriorityAccess                 = 10
	ResumeCause_Mt_SDT_v1810                       = 11
	ResumeCause_Srs_PosConfigOrActivationReq_v1800 = 12
	ResumeCause_Spare3                             = 13
	ResumeCause_Spare2                             = 14
	ResumeCause_Spare1                             = 15
)

var resumeCauseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ResumeCause_Emergency, ResumeCause_HighPriorityAccess, ResumeCause_Mt_Access, ResumeCause_Mo_Signalling, ResumeCause_Mo_Data, ResumeCause_Mo_VoiceCall, ResumeCause_Mo_VideoCall, ResumeCause_Mo_SMS, ResumeCause_Rna_Update, ResumeCause_Mps_PriorityAccess, ResumeCause_Mcs_PriorityAccess, ResumeCause_Mt_SDT_v1810, ResumeCause_Srs_PosConfigOrActivationReq_v1800, ResumeCause_Spare3, ResumeCause_Spare2, ResumeCause_Spare1},
}

type ResumeCause struct {
	Value int64
}

func (v *ResumeCause) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, resumeCauseConstraints)
}

func (v *ResumeCause) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(resumeCauseConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
