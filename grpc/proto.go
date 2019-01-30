/*
   @Time : 2019-01-30 16:42
   @Author : frozenchen
   @File : proto.go
   @Software: DailyServer
*/
package grpc

import (
	"DailyServer/commons/glog"
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
)

const (
	MaxBodySize = uint32(1 << 10)
)
const (
	// size
	PackSize      = 4
	HeaderSize    = 2
	VerSize       = 2
	OperationSize = 4
	SeqIdSize     = 4
	RawHeaderSize = PackSize + HeaderSize + VerSize + OperationSize + SeqIdSize
	MaxPackSize   = MaxBodySize + uint32(RawHeaderSize)
	// offset
	PackOffset      = 0
	HeaderOffset    = PackOffset + PackSize
	VerOffset       = HeaderOffset + HeaderSize
	OperationOffset = VerOffset + VerSize
	SeqIdOffset     = OperationOffset + OperationSize
	BodyOffset      = SeqIdOffset + SeqIdSize
)
const (
	// OpHandshake handshake
	OpHandshake = int32(0)
	// OpHandshakeReply handshake reply
	OpHandshakeReply = int32(1)

	// OpHeartbeat heartbeat
	OpHeartbeat = int32(2)
	// OpHeartbeatReply heartbeat reply
	OpHeartbeatReply = int32(3)

	// OpSendMsg send message.
	OpSendMsg = int32(4)
	// OpSendMsgReply  send message reply
	OpSendMsgReply = int32(5)

	// OpDisconnectReply disconnect reply
	OpDisconnectReply = int32(6)

	// OpAuth auth connnect
	OpAuth = int32(7)
	// OpAuthReply auth connect reply
	OpAuthReply = int32(8)

	// OpRaw raw message
	OpRaw = int32(9)

	// OpProtoReady proto ready
	OpProtoReady = int32(10)
	// OpProtoFinish proto finish
	OpProtoFinish = int32(11)

	// OpChangeRoom change room
	OpChangeRoom = int32(12)
	// OpChangeRoomReply change room reply
	OpChangeRoomReply = int32(13)

	// OpSub subscribe operation
	OpSub = int32(14)
	// OpSubReply subscribe operation
	OpSubReply = int32(15)

	// OpUnsub unsubscribe operation
	OpUnsub = int32(16)
	// OpUnsubReply unsubscribe operation reply
	OpUnsubReply = int32(17)
)

var (
	ProtoFinish = &Proto{Opr: OpProtoFinish}
	ProtoReady  = &Proto{Opr: OpProtoReady}

	ErrMsgPackLen   = errors.New("default Server codec pack length error")
	ErrMsgHeaderLen = errors.New("default Server codec header length error")
	ErrMsgNotCheck  = errors.New("connect not check")
)

/**
协议 [0 0 0 0, 0 0, 0 0, 0 0 0 0, 0 0 0 0, ...]
	包总长-	协议长度(16),版本,通讯代号,身份,消息体
*/
/*type Msg struct {
	Ver       uint16          `json:"ver"`
	Operation uint32          `json:"operation"`
	SeqId     uint32          `json:"seqId"`
	Body      json.RawMessage `json:"body"`
	Code      int32           `json:"code"`
}
*/
func (p *Proto) ReadTCP(r *bufio.Reader) (err error) {
	var (
		bodyLen   uint32
		headerLen uint16
		packLen   uint32
		headBuf   []byte = make([]byte, RawHeaderSize)
		bodyBuf   []byte
	)

	n, err := r.Read(headBuf)
	fmt.Println("head", headBuf)
	if n != int(RawHeaderSize) {
		err = ErrMsgHeaderLen
		return
	}
	packLen = binary.BigEndian.Uint32(headBuf[PackOffset:HeaderOffset])
	headerLen = binary.BigEndian.Uint16(headBuf[HeaderOffset:VerOffset])
	p.Ver = int32(binary.BigEndian.Uint16(headBuf[VerOffset:OperationOffset]))
	p.Opr = int32(binary.BigEndian.Uint32(headBuf[OperationOffset:SeqIdOffset]))
	p.Id = int32(binary.BigEndian.Uint32(headBuf[SeqIdOffset:]))

	glog.Info(p)

	if packLen > MaxPackSize {
		return ErrMsgPackLen
	}
	if uint(headerLen) != RawHeaderSize {
		return ErrMsgHeaderLen
	}
	if bodyLen = packLen - uint32(headerLen); bodyLen > 0 {
		bodyBuf = make([]byte, bodyLen)
		_, err = r.Read(bodyBuf)
		p.Body = bodyBuf
	} else {
		p.Body = nil
	}

	return
}

func (p *Proto) WriteTCP(w *bufio.Writer) (err error) {
	var (
		packLen uint32
	)
	packLen = uint32(RawHeaderSize) + uint32(len(p.Body))
	binary.Write(w, binary.BigEndian, packLen)
	binary.Write(w, binary.BigEndian, uint16(RawHeaderSize))
	binary.Write(w, binary.BigEndian, uint16(p.Ver))
	binary.Write(w, binary.BigEndian, uint32(p.Opr))
	binary.Write(w, binary.BigEndian, uint32(p.Id))
	binary.Write(w, binary.BigEndian, p.Body)
	err = w.Flush()
	return
}

func (p *Proto) ReadWebSocket(ws *websocket.Conn) (err error) {
	var (
		bodyLen   uint32
		headerLen uint16
		packLen   uint32
		allBuf    []byte
	)

	_, allBuf, err = ws.ReadMessage()
	if err != nil {
		return
	}

	if len(allBuf) < (RawHeaderSize) {
		return ErrMsgHeaderLen
	}

	packLen = binary.BigEndian.Uint32(allBuf[PackOffset:HeaderOffset])
	headerLen = binary.BigEndian.Uint16(allBuf[HeaderOffset:VerOffset])
	p.Ver = int32(binary.BigEndian.Uint16(allBuf[VerOffset:OperationOffset]))
	p.Opr = int32(binary.BigEndian.Uint32(allBuf[OperationOffset:SeqIdOffset]))
	p.Id = int32(binary.BigEndian.Uint32(allBuf[SeqIdOffset:]))

	if packLen > MaxPackSize {
		return ErrMsgPackLen
	}

	if headerLen != RawHeaderSize {
		return ErrMsgHeaderLen
	}
	if bodyLen = packLen - uint32(headerLen); bodyLen > 0 {

		p.Body = allBuf[headerLen:packLen]
		glog.Info("websocket msg:", string(p.Body))
	} else {
		p.Body = nil
	}
	return nil
}

func (p *Proto) WriteWebSocket(ws *websocket.Conn) (err error) {
	var (
		buf     = make([]byte, RawHeaderSize)
		packLen = uint32(RawHeaderSize) + uint32(len(p.Body))
	)
	binary.BigEndian.PutUint32(buf, packLen)
	binary.BigEndian.PutUint16(buf[HeaderOffset:], RawHeaderSize)
	binary.BigEndian.PutUint16(buf[VerOffset:], uint16(p.Ver))
	binary.BigEndian.PutUint32(buf[OperationOffset:], uint32(p.Opr))
	binary.BigEndian.PutUint32(buf[SeqIdOffset:], uint32(p.Id))

	fmt.Println(packLen, len(p.Body), len(buf))

	buf = append(buf, p.Body...)
	fmt.Println(len(buf))

	err = ws.WriteMessage(websocket.BinaryMessage, buf)
	return
}
