package controller

type AppController struct {
	Store                   interface{ Store }
	Woman                   interface{ Woman }
	WomanDistrict           interface{ WomanDistrict }
	WomanDistrictCount      interface{ WomanDistrictCount }
	ImmediateAvailableWoman interface{ ImmediateAvailableWoman }
	Prefecture              interface{ Prefecture }
	District                interface{ District }
	BusinessType            interface{ BusinessType }
}
