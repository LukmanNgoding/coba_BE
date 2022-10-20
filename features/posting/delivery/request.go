package delivery

import "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/posting/domain"

type AddFormat struct {
	ID       uint   `json:"id" form:"id"`
	Content  string `json:"content" form:"content"`
	Photo    string `json:"photo" form:"photo"`
	Username string `json:"username" form:"username"`
}

type UpdateFormat struct {
	ID       uint   `json:"id" form:"id"`
	Content  string `json:"content" form:"content"`
	Photo    string `json:"photo" form:"photo"`
	Username string `json:"username" form:"username"`
}

type ShowAllFormat struct {
	ID       uint   `json:"id" form:"id"`
	Content  string `json:"content" form:"content"`
	Owner    string `json:"owner" form:"owner"`
	Comment  string `json:"comment" form:"comment"`
	Username string `json:"username" form:"username"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case AddFormat:
		cnv := i.(AddFormat)
		return domain.Core{Content: cnv.Content, Photo: cnv.Photo, Username: cnv.Username}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, Content: cnv.Content, Photo: cnv.Photo, Username: cnv.Username}
	case ShowAllFormat:
		cnv := i.(ShowAllFormat)
		return domain.Core{ID: cnv.ID, Content: cnv.Content, Owner: cnv.Owner, Comment: cnv.Comment, Username: cnv.Username}

	}

	return domain.Core{}
}
