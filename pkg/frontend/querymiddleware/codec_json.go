// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/querier/queryrange/query_range.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

package querymiddleware

import v1 "github.com/prometheus/prometheus/web/api/v1"

const jsonMimeType = "application/json"

type jsonFormat struct{}

func (j jsonFormat) EncodeResponse(resp *PrometheusResponse) ([]byte, error) {
	return json.Marshal(resp)
}

func (j jsonFormat) DecodeResponse(buf []byte) (*PrometheusResponse, error) {
	var resp PrometheusResponse

	if err := json.Unmarshal(buf, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (j jsonFormat) Name() string {
	return formatJSON
}

func (j jsonFormat) ContentType() v1.MIMEType {
	return v1.MIMEType{Type: "application", SubType: "json"}
}
