package helper

type AjaxReturn struct {
	Success  bool
	Title    string
	Icon     string
	Body     string
	Optional interface{}
}

func AjaxReturnError(body string) AjaxReturn {
	res := AjaxReturn{
		Success:  false,
		Title:    "Gagal",
		Icon:     "error",
		Body:     body,
		Optional: nil,
	}

	return res
}
func AjaxReturnSuccess(body string) AjaxReturn {
	res := AjaxReturn{
		Success:  true,
		Title:    "Berhasil",
		Icon:     "success",
		Body:     body,
		Optional: nil,
	}

	return res
}
