//
// client.go
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

type Client struct {
	ctx           context.Context
	projectID     string
	secretManager *api.Client
}

func NewClient() (*Client, error) {
	ctx := context.Background()
	id, err := fn.GetProjectID()
	if err != nil {
		return nil, err
	}

	client, err := api.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		ctx:           ctx,
		projectID:     id,
		secretManager: client,
	}, nil
}

func (client *Client) Create(name string, data []byte) error {
	createReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", client.projectID),
		SecretId: name,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}

	secret, err := client.secretManager.CreateSecret(client.ctx, createReq)
	if err != nil {
		return err
	}

	addReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: data,
		},
	}

	_, err = client.secretManager.AddSecretVersion(client.ctx, addReq)
	return err
}

func (client *Client) Get(name, version string) ([]byte, error) {
	if len(version) == 0 {
		version = "latest"
	}
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/%s",
			client.projectID, name, version),
	}
	resp, err := client.secretManager.AccessSecretVersion(client.ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Data, nil
}
