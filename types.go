package main

type envType int

const (
	PrivateKey envType = iota
	PublicKey
	Messages
)

func (e *envType) String() string {
	switch *e {
	case PrivateKey:
		return "PRIVATE_KEY"
	case PublicKey:
		return "PUBLIC_KEY"
	case Messages:
		return "MESSAGES"
	}
	return "unknown"
}
