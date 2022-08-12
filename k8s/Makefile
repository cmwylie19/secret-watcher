# Environment variables 
# ---------------------------
SHELL=bash
export HOSTNAME?= $(shell hostname)


#---------------------------
# Deploy to dev (Kind)
.PHONY: deploy/dev
deploy/dev:
	@cat stage/custom.yaml |  sed 's/<hostname>/$(HOSTNAME)/g' > stage/kustomization.yaml
	@kubectl apply -k stage --dry-run=client -oyaml
	@rm stage/kustomization.yaml

# Deploy to stage (arm64)
.PHONY: deploy/stage
deploy/stage:
	@cat stage/custom.yaml |  sed 's/<hostname>/$(HOSTNAME)/g' > stage/kustomization.yaml
	@kubectl apply -k stage --dry-run=client -oyaml
	@rm stage/kustomization.yaml


# Deploy to prod (OpenShift)
.PHONY: deploy/prod
deploy/prod:
	@cat prod/custom.yaml |  sed 's/<hostname>/$(HOSTNAME)/g' > prod/kustomization.yaml
	@kubectl apply -k prod --dry-run=client -oyaml
	@rm prod/kustomization.yaml
