package main

import (
	"github.com/puneet105/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main(){
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Couldn't connect with grpc server %v",err)
	}
	defer conn.Close()

	/*client := proto.NewDemoServiceClient(conn)

	names := &proto.NameList{
		Name: []string{"foo","boo","too"},
	}*/

	//callSayHello(client)
	//callSayHelloStream(client, names)
	//callSayHelloClientStream(client, names)
	//callSayHelloBidirStream(client, names)

	client := proto.NewEmployeeServiceClient(conn)
	/*employee := &proto.EmployeeRequest{
		Id: 1,
		Name: "Puneet Sharma",
		Email: "puneet@gmail.com",
		Phone: "9677355991",
	}
*/
	//callCreateEmployee(client, employee)

	/*employee = &proto.EmployeeRequest{
		Id: 2,
		Name: "Mahesh Sharma",
		Email: "mahesh@gmail.com",
		Phone: "9828279408",
	}*/

	//callCreateEmployee(client, employee)
	time.Sleep(2*time.Second)
	var employees = &[]proto.EmployeeRequest{
		{	Id: 3,
			Name: "abc",
			Email: "abc@gmail.com",
			Phone: "9999999999",
		},
		{	Id: 4,
			Name: "xyz",
			Email: "axyzbc@gmail.com",
			Phone: "8888888888",
		},
		{	Id: 5,
			Name: "asd",
			Email: "asd@gmail.com",
			Phone: "7777777777",
		},
	}

	callCreateMultipleEmployee(client, employees)
	callGetEmployees(client)

}