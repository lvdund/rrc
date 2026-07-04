// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DataCollectionCandidateConfigParameters-r19 (line 26606).

var dataCollectionCandidateConfigParametersR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dataCollectionCandidateConfigId-r19"},
		{Name: "parameters-r19"},
	},
}

var dataCollectionCandidateConfigParameters_r19ParametersR19Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "parametersForBM-r19"},
		{Name: "parametersForCSI-InferencePrediction-r19"},
	},
}

const (
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19                      = 0
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForCSI_InferencePrediction_r19 = 1
)

var dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourcesForChannelMeasurement-r19", Optional: true},
		{Name: "resourcesForChannelPrediction-r19", Optional: true},
		{Name: "associatedIdForChannelMeasurement-r19", Optional: true},
		{Name: "associatedIdForChannelPrediction-r19", Optional: true},
		{Name: "timeGap-r19", Optional: true},
		{Name: "nrofTimeInstance-r19", Optional: true},
	},
}

const (
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Ms10   = 0
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Ms20   = 1
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Ms40   = 2
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Ms80   = 3
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Ms160  = 4
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Spare3 = 5
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Spare2 = 6
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Spare1 = 7
)

var dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19TimeGapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Ms10, DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Ms20, DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Ms40, DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Ms80, DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Ms160, DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Spare3, DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Spare2, DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_TimeGap_r19_Spare1},
}

const (
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_NrofTimeInstance_r19_N1 = 0
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_NrofTimeInstance_r19_N2 = 1
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_NrofTimeInstance_r19_N4 = 2
	DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_NrofTimeInstance_r19_N8 = 3
)

var dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19NrofTimeInstanceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_NrofTimeInstance_r19_N1, DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_NrofTimeInstance_r19_N2, DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_NrofTimeInstance_r19_N4, DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19_NrofTimeInstance_r19_N8},
}

var dataCollectionCandidateConfigParametersR19ParametersR19ParametersForCSIInferencePredictionR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourcesForChannelMeasurement-r19", Optional: true},
		{Name: "codebookConfig-r19", Optional: true},
	},
}

type DataCollectionCandidateConfigParameters_r19 struct {
	DataCollectionCandidateConfigId_r19 DataCollectionCandidateConfigId_r19
	Parameters_r19                      struct {
		Choice              int
		ParametersForBM_r19 *struct {
			ResourcesForChannelMeasurement_r19    *CSI_ResourceConfigId
			ResourcesForChannelPrediction_r19     *CSI_ResourceConfigId
			AssociatedIdForChannelMeasurement_r19 *AssociatedId_r19
			AssociatedIdForChannelPrediction_r19  *AssociatedId_r19
			TimeGap_r19                           *int64
			NrofTimeInstance_r19                  *int64
		}
		ParametersForCSI_InferencePrediction_r19 *struct {
			ResourcesForChannelMeasurement_r19 *CSI_ResourceConfigId
			CodebookConfig_r19                 *CodebookConfig_r18
		}
	}
}

func (ie *DataCollectionCandidateConfigParameters_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataCollectionCandidateConfigParametersR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. dataCollectionCandidateConfigId-r19: ref
	{
		if err := ie.DataCollectionCandidateConfigId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 3. parameters-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(dataCollectionCandidateConfigParameters_r19ParametersR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Parameters_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Parameters_r19.Choice {
		case DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19:
			c := (*ie.Parameters_r19.ParametersForBM_r19)
			dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Seq := e.NewSequenceEncoder(dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Constraints)
			if err := dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Seq.EncodePreamble([]bool{c.ResourcesForChannelMeasurement_r19 != nil, c.ResourcesForChannelPrediction_r19 != nil, c.AssociatedIdForChannelMeasurement_r19 != nil, c.AssociatedIdForChannelPrediction_r19 != nil, c.TimeGap_r19 != nil, c.NrofTimeInstance_r19 != nil}); err != nil {
				return err
			}
			if c.ResourcesForChannelMeasurement_r19 != nil {
				if err := c.ResourcesForChannelMeasurement_r19.Encode(e); err != nil {
					return err
				}
			}
			if c.ResourcesForChannelPrediction_r19 != nil {
				if err := c.ResourcesForChannelPrediction_r19.Encode(e); err != nil {
					return err
				}
			}
			if c.AssociatedIdForChannelMeasurement_r19 != nil {
				if err := c.AssociatedIdForChannelMeasurement_r19.Encode(e); err != nil {
					return err
				}
			}
			if c.AssociatedIdForChannelPrediction_r19 != nil {
				if err := c.AssociatedIdForChannelPrediction_r19.Encode(e); err != nil {
					return err
				}
			}
			if c.TimeGap_r19 != nil {
				if err := e.EncodeEnumerated((*c.TimeGap_r19), dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19TimeGapR19Constraints); err != nil {
					return err
				}
			}
			if c.NrofTimeInstance_r19 != nil {
				if err := e.EncodeEnumerated((*c.NrofTimeInstance_r19), dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19NrofTimeInstanceR19Constraints); err != nil {
					return err
				}
			}
		case DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForCSI_InferencePrediction_r19:
			c := (*ie.Parameters_r19.ParametersForCSI_InferencePrediction_r19)
			dataCollectionCandidateConfigParametersR19ParametersR19ParametersForCSIInferencePredictionR19Seq := e.NewSequenceEncoder(dataCollectionCandidateConfigParametersR19ParametersR19ParametersForCSIInferencePredictionR19Constraints)
			if err := dataCollectionCandidateConfigParametersR19ParametersR19ParametersForCSIInferencePredictionR19Seq.EncodePreamble([]bool{c.ResourcesForChannelMeasurement_r19 != nil, c.CodebookConfig_r19 != nil}); err != nil {
				return err
			}
			if c.ResourcesForChannelMeasurement_r19 != nil {
				if err := c.ResourcesForChannelMeasurement_r19.Encode(e); err != nil {
					return err
				}
			}
			if c.CodebookConfig_r19 != nil {
				if err := c.CodebookConfig_r19.Encode(e); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Parameters_r19.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *DataCollectionCandidateConfigParameters_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataCollectionCandidateConfigParametersR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. dataCollectionCandidateConfigId-r19: ref
	{
		if err := ie.DataCollectionCandidateConfigId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 3. parameters-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(dataCollectionCandidateConfigParameters_r19ParametersR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Parameters_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForBM_r19:
			ie.Parameters_r19.ParametersForBM_r19 = &struct {
				ResourcesForChannelMeasurement_r19    *CSI_ResourceConfigId
				ResourcesForChannelPrediction_r19     *CSI_ResourceConfigId
				AssociatedIdForChannelMeasurement_r19 *AssociatedId_r19
				AssociatedIdForChannelPrediction_r19  *AssociatedId_r19
				TimeGap_r19                           *int64
				NrofTimeInstance_r19                  *int64
			}{}
			c := (*ie.Parameters_r19.ParametersForBM_r19)
			dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Seq := d.NewSequenceDecoder(dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Constraints)
			if err := dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Seq.IsComponentPresent(0) {
				c.ResourcesForChannelMeasurement_r19 = new(CSI_ResourceConfigId)
				if err := (*c.ResourcesForChannelMeasurement_r19).Decode(d); err != nil {
					return err
				}
			}
			if dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Seq.IsComponentPresent(1) {
				c.ResourcesForChannelPrediction_r19 = new(CSI_ResourceConfigId)
				if err := (*c.ResourcesForChannelPrediction_r19).Decode(d); err != nil {
					return err
				}
			}
			if dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Seq.IsComponentPresent(2) {
				c.AssociatedIdForChannelMeasurement_r19 = new(AssociatedId_r19)
				if err := (*c.AssociatedIdForChannelMeasurement_r19).Decode(d); err != nil {
					return err
				}
			}
			if dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Seq.IsComponentPresent(3) {
				c.AssociatedIdForChannelPrediction_r19 = new(AssociatedId_r19)
				if err := (*c.AssociatedIdForChannelPrediction_r19).Decode(d); err != nil {
					return err
				}
			}
			if dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Seq.IsComponentPresent(4) {
				c.TimeGap_r19 = new(int64)
				v, err := d.DecodeEnumerated(dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19TimeGapR19Constraints)
				if err != nil {
					return err
				}
				(*c.TimeGap_r19) = v
			}
			if dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19Seq.IsComponentPresent(5) {
				c.NrofTimeInstance_r19 = new(int64)
				v, err := d.DecodeEnumerated(dataCollectionCandidateConfigParametersR19ParametersR19ParametersForBMR19NrofTimeInstanceR19Constraints)
				if err != nil {
					return err
				}
				(*c.NrofTimeInstance_r19) = v
			}
		case DataCollectionCandidateConfigParameters_r19_Parameters_r19_ParametersForCSI_InferencePrediction_r19:
			ie.Parameters_r19.ParametersForCSI_InferencePrediction_r19 = &struct {
				ResourcesForChannelMeasurement_r19 *CSI_ResourceConfigId
				CodebookConfig_r19                 *CodebookConfig_r18
			}{}
			c := (*ie.Parameters_r19.ParametersForCSI_InferencePrediction_r19)
			dataCollectionCandidateConfigParametersR19ParametersR19ParametersForCSIInferencePredictionR19Seq := d.NewSequenceDecoder(dataCollectionCandidateConfigParametersR19ParametersR19ParametersForCSIInferencePredictionR19Constraints)
			if err := dataCollectionCandidateConfigParametersR19ParametersR19ParametersForCSIInferencePredictionR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if dataCollectionCandidateConfigParametersR19ParametersR19ParametersForCSIInferencePredictionR19Seq.IsComponentPresent(0) {
				c.ResourcesForChannelMeasurement_r19 = new(CSI_ResourceConfigId)
				if err := (*c.ResourcesForChannelMeasurement_r19).Decode(d); err != nil {
					return err
				}
			}
			if dataCollectionCandidateConfigParametersR19ParametersR19ParametersForCSIInferencePredictionR19Seq.IsComponentPresent(1) {
				c.CodebookConfig_r19 = new(CodebookConfig_r18)
				if err := (*c.CodebookConfig_r19).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
