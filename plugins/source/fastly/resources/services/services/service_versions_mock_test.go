package services

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/fastly/go-fastly/v7/fastly"
	"github.com/golang/mock/gomock"
)

func buildServiceVersionsMock(t *testing.T, m *mocks.MockFastlyClient) {
	d := make([]*fastly.Version, 0, 1)
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListVersions(gomock.Any()).Return(d, nil)
}
