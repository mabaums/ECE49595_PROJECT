package main

import (
	// ssh "/ssh"
	// b "ECE49595_PROJECT/basic"
	// ssh "ECE49595_PROJECT/ssh"
	q "ECE49595_PROJECT/queue"
	// // "context"
	d "ECE49595_PROJECT/dock"
	"github.com/go-redis/redis"
	"fmt"
)

func main(){

	fmt.Println(d.CreateNewContainer("redis/redis-stack","127.0.0.1","6379","6379"))
	q.MakeQueue(&redis.Options{ 
		Addr: "localhost:6379", 
		Password: "", 
		 DB: 0, 
   }, &redis.Options{ 
	Addr: "localhost:6379", 
	Password: "", 
	 DB: 0, 
} )
	err := q.AddRequestToQueue( "test94", q.Queue_Request{NAME: "sabash", 
	EMAIL: "sabutdana@gmail.com",
	CURRENT_IP: "mei nahi bataonga",
	LOCATION: "jhadio k peeche",
	CREATED_AT: "cake murder day", 
	LASTSEEN:"don ko dhundna mushkil hi nahi namumkin hai"})
	
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(q.GetRequestFromQueue( "test94"))
	fmt.Println(q.RemoveRequestFromQueue("test94"))
	fmt.Println(q.GetRequestFromQueue( "test94"))


}

