package main  
  
import (  
   "github.com/gin-gonic/gin"  
   "github.com/gorilla/websocket" 
   "net/http"
   "fmt"
   // "time"
)  
  
  
var upGrader = websocket.Upgrader{  
   CheckOrigin: func (r *http.Request) bool {  
      return true  
   },  
}  
  
//webSocket请求ping 返回pong  
func ping(c *gin.Context) {  
   //升级get请求为webSocket协议

   ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)  
   if err != nil {  
      return  
   }  
   // defer ws.Close()  
   go func() {
      for {
         //读取ws中的数据  
         mt, message, err := ws.ReadMessage()  
         if err != nil {  
            fmt.Println("read",err)  
            break  
         }  
         if string(message) == "ping" {  
            message = []byte("pong")  
         }  
         //写入ws数据
         err = ws.WriteMessage(mt, message)  
         if err != nil {
            fmt.Println("Writer",err)  
            break  
         }  

         err = ws.WriteMessage(mt, []byte("test111"))
      }  
   }()

   fmt.Println("---end-11-----")
   // time.Sleep(3*time.Second)
   fmt.Println("---end-22-----")
}  
    
func main() {  
   fmt.Println("---------------",websocket.PingMessage)
   bindAddress := "0.0.0.0:2303"  
   r := gin.Default()  
   r.GET("/ping", ping)  
   r.Run(bindAddress)  
}