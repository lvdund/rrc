package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_Config struct {
	ControlResourceSetToAddModList              []ControlResourceSet                           `lb:1,ub:3,optional`
	ControlResourceSetToReleaseList             []ControlResourceSetId                         `lb:1,ub:3,optional`
	SearchSpacesToAddModList                    []SearchSpace                                  `lb:1,ub:10,optional`
	SearchSpacesToReleaseList                   []SearchSpaceId                                `lb:1,ub:10,optional`
	DownlinkPreemption                          *DownlinkPreemption                            `optional,setuprelease`
	Tpc_PUSCH                                   *PUSCH_TPC_CommandConfig                       `optional,setuprelease`
	Tpc_PUCCH                                   *PUCCH_TPC_CommandConfig                       `optional,setuprelease`
	Tpc_SRS                                     *SRS_TPC_CommandConfig                         `optional,setuprelease`
	ControlResourceSetToAddModListSizeExt_v1610 []ControlResourceSet                           `lb:1,ub:2,optional,ext-1`
	ControlResourceSetToReleaseListSizeExt_r16  []ControlResourceSetId_r16                     `lb:1,ub:5,optional,ext-1`
	SearchSpacesToAddModListExt_r16             []SearchSpaceExt_r16                           `lb:1,ub:10,optional,ext-1`
	UplinkCancellation_r16                      *UplinkCancellation_r16                        `optional,ext-1,setuprelease`
	MonitoringCapabilityConfig_r16              *PDCCH_Config_monitoringCapabilityConfig_r16   `optional,ext-1`
	SearchSpaceSwitchConfig_r16                 *SearchSpaceSwitchConfig_r16                   `optional,ext-1`
	SearchSpacesToAddModListExt_v1700           []SearchSpaceExt_v1700                         `lb:1,ub:10,optional,ext-2`
	MonitoringCapabilityConfig_v1710            *PDCCH_Config_monitoringCapabilityConfig_v1710 `optional,ext-2`
	SearchSpaceSwitchConfig_r17                 *SearchSpaceSwitchConfig_r17                   `optional,ext-2`
	Pdcch_SkippingDurationList_r17              []SCS_SpecificDuration_r17                     `lb:1,ub:3,optional,ext-2`
}

func (ie *PDCCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.ControlResourceSetToAddModListSizeExt_v1610) > 0 || len(ie.ControlResourceSetToReleaseListSizeExt_r16) > 0 || len(ie.SearchSpacesToAddModListExt_r16) > 0 || ie.UplinkCancellation_r16 != nil || ie.MonitoringCapabilityConfig_r16 != nil || ie.SearchSpaceSwitchConfig_r16 != nil || len(ie.SearchSpacesToAddModListExt_v1700) > 0 || ie.MonitoringCapabilityConfig_v1710 != nil || ie.SearchSpaceSwitchConfig_r17 != nil || len(ie.Pdcch_SkippingDurationList_r17) > 0
	preambleBits := []bool{hasExtensions, len(ie.ControlResourceSetToAddModList) > 0, len(ie.ControlResourceSetToReleaseList) > 0, len(ie.SearchSpacesToAddModList) > 0, len(ie.SearchSpacesToReleaseList) > 0, ie.DownlinkPreemption != nil, ie.Tpc_PUSCH != nil, ie.Tpc_PUCCH != nil, ie.Tpc_SRS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.ControlResourceSetToAddModList) > 0 {
		tmp_ControlResourceSetToAddModList := utils.NewSequence[*ControlResourceSet]([]*ControlResourceSet{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		for _, i := range ie.ControlResourceSetToAddModList {
			tmp_ControlResourceSetToAddModList.Value = append(tmp_ControlResourceSetToAddModList.Value, &i)
		}
		if err = tmp_ControlResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode ControlResourceSetToAddModList", err)
		}
	}
	if len(ie.ControlResourceSetToReleaseList) > 0 {
		tmp_ControlResourceSetToReleaseList := utils.NewSequence[*ControlResourceSetId]([]*ControlResourceSetId{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		for _, i := range ie.ControlResourceSetToReleaseList {
			tmp_ControlResourceSetToReleaseList.Value = append(tmp_ControlResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_ControlResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode ControlResourceSetToReleaseList", err)
		}
	}
	if len(ie.SearchSpacesToAddModList) > 0 {
		tmp_SearchSpacesToAddModList := utils.NewSequence[*SearchSpace]([]*SearchSpace{}, uper.Constraint{Lb: 1, Ub: 10}, false)
		for _, i := range ie.SearchSpacesToAddModList {
			tmp_SearchSpacesToAddModList.Value = append(tmp_SearchSpacesToAddModList.Value, &i)
		}
		if err = tmp_SearchSpacesToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpacesToAddModList", err)
		}
	}
	if len(ie.SearchSpacesToReleaseList) > 0 {
		tmp_SearchSpacesToReleaseList := utils.NewSequence[*SearchSpaceId]([]*SearchSpaceId{}, uper.Constraint{Lb: 1, Ub: 10}, false)
		for _, i := range ie.SearchSpacesToReleaseList {
			tmp_SearchSpacesToReleaseList.Value = append(tmp_SearchSpacesToReleaseList.Value, &i)
		}
		if err = tmp_SearchSpacesToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpacesToReleaseList", err)
		}
	}
	if ie.DownlinkPreemption != nil {
		tmp_DownlinkPreemption := utils.SetupRelease[*DownlinkPreemption]{
			Setup: ie.DownlinkPreemption,
		}
		if err = tmp_DownlinkPreemption.Encode(w); err != nil {
			return utils.WrapError("Encode DownlinkPreemption", err)
		}
	}
	if ie.Tpc_PUSCH != nil {
		tmp_Tpc_PUSCH := utils.SetupRelease[*PUSCH_TPC_CommandConfig]{
			Setup: ie.Tpc_PUSCH,
		}
		if err = tmp_Tpc_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_PUSCH", err)
		}
	}
	if ie.Tpc_PUCCH != nil {
		tmp_Tpc_PUCCH := utils.SetupRelease[*PUCCH_TPC_CommandConfig]{
			Setup: ie.Tpc_PUCCH,
		}
		if err = tmp_Tpc_PUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_PUCCH", err)
		}
	}
	if ie.Tpc_SRS != nil {
		tmp_Tpc_SRS := utils.SetupRelease[*SRS_TPC_CommandConfig]{
			Setup: ie.Tpc_SRS,
		}
		if err = tmp_Tpc_SRS.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_SRS", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{len(ie.ControlResourceSetToAddModListSizeExt_v1610) > 0 || len(ie.ControlResourceSetToReleaseListSizeExt_r16) > 0 || len(ie.SearchSpacesToAddModListExt_r16) > 0 || ie.UplinkCancellation_r16 != nil || ie.MonitoringCapabilityConfig_r16 != nil || ie.SearchSpaceSwitchConfig_r16 != nil, len(ie.SearchSpacesToAddModListExt_v1700) > 0 || ie.MonitoringCapabilityConfig_v1710 != nil || ie.SearchSpaceSwitchConfig_r17 != nil || len(ie.Pdcch_SkippingDurationList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDCCH_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.ControlResourceSetToAddModListSizeExt_v1610) > 0, len(ie.ControlResourceSetToReleaseListSizeExt_r16) > 0, len(ie.SearchSpacesToAddModListExt_r16) > 0, ie.UplinkCancellation_r16 != nil, ie.MonitoringCapabilityConfig_r16 != nil, ie.SearchSpaceSwitchConfig_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ControlResourceSetToAddModListSizeExt_v1610 optional
			if len(ie.ControlResourceSetToAddModListSizeExt_v1610) > 0 {
				tmp_ControlResourceSetToAddModListSizeExt_v1610 := utils.NewSequence[*ControlResourceSet]([]*ControlResourceSet{}, uper.Constraint{Lb: 1, Ub: 2}, false)
				for _, i := range ie.ControlResourceSetToAddModListSizeExt_v1610 {
					tmp_ControlResourceSetToAddModListSizeExt_v1610.Value = append(tmp_ControlResourceSetToAddModListSizeExt_v1610.Value, &i)
				}
				if err = tmp_ControlResourceSetToAddModListSizeExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ControlResourceSetToAddModListSizeExt_v1610", err)
				}
			}
			// encode ControlResourceSetToReleaseListSizeExt_r16 optional
			if len(ie.ControlResourceSetToReleaseListSizeExt_r16) > 0 {
				tmp_ControlResourceSetToReleaseListSizeExt_r16 := utils.NewSequence[*ControlResourceSetId_r16]([]*ControlResourceSetId_r16{}, uper.Constraint{Lb: 1, Ub: 5}, false)
				for _, i := range ie.ControlResourceSetToReleaseListSizeExt_r16 {
					tmp_ControlResourceSetToReleaseListSizeExt_r16.Value = append(tmp_ControlResourceSetToReleaseListSizeExt_r16.Value, &i)
				}
				if err = tmp_ControlResourceSetToReleaseListSizeExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ControlResourceSetToReleaseListSizeExt_r16", err)
				}
			}
			// encode SearchSpacesToAddModListExt_r16 optional
			if len(ie.SearchSpacesToAddModListExt_r16) > 0 {
				tmp_SearchSpacesToAddModListExt_r16 := utils.NewSequence[*SearchSpaceExt_r16]([]*SearchSpaceExt_r16{}, uper.Constraint{Lb: 1, Ub: 10}, false)
				for _, i := range ie.SearchSpacesToAddModListExt_r16 {
					tmp_SearchSpacesToAddModListExt_r16.Value = append(tmp_SearchSpacesToAddModListExt_r16.Value, &i)
				}
				if err = tmp_SearchSpacesToAddModListExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SearchSpacesToAddModListExt_r16", err)
				}
			}
			// encode UplinkCancellation_r16 optional
			if ie.UplinkCancellation_r16 != nil {
				tmp_UplinkCancellation_r16 := utils.SetupRelease[*UplinkCancellation_r16]{
					Setup: ie.UplinkCancellation_r16,
				}
				if err = tmp_UplinkCancellation_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkCancellation_r16", err)
				}
			}
			// encode MonitoringCapabilityConfig_r16 optional
			if ie.MonitoringCapabilityConfig_r16 != nil {
				if err = ie.MonitoringCapabilityConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MonitoringCapabilityConfig_r16", err)
				}
			}
			// encode SearchSpaceSwitchConfig_r16 optional
			if ie.SearchSpaceSwitchConfig_r16 != nil {
				if err = ie.SearchSpaceSwitchConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SearchSpaceSwitchConfig_r16", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{len(ie.SearchSpacesToAddModListExt_v1700) > 0, ie.MonitoringCapabilityConfig_v1710 != nil, ie.SearchSpaceSwitchConfig_r17 != nil, len(ie.Pdcch_SkippingDurationList_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SearchSpacesToAddModListExt_v1700 optional
			if len(ie.SearchSpacesToAddModListExt_v1700) > 0 {
				tmp_SearchSpacesToAddModListExt_v1700 := utils.NewSequence[*SearchSpaceExt_v1700]([]*SearchSpaceExt_v1700{}, uper.Constraint{Lb: 1, Ub: 10}, false)
				for _, i := range ie.SearchSpacesToAddModListExt_v1700 {
					tmp_SearchSpacesToAddModListExt_v1700.Value = append(tmp_SearchSpacesToAddModListExt_v1700.Value, &i)
				}
				if err = tmp_SearchSpacesToAddModListExt_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SearchSpacesToAddModListExt_v1700", err)
				}
			}
			// encode MonitoringCapabilityConfig_v1710 optional
			if ie.MonitoringCapabilityConfig_v1710 != nil {
				if err = ie.MonitoringCapabilityConfig_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MonitoringCapabilityConfig_v1710", err)
				}
			}
			// encode SearchSpaceSwitchConfig_r17 optional
			if ie.SearchSpaceSwitchConfig_r17 != nil {
				if err = ie.SearchSpaceSwitchConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SearchSpaceSwitchConfig_r17", err)
				}
			}
			// encode Pdcch_SkippingDurationList_r17 optional
			if len(ie.Pdcch_SkippingDurationList_r17) > 0 {
				tmp_Pdcch_SkippingDurationList_r17 := utils.NewSequence[*SCS_SpecificDuration_r17]([]*SCS_SpecificDuration_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
				for _, i := range ie.Pdcch_SkippingDurationList_r17 {
					tmp_Pdcch_SkippingDurationList_r17.Value = append(tmp_Pdcch_SkippingDurationList_r17.Value, &i)
				}
				if err = tmp_Pdcch_SkippingDurationList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_SkippingDurationList_r17", err)
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

func (ie *PDCCH_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ControlResourceSetToAddModListPresent bool
	if ControlResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ControlResourceSetToReleaseListPresent bool
	if ControlResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpacesToAddModListPresent bool
	if SearchSpacesToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpacesToReleaseListPresent bool
	if SearchSpacesToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DownlinkPreemptionPresent bool
	if DownlinkPreemptionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_PUSCHPresent bool
	if Tpc_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_PUCCHPresent bool
	if Tpc_PUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_SRSPresent bool
	if Tpc_SRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ControlResourceSetToAddModListPresent {
		tmp_ControlResourceSetToAddModList := utils.NewSequence[*ControlResourceSet]([]*ControlResourceSet{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		fn_ControlResourceSetToAddModList := func() *ControlResourceSet {
			return new(ControlResourceSet)
		}
		if err = tmp_ControlResourceSetToAddModList.Decode(r, fn_ControlResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode ControlResourceSetToAddModList", err)
		}
		ie.ControlResourceSetToAddModList = []ControlResourceSet{}
		for _, i := range tmp_ControlResourceSetToAddModList.Value {
			ie.ControlResourceSetToAddModList = append(ie.ControlResourceSetToAddModList, *i)
		}
	}
	if ControlResourceSetToReleaseListPresent {
		tmp_ControlResourceSetToReleaseList := utils.NewSequence[*ControlResourceSetId]([]*ControlResourceSetId{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		fn_ControlResourceSetToReleaseList := func() *ControlResourceSetId {
			return new(ControlResourceSetId)
		}
		if err = tmp_ControlResourceSetToReleaseList.Decode(r, fn_ControlResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode ControlResourceSetToReleaseList", err)
		}
		ie.ControlResourceSetToReleaseList = []ControlResourceSetId{}
		for _, i := range tmp_ControlResourceSetToReleaseList.Value {
			ie.ControlResourceSetToReleaseList = append(ie.ControlResourceSetToReleaseList, *i)
		}
	}
	if SearchSpacesToAddModListPresent {
		tmp_SearchSpacesToAddModList := utils.NewSequence[*SearchSpace]([]*SearchSpace{}, uper.Constraint{Lb: 1, Ub: 10}, false)
		fn_SearchSpacesToAddModList := func() *SearchSpace {
			return new(SearchSpace)
		}
		if err = tmp_SearchSpacesToAddModList.Decode(r, fn_SearchSpacesToAddModList); err != nil {
			return utils.WrapError("Decode SearchSpacesToAddModList", err)
		}
		ie.SearchSpacesToAddModList = []SearchSpace{}
		for _, i := range tmp_SearchSpacesToAddModList.Value {
			ie.SearchSpacesToAddModList = append(ie.SearchSpacesToAddModList, *i)
		}
	}
	if SearchSpacesToReleaseListPresent {
		tmp_SearchSpacesToReleaseList := utils.NewSequence[*SearchSpaceId]([]*SearchSpaceId{}, uper.Constraint{Lb: 1, Ub: 10}, false)
		fn_SearchSpacesToReleaseList := func() *SearchSpaceId {
			return new(SearchSpaceId)
		}
		if err = tmp_SearchSpacesToReleaseList.Decode(r, fn_SearchSpacesToReleaseList); err != nil {
			return utils.WrapError("Decode SearchSpacesToReleaseList", err)
		}
		ie.SearchSpacesToReleaseList = []SearchSpaceId{}
		for _, i := range tmp_SearchSpacesToReleaseList.Value {
			ie.SearchSpacesToReleaseList = append(ie.SearchSpacesToReleaseList, *i)
		}
	}
	if DownlinkPreemptionPresent {
		tmp_DownlinkPreemption := utils.SetupRelease[*DownlinkPreemption]{}
		if err = tmp_DownlinkPreemption.Decode(r); err != nil {
			return utils.WrapError("Decode DownlinkPreemption", err)
		}
		ie.DownlinkPreemption = tmp_DownlinkPreemption.Setup
	}
	if Tpc_PUSCHPresent {
		tmp_Tpc_PUSCH := utils.SetupRelease[*PUSCH_TPC_CommandConfig]{}
		if err = tmp_Tpc_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_PUSCH", err)
		}
		ie.Tpc_PUSCH = tmp_Tpc_PUSCH.Setup
	}
	if Tpc_PUCCHPresent {
		tmp_Tpc_PUCCH := utils.SetupRelease[*PUCCH_TPC_CommandConfig]{}
		if err = tmp_Tpc_PUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_PUCCH", err)
		}
		ie.Tpc_PUCCH = tmp_Tpc_PUCCH.Setup
	}
	if Tpc_SRSPresent {
		tmp_Tpc_SRS := utils.SetupRelease[*SRS_TPC_CommandConfig]{}
		if err = tmp_Tpc_SRS.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_SRS", err)
		}
		ie.Tpc_SRS = tmp_Tpc_SRS.Setup
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ControlResourceSetToAddModListSizeExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ControlResourceSetToReleaseListSizeExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SearchSpacesToAddModListExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkCancellation_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MonitoringCapabilityConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SearchSpaceSwitchConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ControlResourceSetToAddModListSizeExt_v1610 optional
			if ControlResourceSetToAddModListSizeExt_v1610Present {
				tmp_ControlResourceSetToAddModListSizeExt_v1610 := utils.NewSequence[*ControlResourceSet]([]*ControlResourceSet{}, uper.Constraint{Lb: 1, Ub: 2}, false)
				fn_ControlResourceSetToAddModListSizeExt_v1610 := func() *ControlResourceSet {
					return new(ControlResourceSet)
				}
				if err = tmp_ControlResourceSetToAddModListSizeExt_v1610.Decode(extReader, fn_ControlResourceSetToAddModListSizeExt_v1610); err != nil {
					return utils.WrapError("Decode ControlResourceSetToAddModListSizeExt_v1610", err)
				}
				ie.ControlResourceSetToAddModListSizeExt_v1610 = []ControlResourceSet{}
				for _, i := range tmp_ControlResourceSetToAddModListSizeExt_v1610.Value {
					ie.ControlResourceSetToAddModListSizeExt_v1610 = append(ie.ControlResourceSetToAddModListSizeExt_v1610, *i)
				}
			}
			// decode ControlResourceSetToReleaseListSizeExt_r16 optional
			if ControlResourceSetToReleaseListSizeExt_r16Present {
				tmp_ControlResourceSetToReleaseListSizeExt_r16 := utils.NewSequence[*ControlResourceSetId_r16]([]*ControlResourceSetId_r16{}, uper.Constraint{Lb: 1, Ub: 5}, false)
				fn_ControlResourceSetToReleaseListSizeExt_r16 := func() *ControlResourceSetId_r16 {
					return new(ControlResourceSetId_r16)
				}
				if err = tmp_ControlResourceSetToReleaseListSizeExt_r16.Decode(extReader, fn_ControlResourceSetToReleaseListSizeExt_r16); err != nil {
					return utils.WrapError("Decode ControlResourceSetToReleaseListSizeExt_r16", err)
				}
				ie.ControlResourceSetToReleaseListSizeExt_r16 = []ControlResourceSetId_r16{}
				for _, i := range tmp_ControlResourceSetToReleaseListSizeExt_r16.Value {
					ie.ControlResourceSetToReleaseListSizeExt_r16 = append(ie.ControlResourceSetToReleaseListSizeExt_r16, *i)
				}
			}
			// decode SearchSpacesToAddModListExt_r16 optional
			if SearchSpacesToAddModListExt_r16Present {
				tmp_SearchSpacesToAddModListExt_r16 := utils.NewSequence[*SearchSpaceExt_r16]([]*SearchSpaceExt_r16{}, uper.Constraint{Lb: 1, Ub: 10}, false)
				fn_SearchSpacesToAddModListExt_r16 := func() *SearchSpaceExt_r16 {
					return new(SearchSpaceExt_r16)
				}
				if err = tmp_SearchSpacesToAddModListExt_r16.Decode(extReader, fn_SearchSpacesToAddModListExt_r16); err != nil {
					return utils.WrapError("Decode SearchSpacesToAddModListExt_r16", err)
				}
				ie.SearchSpacesToAddModListExt_r16 = []SearchSpaceExt_r16{}
				for _, i := range tmp_SearchSpacesToAddModListExt_r16.Value {
					ie.SearchSpacesToAddModListExt_r16 = append(ie.SearchSpacesToAddModListExt_r16, *i)
				}
			}
			// decode UplinkCancellation_r16 optional
			if UplinkCancellation_r16Present {
				tmp_UplinkCancellation_r16 := utils.SetupRelease[*UplinkCancellation_r16]{}
				if err = tmp_UplinkCancellation_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkCancellation_r16", err)
				}
				ie.UplinkCancellation_r16 = tmp_UplinkCancellation_r16.Setup
			}
			// decode MonitoringCapabilityConfig_r16 optional
			if MonitoringCapabilityConfig_r16Present {
				ie.MonitoringCapabilityConfig_r16 = new(PDCCH_Config_monitoringCapabilityConfig_r16)
				if err = ie.MonitoringCapabilityConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MonitoringCapabilityConfig_r16", err)
				}
			}
			// decode SearchSpaceSwitchConfig_r16 optional
			if SearchSpaceSwitchConfig_r16Present {
				ie.SearchSpaceSwitchConfig_r16 = new(SearchSpaceSwitchConfig_r16)
				if err = ie.SearchSpaceSwitchConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SearchSpaceSwitchConfig_r16", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			SearchSpacesToAddModListExt_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MonitoringCapabilityConfig_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SearchSpaceSwitchConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdcch_SkippingDurationList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SearchSpacesToAddModListExt_v1700 optional
			if SearchSpacesToAddModListExt_v1700Present {
				tmp_SearchSpacesToAddModListExt_v1700 := utils.NewSequence[*SearchSpaceExt_v1700]([]*SearchSpaceExt_v1700{}, uper.Constraint{Lb: 1, Ub: 10}, false)
				fn_SearchSpacesToAddModListExt_v1700 := func() *SearchSpaceExt_v1700 {
					return new(SearchSpaceExt_v1700)
				}
				if err = tmp_SearchSpacesToAddModListExt_v1700.Decode(extReader, fn_SearchSpacesToAddModListExt_v1700); err != nil {
					return utils.WrapError("Decode SearchSpacesToAddModListExt_v1700", err)
				}
				ie.SearchSpacesToAddModListExt_v1700 = []SearchSpaceExt_v1700{}
				for _, i := range tmp_SearchSpacesToAddModListExt_v1700.Value {
					ie.SearchSpacesToAddModListExt_v1700 = append(ie.SearchSpacesToAddModListExt_v1700, *i)
				}
			}
			// decode MonitoringCapabilityConfig_v1710 optional
			if MonitoringCapabilityConfig_v1710Present {
				ie.MonitoringCapabilityConfig_v1710 = new(PDCCH_Config_monitoringCapabilityConfig_v1710)
				if err = ie.MonitoringCapabilityConfig_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode MonitoringCapabilityConfig_v1710", err)
				}
			}
			// decode SearchSpaceSwitchConfig_r17 optional
			if SearchSpaceSwitchConfig_r17Present {
				ie.SearchSpaceSwitchConfig_r17 = new(SearchSpaceSwitchConfig_r17)
				if err = ie.SearchSpaceSwitchConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SearchSpaceSwitchConfig_r17", err)
				}
			}
			// decode Pdcch_SkippingDurationList_r17 optional
			if Pdcch_SkippingDurationList_r17Present {
				tmp_Pdcch_SkippingDurationList_r17 := utils.NewSequence[*SCS_SpecificDuration_r17]([]*SCS_SpecificDuration_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
				fn_Pdcch_SkippingDurationList_r17 := func() *SCS_SpecificDuration_r17 {
					return new(SCS_SpecificDuration_r17)
				}
				if err = tmp_Pdcch_SkippingDurationList_r17.Decode(extReader, fn_Pdcch_SkippingDurationList_r17); err != nil {
					return utils.WrapError("Decode Pdcch_SkippingDurationList_r17", err)
				}
				ie.Pdcch_SkippingDurationList_r17 = []SCS_SpecificDuration_r17{}
				for _, i := range tmp_Pdcch_SkippingDurationList_r17.Value {
					ie.Pdcch_SkippingDurationList_r17 = append(ie.Pdcch_SkippingDurationList_r17, *i)
				}
			}
		}
	}
	return nil
}
