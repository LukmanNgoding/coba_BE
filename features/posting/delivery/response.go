package delivery

import "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/posting/domain"

func SuccessDelete(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}
func SuccessLogin(msg string, token string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
		"token":   token,
	}
}

type AddResponse struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	Photo    string `json:"photo"`
	Username string `json:"username"`
}

type UpdateResponse struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	Photo    string `json:"photo"`
	Username string `json:"username"`
}

type ShowAllResponse struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	Owner    string `json:"owner"`
	Comment  string `json:"comment"`
	Username string `json:"username"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "add":
		cnv := core.(domain.Core)
		res = AddResponse{ID: cnv.ID, Content: cnv.Content, Photo: cnv.Photo, Username: cnv.Username}
	case "upd":
		cnv := core.(domain.Core)
		res = UpdateResponse{ID: cnv.ID, Content: cnv.Content, Photo: cnv.Photo, Username: cnv.Username}
	case "all":
		var arr []ShowAllResponse
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, ShowAllResponse{ID: val.ID, Content: val.Content, Owner: val.Owner, Comment: val.Comment, Username: val.Username})
		}
		res = arr
	}

	return res
}
