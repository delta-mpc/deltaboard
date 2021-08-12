#!bin/bash
echo $1;
if [ $1 and $1 = "init" ];then
python3 gen_config.py;
echo ".open /app/db/delta_dashboard.db" | sqlite3
else
python3 gen_web_config.py;
jupyterhub &
./main migrate --conf /application/config;
./main server --conf /application/config &
sh ./run_node.sh
fi