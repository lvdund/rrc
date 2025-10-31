package ies

// SoundingRS-UL-ConfigDedicatedAperiodic-v1310-setup ::= SEQUENCE
type SoundingrsUlConfigdedicatedaperiodicV1310Setup struct {
	SrsConfigapdciFormat4V1310 *[]SrsConfigapV1310 `lb:1,ub:3`
	SrsActivateapV1310         *SoundingrsUlConfigdedicatedaperiodicV1310SetupSrsActivateapV1310
}
