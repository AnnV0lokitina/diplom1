package db

import (
	"context"
	"flag"
	"github.com/AnnV0lokitina/diplom1/migrations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewDB(t *testing.T) {
	args := flag.Args()
	if len(args) == 0 {
		t.Skip("Skipping testing in CI environment")
	}
	dsn := args[0]
	err := migrations.DoMigrates(dsn)
	require.NoError(t, err)
	ctx := context.Background()

	db, err := NewDB(ctx, dsn)
	assert.Nil(t, err)

	// clear db
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	sqlClearSessions := `DELETE FROM sessions`
	_, err = db.conn.Exec(ctx, sqlClearSessions)
	require.NoError(t, err)

	sqlClearUsers := `DELETE FROM users`
	_, err = db.conn.Exec(ctx, sqlClearUsers)
	require.NoError(t, err)

	err = db.Close(ctx)
	assert.Nil(t, err)
}

func TestUser(t *testing.T) {
	args := flag.Args()
	if len(args) == 0 {
		t.Skip("Skipping testing in CI environment")
	}
	dsn := args[0]
	err := migrations.DoMigrates(dsn)
	require.NoError(t, err)
	ctx := context.Background()

	db, err := NewDB(ctx, dsn)
	assert.Nil(t, err)

	err = db.CreateUser(ctx, "session", "login", "password")
	assert.Nil(t, err)
	err = db.CreateUser(ctx, "session", "login", "password")
	assert.Error(t, err)

	_, err = db.GetUserBySessionID(ctx, "session1")
	assert.Error(t, err)
	user, err := db.GetUserBySessionID(ctx, "session")
	assert.Nil(t, err)

	err = db.AddUserSession(ctx, user)
	assert.Nil(t, err)
	user.ActiveSessionID = "session2"
	err = db.AddUserSession(ctx, user)
	assert.Nil(t, err)

	_, err = db.AuthUser(ctx, "login", "password")
	assert.Nil(t, err)

	_, err = db.AuthUser(ctx, "login", "password1")
	assert.Error(t, err)

	err = db.Close(ctx)
	assert.Nil(t, err)
}
