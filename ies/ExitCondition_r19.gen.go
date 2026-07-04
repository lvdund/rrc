// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ExitCondition-r19 (line 8067).

var exitConditionR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "exitEvaluationOnLR-ForLR-OnLPSS-r19", Optional: true},
		{Name: "exitEvaluationOnLR-ForLR-OnSSB-r19", Optional: true},
	},
}

var exitConditionR19ExitEvaluationOnLRForLROnLPSSR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "thresholdP2-LR-r19"},
		{Name: "thresholdQ2-LR-r19", Optional: true},
	},
}

var exitConditionR19ExitEvaluationOnLRForLROnSSBR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "thresholdP4-LR-r19"},
		{Name: "thresholdQ4-LR-r19", Optional: true},
	},
}

type ExitCondition_r19 struct {
	ExitEvaluationOnLR_ForLR_OnLPSS_r19 *struct {
		ThresholdP2_LR_r19 ThresholdP_LR_r19
		ThresholdQ2_LR_r19 *ThresholdQ_LR_r19
	}
	ExitEvaluationOnLR_ForLR_OnSSB_r19 *struct {
		ThresholdP4_LR_r19 ThresholdP_LR_r19
		ThresholdQ4_LR_r19 *ThresholdQ_LR_r19
	}
}

func (ie *ExitCondition_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(exitConditionR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ExitEvaluationOnLR_ForLR_OnLPSS_r19 != nil, ie.ExitEvaluationOnLR_ForLR_OnSSB_r19 != nil}); err != nil {
		return err
	}

	// 3. exitEvaluationOnLR-ForLR-OnLPSS-r19: sequence
	{
		if ie.ExitEvaluationOnLR_ForLR_OnLPSS_r19 != nil {
			c := ie.ExitEvaluationOnLR_ForLR_OnLPSS_r19
			exitConditionR19ExitEvaluationOnLRForLROnLPSSR19Seq := e.NewSequenceEncoder(exitConditionR19ExitEvaluationOnLRForLROnLPSSR19Constraints)
			if err := exitConditionR19ExitEvaluationOnLRForLROnLPSSR19Seq.EncodePreamble([]bool{c.ThresholdQ2_LR_r19 != nil}); err != nil {
				return err
			}
			if err := c.ThresholdP2_LR_r19.Encode(e); err != nil {
				return err
			}
			if c.ThresholdQ2_LR_r19 != nil {
				if err := c.ThresholdQ2_LR_r19.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. exitEvaluationOnLR-ForLR-OnSSB-r19: sequence
	{
		if ie.ExitEvaluationOnLR_ForLR_OnSSB_r19 != nil {
			c := ie.ExitEvaluationOnLR_ForLR_OnSSB_r19
			exitConditionR19ExitEvaluationOnLRForLROnSSBR19Seq := e.NewSequenceEncoder(exitConditionR19ExitEvaluationOnLRForLROnSSBR19Constraints)
			if err := exitConditionR19ExitEvaluationOnLRForLROnSSBR19Seq.EncodePreamble([]bool{c.ThresholdQ4_LR_r19 != nil}); err != nil {
				return err
			}
			if err := c.ThresholdP4_LR_r19.Encode(e); err != nil {
				return err
			}
			if c.ThresholdQ4_LR_r19 != nil {
				if err := c.ThresholdQ4_LR_r19.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *ExitCondition_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(exitConditionR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. exitEvaluationOnLR-ForLR-OnLPSS-r19: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.ExitEvaluationOnLR_ForLR_OnLPSS_r19 = &struct {
				ThresholdP2_LR_r19 ThresholdP_LR_r19
				ThresholdQ2_LR_r19 *ThresholdQ_LR_r19
			}{}
			c := ie.ExitEvaluationOnLR_ForLR_OnLPSS_r19
			exitConditionR19ExitEvaluationOnLRForLROnLPSSR19Seq := d.NewSequenceDecoder(exitConditionR19ExitEvaluationOnLRForLROnLPSSR19Constraints)
			if err := exitConditionR19ExitEvaluationOnLRForLROnLPSSR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.ThresholdP2_LR_r19.Decode(d); err != nil {
					return err
				}
			}
			if exitConditionR19ExitEvaluationOnLRForLROnLPSSR19Seq.IsComponentPresent(1) {
				c.ThresholdQ2_LR_r19 = new(ThresholdQ_LR_r19)
				if err := (*c.ThresholdQ2_LR_r19).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. exitEvaluationOnLR-ForLR-OnSSB-r19: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.ExitEvaluationOnLR_ForLR_OnSSB_r19 = &struct {
				ThresholdP4_LR_r19 ThresholdP_LR_r19
				ThresholdQ4_LR_r19 *ThresholdQ_LR_r19
			}{}
			c := ie.ExitEvaluationOnLR_ForLR_OnSSB_r19
			exitConditionR19ExitEvaluationOnLRForLROnSSBR19Seq := d.NewSequenceDecoder(exitConditionR19ExitEvaluationOnLRForLROnSSBR19Constraints)
			if err := exitConditionR19ExitEvaluationOnLRForLROnSSBR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.ThresholdP4_LR_r19.Decode(d); err != nil {
					return err
				}
			}
			if exitConditionR19ExitEvaluationOnLRForLROnSSBR19Seq.IsComponentPresent(1) {
				c.ThresholdQ4_LR_r19 = new(ThresholdQ_LR_r19)
				if err := (*c.ThresholdQ4_LR_r19).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
