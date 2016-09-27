DIR := ${CURDIR}
MYSECRET := ${DEMOSHASALT} 
builddredd: 
	docker build -t dreddtest -f dredd/DockerDredd .

dreddtest: builddredd
	docker run\
		-e "DEMOSHASALT=$(MYSECRET)"\
		--entrypoint dredd/rundredd.sh\
		dreddtest
