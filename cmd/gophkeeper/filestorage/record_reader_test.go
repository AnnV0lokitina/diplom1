package filestorage

import (
	"encoding/json"
	"fmt"
	"github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"reflect"
	"testing"
)

const testReaderFileName = "/test_reader"

func genRecord() (*entity.Record, string) {
	credentialsList := []entity.Credentials{
		{
			Login:    "login",
			Password: "password",
			Meta:     "credentials meta",
		},
		{
			Login:    "login1",
			Password: "password1",
			Meta:     "credentials meta 1",
		},
	}
	textFileList := []entity.File{
		{
			Name: "text_name.ext",
			Meta: "text file meta",
		},
		{
			Name: "text_name1.ext",
			Meta: "text file meta1",
		},
	}
	binaryFileList := []entity.File{
		{
			Name: "bin_name.ext",
			Meta: "bin file meta",
		},
		{
			Name: "bin_name1.ext",
			Meta: "bin file meta1",
		},
	}
	bankCardList := []entity.BankCard{
		{
			Number:     "card number",
			ExpDate:    "card exp_date",
			Cardholder: "cardholder",
			Code:       "card code",
			Meta:       "card meta",
		},
		{
			Number:     "card number1",
			ExpDate:    "card exp_date1",
			Cardholder: "cardholder1",
			Code:       "card code1",
			Meta:       "card meta1",
		},
	}
	obj := entity.Record{
		CredentialsList: credentialsList,
		TextFileList:    textFileList,
		BinaryFileList:  binaryFileList,
		BankCardList:    bankCardList,
	}
	a, _ := json.Marshal(&obj)
	jsonString := string(a)
	return &obj, jsonString
}

func TestNewReader(t *testing.T) {
	type args struct {
		filePath    string
		fileContent string
	}
	type resultInterface interface {
		Empty() bool
		ReadRecord() (*entity.Record, error)
		Close() error
	}
	type want struct {
		resultType      string
		interfaceObject interface{}
		record          *entity.Record
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
			name: "new reader positive",
			args: args{
				filePath:    testDir + testReaderFileName,
				fileContent: jsonRecord,
			},
			want: want{
				resultType:      "*filestorage.Reader",
				interfaceObject: (*resultInterface)(nil),
				record:          objRecord,
			},
			wantCreateErr: assert.NoError,
			wantURLErr:    assert.NoError,
			wantCloseErr:  assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Create(tt.args.filePath)
			require.NoError(t, err)
			_, err = file.Write([]byte(tt.args.fileContent))
			require.NoError(t, err)
			err = file.Close()
			require.NoError(t, err)

			r, err := NewReader(tt.args.filePath)
			if !tt.wantCreateErr(t, err, fmt.Sprintf("NewReader(%v)", tt.args.filePath)) {
				return
			}
			assert.Equalf(t, tt.want.resultType, reflect.TypeOf(r).String(), "NewReader(%v)", tt.args.filePath)
			assert.Implements(t, tt.want.interfaceObject, r, "Invalid reader interface")
			rec, err := r.ReadRecord()
			if !tt.wantURLErr(t, err, "decode record") {
				return
			}
			assert.Equal(t, tt.want.record, rec)
			tt.wantCloseErr(t, r.Close(), "Close()")
			os.Remove(tt.args.filePath)
		})
	}
	os.RemoveAll(testDir)
}
