// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MsgA-DMRS-Config-r16 (line 10153).

var msgADMRSConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "msgA-DMRS-AdditionalPosition-r16", Optional: true},
		{Name: "msgA-MaxLength-r16", Optional: true},
		{Name: "msgA-PUSCH-DMRS-CDM-Group-r16", Optional: true},
		{Name: "msgA-PUSCH-NrofPorts-r16", Optional: true},
		{Name: "msgA-ScramblingID0-r16", Optional: true},
		{Name: "msgA-ScramblingID1-r16", Optional: true},
	},
}

const (
	MsgA_DMRS_Config_r16_MsgA_DMRS_AdditionalPosition_r16_Pos0 = 0
	MsgA_DMRS_Config_r16_MsgA_DMRS_AdditionalPosition_r16_Pos1 = 1
	MsgA_DMRS_Config_r16_MsgA_DMRS_AdditionalPosition_r16_Pos3 = 2
)

var msgADMRSConfigR16MsgADMRSAdditionalPositionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MsgA_DMRS_Config_r16_MsgA_DMRS_AdditionalPosition_r16_Pos0, MsgA_DMRS_Config_r16_MsgA_DMRS_AdditionalPosition_r16_Pos1, MsgA_DMRS_Config_r16_MsgA_DMRS_AdditionalPosition_r16_Pos3},
}

const (
	MsgA_DMRS_Config_r16_MsgA_MaxLength_r16_Len2 = 0
)

var msgADMRSConfigR16MsgAMaxLengthR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MsgA_DMRS_Config_r16_MsgA_MaxLength_r16_Len2},
}

var msgADMRSConfigR16MsgAPUSCHDMRSCDMGroupR16Constraints = per.Constrained(0, 1)

var msgADMRSConfigR16MsgAPUSCHNrofPortsR16Constraints = per.Constrained(0, 1)

var msgADMRSConfigR16MsgAScramblingID0R16Constraints = per.Constrained(0, 65535)

var msgADMRSConfigR16MsgAScramblingID1R16Constraints = per.Constrained(0, 65535)

type MsgA_DMRS_Config_r16 struct {
	MsgA_DMRS_AdditionalPosition_r16 *int64
	MsgA_MaxLength_r16               *int64
	MsgA_PUSCH_DMRS_CDM_Group_r16    *int64
	MsgA_PUSCH_NrofPorts_r16         *int64
	MsgA_ScramblingID0_r16           *int64
	MsgA_ScramblingID1_r16           *int64
}

func (ie *MsgA_DMRS_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(msgADMRSConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MsgA_DMRS_AdditionalPosition_r16 != nil, ie.MsgA_MaxLength_r16 != nil, ie.MsgA_PUSCH_DMRS_CDM_Group_r16 != nil, ie.MsgA_PUSCH_NrofPorts_r16 != nil, ie.MsgA_ScramblingID0_r16 != nil, ie.MsgA_ScramblingID1_r16 != nil}); err != nil {
		return err
	}

	// 2. msgA-DMRS-AdditionalPosition-r16: enumerated
	{
		if ie.MsgA_DMRS_AdditionalPosition_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_DMRS_AdditionalPosition_r16, msgADMRSConfigR16MsgADMRSAdditionalPositionR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. msgA-MaxLength-r16: enumerated
	{
		if ie.MsgA_MaxLength_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_MaxLength_r16, msgADMRSConfigR16MsgAMaxLengthR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. msgA-PUSCH-DMRS-CDM-Group-r16: integer
	{
		if ie.MsgA_PUSCH_DMRS_CDM_Group_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_PUSCH_DMRS_CDM_Group_r16, msgADMRSConfigR16MsgAPUSCHDMRSCDMGroupR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. msgA-PUSCH-NrofPorts-r16: integer
	{
		if ie.MsgA_PUSCH_NrofPorts_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_PUSCH_NrofPorts_r16, msgADMRSConfigR16MsgAPUSCHNrofPortsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. msgA-ScramblingID0-r16: integer
	{
		if ie.MsgA_ScramblingID0_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_ScramblingID0_r16, msgADMRSConfigR16MsgAScramblingID0R16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. msgA-ScramblingID1-r16: integer
	{
		if ie.MsgA_ScramblingID1_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_ScramblingID1_r16, msgADMRSConfigR16MsgAScramblingID1R16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MsgA_DMRS_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(msgADMRSConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. msgA-DMRS-AdditionalPosition-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(msgADMRSConfigR16MsgADMRSAdditionalPositionR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_DMRS_AdditionalPosition_r16 = &idx
		}
	}

	// 3. msgA-MaxLength-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(msgADMRSConfigR16MsgAMaxLengthR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_MaxLength_r16 = &idx
		}
	}

	// 4. msgA-PUSCH-DMRS-CDM-Group-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(msgADMRSConfigR16MsgAPUSCHDMRSCDMGroupR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_PUSCH_DMRS_CDM_Group_r16 = &v
		}
	}

	// 5. msgA-PUSCH-NrofPorts-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(msgADMRSConfigR16MsgAPUSCHNrofPortsR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_PUSCH_NrofPorts_r16 = &v
		}
	}

	// 6. msgA-ScramblingID0-r16: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(msgADMRSConfigR16MsgAScramblingID0R16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_ScramblingID0_r16 = &v
		}
	}

	// 7. msgA-ScramblingID1-r16: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(msgADMRSConfigR16MsgAScramblingID1R16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_ScramblingID1_r16 = &v
		}
	}

	return nil
}
