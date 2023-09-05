package pghelper

import (
	"fmt"

	"github.com/maksewsha/wardrobe-back/cfg"
)

func NewConnURL(pgConfig *cfg.PGConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pgConfig.Username, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.DBName)
}
