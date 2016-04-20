package controllers

import(
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/httplib"
//	"github.com/astaxie/beego/context"
	models "cannabot/cannabot-lib/models"
)

type WebOperator struct {
	beego.Controller
}

//type Foo struct {
//	Bar string	`json:"bar"`
//}

func (wo *WebOperator) Get() {
	wo.Data["Website"] = "gmail.com"
	wo.Data["Email"] = "Matthew"
	wo.TplName = "index.tpl"
	wo.Post()
	fmt.Println("Get")
}

func (wo *WebOperator) RequestMachine() {
	fmt.Println("Request Machine")
}

func (wo *WebOperator) Post() {
	fmt.Println("Post")
	//foo1 := new(Foo) // or &Foo{}
	//getJson("http://example.com/?bar=asdfasdf", foo1)
	//println(foo1.Bar)

	wo.HandleNewMachine()
	json := wo.GetString("json")
	if json == "" {
		wo.Ctx.WriteString("json is empty")
	}
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// THIS METHOD SHOULD WORK BUT NEED TO SEND IT SOMETHING!!
func (wo *WebOperator) HandleNewMachine() {
	req := wo.Ctx.Request     //in beego this.Ctx.Request points to the HttpRequest
	p := make([]byte, req.ContentLength)
	_, err := wo.Ctx.Request.Body.Read(p)
	if err == nil {
		var newMachine models.Machine
		err1 := json.Unmarshal(p, &newMachine)
		if err1 == nil {
			fmt.Println(newMachine)
		} else {
			fmt.Println("Unable to unmarshall the JSON request", err1);
		}
	}
}

func (mo *WebOperator) SendInstructions() {
	//req := httplib.Get("localhost:8000/test")
	//str, err := res.String()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(str)


	//ctx := context.Context{Request: req, ResponseWriter: rw}
	//out := context.NewOutput()
	//
	//	out.JSON("{asdf: asdfasdf}", true, false)
	//	report := models.InitTestReport()
	//	out.JSON(report, true, false)
	//
	//	mo.Data["json"] = &json.Marshal(report)
	//	mo.ServeJSON()
	//fmt.Println("SendReport()")
}