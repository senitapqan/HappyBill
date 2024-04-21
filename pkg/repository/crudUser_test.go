package repository

// import (
// 	"happyBill/models"
// 	"log"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	sqlmock "github.com/zhashkevych/go-sqlxmock"
// )

// func CreateUser_Test(t *testing.T) {
// 	db, mock, err := sqlmock.Newx()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	r := NewRepository(db)

// 	tests := []struct {
// 		name 		string
// 		mock 		func()
// 		input		models.User
// 		want 		int
// 		wantErr 	bool
// 	} {
// 		{
// 			name: "Ok",
// 			mock: func() {
// 				rows:= sqlmock.NewRows([]string{"id"}).AddRow(1)
// 				mock.ExpectQuery("INSERT INTO t_users").WithArgs("Test username", "Test password", "Test name", "Test surname", "testemail@gmail.com").
// 				WillReturnRows(rows)
// 			},
// 			input: models.User{
// 				Username: "Test username",
// 				Password: "Test password",
// 				Name: "Test name",
// 				Surname: "Test surname",
// 				Email: "testemail@gmail.com",
// 			},
// 			want: 1,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mock()

// 			got, err := r.CreateUser(tt.input, )
// 			if tt.wantErr {
// 				assert.Error(t, err)

// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.want, got)
// 			}
// 			assert.NoError(t, mock.ExpectationsWereMet())
// 		})
// 	}

// }
