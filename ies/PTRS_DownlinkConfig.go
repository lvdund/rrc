package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_DownlinkConfig struct {
	FrequencyDensity      []int64                                    `lb:2,ub:2,e_lb:0,e_ub:0,optional`
	TimeDensity           []int64                                    `lb:3,ub:3,e_lb:0,e_ub:0,optional`
	Epre_Ratio            *int64                                     `lb:0,ub:3,optional`
	ResourceElementOffset *PTRS_DownlinkConfig_resourceElementOffset `optional`
	MaxNrofPorts_r16      *PTRS_DownlinkConfig_maxNrofPorts_r16      `optional,ext-1`
}

func (ie *PTRS_DownlinkConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.MaxNrofPorts_r16 != nil
	preambleBits := []bool{hasExtensions, len(ie.FrequencyDensity) > 0, len(ie.TimeDensity) > 0, ie.Epre_Ratio != nil, ie.ResourceElementOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.FrequencyDensity) > 0 {
		tmp_FrequencyDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.FrequencyDensity {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
			tmp_FrequencyDensity.Value = append(tmp_FrequencyDensity.Value, &tmp_ie)
		}
		if err = tmp_FrequencyDensity.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyDensity", err)
		}
	}
	if len(ie.TimeDensity) > 0 {
		tmp_TimeDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		for _, i := range ie.TimeDensity {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
			tmp_TimeDensity.Value = append(tmp_TimeDensity.Value, &tmp_ie)
		}
		if err = tmp_TimeDensity.Encode(w); err != nil {
			return utils.WrapError("Encode TimeDensity", err)
		}
	}
	if ie.Epre_Ratio != nil {
		if err = w.WriteInteger(*ie.Epre_Ratio, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Encode Epre_Ratio", err)
		}
	}
	if ie.ResourceElementOffset != nil {
		if err = ie.ResourceElementOffset.Encode(w); err != nil {
			return utils.WrapError("Encode ResourceElementOffset", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.MaxNrofPorts_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PTRS_DownlinkConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.MaxNrofPorts_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxNrofPorts_r16 optional
			if ie.MaxNrofPorts_r16 != nil {
				if err = ie.MaxNrofPorts_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNrofPorts_r16", err)
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

func (ie *PTRS_DownlinkConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyDensityPresent bool
	if FrequencyDensityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeDensityPresent bool
	if TimeDensityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Epre_RatioPresent bool
	if Epre_RatioPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ResourceElementOffsetPresent bool
	if ResourceElementOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FrequencyDensityPresent {
		tmp_FrequencyDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_FrequencyDensity := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
			return &ie
		}
		if err = tmp_FrequencyDensity.Decode(r, fn_FrequencyDensity); err != nil {
			return utils.WrapError("Decode FrequencyDensity", err)
		}
		ie.FrequencyDensity = []int64{}
		for _, i := range tmp_FrequencyDensity.Value {
			ie.FrequencyDensity = append(ie.FrequencyDensity, int64(i.Value))
		}
	}
	if TimeDensityPresent {
		tmp_TimeDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		fn_TimeDensity := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
			return &ie
		}
		if err = tmp_TimeDensity.Decode(r, fn_TimeDensity); err != nil {
			return utils.WrapError("Decode TimeDensity", err)
		}
		ie.TimeDensity = []int64{}
		for _, i := range tmp_TimeDensity.Value {
			ie.TimeDensity = append(ie.TimeDensity, int64(i.Value))
		}
	}
	if Epre_RatioPresent {
		var tmp_int_Epre_Ratio int64
		if tmp_int_Epre_Ratio, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Epre_Ratio", err)
		}
		ie.Epre_Ratio = &tmp_int_Epre_Ratio
	}
	if ResourceElementOffsetPresent {
		ie.ResourceElementOffset = new(PTRS_DownlinkConfig_resourceElementOffset)
		if err = ie.ResourceElementOffset.Decode(r); err != nil {
			return utils.WrapError("Decode ResourceElementOffset", err)
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

			MaxNrofPorts_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxNrofPorts_r16 optional
			if MaxNrofPorts_r16Present {
				ie.MaxNrofPorts_r16 = new(PTRS_DownlinkConfig_maxNrofPorts_r16)
				if err = ie.MaxNrofPorts_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNrofPorts_r16", err)
				}
			}
		}
	}
	return nil
}
