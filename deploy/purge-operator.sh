#!/bin/bash
kubectl delete -f ./operator.yaml -n operator
kubectl delete -f ./rbac-cluster.yaml -n operator
kubectl delete -f ./crd.yaml
