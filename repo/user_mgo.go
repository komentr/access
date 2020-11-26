package repo

import (
	"context"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gitlab.com/komentr/access/models"
	"gitlab.com/komentr/access/utils/errorcode"
	"gitlab.com/komentr/access/utils/utilmongo"
)

type userMgoRepo struct {
	DB *mgo.Database
}

func NewUserMgoRepo(db *mgo.Database) UserRepo {
	return &userMgoRepo{
		DB: db,
	}
}

func (s *userMgoRepo) Create(ctx context.Context, req models.User) (string, error) {
	if req.ID == "" {
		req.ID = bson.NewObjectId().Hex()
	}
	return req.ID, s.DB.C(models.UserColl).Insert(req)
}

func (s *userMgoRepo) Update(ctx context.Context, req models.User) (string, error) {
	return req.ID, s.DB.C(models.UserColl).UpdateId(req.ID, req)
}

func (s *userMgoRepo) Get(ctx context.Context, req models.IDRequest) (models.User, error) {

	u := models.User{}
	if err := s.DB.C(models.UserColl).FindId(req.ID).One(&u); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return u, errorcode.ErrNotFound
		default:
			return u, err
		}
	}

	return u, nil
}

func (s *userMgoRepo) Search(ctx context.Context, req models.UserSearchRequest) (interface{}, error) {
	// fmt.Printf("search req: %+v\n", req)
	fields := bson.M{}
	query := bson.M{}

	if req.TID != "" {
		query["tid"] = req.TID
	}
	if req.Username != "" {
		query["username"] = bson.M{"$regex": bson.RegEx{Pattern: req.Username, Options: "i"}}
	}
	if req.Name != "" {
		query["name"] = bson.M{"$regex": bson.RegEx{Pattern: req.Name, Options: "i"}}
	}
	if req.Role > 0 {
		query["role"] = req.Role
	}
	if req.Active > 0 {
		query["active"] = req.Active
	}
	if req.Verified > 0 {
		query["verified"] = req.Verified
	}
	// fmt.Printf("search query: %+v\n", query)

	// u := []models.User{}
	u := new([]models.User)
	return utilmongo.Paginate(s.DB, models.UserColl, query, fields, u, req.Page, req.Limit, []string{})
}

func (s *userMgoRepo) Delete(ctx context.Context, req models.IDRequest) error {
	return s.DB.C(models.UserColl).RemoveId(req.ID)
}

func (s *userMgoRepo) Find(ctx context.Context, req models.UserFindRequest) (models.User, error) {
	u := models.User{}
	if len(req.Username) == 0 {
		return u, errorcode.ErrInvalidArgument
	}
	err := s.DB.C(models.UserColl).Find(bson.M{"username": req.Username}).One(&u)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			return u, errorcode.ErrNotFound
		default:
			return u, err
		}
	}

	return u, err
}
