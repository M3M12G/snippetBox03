package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)
	mux := pat.New()
	mux.Get(newrelic.WrapHandle(app.newrelic, "/", dynamicMiddleware.ThenFunc(app.home)))

	mux.Get(newrelic.WrapHandle(app.newrelic, "/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm)))
	mux.Post(newrelic.WrapHandle(app.newrelic, "/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet)))
	mux.Get(newrelic.WrapHandle(app.newrelic, "/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet)))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)
}

//return app.recoverPanic(app.logRequest(secureHeaders(mux)))
