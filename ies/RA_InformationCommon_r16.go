package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RA_InformationCommon_r16 struct {
	AbsoluteFrequencyPointA_r16                    ARFCN_ValueNR                                                            `madatory`
	LocationAndBandwidth_r16                       int64                                                                    `lb:0,ub:37949,madatory`
	SubcarrierSpacing_r16                          SubcarrierSpacing                                                        `madatory`
	Msg1_FrequencyStart_r16                        *int64                                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional`
	Msg1_FrequencyStartCFRA_r16                    *int64                                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional`
	Msg1_SubcarrierSpacing_r16                     *SubcarrierSpacing                                                       `optional`
	Msg1_SubcarrierSpacingCFRA_r16                 *SubcarrierSpacing                                                       `optional`
	Msg1_FDM_r16                                   *RA_InformationCommon_r16_msg1_FDM_r16                                   `optional`
	Msg1_FDMCFRA_r16                               *RA_InformationCommon_r16_msg1_FDMCFRA_r16                               `optional`
	PerRAInfoList_r16                              PerRAInfoList_r16                                                        `madatory`
	PerRAInfoList_v1660                            *PerRAInfoList_v1660                                                     `optional,ext-1`
	Msg1_SCS_From_prach_ConfigurationIndex_r16     *RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndex_r16     `optional,ext-2`
	Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 *RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 `optional,ext-3`
	MsgA_RO_FrequencyStart_r17                     *int64                                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional,ext-4`
	MsgA_RO_FrequencyStartCFRA_r17                 *int64                                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional,ext-4`
	MsgA_SubcarrierSpacing_r17                     *SubcarrierSpacing                                                       `optional,ext-4`
	MsgA_RO_FDM_r17                                *RA_InformationCommon_r16_msgA_RO_FDM_r17                                `optional,ext-4`
	MsgA_RO_FDMCFRA_r17                            *RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17                            `optional,ext-4`
	MsgA_SCS_From_prach_ConfigurationIndex_r17     *RA_InformationCommon_r16_msgA_SCS_From_prach_ConfigurationIndex_r17     `optional,ext-4`
	MsgA_TransMax_r17                              *RA_InformationCommon_r16_msgA_TransMax_r17                              `optional,ext-4`
	MsgA_MCS_r17                                   *int64                                                                   `lb:0,ub:15,optional,ext-4`
	NrofPRBs_PerMsgA_PO_r17                        *int64                                                                   `lb:1,ub:32,optional,ext-4`
	MsgA_PUSCH_TimeDomainAllocation_r17            *int64                                                                   `lb:1,ub:maxNrofUL_Allocations,optional,ext-4`
	FrequencyStartMsgA_PUSCH_r17                   *int64                                                                   `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional,ext-4`
	NrofMsgA_PO_FDM_r17                            *RA_InformationCommon_r16_nrofMsgA_PO_FDM_r17                            `optional,ext-4`
	DlPathlossRSRP_r17                             *RSRP_Range                                                              `optional,ext-4`
	IntendedSIBs_r17                               []SIB_Type_r17                                                           `lb:1,ub:maxSIB,optional,ext-4`
	SsbsForSI_Acquisition_r17                      []SSB_Index                                                              `lb:1,ub:maxNrofSSBs_r16,optional,ext-4`
	MsgA_PUSCH_PayloadSize_r17                     *aper.BitString                                                          `lb:5,ub:5,optional,ext-4`
	OnDemandSISuccess_r17                          *RA_InformationCommon_r16_onDemandSISuccess_r17                          `optional,ext-4`
}

func (ie *RA_InformationCommon_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.PerRAInfoList_v1660 != nil || ie.Msg1_SCS_From_prach_ConfigurationIndex_r16 != nil || ie.Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 != nil || ie.MsgA_RO_FrequencyStart_r17 != nil || ie.MsgA_RO_FrequencyStartCFRA_r17 != nil || ie.MsgA_SubcarrierSpacing_r17 != nil || ie.MsgA_RO_FDM_r17 != nil || ie.MsgA_RO_FDMCFRA_r17 != nil || ie.MsgA_SCS_From_prach_ConfigurationIndex_r17 != nil || ie.MsgA_TransMax_r17 != nil || ie.MsgA_MCS_r17 != nil || ie.NrofPRBs_PerMsgA_PO_r17 != nil || ie.MsgA_PUSCH_TimeDomainAllocation_r17 != nil || ie.FrequencyStartMsgA_PUSCH_r17 != nil || ie.NrofMsgA_PO_FDM_r17 != nil || ie.DlPathlossRSRP_r17 != nil || len(ie.IntendedSIBs_r17) > 0 || len(ie.SsbsForSI_Acquisition_r17) > 0 || ie.MsgA_PUSCH_PayloadSize_r17 != nil || ie.OnDemandSISuccess_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Msg1_FrequencyStart_r16 != nil, ie.Msg1_FrequencyStartCFRA_r16 != nil, ie.Msg1_SubcarrierSpacing_r16 != nil, ie.Msg1_SubcarrierSpacingCFRA_r16 != nil, ie.Msg1_FDM_r16 != nil, ie.Msg1_FDMCFRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.AbsoluteFrequencyPointA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode AbsoluteFrequencyPointA_r16", err)
	}
	if err = w.WriteInteger(ie.LocationAndBandwidth_r16, &aper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("WriteInteger LocationAndBandwidth_r16", err)
	}
	if err = ie.SubcarrierSpacing_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SubcarrierSpacing_r16", err)
	}
	if ie.Msg1_FrequencyStart_r16 != nil {
		if err = w.WriteInteger(*ie.Msg1_FrequencyStart_r16, &aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Encode Msg1_FrequencyStart_r16", err)
		}
	}
	if ie.Msg1_FrequencyStartCFRA_r16 != nil {
		if err = w.WriteInteger(*ie.Msg1_FrequencyStartCFRA_r16, &aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Encode Msg1_FrequencyStartCFRA_r16", err)
		}
	}
	if ie.Msg1_SubcarrierSpacing_r16 != nil {
		if err = ie.Msg1_SubcarrierSpacing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Msg1_SubcarrierSpacing_r16", err)
		}
	}
	if ie.Msg1_SubcarrierSpacingCFRA_r16 != nil {
		if err = ie.Msg1_SubcarrierSpacingCFRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Msg1_SubcarrierSpacingCFRA_r16", err)
		}
	}
	if ie.Msg1_FDM_r16 != nil {
		if err = ie.Msg1_FDM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Msg1_FDM_r16", err)
		}
	}
	if ie.Msg1_FDMCFRA_r16 != nil {
		if err = ie.Msg1_FDMCFRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Msg1_FDMCFRA_r16", err)
		}
	}
	if err = ie.PerRAInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PerRAInfoList_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.PerRAInfoList_v1660 != nil, ie.Msg1_SCS_From_prach_ConfigurationIndex_r16 != nil, ie.Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 != nil, ie.MsgA_RO_FrequencyStart_r17 != nil || ie.MsgA_RO_FrequencyStartCFRA_r17 != nil || ie.MsgA_SubcarrierSpacing_r17 != nil || ie.MsgA_RO_FDM_r17 != nil || ie.MsgA_RO_FDMCFRA_r17 != nil || ie.MsgA_SCS_From_prach_ConfigurationIndex_r17 != nil || ie.MsgA_TransMax_r17 != nil || ie.MsgA_MCS_r17 != nil || ie.NrofPRBs_PerMsgA_PO_r17 != nil || ie.MsgA_PUSCH_TimeDomainAllocation_r17 != nil || ie.FrequencyStartMsgA_PUSCH_r17 != nil || ie.NrofMsgA_PO_FDM_r17 != nil || ie.DlPathlossRSRP_r17 != nil || len(ie.IntendedSIBs_r17) > 0 || len(ie.SsbsForSI_Acquisition_r17) > 0 || ie.MsgA_PUSCH_PayloadSize_r17 != nil || ie.OnDemandSISuccess_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RA_InformationCommon_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.PerRAInfoList_v1660 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PerRAInfoList_v1660 optional
			if ie.PerRAInfoList_v1660 != nil {
				if err = ie.PerRAInfoList_v1660.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PerRAInfoList_v1660", err)
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
			optionals_ext_2 := []bool{ie.Msg1_SCS_From_prach_ConfigurationIndex_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Msg1_SCS_From_prach_ConfigurationIndex_r16 optional
			if ie.Msg1_SCS_From_prach_ConfigurationIndex_r16 != nil {
				if err = ie.Msg1_SCS_From_prach_ConfigurationIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Msg1_SCS_From_prach_ConfigurationIndex_r16", err)
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
			optionals_ext_3 := []bool{ie.Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 optional
			if ie.Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 != nil {
				if err = ie.Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16", err)
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
			optionals_ext_4 := []bool{ie.MsgA_RO_FrequencyStart_r17 != nil, ie.MsgA_RO_FrequencyStartCFRA_r17 != nil, ie.MsgA_SubcarrierSpacing_r17 != nil, ie.MsgA_RO_FDM_r17 != nil, ie.MsgA_RO_FDMCFRA_r17 != nil, ie.MsgA_SCS_From_prach_ConfigurationIndex_r17 != nil, ie.MsgA_TransMax_r17 != nil, ie.MsgA_MCS_r17 != nil, ie.NrofPRBs_PerMsgA_PO_r17 != nil, ie.MsgA_PUSCH_TimeDomainAllocation_r17 != nil, ie.FrequencyStartMsgA_PUSCH_r17 != nil, ie.NrofMsgA_PO_FDM_r17 != nil, ie.DlPathlossRSRP_r17 != nil, len(ie.IntendedSIBs_r17) > 0, len(ie.SsbsForSI_Acquisition_r17) > 0, ie.MsgA_PUSCH_PayloadSize_r17 != nil, ie.OnDemandSISuccess_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MsgA_RO_FrequencyStart_r17 optional
			if ie.MsgA_RO_FrequencyStart_r17 != nil {
				if err = extWriter.WriteInteger(*ie.MsgA_RO_FrequencyStart_r17, &aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Encode MsgA_RO_FrequencyStart_r17", err)
				}
			}
			// encode MsgA_RO_FrequencyStartCFRA_r17 optional
			if ie.MsgA_RO_FrequencyStartCFRA_r17 != nil {
				if err = extWriter.WriteInteger(*ie.MsgA_RO_FrequencyStartCFRA_r17, &aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Encode MsgA_RO_FrequencyStartCFRA_r17", err)
				}
			}
			// encode MsgA_SubcarrierSpacing_r17 optional
			if ie.MsgA_SubcarrierSpacing_r17 != nil {
				if err = ie.MsgA_SubcarrierSpacing_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MsgA_SubcarrierSpacing_r17", err)
				}
			}
			// encode MsgA_RO_FDM_r17 optional
			if ie.MsgA_RO_FDM_r17 != nil {
				if err = ie.MsgA_RO_FDM_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MsgA_RO_FDM_r17", err)
				}
			}
			// encode MsgA_RO_FDMCFRA_r17 optional
			if ie.MsgA_RO_FDMCFRA_r17 != nil {
				if err = ie.MsgA_RO_FDMCFRA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MsgA_RO_FDMCFRA_r17", err)
				}
			}
			// encode MsgA_SCS_From_prach_ConfigurationIndex_r17 optional
			if ie.MsgA_SCS_From_prach_ConfigurationIndex_r17 != nil {
				if err = ie.MsgA_SCS_From_prach_ConfigurationIndex_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MsgA_SCS_From_prach_ConfigurationIndex_r17", err)
				}
			}
			// encode MsgA_TransMax_r17 optional
			if ie.MsgA_TransMax_r17 != nil {
				if err = ie.MsgA_TransMax_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MsgA_TransMax_r17", err)
				}
			}
			// encode MsgA_MCS_r17 optional
			if ie.MsgA_MCS_r17 != nil {
				if err = extWriter.WriteInteger(*ie.MsgA_MCS_r17, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode MsgA_MCS_r17", err)
				}
			}
			// encode NrofPRBs_PerMsgA_PO_r17 optional
			if ie.NrofPRBs_PerMsgA_PO_r17 != nil {
				if err = extWriter.WriteInteger(*ie.NrofPRBs_PerMsgA_PO_r17, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode NrofPRBs_PerMsgA_PO_r17", err)
				}
			}
			// encode MsgA_PUSCH_TimeDomainAllocation_r17 optional
			if ie.MsgA_PUSCH_TimeDomainAllocation_r17 != nil {
				if err = extWriter.WriteInteger(*ie.MsgA_PUSCH_TimeDomainAllocation_r17, &aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false); err != nil {
					return utils.WrapError("Encode MsgA_PUSCH_TimeDomainAllocation_r17", err)
				}
			}
			// encode FrequencyStartMsgA_PUSCH_r17 optional
			if ie.FrequencyStartMsgA_PUSCH_r17 != nil {
				if err = extWriter.WriteInteger(*ie.FrequencyStartMsgA_PUSCH_r17, &aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Encode FrequencyStartMsgA_PUSCH_r17", err)
				}
			}
			// encode NrofMsgA_PO_FDM_r17 optional
			if ie.NrofMsgA_PO_FDM_r17 != nil {
				if err = ie.NrofMsgA_PO_FDM_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NrofMsgA_PO_FDM_r17", err)
				}
			}
			// encode DlPathlossRSRP_r17 optional
			if ie.DlPathlossRSRP_r17 != nil {
				if err = ie.DlPathlossRSRP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DlPathlossRSRP_r17", err)
				}
			}
			// encode IntendedSIBs_r17 optional
			if len(ie.IntendedSIBs_r17) > 0 {
				tmp_IntendedSIBs_r17 := utils.NewSequence[*SIB_Type_r17]([]*SIB_Type_r17{}, aper.Constraint{Lb: 1, Ub: maxSIB}, false)
				for _, i := range ie.IntendedSIBs_r17 {
					tmp_IntendedSIBs_r17.Value = append(tmp_IntendedSIBs_r17.Value, &i)
				}
				if err = tmp_IntendedSIBs_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntendedSIBs_r17", err)
				}
			}
			// encode SsbsForSI_Acquisition_r17 optional
			if len(ie.SsbsForSI_Acquisition_r17) > 0 {
				tmp_SsbsForSI_Acquisition_r17 := utils.NewSequence[*SSB_Index]([]*SSB_Index{}, aper.Constraint{Lb: 1, Ub: maxNrofSSBs_r16}, false)
				for _, i := range ie.SsbsForSI_Acquisition_r17 {
					tmp_SsbsForSI_Acquisition_r17.Value = append(tmp_SsbsForSI_Acquisition_r17.Value, &i)
				}
				if err = tmp_SsbsForSI_Acquisition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SsbsForSI_Acquisition_r17", err)
				}
			}
			// encode MsgA_PUSCH_PayloadSize_r17 optional
			if ie.MsgA_PUSCH_PayloadSize_r17 != nil {
				if err = extWriter.WriteBitString(ie.MsgA_PUSCH_PayloadSize_r17.Bytes, uint(ie.MsgA_PUSCH_PayloadSize_r17.NumBits), &aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode MsgA_PUSCH_PayloadSize_r17", err)
				}
			}
			// encode OnDemandSISuccess_r17 optional
			if ie.OnDemandSISuccess_r17 != nil {
				if err = ie.OnDemandSISuccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode OnDemandSISuccess_r17", err)
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

func (ie *RA_InformationCommon_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg1_FrequencyStart_r16Present bool
	if Msg1_FrequencyStart_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg1_FrequencyStartCFRA_r16Present bool
	if Msg1_FrequencyStartCFRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg1_SubcarrierSpacing_r16Present bool
	if Msg1_SubcarrierSpacing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg1_SubcarrierSpacingCFRA_r16Present bool
	if Msg1_SubcarrierSpacingCFRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg1_FDM_r16Present bool
	if Msg1_FDM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg1_FDMCFRA_r16Present bool
	if Msg1_FDMCFRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.AbsoluteFrequencyPointA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode AbsoluteFrequencyPointA_r16", err)
	}
	var tmp_int_LocationAndBandwidth_r16 int64
	if tmp_int_LocationAndBandwidth_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("ReadInteger LocationAndBandwidth_r16", err)
	}
	ie.LocationAndBandwidth_r16 = tmp_int_LocationAndBandwidth_r16
	if err = ie.SubcarrierSpacing_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SubcarrierSpacing_r16", err)
	}
	if Msg1_FrequencyStart_r16Present {
		var tmp_int_Msg1_FrequencyStart_r16 int64
		if tmp_int_Msg1_FrequencyStart_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Decode Msg1_FrequencyStart_r16", err)
		}
		ie.Msg1_FrequencyStart_r16 = &tmp_int_Msg1_FrequencyStart_r16
	}
	if Msg1_FrequencyStartCFRA_r16Present {
		var tmp_int_Msg1_FrequencyStartCFRA_r16 int64
		if tmp_int_Msg1_FrequencyStartCFRA_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Decode Msg1_FrequencyStartCFRA_r16", err)
		}
		ie.Msg1_FrequencyStartCFRA_r16 = &tmp_int_Msg1_FrequencyStartCFRA_r16
	}
	if Msg1_SubcarrierSpacing_r16Present {
		ie.Msg1_SubcarrierSpacing_r16 = new(SubcarrierSpacing)
		if err = ie.Msg1_SubcarrierSpacing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Msg1_SubcarrierSpacing_r16", err)
		}
	}
	if Msg1_SubcarrierSpacingCFRA_r16Present {
		ie.Msg1_SubcarrierSpacingCFRA_r16 = new(SubcarrierSpacing)
		if err = ie.Msg1_SubcarrierSpacingCFRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Msg1_SubcarrierSpacingCFRA_r16", err)
		}
	}
	if Msg1_FDM_r16Present {
		ie.Msg1_FDM_r16 = new(RA_InformationCommon_r16_msg1_FDM_r16)
		if err = ie.Msg1_FDM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Msg1_FDM_r16", err)
		}
	}
	if Msg1_FDMCFRA_r16Present {
		ie.Msg1_FDMCFRA_r16 = new(RA_InformationCommon_r16_msg1_FDMCFRA_r16)
		if err = ie.Msg1_FDMCFRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Msg1_FDMCFRA_r16", err)
		}
	}
	if err = ie.PerRAInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PerRAInfoList_r16", err)
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
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

			PerRAInfoList_v1660Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PerRAInfoList_v1660 optional
			if PerRAInfoList_v1660Present {
				ie.PerRAInfoList_v1660 = new(PerRAInfoList_v1660)
				if err = ie.PerRAInfoList_v1660.Decode(extReader); err != nil {
					return utils.WrapError("Decode PerRAInfoList_v1660", err)
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

			Msg1_SCS_From_prach_ConfigurationIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Msg1_SCS_From_prach_ConfigurationIndex_r16 optional
			if Msg1_SCS_From_prach_ConfigurationIndex_r16Present {
				ie.Msg1_SCS_From_prach_ConfigurationIndex_r16 = new(RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndex_r16)
				if err = ie.Msg1_SCS_From_prach_ConfigurationIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Msg1_SCS_From_prach_ConfigurationIndex_r16", err)
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

			Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 optional
			if Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16Present {
				ie.Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 = new(RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16)
				if err = ie.Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Msg1_SCS_From_prach_ConfigurationIndexCFRA_r16", err)
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

			MsgA_RO_FrequencyStart_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MsgA_RO_FrequencyStartCFRA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MsgA_SubcarrierSpacing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MsgA_RO_FDM_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MsgA_RO_FDMCFRA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MsgA_SCS_From_prach_ConfigurationIndex_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MsgA_TransMax_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MsgA_MCS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NrofPRBs_PerMsgA_PO_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MsgA_PUSCH_TimeDomainAllocation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FrequencyStartMsgA_PUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NrofMsgA_PO_FDM_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DlPathlossRSRP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IntendedSIBs_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SsbsForSI_Acquisition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MsgA_PUSCH_PayloadSize_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			OnDemandSISuccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MsgA_RO_FrequencyStart_r17 optional
			if MsgA_RO_FrequencyStart_r17Present {
				var tmp_int_MsgA_RO_FrequencyStart_r17 int64
				if tmp_int_MsgA_RO_FrequencyStart_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Decode MsgA_RO_FrequencyStart_r17", err)
				}
				ie.MsgA_RO_FrequencyStart_r17 = &tmp_int_MsgA_RO_FrequencyStart_r17
			}
			// decode MsgA_RO_FrequencyStartCFRA_r17 optional
			if MsgA_RO_FrequencyStartCFRA_r17Present {
				var tmp_int_MsgA_RO_FrequencyStartCFRA_r17 int64
				if tmp_int_MsgA_RO_FrequencyStartCFRA_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Decode MsgA_RO_FrequencyStartCFRA_r17", err)
				}
				ie.MsgA_RO_FrequencyStartCFRA_r17 = &tmp_int_MsgA_RO_FrequencyStartCFRA_r17
			}
			// decode MsgA_SubcarrierSpacing_r17 optional
			if MsgA_SubcarrierSpacing_r17Present {
				ie.MsgA_SubcarrierSpacing_r17 = new(SubcarrierSpacing)
				if err = ie.MsgA_SubcarrierSpacing_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MsgA_SubcarrierSpacing_r17", err)
				}
			}
			// decode MsgA_RO_FDM_r17 optional
			if MsgA_RO_FDM_r17Present {
				ie.MsgA_RO_FDM_r17 = new(RA_InformationCommon_r16_msgA_RO_FDM_r17)
				if err = ie.MsgA_RO_FDM_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MsgA_RO_FDM_r17", err)
				}
			}
			// decode MsgA_RO_FDMCFRA_r17 optional
			if MsgA_RO_FDMCFRA_r17Present {
				ie.MsgA_RO_FDMCFRA_r17 = new(RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17)
				if err = ie.MsgA_RO_FDMCFRA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MsgA_RO_FDMCFRA_r17", err)
				}
			}
			// decode MsgA_SCS_From_prach_ConfigurationIndex_r17 optional
			if MsgA_SCS_From_prach_ConfigurationIndex_r17Present {
				ie.MsgA_SCS_From_prach_ConfigurationIndex_r17 = new(RA_InformationCommon_r16_msgA_SCS_From_prach_ConfigurationIndex_r17)
				if err = ie.MsgA_SCS_From_prach_ConfigurationIndex_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MsgA_SCS_From_prach_ConfigurationIndex_r17", err)
				}
			}
			// decode MsgA_TransMax_r17 optional
			if MsgA_TransMax_r17Present {
				ie.MsgA_TransMax_r17 = new(RA_InformationCommon_r16_msgA_TransMax_r17)
				if err = ie.MsgA_TransMax_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MsgA_TransMax_r17", err)
				}
			}
			// decode MsgA_MCS_r17 optional
			if MsgA_MCS_r17Present {
				var tmp_int_MsgA_MCS_r17 int64
				if tmp_int_MsgA_MCS_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode MsgA_MCS_r17", err)
				}
				ie.MsgA_MCS_r17 = &tmp_int_MsgA_MCS_r17
			}
			// decode NrofPRBs_PerMsgA_PO_r17 optional
			if NrofPRBs_PerMsgA_PO_r17Present {
				var tmp_int_NrofPRBs_PerMsgA_PO_r17 int64
				if tmp_int_NrofPRBs_PerMsgA_PO_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode NrofPRBs_PerMsgA_PO_r17", err)
				}
				ie.NrofPRBs_PerMsgA_PO_r17 = &tmp_int_NrofPRBs_PerMsgA_PO_r17
			}
			// decode MsgA_PUSCH_TimeDomainAllocation_r17 optional
			if MsgA_PUSCH_TimeDomainAllocation_r17Present {
				var tmp_int_MsgA_PUSCH_TimeDomainAllocation_r17 int64
				if tmp_int_MsgA_PUSCH_TimeDomainAllocation_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false); err != nil {
					return utils.WrapError("Decode MsgA_PUSCH_TimeDomainAllocation_r17", err)
				}
				ie.MsgA_PUSCH_TimeDomainAllocation_r17 = &tmp_int_MsgA_PUSCH_TimeDomainAllocation_r17
			}
			// decode FrequencyStartMsgA_PUSCH_r17 optional
			if FrequencyStartMsgA_PUSCH_r17Present {
				var tmp_int_FrequencyStartMsgA_PUSCH_r17 int64
				if tmp_int_FrequencyStartMsgA_PUSCH_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
					return utils.WrapError("Decode FrequencyStartMsgA_PUSCH_r17", err)
				}
				ie.FrequencyStartMsgA_PUSCH_r17 = &tmp_int_FrequencyStartMsgA_PUSCH_r17
			}
			// decode NrofMsgA_PO_FDM_r17 optional
			if NrofMsgA_PO_FDM_r17Present {
				ie.NrofMsgA_PO_FDM_r17 = new(RA_InformationCommon_r16_nrofMsgA_PO_FDM_r17)
				if err = ie.NrofMsgA_PO_FDM_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode NrofMsgA_PO_FDM_r17", err)
				}
			}
			// decode DlPathlossRSRP_r17 optional
			if DlPathlossRSRP_r17Present {
				ie.DlPathlossRSRP_r17 = new(RSRP_Range)
				if err = ie.DlPathlossRSRP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DlPathlossRSRP_r17", err)
				}
			}
			// decode IntendedSIBs_r17 optional
			if IntendedSIBs_r17Present {
				tmp_IntendedSIBs_r17 := utils.NewSequence[*SIB_Type_r17]([]*SIB_Type_r17{}, aper.Constraint{Lb: 1, Ub: maxSIB}, false)
				fn_IntendedSIBs_r17 := func() *SIB_Type_r17 {
					return new(SIB_Type_r17)
				}
				if err = tmp_IntendedSIBs_r17.Decode(extReader, fn_IntendedSIBs_r17); err != nil {
					return utils.WrapError("Decode IntendedSIBs_r17", err)
				}
				ie.IntendedSIBs_r17 = []SIB_Type_r17{}
				for _, i := range tmp_IntendedSIBs_r17.Value {
					ie.IntendedSIBs_r17 = append(ie.IntendedSIBs_r17, *i)
				}
			}
			// decode SsbsForSI_Acquisition_r17 optional
			if SsbsForSI_Acquisition_r17Present {
				tmp_SsbsForSI_Acquisition_r17 := utils.NewSequence[*SSB_Index]([]*SSB_Index{}, aper.Constraint{Lb: 1, Ub: maxNrofSSBs_r16}, false)
				fn_SsbsForSI_Acquisition_r17 := func() *SSB_Index {
					return new(SSB_Index)
				}
				if err = tmp_SsbsForSI_Acquisition_r17.Decode(extReader, fn_SsbsForSI_Acquisition_r17); err != nil {
					return utils.WrapError("Decode SsbsForSI_Acquisition_r17", err)
				}
				ie.SsbsForSI_Acquisition_r17 = []SSB_Index{}
				for _, i := range tmp_SsbsForSI_Acquisition_r17.Value {
					ie.SsbsForSI_Acquisition_r17 = append(ie.SsbsForSI_Acquisition_r17, *i)
				}
			}
			// decode MsgA_PUSCH_PayloadSize_r17 optional
			if MsgA_PUSCH_PayloadSize_r17Present {
				var tmp_bs_MsgA_PUSCH_PayloadSize_r17 []byte
				var n_MsgA_PUSCH_PayloadSize_r17 uint
				if tmp_bs_MsgA_PUSCH_PayloadSize_r17, n_MsgA_PUSCH_PayloadSize_r17, err = extReader.ReadBitString(&aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode MsgA_PUSCH_PayloadSize_r17", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_MsgA_PUSCH_PayloadSize_r17,
					NumBits: uint64(n_MsgA_PUSCH_PayloadSize_r17),
				}
				ie.MsgA_PUSCH_PayloadSize_r17 = &tmp_bitstring
			}
			// decode OnDemandSISuccess_r17 optional
			if OnDemandSISuccess_r17Present {
				ie.OnDemandSISuccess_r17 = new(RA_InformationCommon_r16_onDemandSISuccess_r17)
				if err = ie.OnDemandSISuccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode OnDemandSISuccess_r17", err)
				}
			}
		}
	}
	return nil
}
