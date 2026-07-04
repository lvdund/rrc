// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-ResourceConfig (line 7463).

var cSIResourceConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-ResourceConfigId"},
		{Name: "csi-RS-ResourceSetList"},
		{Name: "bwp-Id"},
		{Name: "resourceType"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var cSI_ResourceConfigCsiRSResourceSetListConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nzp-CSI-RS-SSB"},
		{Name: "csi-IM-ResourceSetList"},
	},
}

const (
	CSI_ResourceConfig_Csi_RS_ResourceSetList_Nzp_CSI_RS_SSB         = 0
	CSI_ResourceConfig_Csi_RS_ResourceSetList_Csi_IM_ResourceSetList = 1
)

const (
	CSI_ResourceConfig_ResourceType_Aperiodic      = 0
	CSI_ResourceConfig_ResourceType_SemiPersistent = 1
	CSI_ResourceConfig_ResourceType_Periodic       = 2
)

var cSIResourceConfigResourceTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ResourceConfig_ResourceType_Aperiodic, CSI_ResourceConfig_ResourceType_SemiPersistent, CSI_ResourceConfig_ResourceType_Periodic},
}

var cSIResourceConfigExtCliMeasResourceSetListR19Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cli-RSSI-MeasResourceSetList-r19"},
		{Name: "srs-RSRP-MeasResourceSetList-r19"},
	},
}

const (
	CSI_ResourceConfig_Ext_Cli_MeasResourceSetList_r19_Cli_RSSI_MeasResourceSetList_r19 = 0
	CSI_ResourceConfig_Ext_Cli_MeasResourceSetList_r19_Srs_RSRP_MeasResourceSetList_r19 = 1
)

var cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nzp-CSI-RS-ResourceSetList", Optional: true},
		{Name: "csi-SSB-ResourceSetList", Optional: true},
	},
}

var cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBNzpCSIRSResourceSetListConstraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourceSetsPerConfig)

var cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBCsiSSBResourceSetListConstraints = per.SizeRange(1, common.MaxNrofCSI_SSB_ResourceSetsPerConfig)

var cSIResourceConfigCsiRSResourceSetListCsiIMResourceSetListConstraints = per.SizeRange(1, common.MaxNrofCSI_IM_ResourceSetsPerConfig)

var cSIResourceConfigExtCliMeasResourceSetListR19CliRSSIMeasResourceSetListR19Constraints = per.SizeRange(1, common.MaxNrofCLI_RSSI_MeasResourceSetsPerConfig_r19)

var cSIResourceConfigExtCliMeasResourceSetListR19SrsRSRPMeasResourceSetListR19Constraints = per.SizeRange(1, common.MaxNrofSRS_RSRP_MeasResourceSetsPerConfig_r19)

type CSI_ResourceConfig struct {
	Csi_ResourceConfigId   CSI_ResourceConfigId
	Csi_RS_ResourceSetList struct {
		Choice         int
		Nzp_CSI_RS_SSB *struct {
			Nzp_CSI_RS_ResourceSetList []NZP_CSI_RS_ResourceSetId
			Csi_SSB_ResourceSetList    []CSI_SSB_ResourceSetId
		}
		Csi_IM_ResourceSetList *[]CSI_IM_ResourceSetId
	}
	Bwp_Id                         BWP_Id
	ResourceType                   int64
	Csi_SSB_ResourceSetListExt_r17 *CSI_SSB_ResourceSetId
	Cli_MeasResourceSetList_r19    *struct {
		Choice                           int
		Cli_RSSI_MeasResourceSetList_r19 *[]CLI_RSSI_MeasResourceSetId_r19
		Srs_RSRP_MeasResourceSetList_r19 *[]SRS_RSRP_MeasResourceSetId_r19
	}
}

func (ie *CSI_ResourceConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIResourceConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Csi_SSB_ResourceSetListExt_r17 != nil
	hasExtGroup1 := ie.Cli_MeasResourceSetList_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. csi-ResourceConfigId: ref
	{
		if err := ie.Csi_ResourceConfigId.Encode(e); err != nil {
			return err
		}
	}

	// 3. csi-RS-ResourceSetList: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_ResourceConfigCsiRSResourceSetListConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Csi_RS_ResourceSetList.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Csi_RS_ResourceSetList.Choice {
		case CSI_ResourceConfig_Csi_RS_ResourceSetList_Nzp_CSI_RS_SSB:
			c := (*ie.Csi_RS_ResourceSetList.Nzp_CSI_RS_SSB)
			cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBSeq := e.NewSequenceEncoder(cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBConstraints)
			if err := cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBSeq.EncodePreamble([]bool{c.Nzp_CSI_RS_ResourceSetList != nil, c.Csi_SSB_ResourceSetList != nil}); err != nil {
				return err
			}
			if c.Nzp_CSI_RS_ResourceSetList != nil {
				seqOf := e.NewSequenceOfEncoder(cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBNzpCSIRSResourceSetListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.Nzp_CSI_RS_ResourceSetList))); err != nil {
					return err
				}
				for i := range c.Nzp_CSI_RS_ResourceSetList {
					if err := c.Nzp_CSI_RS_ResourceSetList[i].Encode(e); err != nil {
						return err
					}
				}
			}
			if c.Csi_SSB_ResourceSetList != nil {
				seqOf := e.NewSequenceOfEncoder(cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBCsiSSBResourceSetListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.Csi_SSB_ResourceSetList))); err != nil {
					return err
				}
				for i := range c.Csi_SSB_ResourceSetList {
					if err := c.Csi_SSB_ResourceSetList[i].Encode(e); err != nil {
						return err
					}
				}
			}
		case CSI_ResourceConfig_Csi_RS_ResourceSetList_Csi_IM_ResourceSetList:
			seqOf := e.NewSequenceOfEncoder(cSIResourceConfigCsiRSResourceSetListCsiIMResourceSetListConstraints)
			if err := seqOf.EncodeLength(int64(len((*ie.Csi_RS_ResourceSetList.Csi_IM_ResourceSetList)))); err != nil {
				return err
			}
			for i := range *ie.Csi_RS_ResourceSetList.Csi_IM_ResourceSetList {
				if err := (*ie.Csi_RS_ResourceSetList.Csi_IM_ResourceSetList)[i].Encode(e); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Csi_RS_ResourceSetList.Choice), Constraint: "index out of range"}
		}
	}

	// 4. bwp-Id: ref
	{
		if err := ie.Bwp_Id.Encode(e); err != nil {
			return err
		}
	}

	// 5. resourceType: enumerated
	{
		if err := e.EncodeEnumerated(ie.ResourceType, cSIResourceConfigResourceTypeConstraints); err != nil {
			return err
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
					{Name: "csi-SSB-ResourceSetListExt-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Csi_SSB_ResourceSetListExt_r17 != nil}); err != nil {
				return err
			}

			if ie.Csi_SSB_ResourceSetListExt_r17 != nil {
				if err := ie.Csi_SSB_ResourceSetListExt_r17.Encode(ex); err != nil {
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
					{Name: "cli-MeasResourceSetList-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cli_MeasResourceSetList_r19 != nil}); err != nil {
				return err
			}

			if ie.Cli_MeasResourceSetList_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(cSIResourceConfigExtCliMeasResourceSetListR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Cli_MeasResourceSetList_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Cli_MeasResourceSetList_r19).Choice {
				case CSI_ResourceConfig_Ext_Cli_MeasResourceSetList_r19_Cli_RSSI_MeasResourceSetList_r19:
					seqOf := ex.NewSequenceOfEncoder(cSIResourceConfigExtCliMeasResourceSetListR19CliRSSIMeasResourceSetListR19Constraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.Cli_MeasResourceSetList_r19).Cli_RSSI_MeasResourceSetList_r19)))); err != nil {
						return err
					}
					for i := range *(*ie.Cli_MeasResourceSetList_r19).Cli_RSSI_MeasResourceSetList_r19 {
						if err := (*(*ie.Cli_MeasResourceSetList_r19).Cli_RSSI_MeasResourceSetList_r19)[i].Encode(ex); err != nil {
							return err
						}
					}
				case CSI_ResourceConfig_Ext_Cli_MeasResourceSetList_r19_Srs_RSRP_MeasResourceSetList_r19:
					seqOf := ex.NewSequenceOfEncoder(cSIResourceConfigExtCliMeasResourceSetListR19SrsRSRPMeasResourceSetListR19Constraints)
					if err := seqOf.EncodeLength(int64(len((*(*ie.Cli_MeasResourceSetList_r19).Srs_RSRP_MeasResourceSetList_r19)))); err != nil {
						return err
					}
					for i := range *(*ie.Cli_MeasResourceSetList_r19).Srs_RSRP_MeasResourceSetList_r19 {
						if err := (*(*ie.Cli_MeasResourceSetList_r19).Srs_RSRP_MeasResourceSetList_r19)[i].Encode(ex); err != nil {
							return err
						}
					}
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

func (ie *CSI_ResourceConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIResourceConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. csi-ResourceConfigId: ref
	{
		if err := ie.Csi_ResourceConfigId.Decode(d); err != nil {
			return err
		}
	}

	// 3. csi-RS-ResourceSetList: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_ResourceConfigCsiRSResourceSetListConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Csi_RS_ResourceSetList.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_ResourceConfig_Csi_RS_ResourceSetList_Nzp_CSI_RS_SSB:
			ie.Csi_RS_ResourceSetList.Nzp_CSI_RS_SSB = &struct {
				Nzp_CSI_RS_ResourceSetList []NZP_CSI_RS_ResourceSetId
				Csi_SSB_ResourceSetList    []CSI_SSB_ResourceSetId
			}{}
			c := (*ie.Csi_RS_ResourceSetList.Nzp_CSI_RS_SSB)
			cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBSeq := d.NewSequenceDecoder(cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBConstraints)
			if err := cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBSeq.DecodePreamble(); err != nil {
				return err
			}
			if cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBSeq.IsComponentPresent(0) {
				seqOf := d.NewSequenceOfDecoder(cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBNzpCSIRSResourceSetListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Nzp_CSI_RS_ResourceSetList = make([]NZP_CSI_RS_ResourceSetId, n)
				for i := int64(0); i < n; i++ {
					if err := c.Nzp_CSI_RS_ResourceSetList[i].Decode(d); err != nil {
						return err
					}
				}
			}
			if cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBSeq.IsComponentPresent(1) {
				seqOf := d.NewSequenceOfDecoder(cSIResourceConfigCsiRSResourceSetListNzpCSIRSSSBCsiSSBResourceSetListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Csi_SSB_ResourceSetList = make([]CSI_SSB_ResourceSetId, n)
				for i := int64(0); i < n; i++ {
					if err := c.Csi_SSB_ResourceSetList[i].Decode(d); err != nil {
						return err
					}
				}
			}
		case CSI_ResourceConfig_Csi_RS_ResourceSetList_Csi_IM_ResourceSetList:
			ie.Csi_RS_ResourceSetList.Csi_IM_ResourceSetList = new([]CSI_IM_ResourceSetId)
			seqOf := d.NewSequenceOfDecoder(cSIResourceConfigCsiRSResourceSetListCsiIMResourceSetListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*ie.Csi_RS_ResourceSetList.Csi_IM_ResourceSetList) = make([]CSI_IM_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := (*ie.Csi_RS_ResourceSetList.Csi_IM_ResourceSetList)[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. bwp-Id: ref
	{
		if err := ie.Bwp_Id.Decode(d); err != nil {
			return err
		}
	}

	// 5. resourceType: enumerated
	{
		v3, err := d.DecodeEnumerated(cSIResourceConfigResourceTypeConstraints)
		if err != nil {
			return err
		}
		ie.ResourceType = v3
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
				{Name: "csi-SSB-ResourceSetListExt-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Csi_SSB_ResourceSetListExt_r17 = new(CSI_SSB_ResourceSetId)
			if err := ie.Csi_SSB_ResourceSetListExt_r17.Decode(dx); err != nil {
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
				{Name: "cli-MeasResourceSetList-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Cli_MeasResourceSetList_r19 = &struct {
				Choice                           int
				Cli_RSSI_MeasResourceSetList_r19 *[]CLI_RSSI_MeasResourceSetId_r19
				Srs_RSRP_MeasResourceSetList_r19 *[]SRS_RSRP_MeasResourceSetId_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(cSIResourceConfigExtCliMeasResourceSetListR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Cli_MeasResourceSetList_r19).Choice = int(index)
			switch index {
			case CSI_ResourceConfig_Ext_Cli_MeasResourceSetList_r19_Cli_RSSI_MeasResourceSetList_r19:
				(*ie.Cli_MeasResourceSetList_r19).Cli_RSSI_MeasResourceSetList_r19 = new([]CLI_RSSI_MeasResourceSetId_r19)
				seqOf := dx.NewSequenceOfDecoder(cSIResourceConfigExtCliMeasResourceSetListR19CliRSSIMeasResourceSetListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.Cli_MeasResourceSetList_r19).Cli_RSSI_MeasResourceSetList_r19) = make([]CLI_RSSI_MeasResourceSetId_r19, n)
				for i := int64(0); i < n; i++ {
					if err := (*(*ie.Cli_MeasResourceSetList_r19).Cli_RSSI_MeasResourceSetList_r19)[i].Decode(dx); err != nil {
						return err
					}
				}
			case CSI_ResourceConfig_Ext_Cli_MeasResourceSetList_r19_Srs_RSRP_MeasResourceSetList_r19:
				(*ie.Cli_MeasResourceSetList_r19).Srs_RSRP_MeasResourceSetList_r19 = new([]SRS_RSRP_MeasResourceSetId_r19)
				seqOf := dx.NewSequenceOfDecoder(cSIResourceConfigExtCliMeasResourceSetListR19SrsRSRPMeasResourceSetListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.Cli_MeasResourceSetList_r19).Srs_RSRP_MeasResourceSetList_r19) = make([]SRS_RSRP_MeasResourceSetId_r19, n)
				for i := int64(0); i < n; i++ {
					if err := (*(*ie.Cli_MeasResourceSetList_r19).Srs_RSRP_MeasResourceSetList_r19)[i].Decode(dx); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
