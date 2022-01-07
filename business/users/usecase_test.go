package users

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// _middleware "final-project/apps/middlewares"
	"time"
)

type UserUsecaseMocking struct {
	UserUsecase
	repository mock.Mock
}

func (d *UserUsecaseMocking) Register(ctx context.Context, data Domain) (int, error) {
	return 1, nil
}
func (d *UserUsecaseMocking) GetByEmail(ctx context.Context, email string) (Domain, error) {
	return Domain{
		Id:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "Rony",
		Email:     "ronyelkahfi@gmail.com",
		Password:  "112233456",
	}, nil
}

type args struct {
	ctx context.Context
}

func TestRegister(t *testing.T) {
	UserUsecase := UserUsecaseMocking{}
	tests := struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	_, err := UserUsecase.Register(tests.args.ctx, Domain{
		Id:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "Rony",
		Email:     "ronyelkahfi@gmail.com",
		Password:  "112233456",
	})
	assert.Equal(t, nil, err, err)
}

// func TestLogin(t *testing.T) {
// 	UserUsecase := UserUsecaseMocking{
// 		repository:  mock.Mock{},
// 	}
// 	tests := struct {
// 		name string
// 		args args
// 		//want    *http.Response
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	_, err := UserUsecase.Login(tests.args.ctx, "ronyelkahfi@gmail.com", "112233456")
// 	fmt.Println(err)
// 	//assert.Equal(t, nil, err, err)
// 	// _, err1 := UserUsecase.Login(tests.args.ctx, "ronyelkahfi@gmail.com", "passwordsalah")
// 	// assert.NotEqual(t, "Password mismatch!", err1, "Login Must Failed because password missmatch")
// }
