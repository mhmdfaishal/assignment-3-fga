package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {

	layout := path.Join("web", "layout.html")
	index := path.Join("web", "index.html")
	footer := path.Join("web", "layouts", "_footer.html")
	header := path.Join("web", "layouts", "_header.html")

	var tpl = template.Must(template.New("layout").ParseFiles(footer, header, index, layout))
	
	data, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	var jsonData map[string]interface{}
	json.NewDecoder(data).Decode(&jsonData)

	jsonData["status"].(map[string]interface{})["water"] = rand.Intn(100)
	jsonData["status"].(map[string]interface{})["wind"] = rand.Intn(100)


	file, _ := json.MarshalIndent(jsonData, "", " ")
	_ = os.WriteFile("data.json", file, 0644)


	water := jsonData["status"].(map[string]interface{})["water"].(int)
	wind := jsonData["status"].(map[string]interface{})["wind"].(int)


	if water < 5 || wind < 6 {
		jsonData["status"].(map[string]interface{})["status"] = "safe"
	} else if (water > 6 && water < 8) || (wind > 7 && wind < 15) {
		jsonData["status"].(map[string]interface{})["status"] = "standby"
	} else if water > 8 || wind > 15 {
		jsonData["status"].(map[string]interface{})["status"] = "danger"
	}

	err = tpl.ExecuteTemplate(c.Writer, "layout", jsonData)
	if err != nil {
		log.Print("template executing error: ", err)
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

}
