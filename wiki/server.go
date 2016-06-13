package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "html/template"
  "encoding/base64"
  "strings"
)

type page struct{
  Title string
  Base64Title string
  Body []byte
}

type pageItem struct {
  Title string
  Href string
}

type mainPage struct{
  Title string
  Body []pageItem
}

const filePrefix string = "wiki_prefix_"

func (p *page) save() error{
  name := base64.StdEncoding.EncodeToString([]byte(p.Title))
  filename := "wiki-docs/"+ filePrefix + name
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadWikiDoc(base64Title string) (*page, error){
  filename := "wiki-docs/"+ filePrefix +base64Title
  body, error := ioutil.ReadFile(filename)
  if(error != nil){
    fmt.Printf("Read file: %v error!, %v\n", filename, error)
    return nil, error
  }

  tmp, _ := base64.StdEncoding.DecodeString(base64Title)
  title := string(tmp)

  return &page{Title: title, Base64Title:base64Title, Body:body}, nil
}

func renderTemplate(w http.ResponseWriter, temp string, p interface{}){
  t, err := template.ParseFiles("public/"+temp+".tpl.html")
  if(err !=nil){
    fmt.Fprintf(w,"%v" , err)
  }else{
    t.Execute(w, p)
  }
}

func createHandler(w http.ResponseWriter, r *http.Request){
  p := &page{Title: "Create new page"}
  renderTemplate(w, "create", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request){
  title := r.URL.Path[len("/view/"):]

  fmt.Println(title)
  p, err := loadWikiDoc(title)
  if(err != nil){
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }

  fmt.Println(p)

  renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request){
  title := r.URL.Path[len("/edit/"):]
  p, err := loadWikiDoc(title)

  if (err !=nil){
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }

  renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request){
  title :=r.FormValue("title")
  body := r.FormValue("body")
  p := &page{Title: title, Body: []byte(body)}
  p.save()
  name := base64.StdEncoding.EncodeToString([]byte(title))
  http.Redirect(w, r, "/edit/"+name, http.StatusFound)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
  path := r.URL.Path[len("/"):]

  if(len(path)>0){
    http.Redirect(w, r, "/", http.StatusFound)
  }else{
    files, err := ioutil.ReadDir("./wiki-docs")
    if(err != nil){

    }

    pageItems := make([]pageItem, 0, 0)

    for _, file := range files{
      fileName:= file.Name()
      if(strings.HasPrefix(fileName, filePrefix)){
        name := string(fileName[len(filePrefix):])
        tmp, _ := base64.StdEncoding.DecodeString(name)
        newTitle := string(tmp)

        pItem := pageItem{Title: newTitle, Href: "/view/"+name}
        pageItems = append(pageItems, pItem)
      }
    }

    renderTemplate(w, "home", pageItems)
  }
}

func main(){
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/create/", createHandler)
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)

  // similar to RESTful API
  http.HandleFunc("/save/", saveHandler)
  // http.HandleFunc("/add/", addHandler)

  http.ListenAndServe(":8080", nil)
}
