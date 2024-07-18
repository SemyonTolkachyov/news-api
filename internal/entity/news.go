package entity

//go:generate reform

// reform:News
type News struct {
	Id      int    `param:"Id" query:"Id" form:"Id" reform:"Id,pk"`
	Title   string `param:"Title" query:"Title" form:"Title" reform:"Title"`
	Content string `param:"Content" query:"Content" form:"Content" reform:"Content"`
}

//go:generate reform

// reform:NewsCategories
type NewsCategory struct {
	NewsId     int `param:"NewsId" query:"NewsId" form:"NewsId" reform:"NewsId"`
	CategoryId int `param:"CategoryId" query:"CategoryId" form:"CategoryId" reform:"CategoryId"`
}
