#!bin/bash
# fork ./run_jupyter.sh
envsubst <"$ENV/config.tmpl"> "$ENV/config.yml";
# ./main $WORK --conf $ENV