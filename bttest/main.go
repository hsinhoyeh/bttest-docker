package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/cloud"
	bgt "google.golang.org/cloud/bigtable"
	"google.golang.org/cloud/bigtable/bttest"
	"google.golang.org/grpc"
)

const (
	defaultTimeout = 5 * time.Second
)

func adminClient(addr string) (*bgt.AdminClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), defaultTimeout)
	adminClient, err := bgt.NewAdminClient(
		ctx,
		"proj",
		"zone",
		"cluster",
		cloud.WithBaseGRPC(conn))
	if err != nil {
		return nil, err
	}

	return adminClient, nil
}

var (
	// FIXME: need to parse schema.hbase to create default tables
	// tables is keyed by $tablename, the value is a slice of column family
	tables = map[string][]string{
		"foo": []string{"bar"},
	}
)

func initTables(addr string) error {
	client, err := adminClient(addr)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), defaultTimeout)
	for table, cFams := range tables {
		if err := client.CreateTable(ctx, table); err != nil {
			// we log only instead of fatal directly
			log.Printf("create table failed:%v\n", err)
		}
		for _, cFam := range cFams {
			if err := client.CreateColumnFamily(ctx, table, cFam); err != nil {
				log.Printf("create column family failed:%v\n", err)
			}

			// this is used to minic hbase behavior where only one recent result is returned
			if err := client.SetGCPolicy(ctx, table, cFam, bgt.MaxVersionsPolicy(1)); err != nil {
				log.Printf("set gcpolicy failed:%v\n", err)
			}
		}
	}
	return nil
}

func main() {
	svr, err := bttest.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("init tables..")
	if err := initTables(svr.Addr); err != nil {
		log.Fatal(err)
	}

	log.Printf("ready to serve address:%v\n", svr.Addr)
	neverStop := make(chan int)
	<-neverStop
}
