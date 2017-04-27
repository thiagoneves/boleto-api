package db

import (
	"net/http"
	"strings"

	"fmt"

	"io/ioutil"

	"encoding/json"

	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/util"
)

type elasticDb struct{}

//SaveBoleto salva um boleto no elasticsearch
func (e *elasticDb) SaveBoleto(id string, boleto models.BoletoView) error {
	client := util.DefaultHTTPClient()
	url := fmt.Sprintf("%s/boletoapi/boleto/%s", config.Get().ElasticURL, id)
	req, err := http.NewRequest("POST", url, strings.NewReader(boleto.ToJSON()))
	if err != nil {
		return err
	}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return errResp
	}
	defer req.Body.Close()
	defer resp.Body.Close()
	return nil
}

//GetBoletoById busca um boleto pelo ID que vem na URL
func (e *elasticDb) GetBoletoByID(id string) (models.BoletoView, error) {
	client := util.DefaultHTTPClient()
	url := fmt.Sprintf("%s/boletoapi/boleto/%s", config.Get().ElasticURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.BoletoView{}, err
	}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return models.BoletoView{}, errResp
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	elasticData := struct {
		Source models.BoletoView `json:"_source"`
	}{}
	json.Unmarshal(data, &elasticData)
	return elasticData.Source, nil
}
