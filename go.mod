module lockbox.dev/cmd/migrations-runner

// TODO: use the real version of these
replace (
	lockbox.dev/accounts v0.0.0 => ../../accounts
	lockbox.dev/grants v0.0.0 => ../../grants
	lockbox.dev/hmac v0.0.0 => ../../hmac
	lockbox.dev/scopes v0.0.0 => ../../scopes
	lockbox.dev/sessions v0.0.0 => ../../sessions
	lockbox.dev/tokens v0.0.0 => ../../tokens
)

require (
	github.com/lib/pq v1.0.0
	github.com/rubenv/sql-migrate v0.0.0-20180704111356-3f452fc0ebeb
	lockbox.dev/accounts v0.0.0
	lockbox.dev/grants v0.0.0
	lockbox.dev/scopes v0.0.0
	lockbox.dev/tokens v0.0.0
	yall.in v0.0.1
)
