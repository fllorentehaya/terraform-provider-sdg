package main

import (
	"bytes"
	"io/ioutil"
	"os"

	//"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	log "github.com/sirupsen/logrus"
)

func resourceBatchTemplate() *schema.Resource {
	return &schema.Resource{
		// Operations required by every Terraform resource.
		Create: resourceBatchCreate,
		Read:   resourceBatchRead,
		Update: resourceBatchUpdate,
		Delete: resourceBatchDelete,

		// Define the fields of this schema.
		Schema: map[string]*schema.Schema{
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"entorno": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceBatchCreate(d *schema.ResourceData, m interface{}) error {
	//("Creating word")
	//log.Printf("HOLA")
	//log.Level:=log.TraceLevel
	//var log = logrus.New()
	//log.Formatter = new(log.JSONFormatter)
	//log.Formatter = new(log.TextFormatter)                        //default
	//log.Formatter.(*log.TextFormatter).DisableColors = true       // remove colors
	//log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	//log.Level = logrus.TraceLevel
	//log.Out = os.Stdout
	log.SetOutput(os.Stdout)
	log.Error("Creating algo")
	log.Info("Creating word")
	location := d.Get("location").(string)
	tenant := d.Get("tenant").(string)
	client_id := d.Get("client_id").(string)
	client_secret := d.Get("client_secret").(string)
	subscription := d.Get("subscription").(string)
	entorno := d.Get("entorno").(string)
	var token = authenticationToken(tenant, client_id, client_secret)
	createResourceGroup(token, subscription, entorno)
	createDataFactory(token, subscription, entorno)
	log.Printf(token)
	d.SetId(location)

	return nil
}

func resourceBatchRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("HOLA2")
	location := d.Get("location").(string)
	d.SetId(location)
	return nil
}

func resourceBatchUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("HOLA3")
	d.Set("location", "Si")
	//d.Set("author", bookResponse.Author)http.ListenAndServe(":8080", nil)
	return nil
}

func resourceBatchDelete(d *schema.ResourceData, m interface{}) error {

	log.SetOutput(os.Stdout)
	log.Error("Creating algo")
	log.Info("Creating word")
	location := d.Get("location").(string)
	tenant := d.Get("tenant").(string)
	client_id := d.Get("client_id").(string)
	client_secret := d.Get("client_secret").(string)
	subscription := d.Get("subscription").(string)
	entorno := d.Get("entorno").(string)
	var token = authenticationToken(tenant, client_id, client_secret)

	deleteDataFactory(token, subscription, entorno)
	deleteResourceGroup(token, subscription, entorno)

	log.Printf(token)
	d.SetId(location)
	return nil
}

// funciones auxiliares
func deleteResourceGroup(access_token string, subscription string, entorno string) {
	var bearer = "Bearer " + access_token
	var jsonStr = []byte(`{"location":"West Europe"}`)
	url := "https://management.azure.com/subscriptions/" + subscription + "/resourcegroups/sdg-rg-" + entorno + "-001?api-version=2021-04-01"
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	//client := &http.Client{}
	client2 := &http.Client{}
	resp2, err2 := client2.Do(req)
	log.Printf(url)
	log.Printf(resp2.Status)
	if err2 != nil {
		panic(err)
	}

}
func createResourceGroup(access_token string, subscription string, entorno string) {
	var bearer = "Bearer " + access_token
	var jsonStr = []byte(`{"location":"West Europe"}`)
	url := "https://management.azure.com/subscriptions/" + subscription + "/resourcegroups/sdg-rg-" + entorno + "-001?api-version=2021-04-01"
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	//client := &http.Client{}
	client2 := &http.Client{}
	resp2, err2 := client2.Do(req)
	log.Printf(url)
	log.Printf(resp2.Status)
	if err2 != nil {
		panic(err)
	}

}
func createDataFactory(access_token string, subscription string, entorno string) {
	var bearer = "Bearer " + access_token
	var jsonStr = []byte(`{"location":"West Europe"}`)
	url := "https://management.azure.com/subscriptions/" + subscription + "/resourceGroups/sdg-rg-" + entorno + "-001/providers/Microsoft.DataFactory/factories/sdg-df-" + entorno + "-001?api-version=2018-06-01"
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	//client := &http.Client{}
	client2 := &http.Client{}
	resp2, err2 := client2.Do(req)
	log.Printf(url)
	log.Printf(resp2.Status)
	if err2 != nil {
		panic(err)
	}

}
func deleteDataFactory(access_token string, subscription string, entorno string) {
	var bearer = "Bearer " + access_token
	var jsonStr = []byte(`{"location":"West Europe"}`)
	url := "https://management.azure.com/subscriptions/" + subscription + "/resourceGroups/sdg-rg-" + entorno + "-001/providers/Microsoft.DataFactory/factories/sdg-df-" + entorno + "-001?api-version=2018-06-01"
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	//client := &http.Client{}
	client2 := &http.Client{}
	resp2, err2 := client2.Do(req)
	log.Printf(url)
	log.Printf(resp2.Status)
	if err2 != nil {
		panic(err)
	}

}
func authenticationToken(tenant string, client_id string, client_secret string) string {
	var auth_url = "https://login.microsoftonline.com/" + tenant + "/oauth2/token"

	data := url.Values{}
	data.Set("client_id", client_id)
	data.Set("client_secret", client_secret)
	data.Set("grant_type", "client_credentials")
	data.Set("resource", "https://management.core.windows.net/")

	client := &http.Client{}
	r, err := http.NewRequest("POST", auth_url, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	//r.Header.Add("Content-type", "application/json")
	res, err := client.Do(r)

	defer res.Body.Close()
	//decoder := json.NewDecoder(res.Body)
	//var auth Autenticacion
	//decoder.Decode(&auth)
	//log.Printf("www" + auth.expires_in)

	//log.Printf(res.Status)

	body, err := ioutil.ReadAll(res.Body)

	log.Printf(string(body))
	res1 := strings.Index(string(body), "access_token")
	limite := len(string(body))
	limite_total := limite - res1 + 15

	log.Printf(strconv.Itoa(res1))
	log.Printf(strconv.Itoa(limite))
	log.Printf(strconv.Itoa(limite_total))
	log.Printf(string(body)[res1+15 : limite-2])
	access_token := string(body)[res1+15 : limite-2]
	return access_token

}
