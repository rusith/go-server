package main

import (
	"./queries"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
	"strings"
)

func main() {
	qs := []string {
		(*queries.Query).GetQuery(nil),
		(*queries.User).GetQuery(nil),

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
	http.Handle("/queries", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
