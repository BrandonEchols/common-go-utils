package routing

import (
	"github.com/BrandonEchols/common-go-utils/common"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
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
