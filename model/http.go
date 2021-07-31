package model

type ListAllTemplateRequest struct {
}

type ListAllTemplateResponse struct {
	List []Template `json:"list"`
}

type SendTemplateToTagRequest struct {
	TemplateId string                 `form:"template_id" json:"template_id"`
	TagId      int                    `form:"tag_id" json:"tag_id"`
	Url        string                 `form:"url" json:"url"`
	Data       map[string]interface{} `form:"data" json:"data"`
}

type SendTemplateToTagResponse struct {
	FailOpenIds []string `json:"fail_open_ids"`
}

type CreateTagRequest struct {
	Tag string `form:"tag" json:"tag"`
}

type CreateTagResponse struct {
	Result bool `json:"result"`
}

type DeleteTagRequest struct {
	TagId int `form:"tag_id" json:"tag_id"`
}

type DeleteTagResponse struct {
	Result bool `json:"result"`
}

type AddTagToUserRequest struct {
	TagId  int    `form:"tag_id" json:"tag_id"`
	OpenId string `form:"open_id" json:"open_id"`
}

type AddTagToUserResponse struct {
	Result bool `json:"result"`
}

type DeleteTagFromUserRequest struct {
	TagId  int    `form:"tag_id" json:"tag_id"`
	OpenId string `form:"open_id" json:"open_id"`
}

type DeleteTagFromUserResponse struct {
	Result bool `json:"result"`
}

type ListAllTagRequest struct {
}

type ListAllTagResponse struct {
	List []Tag `json:"list"`
}

type ListAllUserInfoRequest struct {
}

type ListAllUserInfoResponse struct {
	List []UserInfo `json:"list"`
}
