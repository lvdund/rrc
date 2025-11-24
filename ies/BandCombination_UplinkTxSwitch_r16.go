package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_r16 struct {
	BandCombination_r16                        BandCombination                                                                `madatory`
	BandCombination_v1540                      *BandCombination_v1540                                                         `optional`
	BandCombination_v1560                      *BandCombination_v1560                                                         `optional`
	BandCombination_v1570                      *BandCombination_v1570                                                         `optional`
	BandCombination_v1580                      *BandCombination_v1580                                                         `optional`
	BandCombination_v1590                      *BandCombination_v1590                                                         `optional`
	BandCombination_v1610                      *BandCombination_v1610                                                         `optional`
	SupportedBandPairListNR_r16                []ULTxSwitchingBandPair_r16                                                    `lb:1,ub:maxULTxSwitchingBandPairs,madatory`
	UplinkTxSwitching_OptionSupport_r16        *BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_OptionSupport_r16        `optional`
	UplinkTxSwitching_PowerBoosting_r16        *BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PowerBoosting_r16        `optional`
	UplinkTxSwitching_PUSCH_TransCoherence_r16 *BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16 `optional,ext-1`
}

func (ie *BandCombination_UplinkTxSwitch_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.UplinkTxSwitching_PUSCH_TransCoherence_r16 != nil
	preambleBits := []bool{hasExtensions, ie.BandCombination_v1540 != nil, ie.BandCombination_v1560 != nil, ie.BandCombination_v1570 != nil, ie.BandCombination_v1580 != nil, ie.BandCombination_v1590 != nil, ie.BandCombination_v1610 != nil, ie.UplinkTxSwitching_OptionSupport_r16 != nil, ie.UplinkTxSwitching_PowerBoosting_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.BandCombination_r16.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombination_r16", err)
	}
	if ie.BandCombination_v1540 != nil {
		if err = ie.BandCombination_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v1540", err)
		}
	}
	if ie.BandCombination_v1560 != nil {
		if err = ie.BandCombination_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v1560", err)
		}
	}
	if ie.BandCombination_v1570 != nil {
		if err = ie.BandCombination_v1570.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v1570", err)
		}
	}
	if ie.BandCombination_v1580 != nil {
		if err = ie.BandCombination_v1580.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v1580", err)
		}
	}
	if ie.BandCombination_v1590 != nil {
		if err = ie.BandCombination_v1590.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v1590", err)
		}
	}
	if ie.BandCombination_v1610 != nil {
		if err = ie.BandCombination_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v1610", err)
		}
	}
	tmp_SupportedBandPairListNR_r16 := utils.NewSequence[*ULTxSwitchingBandPair_r16]([]*ULTxSwitchingBandPair_r16{}, uper.Constraint{Lb: 1, Ub: maxULTxSwitchingBandPairs}, false)
	for _, i := range ie.SupportedBandPairListNR_r16 {
		tmp_SupportedBandPairListNR_r16.Value = append(tmp_SupportedBandPairListNR_r16.Value, &i)
	}
	if err = tmp_SupportedBandPairListNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedBandPairListNR_r16", err)
	}
	if ie.UplinkTxSwitching_OptionSupport_r16 != nil {
		if err = ie.UplinkTxSwitching_OptionSupport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkTxSwitching_OptionSupport_r16", err)
		}
	}
	if ie.UplinkTxSwitching_PowerBoosting_r16 != nil {
		if err = ie.UplinkTxSwitching_PowerBoosting_r16.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkTxSwitching_PowerBoosting_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.UplinkTxSwitching_PUSCH_TransCoherence_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BandCombination_UplinkTxSwitch_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.UplinkTxSwitching_PUSCH_TransCoherence_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode UplinkTxSwitching_PUSCH_TransCoherence_r16 optional
			if ie.UplinkTxSwitching_PUSCH_TransCoherence_r16 != nil {
				if err = ie.UplinkTxSwitching_PUSCH_TransCoherence_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkTxSwitching_PUSCH_TransCoherence_r16", err)
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

func (ie *BandCombination_UplinkTxSwitch_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var BandCombination_v1540Present bool
	if BandCombination_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BandCombination_v1560Present bool
	if BandCombination_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BandCombination_v1570Present bool
	if BandCombination_v1570Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BandCombination_v1580Present bool
	if BandCombination_v1580Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BandCombination_v1590Present bool
	if BandCombination_v1590Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BandCombination_v1610Present bool
	if BandCombination_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkTxSwitching_OptionSupport_r16Present bool
	if UplinkTxSwitching_OptionSupport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkTxSwitching_PowerBoosting_r16Present bool
	if UplinkTxSwitching_PowerBoosting_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.BandCombination_r16.Decode(r); err != nil {
		return utils.WrapError("Decode BandCombination_r16", err)
	}
	if BandCombination_v1540Present {
		ie.BandCombination_v1540 = new(BandCombination_v1540)
		if err = ie.BandCombination_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v1540", err)
		}
	}
	if BandCombination_v1560Present {
		ie.BandCombination_v1560 = new(BandCombination_v1560)
		if err = ie.BandCombination_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v1560", err)
		}
	}
	if BandCombination_v1570Present {
		ie.BandCombination_v1570 = new(BandCombination_v1570)
		if err = ie.BandCombination_v1570.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v1570", err)
		}
	}
	if BandCombination_v1580Present {
		ie.BandCombination_v1580 = new(BandCombination_v1580)
		if err = ie.BandCombination_v1580.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v1580", err)
		}
	}
	if BandCombination_v1590Present {
		ie.BandCombination_v1590 = new(BandCombination_v1590)
		if err = ie.BandCombination_v1590.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v1590", err)
		}
	}
	if BandCombination_v1610Present {
		ie.BandCombination_v1610 = new(BandCombination_v1610)
		if err = ie.BandCombination_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v1610", err)
		}
	}
	tmp_SupportedBandPairListNR_r16 := utils.NewSequence[*ULTxSwitchingBandPair_r16]([]*ULTxSwitchingBandPair_r16{}, uper.Constraint{Lb: 1, Ub: maxULTxSwitchingBandPairs}, false)
	fn_SupportedBandPairListNR_r16 := func() *ULTxSwitchingBandPair_r16 {
		return new(ULTxSwitchingBandPair_r16)
	}
	if err = tmp_SupportedBandPairListNR_r16.Decode(r, fn_SupportedBandPairListNR_r16); err != nil {
		return utils.WrapError("Decode SupportedBandPairListNR_r16", err)
	}
	ie.SupportedBandPairListNR_r16 = []ULTxSwitchingBandPair_r16{}
	for _, i := range tmp_SupportedBandPairListNR_r16.Value {
		ie.SupportedBandPairListNR_r16 = append(ie.SupportedBandPairListNR_r16, *i)
	}
	if UplinkTxSwitching_OptionSupport_r16Present {
		ie.UplinkTxSwitching_OptionSupport_r16 = new(BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_OptionSupport_r16)
		if err = ie.UplinkTxSwitching_OptionSupport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkTxSwitching_OptionSupport_r16", err)
		}
	}
	if UplinkTxSwitching_PowerBoosting_r16Present {
		ie.UplinkTxSwitching_PowerBoosting_r16 = new(BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PowerBoosting_r16)
		if err = ie.UplinkTxSwitching_PowerBoosting_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkTxSwitching_PowerBoosting_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
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

			UplinkTxSwitching_PUSCH_TransCoherence_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode UplinkTxSwitching_PUSCH_TransCoherence_r16 optional
			if UplinkTxSwitching_PUSCH_TransCoherence_r16Present {
				ie.UplinkTxSwitching_PUSCH_TransCoherence_r16 = new(BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16)
				if err = ie.UplinkTxSwitching_PUSCH_TransCoherence_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkTxSwitching_PUSCH_TransCoherence_r16", err)
				}
			}
		}
	}
	return nil
}
