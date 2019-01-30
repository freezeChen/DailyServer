package im

var Options *options

type options struct {
	TCPPort        string
	WebSocketPort  string
	TCPKeepalive   bool
	TCPReadBuffer  int
	TCPWriteBuffer int
}

func InitIMConfig() {
	Options = &options{
		TCPPort:        ":8020",
		WebSocketPort:  ":8030",
		TCPKeepalive:   true,
		TCPReadBuffer:  1024,
		TCPWriteBuffer: 1024,
	}
}
