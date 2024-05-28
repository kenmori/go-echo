package usecase

import (
	"go-echo/model"
	"go-echo/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

/*
> こちらはなぜ値渡しにしているのですか？

この場合、model.User構造体を値渡しにしている理由は以下の2点が考えられます。

不変性の確保

値渡しを行うことで、呼び出し元のmodel.Userインスタンスに影響を与えずに、メソッド内で安全に操作できます。ポインタ渡しの場合、メソッド内で受け取ったポインタを介して元のインスタンスの値が変更される可能性があります。値渡しによってこの問題を回避できます。

コピーコストの軽減

model.User構造体のフィールドが少数の場合、値コピーのコストは大きくありません。ポインタを使う場合、ポインタ自体はコピーされますが、実際のデータはヒープ上に存在するため、ポインタの値渡しではメモリコピーは発生しません。しかし、ポインタを介してフィールドにアクセスする際の間接参照コストが発生します。
従って、model.User構造体のフィールド数が少ない場合、値渡しの方がパフォーマンス上有利になる可能性があります。
一方、model.User構造体のサイズが大きい場合は、値コピーのコストが高くなるため、ポインタ渡しの方が適切です。
最終的には、構造体のサイズ、不変性の要求、パフォーマンス要件などを考慮して、値渡しかポインタ渡しかを決める必要があります。ただし、一般的には小さな構造体の場合は値渡しを、大きな構造体の場合はポインタ渡しを選択することが多いです。
*/

type userUsecase struct {
	ur repository.IUserRepository
}

// DIするためのconstructor
func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur} // 実体のpointerを&で取得して返している
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	err = uu.ur.CreateUser(&newUser)
	if err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	storeUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storeUser, user.Email); err != nil {
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(storeUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storeUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
