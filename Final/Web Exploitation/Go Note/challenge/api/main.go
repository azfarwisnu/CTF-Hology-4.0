package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var (
	dbsqlite *sql.DB
	key      = securecookie.GenerateRandomKey(32)
	store    = sessions.NewCookieStore(key)
	SesName  = "cookie"
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, DisableColors: true})
	log.Info("Init sqlite & logger")
	dbsqlite = func() *sql.DB {
		a, err := sql.Open("sqlite3", "./hology")
		if err != nil {
			log.Panic(err)
		}
		return a
	}()
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Humanz!\n"))
}

func main() {
	r := mux.NewRouter()
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: false,
		Secure:   false,
		Domain:   os.Getenv("HOST"),
	}

	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/Register", func(rw http.ResponseWriter, r *http.Request) {
		log.Info("New user register")
		var (
			UserData User
		)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(body, &UserData)
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
			return
		}

		Users, err := UserData.BcryptPass().FixUserName().Add()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusForbidden)
			return
		}
		Note := Note{
			Data:    lipsum,
			UsersID: Users.ID,
			Encrypt: false,
		}

		err = Note.AddNote()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusForbidden)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(Users)
		rw.WriteHeader(http.StatusOK)

	}).Methods("POST")

	r.HandleFunc("/Login", func(rw http.ResponseWriter, r *http.Request) {
		log.Info("User login")
		var (
			UserData User
		)
		session, err := store.Get(r, SesName)
		if err != nil {
			log.Error(err)
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusInternalServerError)
		}

		err = json.Unmarshal(body, &UserData)
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusInternalServerError)
		}

		if len(session.Values) == 0 {
			UserData, err := UserData.BcryptPass().FixUserName().Login()
			if UserData != nil && err == nil {
				session.Values["UserName"] = UserData.UserName
				session.Values["UserID"] = UserData.ID
				err := session.Save(r, rw)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				}

				rw.Header().Set("Content-Type", "application/json")
				json.NewEncoder(rw).Encode(map[string]interface{}{
					"Status": 200,
					"UserID": UserData.ID,
				})
				rw.WriteHeader(http.StatusOK)
				return
			} else if err != nil {
				http.Error(rw, err.Error(), http.StatusForbidden)
				return
			} else {
				http.Error(rw, "User not found", http.StatusForbidden)
				return
			}
		} else {
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(map[string]interface{}{
				"Status": 200,
				"UserID": session.Values["UserID"],
			})
			rw.WriteHeader(http.StatusOK)

		}

	}).Methods("POST")

	r.HandleFunc("/Logout", func(rw http.ResponseWriter, r *http.Request) {
		log.Info("User Logout")
		session, err := store.Get(r, SesName)
		if err != nil {
			log.Error(err)
		}

		session.Values = map[interface{}]interface{}{}
		err = session.Save(r, rw)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

	}).Methods("POST")

	r.HandleFunc("/Notes/{userid}", func(rw http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, SesName)
		if err != nil {
			log.Error(err)
		}

		if len(session.Values) == 0 {
			http.Error(rw, "Login first", http.StatusUnauthorized)
			return
		}
		userid, err := strconv.Atoi(mux.Vars(r)["userid"])
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
		}
		var (
			UserData User
		)

		UserData.AddUserID(int64(userid))
		UserNotes, err := UserData.GetNotes()
		if err != nil {
			http.Error(rw, err.Error(),
				http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(UserNotes)
		rw.WriteHeader(http.StatusOK)

	}).Methods("GET")

	r.HandleFunc("/Notes/{userid}/Add", func(rw http.ResponseWriter, r *http.Request) {
		userid, err := strconv.Atoi(mux.Vars(r)["userid"])
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
		}
		var (
			NoteUser = make(map[string]interface{})
		)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(body, &NoteUser)
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
			return
		}

		if len(NoteUser["Text"].(string)) > 200 {
			http.Error(rw, "Text out of range",
				http.StatusBadRequest)
			return
		}

		NoteData := Note{
			UsersID: int64(userid),
			Data:    NoteUser["Text"].(string),
			Encrypt: NoteUser["Encrypt"].(bool),
			Title:   NoteUser["Title"].(string),
		}
		err = NoteData.AddNote()
		if err != nil {
			http.Error(rw, err.Error(),
				http.StatusInternalServerError)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"Status": 200,
		})
		rw.WriteHeader(http.StatusOK)

	}).Methods("PUT")

	r.HandleFunc("/Notes/{userid}/Update/{noteid}", func(rw http.ResponseWriter, r *http.Request) {
		userid, err := strconv.Atoi(mux.Vars(r)["userid"])
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
		}
		noteid, err := strconv.Atoi(mux.Vars(r)["noteid"])
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
		}
		var (
			NoteUser map[string]interface{}
		)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(body, &NoteUser)
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
			return
		}
		NoteData := Note{
			UsersID: int64(userid),
			Data:    NoteUser["Text"].(string),
			ID:      int64(noteid),
			Title:   NoteUser["Title"].(string),
		}

		err = NoteData.UpdateNote()
		if err != nil {
			http.Error(rw, err.Error(),
				http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"Status": 200,
		})
		rw.WriteHeader(http.StatusOK)

	}).Methods("PATCH")

	r.HandleFunc("/Notes/{userid}/Delete/{noteid}", func(rw http.ResponseWriter, r *http.Request) {
		userid, err := strconv.Atoi(mux.Vars(r)["userid"])
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
		}
		noteid, err := strconv.Atoi(mux.Vars(r)["noteid"])
		if err != nil {
			http.Error(rw, "Error reading request body",
				http.StatusBadRequest)
		}

		NoteData := Note{
			UsersID: int64(userid),
			ID:      int64(noteid),
		}

		err = NoteData.DeleteNote()
		if err != nil {
			http.Error(rw, err.Error(),
				http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"Status": 200,
		})
		rw.WriteHeader(http.StatusOK)

	}).Methods("DELETE")

	// Bind to a port and pass our router in
	var handler http.Handler
	{
		handler = handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:8080"}),
			handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Origin", "Authorization", "Content-Type"}),
			handlers.ExposedHeaders([]string{""}),
			handlers.MaxAge(10),
			handlers.AllowCredentials(),
		)(r)
		handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
	}

	log.Fatal(http.ListenAndServe(":2525", handler))
}
