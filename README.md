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
3. Deploy chart
```
helm install artifactory \
    --namespace mynamespace \
    -f templates/artifactory.yaml \
    --version 8.4.8 \
    jfrog/artifactory
```
4. Login to artifactory
    * Default username: `admin`
    * Default password: `password`
5. Create a new docker registry called `demo`


### Deploy Jenkins

1. Add helm stable repo
```
helm repo add stable https://kubernetes-charts.storage.googleapis.com
```
2.  Install Jenkins
```
helm install jenkins \
    --namespace testing \
    -f templates/jenkins.yaml \
    --version 1.9.17 \
    stable/jenkins
```
3. Get Jenkins password
```
kubectl get secret --namespace testing jenkins -o jsonpath="{.data.jenkins-admin-password}" | base64 --decode | pbcopy
```
4. Login to Jenkins
5. Add Credentials
    * these will be the login credentials to artifactory
    * name them `Artifactory`
6. Run the demo pipeline

### Deploy Custom app

1. Create secret for artifactory
```
kubectl create secret \
    docker-registry regcred \
    --docker-server=artifactory.netp5fki6k.nks.cloud \
    --docker-username=admin \
    --docker-password=start123
```

2. Deploy Chart from repo
```
k create ns rocket
cd charts/website
helm install rocket --namespace rocket .
```