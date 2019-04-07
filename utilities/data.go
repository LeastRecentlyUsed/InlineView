package utilities

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var timeout = 5 * time.Second

// AddPriceRecord stores a single price record in the database
func AddPriceRecord(cli *clientv3.Client, key string, rec string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	_, err := cli.Put(ctx, key, rec)
	cancel()
	if err != nil {
		fmt.Println("Error client v3 put failed")
		return false
	}
	return true
}

// CreateEtcdClient creates a new V3 client and returns it
func CreateEtcdClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// AddPriceRecordold stores each price record in the etcd database
func addpricerecordold(key string, rec string) bool {
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
