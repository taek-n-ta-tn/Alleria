package demo

import (
	"encoding/json"
	"fmt"
	"github.com/caibirdme/yql"
	"github.com/gogf/gf/net/ghttp"
	"github.com/taek-n-ta-tn/Alleria/lib/response"
)

func (c *Controller) Index(r *ghttp.Request) {
	var data *SignInRequest
	if err := r.Parse(&data); err != nil {
		response.JsonERR(r, err.Error())
	}
	if one, err := queryArticle(data); err != nil {
		response.JsonERR(r, err.Error())
	} else {
		response.JsonOK(r, one)
	}
}

func (c *Controller) Rule(r *ghttp.Request) {
	rawYQL := `name='deen' and age>=23 and (hobby in ('soccer', 'swim') or score>90))`
	ruler, _ := yql.Rule(rawYQL)
	result, _ := ruler.Match(map[string]interface{}{
		"name":  "deen",
		"age":   int64(23),
		"hobby": "soccer",
		"score": int64(60),
	})
	fmt.Println(result)
	rawYQL = `score ∩ (7,1,9,5,3)`
	result, _ = ruler.Match(map[string]interface{}{
		"score": []int64{3, 100, 200},
	})
	fmt.Println(result)
	rawYQL = `score in (7,1,9,5,3)`
	result, _ = ruler.Match(map[string]interface{}{
		"score": []int64{3, 5, 2},
	})
	fmt.Println(result)
	rawYQL = `score.sum() > 10`
	result, _ = ruler.Match(map[string]interface{}{
		"score": []int{1, 2, 3, 4, 5},
	})
	fmt.Println(result)
	//Output:
	//true
	//true
	//false
	//true
}

func (c *Controller) DoRule(r *ghttp.Request) {
	rawYQL := r.GetString("entry")
	if rawYQL == "" {
		response.JsonERR(r, "entry must be exist")
	}
	matherStr := r.GetString("matcher")
	if matherStr == "" {
		response.JsonERR(r, "matcher must be exist")
	}

	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(matherStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
		response.JsonERR(r, "JsonToMapDemo err: ")
	}

	ruler, _ := yql.Rule(rawYQL)
	result, _ := ruler.Match(mapResult)
	fmt.Println(result)
}
