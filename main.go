package main

import (
	"database/sql"
	"os"

	accountMigrations "lockbox.dev/accounts/storers/postgres/migrations"
	grantMigrations "lockbox.dev/grants/storers/postgres/migrations"
	scopeMigrations "lockbox.dev/scopes/storers/postgres/migrations"
	tokenMigrations "lockbox.dev/tokens/storers/postgres/migrations"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	yall "yall.in"
	"yall.in/colour"
)

func main() {
	log := yall.New(colour.New(os.Stdout, yall.Debug))

	connString := os.Getenv("PG_DB")
	if connString == "" {
		log.Error("PG_DB must be set")
		os.Exit(1)
	}
	pg, err := sql.Open("postgres", connString)
	if err != nil {
		log.WithError(err).Error("error connecting to postgres")
		os.Exit(1)
	}
	accountsMigrations := &migrate.AssetMigrationSource{
		Asset:    accountMigrations.Asset,
		AssetDir: accountMigrations.AssetDir,
		Dir:      "sql",
	}
	_, err = migrate.Exec(pg, "postgres", accountsMigrations, migrate.Up)
	if err != nil {
		log.WithError(err).Error("error running accounts migrations")
		pg.Close()
		os.Exit(1)
	}
	grantsMigrations := &migrate.AssetMigrationSource{
		Asset:    grantMigrations.Asset,
		AssetDir: grantMigrations.AssetDir,
		Dir:      "sql",
	}
	_, err = migrate.Exec(pg, "postgres", grantsMigrations, migrate.Up)
	if err != nil {
		log.WithError(err).Error("error running grants migrations")
		pg.Close()
		os.Exit(1)
	}
	scopesMigrations := &migrate.AssetMigrationSource{
		Asset:    scopeMigrations.Asset,
		AssetDir: scopeMigrations.AssetDir,
		Dir:      "sql",
	}
	_, err = migrate.Exec(pg, "postgres", scopesMigrations, migrate.Up)
	if err != nil {
		log.WithError(err).Error("error runing scopes migrations")
		pg.Close()
		os.Exit(1)
	}
	tokensMigrations := &migrate.AssetMigrationSource{
		Asset:    tokenMigrations.Asset,
		AssetDir: tokenMigrations.AssetDir,
		Dir:      "sql",
	}
	_, err = migrate.Exec(pg, "postgres", tokensMigrations, migrate.Up)
	if err != nil {
		log.WithError(err).Error("error running tokens migrations")
		pg.Close()
		os.Exit(1)
	}
	pg.Close()
	log.Info("migrations complete")
}
