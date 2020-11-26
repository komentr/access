package repo

import (
	"context"
	"strconv"
	"time"

	"gitlab.com/komentr/access/models"

	"github.com/go-redis/redis"

	"gitlab.com/komentr/utils/errorcode"
)

//UIDKey and other redis params
const (
	UIDKey        = "uid"
	UIDPrefixKey  = "uid:"
	UserKey       = "user"
	UserPrefixKey = "user:"
)

type userRedisRepo struct {
	Client *redis.Client
}

//NewUserRedis func declaration
func NewUserRedis(client *redis.Client) UserRepo {
	return &userRedisRepo{
		Client: client,
	}
}

func (s *userRedisRepo) Create(ctx context.Context, req models.User) (string, error) {
	//get new uid
	uid, err := s.Client.Incr(ctx, UIDKey).Uint64()
	if err != nil {
		return "", err
	}
	req.ID = strconv.FormatUint(uid, 10)

	//create entry userdata
	// data, err := json.Marshal(req)
	// if err != nil {
	// 	s.Client.Decr(ctx, UIDKey)
	// 	return "", err
	// }
	var m = make(map[string]interface{})
	m["id"] = req.ID
	m["username"] = req.Username
	m["password"] = req.Password
	m["name"] = req.Name
	m["role"] = req.Role
	m["image"] = req.Image
	m["active"] = req.Active
	m["verified"] = req.Verified
	m["regby"] = req.RegBy
	m["tid"] = req.TID
	m["created"] = req.Created

	if err := s.Client.HMSet(ctx, UIDPrefixKey+req.ID, m).Err(); err != nil {
		s.Client.Decr(ctx, UIDKey)
		return "", err
	}

	var u = make(map[string]interface{})
	u[req.Username] = req.ID

	if err := s.Client.HSet(ctx, UserKey, u).Err(); err != nil {
		s.Client.Del(ctx, UIDPrefixKey+req.ID)
		s.Client.Decr(ctx, UIDKey)
		return "", err
	}
	return req.ID, nil
}

func (s *userRedisRepo) Update(ctx context.Context, req models.User) (string, error) {
	var m = make(map[string]interface{})

	if req.Name != "" {
		m["name"] = req.Name
	}
	if req.Role > 0 {
		m["role"] = req.Role
	}
	if req.Image != "" {
		m["image"] = req.Image
	}
	if req.Active > 0 {
		m["active"] = req.Active
	}
	if req.Verified > 0 {
		m["verified"] = req.Verified
	}
	if req.TID != "" {
		m["tid"] = req.TID
	}

	if err := s.Client.HMSet(ctx, UIDPrefixKey+req.ID, m).Err(); err != nil {
		return "", err
	}
	return "ok", nil
}

func (s *userRedisRepo) Delete(ctx context.Context, req models.IDRequest) error {

	username, err := s.Client.HGet(ctx, UIDPrefixKey+req.ID, "username").Result()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	if err := s.Client.Del(ctx, UIDPrefixKey+req.ID).Err(); err != nil {
		return err
	}
	if err := s.Client.HDel(ctx, UserKey, username).Err(); err != nil {
		return err
	}
	return nil
}

func (s *userRedisRepo) Search(ctx context.Context, req models.UserSearchRequest) (interface{}, error) {

	return nil, nil
}

func (s *userRedisRepo) Get(ctx context.Context, req models.IDRequest) (models.User, error) {
	space := models.User{}
	res, err := s.Client.HGetAll(ctx, UIDPrefixKey+req.ID).Result()
	if err != nil {
		return space, err
	}
	// bytes := []byte(res)

	// if err := json.Unmarshal(bytes, &space); err != nil {
	// 	return space, err
	// }
	space.ID = res["id"]
	space.Username = res["username"]
	space.Password = res["password"]
	space.Name = res["name"]
	space.Role, _ = strconv.Atoi(res["role"])
	space.Image = res["image"]
	space.Active, _ = strconv.Atoi(res["active"])
	space.Verified, _ = strconv.Atoi(res["verified"])
	space.RegBy = res["regby"]
	space.TID = res["tid"]
	space.Created, _ = time.Parse(time.RFC3339Nano, res["created"])

	return space, err
}

func (s *userRedisRepo) Find(ctx context.Context, req models.UserFindRequest) (models.User, error) {
	resp := models.User{}

	if len(req.Username) == 0 {
		return resp, errorcode.ErrInvalidArgument
	}
	uid, err := s.Client.HGet(ctx, UserKey, req.Username).Result()
	if err != nil {
		return resp, err
	}

	return s.Get(ctx, models.IDRequest{ID: uid})
}
