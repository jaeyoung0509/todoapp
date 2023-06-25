package drivers

import (
	"context"
	"fmt"
	"log"

	"github.com/jaeyoung0509/todo/ent"
	"github.com/jaeyoung0509/todo/util"
)

func NewClient() (*ent.Client, error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Can't load config: %v", err)
	}
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", config.User, config.Pass, config.Host, config.Port, config.DBname, config.SslMode)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to db %v", err)
		return nil, err
	}
	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	return client, nil
}
