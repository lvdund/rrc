package ies

import "rrc/utils"

// PollPDU ::= ENUMERATED
type Pollpdu struct {
	Value utils.ENUMERATED
}

const (
	PollpduEnumeratedNothing = iota
	PollpduEnumeratedP4
	PollpduEnumeratedP8
	PollpduEnumeratedP16
	PollpduEnumeratedP32
	PollpduEnumeratedP64
	PollpduEnumeratedP128
	PollpduEnumeratedP256
	PollpduEnumeratedP512
	PollpduEnumeratedP1024
	PollpduEnumeratedP2048
	PollpduEnumeratedP4096
	PollpduEnumeratedP6144
	PollpduEnumeratedP8192
	PollpduEnumeratedP12288
	PollpduEnumeratedP16384
	PollpduEnumeratedP20480
	PollpduEnumeratedP24576
	PollpduEnumeratedP28672
	PollpduEnumeratedP32768
	PollpduEnumeratedP40960
	PollpduEnumeratedP49152
	PollpduEnumeratedP57344
	PollpduEnumeratedP65536
	PollpduEnumeratedInfinity
	PollpduEnumeratedSpare8
	PollpduEnumeratedSpare7
	PollpduEnumeratedSpare6
	PollpduEnumeratedSpare5
	PollpduEnumeratedSpare4
	PollpduEnumeratedSpare3
	PollpduEnumeratedSpare2
	PollpduEnumeratedSpare1
)
