package myspan

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.opencensus.io/exporter/zipkin"
	"go.opencensus.io/trace"

	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
)
type CustomSpan struct {
	*trace.Span
}

type enumtype string
const (
	always  enumtype = "always"
	prob enumtype = "prob"
	never enumtype = "never"
)

func CreateSpan(name string, ep string, traceSample enumtype, args interface{}) *CustomSpan {
	fmt.Println("Create span")
	localEndpoint, err := openzipkin.NewEndpoint(name, ep)
	if err != nil {
		log.Fatalf("Failed to create the local zipkinEndpoint: %v", err)
	}
	reporter := zipkinHTTP.NewReporter("http://localhost:9411/api/v2/spans")
	ze := zipkin.NewExporter(reporter, localEndpoint)
	trace.RegisterExporter(ze)
	var probility float64
	if value, ok := args.(float64); ok {
		probility = value
	} else {
		fmt.Println("args type is wrong! Please input float")
		os.Exit(1)
	}
	fmt.Println("prob:", prob)
	switch traceSample {
	case always:
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	case prob:
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(probility)})
	case never:
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.NeverSample()})
	}
	_, span := trace.StartSpan(context.Background(), name)
	fmt.Println("finish")
	return &CustomSpan{span}
}


