---
title: Google Cloud Platform Service Broker
type: Service Brokers
---

>**NOTE:** This addon is in the preview mode and is not available in Kyma by default. To enable it, add a proper AddonsConfiguration as described [here](#service-brokers-google-cloud-platform-service-broker-enable-gcp-service-broker).

The Google Cloud Platform (GCP) Service Broker is an open-source, [Open Service Broker](https://www.openservicebrokerapi.org/)-compatible API server that provisions managed services in the Google Cloud Platform public cloud. Kyma provides Namespace-scoped GCP Service Broker. In each Namespace, you can configure the GCP Service Broker against different subscriptions. Install the GCP Service Broker by provisioning the **Google Cloud Platform Service Broker** class provided by the Helm Broker.

![gcp broker class](./assets/gcp-class.png)

Once you provision the **Google Cloud Platform Service Broker** class, the GCP Service Broker classes are available in the **Service Catalog** view in a given Namespace.
The GCP Service Broker provides these Service Classes to use with the Service Catalog:

* Google BigQuery
* Google BigTable
* Google Cloud Storage
* Google CloudSQL for MySQL
* Google CloudSQL for PostgreSQL
* Google Cloud Dataflow
* Google Cloud Datastore
* Google Cloud Dialogflow
* Google Cloud Firestore
* Google Machine Learning APIs
* Google Cloud Memorystore for Redis API
* Google PubSub
* Google Spanner
* Stackdriver Debugger
* Stackdriver Monitoring
* Stackdriver Profiler
* Stackdriver Trace

See the details of each Service Class and its specification in the Service Catalog UI.
For more information about the Service Brokers, see [this](#service-brokers-service-brokers) document.

## Enable GCP Service Broker

The GCP Service Broker is a preview implementation and should not be used in a production environment. To enable the GCP Service Broker, create the following AddonsConfiguration:
```
apiVersion: addons.kyma-project.io/v1alpha1
kind: AddonsConfiguration
metadata:
  name: gcp-broker
  namespace: working-namespace
spec:
  reprocessRequest: 0
  repositories:
    - url: https://github.com/kyma-project/addons/releases/download/0.7.0/index-gcp.yaml
```
>**NOTE:** If you register this addons repository URL as cluster-wide, there will be a conflict with the next release, after the GCP Service Broker is enabled by default.
