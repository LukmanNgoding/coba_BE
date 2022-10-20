package services

import (
	"errors"
	"strings"

	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/posting/domain"
	// "github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type postService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &postService{
		qry: repo,
	}
}

func (us *postService) Delete(ID uint) error {
	err := us.qry.Delete(ID)
	if err != nil {
		log.Error(err.Error())
		return errors.New("no data")
	}

	return nil
}
func (us *postService) AddPost(newPost domain.Core) (domain.Core, error) {

	res, err := us.qry.Insert(newPost)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

func (us *postService) UpdatePost(updatedData domain.Core) (domain.Core, error) {
	if updatedData.Owner != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(updatedData.Owner), 10)
		updatedData.Owner = string(generate)
	}

	res, err := us.qry.Update(updatedData)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
		return domain.Core{}, errors.New("rejected from database")
	}

	return res, nil
}

func (us *postService) ShowAllPost() ([]domain.Core, error) {
	res, err := us.qry.GetAllPost()
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("No data")
		}
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}

	return res, nil
}
