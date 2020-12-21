package client

import (
	"github.com/gogo/protobuf/types"
	"github.com/pachyderm/pachyderm/src/client/admin"
	"github.com/pachyderm/pachyderm/src/client/pkg/grpcutil"
)

// InspectCluster retrieves cluster state
func (c APIClient) InspectCluster() (*admin.ClusterInfo, error) {
	clusterInfo, err := c.AdminAPIClient.InspectCluster(c.Ctx(), &types.Empty{})
	if err != nil {
		return nil, grpcutil.ScrubGRPC(err)
	}
	return clusterInfo, nil
}
