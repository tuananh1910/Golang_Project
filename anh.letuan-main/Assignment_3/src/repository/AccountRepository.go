package repository

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"Assignment_3/src/models/entity"
)

type AccountRepository interface {
	Save(account entity.Account) entity.Account
	VerifyCredential(email, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)

	Profile(idAccount string) entity.Account
}

type accountConnectDB struct {
	connect *gorm.DB
}

func hashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Fail hash password")
	}
	return string(hash)
}

func (a *accountConnectDB) Save(account entity.Account) entity.Account {
	account.Password = hashPassword([]byte(account.Password))

	a.connect.Save(&account)

	return account
}

func (a *accountConnectDB) VerifyCredential(email, password string) interface{} {
	var account entity.Account
	res := a.connect.Where("email = ?", email).Take(&account)

	if res.Error == nil {
		return account
	}

	return nil
}

func (a *accountConnectDB) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var account entity.Account
	return a.connect.Where("email = ?", email).Take(&account)
}

func (a *accountConnectDB) Profile(idAccount string) entity.Account {
	var account entity.Account
	a.connect.Find(&account, idAccount)
	return account
}

// create instance of account repo
func NewAccountRepo(db *gorm.DB) AccountRepository {
	return &accountConnectDB{
		connect: db,
	}
}
