package postgres

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"post-tech-challenge-10soat/internal/infra/config"

	"github.com/golang-migrate/migrate/v4"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

func New(ctx context.Context, config *config.DB) (*DB, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	log.Printf("Connecting to database: %s", maskCredentials(url))
	db, err := pgxpool.New(ctx, url)
	if err != nil {
		log.Fatalf("Error creating connection pool: %v", err)
		return nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
		return nil, err
	}
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return &DB{
		db,
		&psql,
		url,
	}, nil
}

func (db *DB) Migrate() error {
	driver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", driver, db.url)
	if err != nil {
		return err
	}
	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	log.Println("Migrations applied successfully")
	return nil
}

func maskCredentials(url string) string {
	return os.ExpandEnv("${DB_CONNECTION}://****:****@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable")
}
