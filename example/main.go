/*
   @Time : 2019-01-30 14:53
   @Author : frozenchen
   @File : main
   @Software: DailyServer
*/
package example

func main() {


	//glog.InitLogger()
	//conn, err := net.Dial("tcp", "127.0.0.1:8020")
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//writer := bufio.NewWriter(conn)
	//reader := bufio.NewReader(conn)
	//rmsg := new(lib)
	//go func() {
	//	for {
	//		err := rmsg.ReadTCP(reader)
	//		if err != nil {
	//			t.Error(err)
	//		}
	//		var info lib.Info
	//
	//		json.Unmarshal(rmsg.Body, &info)
	//
	//		fmt.Printf("info:%+v", info)
	//
	//	}
	//}()
	//
	//msg := new(Msg)
	//
	//msg.Body = []byte("5")
	//
	//err = msg.WriteTCP(writer)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//info := new(lib.Info)
	//info.Id = 5
	//info.Sid = 5
	//info.Msg = "hello i am client"
	//
	//bytes, _ := json.Marshal(info)
	//
	//msg.Body = bytes
	//
	//msg.WriteTCP(writer)

	select {}
}
