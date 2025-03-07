package memberships

import (
	"testing"

	"github.com/Bagussurya12/catalog-music-simple/source/models/memberships"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_repository_CreateUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	type args struct {
		model memberships.User
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				model: memberships.User{
					Email:     "bagus@mailtest.com",
					Username:  "bagus",
					Password:  "password",
					CreatedBy: "bagus@mailtest.com",
					UpdatedBy: "bagus@mailtest.com",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.*)`).WithArgs(
					args.model.Email,
					args.model.Username,
					args.model.Password,
					args.model.CreatedBy,
					args.model.UpdatedBy,
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				model: memberships.User{
					Email:     "bagus@mailtest.com",
					Username:  "bagus",
					Password:  "password",
					CreatedBy: "bagus@mailtest.com",
					UpdatedBy: "bagus@mailtest.com",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.*)`).WithArgs(
					args.model.Email,
					args.model.Username,
					args.model.Password,
					args.model.CreatedBy,
					args.model.UpdatedBy,
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).
					WillReturnError(assert.AnError)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDb,
			}
			if err := r.CreateUser(tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
