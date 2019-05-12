module impractical.co/auth/cmd/migrations-runner

replace (
	impractical.co/auth/accounts v0.0.0 => ../../accounts
	impractical.co/auth/grants v0.0.0 => ../../grants
	impractical.co/auth/scopes v0.0.0 => ../../scopes
	impractical.co/auth/tokens v0.0.0 => ../../tokens
)

require (
	github.com/fatih/color v1.7.0 // indirect
	github.com/lib/pq v1.0.0
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/rubenv/sql-migrate v0.0.0-20180704111356-3f452fc0ebeb
	gopkg.in/gorp.v1 v1.7.1 // indirect
	impractical.co/auth/accounts v0.0.0
	impractical.co/auth/grants v0.0.0
	impractical.co/auth/scopes v0.0.0
	impractical.co/auth/tokens v0.0.0
	yall.in v0.0.1
)
