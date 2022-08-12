# Environment variables 
# ---------------------------
SHELL=bash
export HOSTNAME?= $(shell hostname)


#---------------------------

# Deploy to stage for dev and test
.PHONY: deploy/stage
deploy/stage:
	@echo "This is $(HOSTNAME)"
	@echo "Deploying secret-watcher to stage"
	@cat stage/custom.yaml |  sed 's/<hostname>/$(HOSTNAME)/g' > stage/kustomization.yaml
	@kubectl apply -k stage --dry-run=client -oyaml
	@rm stage/kustomization.yaml


# Deploy to prod for OpenShift
.PHONY: deploy/prod
deploy/prod:
	@echo "This is $(HOSTNAME)"
	@echo "Deploying secret-watcher to prod"
	@cat prod/custom.yaml |  sed 's/<hostname>/$(HOSTNAME)/g' > prod/kustomization.yaml
	@kubectl apply -k prod --dry-run=client -oyaml
	@rm prod/kustomization.yaml
