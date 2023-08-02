package types

type OBUData struct {
	OBUID int     `json:"id"`
	Long  float64 `json:"long"`
	Lat   float64 `json:"lat"`
}
