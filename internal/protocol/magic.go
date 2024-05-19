package protocol

const (
	// ClientMagic is the magic number that a client sends to the server when establishing a connection.
	ClientMagic uint64 = 0x6e697863
	// ServerMagic is the magic number that a server sends to the client when establishing a connection.
	ServerMagic uint64 = 0x6478696f
)
