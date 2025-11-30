package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasConfig struct {
	MeasObjectToRemoveList         *MeasObjectToRemoveList                    `optional`
	MeasObjectToAddModList         *MeasObjectToAddModList                    `optional`
	ReportConfigToRemoveList       *ReportConfigToRemoveList                  `optional`
	ReportConfigToAddModList       *ReportConfigToAddModList                  `optional`
	MeasIdToRemoveList             *MeasIdToRemoveList                        `optional`
	MeasIdToAddModList             *MeasIdToAddModList                        `optional`
	S_MeasureConfig                *MeasConfig_s_MeasureConfig                `optional`
	QuantityConfig                 *QuantityConfig                            `optional`
	MeasGapConfig                  *MeasGapConfig                             `optional`
	MeasGapSharingConfig           *MeasGapSharingConfig                      `optional`
	InterFrequencyConfig_NoGap_r16 *MeasConfig_interFrequencyConfig_NoGap_r16 `optional,ext-1`
}

func (ie *MeasConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.InterFrequencyConfig_NoGap_r16 != nil
	preambleBits := []bool{hasExtensions, ie.MeasObjectToRemoveList != nil, ie.MeasObjectToAddModList != nil, ie.ReportConfigToRemoveList != nil, ie.ReportConfigToAddModList != nil, ie.MeasIdToRemoveList != nil, ie.MeasIdToAddModList != nil, ie.S_MeasureConfig != nil, ie.QuantityConfig != nil, ie.MeasGapConfig != nil, ie.MeasGapSharingConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasObjectToRemoveList != nil {
		if err = ie.MeasObjectToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode MeasObjectToRemoveList", err)
		}
	}
	if ie.MeasObjectToAddModList != nil {
		if err = ie.MeasObjectToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode MeasObjectToAddModList", err)
		}
	}
	if ie.ReportConfigToRemoveList != nil {
		if err = ie.ReportConfigToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode ReportConfigToRemoveList", err)
		}
	}
	if ie.ReportConfigToAddModList != nil {
		if err = ie.ReportConfigToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode ReportConfigToAddModList", err)
		}
	}
	if ie.MeasIdToRemoveList != nil {
		if err = ie.MeasIdToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode MeasIdToRemoveList", err)
		}
	}
	if ie.MeasIdToAddModList != nil {
		if err = ie.MeasIdToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode MeasIdToAddModList", err)
		}
	}
	if ie.S_MeasureConfig != nil {
		if err = ie.S_MeasureConfig.Encode(w); err != nil {
			return utils.WrapError("Encode S_MeasureConfig", err)
		}
	}
	if ie.QuantityConfig != nil {
		if err = ie.QuantityConfig.Encode(w); err != nil {
			return utils.WrapError("Encode QuantityConfig", err)
		}
	}
	if ie.MeasGapConfig != nil {
		if err = ie.MeasGapConfig.Encode(w); err != nil {
			return utils.WrapError("Encode MeasGapConfig", err)
		}
	}
	if ie.MeasGapSharingConfig != nil {
		if err = ie.MeasGapSharingConfig.Encode(w); err != nil {
			return utils.WrapError("Encode MeasGapSharingConfig", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.InterFrequencyConfig_NoGap_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.InterFrequencyConfig_NoGap_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode InterFrequencyConfig_NoGap_r16 optional
			if ie.InterFrequencyConfig_NoGap_r16 != nil {
				if err = ie.InterFrequencyConfig_NoGap_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InterFrequencyConfig_NoGap_r16", err)
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

func (ie *MeasConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasObjectToRemoveListPresent bool
	if MeasObjectToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasObjectToAddModListPresent bool
	if MeasObjectToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReportConfigToRemoveListPresent bool
	if ReportConfigToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReportConfigToAddModListPresent bool
	if ReportConfigToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasIdToRemoveListPresent bool
	if MeasIdToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasIdToAddModListPresent bool
	if MeasIdToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var S_MeasureConfigPresent bool
	if S_MeasureConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var QuantityConfigPresent bool
	if QuantityConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasGapConfigPresent bool
	if MeasGapConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasGapSharingConfigPresent bool
	if MeasGapSharingConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasObjectToRemoveListPresent {
		ie.MeasObjectToRemoveList = new(MeasObjectToRemoveList)
		if err = ie.MeasObjectToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode MeasObjectToRemoveList", err)
		}
	}
	if MeasObjectToAddModListPresent {
		ie.MeasObjectToAddModList = new(MeasObjectToAddModList)
		if err = ie.MeasObjectToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode MeasObjectToAddModList", err)
		}
	}
	if ReportConfigToRemoveListPresent {
		ie.ReportConfigToRemoveList = new(ReportConfigToRemoveList)
		if err = ie.ReportConfigToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode ReportConfigToRemoveList", err)
		}
	}
	if ReportConfigToAddModListPresent {
		ie.ReportConfigToAddModList = new(ReportConfigToAddModList)
		if err = ie.ReportConfigToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode ReportConfigToAddModList", err)
		}
	}
	if MeasIdToRemoveListPresent {
		ie.MeasIdToRemoveList = new(MeasIdToRemoveList)
		if err = ie.MeasIdToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode MeasIdToRemoveList", err)
		}
	}
	if MeasIdToAddModListPresent {
		ie.MeasIdToAddModList = new(MeasIdToAddModList)
		if err = ie.MeasIdToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode MeasIdToAddModList", err)
		}
	}
	if S_MeasureConfigPresent {
		ie.S_MeasureConfig = new(MeasConfig_s_MeasureConfig)
		if err = ie.S_MeasureConfig.Decode(r); err != nil {
			return utils.WrapError("Decode S_MeasureConfig", err)
		}
	}
	if QuantityConfigPresent {
		ie.QuantityConfig = new(QuantityConfig)
		if err = ie.QuantityConfig.Decode(r); err != nil {
			return utils.WrapError("Decode QuantityConfig", err)
		}
	}
	if MeasGapConfigPresent {
		ie.MeasGapConfig = new(MeasGapConfig)
		if err = ie.MeasGapConfig.Decode(r); err != nil {
			return utils.WrapError("Decode MeasGapConfig", err)
		}
	}
	if MeasGapSharingConfigPresent {
		ie.MeasGapSharingConfig = new(MeasGapSharingConfig)
		if err = ie.MeasGapSharingConfig.Decode(r); err != nil {
			return utils.WrapError("Decode MeasGapSharingConfig", err)
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			InterFrequencyConfig_NoGap_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode InterFrequencyConfig_NoGap_r16 optional
			if InterFrequencyConfig_NoGap_r16Present {
				ie.InterFrequencyConfig_NoGap_r16 = new(MeasConfig_interFrequencyConfig_NoGap_r16)
				if err = ie.InterFrequencyConfig_NoGap_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode InterFrequencyConfig_NoGap_r16", err)
				}
			}
		}
	}
	return nil
}
