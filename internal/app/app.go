package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	cfg "github.com/maksewsha/wardrobe-back/cfg"
	rest "github.com/maksewsha/wardrobe-back/internal/web"
	"github.com/maksewsha/wardrobe-back/pkg/pghelper"
)

func Start(config *cfg.Config) {
	connection, error := pgx.Connect(context.Background(), pghelper.NewConnURL(config.PGConfig))
	if error != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", error)
		os.Exit(1)
	}
	defer connection.Close(context.Background())
	if err := rest.RunRouter(); err != nil {
		log.Fatal(err)
	}
}
