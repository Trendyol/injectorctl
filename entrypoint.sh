#!/usr/bin/env sh
echo "Adding new annotations to a file."
echo "Passing arguments is $@"
injectorctl inject -f $@
echo "Successfully added."