package external

import (
	"bytes"
	"context"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	grpcTest "github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/grpc_test"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestNewExtConnection(t *testing.T) {
	c := NewExtConnection("target", "filepath")
	assert.Equal(t, "target", c.target)
	assert.Equal(t, "filepath", c.filePath)
}

func TestCreateConnection(t *testing.T) {
	c := NewExtConnection("localhost:8888", "filepath")
	//_, err := c.createConnection()
	//assert.Error(t, err)

	ts := grpcTest.NewServer("localhost:8080")
	ctx := context.Background()
	ts.Start(ctx)

	conn, err := c.createConnection()
	assert.Nil(t, err)
	conn.Close()

	ts.Stop()
}

func TestRegister(t *testing.T) {
	c := NewExtConnection("localhost:8080", "filepath")
	ts := grpcTest.NewServer("localhost:8080")
	ctx := context.Background()
	ts.Start(ctx)

	session, err := c.Register(ctx, grpcTest.NewUser, "password")
	assert.Nil(t, err)
	assert.Equal(t, grpcTest.CorrectSession, session)

	_, err = c.Register(ctx, grpcTest.ExistingUser, "password")
	assert.Error(t, err)

	_, err = c.Register(ctx, "unknown user", "password")
	assert.Error(t, err)

	ts.Stop()
}

func TestLogin(t *testing.T) {
	c := NewExtConnection("localhost:8080", "filepath")
	ts := grpcTest.NewServer("localhost:8080")
	ctx := context.Background()
	ts.Start(ctx)

	session, err := c.Login(ctx, grpcTest.ExistingUser, "password")
	assert.Nil(t, err)
	assert.Equal(t, grpcTest.CorrectSession, session)

	_, err = c.Register(ctx, "unknown user", "password")
	assert.Error(t, err)

	ts.Stop()
}

func TestStoreInfo(t *testing.T) {
	c := NewExtConnection("localhost:8080", "filepath")
	ts := grpcTest.NewServer("localhost:8080")
	ctx := context.Background()

	testFileContent := "sookeknfgoerkjgoqi34ujtq9834983q4fi4o3jgi34jgioj34goij34oigjiogjieorgj4830450934i0ft349jf" +
		"oaiehftiauwerhfntiuawrhgiuerhgiuprhgpergherogjherogjeorigjeroig"
	file := strings.NewReader(testFileContent)
	now := time.Now()
	yesterday := now.Add(time.Hour * -24)
	info := &entity.FileInfo{
		UpdateTime: now,
	}
	grpcTest.TestFileDate = yesterday
	err := c.StoreInfo(ctx, grpcTest.CorrectSession, file, info)
	assert.Error(t, err)

	ts.Start(ctx)

	err = c.StoreInfo(ctx, grpcTest.CorrectSession, file, info)
	assert.Nil(t, err)
	assert.Equal(t, testFileContent, grpcTest.TestFileContent)

	file1 := strings.NewReader(testFileContent)
	err = c.StoreInfo(ctx, grpcTest.WithErrorSession, file1, info)
	assert.Error(t, err)
	assert.Equal(t, "External error: Save to storage failed", err.Error())

	file2 := strings.NewReader(testFileContent)
	err = c.StoreInfo(ctx, "incorrect session", file2, info)
	assert.Error(t, err)
	assert.Equal(t, "External error: Save to storage failed", err.Error())

	file3 := strings.NewReader(testFileContent)
	twoDaysAgo := now.Add(time.Hour * -48)
	info1 := &entity.FileInfo{
		UpdateTime: twoDaysAgo,
	}
	err = c.StoreInfo(ctx, grpcTest.CorrectSession, file3, info1)
	assert.Error(t, err)

	ts.Stop()
}

func TestRestoreInfo(t *testing.T) {
	c := NewExtConnection("localhost:8080", "filepath")
	ts := grpcTest.NewServer("localhost:8080")
	ctx := context.Background()

	grpcTest.TestFileContent = "sookeknfgoerkjgoqi34ujtq9834983q4fi4o3jgi34jgioj34goij34oigjiogjieorgj4830450934i0ft349jf" +
		"oaiehftiauwerhfntiuawrhgiuerhgiuprhgpergherogjherogjeorigjeroig"

	now := time.Now()
	yesterday := now.Add(time.Hour * -24)
	info := &entity.FileInfo{
		UpdateTime: yesterday,
	}
	grpcTest.TestFileDate = now
	file := bytes.Buffer{}

	err := c.RestoreInfo(ctx, grpcTest.CorrectSession, &file, info)
	assert.Error(t, err)

	ts.Start(ctx)

	err = c.RestoreInfo(ctx, grpcTest.CorrectSession, &file, info)
	assert.Nil(t, err)
	assert.Equal(t, grpcTest.TestFileContent, file.String())

	file.Reset()
	err = c.RestoreInfo(ctx, grpcTest.WithErrorSession, &file, info)
	assert.Error(t, err)

	file.Reset()
	err = c.RestoreInfo(ctx, "incorrect session", &file, info)
	assert.Error(t, err)

	//file.Reset()
	//info = &entity.FileInfo{
	//	UpdateTime: now,
	//}
	//grpcTest.TestFileDate = yesterday
	//err = c.RestoreInfo(ctx, grpcTest.CorrectSession, &file, info)
	//assert.Error(t, err)

	ts.Stop()
}
