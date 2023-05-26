package main

import (
	"context"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/YutaIke/go-api-experiment/config"
	"github.com/YutaIke/go-api-experiment/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := config.NewConfig(".env.local") // TODO: set file name for each environment
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                          // provide migration directory
		schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
		schema.WithDialect(dialect.MySQL),            // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}
	requiredArgsNum := 2
	if len(os.Args) != requiredArgsNum {
		log.Fatal("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}
	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, config.ConnectionString, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
