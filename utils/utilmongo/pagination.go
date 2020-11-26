package utilmongo

import (
	"fmt"
	"math"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Pagination represents a paging
type Pagination struct {
	HasNext  bool        `json:"has_next"`
	HasPrev  bool        `json:"has_prev"`
	PerPage  int         `json:"perpage"`
	Page     int         `json:"page"`
	LastPage int         `json:"last_page"`
	Total    int         `json:"total"`
	Items    interface{} `json:"items"`
}

//Paginate do paginate query
func Paginate(db *mgo.Database, col string, query, fields bson.M, result interface{}, page, perPage int, sorts []string) (*Pagination, error) {
	var skip int

	if page > 0 {
		skip = (page - 1) * perPage
	}

	if err := db.C(col).Find(query).Select(fields).Sort(sorts...).Skip(skip).Limit(perPage).All(result); err != nil {
		return nil, err
	}

	total, err := db.C(col).Find(query).Count()
	if err != nil {
		return nil, err
	}
	pages := int(math.Ceil(float64(total) / float64(perPage)))
	hasNext := page < pages
	hasPrev := page > 1 && pages > 0

	p := Pagination{
		HasNext:  hasNext,
		HasPrev:  hasPrev,
		PerPage:  perPage,
		LastPage: pages,
		Page:     page,
		Total:    total,
		Items:    &result,
	}
	return &p, nil
}

func strToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// objectIDFromHex safely returns ObjectId from a hex representation.
// The conversion will not raise panic.
func objectIDFromHex(hex string) (bson.ObjectId, error) {
	var oid bson.ObjectId
	// check whether a string is a valid hex representation of an ObjectId
	if !bson.IsObjectIdHex(hex) {
		return oid, fmt.Errorf("invalid hex representation:%s", hex)
	}

	// this will raise panic if a string is not a valid hex, hence we checked the ID in step above
	oid = bson.ObjectIdHex(hex)
	return oid, nil
}

// Pipeline paginate implement paginate in with pipeline operations
func PipelinePaginate(c *mgo.Collection, operations []bson.M, result interface{}, page, perPage, total int) (*Pagination, error) {
	var skip int

	if page > 0 {
		skip = (page - 1) * perPage
	}

	//skip and limit
	skipDoc := bson.M{"$skip": skip}
	limitDoc := bson.M{"$limit": perPage}

	operations = append(operations, skipDoc)
	operations = append(operations, limitDoc)

	if err := c.Pipe(operations).All(result); err != nil {
		return nil, err
	}

	pages := int(math.Ceil(float64(total) / float64(perPage)))
	hasNext := page < pages
	hasPrev := page > 1 && pages > 0

	p := Pagination{
		HasNext: hasNext,
		HasPrev: hasPrev,
		Items:   &result,
	}
	return &p, nil
}
