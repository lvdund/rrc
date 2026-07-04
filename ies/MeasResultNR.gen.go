// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultNR (line 9762).

var measResultNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId", Optional: true},
		{Name: "measResult"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

const (
	MeasResultNR_Ext_ChoCandidate_r17_True = 0
)

var measResultNRExtChoCandidateR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasResultNR_Ext_ChoCandidate_r17_True},
}

var measResultNRExtChoConfigR17Constraints = per.SizeRange(1, 2)

const (
	MeasResultNR_Ext_Entering_r18_True = 0
)

var measResultNRExtEnteringR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasResultNR_Ext_Entering_r18_True},
}

var measResultNRDistanceFromReference2R19Constraints = per.Constrained(0, 65535)

var measResultNRMeasResultConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cellResults"},
		{Name: "rsIndexResults", Optional: true},
	},
}

var measResultNRMeasResultCellResultsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resultsSSB-Cell", Optional: true},
		{Name: "resultsCSI-RS-Cell", Optional: true},
	},
}

var measResultNRMeasResultRsIndexResultsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resultsSSB-Indexes", Optional: true},
		{Name: "resultsCSI-RS-Indexes", Optional: true},
	},
}

var measResultNRExtTriggeredEventR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "timeBetweenEvents-r17", Optional: true},
		{Name: "firstTriggeredEvent-r17", Optional: true},
	},
}

const (
	MeasResultNR_Ext_TriggeredEvent_r17_FirstTriggeredEvent_r17_CondFirstEvent  = 0
	MeasResultNR_Ext_TriggeredEvent_r17_FirstTriggeredEvent_r17_CondSecondEvent = 1
)

var measResultNRExtTriggeredEventR17FirstTriggeredEventR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasResultNR_Ext_TriggeredEvent_r17_FirstTriggeredEvent_r17_CondFirstEvent, MeasResultNR_Ext_TriggeredEvent_r17_FirstTriggeredEvent_r17_CondSecondEvent},
}

type MeasResultNR struct {
	PhysCellId *PhysCellId
	MeasResult struct {
		CellResults struct {
			ResultsSSB_Cell    *MeasQuantityResults
			ResultsCSI_RS_Cell *MeasQuantityResults
		}
		RsIndexResults *struct {
			ResultsSSB_Indexes    *ResultsPerSSB_IndexList
			ResultsCSI_RS_Indexes *ResultsPerCSI_RS_IndexList
		}
	}
	Cgi_Info           *CGI_InfoNR
	ChoCandidate_r17   *int64
	ChoConfig_r17      []CondTriggerConfig_r16
	TriggeredEvent_r17 *struct {
		TimeBetweenEvents_r17   *TimeBetweenEvent_r17
		FirstTriggeredEvent_r17 *int64
	}
	Entering_r18               *int64
	DistanceFromReference2_r19 *int64
}

func (ie *MeasResultNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultNRConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Cgi_Info != nil
	hasExtGroup1 := ie.ChoCandidate_r17 != nil || ie.ChoConfig_r17 != nil || ie.TriggeredEvent_r17 != nil
	hasExtGroup2 := ie.Entering_r18 != nil
	hasExtGroup3 := ie.DistanceFromReference2_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PhysCellId != nil}); err != nil {
		return err
	}

	// 3. physCellId: ref
	{
		if ie.PhysCellId != nil {
			if err := ie.PhysCellId.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. measResult: sequence
	{
		{
			c := &ie.MeasResult
			measResultNRMeasResultSeq := e.NewSequenceEncoder(measResultNRMeasResultConstraints)
			if err := measResultNRMeasResultSeq.EncodePreamble([]bool{c.RsIndexResults != nil}); err != nil {
				return err
			}
			{
				c := &c.CellResults
				measResultNRMeasResultCellResultsSeq := e.NewSequenceEncoder(measResultNRMeasResultCellResultsConstraints)
				if err := measResultNRMeasResultCellResultsSeq.EncodePreamble([]bool{c.ResultsSSB_Cell != nil, c.ResultsCSI_RS_Cell != nil}); err != nil {
					return err
				}
				if c.ResultsSSB_Cell != nil {
					if err := c.ResultsSSB_Cell.Encode(e); err != nil {
						return err
					}
				}
				if c.ResultsCSI_RS_Cell != nil {
					if err := c.ResultsCSI_RS_Cell.Encode(e); err != nil {
						return err
					}
				}
			}
			if c.RsIndexResults != nil {
				c := (*c.RsIndexResults)
				measResultNRMeasResultRsIndexResultsSeq := e.NewSequenceEncoder(measResultNRMeasResultRsIndexResultsConstraints)
				if err := measResultNRMeasResultRsIndexResultsSeq.EncodePreamble([]bool{c.ResultsSSB_Indexes != nil, c.ResultsCSI_RS_Indexes != nil}); err != nil {
					return err
				}
				if c.ResultsSSB_Indexes != nil {
					if err := c.ResultsSSB_Indexes.Encode(e); err != nil {
						return err
					}
				}
				if c.ResultsCSI_RS_Indexes != nil {
					if err := c.ResultsCSI_RS_Indexes.Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "cgi-Info", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cgi_Info != nil}); err != nil {
				return err
			}

			if ie.Cgi_Info != nil {
				if err := ie.Cgi_Info.Encode(ex); err != nil {
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
					{Name: "choCandidate-r17", Optional: true},
					{Name: "choConfig-r17", Optional: true},
					{Name: "triggeredEvent-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ChoCandidate_r17 != nil, ie.ChoConfig_r17 != nil, ie.TriggeredEvent_r17 != nil}); err != nil {
				return err
			}

			if ie.ChoCandidate_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ChoCandidate_r17, measResultNRExtChoCandidateR17Constraints); err != nil {
					return err
				}
			}

			if ie.ChoConfig_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(measResultNRExtChoConfigR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.ChoConfig_r17))); err != nil {
					return err
				}
				for i := range ie.ChoConfig_r17 {
					if err := ie.ChoConfig_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.TriggeredEvent_r17 != nil {
				c := ie.TriggeredEvent_r17
				measResultNRExtTriggeredEventR17Seq := ex.NewSequenceEncoder(measResultNRExtTriggeredEventR17Constraints)
				if err := measResultNRExtTriggeredEventR17Seq.EncodePreamble([]bool{c.TimeBetweenEvents_r17 != nil, c.FirstTriggeredEvent_r17 != nil}); err != nil {
					return err
				}
				if c.TimeBetweenEvents_r17 != nil {
					if err := c.TimeBetweenEvents_r17.Encode(ex); err != nil {
						return err
					}
				}
				if c.FirstTriggeredEvent_r17 != nil {
					if err := ex.EncodeEnumerated((*c.FirstTriggeredEvent_r17), measResultNRExtTriggeredEventR17FirstTriggeredEventR17Constraints); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "entering-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Entering_r18 != nil}); err != nil {
				return err
			}

			if ie.Entering_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Entering_r18, measResultNRExtEnteringR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "distanceFromReference2-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DistanceFromReference2_r19 != nil}); err != nil {
				return err
			}

			if ie.DistanceFromReference2_r19 != nil {
				if err := ex.EncodeInteger(*ie.DistanceFromReference2_r19, measResultNRDistanceFromReference2R19Constraints); err != nil {
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

func (ie *MeasResultNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultNRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. physCellId: ref
	{
		if seq.IsComponentPresent(0) {
			ie.PhysCellId = new(PhysCellId)
			if err := ie.PhysCellId.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. measResult: sequence
	{
		{
			c := &ie.MeasResult
			measResultNRMeasResultSeq := d.NewSequenceDecoder(measResultNRMeasResultConstraints)
			if err := measResultNRMeasResultSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				c := &c.CellResults
				measResultNRMeasResultCellResultsSeq := d.NewSequenceDecoder(measResultNRMeasResultCellResultsConstraints)
				if err := measResultNRMeasResultCellResultsSeq.DecodePreamble(); err != nil {
					return err
				}
				if measResultNRMeasResultCellResultsSeq.IsComponentPresent(0) {
					c.ResultsSSB_Cell = new(MeasQuantityResults)
					if err := (*c.ResultsSSB_Cell).Decode(d); err != nil {
						return err
					}
				}
				if measResultNRMeasResultCellResultsSeq.IsComponentPresent(1) {
					c.ResultsCSI_RS_Cell = new(MeasQuantityResults)
					if err := (*c.ResultsCSI_RS_Cell).Decode(d); err != nil {
						return err
					}
				}
			}
			if measResultNRMeasResultSeq.IsComponentPresent(1) {
				c.RsIndexResults = &struct {
					ResultsSSB_Indexes    *ResultsPerSSB_IndexList
					ResultsCSI_RS_Indexes *ResultsPerCSI_RS_IndexList
				}{}
				c := (*c.RsIndexResults)
				measResultNRMeasResultRsIndexResultsSeq := d.NewSequenceDecoder(measResultNRMeasResultRsIndexResultsConstraints)
				if err := measResultNRMeasResultRsIndexResultsSeq.DecodePreamble(); err != nil {
					return err
				}
				if measResultNRMeasResultRsIndexResultsSeq.IsComponentPresent(0) {
					c.ResultsSSB_Indexes = new(ResultsPerSSB_IndexList)
					if err := (*c.ResultsSSB_Indexes).Decode(d); err != nil {
						return err
					}
				}
				if measResultNRMeasResultRsIndexResultsSeq.IsComponentPresent(1) {
					c.ResultsCSI_RS_Indexes = new(ResultsPerCSI_RS_IndexList)
					if err := (*c.ResultsCSI_RS_Indexes).Decode(d); err != nil {
						return err
					}
				}
			}
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
				{Name: "cgi-Info", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Cgi_Info = new(CGI_InfoNR)
			if err := ie.Cgi_Info.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "choCandidate-r17", Optional: true},
				{Name: "choConfig-r17", Optional: true},
				{Name: "triggeredEvent-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measResultNRExtChoCandidateR17Constraints)
			if err != nil {
				return err
			}
			ie.ChoCandidate_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(measResultNRExtChoConfigR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ChoConfig_r17 = make([]CondTriggerConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ChoConfig_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.TriggeredEvent_r17 = &struct {
				TimeBetweenEvents_r17   *TimeBetweenEvent_r17
				FirstTriggeredEvent_r17 *int64
			}{}
			c := ie.TriggeredEvent_r17
			measResultNRExtTriggeredEventR17Seq := dx.NewSequenceDecoder(measResultNRExtTriggeredEventR17Constraints)
			if err := measResultNRExtTriggeredEventR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if measResultNRExtTriggeredEventR17Seq.IsComponentPresent(0) {
				c.TimeBetweenEvents_r17 = new(TimeBetweenEvent_r17)
				if err := (*c.TimeBetweenEvents_r17).Decode(dx); err != nil {
					return err
				}
			}
			if measResultNRExtTriggeredEventR17Seq.IsComponentPresent(1) {
				c.FirstTriggeredEvent_r17 = new(int64)
				v, err := dx.DecodeEnumerated(measResultNRExtTriggeredEventR17FirstTriggeredEventR17Constraints)
				if err != nil {
					return err
				}
				(*c.FirstTriggeredEvent_r17) = v
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "entering-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measResultNRExtEnteringR18Constraints)
			if err != nil {
				return err
			}
			ie.Entering_r18 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "distanceFromReference2-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(measResultNRDistanceFromReference2R19Constraints)
			if err != nil {
				return err
			}
			ie.DistanceFromReference2_r19 = &v
		}
	}

	return nil
}
