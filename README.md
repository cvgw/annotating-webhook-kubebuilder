# Description
A simple mutating admission webhook implemented using the `kubebuilder` framework

# Requirements
Install `goenv`

Install `dep`

Install `kustomize`
```
$ brew install kustomize
```

Install `kubebuilder`
```
$ ./install_kubebuilder.sh
```

# Build and deploy
````
cd ./src
export IMG=your-image-name
make docker-build
make manifests
kustomize build config/default | kubectl apply -f -
````

# Delete resources
````
kubectl -n annotator-system delete \
    clusterrole.rbac.authorization.k8s.io/annotator-manager-role \
    clusterrolebinding.rbac.authorization.k8s.io/annotator-manager-rolebinding \
    service/annotator-controller-manager-service \
    statefulset.apps/annotator-controller-manager
````
