package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RF_Parameters struct {
	SupportedBandListNR                                    []BandNR                                   `lb:1,ub:maxBands,madatory`
	SupportedBandCombinationList                           *BandCombinationList                       `optional`
	AppliedFreqBandListFilter                              *FreqBandList                              `optional`
	SupportedBandCombinationList_v1540                     *BandCombinationList_v1540                 `optional,ext-1`
	Srs_SwitchingTimeRequested                             *RF_Parameters_srs_SwitchingTimeRequested  `optional,ext-1`
	SupportedBandCombinationList_v1550                     *BandCombinationList_v1550                 `optional,ext-2`
	SupportedBandCombinationList_v1560                     *BandCombinationList_v1560                 `optional,ext-3`
	SupportedBandCombinationList_v1610                     *BandCombinationList_v1610                 `optional,ext-4`
	SupportedBandCombinationListSidelinkEUTRA_NR_r16       *BandCombinationListSidelinkEUTRA_NR_r16   `optional,ext-4`
	SupportedBandCombinationList_UplinkTxSwitch_r16        *BandCombinationList_UplinkTxSwitch_r16    `optional,ext-4`
	SupportedBandCombinationList_v1630                     *BandCombinationList_v1630                 `optional,ext-5`
	SupportedBandCombinationListSidelinkEUTRA_NR_v1630     *BandCombinationListSidelinkEUTRA_NR_v1630 `optional,ext-5`
	SupportedBandCombinationList_UplinkTxSwitch_v1630      *BandCombinationList_UplinkTxSwitch_v1630  `optional,ext-5`
	SupportedBandCombinationList_v1640                     *BandCombinationList_v1640                 `optional,ext-6`
	SupportedBandCombinationList_UplinkTxSwitch_v1640      *BandCombinationList_UplinkTxSwitch_v1640  `optional,ext-6`
	SupportedBandCombinationList_v1650                     *BandCombinationList_v1650                 `optional,ext-7`
	SupportedBandCombinationList_UplinkTxSwitch_v1650      *BandCombinationList_UplinkTxSwitch_v1650  `optional,ext-7`
	ExtendedBand_n77_r16                                   *RF_Parameters_extendedBand_n77_r16        `optional,ext-8`
	SupportedBandCombinationList_UplinkTxSwitch_v1670      *BandCombinationList_UplinkTxSwitch_v1670  `optional,ext-9`
	SupportedBandCombinationList_v1680                     *BandCombinationList_v1680                 `optional,ext-10`
	SupportedBandCombinationList_v1690                     *BandCombinationList_v1690                 `optional,ext-11`
	SupportedBandCombinationList_UplinkTxSwitch_v1690      *BandCombinationList_UplinkTxSwitch_v1690  `optional,ext-11`
	SupportedBandCombinationList_v1700                     *BandCombinationList_v1700                 `optional,ext-12`
	SupportedBandCombinationList_UplinkTxSwitch_v1700      *BandCombinationList_UplinkTxSwitch_v1700  `optional,ext-12`
	SupportedBandCombinationListSL_RelayDiscovery_r17      *[]byte                                    `optional,ext-12`
	SupportedBandCombinationListSL_NonRelayDiscovery_r17   *[]byte                                    `optional,ext-12`
	SupportedBandCombinationListSidelinkEUTRA_NR_v1710     *BandCombinationListSidelinkEUTRA_NR_v1710 `optional,ext-12`
	SidelinkRequested_r17                                  *RF_Parameters_sidelinkRequested_r17       `optional,ext-12`
	ExtendedBand_n77_2_r17                                 *RF_Parameters_extendedBand_n77_2_r17      `optional,ext-12`
	SupportedBandCombinationList_v1720                     *BandCombinationList_v1720                 `optional,ext-13`
	SupportedBandCombinationList_UplinkTxSwitch_v1720      *BandCombinationList_UplinkTxSwitch_v1720  `optional,ext-13`
	SupportedBandCombinationList_v1730                     *BandCombinationList_v1730                 `optional,ext-14`
	SupportedBandCombinationList_UplinkTxSwitch_v1730      *BandCombinationList_UplinkTxSwitch_v1730  `optional,ext-14`
	SupportedBandCombinationListSL_RelayDiscovery_v1730    *BandCombinationListSL_Discovery_r17       `optional,ext-14`
	SupportedBandCombinationListSL_NonRelayDiscovery_v1730 *BandCombinationListSL_Discovery_r17       `optional,ext-14`
}

func (ie *RF_Parameters) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.SupportedBandCombinationList_v1540 != nil || ie.Srs_SwitchingTimeRequested != nil || ie.SupportedBandCombinationList_v1550 != nil || ie.SupportedBandCombinationList_v1560 != nil || ie.SupportedBandCombinationList_v1610 != nil || ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_r16 != nil || ie.SupportedBandCombinationList_v1630 != nil || ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 != nil || ie.SupportedBandCombinationList_v1640 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 != nil || ie.SupportedBandCombinationList_v1650 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1650 != nil || ie.ExtendedBand_n77_r16 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1670 != nil || ie.SupportedBandCombinationList_v1680 != nil || ie.SupportedBandCombinationList_v1690 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1690 != nil || ie.SupportedBandCombinationList_v1700 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 != nil || ie.SupportedBandCombinationListSL_RelayDiscovery_r17 != nil || ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17 != nil || ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil || ie.SidelinkRequested_r17 != nil || ie.ExtendedBand_n77_2_r17 != nil || ie.SupportedBandCombinationList_v1720 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 != nil || ie.SupportedBandCombinationList_v1730 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 != nil || ie.SupportedBandCombinationListSL_RelayDiscovery_v1730 != nil || ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.SupportedBandCombinationList != nil, ie.AppliedFreqBandListFilter != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_SupportedBandListNR := utils.NewSequence[*BandNR]([]*BandNR{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
	for _, i := range ie.SupportedBandListNR {
		tmp_SupportedBandListNR.Value = append(tmp_SupportedBandListNR.Value, &i)
	}
	if err = tmp_SupportedBandListNR.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedBandListNR", err)
	}
	if ie.SupportedBandCombinationList != nil {
		if err = ie.SupportedBandCombinationList.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandCombinationList", err)
		}
	}
	if ie.AppliedFreqBandListFilter != nil {
		if err = ie.AppliedFreqBandListFilter.Encode(w); err != nil {
			return utils.WrapError("Encode AppliedFreqBandListFilter", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 14 bits for 14 extension groups
		extBitmap := []bool{ie.SupportedBandCombinationList_v1540 != nil || ie.Srs_SwitchingTimeRequested != nil, ie.SupportedBandCombinationList_v1550 != nil, ie.SupportedBandCombinationList_v1560 != nil, ie.SupportedBandCombinationList_v1610 != nil || ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_r16 != nil, ie.SupportedBandCombinationList_v1630 != nil || ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 != nil, ie.SupportedBandCombinationList_v1640 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 != nil, ie.SupportedBandCombinationList_v1650 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1650 != nil, ie.ExtendedBand_n77_r16 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1670 != nil, ie.SupportedBandCombinationList_v1680 != nil, ie.SupportedBandCombinationList_v1690 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1690 != nil, ie.SupportedBandCombinationList_v1700 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 != nil || ie.SupportedBandCombinationListSL_RelayDiscovery_r17 != nil || ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17 != nil || ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil || ie.SidelinkRequested_r17 != nil || ie.ExtendedBand_n77_2_r17 != nil, ie.SupportedBandCombinationList_v1720 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 != nil, ie.SupportedBandCombinationList_v1730 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 != nil || ie.SupportedBandCombinationListSL_RelayDiscovery_v1730 != nil || ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RF_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.SupportedBandCombinationList_v1540 != nil, ie.Srs_SwitchingTimeRequested != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1540 optional
			if ie.SupportedBandCombinationList_v1540 != nil {
				if err = ie.SupportedBandCombinationList_v1540.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1540", err)
				}
			}
			// encode Srs_SwitchingTimeRequested optional
			if ie.Srs_SwitchingTimeRequested != nil {
				if err = ie.Srs_SwitchingTimeRequested.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_SwitchingTimeRequested", err)
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
			optionals_ext_2 := []bool{ie.SupportedBandCombinationList_v1550 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1550 optional
			if ie.SupportedBandCombinationList_v1550 != nil {
				if err = ie.SupportedBandCombinationList_v1550.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1550", err)
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
			optionals_ext_3 := []bool{ie.SupportedBandCombinationList_v1560 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1560 optional
			if ie.SupportedBandCombinationList_v1560 != nil {
				if err = ie.SupportedBandCombinationList_v1560.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1560", err)
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
			optionals_ext_4 := []bool{ie.SupportedBandCombinationList_v1610 != nil, ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1610 optional
			if ie.SupportedBandCombinationList_v1610 != nil {
				if err = ie.SupportedBandCombinationList_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1610", err)
				}
			}
			// encode SupportedBandCombinationListSidelinkEUTRA_NR_r16 optional
			if ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16 != nil {
				if err = ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationListSidelinkEUTRA_NR_r16", err)
				}
			}
			// encode SupportedBandCombinationList_UplinkTxSwitch_r16 optional
			if ie.SupportedBandCombinationList_UplinkTxSwitch_r16 != nil {
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_UplinkTxSwitch_r16", err)
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
			optionals_ext_5 := []bool{ie.SupportedBandCombinationList_v1630 != nil, ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1630 optional
			if ie.SupportedBandCombinationList_v1630 != nil {
				if err = ie.SupportedBandCombinationList_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1630", err)
				}
			}
			// encode SupportedBandCombinationListSidelinkEUTRA_NR_v1630 optional
			if ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil {
				if err = ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationListSidelinkEUTRA_NR_v1630", err)
				}
			}
			// encode SupportedBandCombinationList_UplinkTxSwitch_v1630 optional
			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 != nil {
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_UplinkTxSwitch_v1630", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.SupportedBandCombinationList_v1640 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1640 optional
			if ie.SupportedBandCombinationList_v1640 != nil {
				if err = ie.SupportedBandCombinationList_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1640", err)
				}
			}
			// encode SupportedBandCombinationList_UplinkTxSwitch_v1640 optional
			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 != nil {
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_UplinkTxSwitch_v1640", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 7
		if extBitmap[6] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 7
			optionals_ext_7 := []bool{ie.SupportedBandCombinationList_v1650 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1650 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1650 optional
			if ie.SupportedBandCombinationList_v1650 != nil {
				if err = ie.SupportedBandCombinationList_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1650", err)
				}
			}
			// encode SupportedBandCombinationList_UplinkTxSwitch_v1650 optional
			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1650 != nil {
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_UplinkTxSwitch_v1650", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 8
		if extBitmap[7] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 8
			optionals_ext_8 := []bool{ie.ExtendedBand_n77_r16 != nil}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ExtendedBand_n77_r16 optional
			if ie.ExtendedBand_n77_r16 != nil {
				if err = ie.ExtendedBand_n77_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ExtendedBand_n77_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 9
		if extBitmap[8] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 9
			optionals_ext_9 := []bool{ie.SupportedBandCombinationList_UplinkTxSwitch_v1670 != nil}
			for _, bit := range optionals_ext_9 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_UplinkTxSwitch_v1670 optional
			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1670 != nil {
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1670.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_UplinkTxSwitch_v1670", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 10
		if extBitmap[9] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 10
			optionals_ext_10 := []bool{ie.SupportedBandCombinationList_v1680 != nil}
			for _, bit := range optionals_ext_10 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1680 optional
			if ie.SupportedBandCombinationList_v1680 != nil {
				if err = ie.SupportedBandCombinationList_v1680.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1680", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 11
		if extBitmap[10] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 11
			optionals_ext_11 := []bool{ie.SupportedBandCombinationList_v1690 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1690 != nil}
			for _, bit := range optionals_ext_11 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1690 optional
			if ie.SupportedBandCombinationList_v1690 != nil {
				if err = ie.SupportedBandCombinationList_v1690.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1690", err)
				}
			}
			// encode SupportedBandCombinationList_UplinkTxSwitch_v1690 optional
			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1690 != nil {
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1690.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_UplinkTxSwitch_v1690", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 12
		if extBitmap[11] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 12
			optionals_ext_12 := []bool{ie.SupportedBandCombinationList_v1700 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 != nil, ie.SupportedBandCombinationListSL_RelayDiscovery_r17 != nil, ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17 != nil, ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil, ie.SidelinkRequested_r17 != nil, ie.ExtendedBand_n77_2_r17 != nil}
			for _, bit := range optionals_ext_12 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1700 optional
			if ie.SupportedBandCombinationList_v1700 != nil {
				if err = ie.SupportedBandCombinationList_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1700", err)
				}
			}
			// encode SupportedBandCombinationList_UplinkTxSwitch_v1700 optional
			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 != nil {
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_UplinkTxSwitch_v1700", err)
				}
			}
			// encode SupportedBandCombinationListSL_RelayDiscovery_r17 optional
			if ie.SupportedBandCombinationListSL_RelayDiscovery_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.SupportedBandCombinationListSL_RelayDiscovery_r17, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationListSL_RelayDiscovery_r17", err)
				}
			}
			// encode SupportedBandCombinationListSL_NonRelayDiscovery_r17 optional
			if ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationListSL_NonRelayDiscovery_r17", err)
				}
			}
			// encode SupportedBandCombinationListSidelinkEUTRA_NR_v1710 optional
			if ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil {
				if err = ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationListSidelinkEUTRA_NR_v1710", err)
				}
			}
			// encode SidelinkRequested_r17 optional
			if ie.SidelinkRequested_r17 != nil {
				if err = ie.SidelinkRequested_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SidelinkRequested_r17", err)
				}
			}
			// encode ExtendedBand_n77_2_r17 optional
			if ie.ExtendedBand_n77_2_r17 != nil {
				if err = ie.ExtendedBand_n77_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ExtendedBand_n77_2_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 13
		if extBitmap[12] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 13
			optionals_ext_13 := []bool{ie.SupportedBandCombinationList_v1720 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 != nil}
			for _, bit := range optionals_ext_13 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1720 optional
			if ie.SupportedBandCombinationList_v1720 != nil {
				if err = ie.SupportedBandCombinationList_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1720", err)
				}
			}
			// encode SupportedBandCombinationList_UplinkTxSwitch_v1720 optional
			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 != nil {
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_UplinkTxSwitch_v1720", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 14
		if extBitmap[13] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 14
			optionals_ext_14 := []bool{ie.SupportedBandCombinationList_v1730 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 != nil, ie.SupportedBandCombinationListSL_RelayDiscovery_v1730 != nil, ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil}
			for _, bit := range optionals_ext_14 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedBandCombinationList_v1730 optional
			if ie.SupportedBandCombinationList_v1730 != nil {
				if err = ie.SupportedBandCombinationList_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_v1730", err)
				}
			}
			// encode SupportedBandCombinationList_UplinkTxSwitch_v1730 optional
			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 != nil {
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationList_UplinkTxSwitch_v1730", err)
				}
			}
			// encode SupportedBandCombinationListSL_RelayDiscovery_v1730 optional
			if ie.SupportedBandCombinationListSL_RelayDiscovery_v1730 != nil {
				if err = ie.SupportedBandCombinationListSL_RelayDiscovery_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationListSL_RelayDiscovery_v1730", err)
				}
			}
			// encode SupportedBandCombinationListSL_NonRelayDiscovery_v1730 optional
			if ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil {
				if err = ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedBandCombinationListSL_NonRelayDiscovery_v1730", err)
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

func (ie *RF_Parameters) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandCombinationListPresent bool
	if SupportedBandCombinationListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AppliedFreqBandListFilterPresent bool
	if AppliedFreqBandListFilterPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_SupportedBandListNR := utils.NewSequence[*BandNR]([]*BandNR{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
	fn_SupportedBandListNR := func() *BandNR {
		return new(BandNR)
	}
	if err = tmp_SupportedBandListNR.Decode(r, fn_SupportedBandListNR); err != nil {
		return utils.WrapError("Decode SupportedBandListNR", err)
	}
	ie.SupportedBandListNR = []BandNR{}
	for _, i := range tmp_SupportedBandListNR.Value {
		ie.SupportedBandListNR = append(ie.SupportedBandListNR, *i)
	}
	if SupportedBandCombinationListPresent {
		ie.SupportedBandCombinationList = new(BandCombinationList)
		if err = ie.SupportedBandCombinationList.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedBandCombinationList", err)
		}
	}
	if AppliedFreqBandListFilterPresent {
		ie.AppliedFreqBandListFilter = new(FreqBandList)
		if err = ie.AppliedFreqBandListFilter.Decode(r); err != nil {
			return utils.WrapError("Decode AppliedFreqBandListFilter", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 14 bits for 14 extension groups
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

			SupportedBandCombinationList_v1540Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_SwitchingTimeRequestedPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1540 optional
			if SupportedBandCombinationList_v1540Present {
				ie.SupportedBandCombinationList_v1540 = new(BandCombinationList_v1540)
				if err = ie.SupportedBandCombinationList_v1540.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1540", err)
				}
			}
			// decode Srs_SwitchingTimeRequested optional
			if Srs_SwitchingTimeRequestedPresent {
				ie.Srs_SwitchingTimeRequested = new(RF_Parameters_srs_SwitchingTimeRequested)
				if err = ie.Srs_SwitchingTimeRequested.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_SwitchingTimeRequested", err)
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

			SupportedBandCombinationList_v1550Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1550 optional
			if SupportedBandCombinationList_v1550Present {
				ie.SupportedBandCombinationList_v1550 = new(BandCombinationList_v1550)
				if err = ie.SupportedBandCombinationList_v1550.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1550", err)
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

			SupportedBandCombinationList_v1560Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1560 optional
			if SupportedBandCombinationList_v1560Present {
				ie.SupportedBandCombinationList_v1560 = new(BandCombinationList_v1560)
				if err = ie.SupportedBandCombinationList_v1560.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1560", err)
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

			SupportedBandCombinationList_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationListSidelinkEUTRA_NR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationList_UplinkTxSwitch_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1610 optional
			if SupportedBandCombinationList_v1610Present {
				ie.SupportedBandCombinationList_v1610 = new(BandCombinationList_v1610)
				if err = ie.SupportedBandCombinationList_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1610", err)
				}
			}
			// decode SupportedBandCombinationListSidelinkEUTRA_NR_r16 optional
			if SupportedBandCombinationListSidelinkEUTRA_NR_r16Present {
				ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16 = new(BandCombinationListSidelinkEUTRA_NR_r16)
				if err = ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationListSidelinkEUTRA_NR_r16", err)
				}
			}
			// decode SupportedBandCombinationList_UplinkTxSwitch_r16 optional
			if SupportedBandCombinationList_UplinkTxSwitch_r16Present {
				ie.SupportedBandCombinationList_UplinkTxSwitch_r16 = new(BandCombinationList_UplinkTxSwitch_r16)
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_UplinkTxSwitch_r16", err)
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

			SupportedBandCombinationList_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationListSidelinkEUTRA_NR_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationList_UplinkTxSwitch_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1630 optional
			if SupportedBandCombinationList_v1630Present {
				ie.SupportedBandCombinationList_v1630 = new(BandCombinationList_v1630)
				if err = ie.SupportedBandCombinationList_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1630", err)
				}
			}
			// decode SupportedBandCombinationListSidelinkEUTRA_NR_v1630 optional
			if SupportedBandCombinationListSidelinkEUTRA_NR_v1630Present {
				ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630 = new(BandCombinationListSidelinkEUTRA_NR_v1630)
				if err = ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationListSidelinkEUTRA_NR_v1630", err)
				}
			}
			// decode SupportedBandCombinationList_UplinkTxSwitch_v1630 optional
			if SupportedBandCombinationList_UplinkTxSwitch_v1630Present {
				ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 = new(BandCombinationList_UplinkTxSwitch_v1630)
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_UplinkTxSwitch_v1630", err)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SupportedBandCombinationList_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationList_UplinkTxSwitch_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1640 optional
			if SupportedBandCombinationList_v1640Present {
				ie.SupportedBandCombinationList_v1640 = new(BandCombinationList_v1640)
				if err = ie.SupportedBandCombinationList_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1640", err)
				}
			}
			// decode SupportedBandCombinationList_UplinkTxSwitch_v1640 optional
			if SupportedBandCombinationList_UplinkTxSwitch_v1640Present {
				ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 = new(BandCombinationList_UplinkTxSwitch_v1640)
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_UplinkTxSwitch_v1640", err)
				}
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SupportedBandCombinationList_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationList_UplinkTxSwitch_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1650 optional
			if SupportedBandCombinationList_v1650Present {
				ie.SupportedBandCombinationList_v1650 = new(BandCombinationList_v1650)
				if err = ie.SupportedBandCombinationList_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1650", err)
				}
			}
			// decode SupportedBandCombinationList_UplinkTxSwitch_v1650 optional
			if SupportedBandCombinationList_UplinkTxSwitch_v1650Present {
				ie.SupportedBandCombinationList_UplinkTxSwitch_v1650 = new(BandCombinationList_UplinkTxSwitch_v1650)
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_UplinkTxSwitch_v1650", err)
				}
			}
		}
		// decode extension group 8
		if len(extBitmap) > 7 && extBitmap[7] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			ExtendedBand_n77_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ExtendedBand_n77_r16 optional
			if ExtendedBand_n77_r16Present {
				ie.ExtendedBand_n77_r16 = new(RF_Parameters_extendedBand_n77_r16)
				if err = ie.ExtendedBand_n77_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ExtendedBand_n77_r16", err)
				}
			}
		}
		// decode extension group 9
		if len(extBitmap) > 8 && extBitmap[8] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SupportedBandCombinationList_UplinkTxSwitch_v1670Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_UplinkTxSwitch_v1670 optional
			if SupportedBandCombinationList_UplinkTxSwitch_v1670Present {
				ie.SupportedBandCombinationList_UplinkTxSwitch_v1670 = new(BandCombinationList_UplinkTxSwitch_v1670)
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1670.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_UplinkTxSwitch_v1670", err)
				}
			}
		}
		// decode extension group 10
		if len(extBitmap) > 9 && extBitmap[9] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SupportedBandCombinationList_v1680Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1680 optional
			if SupportedBandCombinationList_v1680Present {
				ie.SupportedBandCombinationList_v1680 = new(BandCombinationList_v1680)
				if err = ie.SupportedBandCombinationList_v1680.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1680", err)
				}
			}
		}
		// decode extension group 11
		if len(extBitmap) > 10 && extBitmap[10] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SupportedBandCombinationList_v1690Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationList_UplinkTxSwitch_v1690Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1690 optional
			if SupportedBandCombinationList_v1690Present {
				ie.SupportedBandCombinationList_v1690 = new(BandCombinationList_v1690)
				if err = ie.SupportedBandCombinationList_v1690.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1690", err)
				}
			}
			// decode SupportedBandCombinationList_UplinkTxSwitch_v1690 optional
			if SupportedBandCombinationList_UplinkTxSwitch_v1690Present {
				ie.SupportedBandCombinationList_UplinkTxSwitch_v1690 = new(BandCombinationList_UplinkTxSwitch_v1690)
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1690.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_UplinkTxSwitch_v1690", err)
				}
			}
		}
		// decode extension group 12
		if len(extBitmap) > 11 && extBitmap[11] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SupportedBandCombinationList_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationList_UplinkTxSwitch_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationListSL_RelayDiscovery_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationListSL_NonRelayDiscovery_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationListSidelinkEUTRA_NR_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SidelinkRequested_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ExtendedBand_n77_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1700 optional
			if SupportedBandCombinationList_v1700Present {
				ie.SupportedBandCombinationList_v1700 = new(BandCombinationList_v1700)
				if err = ie.SupportedBandCombinationList_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1700", err)
				}
			}
			// decode SupportedBandCombinationList_UplinkTxSwitch_v1700 optional
			if SupportedBandCombinationList_UplinkTxSwitch_v1700Present {
				ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 = new(BandCombinationList_UplinkTxSwitch_v1700)
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_UplinkTxSwitch_v1700", err)
				}
			}
			// decode SupportedBandCombinationListSL_RelayDiscovery_r17 optional
			if SupportedBandCombinationListSL_RelayDiscovery_r17Present {
				var tmp_os_SupportedBandCombinationListSL_RelayDiscovery_r17 []byte
				if tmp_os_SupportedBandCombinationListSL_RelayDiscovery_r17, err = extReader.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationListSL_RelayDiscovery_r17", err)
				}
				ie.SupportedBandCombinationListSL_RelayDiscovery_r17 = &tmp_os_SupportedBandCombinationListSL_RelayDiscovery_r17
			}
			// decode SupportedBandCombinationListSL_NonRelayDiscovery_r17 optional
			if SupportedBandCombinationListSL_NonRelayDiscovery_r17Present {
				var tmp_os_SupportedBandCombinationListSL_NonRelayDiscovery_r17 []byte
				if tmp_os_SupportedBandCombinationListSL_NonRelayDiscovery_r17, err = extReader.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationListSL_NonRelayDiscovery_r17", err)
				}
				ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17 = &tmp_os_SupportedBandCombinationListSL_NonRelayDiscovery_r17
			}
			// decode SupportedBandCombinationListSidelinkEUTRA_NR_v1710 optional
			if SupportedBandCombinationListSidelinkEUTRA_NR_v1710Present {
				ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710 = new(BandCombinationListSidelinkEUTRA_NR_v1710)
				if err = ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationListSidelinkEUTRA_NR_v1710", err)
				}
			}
			// decode SidelinkRequested_r17 optional
			if SidelinkRequested_r17Present {
				ie.SidelinkRequested_r17 = new(RF_Parameters_sidelinkRequested_r17)
				if err = ie.SidelinkRequested_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SidelinkRequested_r17", err)
				}
			}
			// decode ExtendedBand_n77_2_r17 optional
			if ExtendedBand_n77_2_r17Present {
				ie.ExtendedBand_n77_2_r17 = new(RF_Parameters_extendedBand_n77_2_r17)
				if err = ie.ExtendedBand_n77_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ExtendedBand_n77_2_r17", err)
				}
			}
		}
		// decode extension group 13
		if len(extBitmap) > 12 && extBitmap[12] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SupportedBandCombinationList_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationList_UplinkTxSwitch_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1720 optional
			if SupportedBandCombinationList_v1720Present {
				ie.SupportedBandCombinationList_v1720 = new(BandCombinationList_v1720)
				if err = ie.SupportedBandCombinationList_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1720", err)
				}
			}
			// decode SupportedBandCombinationList_UplinkTxSwitch_v1720 optional
			if SupportedBandCombinationList_UplinkTxSwitch_v1720Present {
				ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 = new(BandCombinationList_UplinkTxSwitch_v1720)
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_UplinkTxSwitch_v1720", err)
				}
			}
		}
		// decode extension group 14
		if len(extBitmap) > 13 && extBitmap[13] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SupportedBandCombinationList_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationList_UplinkTxSwitch_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationListSL_RelayDiscovery_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedBandCombinationListSL_NonRelayDiscovery_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedBandCombinationList_v1730 optional
			if SupportedBandCombinationList_v1730Present {
				ie.SupportedBandCombinationList_v1730 = new(BandCombinationList_v1730)
				if err = ie.SupportedBandCombinationList_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_v1730", err)
				}
			}
			// decode SupportedBandCombinationList_UplinkTxSwitch_v1730 optional
			if SupportedBandCombinationList_UplinkTxSwitch_v1730Present {
				ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 = new(BandCombinationList_UplinkTxSwitch_v1730)
				if err = ie.SupportedBandCombinationList_UplinkTxSwitch_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationList_UplinkTxSwitch_v1730", err)
				}
			}
			// decode SupportedBandCombinationListSL_RelayDiscovery_v1730 optional
			if SupportedBandCombinationListSL_RelayDiscovery_v1730Present {
				ie.SupportedBandCombinationListSL_RelayDiscovery_v1730 = new(BandCombinationListSL_Discovery_r17)
				if err = ie.SupportedBandCombinationListSL_RelayDiscovery_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationListSL_RelayDiscovery_v1730", err)
				}
			}
			// decode SupportedBandCombinationListSL_NonRelayDiscovery_v1730 optional
			if SupportedBandCombinationListSL_NonRelayDiscovery_v1730Present {
				ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730 = new(BandCombinationListSL_Discovery_r17)
				if err = ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedBandCombinationListSL_NonRelayDiscovery_v1730", err)
				}
			}
		}
	}
	return nil
}
