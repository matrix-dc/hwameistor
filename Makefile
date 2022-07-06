include Makefile.variables
# local-disk-manager build definitions
LDM_MODULE_NAME = local-disk-manager
LDM_BUILD_INPUT = ${CMDS_DIR}/${LDM_MODULE_NAME}/diskmanager.go

# local-storage build definitions
LS_MODULE_NAME = local-storage
LS_BUILD_INPUT = ${CMDS_DIR}/${LS_MODULE_NAME}/storage.go

.PHONY: compile
compile: compile_ldm compile_ls

.PHONY: image
image: build_ldm_image build_ls_image

.PHONY: release
release: release_ldm release_ls

.PHONY: vendor
vendor:
	go mod tidy -compat=1.17
	go mod vendor

.PHONY: release_ldm
release_ldm:
	# build for amd64 version
	${DOCKER_MAKE_CMD} make compile_ldm
	${DOCKER_BUILDX_CMD_AMD64} -t ${LDM_IMAGE_NAME}:${RELEASE_TAG}-amd64 -f ${LDM_IMAGE_DOCKERFILE} ${PROJECT_SOURCE_CODE_DIR}
	# build for arm64 version
	${DOCKER_MAKE_CMD} make compile_arm64
	${DOCKER_BUILDX_CMD_ARM64} -t ${LDM_IMAGE_NAME}:${RELEASE_TAG}-arm64 -f ${LDM_IMAGE_DOCKERFILE} ${PROJECT_SOURCE_CODE_DIR}
	# push to a public registry
	${MUILT_ARCH_PUSH_CMD} ${LDM_IMAGE_NAME}:${RELEASE_TAG}

.PHONY: release_ls
release_ls:
	# build for amd64 version
	${DOCKER_MAKE_CMD} make compile_ls
	${DOCKER_BUILDX_CMD_AMD64} -t ${LS_IMAGE_NAME}:${RELEASE_TAG}-amd64 -f ${LS_IMAGE_DOCKERFILE} ${PROJECT_SOURCE_CODE_DIR}
	# build for arm64 version
	${DOCKER_MAKE_CMD} make compile_arm64
	${DOCKER_BUILDX_CMD_ARM64} -t ${LS_IMAGE_NAME}:${RELEASE_TAG}-arm64 -f ${LS_IMAGE_DOCKERFILE} ${PROJECT_SOURCE_CODE_DIR}
	# push to a public registry
	${MUILT_ARCH_PUSH_CMD} ${LS_IMAGE_NAME}:${RELEASE_TAG}

.PHONY: build_ldm_image
build_ldm_image:
	@echo "Build local-disk-manager image ${LDM_IMAGE_NAME}:${IMAGE_TAG}"
	${DOCKER_MAKE_CMD} make compile_ldm
	docker build -t ${LDM_IMAGE_NAME}:${IMAGE_TAG} -f ${LDM_IMAGE_DOCKERFILE} ${PROJECT_SOURCE_CODE_DIR}

.PHONY: build_ls_image
build_ls_image:
	@echo "Build local-storage image ${LDM_IMAGE_NAME}:${IMAGE_TAG}"
	${DOCKER_MAKE_CMD} make compile_ls
	docker build -t ${LS_IMAGE_NAME}:${IMAGE_TAG} -f ${LS_IMAGE_DOCKERFILE} ${PROJECT_SOURCE_CODE_DIR}

.PHONY: apis
apis:
	${DOCKER_MAKE_CMD} make _gen-apis

.PHONY: builder
builder:
	docker build -t ${BUILDER_NAME}:${BUILDER_TAG} -f ${BUILDER_DOCKERFILE} ${PROJECT_SOURCE_CODE_DIR}

.PHONY: _gen-apis
_gen-apis:
	${OPERATOR_CMD} generate k8s
	${OPERATOR_CMD} generate crds
	bash hack/update-codegen.sh

.PHONY: compile_ldm
compile_ldm:
	GOARCH=amd64 ${BUILD_ENVS} ${BUILD_CMD} ${BUILD_OPTIONS} -o ${LDM_BUILD_OUTPUT} ${LDM_BUILD_INPUT}

.PHONY: compile_ls
compile_ls:
	GOARCH=amd64 ${BUILD_ENVS} ${BUILD_CMD} ${BUILD_OPTIONS} -o ${LS_BUILD_OUTPUT} ${LS_BUILD_INPUT}