package entity

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"reflect"
	"testing"
)

const testWriterFileName = "/test_writer"

func TestNewWriter(t *testing.T) {
	type args struct {
		filePath string
		record   *Record
	}
	type resultInterface interface {
		WriteRecord(record *Record) error
		Close() error
	}
	type want struct {
		resultType      string
		interfaceObject interface{}
		record          string
	}
	tmpDir := os.TempDir()
	testDir, err := os.MkdirTemp(tmpDir, "test")
	require.NoError(t, err)

	objRecord, jsonRecord := genRecord()

	tests := []struct {
		name          string
		args          args
		want          want
		wantCreateErr assert.ErrorAssertionFunc
		wantURLErr    assert.ErrorAssertionFunc
		wantCloseErr  assert.ErrorAssertionFunc
	}{
		{
			name: "create writer",
			args: args{
				filePath: testDir + testWriterFileName,
				record:   objRecord,
			},
			want: want{
				resultType:      "*entity.Writer",
				interfaceObject: (*resultInterface)(nil),
				record:          jsonRecord + "\n",
			},
			wantCreateErr: assert.NoError,
			wantURLErr:    assert.NoError,
			wantCloseErr:  assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := NewWriter(tt.args.filePath)
			if !tt.wantCreateErr(t, err, fmt.Sprintf("NewWriter(%v)", tt.args.filePath)) {
				return
			}
			assert.Equalf(t, tt.want.resultType, reflect.TypeOf(w).String(), "NewWriter(%v)", tt.args.filePath)
			assert.Implements(t, tt.want.interfaceObject, w, "Invalid writer interface")
			assert.FileExistsf(t, tt.args.filePath, "file path %v", tt.args.filePath)
			tt.wantURLErr(t, w.WriteRecord(tt.args.record), fmt.Sprintf("WriteRecord(%v)", tt.args.record))
			tt.wantCloseErr(t, w.Close(), "Close()")

			file, err := os.Open(tt.args.filePath)
			require.NoError(t, err)
			buf := make([]byte, 16)
			data := bytes.NewBuffer([]byte{})
			for {
				n, err := file.Read(buf)
				if err == io.EOF {
					break
				}
				data.Write(buf[:n])
			}
			assert.Equal(t, tt.want.record, data.String())
			err = file.Close()
			require.NoError(t, err)
			os.Remove(tt.args.filePath)
		})
	}
	os.RemoveAll(testDir)
}
