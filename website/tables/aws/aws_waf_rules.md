# Table: aws_waf_rules

This table shows data for WAF Rules.

https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_RuleSummary.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|tags|JSON|
|predicates|JSON|
|rule_id|String|
|metric_name|String|
|name|String|