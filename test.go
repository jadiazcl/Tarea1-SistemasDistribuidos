package main

import(
  "fmt"
  "time"
  "log"
  "net"
  "google.golang.org/grpc"
  "github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
  )

  const (
  	port = ":50051"
  )



  type orden struct {
      created_time time.Time
      id_paquete string
      tipo string
      nombre string
      valor  int
      origen string
      destino string
      seguimiento int
  }

func NewOrden(ordenes []*orden, id_paquete string, tipo string, nombre string,
  valor  int, origen string, destino string ) *orden {
    orden := orden{id_paquete: id_paquete,tipo:tipo,nombre:nombre,valor:valor,
    origen:origen,destino:destino}
    orden.created_time = time.Now()
    orden.seguimiento = NewCodeSeguimiento(ordenes)
    return &orden
}

func NewCodeSeguimiento(ordenes []*orden) int{
    if len(ordenes)==0 {
      return 1
    }
    return ordenes[len(ordenes)-1].seguimiento+1
}


func main() {
    fmt.Println("Go gRPC Beginners Tutorial!")

  	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
  	if err != nil {
  		log.Fatalf("failed to listen: %v", err)
  	}

  	s := chat.Server{}

  	grpcServer := grpc.NewServer()

  	chat.RegisterChatServiceServer(grpcServer, &s)

  	if err := grpcServer.Serve(lis); err != nil {
  		log.Fatalf("failed to serve: %s", err)
  	}
    fmt.Println("Wena profe")
    ordenes := []*orden{}
    aux:=NewOrden(ordenes,"Paquete1","mochila","Jorgekun",1000,"chilito","membrillo")
    ordenes=append(ordenes,aux)
    aux=NewOrden(ordenes,"Paquete2","Bebida","Iñakikun",2000,"chilito","Corea")
    ordenes=append(ordenes,aux)
    for i := 0; i < len(ordenes); i++ {
      fmt.Println(ordenes[i])
      fmt.Println(ordenes[i].created_time.Format(time.ANSIC))
      fmt.Println("////")
    }
    //fmt.Println(aux.created_time)
    //fmt.Println(aux.created_time.Format(time.ANSIC))
}