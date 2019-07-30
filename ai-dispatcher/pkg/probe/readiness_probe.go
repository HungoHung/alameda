package probe

import (
	"context"
	"fmt"

	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"google.golang.org/grpc"
)

type ReadinessProbeConfig struct {
	DatahubAddr string
}

func queryDatahub(datahubAddr string) error {
	conn, err := grpc.Dial(datahubAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	defer conn.Close()
	datahubServiceClnt := datahub_v1alpha1.NewDatahubServiceClient(conn)
	res, err := datahubServiceClnt.ListAlamedaNodes(context.Background(), &datahub_v1alpha1.ListAlamedaNodesRequest{})
	if err != nil {
		return err
	}

	if len(res.GetNodes()) == 0 {
		return fmt.Errorf("No nodes found in datahub")
	}

	return err
}