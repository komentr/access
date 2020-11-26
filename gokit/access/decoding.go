package access

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"gitlab.com/komentr/access/models"

	"github.com/julienschmidt/httprouter"
)

func decodeNewUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body = new(models.UserRequest)
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	// jwt := r.Context().Value("jwtclaim").(models.JwtClaim)
	// body.TID = jwt.TID
	// return newRequest{Req: body, Ctx: r.Context()}, nil
	return body, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body = new(models.User)
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func decodeIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := r.Context().Value("params").(httprouter.Params)
	sid := vars[0].Value
	var idreq = new(models.IDRequest)
	idreq.ID = sid
	return idreq, nil
}

func decodeSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {

	p := new(models.UserSearchRequest)
	var err error

	slimit := r.URL.Query()["limit"]
	if len(slimit) > 0 {
		p.Limit, err = strconv.Atoi(slimit[0])
		if err != nil {
			return p, errors.New("fail parsing limit")
		}
	} else {
		p.Limit = defaultLimit
	}

	spage := r.URL.Query()["page"]
	if len(spage) > 0 {
		p.Page, _ = strconv.Atoi(spage[0])
	} else {
		p.Page = defaultPage
	}

	sbssid := r.URL.Query()["tid"]
	if len(sbssid) > 0 {
		p.TID = sbssid[0]
	}

	srole := r.URL.Query()["role"]
	if len(srole) > 0 {
		p.Role, _ = strconv.Atoi(srole[0])
	}

	suname := r.URL.Query()["username"]
	if len(suname) > 0 {
		p.Username = suname[0]
	}

	sname := r.URL.Query()["name"]
	if len(sname) > 0 {
		p.Name = sname[0]
	}

	sactive := r.URL.Query()["active"]
	if len(sactive) > 0 {
		p.Active, _ = strconv.Atoi(sactive[0])
	}

	sver := r.URL.Query()["verified"]
	if len(sver) > 0 {
		p.Verified, _ = strconv.Atoi(sver[0])
	}

	sregby := r.URL.Query()["regby"]
	if len(sregby) > 0 {
		p.RegBy = sregby[0]
	}

	ssort := r.URL.Query()["sort"]
	if len(ssort) > 0 {
		p.Sortby = ssort[0]
	}

	sorder := r.URL.Query()["order"]
	if len(sorder) > 0 {
		p.Order = sorder[0]
	}

	return p, nil
}
