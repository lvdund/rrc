// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EntryCondition-r19 (line 8047).

var entryConditionR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "entryEvaluationOnMR-ForLR-OnLPSS-r19", Optional: true},
		{Name: "entryEvaluationOnMR-ForLR-OnSSB-r19", Optional: true},
		{Name: "entryEvaluationOnLR-ForLR-OnSSB-r19", Optional: true},
		{Name: "entryEvaluationOnLR-ForLR-OnLPSS-r19", Optional: true},
	},
}

var entryConditionR19EntryEvaluationOnMRForLROnLPSSR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "thresholdP1-r19"},
		{Name: "thresholdQ1-r19", Optional: true},
	},
}

var entryConditionR19EntryEvaluationOnMRForLROnSSBR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "thresholdP2-r19"},
		{Name: "thresholdQ2-r19", Optional: true},
	},
}

var entryConditionR19EntryEvaluationOnLRForLROnSSBR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "thresholdP1-LR-r19"},
		{Name: "thresholdQ1-LR-r19", Optional: true},
	},
}

var entryConditionR19EntryEvaluationOnLRForLROnLPSSR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "thresholdP3-LR-r19"},
		{Name: "thresholdQ3-LR-r19", Optional: true},
	},
}

type EntryCondition_r19 struct {
	EntryEvaluationOnMR_ForLR_OnLPSS_r19 *struct {
		ThresholdP1_r19 ReselectionThreshold
		ThresholdQ1_r19 *ReselectionThresholdQ
	}
	EntryEvaluationOnMR_ForLR_OnSSB_r19 *struct {
		ThresholdP2_r19 ReselectionThreshold
		ThresholdQ2_r19 *ReselectionThresholdQ
	}
	EntryEvaluationOnLR_ForLR_OnSSB_r19 *struct {
		ThresholdP1_LR_r19 ThresholdP_LR_r19
		ThresholdQ1_LR_r19 *ThresholdQ_LR_r19
	}
	EntryEvaluationOnLR_ForLR_OnLPSS_r19 *struct {
		ThresholdP3_LR_r19 ThresholdP_LR_r19
		ThresholdQ3_LR_r19 *ThresholdQ_LR_r19
	}
}

func (ie *EntryCondition_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(entryConditionR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.EntryEvaluationOnMR_ForLR_OnLPSS_r19 != nil, ie.EntryEvaluationOnMR_ForLR_OnSSB_r19 != nil, ie.EntryEvaluationOnLR_ForLR_OnSSB_r19 != nil, ie.EntryEvaluationOnLR_ForLR_OnLPSS_r19 != nil}); err != nil {
		return err
	}

	// 3. entryEvaluationOnMR-ForLR-OnLPSS-r19: sequence
	{
		if ie.EntryEvaluationOnMR_ForLR_OnLPSS_r19 != nil {
			c := ie.EntryEvaluationOnMR_ForLR_OnLPSS_r19
			entryConditionR19EntryEvaluationOnMRForLROnLPSSR19Seq := e.NewSequenceEncoder(entryConditionR19EntryEvaluationOnMRForLROnLPSSR19Constraints)
			if err := entryConditionR19EntryEvaluationOnMRForLROnLPSSR19Seq.EncodePreamble([]bool{c.ThresholdQ1_r19 != nil}); err != nil {
				return err
			}
			if err := c.ThresholdP1_r19.Encode(e); err != nil {
				return err
			}
			if c.ThresholdQ1_r19 != nil {
				if err := c.ThresholdQ1_r19.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. entryEvaluationOnMR-ForLR-OnSSB-r19: sequence
	{
		if ie.EntryEvaluationOnMR_ForLR_OnSSB_r19 != nil {
			c := ie.EntryEvaluationOnMR_ForLR_OnSSB_r19
			entryConditionR19EntryEvaluationOnMRForLROnSSBR19Seq := e.NewSequenceEncoder(entryConditionR19EntryEvaluationOnMRForLROnSSBR19Constraints)
			if err := entryConditionR19EntryEvaluationOnMRForLROnSSBR19Seq.EncodePreamble([]bool{c.ThresholdQ2_r19 != nil}); err != nil {
				return err
			}
			if err := c.ThresholdP2_r19.Encode(e); err != nil {
				return err
			}
			if c.ThresholdQ2_r19 != nil {
				if err := c.ThresholdQ2_r19.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. entryEvaluationOnLR-ForLR-OnSSB-r19: sequence
	{
		if ie.EntryEvaluationOnLR_ForLR_OnSSB_r19 != nil {
			c := ie.EntryEvaluationOnLR_ForLR_OnSSB_r19
			entryConditionR19EntryEvaluationOnLRForLROnSSBR19Seq := e.NewSequenceEncoder(entryConditionR19EntryEvaluationOnLRForLROnSSBR19Constraints)
			if err := entryConditionR19EntryEvaluationOnLRForLROnSSBR19Seq.EncodePreamble([]bool{c.ThresholdQ1_LR_r19 != nil}); err != nil {
				return err
			}
			if err := c.ThresholdP1_LR_r19.Encode(e); err != nil {
				return err
			}
			if c.ThresholdQ1_LR_r19 != nil {
				if err := c.ThresholdQ1_LR_r19.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. entryEvaluationOnLR-ForLR-OnLPSS-r19: sequence
	{
		if ie.EntryEvaluationOnLR_ForLR_OnLPSS_r19 != nil {
			c := ie.EntryEvaluationOnLR_ForLR_OnLPSS_r19
			entryConditionR19EntryEvaluationOnLRForLROnLPSSR19Seq := e.NewSequenceEncoder(entryConditionR19EntryEvaluationOnLRForLROnLPSSR19Constraints)
			if err := entryConditionR19EntryEvaluationOnLRForLROnLPSSR19Seq.EncodePreamble([]bool{c.ThresholdQ3_LR_r19 != nil}); err != nil {
				return err
			}
			if err := c.ThresholdP3_LR_r19.Encode(e); err != nil {
				return err
			}
			if c.ThresholdQ3_LR_r19 != nil {
				if err := c.ThresholdQ3_LR_r19.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *EntryCondition_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(entryConditionR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. entryEvaluationOnMR-ForLR-OnLPSS-r19: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.EntryEvaluationOnMR_ForLR_OnLPSS_r19 = &struct {
				ThresholdP1_r19 ReselectionThreshold
				ThresholdQ1_r19 *ReselectionThresholdQ
			}{}
			c := ie.EntryEvaluationOnMR_ForLR_OnLPSS_r19
			entryConditionR19EntryEvaluationOnMRForLROnLPSSR19Seq := d.NewSequenceDecoder(entryConditionR19EntryEvaluationOnMRForLROnLPSSR19Constraints)
			if err := entryConditionR19EntryEvaluationOnMRForLROnLPSSR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.ThresholdP1_r19.Decode(d); err != nil {
					return err
				}
			}
			if entryConditionR19EntryEvaluationOnMRForLROnLPSSR19Seq.IsComponentPresent(1) {
				c.ThresholdQ1_r19 = new(ReselectionThresholdQ)
				if err := (*c.ThresholdQ1_r19).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. entryEvaluationOnMR-ForLR-OnSSB-r19: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.EntryEvaluationOnMR_ForLR_OnSSB_r19 = &struct {
				ThresholdP2_r19 ReselectionThreshold
				ThresholdQ2_r19 *ReselectionThresholdQ
			}{}
			c := ie.EntryEvaluationOnMR_ForLR_OnSSB_r19
			entryConditionR19EntryEvaluationOnMRForLROnSSBR19Seq := d.NewSequenceDecoder(entryConditionR19EntryEvaluationOnMRForLROnSSBR19Constraints)
			if err := entryConditionR19EntryEvaluationOnMRForLROnSSBR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.ThresholdP2_r19.Decode(d); err != nil {
					return err
				}
			}
			if entryConditionR19EntryEvaluationOnMRForLROnSSBR19Seq.IsComponentPresent(1) {
				c.ThresholdQ2_r19 = new(ReselectionThresholdQ)
				if err := (*c.ThresholdQ2_r19).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. entryEvaluationOnLR-ForLR-OnSSB-r19: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.EntryEvaluationOnLR_ForLR_OnSSB_r19 = &struct {
				ThresholdP1_LR_r19 ThresholdP_LR_r19
				ThresholdQ1_LR_r19 *ThresholdQ_LR_r19
			}{}
			c := ie.EntryEvaluationOnLR_ForLR_OnSSB_r19
			entryConditionR19EntryEvaluationOnLRForLROnSSBR19Seq := d.NewSequenceDecoder(entryConditionR19EntryEvaluationOnLRForLROnSSBR19Constraints)
			if err := entryConditionR19EntryEvaluationOnLRForLROnSSBR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.ThresholdP1_LR_r19.Decode(d); err != nil {
					return err
				}
			}
			if entryConditionR19EntryEvaluationOnLRForLROnSSBR19Seq.IsComponentPresent(1) {
				c.ThresholdQ1_LR_r19 = new(ThresholdQ_LR_r19)
				if err := (*c.ThresholdQ1_LR_r19).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. entryEvaluationOnLR-ForLR-OnLPSS-r19: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.EntryEvaluationOnLR_ForLR_OnLPSS_r19 = &struct {
				ThresholdP3_LR_r19 ThresholdP_LR_r19
				ThresholdQ3_LR_r19 *ThresholdQ_LR_r19
			}{}
			c := ie.EntryEvaluationOnLR_ForLR_OnLPSS_r19
			entryConditionR19EntryEvaluationOnLRForLROnLPSSR19Seq := d.NewSequenceDecoder(entryConditionR19EntryEvaluationOnLRForLROnLPSSR19Constraints)
			if err := entryConditionR19EntryEvaluationOnLRForLROnLPSSR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.ThresholdP3_LR_r19.Decode(d); err != nil {
					return err
				}
			}
			if entryConditionR19EntryEvaluationOnLRForLROnLPSSR19Seq.IsComponentPresent(1) {
				c.ThresholdQ3_LR_r19 = new(ThresholdQ_LR_r19)
				if err := (*c.ThresholdQ3_LR_r19).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
