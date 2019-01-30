package im
//
//import (
//	"DailyServer/commons/glog"
//	"bufio"
//	"encoding/binary"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/gorilla/websocket"
//)
//
//const (
//	MaxBodySize = uint32(1 << 10)
//	OPer_check  = 0000
//	OPer_msg    = 0001
//)
//const (
//	// size
//	PackSize      = 4
//	HeaderSize    = 2
//	VerSize       = 2
//	OperationSize = 4
//	SeqIdSize     = 4
//	RawHeaderSize = PackSize + HeaderSize + VerSize + OperationSize + SeqIdSize
//	MaxPackSize   = MaxBodySize + uint32(RawHeaderSize)
//	// offset
//	PackOffset      = 0
//	HeaderOffset    = PackOffset + PackSize
//	VerOffset       = HeaderOffset + HeaderSize
//	OperationOffset = VerOffset + VerSize
//	SeqIdOffset     = OperationOffset + OperationSize
//	BodyOffset      = SeqIdOffset + SeqIdSize
//)
//
//var (
//	MsgFinish = &Msg{Code: 3}
//	MsgReady  = &Msg{Code: 2}
//
//	ErrMsgPackLen   = errors.New("default Server codec pack length error")
//	ErrMsgHeaderLen = errors.New("default Server codec header length error")
//	ErrMsgNotCheck  = errors.New("Online not check")
//)
//
///**
//协议 [0 0 0 0, 0 0, 0 0, 0 0 0 0, 0 0 0 0, ...]
//	包总长-	协议长度(16),版本,通讯代号,身份,消息体
//*/
//type Msg struct {
//	Ver       uint16          `json:"ver"`
//	Operation uint32          `json:"operation"`
//	SeqId     uint32          `json:"seqId"`
//	Body      json.RawMessage `json:"body"`
//	Code      int32           `json:"code"`
//}
//
//func (m *Msg) ReadTCP(r *bufio.Reader) (err error) {
//	var (
//		bodyLen   uint32
//		headerLen uint16
//		packLen   uint32
//		headBuf   []byte = make([]byte, RawHeaderSize)
//		bodyBuf   []byte
//	)
//
//	n, err := r.Read(headBuf)
//	fmt.Println("head", headBuf)
//	if n != int(RawHeaderSize) {
//		err = ErrMsgHeaderLen
//		return
//	}
//	packLen = binary.BigEndian.Uint32(headBuf[PackOffset:HeaderOffset])
//	headerLen = binary.BigEndian.Uint16(headBuf[HeaderOffset:VerOffset])
//	m.Ver = binary.BigEndian.Uint16(headBuf[VerOffset:OperationOffset])
//	m.Operation = binary.BigEndian.Uint32(headBuf[OperationOffset:SeqIdOffset])
//	m.SeqId = binary.BigEndian.Uint32(headBuf[SeqIdOffset:])
//
//	if packLen > MaxPackSize {
//		return ErrMsgPackLen
//	}
//	if uint(headerLen) != RawHeaderSize {
//		return ErrMsgHeaderLen
//	}
//	if bodyLen = packLen - uint32(headerLen); bodyLen > 0 {
//		bodyBuf = make([]byte, bodyLen)
//		_, err = r.Read(bodyBuf)
//		m.Body = bodyBuf
//	} else {
//		m.Body = nil
//	}
//
//	fmt.Println(string(m.Body))
//
//	return
//}
//
//func (m *Msg) WriteTCP(w *bufio.Writer) (err error) {
//	var (
//		packLen uint32
//	)
//	packLen = uint32(RawHeaderSize) + uint32(len(m.Body))
//	binary.Write(w, binary.BigEndian, packLen)
//	binary.Write(w, binary.BigEndian, uint16(RawHeaderSize))
//	binary.Write(w, binary.BigEndian, m.Ver)
//	binary.Write(w, binary.BigEndian, m.Operation)
//	binary.Write(w, binary.BigEndian, m.SeqId)
//	binary.Write(w, binary.BigEndian, m.Body)
//	err = w.Flush()
//	return
//}
//
//func (m *Msg) ReadWebSocket(ws *websocket.Conn) (err error) {
//	var (
//		bodyLen   uint32
//		headerLen uint16
//		packLen   uint32
//		allBuf    []byte
//	)
//
//	_, allBuf, err = ws.ReadMessage()
//	if err != nil {
//		return
//	}
//
//	if len(allBuf) < (RawHeaderSize) {
//		return ErrMsgHeaderLen
//	}
//
//	packLen = binary.BigEndian.Uint32(allBuf[PackOffset:HeaderOffset])
//	headerLen = binary.BigEndian.Uint16(allBuf[HeaderOffset:VerOffset])
//	m.Ver = binary.BigEndian.Uint16(allBuf[VerOffset:OperationOffset])
//	m.Operation = binary.BigEndian.Uint32(allBuf[OperationOffset:SeqIdOffset])
//	m.SeqId = binary.BigEndian.Uint32(allBuf[SeqIdOffset:])
//
//	if packLen > MaxPackSize {
//		return ErrMsgPackLen
//	}
//
//	if headerLen != RawHeaderSize {
//		return ErrMsgHeaderLen
//	}
//	if bodyLen = packLen - uint32(headerLen); bodyLen > 0 {
//
//		m.Body = allBuf[headerLen:packLen]
//		glog.Info("websocket msg:", string(m.Body))
//	} else {
//		m.Body = nil
//	}
//	return nil
//}
//
//func (m *Msg) WriteWebSocket(ws *websocket.Conn) (err error) {
//	var (
//		buf     = make([]byte, RawHeaderSize)
//		packLen = uint32(RawHeaderSize) + uint32(len(m.Body))
//	)
//	binary.BigEndian.PutUint32(buf, packLen)
//	binary.BigEndian.PutUint16(buf[HeaderOffset:], RawHeaderSize)
//	binary.BigEndian.PutUint16(buf[VerOffset:], uint16(m.Ver))
//	binary.BigEndian.PutUint32(buf[OperationOffset:], m.Operation)
//	binary.BigEndian.PutUint32(buf[SeqIdOffset:], m.SeqId)
//
//	fmt.Println(packLen, len(m.Body), len(buf))
//
//	buf = append(buf, m.Body...)
//	fmt.Println(len(buf))
//
//	err = ws.WriteMessage(websocket.BinaryMessage, buf)
//	return
//}
