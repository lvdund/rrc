package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersXDD_Diff struct {
	SkipUplinkTxDynamic                *MAC_ParametersXDD_Diff_skipUplinkTxDynamic                `optional`
	LogicalChannelSR_DelayTimer        *MAC_ParametersXDD_Diff_logicalChannelSR_DelayTimer        `optional`
	LongDRX_Cycle                      *MAC_ParametersXDD_Diff_longDRX_Cycle                      `optional`
	ShortDRX_Cycle                     *MAC_ParametersXDD_Diff_shortDRX_Cycle                     `optional`
	MultipleSR_Configurations          *MAC_ParametersXDD_Diff_multipleSR_Configurations          `optional`
	MultipleConfiguredGrants           *MAC_ParametersXDD_Diff_multipleConfiguredGrants           `optional`
	SecondaryDRX_Group_r16             *MAC_ParametersXDD_Diff_secondaryDRX_Group_r16             `optional,ext-1`
	EnhancedSkipUplinkTxDynamic_r16    *MAC_ParametersXDD_Diff_enhancedSkipUplinkTxDynamic_r16    `optional,ext-2`
	EnhancedSkipUplinkTxConfigured_r16 *MAC_ParametersXDD_Diff_enhancedSkipUplinkTxConfigured_r16 `optional,ext-2`
}

func (ie *MAC_ParametersXDD_Diff) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.SecondaryDRX_Group_r16 != nil || ie.EnhancedSkipUplinkTxDynamic_r16 != nil || ie.EnhancedSkipUplinkTxConfigured_r16 != nil
	preambleBits := []bool{hasExtensions, ie.SkipUplinkTxDynamic != nil, ie.LogicalChannelSR_DelayTimer != nil, ie.LongDRX_Cycle != nil, ie.ShortDRX_Cycle != nil, ie.MultipleSR_Configurations != nil, ie.MultipleConfiguredGrants != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SkipUplinkTxDynamic != nil {
		if err = ie.SkipUplinkTxDynamic.Encode(w); err != nil {
			return utils.WrapError("Encode SkipUplinkTxDynamic", err)
		}
	}
	if ie.LogicalChannelSR_DelayTimer != nil {
		if err = ie.LogicalChannelSR_DelayTimer.Encode(w); err != nil {
			return utils.WrapError("Encode LogicalChannelSR_DelayTimer", err)
		}
	}
	if ie.LongDRX_Cycle != nil {
		if err = ie.LongDRX_Cycle.Encode(w); err != nil {
			return utils.WrapError("Encode LongDRX_Cycle", err)
		}
	}
	if ie.ShortDRX_Cycle != nil {
		if err = ie.ShortDRX_Cycle.Encode(w); err != nil {
			return utils.WrapError("Encode ShortDRX_Cycle", err)
		}
	}
	if ie.MultipleSR_Configurations != nil {
		if err = ie.MultipleSR_Configurations.Encode(w); err != nil {
			return utils.WrapError("Encode MultipleSR_Configurations", err)
		}
	}
	if ie.MultipleConfiguredGrants != nil {
		if err = ie.MultipleConfiguredGrants.Encode(w); err != nil {
			return utils.WrapError("Encode MultipleConfiguredGrants", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.SecondaryDRX_Group_r16 != nil, ie.EnhancedSkipUplinkTxDynamic_r16 != nil || ie.EnhancedSkipUplinkTxConfigured_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MAC_ParametersXDD_Diff", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.SecondaryDRX_Group_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SecondaryDRX_Group_r16 optional
			if ie.SecondaryDRX_Group_r16 != nil {
				if err = ie.SecondaryDRX_Group_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SecondaryDRX_Group_r16", err)
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
			optionals_ext_2 := []bool{ie.EnhancedSkipUplinkTxDynamic_r16 != nil, ie.EnhancedSkipUplinkTxConfigured_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode EnhancedSkipUplinkTxDynamic_r16 optional
			if ie.EnhancedSkipUplinkTxDynamic_r16 != nil {
				if err = ie.EnhancedSkipUplinkTxDynamic_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedSkipUplinkTxDynamic_r16", err)
				}
			}
			// encode EnhancedSkipUplinkTxConfigured_r16 optional
			if ie.EnhancedSkipUplinkTxConfigured_r16 != nil {
				if err = ie.EnhancedSkipUplinkTxConfigured_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedSkipUplinkTxConfigured_r16", err)
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

func (ie *MAC_ParametersXDD_Diff) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var SkipUplinkTxDynamicPresent bool
	if SkipUplinkTxDynamicPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LogicalChannelSR_DelayTimerPresent bool
	if LogicalChannelSR_DelayTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LongDRX_CyclePresent bool
	if LongDRX_CyclePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ShortDRX_CyclePresent bool
	if ShortDRX_CyclePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MultipleSR_ConfigurationsPresent bool
	if MultipleSR_ConfigurationsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MultipleConfiguredGrantsPresent bool
	if MultipleConfiguredGrantsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SkipUplinkTxDynamicPresent {
		ie.SkipUplinkTxDynamic = new(MAC_ParametersXDD_Diff_skipUplinkTxDynamic)
		if err = ie.SkipUplinkTxDynamic.Decode(r); err != nil {
			return utils.WrapError("Decode SkipUplinkTxDynamic", err)
		}
	}
	if LogicalChannelSR_DelayTimerPresent {
		ie.LogicalChannelSR_DelayTimer = new(MAC_ParametersXDD_Diff_logicalChannelSR_DelayTimer)
		if err = ie.LogicalChannelSR_DelayTimer.Decode(r); err != nil {
			return utils.WrapError("Decode LogicalChannelSR_DelayTimer", err)
		}
	}
	if LongDRX_CyclePresent {
		ie.LongDRX_Cycle = new(MAC_ParametersXDD_Diff_longDRX_Cycle)
		if err = ie.LongDRX_Cycle.Decode(r); err != nil {
			return utils.WrapError("Decode LongDRX_Cycle", err)
		}
	}
	if ShortDRX_CyclePresent {
		ie.ShortDRX_Cycle = new(MAC_ParametersXDD_Diff_shortDRX_Cycle)
		if err = ie.ShortDRX_Cycle.Decode(r); err != nil {
			return utils.WrapError("Decode ShortDRX_Cycle", err)
		}
	}
	if MultipleSR_ConfigurationsPresent {
		ie.MultipleSR_Configurations = new(MAC_ParametersXDD_Diff_multipleSR_Configurations)
		if err = ie.MultipleSR_Configurations.Decode(r); err != nil {
			return utils.WrapError("Decode MultipleSR_Configurations", err)
		}
	}
	if MultipleConfiguredGrantsPresent {
		ie.MultipleConfiguredGrants = new(MAC_ParametersXDD_Diff_multipleConfiguredGrants)
		if err = ie.MultipleConfiguredGrants.Decode(r); err != nil {
			return utils.WrapError("Decode MultipleConfiguredGrants", err)
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

			SecondaryDRX_Group_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SecondaryDRX_Group_r16 optional
			if SecondaryDRX_Group_r16Present {
				ie.SecondaryDRX_Group_r16 = new(MAC_ParametersXDD_Diff_secondaryDRX_Group_r16)
				if err = ie.SecondaryDRX_Group_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SecondaryDRX_Group_r16", err)
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

			EnhancedSkipUplinkTxDynamic_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnhancedSkipUplinkTxConfigured_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode EnhancedSkipUplinkTxDynamic_r16 optional
			if EnhancedSkipUplinkTxDynamic_r16Present {
				ie.EnhancedSkipUplinkTxDynamic_r16 = new(MAC_ParametersXDD_Diff_enhancedSkipUplinkTxDynamic_r16)
				if err = ie.EnhancedSkipUplinkTxDynamic_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedSkipUplinkTxDynamic_r16", err)
				}
			}
			// decode EnhancedSkipUplinkTxConfigured_r16 optional
			if EnhancedSkipUplinkTxConfigured_r16Present {
				ie.EnhancedSkipUplinkTxConfigured_r16 = new(MAC_ParametersXDD_Diff_enhancedSkipUplinkTxConfigured_r16)
				if err = ie.EnhancedSkipUplinkTxConfigured_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedSkipUplinkTxConfigured_r16", err)
				}
			}
		}
	}
	return nil
}
