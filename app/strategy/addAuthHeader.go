package strategy

import (
	"net/http"
	"net/http/httputil"
	"os"

	"context"
	"log"

	"google.golang.org/api/idtoken"

	"github.com/julienschmidt/httprouter"
)

func AddAuthHeader(rp *httputil.ReverseProxy) httprouter.Handle {

	//Getting envs
	aud := os.Getenv("SCHEME") + "://" + os.Getenv("HOST")
	jsonCert := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	ts, err := idtoken.NewTokenSource(context.Background(), aud, idtoken.WithCredentialsFile(jsonCert))
	if err != nil {
		log.Fatalf("unable to create TokenSource: %v", err)
	}

	//Returning our handler
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		tok, err := ts.Token()
		if err != nil {
			log.Fatalf("unable to retrieve Token: %v", err)
		}

		validTok, err := idtoken.Validate(context.Background(), tok.AccessToken, aud)

		if err != nil {
			log.Fatalf("token validation failed: %v", err)
		}

		if validTok.Audience != aud {
			log.Fatalf("got %q, want %q", validTok.Audience, aud)
		}

		r.Header.Add("Authorization", "Bearer "+tok.AccessToken)

		log.Println(aud + r.RequestURI)

		rp.ServeHTTP(w, r)
	}
}
