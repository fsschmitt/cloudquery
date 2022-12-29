// Code generated by codegen; DO NOT EDIT.

package aiplatform

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"google.golang.org/api/option"

	"google.golang.org/genproto/googleapis/cloud/location"

	"cloud.google.com/go/aiplatform/apiv1"
)

func ModelDeploymentMonitoringJobs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_model_deployment_monitoring_jobs",
		Description: `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.modelDeploymentMonitoringJobs#ModelDeploymentMonitoringJob`,
		Resolver:    fetchModelDeploymentMonitoringJobs,
		Multiplex:   client.ProjectMultiplexEnabledServices("aiplatform.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Endpoint"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("State"),
			},
			{
				Name:     "schedule_state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("ScheduleState"),
			},
			{
				Name:     "latest_monitoring_pipeline_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LatestMonitoringPipelineMetadata"),
			},
			{
				Name:     "model_deployment_monitoring_objective_configs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ModelDeploymentMonitoringObjectiveConfigs"),
			},
			{
				Name:     "model_deployment_monitoring_schedule_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ModelDeploymentMonitoringScheduleConfig"),
			},
			{
				Name:     "logging_sampling_strategy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LoggingSamplingStrategy"),
			},
			{
				Name:     "model_monitoring_alert_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ModelMonitoringAlertConfig"),
			},
			{
				Name:     "predict_instance_schema_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PredictInstanceSchemaUri"),
			},
			{
				Name:          "sample_predict_instance",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("SamplePredictInstance"),
				IgnoreInTests: true,
			},
			{
				Name:     "analysis_instance_schema_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AnalysisInstanceSchemaUri"),
			},
			{
				Name:     "bigquery_tables",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BigqueryTables"),
			},
			{
				Name:     "log_ttl",
				Type:     schema.TypeInt,
				Resolver: client.ResolveProtoDuration("LogTtl"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
			{
				Name:     "next_schedule_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("NextScheduleTime"),
			},
			{
				Name:     "stats_anomalies_base_directory",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StatsAnomaliesBaseDirectory"),
			},
			{
				Name:     "encryption_spec",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionSpec"),
			},
			{
				Name:     "enable_monitoring_pipeline_logs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableMonitoringPipelineLogs"),
			},
			{
				Name:     "error",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Error"),
			},
		},
	}
}

func fetchModelDeploymentMonitoringJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListModelDeploymentMonitoringJobsRequest{
		Parent: parent.Item.(*location.Location).Name,
	}
	if filterLocation(parent) {
		return nil
	}

	clientOptions := c.ClientOptions
	clientOptions = append([]option.ClientOption{option.WithEndpoint(parent.Item.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
	gcpClient, err := aiplatform.NewJobClient(ctx, clientOptions...)

	if err != nil {
		return err
	}
	it := gcpClient.ListModelDeploymentMonitoringJobs(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}