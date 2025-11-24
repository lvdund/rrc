package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16 struct {
	FreqBandSidelink_r16                             FreqBandIndicatorNR                                                `madatory`
	Sl_Reception_r16                                 *BandSidelink_r16_sl_Reception_r16                                 `optional`
	Sl_TransmissionMode1_r16                         *BandSidelink_r16_sl_TransmissionMode1_r16                         `optional`
	Sync_Sidelink_r16                                *BandSidelink_r16_sync_Sidelink_r16                                `optional`
	Sl_Tx_256QAM_r16                                 *BandSidelink_r16_sl_Tx_256QAM_r16                                 `optional`
	Psfch_FormatZeroSidelink_r16                     *BandSidelink_r16_psfch_FormatZeroSidelink_r16                     `optional`
	LowSE_64QAM_MCS_TableSidelink_r16                *BandSidelink_r16_lowSE_64QAM_MCS_TableSidelink_r16                `optional`
	Enb_sync_Sidelink_r16                            *BandSidelink_r16_enb_sync_Sidelink_r16                            `optional`
	Sl_TransmissionMode2_r16                         *BandSidelink_r16_sl_TransmissionMode2_r16                         `optional,ext-1`
	CongestionControlSidelink_r16                    *BandSidelink_r16_congestionControlSidelink_r16                    `optional,ext-1`
	FewerSymbolSlotSidelink_r16                      *BandSidelink_r16_fewerSymbolSlotSidelink_r16                      `optional,ext-1`
	Sl_openLoopPC_RSRP_ReportSidelink_r16            *BandSidelink_r16_sl_openLoopPC_RSRP_ReportSidelink_r16            `optional,ext-1`
	Sl_Rx_256QAM_r16                                 *BandSidelink_r16_sl_Rx_256QAM_r16                                 `optional,ext-1`
	Ue_PowerClassSidelink_r16                        *BandSidelink_r16_ue_PowerClassSidelink_r16                        `optional,ext-2`
	Sl_TransmissionMode2_RandomResourceSelection_r17 *BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17 `optional,ext-3`
	Sync_Sidelink_v1710                              *BandSidelink_r16_sync_Sidelink_v1710                              `optional,ext-3`
	Enb_sync_Sidelink_v1710                          *BandSidelink_r16_enb_sync_Sidelink_v1710                          `optional,ext-3`
	Rx_IUC_Scheme1_PreferredMode2Sidelink_r17        *BandSidelink_r16_rx_IUC_Scheme1_PreferredMode2Sidelink_r17        `optional,ext-3`
	Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17     *BandSidelink_r16_rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17     `optional,ext-3`
	Rx_IUC_Scheme2_Mode2Sidelink_r17                 *BandSidelink_r16_rx_IUC_Scheme2_Mode2Sidelink_r17                 `optional,ext-3`
	Rx_IUC_Scheme1_SCI_r17                           *BandSidelink_r16_rx_IUC_Scheme1_SCI_r17                           `optional,ext-3`
	Rx_IUC_Scheme1_SCI_ExplicitReq_r17               *BandSidelink_r16_rx_IUC_Scheme1_SCI_ExplicitReq_r17               `optional,ext-3`
}

func (ie *BandSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Sl_TransmissionMode2_r16 != nil || ie.CongestionControlSidelink_r16 != nil || ie.FewerSymbolSlotSidelink_r16 != nil || ie.Sl_openLoopPC_RSRP_ReportSidelink_r16 != nil || ie.Sl_Rx_256QAM_r16 != nil || ie.Ue_PowerClassSidelink_r16 != nil || ie.Sl_TransmissionMode2_RandomResourceSelection_r17 != nil || ie.Sync_Sidelink_v1710 != nil || ie.Enb_sync_Sidelink_v1710 != nil || ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme1_SCI_r17 != nil || ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Sl_Reception_r16 != nil, ie.Sl_TransmissionMode1_r16 != nil, ie.Sync_Sidelink_r16 != nil, ie.Sl_Tx_256QAM_r16 != nil, ie.Psfch_FormatZeroSidelink_r16 != nil, ie.LowSE_64QAM_MCS_TableSidelink_r16 != nil, ie.Enb_sync_Sidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.FreqBandSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FreqBandSidelink_r16", err)
	}
	if ie.Sl_Reception_r16 != nil {
		if err = ie.Sl_Reception_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_Reception_r16", err)
		}
	}
	if ie.Sl_TransmissionMode1_r16 != nil {
		if err = ie.Sl_TransmissionMode1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TransmissionMode1_r16", err)
		}
	}
	if ie.Sync_Sidelink_r16 != nil {
		if err = ie.Sync_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sync_Sidelink_r16", err)
		}
	}
	if ie.Sl_Tx_256QAM_r16 != nil {
		if err = ie.Sl_Tx_256QAM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_Tx_256QAM_r16", err)
		}
	}
	if ie.Psfch_FormatZeroSidelink_r16 != nil {
		if err = ie.Psfch_FormatZeroSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Psfch_FormatZeroSidelink_r16", err)
		}
	}
	if ie.LowSE_64QAM_MCS_TableSidelink_r16 != nil {
		if err = ie.LowSE_64QAM_MCS_TableSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LowSE_64QAM_MCS_TableSidelink_r16", err)
		}
	}
	if ie.Enb_sync_Sidelink_r16 != nil {
		if err = ie.Enb_sync_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Enb_sync_Sidelink_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.Sl_TransmissionMode2_r16 != nil || ie.CongestionControlSidelink_r16 != nil || ie.FewerSymbolSlotSidelink_r16 != nil || ie.Sl_openLoopPC_RSRP_ReportSidelink_r16 != nil || ie.Sl_Rx_256QAM_r16 != nil, ie.Ue_PowerClassSidelink_r16 != nil, ie.Sl_TransmissionMode2_RandomResourceSelection_r17 != nil || ie.Sync_Sidelink_v1710 != nil || ie.Enb_sync_Sidelink_v1710 != nil || ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme1_SCI_r17 != nil || ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BandSidelink_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_TransmissionMode2_r16 != nil, ie.CongestionControlSidelink_r16 != nil, ie.FewerSymbolSlotSidelink_r16 != nil, ie.Sl_openLoopPC_RSRP_ReportSidelink_r16 != nil, ie.Sl_Rx_256QAM_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_TransmissionMode2_r16 optional
			if ie.Sl_TransmissionMode2_r16 != nil {
				if err = ie.Sl_TransmissionMode2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_TransmissionMode2_r16", err)
				}
			}
			// encode CongestionControlSidelink_r16 optional
			if ie.CongestionControlSidelink_r16 != nil {
				if err = ie.CongestionControlSidelink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CongestionControlSidelink_r16", err)
				}
			}
			// encode FewerSymbolSlotSidelink_r16 optional
			if ie.FewerSymbolSlotSidelink_r16 != nil {
				if err = ie.FewerSymbolSlotSidelink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FewerSymbolSlotSidelink_r16", err)
				}
			}
			// encode Sl_openLoopPC_RSRP_ReportSidelink_r16 optional
			if ie.Sl_openLoopPC_RSRP_ReportSidelink_r16 != nil {
				if err = ie.Sl_openLoopPC_RSRP_ReportSidelink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_openLoopPC_RSRP_ReportSidelink_r16", err)
				}
			}
			// encode Sl_Rx_256QAM_r16 optional
			if ie.Sl_Rx_256QAM_r16 != nil {
				if err = ie.Sl_Rx_256QAM_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_Rx_256QAM_r16", err)
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
			optionals_ext_2 := []bool{ie.Ue_PowerClassSidelink_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ue_PowerClassSidelink_r16 optional
			if ie.Ue_PowerClassSidelink_r16 != nil {
				if err = ie.Ue_PowerClassSidelink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ue_PowerClassSidelink_r16", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.Sl_TransmissionMode2_RandomResourceSelection_r17 != nil, ie.Sync_Sidelink_v1710 != nil, ie.Enb_sync_Sidelink_v1710 != nil, ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil, ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil, ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 != nil, ie.Rx_IUC_Scheme1_SCI_r17 != nil, ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_TransmissionMode2_RandomResourceSelection_r17 optional
			if ie.Sl_TransmissionMode2_RandomResourceSelection_r17 != nil {
				if err = ie.Sl_TransmissionMode2_RandomResourceSelection_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_TransmissionMode2_RandomResourceSelection_r17", err)
				}
			}
			// encode Sync_Sidelink_v1710 optional
			if ie.Sync_Sidelink_v1710 != nil {
				if err = ie.Sync_Sidelink_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sync_Sidelink_v1710", err)
				}
			}
			// encode Enb_sync_Sidelink_v1710 optional
			if ie.Enb_sync_Sidelink_v1710 != nil {
				if err = ie.Enb_sync_Sidelink_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Enb_sync_Sidelink_v1710", err)
				}
			}
			// encode Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 optional
			if ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil {
				if err = ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rx_IUC_Scheme1_PreferredMode2Sidelink_r17", err)
				}
			}
			// encode Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 optional
			if ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil {
				if err = ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17", err)
				}
			}
			// encode Rx_IUC_Scheme2_Mode2Sidelink_r17 optional
			if ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 != nil {
				if err = ie.Rx_IUC_Scheme2_Mode2Sidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rx_IUC_Scheme2_Mode2Sidelink_r17", err)
				}
			}
			// encode Rx_IUC_Scheme1_SCI_r17 optional
			if ie.Rx_IUC_Scheme1_SCI_r17 != nil {
				if err = ie.Rx_IUC_Scheme1_SCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rx_IUC_Scheme1_SCI_r17", err)
				}
			}
			// encode Rx_IUC_Scheme1_SCI_ExplicitReq_r17 optional
			if ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil {
				if err = ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rx_IUC_Scheme1_SCI_ExplicitReq_r17", err)
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

func (ie *BandSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Reception_r16Present bool
	if Sl_Reception_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TransmissionMode1_r16Present bool
	if Sl_TransmissionMode1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sync_Sidelink_r16Present bool
	if Sync_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Tx_256QAM_r16Present bool
	if Sl_Tx_256QAM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Psfch_FormatZeroSidelink_r16Present bool
	if Psfch_FormatZeroSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LowSE_64QAM_MCS_TableSidelink_r16Present bool
	if LowSE_64QAM_MCS_TableSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Enb_sync_Sidelink_r16Present bool
	if Enb_sync_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.FreqBandSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FreqBandSidelink_r16", err)
	}
	if Sl_Reception_r16Present {
		ie.Sl_Reception_r16 = new(BandSidelink_r16_sl_Reception_r16)
		if err = ie.Sl_Reception_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Reception_r16", err)
		}
	}
	if Sl_TransmissionMode1_r16Present {
		ie.Sl_TransmissionMode1_r16 = new(BandSidelink_r16_sl_TransmissionMode1_r16)
		if err = ie.Sl_TransmissionMode1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TransmissionMode1_r16", err)
		}
	}
	if Sync_Sidelink_r16Present {
		ie.Sync_Sidelink_r16 = new(BandSidelink_r16_sync_Sidelink_r16)
		if err = ie.Sync_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sync_Sidelink_r16", err)
		}
	}
	if Sl_Tx_256QAM_r16Present {
		ie.Sl_Tx_256QAM_r16 = new(BandSidelink_r16_sl_Tx_256QAM_r16)
		if err = ie.Sl_Tx_256QAM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Tx_256QAM_r16", err)
		}
	}
	if Psfch_FormatZeroSidelink_r16Present {
		ie.Psfch_FormatZeroSidelink_r16 = new(BandSidelink_r16_psfch_FormatZeroSidelink_r16)
		if err = ie.Psfch_FormatZeroSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Psfch_FormatZeroSidelink_r16", err)
		}
	}
	if LowSE_64QAM_MCS_TableSidelink_r16Present {
		ie.LowSE_64QAM_MCS_TableSidelink_r16 = new(BandSidelink_r16_lowSE_64QAM_MCS_TableSidelink_r16)
		if err = ie.LowSE_64QAM_MCS_TableSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LowSE_64QAM_MCS_TableSidelink_r16", err)
		}
	}
	if Enb_sync_Sidelink_r16Present {
		ie.Enb_sync_Sidelink_r16 = new(BandSidelink_r16_enb_sync_Sidelink_r16)
		if err = ie.Enb_sync_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Enb_sync_Sidelink_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			Sl_TransmissionMode2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CongestionControlSidelink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FewerSymbolSlotSidelink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_openLoopPC_RSRP_ReportSidelink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_Rx_256QAM_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_TransmissionMode2_r16 optional
			if Sl_TransmissionMode2_r16Present {
				ie.Sl_TransmissionMode2_r16 = new(BandSidelink_r16_sl_TransmissionMode2_r16)
				if err = ie.Sl_TransmissionMode2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_TransmissionMode2_r16", err)
				}
			}
			// decode CongestionControlSidelink_r16 optional
			if CongestionControlSidelink_r16Present {
				ie.CongestionControlSidelink_r16 = new(BandSidelink_r16_congestionControlSidelink_r16)
				if err = ie.CongestionControlSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CongestionControlSidelink_r16", err)
				}
			}
			// decode FewerSymbolSlotSidelink_r16 optional
			if FewerSymbolSlotSidelink_r16Present {
				ie.FewerSymbolSlotSidelink_r16 = new(BandSidelink_r16_fewerSymbolSlotSidelink_r16)
				if err = ie.FewerSymbolSlotSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode FewerSymbolSlotSidelink_r16", err)
				}
			}
			// decode Sl_openLoopPC_RSRP_ReportSidelink_r16 optional
			if Sl_openLoopPC_RSRP_ReportSidelink_r16Present {
				ie.Sl_openLoopPC_RSRP_ReportSidelink_r16 = new(BandSidelink_r16_sl_openLoopPC_RSRP_ReportSidelink_r16)
				if err = ie.Sl_openLoopPC_RSRP_ReportSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_openLoopPC_RSRP_ReportSidelink_r16", err)
				}
			}
			// decode Sl_Rx_256QAM_r16 optional
			if Sl_Rx_256QAM_r16Present {
				ie.Sl_Rx_256QAM_r16 = new(BandSidelink_r16_sl_Rx_256QAM_r16)
				if err = ie.Sl_Rx_256QAM_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_Rx_256QAM_r16", err)
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

			Ue_PowerClassSidelink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ue_PowerClassSidelink_r16 optional
			if Ue_PowerClassSidelink_r16Present {
				ie.Ue_PowerClassSidelink_r16 = new(BandSidelink_r16_ue_PowerClassSidelink_r16)
				if err = ie.Ue_PowerClassSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ue_PowerClassSidelink_r16", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Sl_TransmissionMode2_RandomResourceSelection_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sync_Sidelink_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Enb_sync_Sidelink_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Rx_IUC_Scheme1_PreferredMode2Sidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Rx_IUC_Scheme2_Mode2Sidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Rx_IUC_Scheme1_SCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Rx_IUC_Scheme1_SCI_ExplicitReq_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_TransmissionMode2_RandomResourceSelection_r17 optional
			if Sl_TransmissionMode2_RandomResourceSelection_r17Present {
				ie.Sl_TransmissionMode2_RandomResourceSelection_r17 = new(BandSidelink_r16_sl_TransmissionMode2_RandomResourceSelection_r17)
				if err = ie.Sl_TransmissionMode2_RandomResourceSelection_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_TransmissionMode2_RandomResourceSelection_r17", err)
				}
			}
			// decode Sync_Sidelink_v1710 optional
			if Sync_Sidelink_v1710Present {
				ie.Sync_Sidelink_v1710 = new(BandSidelink_r16_sync_Sidelink_v1710)
				if err = ie.Sync_Sidelink_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sync_Sidelink_v1710", err)
				}
			}
			// decode Enb_sync_Sidelink_v1710 optional
			if Enb_sync_Sidelink_v1710Present {
				ie.Enb_sync_Sidelink_v1710 = new(BandSidelink_r16_enb_sync_Sidelink_v1710)
				if err = ie.Enb_sync_Sidelink_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode Enb_sync_Sidelink_v1710", err)
				}
			}
			// decode Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 optional
			if Rx_IUC_Scheme1_PreferredMode2Sidelink_r17Present {
				ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 = new(BandSidelink_r16_rx_IUC_Scheme1_PreferredMode2Sidelink_r17)
				if err = ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rx_IUC_Scheme1_PreferredMode2Sidelink_r17", err)
				}
			}
			// decode Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 optional
			if Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17Present {
				ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 = new(BandSidelink_r16_rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17)
				if err = ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17", err)
				}
			}
			// decode Rx_IUC_Scheme2_Mode2Sidelink_r17 optional
			if Rx_IUC_Scheme2_Mode2Sidelink_r17Present {
				ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 = new(BandSidelink_r16_rx_IUC_Scheme2_Mode2Sidelink_r17)
				if err = ie.Rx_IUC_Scheme2_Mode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rx_IUC_Scheme2_Mode2Sidelink_r17", err)
				}
			}
			// decode Rx_IUC_Scheme1_SCI_r17 optional
			if Rx_IUC_Scheme1_SCI_r17Present {
				ie.Rx_IUC_Scheme1_SCI_r17 = new(BandSidelink_r16_rx_IUC_Scheme1_SCI_r17)
				if err = ie.Rx_IUC_Scheme1_SCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rx_IUC_Scheme1_SCI_r17", err)
				}
			}
			// decode Rx_IUC_Scheme1_SCI_ExplicitReq_r17 optional
			if Rx_IUC_Scheme1_SCI_ExplicitReq_r17Present {
				ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 = new(BandSidelink_r16_rx_IUC_Scheme1_SCI_ExplicitReq_r17)
				if err = ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rx_IUC_Scheme1_SCI_ExplicitReq_r17", err)
				}
			}
		}
	}
	return nil
}
