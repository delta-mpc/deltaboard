#!bin/bash
envsubst <"$ENV/config.tmpl"> "$ENV/config.yml";
./main migrate --conf config;
./main server --conf config;
