package response

type DataResult struct {
	Result     []DetectionResult `json:"result"`
	ImgPath    string            `json:"img_path"`
	ResImgPath string            `json:"res_img_path"`
}

type DetectionResult struct {
	Name    string   `json:"name"`
	Percent string   `json:"percent"`
	Box     [4]int64 `json:"box"`
}
