package util

type ClientStateEnum int

const (
	_ ClientStateEnum = iota
	ClientStateNotVerify
	ClientStateVerified
	ClientStateInited
	ClientStateClosed
)
