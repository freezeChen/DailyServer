package main

import (
	"DailyServer/commons/glog"
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	MaxBodySize = uint32(1 << 10)
	OPer_check  = 0000
	OPer_msg    = 0001
)
const (
	// size
	PackSize      uint = 4
	HeaderSize    uint = 2
	VerSize       uint = 2
	OperationSize uint = 4
	SeqIdSize     uint = 4
	RawHeaderSize uint = PackSize + HeaderSize + VerSize + OperationSize + SeqIdSize
	MaxPackSize        = MaxBodySize + uint32(RawHeaderSize)
	// offset
	PackOffset      = 0
	HeaderOffset    = PackOffset + PackSize
	VerOffset       = HeaderOffset + HeaderSize
	OperationOffset = VerOffset + VerSize
	SeqIdOffset     = OperationOffset + OperationSize
)

var (
	MsgFinish = &Msg{Code: 3}
	MsgReady  = &Msg{Code: 2}

	ErrMsgPackLen   = errors.New("default server codec pack length error")
	ErrMsgHeaderLen = errors.New("default server codec header length error")
	ErrMsgNotCheck  = errors.New("connect not check")
)

/**
协议 [0 0 0 0, 0 0, 0 0, 0 0 0 0, 0 0 0 0, ...]
	包总长-	协议长度(16),版本,通讯代号,身份,消息体
 */
type Msg struct {
	Ver       uint16          `json:"ver"`
	Operation uint32          `json:"operation"`
	SeqId     uint32          `json:"seqId"`
	Body      json.RawMessage `json:"body"`
	Code      int32           `json:"code"`
}

func (m *Msg) ReadTCP(r *bufio.Reader) (err error) {
	var (
		bodyLen   uint32
		headerLen uint16
		packLen   uint32
		headBuf   []byte = make([]byte, RawHeaderSize)
		bodyBuf   []byte
	)
	glog.Info("readtcp start")
	//all, _ := ioutil.ReadAll(r)
	//fmt.Println("all", string(all))
	n, err := r.Read(headBuf)

	glog.Info("readtcp read")


	fmt.Println("head", headBuf)
	if n != int(RawHeaderSize) {
		err = ErrMsgHeaderLen
		return
	}
	packLen = binary.BigEndian.Uint32(headBuf[PackOffset:HeaderOffset])
	headerLen = binary.BigEndian.Uint16(headBuf[HeaderOffset:VerOffset])
	m.Ver = binary.BigEndian.Uint16(headBuf[VerOffset:OperationOffset])
	m.Operation = binary.BigEndian.Uint32(headBuf[OperationOffset:SeqIdOffset])
	m.SeqId = binary.BigEndian.Uint32(headBuf[SeqIdOffset:])

	if packLen > MaxPackSize {
		return ErrMsgPackLen
	}
	if uint(headerLen) != RawHeaderSize {
		return ErrMsgHeaderLen
	}
	if bodyLen = packLen - uint32(headerLen); bodyLen > 0 {
		bodyBuf = make([]byte, bodyLen)
		_, err = r.Read(bodyBuf)
		m.Body = bodyBuf
	} else {
		m.Body = nil
	}

	fmt.Println(string(m.Body))

	return
}

func (m *Msg) WriteTCP(w *bufio.Writer) (err error) {
	var (
		packLen uint32
	)
	packLen = uint32(RawHeaderSize) + uint32(len(m.Body))
	binary.Write(w, binary.BigEndian, packLen)
	binary.Write(w, binary.BigEndian, uint16(RawHeaderSize))
	binary.Write(w, binary.BigEndian, m.Ver)
	binary.Write(w, binary.BigEndian, m.Operation)
	binary.Write(w, binary.BigEndian, m.SeqId)
	binary.Write(w, binary.BigEndian, m.Body)
	err = w.Flush()
	return
}

func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}
