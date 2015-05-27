package main

import (
	"os"
//	"time"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"github.com/codegangsta/martini"
)


// The one and only martini instance.
var m *martini.Martini

func init() {
	m = martini.New()
	// Setup middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())

	m.Use(MapEncoder)
	// Setup routes
	r := martini.NewRouter()
	r.Get(``, func() string {
		return "Welcome to restApi framework in Continuum blaffa\n"
	})
//------------------------------------------------------------------------
//----------------------- PORT -------------------------------------------
//------------------------------------------------------------------------
	r.Post   (`/fil/:data`   ,filData)
	r.Get    (`/fil`         ,getDate)
	r.Put    (`/fil/:data`   ,putData)
	r.Delete (`/fil`         ,delData)
//------------------------------------------------------------------------
//----------------------- PORT -------------------------------------------
//------------------------------------------------------------------------
	r.Post   (`/port/:portno` ,postPort)
	r.Get    (`/port`         ,getPort)
	r.Put    (`/port/:portno` ,putPort)
	r.Delete (`/port`         ,delPort)
//------------------------------------------------------------------------
//----------------------- Response String---------------------------------
//------------------------------------------------------------------------
	r.Post   (`/response/:string`, postResponse)
	r.Put    (`/response/:string`, putResponse)
	r.Get    (`/response`,         getResponse)
	r.Delete (`/response`,         delResponse)
	// Add the router action
	m.Action(r.Handle)
}

// The regex to check for the requested format (allows an optional trailing
// slash).
var rxExt = regexp.MustCompile(`(\.(?:xml|text|json))\/?$`)

// MapEncoder intercepts the request's URL, detects the requested format,
// and injects the correct encoder dependency for this request. It rewrites
// the URL to remove the format extension, so that routes can be defined
// without it.
func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request) {
	// Get the format extension
	matches := rxExt.FindStringSubmatch(r.URL.Path)
	ft := ".json"
	if len(matches) > 1 {
		// Rewrite the URL without the format extension
		l := len(r.URL.Path) - len(matches[1])
		if strings.HasSuffix(r.URL.Path, "/") {
			l--
		}
		r.URL.Path = r.URL.Path[:l]
		ft = matches[1]
	}
	// Inject the requested encoder
	switch ft {
	case ".xml":
		c.MapTo(xmlEncoder{}, (*Encoder)(nil))
		w.Header().Set("Content-Type", "application/xml")
	case ".text":
		c.MapTo(textEncoder{}, (*Encoder)(nil))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	default:
		c.MapTo(jsonEncoder{}, (*Encoder)(nil))
		w.Header().Set("Content-Type", "application/json")
	}
}

func main() {
	go func() {

	}()
	portAPI := os.Getenv("PORT")
	// When running outside CF
	if portAPI == "" {
		portAPI = "40444" 
	}
	fmt.Println("Listening on PORT:", portAPI)
	
	if err := http.ListenAndServe (":"+portAPI, m);  err != nil {
		log.Fatal("Not so good", err)
	}

}
