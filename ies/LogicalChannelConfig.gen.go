// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LogicalChannelConfig (line 8593).

var logicalChannelConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-SpecificParameters", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var logicalChannelConfigChannelAccessPriorityR16Constraints = per.Constrained(1, 4)

const (
	LogicalChannelConfig_Ext_BitRateMultiplier_r16_X40  = 0
	LogicalChannelConfig_Ext_BitRateMultiplier_r16_X70  = 1
	LogicalChannelConfig_Ext_BitRateMultiplier_r16_X100 = 2
	LogicalChannelConfig_Ext_BitRateMultiplier_r16_X200 = 3
)

var logicalChannelConfigExtBitRateMultiplierR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LogicalChannelConfig_Ext_BitRateMultiplier_r16_X40, LogicalChannelConfig_Ext_BitRateMultiplier_r16_X70, LogicalChannelConfig_Ext_BitRateMultiplier_r16_X100, LogicalChannelConfig_Ext_BitRateMultiplier_r16_X200},
}

var logicalChannelConfigUlSpecificParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "priority"},
		{Name: "prioritisedBitRate"},
		{Name: "bucketSizeDuration"},
		{Name: "allowedServingCells", Optional: true},
		{Name: "allowedSCS-List", Optional: true},
		{Name: "maxPUSCH-Duration", Optional: true},
		{Name: "configuredGrantType1Allowed", Optional: true},
		{Name: "logicalChannelGroup", Optional: true},
		{Name: "schedulingRequestID", Optional: true},
		{Name: "logicalChannelSR-Mask"},
		{Name: "logicalChannelSR-DelayTimerApplied"},
	},
}

const (
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps0     = 0
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps8     = 1
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps16    = 2
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps32    = 3
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps64    = 4
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps128   = 5
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps256   = 6
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps512   = 7
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps1024  = 8
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps2048  = 9
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps4096  = 10
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps8192  = 11
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps16384 = 12
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps32768 = 13
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps65536 = 14
	LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_Infinity  = 15
)

var logicalChannelConfigUlSpecificParametersPrioritisedBitRateConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps0, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps8, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps16, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps32, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps64, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps128, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps256, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps512, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps1024, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps2048, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps4096, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps8192, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps16384, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps32768, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_KBps65536, LogicalChannelConfig_Ul_SpecificParameters_PrioritisedBitRate_Infinity},
}

const (
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms5    = 0
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms10   = 1
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms20   = 2
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms50   = 3
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms100  = 4
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms150  = 5
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms300  = 6
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms500  = 7
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms1000 = 8
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare7 = 9
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare6 = 10
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare5 = 11
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare4 = 12
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare3 = 13
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare2 = 14
	LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare1 = 15
)

var logicalChannelConfigUlSpecificParametersBucketSizeDurationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms5, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms10, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms20, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms50, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms100, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms150, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms300, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms500, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Ms1000, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare7, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare6, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare5, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare4, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare3, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare2, LogicalChannelConfig_Ul_SpecificParameters_BucketSizeDuration_Spare1},
}

var logicalChannelConfigUlSpecificParametersAllowedServingCellsConstraints = per.SizeRange(1, common.MaxNrofServingCells_1)

var logicalChannelConfigUlSpecificParametersAllowedSCSListConstraints = per.SizeRange(1, common.MaxSCSs)

const (
	LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p02       = 0
	LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p04       = 1
	LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p0625     = 2
	LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p125      = 3
	LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p25       = 4
	LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p5        = 5
	LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p01_v1700 = 6
	LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Spare1       = 7
)

var logicalChannelConfigUlSpecificParametersMaxPUSCHDurationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p02, LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p04, LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p0625, LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p125, LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p25, LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p5, LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Ms0p01_v1700, LogicalChannelConfig_Ul_SpecificParameters_MaxPUSCH_Duration_Spare1},
}

const (
	LogicalChannelConfig_Ul_SpecificParameters_ConfiguredGrantType1Allowed_True = 0
)

var logicalChannelConfigUlSpecificParametersConfiguredGrantType1AllowedConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LogicalChannelConfig_Ul_SpecificParameters_ConfiguredGrantType1Allowed_True},
}

type LogicalChannelConfig struct {
	Ul_SpecificParameters *struct {
		Priority                           int64
		PrioritisedBitRate                 int64
		BucketSizeDuration                 int64
		AllowedServingCells                []ServCellIndex
		AllowedSCS_List                    []SubcarrierSpacing
		MaxPUSCH_Duration                  *int64
		ConfiguredGrantType1Allowed        *int64
		LogicalChannelGroup                *int64
		SchedulingRequestID                *SchedulingRequestId
		LogicalChannelSR_Mask              bool
		LogicalChannelSR_DelayTimerApplied bool
	}
	ChannelAccessPriority_r16 *int64
	BitRateMultiplier_r16     *int64
}

func (ie *LogicalChannelConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(logicalChannelConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ChannelAccessPriority_r16 != nil || ie.BitRateMultiplier_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_SpecificParameters != nil}); err != nil {
		return err
	}

	// 3. ul-SpecificParameters: sequence
	{
		if ie.Ul_SpecificParameters != nil {
			c := ie.Ul_SpecificParameters
			logicalChannelConfigUlSpecificParametersSeq := e.NewSequenceEncoder(logicalChannelConfigUlSpecificParametersConstraints)
			if err := logicalChannelConfigUlSpecificParametersSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := logicalChannelConfigUlSpecificParametersSeq.EncodePreamble([]bool{c.AllowedServingCells != nil, c.AllowedSCS_List != nil, c.MaxPUSCH_Duration != nil, c.ConfiguredGrantType1Allowed != nil, c.LogicalChannelGroup != nil, c.SchedulingRequestID != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Priority, per.Constrained(1, 16)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.PrioritisedBitRate, logicalChannelConfigUlSpecificParametersPrioritisedBitRateConstraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.BucketSizeDuration, logicalChannelConfigUlSpecificParametersBucketSizeDurationConstraints); err != nil {
				return err
			}
			if c.AllowedServingCells != nil {
				seqOf := e.NewSequenceOfEncoder(logicalChannelConfigUlSpecificParametersAllowedServingCellsConstraints)
				if err := seqOf.EncodeLength(int64(len(c.AllowedServingCells))); err != nil {
					return err
				}
				for i := range c.AllowedServingCells {
					if err := c.AllowedServingCells[i].Encode(e); err != nil {
						return err
					}
				}
			}
			if c.AllowedSCS_List != nil {
				seqOf := e.NewSequenceOfEncoder(logicalChannelConfigUlSpecificParametersAllowedSCSListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.AllowedSCS_List))); err != nil {
					return err
				}
				for i := range c.AllowedSCS_List {
					if err := c.AllowedSCS_List[i].Encode(e); err != nil {
						return err
					}
				}
			}
			if c.MaxPUSCH_Duration != nil {
				if err := e.EncodeEnumerated((*c.MaxPUSCH_Duration), logicalChannelConfigUlSpecificParametersMaxPUSCHDurationConstraints); err != nil {
					return err
				}
			}
			if c.ConfiguredGrantType1Allowed != nil {
				if err := e.EncodeEnumerated((*c.ConfiguredGrantType1Allowed), logicalChannelConfigUlSpecificParametersConfiguredGrantType1AllowedConstraints); err != nil {
					return err
				}
			}
			if c.LogicalChannelGroup != nil {
				if err := e.EncodeInteger((*c.LogicalChannelGroup), per.Constrained(0, common.MaxLCG_ID)); err != nil {
					return err
				}
			}
			if c.SchedulingRequestID != nil {
				if err := c.SchedulingRequestID.Encode(e); err != nil {
					return err
				}
			}
			if err := e.EncodeBoolean(c.LogicalChannelSR_Mask); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.LogicalChannelSR_DelayTimerApplied); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "channelAccessPriority-r16", Optional: true},
					{Name: "bitRateMultiplier-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ChannelAccessPriority_r16 != nil, ie.BitRateMultiplier_r16 != nil}); err != nil {
				return err
			}

			if ie.ChannelAccessPriority_r16 != nil {
				if err := ex.EncodeInteger(*ie.ChannelAccessPriority_r16, logicalChannelConfigChannelAccessPriorityR16Constraints); err != nil {
					return err
				}
			}

			if ie.BitRateMultiplier_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.BitRateMultiplier_r16, logicalChannelConfigExtBitRateMultiplierR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *LogicalChannelConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(logicalChannelConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ul-SpecificParameters: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Ul_SpecificParameters = &struct {
				Priority                           int64
				PrioritisedBitRate                 int64
				BucketSizeDuration                 int64
				AllowedServingCells                []ServCellIndex
				AllowedSCS_List                    []SubcarrierSpacing
				MaxPUSCH_Duration                  *int64
				ConfiguredGrantType1Allowed        *int64
				LogicalChannelGroup                *int64
				SchedulingRequestID                *SchedulingRequestId
				LogicalChannelSR_Mask              bool
				LogicalChannelSR_DelayTimerApplied bool
			}{}
			c := ie.Ul_SpecificParameters
			logicalChannelConfigUlSpecificParametersSeq := d.NewSequenceDecoder(logicalChannelConfigUlSpecificParametersConstraints)
			if err := logicalChannelConfigUlSpecificParametersSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := logicalChannelConfigUlSpecificParametersSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.Priority = v
			}
			{
				v, err := d.DecodeEnumerated(logicalChannelConfigUlSpecificParametersPrioritisedBitRateConstraints)
				if err != nil {
					return err
				}
				c.PrioritisedBitRate = v
			}
			{
				v, err := d.DecodeEnumerated(logicalChannelConfigUlSpecificParametersBucketSizeDurationConstraints)
				if err != nil {
					return err
				}
				c.BucketSizeDuration = v
			}
			if logicalChannelConfigUlSpecificParametersSeq.IsComponentPresent(3) {
				seqOf := d.NewSequenceOfDecoder(logicalChannelConfigUlSpecificParametersAllowedServingCellsConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.AllowedServingCells = make([]ServCellIndex, n)
				for i := int64(0); i < n; i++ {
					if err := c.AllowedServingCells[i].Decode(d); err != nil {
						return err
					}
				}
			}
			if logicalChannelConfigUlSpecificParametersSeq.IsComponentPresent(4) {
				seqOf := d.NewSequenceOfDecoder(logicalChannelConfigUlSpecificParametersAllowedSCSListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.AllowedSCS_List = make([]SubcarrierSpacing, n)
				for i := int64(0); i < n; i++ {
					if err := c.AllowedSCS_List[i].Decode(d); err != nil {
						return err
					}
				}
			}
			if logicalChannelConfigUlSpecificParametersSeq.IsComponentPresent(5) {
				c.MaxPUSCH_Duration = new(int64)
				v, err := d.DecodeEnumerated(logicalChannelConfigUlSpecificParametersMaxPUSCHDurationConstraints)
				if err != nil {
					return err
				}
				(*c.MaxPUSCH_Duration) = v
			}
			if logicalChannelConfigUlSpecificParametersSeq.IsComponentPresent(6) {
				c.ConfiguredGrantType1Allowed = new(int64)
				v, err := d.DecodeEnumerated(logicalChannelConfigUlSpecificParametersConfiguredGrantType1AllowedConstraints)
				if err != nil {
					return err
				}
				(*c.ConfiguredGrantType1Allowed) = v
			}
			if logicalChannelConfigUlSpecificParametersSeq.IsComponentPresent(7) {
				c.LogicalChannelGroup = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxLCG_ID))
				if err != nil {
					return err
				}
				(*c.LogicalChannelGroup) = v
			}
			if logicalChannelConfigUlSpecificParametersSeq.IsComponentPresent(8) {
				c.SchedulingRequestID = new(SchedulingRequestId)
				if err := (*c.SchedulingRequestID).Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.LogicalChannelSR_Mask = v
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.LogicalChannelSR_DelayTimerApplied = v
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "channelAccessPriority-r16", Optional: true},
				{Name: "bitRateMultiplier-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(logicalChannelConfigChannelAccessPriorityR16Constraints)
			if err != nil {
				return err
			}
			ie.ChannelAccessPriority_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(logicalChannelConfigExtBitRateMultiplierR16Constraints)
			if err != nil {
				return err
			}
			ie.BitRateMultiplier_r16 = &v
		}
	}

	return nil
}
