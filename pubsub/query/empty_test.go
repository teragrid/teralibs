package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teragrid/teralibs/pubsub"
	"github.com/teragrid/teralibs/pubsub/query"
)

func TestEmptyQueryMatchesAnything(t *testing.T) {
	q := query.Empty{}
	assert.True(t, q.Matches(pubsub.NewTagMap(map[string]interface{}{})))
	assert.True(t, q.Matches(pubsub.NewTagMap(map[string]interface{}{"Asher": "Roth"})))
	assert.True(t, q.Matches(pubsub.NewTagMap(map[string]interface{}{"Route": 66})))
	assert.True(t, q.Matches(pubsub.NewTagMap(map[string]interface{}{"Route": 66, "Billy": "Blue"})))
}
