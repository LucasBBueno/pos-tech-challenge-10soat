package postgres

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"

	"post-tech-challenge-10soat/internal/adapter/config"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

func New(ctx context.Context, config *config.DB) (*DB, error) {
	// Criação da URL de conexão
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	// Log da URL de conexão para debugging (oculte credenciais em produção)
	log.Printf("Connecting to database: %s", maskCredentials(url))

	// Criação do pool de conexões
	db, err := pgxpool.New(ctx, url)
	if err != nil {
		log.Fatalf("Error creating connection pool: %v", err)
		return nil, err
	}

	// Testa a conexão com o banco
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
	// Executa as migrações
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

// Função auxiliar para mascarar credenciais da URL de conexão nos logs
func maskCredentials(url string) string {
	return os.ExpandEnv("${DB_CONNECTION}://****:****@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable")
}
