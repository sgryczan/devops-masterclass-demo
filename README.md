## Steps

### Requirements
* A Kubernetes cluster
* `kubectl` configured to said cluster
* `helm` installed. In this example i'm using helm 3, but helm 2 with a valid tiller deployment configured should work too
* Artifactory trial license key - [See here](https://jfrog.com/artifactory/free-trial/)


### Deploy Artifactory

1. Add JFrog chart repo
```
helm repo add jfrog https://charts.jfrog.io
```
2. Update repos
```
helm repo update
```
3. Create a namespace for artifactory
```
k create ns artifactory
```
4. Deploy chart
```
helm install artifactory \
    --namespace artifactory \
    -f templates/artifactory.yaml \
    --version 8.4.8 \
    jfrog/artifactory
```
5. Login to artifactory by navigating to the hostname set for the ingress
    * Default username: `admin`
    * Default password: `password`
6. You'll be prompted to change the default admin password
7. Create a new docker registry called `demo`


### Deploy Jenkins

1. Add helm stable repo
```
helm repo add stable https://kubernetes-charts.storage.googleapis.com
```

2.  Install Jenkins
```
helm install jenkins \
    --namespace jenkins \
    -f templates/jenkins.yaml \
    --version 1.9.17 \
    stable/jenkins
```
3. Get Jenkins password
```
kubectl get secret --namespace jenkins jenkins -o jsonpath="{.data.jenkins-admin-password}" | base64 --decode
```
4. Login to Jenkins by navigating to the hostname specified for the ingress
5. Add Credentials
    * these will be the login credentials to artifactory
    * name them `Artifactory`
6. Run the `Demo` pipeline

### Deploy Custom app

1. Create a namespace for the application
```
k create ns rocket
```

1. Create secret for artifactory
```
kubectl create secret \
    docker-registry regcred \
    --docker-server=<ARTIFACTORY INGRESS HOSTNAME> \
    --docker-username=admin \
    --docker-password=start123
```

2. Deploy Chart from repo directory
```
cd charts/website
helm install rocket --namespace rocket .
```

3. Access the application by navigating to the ingress hostname set in the chart.