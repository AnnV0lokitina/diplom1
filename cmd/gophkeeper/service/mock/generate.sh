#!/bin/bash

mockgen -source=../interface/repo_interface.go -destination=repo_interface.go \
-aux_files=github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service/interface=../interface/repo_data_interface.go,\
github.com/AnnV0lokitina/diplom1/cmd/gophkeeper/service/interface=../interface/repo_zip_interface.go \
-package mock
mockgen -source=../interface/remote_interface.go -destination=internal_interfaces.go -package mock
mockgen -source=../interface/remote_use_interface.go -destination=remote_use_interfaces.go -package mock
mockgen -source=../interface/external_interface.go -destination=external_interfaces.go -package mock
