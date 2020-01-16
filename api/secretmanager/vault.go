//
// secretmanager.go
//
// Copyright (c) 2019 Markku Rossi
//
// All rights reserved.
//

package secretmanager

import (
	"context"
	"fmt"

	api "cloud.google.com/go/secretmanager/apiv1beta1"
	"github.com/markkurossi/go-libs/fn"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
)

type SecretManager struct {
	ctx       context.Context
	projectID string
	client    *api.Client
}

func NewSecretManager() (*SecretManager, error) {
	ctx := context.Background()
	id, err := fn.GetProjectID()
	if err != nil {
		return nil, err
	}

	client, err := api.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &SecretManager{
		ctx:       ctx,
		projectID: id,
		client:    client,
	}, nil
}

func (sm *SecretManager) Create(name string, data []byte) error {
	createReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", sm.projectID),
		SecretId: name,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}

	secret, err := sm.client.CreateSecret(sm.ctx, createReq)
	if err != nil {
		return err
	}

	addReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: data,
		},
	}

	_, err = sm.client.AddSecretVersion(sm.ctx, addReq)
	return err
}

func (sm *SecretManager) Get(name, version string) ([]byte, error) {
	if len(version) == 0 {
		version = "latest"
	}
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/%s",
			sm.projectID, name, version),
	}
	resp, err := sm.client.AccessSecretVersion(sm.ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Data, nil
}
