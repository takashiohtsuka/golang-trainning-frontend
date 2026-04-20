package women

type DistrictCountResponse struct {
	Total uint `json:"total"`
}

func NewDistrictCountResponse(total uint) DistrictCountResponse {
	return DistrictCountResponse{Total: total}
}
