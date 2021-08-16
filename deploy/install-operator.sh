#!/bin/bash
kubectl apply -f ./crd.yaml
kubectl apply -f ./rbac-cluster.yaml -n operator
kubectl apply -f ./operator.yaml -n operator