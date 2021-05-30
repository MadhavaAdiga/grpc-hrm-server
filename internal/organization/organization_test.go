package organization_test

import (
	"testing"
)

func TestCreateOrganization(t *testing.T) {
	t.Parallel()

	// 	ctrl := gomock.NewController(t)
	// 	defer ctrl.Finish()

	// 	store := mockdb.NewMockStore(ctrl)

	// 	arg := db.CreateOrganizationParam{
	// 		Name:      utils.RandomName(),
	// 		CreatedBy: utils.RandomName(),
	// 		Status:    0,
	// 		UpdatedBy: utils.RandomName(),
	// 		CreatorID: uuid.New().String(),
	// 		UpdaterID: uuid.New().String(),
	// 	}

	// 	store.EXPECT().CreateOrganization(gomock.Any(), arg).
	// 		Times(1).Return(db.Organization, nil)
}
