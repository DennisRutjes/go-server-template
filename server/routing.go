package server

import (
	"net/http"
	"regexp"
	"sort"
)

type routes []*route

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	Routes routes
}

func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.Routes = append(h.Routes, &route{pattern, handler})
	sort.Sort(sort.Reverse(h.Routes))
}

func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.Routes = append(h.Routes, &route{pattern, http.HandlerFunc(handler)})
	sort.Sort(sort.Reverse(h.Routes))
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.Routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

func (s routes) Len() int {
	return len(s)
}
func (s routes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s routes) Less(i, j int) bool {
	return len(s[i].pattern.String()) < len(s[j].pattern.String())
}

func (server *Server) Router() *RegexpHandler {
	router := &RegexpHandler{}
	router.HandleFunc(regexp.MustCompile("/hello"), server.HandlerForSayHello())
	return router
}
