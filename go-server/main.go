package main

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
)

type query struct{}
type someOtherType struct {
	_name string
}

func (_ *query) Hello() string { return "Hello, world!" }
func (_ *query) Some() *someOtherType { return &someOtherType{ _name: "Test"} }
func (s *someOtherType) Name() string { return s._name; }

func main() {
	s := `
                schema {
					query: Query
                }
                type Query {
					hello: String!
					some: SomeOtherType!
                }
				type SomeOtherType {
					name: String!
				}
        `
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	//message := strings.TrimPrefix(r.URL.Path, "/")
	//message = "Hello " + message + "!"
	//w.Header().Set("Content-Type", "application/json")
	//jsonD, jsonErr := json.Marshal(message)
	//if jsonErr != nil {
	//	w.Write([]byte("\"An error occurred while trying to process the request\""))
	//	w.WriteHeader(500)
	//}
	//w.Write(jsonD)
	//
	//
	//
	//
	//
	//					fields := graphql.Fields{
	//						"hello": &graphql.Field{
	//		Type: graphql.String,
	//		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	//			return "world", nil
	//		},
	//	},
	//}
	//rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	//schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	//schema, err := graphql.NewSchema(schemaConfig)
	//if err != nil {
	//	log.Fatalf("failed to create new schema, error: %v", err)
	//}
	//
	//// Query
	//query := `
	//	{
	//		hello
	//	}
	//`
	//params := graphql.Params{Schema: schema, RequestString: query}
	//rr := graphql.Do(params)
	//if len(rr.Errors) > 0 {
	//	log.Fatalf("failed to execute graphql operation, errors: %+v", rr.Errors)
	//}
	//w.Header().Set("Content-Type", "application/json")
	//rJSON, _ := json.Marshal(rr)
	//w.Write([]byte(rJSON))
}