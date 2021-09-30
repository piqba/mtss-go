package mtss_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/piqba/mtss-go"
	"github.com/stretchr/testify/assert"
)

func Test_Invalid_URL(t *testing.T) {

	api := mtss.NewAPIClient(
		"ht&@-tp://:aa",
		true,
		nil,
		nil,
	)
	actual, err := api.MtssJobs(context.Background())
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func Test_Metrics(t *testing.T) {

	expectedElement := mtss.Mtss{
		ID:           31215,
		Company:      "GP Matanzas",
		Position:     "LIMPIEZA DE CALLE",
		RegisterDate: "2021-09-16T00:00:00",
	}

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					{"id":31215,"organismo_id":31400,"organismo":"GP Matanzas","cargo_id":30179,"cargo":"LIMPIEZA DE CALLE","cantidad":3,"entidad_id":"31406538","entidad":"UP Municipal Direccion De Servicios Comunales De Col","provincia_id":6,"provincia":"Matanzas","municipio_id":52,"municipio":"Colón","ocupadas":0,"actividad":"limpieza de Calles picket","salario":2446.00,"nivel_escolar_id":3,"nivelEscolar":"9no Grado","observaciones":"","correo_entidad":"","direccion_entidad":"Calle Colon e/c Rafael Aguila y Luz caballero #230-B","telefono_entidad":"45-316236","fecha_registro":"2021-09-16T00:00:00","unique_stamp":"20210916112319","habilitada":true},
					{"id":31220,"organismo_id":16100,"organismo":"Ministerio de Comunicaciones","cargo_id":10029,"cargo":"Gestor “A” Comercial Postal","cantidad":1,"entidad_id":"16141449","entidad":"Correos","provincia_id":6,"provincia":"Matanzas","municipio_id":52,"municipio":"Colón","ocupadas":0,"actividad":"gestorA Com Postal","salario":2660.00,"nivel_escolar_id":4,"nivelEscolar":"12mo Grado","observaciones":"","correo_entidad":null,"direccion_entidad":"Camilo Cienfuegos # 34","telefono_entidad":"45313133","fecha_registro":"2021-09-16T00:00:00","unique_stamp":"20210916131852","habilitada":true}
					]`,
			))
		}),
	)
	defer s.Close()
	api := mtss.NewAPIClient(
		s.URL,
		false,
		nil,
		nil,
	)
	jobs, err := api.MtssJobs(context.TODO())
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement.ID, jobs[0].ID)
	assert.Equal(t, expectedElement.Company, jobs[0].Company)
	assert.Equal(t, expectedElement.Position, jobs[0].Position)
	assert.Equal(t, expectedElement.RegisterDate, jobs[0].RegisterDate)
}
