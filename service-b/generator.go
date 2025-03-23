package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pcandido/rest-vs-grpc/proto/pb"
)

func GenerateComplexData() *pb.ComplexDataResponse {
	attributes := make(map[string]string)
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key_%d", i)
		attributes[key] = fmt.Sprintf("value_%d", rand.Intn(1000))
	}

	flags := make(map[string]bool)
	for i := 0; i < 5; i++ {
		flags[fmt.Sprintf("flag_%d", i)] = rand.Intn(2) == 0
	}

	items := make([]*pb.Item, 100)
	for i := range items {
		items[i] = &pb.Item{
			Id:     fmt.Sprintf("item_%d", i),
			Amount: rand.Float64() * 1000,
			Active: rand.Intn(2) == 0,
		}
	}

	transactions := make([]*pb.Transaction, 50)
	for i := range transactions {
		transactions[i] = &pb.Transaction{
			TransactionId: fmt.Sprintf("txn_%d", i),
			Value:         rand.Float64() * 5000,
			Timestamp:     time.Now().Add(time.Duration(-rand.Intn(1000)) * time.Hour).Format(time.RFC3339),
		}
	}

	nestedComplex := make([]*pb.NestedComplex, 30)
	for i := range nestedComplex {
		nestedComplex[i] = &pb.NestedComplex{
			Index:  int32(i),
			Weight: rand.Float32() * 100,
			Detail: &pb.Detail{
				Description: fmt.Sprintf("detail_%d", i),
				Level:       rand.Int31(),
				Valid:       rand.Intn(2) == 0,
			},
		}
	}

	return &pb.ComplexDataResponse{
		Id:         fmt.Sprintf("id_%d", rand.Intn(10000)),
		Name:       fmt.Sprintf("name_%d", rand.Intn(10000)),
		Timestamp:  time.Now().Format(time.RFC3339),
		Attributes: attributes,
		Metadata: &pb.Metadata{
			Version:     rand.Int31(),
			Tags:        []string{"tag1", "tag2", "tag3"},
			Source:      "auto-generated",
			Permissions: []int64{1, 2, 3, 4, 5},
		},
		Items:         items,
		Transactions:  transactions,
		Flags:         flags,
		LargePayload:  make([]byte, 1000),
		NestedComplex: nestedComplex,
	}
}
