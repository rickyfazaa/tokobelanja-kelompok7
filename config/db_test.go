package config

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func TestInitDB(t *testing.T) {
	type args struct {
		username string
		password string
		host     string
		port     string
		dbName   string
	}
	tests := []struct {
		name string
		args args
		want *gorm.DB
	}{
		{
			name: "init db",
			args: args{
				username: "root",
				password: "",
				host:     "127.0.0.1",
				port:     "3306",
				dbName:   "hacktiv8_go_fp4",
			},
		},
	}
	fmt.Println(len(tests)) /* Check apakah test table driven berhasil*/
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitDB(tt.args.username, tt.args.password, tt.args.host, tt.args.port, tt.args.dbName); got == nil {
				t.Errorf("InitDB() = %v, want %v", got, "not nil")
			}
		})
	}
}
