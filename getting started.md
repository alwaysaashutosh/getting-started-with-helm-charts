# Getting Started With Kubernetes Helm Charts
## Some Basic Helm Commands



1. ### helm create  [name-of-the-chart]
> ex:  helm create helloworld -----▶️   This will create a boilerplate for our  chart . A repo with name helloworld will be created along with the default files .


2. ### helm install [custom-release] [name-of-the-chart]
> ex: helm install myhelloworldrelease helloworld ------▶️ This will install the helloworld chart with the release name myhelloworldrelease to our cluster.


3. ### helm upgrade [release-name-which-is-already-deployed] [name-of-the-chart]
> ex: helm upgrade myhelloworldrelease helloworld -----▶️ This will upgrade the existing release .

4. ### helm rollback [release-whhich-is-already-deployed] [chart-version/revison]
> ex: helm rollback myhelloworldrelease 1 -----▶️ Here we want to skip to the 1st revison .

#### Note: The revison number will be incremented in a sequential way .

5. ### helm install [release-name] --debug --dry-run [chart-name]
> ex: helm install myhelloworldrelease --debug --dry-run helloworld  -----▶️ Here its debugging and doing a dry run, its recommended to use it before installing any helm chart .


<!-- Bash script block -->

```bash
     👨‍💻install.go:222: [debug] Original chart version: ""
     👨‍💻install.go:239: [debug] CHART PATH: /home/ashutosh/go/src/github.com/Golang/Getting started with HELM/helloworld
     👨‍💻NAME: myhelloworldrelease
     👨‍💻LAST DEPLOYED: Thu Aug  8 15:36:07 2024
     👨‍💻NAMESPACE: default
     👨‍💻STATUS: pending-install
    
```
#### It will validate with your kube-api server results you in the values which are going to be a part of that chart . The status will be pending-install 




6. ### helm template [chart-name]
> ex: helm template helloworld -----▶️ This command aims to render the charts locally .

#### Diffrence between --debug --dry-run and template is template validates the yaml without connecting to the k8s api server whereas in case of --debug --dry-run it will connect with the kube-api server .

7. ### helm lint [chart-name]
> ex: helm lint helloworld  -----▶️ This aims to find any errors or misconfiguration in the helm chart .

8. ### helm uninstall [release-name]
> ex: helm uninstall myhelloworldrelease -----▶️ This will remove myhelloworldrelease chart from k8s.

9. ### Installing a Golang webserver using Helm

10. ### Helmfile is an abstraction over helm chart commands, with a single command [helm sync] you will be able to do all the operations .
> ex: writing a helm file 

```yaml
        releases:
            - name: <release-name>
              chart: <location-of-your-chart/chart_name>
              install:   true/false (true in case if you want to install, false in case of uninstalling) 
```

11. ### Using github reposiory for installing helmchart using helmfile .

```yaml
        repositories:
            - name: <name-of-your-repo>
              url : <url-to-the-github-repo>
        releases:
            - name: <release-name>
              chart: <chart_name/location>
              installed:  true/false
```       
12. ### Install multiple helm charts using helmfile
``` yaml
        releases:
            - name: <release-name-1>
              chart: <location-of-your-chart/chart_name>
              install:   true/false (true in case if you want to install, false in case of uninstalling) 


            - name: <release-name-2>
              chart: <location-of-your-chart/chart_name>
              install:   true/false (true in case if you want to install, false in case of uninstalling) 
```
* ### Helm Repo (Used to import charts from various opensouce repositories)

13. ### helm search hub [repo-name] 
> ex: helm search hub wordpress  -----▶️ using search hub we can search for various charts and import it for our use .
    
    
 * #### helm search hub wordpress  --max-col-width=0 -----▶️ this will show us the full repo url.

14. ### helm repo add [name] [chart-url]
> ex: helm repo add bitnami https://charts.bitnami.com/bitnami -----▶️ In this example i am using bitnami as they provide the most stable chart in the helm world.



15. ### helm search repo [chart-name-to-be-searched] --versions
> ex: helm search repo wordpress --versions

16. ### helm show readme [chart-name] --version [chart-version]
>ex: helm show readme bitnami/wordpress --version 10.0.3 -----▶️ This command is basically used to get instruction from the readme file



17. ### Setup a user and Password for wordpress
> touch wordpress-values.yaml

```yaml
wordpressUsername: alwaysashutosh
wordpressPassword: ashutosh
wordpressEmail: contact@ashutosh.com
wordpressFirstName: Ashutosh
wordpressLastName: Pandey
wordpressBlogName: alwaysashutosh.com
service: 
  type: LoadBalancer
```


18. ### Installing wordpress Helm Chart 
```bash
helm install [release-name] [chart-name] --values=<yaml-used-to-override-the-chart> --namespace <namespace-name>  --version <helm-chart-version>
```
>ex: helm install wordpress bitnami/wordpress --values=wordpress-values.yaml --namespace nswordpress --version 10.0.3

*  --values OR -f flag in your Helm commands is used to override the values in a chart and pass in a new file.
* --namespace or -n flag is used to add namespace scope for this request.

19. ### Helm hooks  - helm hooks can be any kubernates resources ie; (Batch, Deployment, Service)

> hooks folder can be created inside the templates directory .

```
template
│      
│
└───hooks
│   │   pre-install.yaml | post-install.yaml
│   │

```
> For more info on hooks please visit this  (https://helm.sh/docs/topics/charts_hooks/)

20. ### Annotation are the most important part of hooks ---▶️ What makes them special is annotation we add to the metadata: section.
#### Pre-Install Hook.
```bash
        "helm.sh/hook": "pre-install"
        "helm.sh/hook-weight": "0"
        "helm.sh/hook-delete-policy": hook-succeeded
                                
```
#### Post-Install Hook.
```bash
        "helm.sh/hook": "post-install"
        "helm.sh/hook-weight": "5"
        "helm.sh/hook-delete-policy": hook-succeeded
```

21. ### helm test [deployed-release-name]  
>▶️ Its the kind of running an unit test in your release . You can find the test file inside templates/tests .


22. ### Push Helm charts to Dockerhub OCI(Open Container Initiative) .

    Login    ❎  
    ```bash
    helm registry login <host> -u <username> -p <password>  docker login command will also work .
    ```
    >eg :     helm registry login registry-1.docker.io -u your_username -p your_password

    Package  📦  
    ```bash
    helm package <chart-name> --version 0.1.2          | you can specify the version otherwise default version will be 0.1.0  | creates a tar file with <chart-name>-<version>.tgz
    ```
    
    Push|Pull📍 
    ```bash
    helm push <tarball_file_.tgz> <repourl>
    ```
    >eg: helm push charts-v0.1.0.tgz oci://registry-1.docker.io/<username/namespacename>
                         with the above command it will get pushed to the charts repo in your namespace with 0.1.0 version/tag | ashutoshk1/charts:v0.1.0
