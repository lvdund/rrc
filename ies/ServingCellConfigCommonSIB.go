package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServingCellConfigCommonSIB struct {
	DownlinkConfigCommon             DownlinkConfigCommonSIB                                      `madatory`
	UplinkConfigCommon               *UplinkConfigCommonSIB                                       `optional`
	SupplementaryUplink              *UplinkConfigCommonSIB                                       `optional`
	N_TimingAdvanceOffset            *ServingCellConfigCommonSIB_n_TimingAdvanceOffset            `optional`
	Ssb_PositionsInBurst             *ServingCellConfigCommonSIB_ssb_PositionsInBurst             `lb:8,ub:8,optional`
	Ssb_PeriodicityServingCell       ServingCellConfigCommonSIB_ssb_PeriodicityServingCell        `madatory`
	Tdd_UL_DL_ConfigurationCommon    *TDD_UL_DL_ConfigCommon                                      `optional`
	Ss_PBCH_BlockPower               int64                                                        `lb:-60,ub:50,madatory`
	ChannelAccessMode_r16            *ServingCellConfigCommonSIB_channelAccessMode_r16            `optional,ext-1`
	DiscoveryBurstWindowLength_r16   *ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16   `optional,ext-1`
	HighSpeedConfig_r16              *HighSpeedConfig_r16                                         `optional,ext-1`
	ChannelAccessMode2_r17           *ServingCellConfigCommonSIB_channelAccessMode2_r17           `optional,ext-2`
	DiscoveryBurstWindowLength_v1700 *ServingCellConfigCommonSIB_discoveryBurstWindowLength_v1700 `optional,ext-2`
	HighSpeedConfigFR2_r17           *HighSpeedConfigFR2_r17                                      `optional,ext-2`
	UplinkConfigCommon_v1700         *UplinkConfigCommonSIB_v1700                                 `optional,ext-2`
	EnhancedMeasurementLEO_r17       *ServingCellConfigCommonSIB_enhancedMeasurementLEO_r17       `optional,ext-3`
}

func (ie *ServingCellConfigCommonSIB) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ChannelAccessMode_r16 != nil || ie.DiscoveryBurstWindowLength_r16 != nil || ie.HighSpeedConfig_r16 != nil || ie.ChannelAccessMode2_r17 != nil || ie.DiscoveryBurstWindowLength_v1700 != nil || ie.HighSpeedConfigFR2_r17 != nil || ie.UplinkConfigCommon_v1700 != nil || ie.EnhancedMeasurementLEO_r17 != nil
	preambleBits := []bool{hasExtensions, ie.UplinkConfigCommon != nil, ie.SupplementaryUplink != nil, ie.N_TimingAdvanceOffset != nil, ie.Ssb_PositionsInBurst != nil, ie.Tdd_UL_DL_ConfigurationCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.DownlinkConfigCommon.Encode(w); err != nil {
		return utils.WrapError("Encode DownlinkConfigCommon", err)
	}
	if ie.UplinkConfigCommon != nil {
		if err = ie.UplinkConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkConfigCommon", err)
		}
	}
	if ie.SupplementaryUplink != nil {
		if err = ie.SupplementaryUplink.Encode(w); err != nil {
			return utils.WrapError("Encode SupplementaryUplink", err)
		}
	}
	if ie.N_TimingAdvanceOffset != nil {
		if err = ie.N_TimingAdvanceOffset.Encode(w); err != nil {
			return utils.WrapError("Encode N_TimingAdvanceOffset", err)
		}
	}
	if ie.Ssb_PositionsInBurst != nil {
		if err = ie.Ssb_PositionsInBurst.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_PositionsInBurst", err)
		}
	}
	if err = ie.Ssb_PeriodicityServingCell.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_PeriodicityServingCell", err)
	}
	if ie.Tdd_UL_DL_ConfigurationCommon != nil {
		if err = ie.Tdd_UL_DL_ConfigurationCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_UL_DL_ConfigurationCommon", err)
		}
	}
	if err = w.WriteInteger(ie.Ss_PBCH_BlockPower, &uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("WriteInteger Ss_PBCH_BlockPower", err)
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.ChannelAccessMode_r16 != nil || ie.DiscoveryBurstWindowLength_r16 != nil || ie.HighSpeedConfig_r16 != nil, ie.ChannelAccessMode2_r17 != nil || ie.DiscoveryBurstWindowLength_v1700 != nil || ie.HighSpeedConfigFR2_r17 != nil || ie.UplinkConfigCommon_v1700 != nil, ie.EnhancedMeasurementLEO_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ServingCellConfigCommonSIB", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ChannelAccessMode_r16 != nil, ie.DiscoveryBurstWindowLength_r16 != nil, ie.HighSpeedConfig_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ChannelAccessMode_r16 optional
			if ie.ChannelAccessMode_r16 != nil {
				if err = ie.ChannelAccessMode_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelAccessMode_r16", err)
				}
			}
			// encode DiscoveryBurstWindowLength_r16 optional
			if ie.DiscoveryBurstWindowLength_r16 != nil {
				if err = ie.DiscoveryBurstWindowLength_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DiscoveryBurstWindowLength_r16", err)
				}
			}
			// encode HighSpeedConfig_r16 optional
			if ie.HighSpeedConfig_r16 != nil {
				if err = ie.HighSpeedConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HighSpeedConfig_r16", err)
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
			optionals_ext_2 := []bool{ie.ChannelAccessMode2_r17 != nil, ie.DiscoveryBurstWindowLength_v1700 != nil, ie.HighSpeedConfigFR2_r17 != nil, ie.UplinkConfigCommon_v1700 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ChannelAccessMode2_r17 optional
			if ie.ChannelAccessMode2_r17 != nil {
				if err = ie.ChannelAccessMode2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelAccessMode2_r17", err)
				}
			}
			// encode DiscoveryBurstWindowLength_v1700 optional
			if ie.DiscoveryBurstWindowLength_v1700 != nil {
				if err = ie.DiscoveryBurstWindowLength_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DiscoveryBurstWindowLength_v1700", err)
				}
			}
			// encode HighSpeedConfigFR2_r17 optional
			if ie.HighSpeedConfigFR2_r17 != nil {
				if err = ie.HighSpeedConfigFR2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HighSpeedConfigFR2_r17", err)
				}
			}
			// encode UplinkConfigCommon_v1700 optional
			if ie.UplinkConfigCommon_v1700 != nil {
				if err = ie.UplinkConfigCommon_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkConfigCommon_v1700", err)
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
			optionals_ext_3 := []bool{ie.EnhancedMeasurementLEO_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode EnhancedMeasurementLEO_r17 optional
			if ie.EnhancedMeasurementLEO_r17 != nil {
				if err = ie.EnhancedMeasurementLEO_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedMeasurementLEO_r17", err)
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

func (ie *ServingCellConfigCommonSIB) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkConfigCommonPresent bool
	if UplinkConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SupplementaryUplinkPresent bool
	if SupplementaryUplinkPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var N_TimingAdvanceOffsetPresent bool
	if N_TimingAdvanceOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_PositionsInBurstPresent bool
	if Ssb_PositionsInBurstPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_UL_DL_ConfigurationCommonPresent bool
	if Tdd_UL_DL_ConfigurationCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.DownlinkConfigCommon.Decode(r); err != nil {
		return utils.WrapError("Decode DownlinkConfigCommon", err)
	}
	if UplinkConfigCommonPresent {
		ie.UplinkConfigCommon = new(UplinkConfigCommonSIB)
		if err = ie.UplinkConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkConfigCommon", err)
		}
	}
	if SupplementaryUplinkPresent {
		ie.SupplementaryUplink = new(UplinkConfigCommonSIB)
		if err = ie.SupplementaryUplink.Decode(r); err != nil {
			return utils.WrapError("Decode SupplementaryUplink", err)
		}
	}
	if N_TimingAdvanceOffsetPresent {
		ie.N_TimingAdvanceOffset = new(ServingCellConfigCommonSIB_n_TimingAdvanceOffset)
		if err = ie.N_TimingAdvanceOffset.Decode(r); err != nil {
			return utils.WrapError("Decode N_TimingAdvanceOffset", err)
		}
	}
	if Ssb_PositionsInBurstPresent {
		ie.Ssb_PositionsInBurst = new(ServingCellConfigCommonSIB_ssb_PositionsInBurst)
		if err = ie.Ssb_PositionsInBurst.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_PositionsInBurst", err)
		}
	}
	if err = ie.Ssb_PeriodicityServingCell.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_PeriodicityServingCell", err)
	}
	if Tdd_UL_DL_ConfigurationCommonPresent {
		ie.Tdd_UL_DL_ConfigurationCommon = new(TDD_UL_DL_ConfigCommon)
		if err = ie.Tdd_UL_DL_ConfigurationCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_UL_DL_ConfigurationCommon", err)
		}
	}
	var tmp_int_Ss_PBCH_BlockPower int64
	if tmp_int_Ss_PBCH_BlockPower, err = r.ReadInteger(&uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("ReadInteger Ss_PBCH_BlockPower", err)
	}
	ie.Ss_PBCH_BlockPower = tmp_int_Ss_PBCH_BlockPower

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

			ChannelAccessMode_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DiscoveryBurstWindowLength_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HighSpeedConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ChannelAccessMode_r16 optional
			if ChannelAccessMode_r16Present {
				ie.ChannelAccessMode_r16 = new(ServingCellConfigCommonSIB_channelAccessMode_r16)
				if err = ie.ChannelAccessMode_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelAccessMode_r16", err)
				}
			}
			// decode DiscoveryBurstWindowLength_r16 optional
			if DiscoveryBurstWindowLength_r16Present {
				ie.DiscoveryBurstWindowLength_r16 = new(ServingCellConfigCommonSIB_discoveryBurstWindowLength_r16)
				if err = ie.DiscoveryBurstWindowLength_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode DiscoveryBurstWindowLength_r16", err)
				}
			}
			// decode HighSpeedConfig_r16 optional
			if HighSpeedConfig_r16Present {
				ie.HighSpeedConfig_r16 = new(HighSpeedConfig_r16)
				if err = ie.HighSpeedConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode HighSpeedConfig_r16", err)
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

			ChannelAccessMode2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DiscoveryBurstWindowLength_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HighSpeedConfigFR2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkConfigCommon_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ChannelAccessMode2_r17 optional
			if ChannelAccessMode2_r17Present {
				ie.ChannelAccessMode2_r17 = new(ServingCellConfigCommonSIB_channelAccessMode2_r17)
				if err = ie.ChannelAccessMode2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelAccessMode2_r17", err)
				}
			}
			// decode DiscoveryBurstWindowLength_v1700 optional
			if DiscoveryBurstWindowLength_v1700Present {
				ie.DiscoveryBurstWindowLength_v1700 = new(ServingCellConfigCommonSIB_discoveryBurstWindowLength_v1700)
				if err = ie.DiscoveryBurstWindowLength_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode DiscoveryBurstWindowLength_v1700", err)
				}
			}
			// decode HighSpeedConfigFR2_r17 optional
			if HighSpeedConfigFR2_r17Present {
				ie.HighSpeedConfigFR2_r17 = new(HighSpeedConfigFR2_r17)
				if err = ie.HighSpeedConfigFR2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode HighSpeedConfigFR2_r17", err)
				}
			}
			// decode UplinkConfigCommon_v1700 optional
			if UplinkConfigCommon_v1700Present {
				ie.UplinkConfigCommon_v1700 = new(UplinkConfigCommonSIB_v1700)
				if err = ie.UplinkConfigCommon_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkConfigCommon_v1700", err)
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

			EnhancedMeasurementLEO_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode EnhancedMeasurementLEO_r17 optional
			if EnhancedMeasurementLEO_r17Present {
				ie.EnhancedMeasurementLEO_r17 = new(ServingCellConfigCommonSIB_enhancedMeasurementLEO_r17)
				if err = ie.EnhancedMeasurementLEO_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedMeasurementLEO_r17", err)
				}
			}
		}
	}
	return nil
}
