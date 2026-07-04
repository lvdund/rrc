// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CondReconfigToAddMod-r16 (line 6465).

var condReconfigToAddModR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "condReconfigId-r16"},
		{Name: "condExecutionCond-r16", Optional: true},
		{Name: "condRRCReconfig-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var condReconfigToAddModR16CondExecutionCondR16Constraints = per.SizeRange(1, 2)

var condReconfigToAddModR16CondRRCReconfigR16Constraints = per.SizeConstraints{}

var condReconfigToAddModR16CondExecutionCondSCGR17Constraints = per.SizeConstraints{}

var condReconfigToAddModR16ExtCondExecutionCondPSCellR18Constraints = per.SizeRange(1, 2)

const (
	CondReconfigToAddMod_r16_Ext_Scpac_ConfigComplete_r18_True = 0
)

var condReconfigToAddModR16ExtScpacConfigCompleteR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CondReconfigToAddMod_r16_Ext_Scpac_ConfigComplete_r18_True},
}

type CondReconfigToAddMod_r16 struct {
	CondReconfigId_r16          CondReconfigId_r16
	CondExecutionCond_r16       []MeasId
	CondRRCReconfig_r16         []byte
	CondExecutionCondSCG_r17    []byte
	CondExecutionCondPSCell_r18 []MeasId
	SubsequentCondReconfig_r18  *SubsequentCondReconfig_r18
	SecurityCellSetId_r18       *SecurityCellSetId_r18
	Scpac_ConfigComplete_r18    *int64
}

func (ie *CondReconfigToAddMod_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(condReconfigToAddModR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.CondExecutionCondSCG_r17 != nil
	hasExtGroup1 := ie.CondExecutionCondPSCell_r18 != nil || ie.SubsequentCondReconfig_r18 != nil || ie.SecurityCellSetId_r18 != nil || ie.Scpac_ConfigComplete_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CondExecutionCond_r16 != nil, ie.CondRRCReconfig_r16 != nil}); err != nil {
		return err
	}

	// 3. condReconfigId-r16: ref
	{
		if err := ie.CondReconfigId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. condExecutionCond-r16: sequence-of
	{
		if ie.CondExecutionCond_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(condReconfigToAddModR16CondExecutionCondR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.CondExecutionCond_r16))); err != nil {
				return err
			}
			for i := range ie.CondExecutionCond_r16 {
				if err := ie.CondExecutionCond_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. condRRCReconfig-r16: octet-string
	{
		if ie.CondRRCReconfig_r16 != nil {
			if err := e.EncodeOctetString(ie.CondRRCReconfig_r16, condReconfigToAddModR16CondRRCReconfigR16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "condExecutionCondSCG-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CondExecutionCondSCG_r17 != nil}); err != nil {
				return err
			}

			if ie.CondExecutionCondSCG_r17 != nil {
				if err := ex.EncodeOctetString(ie.CondExecutionCondSCG_r17, condReconfigToAddModR16CondExecutionCondSCGR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "condExecutionCondPSCell-r18", Optional: true},
					{Name: "subsequentCondReconfig-r18", Optional: true},
					{Name: "securityCellSetId-r18", Optional: true},
					{Name: "scpac-ConfigComplete-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CondExecutionCondPSCell_r18 != nil, ie.SubsequentCondReconfig_r18 != nil, ie.SecurityCellSetId_r18 != nil, ie.Scpac_ConfigComplete_r18 != nil}); err != nil {
				return err
			}

			if ie.CondExecutionCondPSCell_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(condReconfigToAddModR16ExtCondExecutionCondPSCellR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.CondExecutionCondPSCell_r18))); err != nil {
					return err
				}
				for i := range ie.CondExecutionCondPSCell_r18 {
					if err := ie.CondExecutionCondPSCell_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SubsequentCondReconfig_r18 != nil {
				if err := ie.SubsequentCondReconfig_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SecurityCellSetId_r18 != nil {
				if err := ie.SecurityCellSetId_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Scpac_ConfigComplete_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Scpac_ConfigComplete_r18, condReconfigToAddModR16ExtScpacConfigCompleteR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CondReconfigToAddMod_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(condReconfigToAddModR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. condReconfigId-r16: ref
	{
		if err := ie.CondReconfigId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. condExecutionCond-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(condReconfigToAddModR16CondExecutionCondR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CondExecutionCond_r16 = make([]MeasId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CondExecutionCond_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. condRRCReconfig-r16: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(condReconfigToAddModR16CondRRCReconfigR16Constraints)
			if err != nil {
				return err
			}
			ie.CondRRCReconfig_r16 = v
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "condExecutionCondSCG-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeOctetString(condReconfigToAddModR16CondExecutionCondSCGR17Constraints)
			if err != nil {
				return err
			}
			ie.CondExecutionCondSCG_r17 = v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "condExecutionCondPSCell-r18", Optional: true},
				{Name: "subsequentCondReconfig-r18", Optional: true},
				{Name: "securityCellSetId-r18", Optional: true},
				{Name: "scpac-ConfigComplete-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(condReconfigToAddModR16ExtCondExecutionCondPSCellR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CondExecutionCondPSCell_r18 = make([]MeasId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CondExecutionCondPSCell_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SubsequentCondReconfig_r18 = new(SubsequentCondReconfig_r18)
			if err := ie.SubsequentCondReconfig_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SecurityCellSetId_r18 = new(SecurityCellSetId_r18)
			if err := ie.SecurityCellSetId_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(condReconfigToAddModR16ExtScpacConfigCompleteR18Constraints)
			if err != nil {
				return err
			}
			ie.Scpac_ConfigComplete_r18 = &v
		}
	}

	return nil
}
