package entities

type Forecast struct {
	List []Weather `json:"list"`
}
