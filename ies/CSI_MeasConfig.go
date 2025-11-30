package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_MeasConfig struct {
	Nzp_CSI_RS_ResourceToAddModList           []NZP_CSI_RS_Resource                       `lb:1,ub:maxNrofNZP_CSI_RS_Resources,optional`
	Nzp_CSI_RS_ResourceToReleaseList          []NZP_CSI_RS_ResourceId                     `lb:1,ub:maxNrofNZP_CSI_RS_Resources,optional`
	Nzp_CSI_RS_ResourceSetToAddModList        []NZP_CSI_RS_ResourceSet                    `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSets,optional`
	Nzp_CSI_RS_ResourceSetToReleaseList       []NZP_CSI_RS_ResourceSetId                  `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSets,optional`
	Csi_IM_ResourceToAddModList               []CSI_IM_Resource                           `lb:1,ub:maxNrofCSI_IM_Resources,optional`
	Csi_IM_ResourceToReleaseList              []CSI_IM_ResourceId                         `lb:1,ub:maxNrofCSI_IM_Resources,optional`
	Csi_IM_ResourceSetToAddModList            []CSI_IM_ResourceSet                        `lb:1,ub:maxNrofCSI_IM_ResourceSets,optional`
	Csi_IM_ResourceSetToReleaseList           []CSI_IM_ResourceSetId                      `lb:1,ub:maxNrofCSI_IM_ResourceSets,optional`
	Csi_SSB_ResourceSetToAddModList           []CSI_SSB_ResourceSet                       `lb:1,ub:maxNrofCSI_SSB_ResourceSets,optional`
	Csi_SSB_ResourceSetToReleaseList          []CSI_SSB_ResourceSetId                     `lb:1,ub:maxNrofCSI_SSB_ResourceSets,optional`
	Csi_ResourceConfigToAddModList            []CSI_ResourceConfig                        `lb:1,ub:maxNrofCSI_ResourceConfigurations,optional`
	Csi_ResourceConfigToReleaseList           []CSI_ResourceConfigId                      `lb:1,ub:maxNrofCSI_ResourceConfigurations,optional`
	Csi_ReportConfigToAddModList              []CSI_ReportConfig                          `lb:1,ub:maxNrofCSI_ReportConfigurations,optional`
	Csi_ReportConfigToReleaseList             []CSI_ReportConfigId                        `lb:1,ub:maxNrofCSI_ReportConfigurations,optional`
	ReportTriggerSize                         *int64                                      `lb:0,ub:6,optional`
	AperiodicTriggerStateList                 *CSI_AperiodicTriggerStateList              `optional,setuprelease`
	SemiPersistentOnPUSCH_TriggerStateList    *CSI_SemiPersistentOnPUSCH_TriggerStateList `optional,setuprelease`
	ReportTriggerSizeDCI_0_2_r16              *int64                                      `lb:0,ub:6,optional,ext-1`
	SCellActivationRS_ConfigToAddModList_r17  []SCellActivationRS_Config_r17              `lb:1,ub:maxNrofSCellActRS_r17,optional,ext-2`
	SCellActivationRS_ConfigToReleaseList_r17 []SCellActivationRS_ConfigId_r17            `lb:1,ub:maxNrofSCellActRS_r17,optional,ext-2`
}

func (ie *CSI_MeasConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.ReportTriggerSizeDCI_0_2_r16 != nil || len(ie.SCellActivationRS_ConfigToAddModList_r17) > 0 || len(ie.SCellActivationRS_ConfigToReleaseList_r17) > 0
	preambleBits := []bool{hasExtensions, len(ie.Nzp_CSI_RS_ResourceToAddModList) > 0, len(ie.Nzp_CSI_RS_ResourceToReleaseList) > 0, len(ie.Nzp_CSI_RS_ResourceSetToAddModList) > 0, len(ie.Nzp_CSI_RS_ResourceSetToReleaseList) > 0, len(ie.Csi_IM_ResourceToAddModList) > 0, len(ie.Csi_IM_ResourceToReleaseList) > 0, len(ie.Csi_IM_ResourceSetToAddModList) > 0, len(ie.Csi_IM_ResourceSetToReleaseList) > 0, len(ie.Csi_SSB_ResourceSetToAddModList) > 0, len(ie.Csi_SSB_ResourceSetToReleaseList) > 0, len(ie.Csi_ResourceConfigToAddModList) > 0, len(ie.Csi_ResourceConfigToReleaseList) > 0, len(ie.Csi_ReportConfigToAddModList) > 0, len(ie.Csi_ReportConfigToReleaseList) > 0, ie.ReportTriggerSize != nil, ie.AperiodicTriggerStateList != nil, ie.SemiPersistentOnPUSCH_TriggerStateList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Nzp_CSI_RS_ResourceToAddModList) > 0 {
		tmp_Nzp_CSI_RS_ResourceToAddModList := utils.NewSequence[*NZP_CSI_RS_Resource]([]*NZP_CSI_RS_Resource{}, aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_Resources}, false)
		for _, i := range ie.Nzp_CSI_RS_ResourceToAddModList {
			tmp_Nzp_CSI_RS_ResourceToAddModList.Value = append(tmp_Nzp_CSI_RS_ResourceToAddModList.Value, &i)
		}
		if err = tmp_Nzp_CSI_RS_ResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Nzp_CSI_RS_ResourceToAddModList", err)
		}
	}
	if len(ie.Nzp_CSI_RS_ResourceToReleaseList) > 0 {
		tmp_Nzp_CSI_RS_ResourceToReleaseList := utils.NewSequence[*NZP_CSI_RS_ResourceId]([]*NZP_CSI_RS_ResourceId{}, aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_Resources}, false)
		for _, i := range ie.Nzp_CSI_RS_ResourceToReleaseList {
			tmp_Nzp_CSI_RS_ResourceToReleaseList.Value = append(tmp_Nzp_CSI_RS_ResourceToReleaseList.Value, &i)
		}
		if err = tmp_Nzp_CSI_RS_ResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Nzp_CSI_RS_ResourceToReleaseList", err)
		}
	}
	if len(ie.Nzp_CSI_RS_ResourceSetToAddModList) > 0 {
		tmp_Nzp_CSI_RS_ResourceSetToAddModList := utils.NewSequence[*NZP_CSI_RS_ResourceSet]([]*NZP_CSI_RS_ResourceSet{}, aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.Nzp_CSI_RS_ResourceSetToAddModList {
			tmp_Nzp_CSI_RS_ResourceSetToAddModList.Value = append(tmp_Nzp_CSI_RS_ResourceSetToAddModList.Value, &i)
		}
		if err = tmp_Nzp_CSI_RS_ResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Nzp_CSI_RS_ResourceSetToAddModList", err)
		}
	}
	if len(ie.Nzp_CSI_RS_ResourceSetToReleaseList) > 0 {
		tmp_Nzp_CSI_RS_ResourceSetToReleaseList := utils.NewSequence[*NZP_CSI_RS_ResourceSetId]([]*NZP_CSI_RS_ResourceSetId{}, aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.Nzp_CSI_RS_ResourceSetToReleaseList {
			tmp_Nzp_CSI_RS_ResourceSetToReleaseList.Value = append(tmp_Nzp_CSI_RS_ResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_Nzp_CSI_RS_ResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Nzp_CSI_RS_ResourceSetToReleaseList", err)
		}
	}
	if len(ie.Csi_IM_ResourceToAddModList) > 0 {
		tmp_Csi_IM_ResourceToAddModList := utils.NewSequence[*CSI_IM_Resource]([]*CSI_IM_Resource{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_Resources}, false)
		for _, i := range ie.Csi_IM_ResourceToAddModList {
			tmp_Csi_IM_ResourceToAddModList.Value = append(tmp_Csi_IM_ResourceToAddModList.Value, &i)
		}
		if err = tmp_Csi_IM_ResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_IM_ResourceToAddModList", err)
		}
	}
	if len(ie.Csi_IM_ResourceToReleaseList) > 0 {
		tmp_Csi_IM_ResourceToReleaseList := utils.NewSequence[*CSI_IM_ResourceId]([]*CSI_IM_ResourceId{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_Resources}, false)
		for _, i := range ie.Csi_IM_ResourceToReleaseList {
			tmp_Csi_IM_ResourceToReleaseList.Value = append(tmp_Csi_IM_ResourceToReleaseList.Value, &i)
		}
		if err = tmp_Csi_IM_ResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_IM_ResourceToReleaseList", err)
		}
	}
	if len(ie.Csi_IM_ResourceSetToAddModList) > 0 {
		tmp_Csi_IM_ResourceSetToAddModList := utils.NewSequence[*CSI_IM_ResourceSet]([]*CSI_IM_ResourceSet{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSets}, false)
		for _, i := range ie.Csi_IM_ResourceSetToAddModList {
			tmp_Csi_IM_ResourceSetToAddModList.Value = append(tmp_Csi_IM_ResourceSetToAddModList.Value, &i)
		}
		if err = tmp_Csi_IM_ResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_IM_ResourceSetToAddModList", err)
		}
	}
	if len(ie.Csi_IM_ResourceSetToReleaseList) > 0 {
		tmp_Csi_IM_ResourceSetToReleaseList := utils.NewSequence[*CSI_IM_ResourceSetId]([]*CSI_IM_ResourceSetId{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSets}, false)
		for _, i := range ie.Csi_IM_ResourceSetToReleaseList {
			tmp_Csi_IM_ResourceSetToReleaseList.Value = append(tmp_Csi_IM_ResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_Csi_IM_ResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_IM_ResourceSetToReleaseList", err)
		}
	}
	if len(ie.Csi_SSB_ResourceSetToAddModList) > 0 {
		tmp_Csi_SSB_ResourceSetToAddModList := utils.NewSequence[*CSI_SSB_ResourceSet]([]*CSI_SSB_ResourceSet{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSets}, false)
		for _, i := range ie.Csi_SSB_ResourceSetToAddModList {
			tmp_Csi_SSB_ResourceSetToAddModList.Value = append(tmp_Csi_SSB_ResourceSetToAddModList.Value, &i)
		}
		if err = tmp_Csi_SSB_ResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_SSB_ResourceSetToAddModList", err)
		}
	}
	if len(ie.Csi_SSB_ResourceSetToReleaseList) > 0 {
		tmp_Csi_SSB_ResourceSetToReleaseList := utils.NewSequence[*CSI_SSB_ResourceSetId]([]*CSI_SSB_ResourceSetId{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSets}, false)
		for _, i := range ie.Csi_SSB_ResourceSetToReleaseList {
			tmp_Csi_SSB_ResourceSetToReleaseList.Value = append(tmp_Csi_SSB_ResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_Csi_SSB_ResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_SSB_ResourceSetToReleaseList", err)
		}
	}
	if len(ie.Csi_ResourceConfigToAddModList) > 0 {
		tmp_Csi_ResourceConfigToAddModList := utils.NewSequence[*CSI_ResourceConfig]([]*CSI_ResourceConfig{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_ResourceConfigurations}, false)
		for _, i := range ie.Csi_ResourceConfigToAddModList {
			tmp_Csi_ResourceConfigToAddModList.Value = append(tmp_Csi_ResourceConfigToAddModList.Value, &i)
		}
		if err = tmp_Csi_ResourceConfigToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_ResourceConfigToAddModList", err)
		}
	}
	if len(ie.Csi_ResourceConfigToReleaseList) > 0 {
		tmp_Csi_ResourceConfigToReleaseList := utils.NewSequence[*CSI_ResourceConfigId]([]*CSI_ResourceConfigId{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_ResourceConfigurations}, false)
		for _, i := range ie.Csi_ResourceConfigToReleaseList {
			tmp_Csi_ResourceConfigToReleaseList.Value = append(tmp_Csi_ResourceConfigToReleaseList.Value, &i)
		}
		if err = tmp_Csi_ResourceConfigToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_ResourceConfigToReleaseList", err)
		}
	}
	if len(ie.Csi_ReportConfigToAddModList) > 0 {
		tmp_Csi_ReportConfigToAddModList := utils.NewSequence[*CSI_ReportConfig]([]*CSI_ReportConfig{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_ReportConfigurations}, false)
		for _, i := range ie.Csi_ReportConfigToAddModList {
			tmp_Csi_ReportConfigToAddModList.Value = append(tmp_Csi_ReportConfigToAddModList.Value, &i)
		}
		if err = tmp_Csi_ReportConfigToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_ReportConfigToAddModList", err)
		}
	}
	if len(ie.Csi_ReportConfigToReleaseList) > 0 {
		tmp_Csi_ReportConfigToReleaseList := utils.NewSequence[*CSI_ReportConfigId]([]*CSI_ReportConfigId{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_ReportConfigurations}, false)
		for _, i := range ie.Csi_ReportConfigToReleaseList {
			tmp_Csi_ReportConfigToReleaseList.Value = append(tmp_Csi_ReportConfigToReleaseList.Value, &i)
		}
		if err = tmp_Csi_ReportConfigToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_ReportConfigToReleaseList", err)
		}
	}
	if ie.ReportTriggerSize != nil {
		if err = w.WriteInteger(*ie.ReportTriggerSize, &aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode ReportTriggerSize", err)
		}
	}
	if ie.AperiodicTriggerStateList != nil {
		tmp_AperiodicTriggerStateList := utils.SetupRelease[*CSI_AperiodicTriggerStateList]{
			Setup: ie.AperiodicTriggerStateList,
		}
		if err = tmp_AperiodicTriggerStateList.Encode(w); err != nil {
			return utils.WrapError("Encode AperiodicTriggerStateList", err)
		}
	}
	if ie.SemiPersistentOnPUSCH_TriggerStateList != nil {
		tmp_SemiPersistentOnPUSCH_TriggerStateList := utils.SetupRelease[*CSI_SemiPersistentOnPUSCH_TriggerStateList]{
			Setup: ie.SemiPersistentOnPUSCH_TriggerStateList,
		}
		if err = tmp_SemiPersistentOnPUSCH_TriggerStateList.Encode(w); err != nil {
			return utils.WrapError("Encode SemiPersistentOnPUSCH_TriggerStateList", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.ReportTriggerSizeDCI_0_2_r16 != nil, len(ie.SCellActivationRS_ConfigToAddModList_r17) > 0 || len(ie.SCellActivationRS_ConfigToReleaseList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_MeasConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ReportTriggerSizeDCI_0_2_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ReportTriggerSizeDCI_0_2_r16 optional
			if ie.ReportTriggerSizeDCI_0_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.ReportTriggerSizeDCI_0_2_r16, &aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
					return utils.WrapError("Encode ReportTriggerSizeDCI_0_2_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{len(ie.SCellActivationRS_ConfigToAddModList_r17) > 0, len(ie.SCellActivationRS_ConfigToReleaseList_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SCellActivationRS_ConfigToAddModList_r17 optional
			if len(ie.SCellActivationRS_ConfigToAddModList_r17) > 0 {
				tmp_SCellActivationRS_ConfigToAddModList_r17 := utils.NewSequence[*SCellActivationRS_Config_r17]([]*SCellActivationRS_Config_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofSCellActRS_r17}, false)
				for _, i := range ie.SCellActivationRS_ConfigToAddModList_r17 {
					tmp_SCellActivationRS_ConfigToAddModList_r17.Value = append(tmp_SCellActivationRS_ConfigToAddModList_r17.Value, &i)
				}
				if err = tmp_SCellActivationRS_ConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SCellActivationRS_ConfigToAddModList_r17", err)
				}
			}
			// encode SCellActivationRS_ConfigToReleaseList_r17 optional
			if len(ie.SCellActivationRS_ConfigToReleaseList_r17) > 0 {
				tmp_SCellActivationRS_ConfigToReleaseList_r17 := utils.NewSequence[*SCellActivationRS_ConfigId_r17]([]*SCellActivationRS_ConfigId_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofSCellActRS_r17}, false)
				for _, i := range ie.SCellActivationRS_ConfigToReleaseList_r17 {
					tmp_SCellActivationRS_ConfigToReleaseList_r17.Value = append(tmp_SCellActivationRS_ConfigToReleaseList_r17.Value, &i)
				}
				if err = tmp_SCellActivationRS_ConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SCellActivationRS_ConfigToReleaseList_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *CSI_MeasConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Nzp_CSI_RS_ResourceToAddModListPresent bool
	if Nzp_CSI_RS_ResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Nzp_CSI_RS_ResourceToReleaseListPresent bool
	if Nzp_CSI_RS_ResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Nzp_CSI_RS_ResourceSetToAddModListPresent bool
	if Nzp_CSI_RS_ResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Nzp_CSI_RS_ResourceSetToReleaseListPresent bool
	if Nzp_CSI_RS_ResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_IM_ResourceToAddModListPresent bool
	if Csi_IM_ResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_IM_ResourceToReleaseListPresent bool
	if Csi_IM_ResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_IM_ResourceSetToAddModListPresent bool
	if Csi_IM_ResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_IM_ResourceSetToReleaseListPresent bool
	if Csi_IM_ResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_SSB_ResourceSetToAddModListPresent bool
	if Csi_SSB_ResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_SSB_ResourceSetToReleaseListPresent bool
	if Csi_SSB_ResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_ResourceConfigToAddModListPresent bool
	if Csi_ResourceConfigToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_ResourceConfigToReleaseListPresent bool
	if Csi_ResourceConfigToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_ReportConfigToAddModListPresent bool
	if Csi_ReportConfigToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_ReportConfigToReleaseListPresent bool
	if Csi_ReportConfigToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReportTriggerSizePresent bool
	if ReportTriggerSizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AperiodicTriggerStateListPresent bool
	if AperiodicTriggerStateListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SemiPersistentOnPUSCH_TriggerStateListPresent bool
	if SemiPersistentOnPUSCH_TriggerStateListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Nzp_CSI_RS_ResourceToAddModListPresent {
		tmp_Nzp_CSI_RS_ResourceToAddModList := utils.NewSequence[*NZP_CSI_RS_Resource]([]*NZP_CSI_RS_Resource{}, aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_Resources}, false)
		fn_Nzp_CSI_RS_ResourceToAddModList := func() *NZP_CSI_RS_Resource {
			return new(NZP_CSI_RS_Resource)
		}
		if err = tmp_Nzp_CSI_RS_ResourceToAddModList.Decode(r, fn_Nzp_CSI_RS_ResourceToAddModList); err != nil {
			return utils.WrapError("Decode Nzp_CSI_RS_ResourceToAddModList", err)
		}
		ie.Nzp_CSI_RS_ResourceToAddModList = []NZP_CSI_RS_Resource{}
		for _, i := range tmp_Nzp_CSI_RS_ResourceToAddModList.Value {
			ie.Nzp_CSI_RS_ResourceToAddModList = append(ie.Nzp_CSI_RS_ResourceToAddModList, *i)
		}
	}
	if Nzp_CSI_RS_ResourceToReleaseListPresent {
		tmp_Nzp_CSI_RS_ResourceToReleaseList := utils.NewSequence[*NZP_CSI_RS_ResourceId]([]*NZP_CSI_RS_ResourceId{}, aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_Resources}, false)
		fn_Nzp_CSI_RS_ResourceToReleaseList := func() *NZP_CSI_RS_ResourceId {
			return new(NZP_CSI_RS_ResourceId)
		}
		if err = tmp_Nzp_CSI_RS_ResourceToReleaseList.Decode(r, fn_Nzp_CSI_RS_ResourceToReleaseList); err != nil {
			return utils.WrapError("Decode Nzp_CSI_RS_ResourceToReleaseList", err)
		}
		ie.Nzp_CSI_RS_ResourceToReleaseList = []NZP_CSI_RS_ResourceId{}
		for _, i := range tmp_Nzp_CSI_RS_ResourceToReleaseList.Value {
			ie.Nzp_CSI_RS_ResourceToReleaseList = append(ie.Nzp_CSI_RS_ResourceToReleaseList, *i)
		}
	}
	if Nzp_CSI_RS_ResourceSetToAddModListPresent {
		tmp_Nzp_CSI_RS_ResourceSetToAddModList := utils.NewSequence[*NZP_CSI_RS_ResourceSet]([]*NZP_CSI_RS_ResourceSet{}, aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSets}, false)
		fn_Nzp_CSI_RS_ResourceSetToAddModList := func() *NZP_CSI_RS_ResourceSet {
			return new(NZP_CSI_RS_ResourceSet)
		}
		if err = tmp_Nzp_CSI_RS_ResourceSetToAddModList.Decode(r, fn_Nzp_CSI_RS_ResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode Nzp_CSI_RS_ResourceSetToAddModList", err)
		}
		ie.Nzp_CSI_RS_ResourceSetToAddModList = []NZP_CSI_RS_ResourceSet{}
		for _, i := range tmp_Nzp_CSI_RS_ResourceSetToAddModList.Value {
			ie.Nzp_CSI_RS_ResourceSetToAddModList = append(ie.Nzp_CSI_RS_ResourceSetToAddModList, *i)
		}
	}
	if Nzp_CSI_RS_ResourceSetToReleaseListPresent {
		tmp_Nzp_CSI_RS_ResourceSetToReleaseList := utils.NewSequence[*NZP_CSI_RS_ResourceSetId]([]*NZP_CSI_RS_ResourceSetId{}, aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSets}, false)
		fn_Nzp_CSI_RS_ResourceSetToReleaseList := func() *NZP_CSI_RS_ResourceSetId {
			return new(NZP_CSI_RS_ResourceSetId)
		}
		if err = tmp_Nzp_CSI_RS_ResourceSetToReleaseList.Decode(r, fn_Nzp_CSI_RS_ResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode Nzp_CSI_RS_ResourceSetToReleaseList", err)
		}
		ie.Nzp_CSI_RS_ResourceSetToReleaseList = []NZP_CSI_RS_ResourceSetId{}
		for _, i := range tmp_Nzp_CSI_RS_ResourceSetToReleaseList.Value {
			ie.Nzp_CSI_RS_ResourceSetToReleaseList = append(ie.Nzp_CSI_RS_ResourceSetToReleaseList, *i)
		}
	}
	if Csi_IM_ResourceToAddModListPresent {
		tmp_Csi_IM_ResourceToAddModList := utils.NewSequence[*CSI_IM_Resource]([]*CSI_IM_Resource{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_Resources}, false)
		fn_Csi_IM_ResourceToAddModList := func() *CSI_IM_Resource {
			return new(CSI_IM_Resource)
		}
		if err = tmp_Csi_IM_ResourceToAddModList.Decode(r, fn_Csi_IM_ResourceToAddModList); err != nil {
			return utils.WrapError("Decode Csi_IM_ResourceToAddModList", err)
		}
		ie.Csi_IM_ResourceToAddModList = []CSI_IM_Resource{}
		for _, i := range tmp_Csi_IM_ResourceToAddModList.Value {
			ie.Csi_IM_ResourceToAddModList = append(ie.Csi_IM_ResourceToAddModList, *i)
		}
	}
	if Csi_IM_ResourceToReleaseListPresent {
		tmp_Csi_IM_ResourceToReleaseList := utils.NewSequence[*CSI_IM_ResourceId]([]*CSI_IM_ResourceId{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_Resources}, false)
		fn_Csi_IM_ResourceToReleaseList := func() *CSI_IM_ResourceId {
			return new(CSI_IM_ResourceId)
		}
		if err = tmp_Csi_IM_ResourceToReleaseList.Decode(r, fn_Csi_IM_ResourceToReleaseList); err != nil {
			return utils.WrapError("Decode Csi_IM_ResourceToReleaseList", err)
		}
		ie.Csi_IM_ResourceToReleaseList = []CSI_IM_ResourceId{}
		for _, i := range tmp_Csi_IM_ResourceToReleaseList.Value {
			ie.Csi_IM_ResourceToReleaseList = append(ie.Csi_IM_ResourceToReleaseList, *i)
		}
	}
	if Csi_IM_ResourceSetToAddModListPresent {
		tmp_Csi_IM_ResourceSetToAddModList := utils.NewSequence[*CSI_IM_ResourceSet]([]*CSI_IM_ResourceSet{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSets}, false)
		fn_Csi_IM_ResourceSetToAddModList := func() *CSI_IM_ResourceSet {
			return new(CSI_IM_ResourceSet)
		}
		if err = tmp_Csi_IM_ResourceSetToAddModList.Decode(r, fn_Csi_IM_ResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode Csi_IM_ResourceSetToAddModList", err)
		}
		ie.Csi_IM_ResourceSetToAddModList = []CSI_IM_ResourceSet{}
		for _, i := range tmp_Csi_IM_ResourceSetToAddModList.Value {
			ie.Csi_IM_ResourceSetToAddModList = append(ie.Csi_IM_ResourceSetToAddModList, *i)
		}
	}
	if Csi_IM_ResourceSetToReleaseListPresent {
		tmp_Csi_IM_ResourceSetToReleaseList := utils.NewSequence[*CSI_IM_ResourceSetId]([]*CSI_IM_ResourceSetId{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSets}, false)
		fn_Csi_IM_ResourceSetToReleaseList := func() *CSI_IM_ResourceSetId {
			return new(CSI_IM_ResourceSetId)
		}
		if err = tmp_Csi_IM_ResourceSetToReleaseList.Decode(r, fn_Csi_IM_ResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode Csi_IM_ResourceSetToReleaseList", err)
		}
		ie.Csi_IM_ResourceSetToReleaseList = []CSI_IM_ResourceSetId{}
		for _, i := range tmp_Csi_IM_ResourceSetToReleaseList.Value {
			ie.Csi_IM_ResourceSetToReleaseList = append(ie.Csi_IM_ResourceSetToReleaseList, *i)
		}
	}
	if Csi_SSB_ResourceSetToAddModListPresent {
		tmp_Csi_SSB_ResourceSetToAddModList := utils.NewSequence[*CSI_SSB_ResourceSet]([]*CSI_SSB_ResourceSet{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSets}, false)
		fn_Csi_SSB_ResourceSetToAddModList := func() *CSI_SSB_ResourceSet {
			return new(CSI_SSB_ResourceSet)
		}
		if err = tmp_Csi_SSB_ResourceSetToAddModList.Decode(r, fn_Csi_SSB_ResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode Csi_SSB_ResourceSetToAddModList", err)
		}
		ie.Csi_SSB_ResourceSetToAddModList = []CSI_SSB_ResourceSet{}
		for _, i := range tmp_Csi_SSB_ResourceSetToAddModList.Value {
			ie.Csi_SSB_ResourceSetToAddModList = append(ie.Csi_SSB_ResourceSetToAddModList, *i)
		}
	}
	if Csi_SSB_ResourceSetToReleaseListPresent {
		tmp_Csi_SSB_ResourceSetToReleaseList := utils.NewSequence[*CSI_SSB_ResourceSetId]([]*CSI_SSB_ResourceSetId{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSets}, false)
		fn_Csi_SSB_ResourceSetToReleaseList := func() *CSI_SSB_ResourceSetId {
			return new(CSI_SSB_ResourceSetId)
		}
		if err = tmp_Csi_SSB_ResourceSetToReleaseList.Decode(r, fn_Csi_SSB_ResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode Csi_SSB_ResourceSetToReleaseList", err)
		}
		ie.Csi_SSB_ResourceSetToReleaseList = []CSI_SSB_ResourceSetId{}
		for _, i := range tmp_Csi_SSB_ResourceSetToReleaseList.Value {
			ie.Csi_SSB_ResourceSetToReleaseList = append(ie.Csi_SSB_ResourceSetToReleaseList, *i)
		}
	}
	if Csi_ResourceConfigToAddModListPresent {
		tmp_Csi_ResourceConfigToAddModList := utils.NewSequence[*CSI_ResourceConfig]([]*CSI_ResourceConfig{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_ResourceConfigurations}, false)
		fn_Csi_ResourceConfigToAddModList := func() *CSI_ResourceConfig {
			return new(CSI_ResourceConfig)
		}
		if err = tmp_Csi_ResourceConfigToAddModList.Decode(r, fn_Csi_ResourceConfigToAddModList); err != nil {
			return utils.WrapError("Decode Csi_ResourceConfigToAddModList", err)
		}
		ie.Csi_ResourceConfigToAddModList = []CSI_ResourceConfig{}
		for _, i := range tmp_Csi_ResourceConfigToAddModList.Value {
			ie.Csi_ResourceConfigToAddModList = append(ie.Csi_ResourceConfigToAddModList, *i)
		}
	}
	if Csi_ResourceConfigToReleaseListPresent {
		tmp_Csi_ResourceConfigToReleaseList := utils.NewSequence[*CSI_ResourceConfigId]([]*CSI_ResourceConfigId{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_ResourceConfigurations}, false)
		fn_Csi_ResourceConfigToReleaseList := func() *CSI_ResourceConfigId {
			return new(CSI_ResourceConfigId)
		}
		if err = tmp_Csi_ResourceConfigToReleaseList.Decode(r, fn_Csi_ResourceConfigToReleaseList); err != nil {
			return utils.WrapError("Decode Csi_ResourceConfigToReleaseList", err)
		}
		ie.Csi_ResourceConfigToReleaseList = []CSI_ResourceConfigId{}
		for _, i := range tmp_Csi_ResourceConfigToReleaseList.Value {
			ie.Csi_ResourceConfigToReleaseList = append(ie.Csi_ResourceConfigToReleaseList, *i)
		}
	}
	if Csi_ReportConfigToAddModListPresent {
		tmp_Csi_ReportConfigToAddModList := utils.NewSequence[*CSI_ReportConfig]([]*CSI_ReportConfig{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_ReportConfigurations}, false)
		fn_Csi_ReportConfigToAddModList := func() *CSI_ReportConfig {
			return new(CSI_ReportConfig)
		}
		if err = tmp_Csi_ReportConfigToAddModList.Decode(r, fn_Csi_ReportConfigToAddModList); err != nil {
			return utils.WrapError("Decode Csi_ReportConfigToAddModList", err)
		}
		ie.Csi_ReportConfigToAddModList = []CSI_ReportConfig{}
		for _, i := range tmp_Csi_ReportConfigToAddModList.Value {
			ie.Csi_ReportConfigToAddModList = append(ie.Csi_ReportConfigToAddModList, *i)
		}
	}
	if Csi_ReportConfigToReleaseListPresent {
		tmp_Csi_ReportConfigToReleaseList := utils.NewSequence[*CSI_ReportConfigId]([]*CSI_ReportConfigId{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_ReportConfigurations}, false)
		fn_Csi_ReportConfigToReleaseList := func() *CSI_ReportConfigId {
			return new(CSI_ReportConfigId)
		}
		if err = tmp_Csi_ReportConfigToReleaseList.Decode(r, fn_Csi_ReportConfigToReleaseList); err != nil {
			return utils.WrapError("Decode Csi_ReportConfigToReleaseList", err)
		}
		ie.Csi_ReportConfigToReleaseList = []CSI_ReportConfigId{}
		for _, i := range tmp_Csi_ReportConfigToReleaseList.Value {
			ie.Csi_ReportConfigToReleaseList = append(ie.Csi_ReportConfigToReleaseList, *i)
		}
	}
	if ReportTriggerSizePresent {
		var tmp_int_ReportTriggerSize int64
		if tmp_int_ReportTriggerSize, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode ReportTriggerSize", err)
		}
		ie.ReportTriggerSize = &tmp_int_ReportTriggerSize
	}
	if AperiodicTriggerStateListPresent {
		tmp_AperiodicTriggerStateList := utils.SetupRelease[*CSI_AperiodicTriggerStateList]{}
		if err = tmp_AperiodicTriggerStateList.Decode(r); err != nil {
			return utils.WrapError("Decode AperiodicTriggerStateList", err)
		}
		ie.AperiodicTriggerStateList = tmp_AperiodicTriggerStateList.Setup
	}
	if SemiPersistentOnPUSCH_TriggerStateListPresent {
		tmp_SemiPersistentOnPUSCH_TriggerStateList := utils.SetupRelease[*CSI_SemiPersistentOnPUSCH_TriggerStateList]{}
		if err = tmp_SemiPersistentOnPUSCH_TriggerStateList.Decode(r); err != nil {
			return utils.WrapError("Decode SemiPersistentOnPUSCH_TriggerStateList", err)
		}
		ie.SemiPersistentOnPUSCH_TriggerStateList = tmp_SemiPersistentOnPUSCH_TriggerStateList.Setup
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			ReportTriggerSizeDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ReportTriggerSizeDCI_0_2_r16 optional
			if ReportTriggerSizeDCI_0_2_r16Present {
				var tmp_int_ReportTriggerSizeDCI_0_2_r16 int64
				if tmp_int_ReportTriggerSizeDCI_0_2_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
					return utils.WrapError("Decode ReportTriggerSizeDCI_0_2_r16", err)
				}
				ie.ReportTriggerSizeDCI_0_2_r16 = &tmp_int_ReportTriggerSizeDCI_0_2_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SCellActivationRS_ConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SCellActivationRS_ConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SCellActivationRS_ConfigToAddModList_r17 optional
			if SCellActivationRS_ConfigToAddModList_r17Present {
				tmp_SCellActivationRS_ConfigToAddModList_r17 := utils.NewSequence[*SCellActivationRS_Config_r17]([]*SCellActivationRS_Config_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofSCellActRS_r17}, false)
				fn_SCellActivationRS_ConfigToAddModList_r17 := func() *SCellActivationRS_Config_r17 {
					return new(SCellActivationRS_Config_r17)
				}
				if err = tmp_SCellActivationRS_ConfigToAddModList_r17.Decode(extReader, fn_SCellActivationRS_ConfigToAddModList_r17); err != nil {
					return utils.WrapError("Decode SCellActivationRS_ConfigToAddModList_r17", err)
				}
				ie.SCellActivationRS_ConfigToAddModList_r17 = []SCellActivationRS_Config_r17{}
				for _, i := range tmp_SCellActivationRS_ConfigToAddModList_r17.Value {
					ie.SCellActivationRS_ConfigToAddModList_r17 = append(ie.SCellActivationRS_ConfigToAddModList_r17, *i)
				}
			}
			// decode SCellActivationRS_ConfigToReleaseList_r17 optional
			if SCellActivationRS_ConfigToReleaseList_r17Present {
				tmp_SCellActivationRS_ConfigToReleaseList_r17 := utils.NewSequence[*SCellActivationRS_ConfigId_r17]([]*SCellActivationRS_ConfigId_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofSCellActRS_r17}, false)
				fn_SCellActivationRS_ConfigToReleaseList_r17 := func() *SCellActivationRS_ConfigId_r17 {
					return new(SCellActivationRS_ConfigId_r17)
				}
				if err = tmp_SCellActivationRS_ConfigToReleaseList_r17.Decode(extReader, fn_SCellActivationRS_ConfigToReleaseList_r17); err != nil {
					return utils.WrapError("Decode SCellActivationRS_ConfigToReleaseList_r17", err)
				}
				ie.SCellActivationRS_ConfigToReleaseList_r17 = []SCellActivationRS_ConfigId_r17{}
				for _, i := range tmp_SCellActivationRS_ConfigToReleaseList_r17.Value {
					ie.SCellActivationRS_ConfigToReleaseList_r17 = append(ie.SCellActivationRS_ConfigToReleaseList_r17, *i)
				}
			}
		}
	}
	return nil
}
