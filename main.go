package main

import (
	"database/sql"
	"os"

	accountMigrations "lockbox.dev/accounts/storers/postgres/migrations"
	clientMigrations "lockbox.dev/clients/storers/postgres/migrations"
	grantsPostgres "lockbox.dev/grants/storers/postgres"
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

	m := map[string]migrate.MigrationSource{
		"accounts": &migrate.AssetMigrationSource{
			Asset:    accountMigrations.Asset,
			AssetDir: accountMigrations.AssetDir,
			Dir:      "sql",
		},
		"clients": &migrate.AssetMigrationSource{
			Asset:    clientMigrations.Asset,
			AssetDir: clientMigrations.AssetDir,
			Dir:      "sql",
		},
		"grants": grantsPostgres.MigrationSource(),
		"scopes": &migrate.AssetMigrationSource{
			Asset:    scopeMigrations.Asset,
			AssetDir: scopeMigrations.AssetDir,
			Dir:      "sql",
		},
		"tokens": &migrate.AssetMigrationSource{
			Asset:    tokenMigrations.Asset,
			AssetDir: tokenMigrations.AssetDir,
			Dir:      "sql",
		},
	}
	for pkg, source := range m {
		log.WithField("pkg", pkg).Debug("running migrations")
		migrate.SetTable("migrations_" + pkg)
		_, err = migrate.Exec(pg, "postgres", source, migrate.Up)
		if err != nil {
			log.WithError(err).WithField("pkg", pkg).Error("error running migrations")
			pg.Close()
			os.Exit(1)
		}
	}
	pg.Close()
	log.Info("migrations complete")
}
