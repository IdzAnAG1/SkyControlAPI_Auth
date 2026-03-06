REPO=https://github.com/IdzAnAG1/SkyControlAPI_Contracts.git#branch=main


buf_gen:
	buf generate $(REPO) --template buf.gen.yaml --path proto/auth/v1