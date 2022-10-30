#!/bin/bash

mockgen -source=../../interface/remote_use_interface.go -destination=remote_use_interfaces.go -package mock
mockgen -source=../../interface/repo_zip_interface.go -destination=repo_zip_interfaces.go -package mock
