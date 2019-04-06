package utilities

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var timeout = 5 * time.Second

// AddPriceRecord stores each price record in the etcd database
func AddPriceRecord(key string, rec string) bool {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		fmt.Println("Error create etcd client v3.")
		return false
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	_, err = cli.Put(ctx, key, rec)
	cancel()
	if err != nil {
		fmt.Println("Error client v3 put failed")
	}
	return true
}
