var ws;
const rawHeaderLen = 16;
const packetOffset = 0;
const headerOffset = 4;
const verOffset = 6;
const opOffset = 8;
const seqOffset = 12;
var textEncoder = new TextEncoder();
var textDecoder = new TextDecoder();

function auth() {
    var bodyBuf = textEncoder.encode("5");
    let headBuf = new ArrayBuffer(rawHeaderLen);
    var headerView = new DataView(headBuf, 0);
    headerView.setInt32(packetOffset, rawHeaderLen + bodyBuf.byteLength);
    headerView.setInt16(headerOffset, rawHeaderLen);
    headerView.setInt16(verOffset, 1);
    headerView.setInt32(opOffset, 7);
    headerView.setInt32(seqOffset, 1);
    ws.send(mergeArrayBuffer(headBuf, bodyBuf));
}


function send() {
    var input = $("#msg").val();


    console.log(input);
    var msg = {
        id: 5,
        sid: 5,
        msg: input
    };

    var jsonStr = JSON.stringify(msg);
    var msgBuf = textEncoder.encode(jsonStr);


    /**
     协议 [0 0 0 0, 0 0, 0 0, 0 0 0 0, 0 0 0 0, ...]
     包总长-    协议长度(16),版本,通讯代号,身份,消息体
     */


    let headBuf = new ArrayBuffer(rawHeaderLen);
    var headerView = new DataView(headBuf, 0);
    headerView.setInt32(packetOffset, rawHeaderLen + msgBuf.byteLength);
    headerView.setInt16(headerOffset, rawHeaderLen);
    headerView.setInt16(verOffset, 1);
    headerView.setInt32(opOffset, 7);
    headerView.setInt32(seqOffset, 1);

    ws.send(mergeArrayBuffer(headBuf, msgBuf));

}


function mergeArrayBuffer(ab1, ab2) {
    var u81 = new Uint8Array(ab1),
        u82 = new Uint8Array(ab2),
        res = new Uint8Array(ab1.byteLength + ab2.byteLength);
    res.set(u81, 0);
    res.set(u82, ab1.byteLength);
    return res.buffer;
}

function link() {
    ws = new WebSocket("ws://localhost:8888/ws");
    ws.binaryType = 'arraybuffer';
    ws.onopen = function () {
        console.log("open");
        // auth();
    };
    ws.onmessage = function (evt) {
        var data = evt.data;
        var dataView = new DataView(data, 0);
        var packetLen = dataView.getInt32(packetOffset);
        var headerLen = dataView.getInt16(headerOffset);
        var ver = dataView.getInt16(verOffset);
        var op = dataView.getInt32(opOffset);
        var seq = dataView.getInt32(seqOffset);
        var msgBody = '{"id":5,"sid":5,"msg":"发送内容"}';

        msgBody = textDecoder.decode(data.slice(headerLen, packetLen));
        var ss = '{"id":5,"sid":5,"msg":"发送内容"}';
        console.log(msgBody.length, ss.length);

        console.log("receiveHeader: packetLen=" + packetLen, "headerLen=" + headerLen, "ver=" + ver, "op=" + op, "seq=" + seq, "body=|" + msgBody.toString() + "|");
        // var parseJSON = jQuery.parseJSON(msgBody.trim());
        var parse = JSON.parse(msgBody);
        console.log(parse)

    };
    ws.onclose = function (evt) {
        console.log("close");
    };
    ws.onerror = function (evt) {
        console.log("onerror")
    };
}

function byteToString(arr) {
    if (typeof arr === 'string') {
        return arr;
    }
    var str = '',
        _arr = arr;
    for (var i = 0; i < _arr.length; i++) {
        var one = _arr[i].toString(2),
            v = one.match(/^1+?(?=0)/);
        if (v && one.length == 8) {
            var bytesLength = v[0].length;
            var store = _arr[i].toString(2).slice(7 - bytesLength);
            for (var st = 1; st < bytesLength; st++) {
                store += _arr[st + i].toString(2).slice(2);
            }
            str += String.fromCharCode(parseInt(store, 2));
            i += bytesLength - 1;
        } else {
            str += String.fromCharCode(_arr[i]);
        }
    }
    return str;

}