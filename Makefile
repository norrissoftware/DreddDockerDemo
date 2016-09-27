builddredd: 
	docker build -t dreddtest -f dredd/DockerDredd .

dreddtest: builddredd
	@docker run\
		-e "DEMOSHASALT=$(DEMOSHASALT)"\
		--entrypoint dredd/rundredd.sh\
		dreddtest
