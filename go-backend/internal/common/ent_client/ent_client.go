package ent_client

import (
	"context"
	"fmt"
	"go-backend/ent"
	"go-backend/internal/common/env"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func New(env *env.Env) *ent.Client {
	client, err := ent.Open("mysql", env.DatabaseUrl)
	if err != nil {
		log.Fatalf("❌ [ENT] failed opening connection to mysql: %v", err)
	}

	ctx := context.Background()

	// bật tính tăng sử dụng câu lệnh sql trong code (--feature sql/execquery)
	// patth file: go-backend/ent/generate.go
	// //go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/execquery ./schema
	_, err = client.QueryContext(ctx, "SELECT 1 + 1")
	if err != nil {
		log.Fatalf("❌ [ENT] failed connection to mysql: %v", err)
	}

	fmt.Println("✅ [ENT] Connection To MySQL Successfully")

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("❌ [ENT] failed creating schema resources: %v", err)
	}

	return client
}
