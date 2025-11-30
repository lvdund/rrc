package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelinkPC5_r16 struct {
	FreqBandSidelink_r16                         FreqBandIndicatorNR                                               `madatory`
	Sl_Reception_r16                             *BandSidelinkPC5_r16_sl_Reception_r16                             `optional`
	Sl_Tx_256QAM_r16                             *BandSidelinkPC5_r16_sl_Tx_256QAM_r16                             `optional`
	LowSE_64QAM_MCS_TableSidelink_r16            *BandSidelinkPC5_r16_lowSE_64QAM_MCS_TableSidelink_r16            `optional`
	Csi_ReportSidelink_r16                       *BandSidelinkPC5_r16_csi_ReportSidelink_r16                       `optional,ext-1`
	RankTwoReception_r16                         *BandSidelinkPC5_r16_rankTwoReception_r16                         `optional,ext-1`
	Sl_openLoopPC_RSRP_ReportSidelink_r16        *BandSidelinkPC5_r16_sl_openLoopPC_RSRP_ReportSidelink_r16        `optional,ext-1`
	Sl_Rx_256QAM_r16                             *BandSidelinkPC5_r16_sl_Rx_256QAM_r16                             `optional,ext-1`
	Rx_IUC_Scheme1_PreferredMode2Sidelink_r17    *BandSidelinkPC5_r16_rx_IUC_Scheme1_PreferredMode2Sidelink_r17    `optional,ext-2`
	Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 *BandSidelinkPC5_r16_rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 `optional,ext-2`
	Rx_IUC_Scheme2_Mode2Sidelink_r17             *BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17             `optional,ext-2`
	Rx_IUC_Scheme1_SCI_r17                       *BandSidelinkPC5_r16_rx_IUC_Scheme1_SCI_r17                       `optional,ext-2`
	Rx_IUC_Scheme1_SCI_ExplicitReq_r17           *BandSidelinkPC5_r16_rx_IUC_Scheme1_SCI_ExplicitReq_r17           `optional,ext-2`
	Scheme2_ConflictDeterminationRSRP_r17        *BandSidelinkPC5_r16_scheme2_ConflictDeterminationRSRP_r17        `optional,ext-2`
}

func (ie *BandSidelinkPC5_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Csi_ReportSidelink_r16 != nil || ie.RankTwoReception_r16 != nil || ie.Sl_openLoopPC_RSRP_ReportSidelink_r16 != nil || ie.Sl_Rx_256QAM_r16 != nil || ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme1_SCI_r17 != nil || ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil || ie.Scheme2_ConflictDeterminationRSRP_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Sl_Reception_r16 != nil, ie.Sl_Tx_256QAM_r16 != nil, ie.LowSE_64QAM_MCS_TableSidelink_r16 != nil}
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
	if ie.Sl_Tx_256QAM_r16 != nil {
		if err = ie.Sl_Tx_256QAM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_Tx_256QAM_r16", err)
		}
	}
	if ie.LowSE_64QAM_MCS_TableSidelink_r16 != nil {
		if err = ie.LowSE_64QAM_MCS_TableSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LowSE_64QAM_MCS_TableSidelink_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Csi_ReportSidelink_r16 != nil || ie.RankTwoReception_r16 != nil || ie.Sl_openLoopPC_RSRP_ReportSidelink_r16 != nil || ie.Sl_Rx_256QAM_r16 != nil, ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme1_SCI_r17 != nil || ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil || ie.Scheme2_ConflictDeterminationRSRP_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BandSidelinkPC5_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Csi_ReportSidelink_r16 != nil, ie.RankTwoReception_r16 != nil, ie.Sl_openLoopPC_RSRP_ReportSidelink_r16 != nil, ie.Sl_Rx_256QAM_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Csi_ReportSidelink_r16 optional
			if ie.Csi_ReportSidelink_r16 != nil {
				if err = ie.Csi_ReportSidelink_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_ReportSidelink_r16", err)
				}
			}
			// encode RankTwoReception_r16 optional
			if ie.RankTwoReception_r16 != nil {
				if err = ie.RankTwoReception_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RankTwoReception_r16", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil, ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil, ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 != nil, ie.Rx_IUC_Scheme1_SCI_r17 != nil, ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil, ie.Scheme2_ConflictDeterminationRSRP_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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
			// encode Scheme2_ConflictDeterminationRSRP_r17 optional
			if ie.Scheme2_ConflictDeterminationRSRP_r17 != nil {
				if err = ie.Scheme2_ConflictDeterminationRSRP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Scheme2_ConflictDeterminationRSRP_r17", err)
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

func (ie *BandSidelinkPC5_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Reception_r16Present bool
	if Sl_Reception_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Tx_256QAM_r16Present bool
	if Sl_Tx_256QAM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LowSE_64QAM_MCS_TableSidelink_r16Present bool
	if LowSE_64QAM_MCS_TableSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.FreqBandSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FreqBandSidelink_r16", err)
	}
	if Sl_Reception_r16Present {
		ie.Sl_Reception_r16 = new(BandSidelinkPC5_r16_sl_Reception_r16)
		if err = ie.Sl_Reception_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Reception_r16", err)
		}
	}
	if Sl_Tx_256QAM_r16Present {
		ie.Sl_Tx_256QAM_r16 = new(BandSidelinkPC5_r16_sl_Tx_256QAM_r16)
		if err = ie.Sl_Tx_256QAM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Tx_256QAM_r16", err)
		}
	}
	if LowSE_64QAM_MCS_TableSidelink_r16Present {
		ie.LowSE_64QAM_MCS_TableSidelink_r16 = new(BandSidelinkPC5_r16_lowSE_64QAM_MCS_TableSidelink_r16)
		if err = ie.LowSE_64QAM_MCS_TableSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LowSE_64QAM_MCS_TableSidelink_r16", err)
		}
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

			Csi_ReportSidelink_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RankTwoReception_r16Present, err := extReader.ReadBool()
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
			// decode Csi_ReportSidelink_r16 optional
			if Csi_ReportSidelink_r16Present {
				ie.Csi_ReportSidelink_r16 = new(BandSidelinkPC5_r16_csi_ReportSidelink_r16)
				if err = ie.Csi_ReportSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_ReportSidelink_r16", err)
				}
			}
			// decode RankTwoReception_r16 optional
			if RankTwoReception_r16Present {
				ie.RankTwoReception_r16 = new(BandSidelinkPC5_r16_rankTwoReception_r16)
				if err = ie.RankTwoReception_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode RankTwoReception_r16", err)
				}
			}
			// decode Sl_openLoopPC_RSRP_ReportSidelink_r16 optional
			if Sl_openLoopPC_RSRP_ReportSidelink_r16Present {
				ie.Sl_openLoopPC_RSRP_ReportSidelink_r16 = new(BandSidelinkPC5_r16_sl_openLoopPC_RSRP_ReportSidelink_r16)
				if err = ie.Sl_openLoopPC_RSRP_ReportSidelink_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_openLoopPC_RSRP_ReportSidelink_r16", err)
				}
			}
			// decode Sl_Rx_256QAM_r16 optional
			if Sl_Rx_256QAM_r16Present {
				ie.Sl_Rx_256QAM_r16 = new(BandSidelinkPC5_r16_sl_Rx_256QAM_r16)
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

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
			Scheme2_ConflictDeterminationRSRP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 optional
			if Rx_IUC_Scheme1_PreferredMode2Sidelink_r17Present {
				ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 = new(BandSidelinkPC5_r16_rx_IUC_Scheme1_PreferredMode2Sidelink_r17)
				if err = ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rx_IUC_Scheme1_PreferredMode2Sidelink_r17", err)
				}
			}
			// decode Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 optional
			if Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17Present {
				ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 = new(BandSidelinkPC5_r16_rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17)
				if err = ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17", err)
				}
			}
			// decode Rx_IUC_Scheme2_Mode2Sidelink_r17 optional
			if Rx_IUC_Scheme2_Mode2Sidelink_r17Present {
				ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 = new(BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17)
				if err = ie.Rx_IUC_Scheme2_Mode2Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rx_IUC_Scheme2_Mode2Sidelink_r17", err)
				}
			}
			// decode Rx_IUC_Scheme1_SCI_r17 optional
			if Rx_IUC_Scheme1_SCI_r17Present {
				ie.Rx_IUC_Scheme1_SCI_r17 = new(BandSidelinkPC5_r16_rx_IUC_Scheme1_SCI_r17)
				if err = ie.Rx_IUC_Scheme1_SCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rx_IUC_Scheme1_SCI_r17", err)
				}
			}
			// decode Rx_IUC_Scheme1_SCI_ExplicitReq_r17 optional
			if Rx_IUC_Scheme1_SCI_ExplicitReq_r17Present {
				ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 = new(BandSidelinkPC5_r16_rx_IUC_Scheme1_SCI_ExplicitReq_r17)
				if err = ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rx_IUC_Scheme1_SCI_ExplicitReq_r17", err)
				}
			}
			// decode Scheme2_ConflictDeterminationRSRP_r17 optional
			if Scheme2_ConflictDeterminationRSRP_r17Present {
				ie.Scheme2_ConflictDeterminationRSRP_r17 = new(BandSidelinkPC5_r16_scheme2_ConflictDeterminationRSRP_r17)
				if err = ie.Scheme2_ConflictDeterminationRSRP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Scheme2_ConflictDeterminationRSRP_r17", err)
				}
			}
		}
	}
	return nil
}
