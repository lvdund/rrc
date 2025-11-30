package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SharedSpectrumChAccessParamsPerBand_r16 struct {
	Ul_DynamicChAccess_r16                   *SharedSpectrumChAccessParamsPerBand_r16_ul_DynamicChAccess_r16                   `optional`
	Ul_Semi_StaticChAccess_r16               *SharedSpectrumChAccessParamsPerBand_r16_ul_Semi_StaticChAccess_r16               `optional`
	Ssb_RRM_DynamicChAccess_r16              *SharedSpectrumChAccessParamsPerBand_r16_ssb_RRM_DynamicChAccess_r16              `optional`
	Ssb_RRM_Semi_StaticChAccess_r16          *SharedSpectrumChAccessParamsPerBand_r16_ssb_RRM_Semi_StaticChAccess_r16          `optional`
	Mib_Acquisition_r16                      *SharedSpectrumChAccessParamsPerBand_r16_mib_Acquisition_r16                      `optional`
	Ssb_RLM_DynamicChAccess_r16              *SharedSpectrumChAccessParamsPerBand_r16_ssb_RLM_DynamicChAccess_r16              `optional`
	Ssb_RLM_Semi_StaticChAccess_r16          *SharedSpectrumChAccessParamsPerBand_r16_ssb_RLM_Semi_StaticChAccess_r16          `optional`
	Sib1_Acquisition_r16                     *SharedSpectrumChAccessParamsPerBand_r16_sib1_Acquisition_r16                     `optional`
	ExtRA_ResponseWindow_r16                 *SharedSpectrumChAccessParamsPerBand_r16_extRA_ResponseWindow_r16                 `optional`
	Ssb_BFD_CBD_dynamicChannelAccess_r16     *SharedSpectrumChAccessParamsPerBand_r16_ssb_BFD_CBD_dynamicChannelAccess_r16     `optional`
	Ssb_BFD_CBD_semi_staticChannelAccess_r16 *SharedSpectrumChAccessParamsPerBand_r16_ssb_BFD_CBD_semi_staticChannelAccess_r16 `optional`
	Csi_RS_BFD_CBD_r16                       *SharedSpectrumChAccessParamsPerBand_r16_csi_RS_BFD_CBD_r16                       `optional`
	Ul_ChannelBW_SCell_10mhz_r16             *SharedSpectrumChAccessParamsPerBand_r16_ul_ChannelBW_SCell_10mhz_r16             `optional`
	Rssi_ChannelOccupancyReporting_r16       *SharedSpectrumChAccessParamsPerBand_r16_rssi_ChannelOccupancyReporting_r16       `optional`
	Srs_StartAnyOFDM_Symbol_r16              *SharedSpectrumChAccessParamsPerBand_r16_srs_StartAnyOFDM_Symbol_r16              `optional`
	SearchSpaceFreqMonitorLocation_r16       *int64                                                                            `lb:1,ub:5,optional`
	Coreset_RB_Offset_r16                    *SharedSpectrumChAccessParamsPerBand_r16_coreset_RB_Offset_r16                    `optional`
	Cgi_Acquisition_r16                      *SharedSpectrumChAccessParamsPerBand_r16_cgi_Acquisition_r16                      `optional`
	ConfiguredUL_Tx_r16                      *SharedSpectrumChAccessParamsPerBand_r16_configuredUL_Tx_r16                      `optional`
	Prach_Wideband_r16                       *SharedSpectrumChAccessParamsPerBand_r16_prach_Wideband_r16                       `optional`
	Dci_AvailableRB_Set_r16                  *SharedSpectrumChAccessParamsPerBand_r16_dci_AvailableRB_Set_r16                  `optional`
	Dci_ChOccupancyDuration_r16              *SharedSpectrumChAccessParamsPerBand_r16_dci_ChOccupancyDuration_r16              `optional`
	TypeB_PDSCH_length_r16                   *SharedSpectrumChAccessParamsPerBand_r16_typeB_PDSCH_length_r16                   `optional`
	SearchSpaceSwitchWithDCI_r16             *SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithDCI_r16             `optional`
	SearchSpaceSwitchWithoutDCI_r16          *SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithoutDCI_r16          `optional`
	SearchSpaceSwitchCapability2_r16         *SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchCapability2_r16         `optional`
	Non_numericalPDSCH_HARQ_timing_r16       *SharedSpectrumChAccessParamsPerBand_r16_non_numericalPDSCH_HARQ_timing_r16       `optional`
	EnhancedDynamicHARQ_codebook_r16         *SharedSpectrumChAccessParamsPerBand_r16_enhancedDynamicHARQ_codebook_r16         `optional`
	OneShotHARQ_feedback_r16                 *SharedSpectrumChAccessParamsPerBand_r16_oneShotHARQ_feedback_r16                 `optional`
	MultiPUSCH_UL_grant_r16                  *SharedSpectrumChAccessParamsPerBand_r16_multiPUSCH_UL_grant_r16                  `optional`
	Csi_RS_RLM_r16                           *SharedSpectrumChAccessParamsPerBand_r16_csi_RS_RLM_r16                           `optional`
	Dummy                                    *SharedSpectrumChAccessParamsPerBand_r16_dummy                                    `optional`
	PeriodicAndSemi_PersistentCSI_RS_r16     *SharedSpectrumChAccessParamsPerBand_r16_periodicAndSemi_PersistentCSI_RS_r16     `optional`
	Pusch_PRB_interlace_r16                  *SharedSpectrumChAccessParamsPerBand_r16_pusch_PRB_interlace_r16                  `optional`
	Pucch_F0_F1_PRB_Interlace_r16            *SharedSpectrumChAccessParamsPerBand_r16_pucch_F0_F1_PRB_Interlace_r16            `optional`
	Occ_PRB_PF2_PF3_r16                      *SharedSpectrumChAccessParamsPerBand_r16_occ_PRB_PF2_PF3_r16                      `optional`
	ExtCP_rangeCG_PUSCH_r16                  *SharedSpectrumChAccessParamsPerBand_r16_extCP_rangeCG_PUSCH_r16                  `optional`
	ConfiguredGrantWithReTx_r16              *SharedSpectrumChAccessParamsPerBand_r16_configuredGrantWithReTx_r16              `optional`
	Ed_Threshold_r16                         *SharedSpectrumChAccessParamsPerBand_r16_ed_Threshold_r16                         `optional`
	Ul_DL_COT_Sharing_r16                    *SharedSpectrumChAccessParamsPerBand_r16_ul_DL_COT_Sharing_r16                    `optional`
	Mux_CG_UCI_HARQ_ACK_r16                  *SharedSpectrumChAccessParamsPerBand_r16_mux_CG_UCI_HARQ_ACK_r16                  `optional`
	Cg_resourceConfig_r16                    *SharedSpectrumChAccessParamsPerBand_r16_cg_resourceConfig_r16                    `optional`
}

func (ie *SharedSpectrumChAccessParamsPerBand_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ul_DynamicChAccess_r16 != nil, ie.Ul_Semi_StaticChAccess_r16 != nil, ie.Ssb_RRM_DynamicChAccess_r16 != nil, ie.Ssb_RRM_Semi_StaticChAccess_r16 != nil, ie.Mib_Acquisition_r16 != nil, ie.Ssb_RLM_DynamicChAccess_r16 != nil, ie.Ssb_RLM_Semi_StaticChAccess_r16 != nil, ie.Sib1_Acquisition_r16 != nil, ie.ExtRA_ResponseWindow_r16 != nil, ie.Ssb_BFD_CBD_dynamicChannelAccess_r16 != nil, ie.Ssb_BFD_CBD_semi_staticChannelAccess_r16 != nil, ie.Csi_RS_BFD_CBD_r16 != nil, ie.Ul_ChannelBW_SCell_10mhz_r16 != nil, ie.Rssi_ChannelOccupancyReporting_r16 != nil, ie.Srs_StartAnyOFDM_Symbol_r16 != nil, ie.SearchSpaceFreqMonitorLocation_r16 != nil, ie.Coreset_RB_Offset_r16 != nil, ie.Cgi_Acquisition_r16 != nil, ie.ConfiguredUL_Tx_r16 != nil, ie.Prach_Wideband_r16 != nil, ie.Dci_AvailableRB_Set_r16 != nil, ie.Dci_ChOccupancyDuration_r16 != nil, ie.TypeB_PDSCH_length_r16 != nil, ie.SearchSpaceSwitchWithDCI_r16 != nil, ie.SearchSpaceSwitchWithoutDCI_r16 != nil, ie.SearchSpaceSwitchCapability2_r16 != nil, ie.Non_numericalPDSCH_HARQ_timing_r16 != nil, ie.EnhancedDynamicHARQ_codebook_r16 != nil, ie.OneShotHARQ_feedback_r16 != nil, ie.MultiPUSCH_UL_grant_r16 != nil, ie.Csi_RS_RLM_r16 != nil, ie.Dummy != nil, ie.PeriodicAndSemi_PersistentCSI_RS_r16 != nil, ie.Pusch_PRB_interlace_r16 != nil, ie.Pucch_F0_F1_PRB_Interlace_r16 != nil, ie.Occ_PRB_PF2_PF3_r16 != nil, ie.ExtCP_rangeCG_PUSCH_r16 != nil, ie.ConfiguredGrantWithReTx_r16 != nil, ie.Ed_Threshold_r16 != nil, ie.Ul_DL_COT_Sharing_r16 != nil, ie.Mux_CG_UCI_HARQ_ACK_r16 != nil, ie.Cg_resourceConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ul_DynamicChAccess_r16 != nil {
		if err = ie.Ul_DynamicChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_DynamicChAccess_r16", err)
		}
	}
	if ie.Ul_Semi_StaticChAccess_r16 != nil {
		if err = ie.Ul_Semi_StaticChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_Semi_StaticChAccess_r16", err)
		}
	}
	if ie.Ssb_RRM_DynamicChAccess_r16 != nil {
		if err = ie.Ssb_RRM_DynamicChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_RRM_DynamicChAccess_r16", err)
		}
	}
	if ie.Ssb_RRM_Semi_StaticChAccess_r16 != nil {
		if err = ie.Ssb_RRM_Semi_StaticChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_RRM_Semi_StaticChAccess_r16", err)
		}
	}
	if ie.Mib_Acquisition_r16 != nil {
		if err = ie.Mib_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mib_Acquisition_r16", err)
		}
	}
	if ie.Ssb_RLM_DynamicChAccess_r16 != nil {
		if err = ie.Ssb_RLM_DynamicChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_RLM_DynamicChAccess_r16", err)
		}
	}
	if ie.Ssb_RLM_Semi_StaticChAccess_r16 != nil {
		if err = ie.Ssb_RLM_Semi_StaticChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_RLM_Semi_StaticChAccess_r16", err)
		}
	}
	if ie.Sib1_Acquisition_r16 != nil {
		if err = ie.Sib1_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sib1_Acquisition_r16", err)
		}
	}
	if ie.ExtRA_ResponseWindow_r16 != nil {
		if err = ie.ExtRA_ResponseWindow_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ExtRA_ResponseWindow_r16", err)
		}
	}
	if ie.Ssb_BFD_CBD_dynamicChannelAccess_r16 != nil {
		if err = ie.Ssb_BFD_CBD_dynamicChannelAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_BFD_CBD_dynamicChannelAccess_r16", err)
		}
	}
	if ie.Ssb_BFD_CBD_semi_staticChannelAccess_r16 != nil {
		if err = ie.Ssb_BFD_CBD_semi_staticChannelAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_BFD_CBD_semi_staticChannelAccess_r16", err)
		}
	}
	if ie.Csi_RS_BFD_CBD_r16 != nil {
		if err = ie.Csi_RS_BFD_CBD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RS_BFD_CBD_r16", err)
		}
	}
	if ie.Ul_ChannelBW_SCell_10mhz_r16 != nil {
		if err = ie.Ul_ChannelBW_SCell_10mhz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_ChannelBW_SCell_10mhz_r16", err)
		}
	}
	if ie.Rssi_ChannelOccupancyReporting_r16 != nil {
		if err = ie.Rssi_ChannelOccupancyReporting_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rssi_ChannelOccupancyReporting_r16", err)
		}
	}
	if ie.Srs_StartAnyOFDM_Symbol_r16 != nil {
		if err = ie.Srs_StartAnyOFDM_Symbol_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_StartAnyOFDM_Symbol_r16", err)
		}
	}
	if ie.SearchSpaceFreqMonitorLocation_r16 != nil {
		if err = w.WriteInteger(*ie.SearchSpaceFreqMonitorLocation_r16, &aper.Constraint{Lb: 1, Ub: 5}, false); err != nil {
			return utils.WrapError("Encode SearchSpaceFreqMonitorLocation_r16", err)
		}
	}
	if ie.Coreset_RB_Offset_r16 != nil {
		if err = ie.Coreset_RB_Offset_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Coreset_RB_Offset_r16", err)
		}
	}
	if ie.Cgi_Acquisition_r16 != nil {
		if err = ie.Cgi_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Cgi_Acquisition_r16", err)
		}
	}
	if ie.ConfiguredUL_Tx_r16 != nil {
		if err = ie.ConfiguredUL_Tx_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredUL_Tx_r16", err)
		}
	}
	if ie.Prach_Wideband_r16 != nil {
		if err = ie.Prach_Wideband_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Prach_Wideband_r16", err)
		}
	}
	if ie.Dci_AvailableRB_Set_r16 != nil {
		if err = ie.Dci_AvailableRB_Set_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_AvailableRB_Set_r16", err)
		}
	}
	if ie.Dci_ChOccupancyDuration_r16 != nil {
		if err = ie.Dci_ChOccupancyDuration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dci_ChOccupancyDuration_r16", err)
		}
	}
	if ie.TypeB_PDSCH_length_r16 != nil {
		if err = ie.TypeB_PDSCH_length_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TypeB_PDSCH_length_r16", err)
		}
	}
	if ie.SearchSpaceSwitchWithDCI_r16 != nil {
		if err = ie.SearchSpaceSwitchWithDCI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceSwitchWithDCI_r16", err)
		}
	}
	if ie.SearchSpaceSwitchWithoutDCI_r16 != nil {
		if err = ie.SearchSpaceSwitchWithoutDCI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceSwitchWithoutDCI_r16", err)
		}
	}
	if ie.SearchSpaceSwitchCapability2_r16 != nil {
		if err = ie.SearchSpaceSwitchCapability2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceSwitchCapability2_r16", err)
		}
	}
	if ie.Non_numericalPDSCH_HARQ_timing_r16 != nil {
		if err = ie.Non_numericalPDSCH_HARQ_timing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Non_numericalPDSCH_HARQ_timing_r16", err)
		}
	}
	if ie.EnhancedDynamicHARQ_codebook_r16 != nil {
		if err = ie.EnhancedDynamicHARQ_codebook_r16.Encode(w); err != nil {
			return utils.WrapError("Encode EnhancedDynamicHARQ_codebook_r16", err)
		}
	}
	if ie.OneShotHARQ_feedback_r16 != nil {
		if err = ie.OneShotHARQ_feedback_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OneShotHARQ_feedback_r16", err)
		}
	}
	if ie.MultiPUSCH_UL_grant_r16 != nil {
		if err = ie.MultiPUSCH_UL_grant_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MultiPUSCH_UL_grant_r16", err)
		}
	}
	if ie.Csi_RS_RLM_r16 != nil {
		if err = ie.Csi_RS_RLM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RS_RLM_r16", err)
		}
	}
	if ie.Dummy != nil {
		if err = ie.Dummy.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	if ie.PeriodicAndSemi_PersistentCSI_RS_r16 != nil {
		if err = ie.PeriodicAndSemi_PersistentCSI_RS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PeriodicAndSemi_PersistentCSI_RS_r16", err)
		}
	}
	if ie.Pusch_PRB_interlace_r16 != nil {
		if err = ie.Pusch_PRB_interlace_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_PRB_interlace_r16", err)
		}
	}
	if ie.Pucch_F0_F1_PRB_Interlace_r16 != nil {
		if err = ie.Pucch_F0_F1_PRB_Interlace_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_F0_F1_PRB_Interlace_r16", err)
		}
	}
	if ie.Occ_PRB_PF2_PF3_r16 != nil {
		if err = ie.Occ_PRB_PF2_PF3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Occ_PRB_PF2_PF3_r16", err)
		}
	}
	if ie.ExtCP_rangeCG_PUSCH_r16 != nil {
		if err = ie.ExtCP_rangeCG_PUSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ExtCP_rangeCG_PUSCH_r16", err)
		}
	}
	if ie.ConfiguredGrantWithReTx_r16 != nil {
		if err = ie.ConfiguredGrantWithReTx_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredGrantWithReTx_r16", err)
		}
	}
	if ie.Ed_Threshold_r16 != nil {
		if err = ie.Ed_Threshold_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ed_Threshold_r16", err)
		}
	}
	if ie.Ul_DL_COT_Sharing_r16 != nil {
		if err = ie.Ul_DL_COT_Sharing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_DL_COT_Sharing_r16", err)
		}
	}
	if ie.Mux_CG_UCI_HARQ_ACK_r16 != nil {
		if err = ie.Mux_CG_UCI_HARQ_ACK_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mux_CG_UCI_HARQ_ACK_r16", err)
		}
	}
	if ie.Cg_resourceConfig_r16 != nil {
		if err = ie.Cg_resourceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Cg_resourceConfig_r16", err)
		}
	}
	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_r16) Decode(r *aper.AperReader) error {
	var err error
	var Ul_DynamicChAccess_r16Present bool
	if Ul_DynamicChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_Semi_StaticChAccess_r16Present bool
	if Ul_Semi_StaticChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_RRM_DynamicChAccess_r16Present bool
	if Ssb_RRM_DynamicChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_RRM_Semi_StaticChAccess_r16Present bool
	if Ssb_RRM_Semi_StaticChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mib_Acquisition_r16Present bool
	if Mib_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_RLM_DynamicChAccess_r16Present bool
	if Ssb_RLM_DynamicChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_RLM_Semi_StaticChAccess_r16Present bool
	if Ssb_RLM_Semi_StaticChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sib1_Acquisition_r16Present bool
	if Sib1_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ExtRA_ResponseWindow_r16Present bool
	if ExtRA_ResponseWindow_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_BFD_CBD_dynamicChannelAccess_r16Present bool
	if Ssb_BFD_CBD_dynamicChannelAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_BFD_CBD_semi_staticChannelAccess_r16Present bool
	if Ssb_BFD_CBD_semi_staticChannelAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RS_BFD_CBD_r16Present bool
	if Csi_RS_BFD_CBD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_ChannelBW_SCell_10mhz_r16Present bool
	if Ul_ChannelBW_SCell_10mhz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rssi_ChannelOccupancyReporting_r16Present bool
	if Rssi_ChannelOccupancyReporting_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_StartAnyOFDM_Symbol_r16Present bool
	if Srs_StartAnyOFDM_Symbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceFreqMonitorLocation_r16Present bool
	if SearchSpaceFreqMonitorLocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Coreset_RB_Offset_r16Present bool
	if Coreset_RB_Offset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cgi_Acquisition_r16Present bool
	if Cgi_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredUL_Tx_r16Present bool
	if ConfiguredUL_Tx_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Prach_Wideband_r16Present bool
	if Prach_Wideband_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dci_AvailableRB_Set_r16Present bool
	if Dci_AvailableRB_Set_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dci_ChOccupancyDuration_r16Present bool
	if Dci_ChOccupancyDuration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TypeB_PDSCH_length_r16Present bool
	if TypeB_PDSCH_length_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceSwitchWithDCI_r16Present bool
	if SearchSpaceSwitchWithDCI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceSwitchWithoutDCI_r16Present bool
	if SearchSpaceSwitchWithoutDCI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceSwitchCapability2_r16Present bool
	if SearchSpaceSwitchCapability2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Non_numericalPDSCH_HARQ_timing_r16Present bool
	if Non_numericalPDSCH_HARQ_timing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var EnhancedDynamicHARQ_codebook_r16Present bool
	if EnhancedDynamicHARQ_codebook_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OneShotHARQ_feedback_r16Present bool
	if OneShotHARQ_feedback_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MultiPUSCH_UL_grant_r16Present bool
	if MultiPUSCH_UL_grant_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RS_RLM_r16Present bool
	if Csi_RS_RLM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PeriodicAndSemi_PersistentCSI_RS_r16Present bool
	if PeriodicAndSemi_PersistentCSI_RS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_PRB_interlace_r16Present bool
	if Pusch_PRB_interlace_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_F0_F1_PRB_Interlace_r16Present bool
	if Pucch_F0_F1_PRB_Interlace_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Occ_PRB_PF2_PF3_r16Present bool
	if Occ_PRB_PF2_PF3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ExtCP_rangeCG_PUSCH_r16Present bool
	if ExtCP_rangeCG_PUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredGrantWithReTx_r16Present bool
	if ConfiguredGrantWithReTx_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ed_Threshold_r16Present bool
	if Ed_Threshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_DL_COT_Sharing_r16Present bool
	if Ul_DL_COT_Sharing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mux_CG_UCI_HARQ_ACK_r16Present bool
	if Mux_CG_UCI_HARQ_ACK_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Cg_resourceConfig_r16Present bool
	if Cg_resourceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ul_DynamicChAccess_r16Present {
		ie.Ul_DynamicChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ul_DynamicChAccess_r16)
		if err = ie.Ul_DynamicChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_DynamicChAccess_r16", err)
		}
	}
	if Ul_Semi_StaticChAccess_r16Present {
		ie.Ul_Semi_StaticChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ul_Semi_StaticChAccess_r16)
		if err = ie.Ul_Semi_StaticChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_Semi_StaticChAccess_r16", err)
		}
	}
	if Ssb_RRM_DynamicChAccess_r16Present {
		ie.Ssb_RRM_DynamicChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_RRM_DynamicChAccess_r16)
		if err = ie.Ssb_RRM_DynamicChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_RRM_DynamicChAccess_r16", err)
		}
	}
	if Ssb_RRM_Semi_StaticChAccess_r16Present {
		ie.Ssb_RRM_Semi_StaticChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_RRM_Semi_StaticChAccess_r16)
		if err = ie.Ssb_RRM_Semi_StaticChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_RRM_Semi_StaticChAccess_r16", err)
		}
	}
	if Mib_Acquisition_r16Present {
		ie.Mib_Acquisition_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_mib_Acquisition_r16)
		if err = ie.Mib_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mib_Acquisition_r16", err)
		}
	}
	if Ssb_RLM_DynamicChAccess_r16Present {
		ie.Ssb_RLM_DynamicChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_RLM_DynamicChAccess_r16)
		if err = ie.Ssb_RLM_DynamicChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_RLM_DynamicChAccess_r16", err)
		}
	}
	if Ssb_RLM_Semi_StaticChAccess_r16Present {
		ie.Ssb_RLM_Semi_StaticChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_RLM_Semi_StaticChAccess_r16)
		if err = ie.Ssb_RLM_Semi_StaticChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_RLM_Semi_StaticChAccess_r16", err)
		}
	}
	if Sib1_Acquisition_r16Present {
		ie.Sib1_Acquisition_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_sib1_Acquisition_r16)
		if err = ie.Sib1_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sib1_Acquisition_r16", err)
		}
	}
	if ExtRA_ResponseWindow_r16Present {
		ie.ExtRA_ResponseWindow_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_extRA_ResponseWindow_r16)
		if err = ie.ExtRA_ResponseWindow_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ExtRA_ResponseWindow_r16", err)
		}
	}
	if Ssb_BFD_CBD_dynamicChannelAccess_r16Present {
		ie.Ssb_BFD_CBD_dynamicChannelAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_BFD_CBD_dynamicChannelAccess_r16)
		if err = ie.Ssb_BFD_CBD_dynamicChannelAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_BFD_CBD_dynamicChannelAccess_r16", err)
		}
	}
	if Ssb_BFD_CBD_semi_staticChannelAccess_r16Present {
		ie.Ssb_BFD_CBD_semi_staticChannelAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_BFD_CBD_semi_staticChannelAccess_r16)
		if err = ie.Ssb_BFD_CBD_semi_staticChannelAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_BFD_CBD_semi_staticChannelAccess_r16", err)
		}
	}
	if Csi_RS_BFD_CBD_r16Present {
		ie.Csi_RS_BFD_CBD_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_csi_RS_BFD_CBD_r16)
		if err = ie.Csi_RS_BFD_CBD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_BFD_CBD_r16", err)
		}
	}
	if Ul_ChannelBW_SCell_10mhz_r16Present {
		ie.Ul_ChannelBW_SCell_10mhz_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ul_ChannelBW_SCell_10mhz_r16)
		if err = ie.Ul_ChannelBW_SCell_10mhz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_ChannelBW_SCell_10mhz_r16", err)
		}
	}
	if Rssi_ChannelOccupancyReporting_r16Present {
		ie.Rssi_ChannelOccupancyReporting_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_rssi_ChannelOccupancyReporting_r16)
		if err = ie.Rssi_ChannelOccupancyReporting_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rssi_ChannelOccupancyReporting_r16", err)
		}
	}
	if Srs_StartAnyOFDM_Symbol_r16Present {
		ie.Srs_StartAnyOFDM_Symbol_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_srs_StartAnyOFDM_Symbol_r16)
		if err = ie.Srs_StartAnyOFDM_Symbol_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_StartAnyOFDM_Symbol_r16", err)
		}
	}
	if SearchSpaceFreqMonitorLocation_r16Present {
		var tmp_int_SearchSpaceFreqMonitorLocation_r16 int64
		if tmp_int_SearchSpaceFreqMonitorLocation_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 5}, false); err != nil {
			return utils.WrapError("Decode SearchSpaceFreqMonitorLocation_r16", err)
		}
		ie.SearchSpaceFreqMonitorLocation_r16 = &tmp_int_SearchSpaceFreqMonitorLocation_r16
	}
	if Coreset_RB_Offset_r16Present {
		ie.Coreset_RB_Offset_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_coreset_RB_Offset_r16)
		if err = ie.Coreset_RB_Offset_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Coreset_RB_Offset_r16", err)
		}
	}
	if Cgi_Acquisition_r16Present {
		ie.Cgi_Acquisition_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_cgi_Acquisition_r16)
		if err = ie.Cgi_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Cgi_Acquisition_r16", err)
		}
	}
	if ConfiguredUL_Tx_r16Present {
		ie.ConfiguredUL_Tx_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_configuredUL_Tx_r16)
		if err = ie.ConfiguredUL_Tx_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredUL_Tx_r16", err)
		}
	}
	if Prach_Wideband_r16Present {
		ie.Prach_Wideband_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_prach_Wideband_r16)
		if err = ie.Prach_Wideband_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Prach_Wideband_r16", err)
		}
	}
	if Dci_AvailableRB_Set_r16Present {
		ie.Dci_AvailableRB_Set_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_dci_AvailableRB_Set_r16)
		if err = ie.Dci_AvailableRB_Set_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_AvailableRB_Set_r16", err)
		}
	}
	if Dci_ChOccupancyDuration_r16Present {
		ie.Dci_ChOccupancyDuration_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_dci_ChOccupancyDuration_r16)
		if err = ie.Dci_ChOccupancyDuration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dci_ChOccupancyDuration_r16", err)
		}
	}
	if TypeB_PDSCH_length_r16Present {
		ie.TypeB_PDSCH_length_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_typeB_PDSCH_length_r16)
		if err = ie.TypeB_PDSCH_length_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TypeB_PDSCH_length_r16", err)
		}
	}
	if SearchSpaceSwitchWithDCI_r16Present {
		ie.SearchSpaceSwitchWithDCI_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithDCI_r16)
		if err = ie.SearchSpaceSwitchWithDCI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceSwitchWithDCI_r16", err)
		}
	}
	if SearchSpaceSwitchWithoutDCI_r16Present {
		ie.SearchSpaceSwitchWithoutDCI_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithoutDCI_r16)
		if err = ie.SearchSpaceSwitchWithoutDCI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceSwitchWithoutDCI_r16", err)
		}
	}
	if SearchSpaceSwitchCapability2_r16Present {
		ie.SearchSpaceSwitchCapability2_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchCapability2_r16)
		if err = ie.SearchSpaceSwitchCapability2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceSwitchCapability2_r16", err)
		}
	}
	if Non_numericalPDSCH_HARQ_timing_r16Present {
		ie.Non_numericalPDSCH_HARQ_timing_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_non_numericalPDSCH_HARQ_timing_r16)
		if err = ie.Non_numericalPDSCH_HARQ_timing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Non_numericalPDSCH_HARQ_timing_r16", err)
		}
	}
	if EnhancedDynamicHARQ_codebook_r16Present {
		ie.EnhancedDynamicHARQ_codebook_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_enhancedDynamicHARQ_codebook_r16)
		if err = ie.EnhancedDynamicHARQ_codebook_r16.Decode(r); err != nil {
			return utils.WrapError("Decode EnhancedDynamicHARQ_codebook_r16", err)
		}
	}
	if OneShotHARQ_feedback_r16Present {
		ie.OneShotHARQ_feedback_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_oneShotHARQ_feedback_r16)
		if err = ie.OneShotHARQ_feedback_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OneShotHARQ_feedback_r16", err)
		}
	}
	if MultiPUSCH_UL_grant_r16Present {
		ie.MultiPUSCH_UL_grant_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_multiPUSCH_UL_grant_r16)
		if err = ie.MultiPUSCH_UL_grant_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MultiPUSCH_UL_grant_r16", err)
		}
	}
	if Csi_RS_RLM_r16Present {
		ie.Csi_RS_RLM_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_csi_RS_RLM_r16)
		if err = ie.Csi_RS_RLM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_RLM_r16", err)
		}
	}
	if DummyPresent {
		ie.Dummy = new(SharedSpectrumChAccessParamsPerBand_r16_dummy)
		if err = ie.Dummy.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
	}
	if PeriodicAndSemi_PersistentCSI_RS_r16Present {
		ie.PeriodicAndSemi_PersistentCSI_RS_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_periodicAndSemi_PersistentCSI_RS_r16)
		if err = ie.PeriodicAndSemi_PersistentCSI_RS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicAndSemi_PersistentCSI_RS_r16", err)
		}
	}
	if Pusch_PRB_interlace_r16Present {
		ie.Pusch_PRB_interlace_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_pusch_PRB_interlace_r16)
		if err = ie.Pusch_PRB_interlace_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_PRB_interlace_r16", err)
		}
	}
	if Pucch_F0_F1_PRB_Interlace_r16Present {
		ie.Pucch_F0_F1_PRB_Interlace_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_pucch_F0_F1_PRB_Interlace_r16)
		if err = ie.Pucch_F0_F1_PRB_Interlace_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_F0_F1_PRB_Interlace_r16", err)
		}
	}
	if Occ_PRB_PF2_PF3_r16Present {
		ie.Occ_PRB_PF2_PF3_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_occ_PRB_PF2_PF3_r16)
		if err = ie.Occ_PRB_PF2_PF3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Occ_PRB_PF2_PF3_r16", err)
		}
	}
	if ExtCP_rangeCG_PUSCH_r16Present {
		ie.ExtCP_rangeCG_PUSCH_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_extCP_rangeCG_PUSCH_r16)
		if err = ie.ExtCP_rangeCG_PUSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ExtCP_rangeCG_PUSCH_r16", err)
		}
	}
	if ConfiguredGrantWithReTx_r16Present {
		ie.ConfiguredGrantWithReTx_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_configuredGrantWithReTx_r16)
		if err = ie.ConfiguredGrantWithReTx_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredGrantWithReTx_r16", err)
		}
	}
	if Ed_Threshold_r16Present {
		ie.Ed_Threshold_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ed_Threshold_r16)
		if err = ie.Ed_Threshold_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ed_Threshold_r16", err)
		}
	}
	if Ul_DL_COT_Sharing_r16Present {
		ie.Ul_DL_COT_Sharing_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ul_DL_COT_Sharing_r16)
		if err = ie.Ul_DL_COT_Sharing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_DL_COT_Sharing_r16", err)
		}
	}
	if Mux_CG_UCI_HARQ_ACK_r16Present {
		ie.Mux_CG_UCI_HARQ_ACK_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_mux_CG_UCI_HARQ_ACK_r16)
		if err = ie.Mux_CG_UCI_HARQ_ACK_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mux_CG_UCI_HARQ_ACK_r16", err)
		}
	}
	if Cg_resourceConfig_r16Present {
		ie.Cg_resourceConfig_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_cg_resourceConfig_r16)
		if err = ie.Cg_resourceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Cg_resourceConfig_r16", err)
		}
	}
	return nil
}
