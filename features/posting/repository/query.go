package repository

import (
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/posting/domain"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Delete(ID uint) error {
	var resQry Post
	if err := rq.db.Delete(&resQry, "ID = ?", ID).Error; err != nil {
		return err
	}

	return nil
}

func (rq *repoQuery) Insert(newPost domain.Core) (domain.Core, error) {
	var cnv Post
	cnv = FromDomain(newPost)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	newPost = ToDomain(cnv)
	return newPost, nil
}

func (rq *repoQuery) Update(updatedData domain.Core) (domain.Core, error) {
	var cnv Post
	cnv = FromDomain(updatedData)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	updatedData = ToDomain(cnv)
	return updatedData, nil
}
func (rq *repoQuery) GetAllPost() ([]domain.Core, error) {
	var resQry []Post
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArray(resQry)
	return res, nil
}
