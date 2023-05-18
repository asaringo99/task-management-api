package request

type TaskModel struct {
	Userid   int    `json:"id" form:"id" query:"id"`
	Status   string `json:"status" form:"status" query:"status"`
	Priority int    `json:"priority" form:"priority" query:"priority"`
	Contents string `json:"contents" form:"contents" query:"contents"`
}
