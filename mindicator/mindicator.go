package mindicator

import(
	"time"
	"net/http"
	"io/ioutil"
	"log"
    "encoding/json"
    "fmt"
    "yagan.bot/telegram"
)

type Response struct {
    Version string `json:"version"`
    Autor string `json:"autor"`
    Fecha string `json:"fecha"`
    UF Indicator `json:"uf"`
    USD Indicator `json:"dolar"`
    EUR Indicator `json:"euro"`
}

type Indicator struct {
    Codigo string `json:"codigo"`
    Nombre string `json:"nombre"`
    Unidad string `json:"unidad_medida"`
    Fecha string `json:"fecha"`
    Valor float64 `json:"valor"`
}


func PrintIndicators() {
	time.Sleep(10 * time.Second)

	for {
        client := &http.Client{}
        req, err := http.NewRequest("GET", "https://mindicador.cl/api", nil)

        req.Header.Add("Accept", "application/json")
        req.Header.Add("Content-Type", "application/json")

        resp, err := client.Do(req)
        if err != nil {
           log.Fatal(err.Error())
        }

        defer resp.Body.Close()
        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err.Error())
        }

        var responseObject Response
        json.Unmarshal(bodyBytes, &responseObject)
        //log.Printf("API Response as struct %+v\n", responseObject)
        //log.Println(createMessage(&responseObject))
		telegram.SendNotification(createMessage(&responseObject))

		time.Sleep(1 * time.Hour)
	}
}

func createMessage(response *Response) string {
    return "UF: "+fmt.Sprintf("%.2f",response.UF.Valor)+" | USD: "+fmt.Sprintf("%.2f",response.USD.Valor)+" | EUR: "+fmt.Sprintf("%.2f",response.EUR.Valor)
}
