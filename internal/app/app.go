package app

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	cfg "github.com/maksewsha/wardrobe-back/cfg"
	"github.com/maksewsha/wardrobe-back/internal/repository"
	"github.com/maksewsha/wardrobe-back/internal/usecases"
	"github.com/maksewsha/wardrobe-back/internal/web"
	"github.com/maksewsha/wardrobe-back/pkg/pghelper"
)

func Start(config *cfg.Config) {
	connection, error := pgx.Connect(context.Background(), pghelper.NewConnURL(config.PGConfig))
	if error != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", error)
		os.Exit(1)
	}
	defer connection.Close(context.Background())
	uc := usecases.NewUserUseCase(repository.NewUserRepository(connection))
	handlers := web.NewHandlers(uc)
	r := gin.New()
	handlers.SetRouter(r)
	r.Run()
}
