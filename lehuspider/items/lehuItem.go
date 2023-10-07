package items

type LehuIndexItem struct {
	Avatar  string            `json:"avatar"`
	Author  string            `json:"author"`
	Title   string            `json:"title,omitempty"`
	Content string            `json:"content"`
	Hot     string            `json:"hot"`
	CmtNum  string            `json:"cmtNum"`
	Date    string            `json:"date"`
	ImgURL  string            `json:"imgURL"`
	Tags    map[string]string `json:"Tags"`
	Cmt     []string          `json:"cmt"`
}
