package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig struct {
	ReportConfigId                             CSI_ReportConfigId                                          `madatory`
	Carrier                                    *ServCellIndex                                              `optional`
	ResourcesForChannelMeasurement             CSI_ResourceConfigId                                        `madatory`
	Csi_IM_ResourcesForInterference            *CSI_ResourceConfigId                                       `optional`
	Nzp_CSI_RS_ResourcesForInterference        *CSI_ResourceConfigId                                       `optional`
	ReportConfigType                           []int64                                                     `lb:1,ub:maxNrofUL_Allocations,e_lb:0,e_ub:32,madatory`
	ReportQuantity                             *CSI_ReportConfig_reportQuantity                            `optional`
	ReportFreqConfiguration                    *CSI_ReportConfig_reportFreqConfiguration                   `lb:3,ub:3,optional`
	TimeRestrictionForChannelMeasurements      CSI_ReportConfig_timeRestrictionForChannelMeasurements      `madatory,ext`
	TimeRestrictionForInterferenceMeasurements CSI_ReportConfig_timeRestrictionForInterferenceMeasurements `madatory,ext`
	CodebookConfig                             *CodebookConfig                                             `optional,ext`
	Dummy                                      *CSI_ReportConfig_dummy                                     `optional,ext`
	GroupBasedBeamReporting                    *CSI_ReportConfig_groupBasedBeamReporting                   `optional,ext`
	Cqi_Table                                  *CSI_ReportConfig_cqi_Table                                 `optional,ext`
	SubbandSize                                CSI_ReportConfig_subbandSize                                `madatory,ext`
	Non_PMI_PortIndication                     []PortIndexFor8Ranks                                        `lb:1,ub:maxNrofNZP_CSI_RS_ResourcesPerConfig,optional,ext`
	SemiPersistentOnPUSCH_v1530                *CSI_ReportConfig_semiPersistentOnPUSCH_v1530               `optional,ext-1`
	SemiPersistentOnPUSCH_v1610                []int64                                                     `lb:1,ub:maxNrofUL_Allocations_r16,e_lb:0,e_ub:32,optional,ext-2`
	Aperiodic_v1610                            []int64                                                     `lb:1,ub:maxNrofUL_Allocations_r16,e_lb:0,e_ub:32,optional,ext-2`
	ReportQuantity_r16                         *CSI_ReportConfig_reportQuantity_r16                        `optional,ext-2`
	CodebookConfig_r16                         *CodebookConfig_r16                                         `optional,ext-2`
	Cqi_BitsPerSubband_r17                     *CSI_ReportConfig_cqi_BitsPerSubband_r17                    `optional,ext-3`
	GroupBasedBeamReporting_v1710              *CSI_ReportConfig_groupBasedBeamReporting_v1710             `optional,ext-3`
	CodebookConfig_r17                         *CodebookConfig_r17                                         `optional,ext-3`
	SharedCMR_r17                              *CSI_ReportConfig_sharedCMR_r17                             `optional,ext-3`
	Csi_ReportMode_r17                         *CSI_ReportConfig_csi_ReportMode_r17                        `optional,ext-3`
	NumberOfSingleTRP_CSI_Mode1_r17            *CSI_ReportConfig_numberOfSingleTRP_CSI_Mode1_r17           `optional,ext-3`
	ReportQuantity_r17                         *CSI_ReportConfig_reportQuantity_r17                        `optional,ext-3`
	SemiPersistentOnPUSCH_v1720                []int64                                                     `lb:1,ub:maxNrofUL_Allocations_r16,e_lb:0,e_ub:128,optional,ext-4`
	Aperiodic_v1720                            []int64                                                     `lb:1,ub:maxNrofUL_Allocations_r16,e_lb:0,e_ub:128,optional,ext-4`
	CodebookConfig_v1730                       *CodebookConfig_v1730                                       `optional,ext-5`
}

func (ie *CSI_ReportConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.SemiPersistentOnPUSCH_v1530 != nil || len(ie.SemiPersistentOnPUSCH_v1610) > 0 || len(ie.Aperiodic_v1610) > 0 || ie.ReportQuantity_r16 != nil || ie.CodebookConfig_r16 != nil || ie.Cqi_BitsPerSubband_r17 != nil || ie.GroupBasedBeamReporting_v1710 != nil || ie.CodebookConfig_r17 != nil || ie.SharedCMR_r17 != nil || ie.Csi_ReportMode_r17 != nil || ie.NumberOfSingleTRP_CSI_Mode1_r17 != nil || ie.ReportQuantity_r17 != nil || len(ie.SemiPersistentOnPUSCH_v1720) > 0 || len(ie.Aperiodic_v1720) > 0 || ie.CodebookConfig_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.Carrier != nil, ie.Csi_IM_ResourcesForInterference != nil, ie.Nzp_CSI_RS_ResourcesForInterference != nil, ie.ReportQuantity != nil, ie.ReportFreqConfiguration != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ReportConfigId.Encode(w); err != nil {
		return utils.WrapError("Encode ReportConfigId", err)
	}
	if ie.Carrier != nil {
		if err = ie.Carrier.Encode(w); err != nil {
			return utils.WrapError("Encode Carrier", err)
		}
	}
	if err = ie.ResourcesForChannelMeasurement.Encode(w); err != nil {
		return utils.WrapError("Encode ResourcesForChannelMeasurement", err)
	}
	if ie.Csi_IM_ResourcesForInterference != nil {
		if err = ie.Csi_IM_ResourcesForInterference.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_IM_ResourcesForInterference", err)
		}
	}
	if ie.Nzp_CSI_RS_ResourcesForInterference != nil {
		if err = ie.Nzp_CSI_RS_ResourcesForInterference.Encode(w); err != nil {
			return utils.WrapError("Encode Nzp_CSI_RS_ResourcesForInterference", err)
		}
	}
	tmp_ReportConfigType := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false)
	for _, i := range ie.ReportConfigType {
		tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 32}, false)
		tmp_ReportConfigType.Value = append(tmp_ReportConfigType.Value, &tmp_ie)
	}
	if err = tmp_ReportConfigType.Encode(w); err != nil {
		return utils.WrapError("Encode ReportConfigType", err)
	}
	if ie.ReportQuantity != nil {
		if err = ie.ReportQuantity.Encode(w); err != nil {
			return utils.WrapError("Encode ReportQuantity", err)
		}
	}
	if ie.ReportFreqConfiguration != nil {
		if err = ie.ReportFreqConfiguration.Encode(w); err != nil {
			return utils.WrapError("Encode ReportFreqConfiguration", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.SemiPersistentOnPUSCH_v1530 != nil, len(ie.SemiPersistentOnPUSCH_v1610) > 0 || len(ie.Aperiodic_v1610) > 0 || ie.ReportQuantity_r16 != nil || ie.CodebookConfig_r16 != nil, ie.Cqi_BitsPerSubband_r17 != nil || ie.GroupBasedBeamReporting_v1710 != nil || ie.CodebookConfig_r17 != nil || ie.SharedCMR_r17 != nil || ie.Csi_ReportMode_r17 != nil || ie.NumberOfSingleTRP_CSI_Mode1_r17 != nil || ie.ReportQuantity_r17 != nil, len(ie.SemiPersistentOnPUSCH_v1720) > 0 || len(ie.Aperiodic_v1720) > 0, ie.CodebookConfig_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_ReportConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.SemiPersistentOnPUSCH_v1530 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SemiPersistentOnPUSCH_v1530 optional
			if ie.SemiPersistentOnPUSCH_v1530 != nil {
				if err = ie.SemiPersistentOnPUSCH_v1530.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SemiPersistentOnPUSCH_v1530", err)
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
			optionals_ext_2 := []bool{len(ie.SemiPersistentOnPUSCH_v1610) > 0, len(ie.Aperiodic_v1610) > 0, ie.ReportQuantity_r16 != nil, ie.CodebookConfig_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SemiPersistentOnPUSCH_v1610 optional
			if len(ie.SemiPersistentOnPUSCH_v1610) > 0 {
				tmp_SemiPersistentOnPUSCH_v1610 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				for _, i := range ie.SemiPersistentOnPUSCH_v1610 {
					tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 32}, false)
					tmp_SemiPersistentOnPUSCH_v1610.Value = append(tmp_SemiPersistentOnPUSCH_v1610.Value, &tmp_ie)
				}
				if err = tmp_SemiPersistentOnPUSCH_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SemiPersistentOnPUSCH_v1610", err)
				}
			}
			// encode Aperiodic_v1610 optional
			if len(ie.Aperiodic_v1610) > 0 {
				tmp_Aperiodic_v1610 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				for _, i := range ie.Aperiodic_v1610 {
					tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 32}, false)
					tmp_Aperiodic_v1610.Value = append(tmp_Aperiodic_v1610.Value, &tmp_ie)
				}
				if err = tmp_Aperiodic_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Aperiodic_v1610", err)
				}
			}
			// encode ReportQuantity_r16 optional
			if ie.ReportQuantity_r16 != nil {
				if err = ie.ReportQuantity_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReportQuantity_r16", err)
				}
			}
			// encode CodebookConfig_r16 optional
			if ie.CodebookConfig_r16 != nil {
				if err = ie.CodebookConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookConfig_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.Cqi_BitsPerSubband_r17 != nil, ie.GroupBasedBeamReporting_v1710 != nil, ie.CodebookConfig_r17 != nil, ie.SharedCMR_r17 != nil, ie.Csi_ReportMode_r17 != nil, ie.NumberOfSingleTRP_CSI_Mode1_r17 != nil, ie.ReportQuantity_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Cqi_BitsPerSubband_r17 optional
			if ie.Cqi_BitsPerSubband_r17 != nil {
				if err = ie.Cqi_BitsPerSubband_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cqi_BitsPerSubband_r17", err)
				}
			}
			// encode GroupBasedBeamReporting_v1710 optional
			if ie.GroupBasedBeamReporting_v1710 != nil {
				if err = ie.GroupBasedBeamReporting_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GroupBasedBeamReporting_v1710", err)
				}
			}
			// encode CodebookConfig_r17 optional
			if ie.CodebookConfig_r17 != nil {
				if err = ie.CodebookConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookConfig_r17", err)
				}
			}
			// encode SharedCMR_r17 optional
			if ie.SharedCMR_r17 != nil {
				if err = ie.SharedCMR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SharedCMR_r17", err)
				}
			}
			// encode Csi_ReportMode_r17 optional
			if ie.Csi_ReportMode_r17 != nil {
				if err = ie.Csi_ReportMode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_ReportMode_r17", err)
				}
			}
			// encode NumberOfSingleTRP_CSI_Mode1_r17 optional
			if ie.NumberOfSingleTRP_CSI_Mode1_r17 != nil {
				if err = ie.NumberOfSingleTRP_CSI_Mode1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NumberOfSingleTRP_CSI_Mode1_r17", err)
				}
			}
			// encode ReportQuantity_r17 optional
			if ie.ReportQuantity_r17 != nil {
				if err = ie.ReportQuantity_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReportQuantity_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{len(ie.SemiPersistentOnPUSCH_v1720) > 0, len(ie.Aperiodic_v1720) > 0}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SemiPersistentOnPUSCH_v1720 optional
			if len(ie.SemiPersistentOnPUSCH_v1720) > 0 {
				tmp_SemiPersistentOnPUSCH_v1720 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				for _, i := range ie.SemiPersistentOnPUSCH_v1720 {
					tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 128}, false)
					tmp_SemiPersistentOnPUSCH_v1720.Value = append(tmp_SemiPersistentOnPUSCH_v1720.Value, &tmp_ie)
				}
				if err = tmp_SemiPersistentOnPUSCH_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SemiPersistentOnPUSCH_v1720", err)
				}
			}
			// encode Aperiodic_v1720 optional
			if len(ie.Aperiodic_v1720) > 0 {
				tmp_Aperiodic_v1720 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				for _, i := range ie.Aperiodic_v1720 {
					tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 128}, false)
					tmp_Aperiodic_v1720.Value = append(tmp_Aperiodic_v1720.Value, &tmp_ie)
				}
				if err = tmp_Aperiodic_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Aperiodic_v1720", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.CodebookConfig_v1730 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CodebookConfig_v1730 optional
			if ie.CodebookConfig_v1730 != nil {
				if err = ie.CodebookConfig_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookConfig_v1730", err)
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

func (ie *CSI_ReportConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var CarrierPresent bool
	if CarrierPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_IM_ResourcesForInterferencePresent bool
	if Csi_IM_ResourcesForInterferencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Nzp_CSI_RS_ResourcesForInterferencePresent bool
	if Nzp_CSI_RS_ResourcesForInterferencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReportQuantityPresent bool
	if ReportQuantityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReportFreqConfigurationPresent bool
	if ReportFreqConfigurationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ReportConfigId.Decode(r); err != nil {
		return utils.WrapError("Decode ReportConfigId", err)
	}
	if CarrierPresent {
		ie.Carrier = new(ServCellIndex)
		if err = ie.Carrier.Decode(r); err != nil {
			return utils.WrapError("Decode Carrier", err)
		}
	}
	if err = ie.ResourcesForChannelMeasurement.Decode(r); err != nil {
		return utils.WrapError("Decode ResourcesForChannelMeasurement", err)
	}
	if Csi_IM_ResourcesForInterferencePresent {
		ie.Csi_IM_ResourcesForInterference = new(CSI_ResourceConfigId)
		if err = ie.Csi_IM_ResourcesForInterference.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_IM_ResourcesForInterference", err)
		}
	}
	if Nzp_CSI_RS_ResourcesForInterferencePresent {
		ie.Nzp_CSI_RS_ResourcesForInterference = new(CSI_ResourceConfigId)
		if err = ie.Nzp_CSI_RS_ResourcesForInterference.Decode(r); err != nil {
			return utils.WrapError("Decode Nzp_CSI_RS_ResourcesForInterference", err)
		}
	}
	tmp_ReportConfigType := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false)
	fn_ReportConfigType := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 32}, false)
		return &ie
	}
	if err = tmp_ReportConfigType.Decode(r, fn_ReportConfigType); err != nil {
		return utils.WrapError("Decode ReportConfigType", err)
	}
	ie.ReportConfigType = []int64{}
	for _, i := range tmp_ReportConfigType.Value {
		ie.ReportConfigType = append(ie.ReportConfigType, int64(i.Value))
	}
	if ReportQuantityPresent {
		ie.ReportQuantity = new(CSI_ReportConfig_reportQuantity)
		if err = ie.ReportQuantity.Decode(r); err != nil {
			return utils.WrapError("Decode ReportQuantity", err)
		}
	}
	if ReportFreqConfigurationPresent {
		ie.ReportFreqConfiguration = new(CSI_ReportConfig_reportFreqConfiguration)
		if err = ie.ReportFreqConfiguration.Decode(r); err != nil {
			return utils.WrapError("Decode ReportFreqConfiguration", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 5 bits for 5 extension groups
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

			SemiPersistentOnPUSCH_v1530Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SemiPersistentOnPUSCH_v1530 optional
			if SemiPersistentOnPUSCH_v1530Present {
				ie.SemiPersistentOnPUSCH_v1530 = new(CSI_ReportConfig_semiPersistentOnPUSCH_v1530)
				if err = ie.SemiPersistentOnPUSCH_v1530.Decode(extReader); err != nil {
					return utils.WrapError("Decode SemiPersistentOnPUSCH_v1530", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SemiPersistentOnPUSCH_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Aperiodic_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ReportQuantity_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SemiPersistentOnPUSCH_v1610 optional
			if SemiPersistentOnPUSCH_v1610Present {
				tmp_SemiPersistentOnPUSCH_v1610 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				fn_SemiPersistentOnPUSCH_v1610 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 32}, false)
					return &ie
				}
				if err = tmp_SemiPersistentOnPUSCH_v1610.Decode(extReader, fn_SemiPersistentOnPUSCH_v1610); err != nil {
					return utils.WrapError("Decode SemiPersistentOnPUSCH_v1610", err)
				}
				ie.SemiPersistentOnPUSCH_v1610 = []int64{}
				for _, i := range tmp_SemiPersistentOnPUSCH_v1610.Value {
					ie.SemiPersistentOnPUSCH_v1610 = append(ie.SemiPersistentOnPUSCH_v1610, int64(i.Value))
				}
			}
			// decode Aperiodic_v1610 optional
			if Aperiodic_v1610Present {
				tmp_Aperiodic_v1610 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				fn_Aperiodic_v1610 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 32}, false)
					return &ie
				}
				if err = tmp_Aperiodic_v1610.Decode(extReader, fn_Aperiodic_v1610); err != nil {
					return utils.WrapError("Decode Aperiodic_v1610", err)
				}
				ie.Aperiodic_v1610 = []int64{}
				for _, i := range tmp_Aperiodic_v1610.Value {
					ie.Aperiodic_v1610 = append(ie.Aperiodic_v1610, int64(i.Value))
				}
			}
			// decode ReportQuantity_r16 optional
			if ReportQuantity_r16Present {
				ie.ReportQuantity_r16 = new(CSI_ReportConfig_reportQuantity_r16)
				if err = ie.ReportQuantity_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReportQuantity_r16", err)
				}
			}
			// decode CodebookConfig_r16 optional
			if CodebookConfig_r16Present {
				ie.CodebookConfig_r16 = new(CodebookConfig_r16)
				if err = ie.CodebookConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookConfig_r16", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Cqi_BitsPerSubband_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GroupBasedBeamReporting_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SharedCMR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_ReportMode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NumberOfSingleTRP_CSI_Mode1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ReportQuantity_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Cqi_BitsPerSubband_r17 optional
			if Cqi_BitsPerSubband_r17Present {
				ie.Cqi_BitsPerSubband_r17 = new(CSI_ReportConfig_cqi_BitsPerSubband_r17)
				if err = ie.Cqi_BitsPerSubband_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cqi_BitsPerSubband_r17", err)
				}
			}
			// decode GroupBasedBeamReporting_v1710 optional
			if GroupBasedBeamReporting_v1710Present {
				ie.GroupBasedBeamReporting_v1710 = new(CSI_ReportConfig_groupBasedBeamReporting_v1710)
				if err = ie.GroupBasedBeamReporting_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode GroupBasedBeamReporting_v1710", err)
				}
			}
			// decode CodebookConfig_r17 optional
			if CodebookConfig_r17Present {
				ie.CodebookConfig_r17 = new(CodebookConfig_r17)
				if err = ie.CodebookConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookConfig_r17", err)
				}
			}
			// decode SharedCMR_r17 optional
			if SharedCMR_r17Present {
				ie.SharedCMR_r17 = new(CSI_ReportConfig_sharedCMR_r17)
				if err = ie.SharedCMR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SharedCMR_r17", err)
				}
			}
			// decode Csi_ReportMode_r17 optional
			if Csi_ReportMode_r17Present {
				ie.Csi_ReportMode_r17 = new(CSI_ReportConfig_csi_ReportMode_r17)
				if err = ie.Csi_ReportMode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_ReportMode_r17", err)
				}
			}
			// decode NumberOfSingleTRP_CSI_Mode1_r17 optional
			if NumberOfSingleTRP_CSI_Mode1_r17Present {
				ie.NumberOfSingleTRP_CSI_Mode1_r17 = new(CSI_ReportConfig_numberOfSingleTRP_CSI_Mode1_r17)
				if err = ie.NumberOfSingleTRP_CSI_Mode1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode NumberOfSingleTRP_CSI_Mode1_r17", err)
				}
			}
			// decode ReportQuantity_r17 optional
			if ReportQuantity_r17Present {
				ie.ReportQuantity_r17 = new(CSI_ReportConfig_reportQuantity_r17)
				if err = ie.ReportQuantity_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReportQuantity_r17", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SemiPersistentOnPUSCH_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Aperiodic_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SemiPersistentOnPUSCH_v1720 optional
			if SemiPersistentOnPUSCH_v1720Present {
				tmp_SemiPersistentOnPUSCH_v1720 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				fn_SemiPersistentOnPUSCH_v1720 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 128}, false)
					return &ie
				}
				if err = tmp_SemiPersistentOnPUSCH_v1720.Decode(extReader, fn_SemiPersistentOnPUSCH_v1720); err != nil {
					return utils.WrapError("Decode SemiPersistentOnPUSCH_v1720", err)
				}
				ie.SemiPersistentOnPUSCH_v1720 = []int64{}
				for _, i := range tmp_SemiPersistentOnPUSCH_v1720.Value {
					ie.SemiPersistentOnPUSCH_v1720 = append(ie.SemiPersistentOnPUSCH_v1720, int64(i.Value))
				}
			}
			// decode Aperiodic_v1720 optional
			if Aperiodic_v1720Present {
				tmp_Aperiodic_v1720 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations_r16}, false)
				fn_Aperiodic_v1720 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 128}, false)
					return &ie
				}
				if err = tmp_Aperiodic_v1720.Decode(extReader, fn_Aperiodic_v1720); err != nil {
					return utils.WrapError("Decode Aperiodic_v1720", err)
				}
				ie.Aperiodic_v1720 = []int64{}
				for _, i := range tmp_Aperiodic_v1720.Value {
					ie.Aperiodic_v1720 = append(ie.Aperiodic_v1720, int64(i.Value))
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			CodebookConfig_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CodebookConfig_v1730 optional
			if CodebookConfig_v1730Present {
				ie.CodebookConfig_v1730 = new(CodebookConfig_v1730)
				if err = ie.CodebookConfig_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookConfig_v1730", err)
				}
			}
		}
	}
	return nil
}
