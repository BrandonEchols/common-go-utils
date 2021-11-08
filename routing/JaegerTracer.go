/**
 * Copyright 2017,2019 InsideSales.com Inc.
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is the property of InsideSales.com, Inc. and its suppliers, if
 * any. The intellectual and technical concepts contained herein are proprietary and are protected by
 * trade secret or copyright law, and may be covered by U.S. and foreign patents and patents pending.
 * Dissemination of this information or reproduction of this material is strictly forbidden without prior
 * written permission from InsideSales.com Inc.
 *
 * Requests for permission should be addressed to the Legal Department, InsideSales.com,
 * 1712 South East Bay Blvd. Provo, UT 84606.
 *
 * The software and any accompanying documentation are provided "as is" with no warranty.
 * InsideSales.com, Inc. shall not be liable for direct, indirect, special, incidental, consequential, or other
 * damages, under any theory of liability.
 */
package routing

import (
	"bitbucket.xant.tech/ci/ci-go-utils/common"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const SPAN_TAG_KEY_IDM_GUID = "idm_guid"
const SPAN_TAG_KEY_IDM_ID = "idm_id"
const SPAN_TAG_KEY_ORG_ID = "org_id"
const SPAN_TAG_KEY_MESSAGE = "message"
const SPAN_TAG_KEY_RESPONSE_BODY = "http.response.body"
const SPAN_TAG_KEY_STATUS_CODE = "http.status_code"

func InitializeJaegerTracer(configs common.IConfigGetter) (func(), error) {

	collectorEndpoint := configs.MustGetConfigVar("JAEGER_COLLECTOR_ENDPOINT")
	serviceName := configs.MustGetConfigVar("JAEGER_SERVICE_NAME")
	podID := configs.MustGetConfigVar("JAEGER_POD_ID")

	// Create the a jaeger exporter
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(collectorEndpoint),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: serviceName,
			Tags:        []label.KeyValue{label.String("POD_ID", podID), label.String("service", serviceName)},
		}),
		jaeger.WithSDK(&sdktrace.Config{
			DefaultSampler: sdktrace.AlwaysSample(),
		}),
	)
	return flush, err
}
