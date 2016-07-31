#!/bin/bash

# Exit on any error
set -e

sudo /opt/google-cloud-sdk/bin/gcloud docker push jasonwitting/simple-go-http
sudo chown -R ubuntu:ubuntu /home/ubuntu/.kube
kubectl patch deployment simple-go-http -p '{"spec":{"template":{"spec":{"containers":[{"name":"simple-go-http","image":"jasonwitting/simple-go-http:'"$CIRCLE_SHA1"'"}]}}}}'
