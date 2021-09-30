package mtss

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/piqba/mtss-go/pkg/errors"
)

// Mtss represents a object for Job from Mtss source
type Mtss struct {
	ID            int     `json:"id"`
	Company       string  `json:"organismo"`
	Position      string  `json:"cargo"`
	Taken         int     `json:"cantidad"`
	Entity        string  `json:"entidad"`
	Province      string  `json:"provincia"`
	Municipality  string  `json:"municipio"`
	Availability  int     `json:"ocupadas"`
	Activity      string  `json:"actividad"`
	Pay           float32 `json:"salario"`
	SchoolLevel   string  `json:"nivelEscolar"`
	Details       string  `json:"observaciones"`
	EntityMail    string  `json:"correo_entidad"`
	EntityAddress string  `json:"direccion_entidad"`
	EntityPhone   string  `json:"telefono_entidad"`
	RegisterDate  string  `json:"fecha_registro"`
	UniqueStamp   string  `json:"unique_stamp"`
	Enabled       bool    `json:"habilitada"`
	Source        string  `json:"source"`
	TypeWork      string  `json:"type_work"`
}

// ToMAP conver this struct to a simple map
func (mt *Mtss) ToMAP() (toHashMap map[string]interface{}, err error) {

	fromStruct, err := json.Marshal(mt)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(fromStruct, &toHashMap); err != nil {
		return toHashMap, err
	}

	return toHashMap, nil
}
func (mt *Mtss) MarshalBinary() ([]byte, error) {
	return json.Marshal(mt)
}

func (mt *Mtss) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &mt); err != nil {
		return err
	}

	return nil
}
func (c *client) MtssJobs(ctx context.Context) ([]Mtss, error) {
	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		OFFERS,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.Errorf("mtss/clientHttp: unexpected response status %d: %q", status, res)
	}
	result := []Mtss{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return nil, errors.Errorf("mtss/clientHttp: decoding error for data %s: %v", res, err)
	}
	return result, nil
}
