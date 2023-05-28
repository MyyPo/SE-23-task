package simpdb

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSimpDBProvider_CreateOne(t *testing.T) {
	simpDb := NewSimpDBProvider(dbPath)
	type args struct {
		newRecord string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			fmt.Sprintf("create record: %s", testEmail1),
			args{
				newRecord: testEmail1,
			},
			false,
		},
		{
			fmt.Sprintf("create record: %s", testEmail2),
			args{
				newRecord: testEmail2,
			},
			false,
		},
		{
			fmt.Sprintf("create record: %s", testEmail3),
			args{
				newRecord: testEmail3,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := simpDb.CreateOne(tt.args.newRecord); (err != nil) != tt.wantErr {
				t.Errorf("SimpDBProvider.CreateOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimpDBProvider_GetAll(t *testing.T) {
	simpDb := NewSimpDBProvider(dbPath)
	tests := []struct {
		name    string
		want    *[]string
		wantErr bool
	}{
		{
			"get all previously created emails",
			&[]string{testEmail1, testEmail2, testEmail3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := simpDb.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("SimpDBProvider.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpDBProvider.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	dbPath     = "./db/"
	testEmail1 = "example@gmail.com"
	testEmail2 = "etc@gmail.com"
	testEmail3 = "farm@outlook.com"
)

func cleanupDb() {
	fileInfos, _ := os.ReadDir(dbPath)

	for _, fileInfo := range fileInfos {
		filePath := filepath.Join(dbPath, fileInfo.Name())
		err := os.Remove(filePath)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestMain(m *testing.M) {
	exitCode := m.Run()

	cleanupDb()

	os.Exit(exitCode)
}
