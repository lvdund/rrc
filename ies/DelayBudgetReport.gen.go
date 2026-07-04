// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DelayBudgetReport (line 2403).

var delayBudgetReportConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type1"},
	},
}

const (
	DelayBudgetReport_Type1 = 0
)

const (
	DelayBudgetReport_Type1_MsMinus1280 = 0
	DelayBudgetReport_Type1_MsMinus640  = 1
	DelayBudgetReport_Type1_MsMinus320  = 2
	DelayBudgetReport_Type1_MsMinus160  = 3
	DelayBudgetReport_Type1_MsMinus80   = 4
	DelayBudgetReport_Type1_MsMinus60   = 5
	DelayBudgetReport_Type1_MsMinus40   = 6
	DelayBudgetReport_Type1_MsMinus20   = 7
	DelayBudgetReport_Type1_Ms0         = 8
	DelayBudgetReport_Type1_Ms20        = 9
	DelayBudgetReport_Type1_Ms40        = 10
	DelayBudgetReport_Type1_Ms60        = 11
	DelayBudgetReport_Type1_Ms80        = 12
	DelayBudgetReport_Type1_Ms160       = 13
	DelayBudgetReport_Type1_Ms320       = 14
	DelayBudgetReport_Type1_Ms640       = 15
	DelayBudgetReport_Type1_Ms1280      = 16
)

var delayBudgetReportType1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DelayBudgetReport_Type1_MsMinus1280, DelayBudgetReport_Type1_MsMinus640, DelayBudgetReport_Type1_MsMinus320, DelayBudgetReport_Type1_MsMinus160, DelayBudgetReport_Type1_MsMinus80, DelayBudgetReport_Type1_MsMinus60, DelayBudgetReport_Type1_MsMinus40, DelayBudgetReport_Type1_MsMinus20, DelayBudgetReport_Type1_Ms0, DelayBudgetReport_Type1_Ms20, DelayBudgetReport_Type1_Ms40, DelayBudgetReport_Type1_Ms60, DelayBudgetReport_Type1_Ms80, DelayBudgetReport_Type1_Ms160, DelayBudgetReport_Type1_Ms320, DelayBudgetReport_Type1_Ms640, DelayBudgetReport_Type1_Ms1280},
}

type DelayBudgetReport struct {
	Choice int
	Type1  *int64
}

func (ie *DelayBudgetReport) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(delayBudgetReportConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case DelayBudgetReport_Type1:
		if err := e.EncodeEnumerated((*ie.Type1), delayBudgetReportType1Constraints); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "DelayBudgetReport",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *DelayBudgetReport) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(delayBudgetReportConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "DelayBudgetReport", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case DelayBudgetReport_Type1:
		ie.Type1 = new(int64)
		v, err := d.DecodeEnumerated(delayBudgetReportType1Constraints)
		if err != nil {
			return err
		}
		(*ie.Type1) = v
	default:
		return &per.DecodeError{TypeName: "DelayBudgetReport", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
