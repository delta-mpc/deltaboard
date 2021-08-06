#!bin/bash
jupyterhub &
envsubst <"config/config.tmpl"> "config/config.yml";
./main migrate --conf config &
./main server --conf config & 
sh run_node.sh