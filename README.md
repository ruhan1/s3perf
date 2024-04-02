# s3perf

This program tests the performance of a AWS S3 based storage service. You can simple run it by:
go run cmd/main.go

You can also create Openshift pipeline (and tasks) by below and run the pipeline via 'tkn' or Openshift console.

oc apply -f task-s3perf-prepare.yml

oc apply -f task-s3perf-execute.yml

oc apply -f s3perf-pipeline.yml

oc apply -f s3perf-pvc.yml

You can run below command to start a pipeline:

$ tkn pipeline start s3perf-pipeline -w name=shared-workspace,claimName=s3perf

NOTE: If you haven't installed the 'tkn', install it first. This is CLI for tekton pipelines.