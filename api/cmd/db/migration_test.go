package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"github.com/tektoncd/hub/api/pkg/testutils"
)

func TestMigration(t *testing.T) {
	tc := testutils.Setup(t)

	testutils.LoadFixtures(t, tc.FixturePath())

	api := tc.APIConfig
	logger := api.Logger()
	err := tc.Error()

	db := api.DB()

	if err = Migrate(api); err != nil {
		logger.Errorf("DB initialisation failed !!")
		return
	}
	logger.Info("DB initialisation successful !!")

	assert.NoError(t, err)

	assert.True(t, db.HasTable(&model.Catalog{}))

}
