package commands

import (
    "context"
    "fmt"
    pb "github.com/mix-go/mix-grpc-skeleton/protos"
    "google.golang.org/grpc"
    "time"
)

type GrpcClientCommand struct {
}

func (t *GrpcClientCommand) Main() {
    ctx, _ := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
    conn, err := grpc.DialContext(ctx, Addr, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        panic(err)
    }
    defer func() {
        _ = conn.Close()
    }()

    cli := pb.NewUserClient(conn)
    req := pb.AddRequest{
        Name: "xiaoliu",
    }
    resp, err := cli.Add(ctx, &req)
    if err != nil {
        panic(err)
    }
    fmt.Println(fmt.Sprintf("Add : %d", resp.UserId))
}