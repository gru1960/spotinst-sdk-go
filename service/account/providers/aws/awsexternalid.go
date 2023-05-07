package aws

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/jsonutil"
)

type AwsAccountExternalId struct {
	ExternalId *string `json:"externalId,omitempty"`
	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// forceSendFields is a list of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string

	// nullFields is a list of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string
}

func (o AwsAccountExternalId) MarshalJSON() ([]byte, error) {
	type noMethod AwsAccountExternalId
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

type CreateAWSAccountExternalIdInput struct {
	AWSAccountExternalId *string `json:"AWSAccountExternalId,omitempty"`
}
type CreateAWSAccountExternalIdOutput struct {
	AWSAccountExternalId *AwsAccountExternalId `json:"AWSAccountExternalId,omitempty"`
}

func (s *ServiceOp) CreateAwsAccountExternalId(ctx context.Context, input *string) (*CreateAWSAccountExternalIdOutput, error) {
	r := client.NewRequest(http.MethodPost, "/setup/credentials/aws/externalId")
	r.Params.Set("accountId", spotinst.StringValue(input))

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := awsAccountExternalIdFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateAWSAccountExternalIdOutput)
	if len(gs) > 0 {
		output.AWSAccountExternalId = gs[0]
	}

	return output, nil
}

func awsAccountExternalIdFromHttpResponse(resp *http.Response) ([]*AwsAccountExternalId, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return awsAccountExternalIdsFromJSON(body)
}

func awsAccountExternalIdsFromJSON(in []byte) ([]*AwsAccountExternalId, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*AwsAccountExternalId, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := awsAccountExternalIdFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func awsAccountExternalIdFromJSON(in []byte) (*AwsAccountExternalId, error) {
	b := new(AwsAccountExternalId)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}
