#!bin/bash
envsubst <"$ENV/config.tmpl"> "$ENV/config.yml";
./main $WORK --conf $ENV
