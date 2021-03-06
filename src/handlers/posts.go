package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"../courseconfig"
	"../user"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
)

//Handle posts for /{matrikel}/post/{postnr}
func HandlePosts(w http.ResponseWriter, r *http.Request) {

	conf := courseconfig.GetConfig()
	us, err := studentFromURL(r)
	if err != nil {
		fmt.Println(MsgErrorReadingMat)
	}

	//Check Permission
	if !checkViewPermission(us, r) {
		WriteMsg(w, MsgPermissionDenied)
		return
	}

	params := mux.Vars(r)
	postNrSt := params["postnr"]
	postNr, _ := strconv.Atoi(postNrSt)

	//Parse Markdown to []byte
	md := blackfriday.MarkdownCommon(us.GetPost(postNr))

	//bool to check if there is Feedback in Template
	FbIs := false

	fb, err := user.GetFeedback(us.Matrikel, postNr)
	if err == nil {
		FbIs = true
	}

	_, loggedInUserMat := loggedIn(r)
	loggedInUser, err := user.FromMatrikel(loggedInUserMat)

	//Check if logged in user is authorized to see Feedback to the current Post
	allowFeedback := false
	if loggedInUserMat == us.Matrikel || loggedInUser.IsAuthorized() && err == nil {
		allowFeedback = true
	}

	pageData := struct {
		DisplayFeedback bool
		St              user.User
		Profile         string
		PostNr          int
		FbIs            bool
		Fb              user.Feedback
		Conf            courseconfig.Config
	}{
		DisplayFeedback: allowFeedback,
		St:              us,
		Profile:         string(md[:]),
		PostNr:          postNr,
		FbIs:            FbIs,
		Fb:              fb,
		Conf:            conf,
	}

	tpl := template.Must(template.ParseFiles("./templates/posts.go.html"))
	tpl.Execute(w, pageData)
}
