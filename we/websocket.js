var ws;

function send() {
   ws.send("fsdf")
}

function link() {
    ws = new WebSocket("ws://localhost:8888/ws");
    ws.onopen = function () {
        console.log("open");
    };
    ws.onmessage = function (evt) {
        console.log(evt.data)
    };
    ws.onclose = function (evt) {
        console.log("close");
    };
    ws.onerror = function (evt) {
        console.log("onerror")
    };
}