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

# How it work

Test artifacts are from below folo report. 
https://indy-gateway.psi.redhat.com/api/folo/admin/build-A6RE4WO5CDYAA/report

The program parses the report and get all the 'downloads' and 'uploads' artifacts.

It then downloads those files from original site and stores to a local dir (e.g, ./test/content) so that it need not to download again in following runs.

The step two is to upload all those files on to target storage service which is based on AWS S3. We measure the upload performance during the uploading.

The step three is to download the files from the target storage service and measure the download performance.

We may use a few concurrent threads to do so (via Go rutines).

It prints each up/download like Maven, as:

> foo/bar/1.0/bar-1.0.jar (40Kb at 230kb/s)
> ...

We can also accumulate and make a final summary like:

> Total Upload: 120Mb at 3M/s
>
> Total Download: 500Mb at 25M/s