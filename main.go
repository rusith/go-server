package main

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rusith/go-server/queries"
	"log"
	"net/http"
	"strings"
)

func StartQuery() {
	qs := []string {
		(*queries.Query).GetQuery(nil),
		(*queries.UserType).GetQuery(nil),
	}
	var queryBuilder strings.Builder
	queryBuilder.WriteString(`
     schema {
	    query: Query
	 }
`)
	for _, q := range qs {
		queryBuilder.WriteString(q)
	}
	schema := graphql.MustParseSchema(queryBuilder.String(), &queries.Query{})
	http.Handle("/q", &relay.Handler{Schema: schema})
}

func main() {
	StartQuery()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
