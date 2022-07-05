package types

type ConcurrencyChannels struct {
	Stateful  *ConcurrencyControl
	Stateless *ConcurrencyControl
	Direct    *ConcurrencyControl
}
