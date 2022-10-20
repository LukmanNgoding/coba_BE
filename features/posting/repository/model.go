package repository

import (
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/posting/domain"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content string
	Owner   string
	Comment string
}

func FromDomain(du domain.Core) Post {
	return Post{
		Model:   gorm.Model{ID: du.ID},
		Content: du.Content,
		Owner:   du.Owner,
		Comment: du.Comment,
	}
}

func ToDomain(u Post) domain.Core {
	return domain.Core{

		ID:      u.ID,
		Content: u.Content,
		Owner:   u.Owner,
		Comment: u.Comment,
	}
}

func ToDomainArray(au []Post) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Content: val.Content, Owner: val.Owner, Comment: val.Comment})
	}

	return res
}
