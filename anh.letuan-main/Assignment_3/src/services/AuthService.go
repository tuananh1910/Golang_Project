package services

import (
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"

	"Assignment_3/src/models/dto"
	"Assignment_3/src/models/entity"
	"Assignment_3/src/repository"
)

type AuthService interface {
	VerifyCredential(email, password string) interface{}
	Register(registerAccount dto.RegisterDTO) entity.Account
	IsDuplicateEmail(email string) bool
}

type authService struct {
	accountRepo repository.AccountRepository
}

func (a *authService) VerifyCredential(email, password string) interface{} {
	res := a.accountRepo.VerifyCredential(email, password)
	if v, ok := res.(entity.Account); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func comparePassword(hashedPassword string, password []byte) bool {
	hashedPasswordByte := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(hashedPasswordByte, password)

	if err != nil {
		return false
	}

	return true
}

func (a *authService) Register(registerAccount dto.RegisterDTO) entity.Account {
	regisAccount := entity.Account{}

	err := smapping.FillStruct(&regisAccount, smapping.MapFields(&registerAccount))

	if err != nil {
		log.Println("failed map")
	}

	res := a.accountRepo.Save(regisAccount)
	return res
}

func (a *authService) IsDuplicateEmail(email string) bool {
	res := a.accountRepo.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

// create instance accountService
func NewAuthService(accountRepo repository.AccountRepository) AuthService {
	return &authService{
		accountRepo: accountRepo,
	}
}
